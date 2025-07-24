package entities

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var (
	items   = make(map[int]Item)
	itemsMu sync.Mutex
)

// /items (GET, POST)
func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAllItems(w, r)
	case http.MethodPost:
		createItem(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// /items/{id} (GET, PUT, DELETE)
func ItemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/items/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getItem(w, r, id)
	case http.MethodPut:
		updateItem(w, r, id)
	case http.MethodDelete:
		deleteItem(w, r, id)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	itemsMu.Lock()
	defer itemsMu.Unlock()
	var list []Item
	for _, i := range items {
		list = append(list, i)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func getItem(w http.ResponseWriter, r *http.Request, id int) {
	itemsMu.Lock()
	defer itemsMu.Unlock()
	i, ok := items[id]
	if !ok {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(i)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var i Item
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	itemsMu.Lock()
	i.ID = rand.Intn(1000000)
	items[i.ID] = i
	itemsMu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(i)
}

func updateItem(w http.ResponseWriter, r *http.Request, id int) {
	var i Item
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	itemsMu.Lock()
	defer itemsMu.Unlock()
	if _, ok := items[id]; !ok {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	i.ID = id
	items[id] = i
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(i)
}

func deleteItem(w http.ResponseWriter, r *http.Request, id int) {
	itemsMu.Lock()
	defer itemsMu.Unlock()
	if _, ok := items[id]; !ok {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	delete(items, id)
	w.WriteHeader(http.StatusNoContent)
} 
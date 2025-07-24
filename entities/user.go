package entities

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	users   = make(map[int]User)
	usersMu sync.Mutex
)

// /users (GET, POST)
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAllUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// /users/{id} (GET, PUT, DELETE)
func UserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getUser(w, r, id)
	case http.MethodPut:
		updateUser(w, r, id)
	case http.MethodDelete:
		deleteUser(w, r, id)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	usersMu.Lock()
	defer usersMu.Unlock()
	var list []User
	for _, u := range users {
		list = append(list, u)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func getUser(w http.ResponseWriter, r *http.Request, id int) {
	usersMu.Lock()
	defer usersMu.Unlock()
	u, ok := users[id]
	if !ok {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	usersMu.Lock()
	u.ID = rand.Intn(1000000)
	users[u.ID] = u
	usersMu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	usersMu.Lock()
	defer usersMu.Unlock()
	if _, ok := users[id]; !ok {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	u.ID = id
	users[id] = u
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	usersMu.Lock()
	defer usersMu.Unlock()
	if _, ok := users[id]; !ok {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	delete(users, id)
	w.WriteHeader(http.StatusNoContent)
} 
# CRUD em Go

Este projeto é um exemplo de CRUD (Create, Read, Update, Delete) simples em Go, utilizando a biblioteca padrão (`net/http`, `encoding/json`).

## Estrutura do Projeto

```
├── main.go              # Inicialização do servidor e rotas principais
├── entities/
│   ├── user.go          # Entidade e handlers de Usuário
│   └── item.go          # Entidade e handlers de Item
├── go.mod               # Go Modules
├── run.sh               # Script para rodar o servidor
```

## Como rodar

1. Certifique-se de ter o Go instalado (versão 1.16+).
2. Instale as dependências (se necessário):
   ```sh
   go mod tidy
   ```
3. Inicie o servidor:
   ```sh
   ./run.sh
   ```
   O servidor estará disponível em `http://localhost:8080`.

## Rotas disponíveis

### Usuários
- `GET    /users`         → Lista todos os usuários
- `POST   /users`         → Cria um novo usuário
- `GET    /users/{id}`    → Busca um usuário pelo ID
- `PUT    /users/{id}`    → Atualiza um usuário
- `DELETE /users/{id}`    → Remove um usuário

#### Exemplo de JSON para criar/atualizar usuário
```json
{
  "name": "João",
  "age": 30
}
```

### Itens
- `GET    /items`         → Lista todos os itens
- `POST   /items`         → Cria um novo item
- `GET    /items/{id}`    → Busca um item pelo ID
- `PUT    /items/{id}`    → Atualiza um item
- `DELETE /items/{id}`    → Remove um item

#### Exemplo de JSON para criar/atualizar item
```json
{
  "name": "Notebook",
  "price": 1999.99
}
```

## Observações
- O armazenamento é feito em memória (map), então os dados são perdidos ao reiniciar o servidor.
- O projeto está modularizado usando Go Modules e packages.

---

Sinta-se à vontade para adaptar, melhorar ou usar como base para outros projetos! 
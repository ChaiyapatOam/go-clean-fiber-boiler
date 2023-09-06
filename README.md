# Golang Clean Architecture Bolilorplate

## Built with 
- Go
- Fiber
- MySQL
- sqlx
## Run in Development
Recommend run with devcontainer (extension is vscode)

```bash
go run .
```

## Folder Structure

```md
.
├── 📂config/
│   └── config.go
├── 📂controller/
│   └── user.controller.go
├── 📂db/
│   └── mysql.go
├── 📂domain/
│   └── user.go
├── 📂internal/
│   ├── 📂payload
│   └── 📂validator
├── 📂repository/
│   └── user.repository.go
├── go.mod
├── go.sum
├── main.go
└── README.md
```

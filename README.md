# 🧠 Task Management HTTP Server (Go + REST + Swagger)

This project implements a lightweight HTTP server in Go that allows you to create, check the status, retrieve results, and delete tasks via a RESTful API. It uses in-memory storage and simulates asynchronous task processing. Swagger UI is available for easy API exploration.

---

## 🚀 Features

- ✅ Create or update tasks
- 🔄 Check task status by ID
- 📦 Retrieve task results
- ❌ Delete tasks
- 📑 Auto-generated Swagger documentation
- ⚙️ Configurable via YAML file

---

## 🔧 Tech Stack

- [Go (Golang)](https://go.dev/)
- [Chi](https://github.com/go-chi/chi) — HTTP router
- [Swaggo](https://github.com/swaggo/swag) — Swagger docs
- [cleanenv](https://github.com/ilyakaznacheev/cleanenv) — config loader
- [UUID](https://pkg.go.dev/github.com/google/uuid) — task ID generation

---

## 📂 Project Structure

```
http_server_arch/
├── api/http/                 # HTTP handlers with Swagger annotations
├── cmd/config/               # Flag parsing and config loading
├── docs/                     # Auto-generated Swagger docs
├── domain/                   # Domain models (Task)
├── main.go                   # App entry point
├── pkg/http/                 # Server creation & runner
├── repository/ram_storage/   # In-memory task storage
├── usecases/                 # Business logic layer
└── config.yml                # App config (example)
```

---

## ⚙️ Configuration

Create a file named `config.yml` in the project root:

```yaml
address: ":8080"
```

This sets the address/port where the server will listen.

---

## 🧪 API Endpoints

| Method | Endpoint           | Description                      |
|--------|--------------------|----------------------------------|
| POST   | `/task`            | Create a new task                |
| PUT    | `/task`            | Create or update a task          |
| GET    | `/task/status?id=` | Get task status by ID            |
| GET    | `/task/result?id=` | Get task result by ID            |
| DELETE | `/task/{id}`       | Delete a task by ID              |

---

## 📦 Build & Run

1. Install dependencies:
   ```bash
   go get ./...
   swag init  # Only needed if docs not yet generated
   ```

2. Build the project:
   ```bash
   go build -o server
   ```

3. Run the server:
   ```bash
   ./server -config=config.yml
   ```

---

## 🌐 Swagger UI

After starting the server, open this in your browser:

```
http://localhost:8080/swagger/index.html
```

You will see interactive API documentation generated from annotations in your Go code.

---

## 🧪 Example `curl` Request

```bash
curl -X POST http://localhost:8080/task \
-H "Content-Type: application/json" \
-d '{"code": "print(1+1)", "compiler": "python"}'
```

---

## 📝 Notes

- Tasks simulate asynchronous processing: 10 seconds after creation, status changes to `ready` and result becomes available.
- Storage is in-memory only (RAM). All tasks are lost when the server stops.
- The server is intended for demonstration, testing, or educational use.

---

## 🛠️ Future Improvements

- [ ] Replace RAM storage with persistent database (e.g., PostgreSQL)
- [ ] Add authentication (e.g., JWT)
- [ ] Support multiple languages/compilers with sandboxed execution
- [ ] Add graceful shutdown support

---

## 📄 License

MIT — free to use, modify, and distribute.

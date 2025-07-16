
---

# 📘 Task Management API Documentation

This API allows users to manage tasks through a RESTful interface. You can create, read, update, and delete tasks. It is built using Go and the Gin framework.

---

## 🔗 Published API Documentation

📄 View the full Postman documentation here:
👉 [https://documenter.getpostman.com/view/39808423/2sB34ijK67]

---

## 🛠️ How to Run the API

### 1. **Clone the repository**

```bash
git clone https://github.com/tsigemariamzewdu/Task-Management-REST-API-using-Go-and-Gin-Framework

```

### 2. **Install dependencies**

```bash
go mod tidy
```

### 3. **Run the server**

```bash
go run main.go
```

By default, the server runs on:

```
http://localhost:8080
```

---

## 📬 API Endpoints

### 📥 `GET /tasks`

**Description:** Get a list of all tasks.
**Response:**

```json
[
  {
    "id": 1,
    "title": "morning Plan",
    "description": "finish task 4",
    "dueDate": "2025-07-20T23:59:59Z",
    "Status": "in-progress"
  },
  ...
]
```

---

### 🔍 `GET /tasks/:id`

**Description:** Get a task by its ID.
**Params:**

* `id` (int): Task ID
  **Response:**

```json
{
  "id": 1,
  "title": "morning Plan",
  "description": "finish task 4",
  "dueDate": "2025-07-20T23:59:59Z",
  "Status": "in-progress"
}
```

---

### ➕ `POST /tasks`

**Description:** Create a new task.
**Request Body:**

```json
{
  "title": "New Task",
  "description": "Complete project",
  "dueDate": "2025-07-31T23:59:59Z",
  "Status": "not-started"
}
```

**Response:** Returns the created task with ID.

---

### 🔄 `PUT /tasks/:id`

**Description:** Update a task by ID.
**Params:**

* `id` (int): Task ID
  **Request Body:**

```json
{
  "title": "Updated Title",
  "description": "Updated Description",
  "dueDate": "2025-08-01T23:59:59Z",
  "Status": "completed"
}
```

**Response:** Updated task object.

---

### ❌ `DELETE /tasks/:id`

**Description:** Delete a task by ID.
**Params:**

* `id` (int): Task ID
  **Response:**

```json
{
  "message": "Task deleted successfully"
}
```

---

## 📎 Status Enum

The `Status` field must be one of the following values:

* `not-started`
* `in-progress`
* `completed`

---


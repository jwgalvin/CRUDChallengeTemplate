# CRUD API Starter Template (Go)

A starter framework for a CRUD REST API interview challenge. The boilerplate is set up; you'll implement the core logic during the interview.

## What's Pre-Built (Framework)

✅ **Ready to use:**
- Server startup and graceful shutdown (`cmd/server/main.go`)
- Route registration structure (`internal/api/server.go`)
- Data models (`internal/model/item.go`)
- Store interface definition (`internal/store/store.go`)
- Comprehensive documentation

## What You'll Build (During Interview)

❌ **Stubbed and waiting for implementation:**
- `internal/handlers/items.go` - CRUD request handlers (5 functions)
- `internal/httpjson/json.go` - Request/response JSON encoding (3 functions)
- `internal/validation/validation.go` - Input validation (2 functions)
- `internal/store/memory/memory.go` - In-memory data storage (5 methods)

Each file has detailed TODO comments explaining what to implement.

## Test Files (Reference & Validation)

The test files are included as **reference implementations** and to validate your work:

- `internal/handlers/items_test.go` - Examples of testing CRUD operations
- `internal/store/memory/memory_test.go` - Examples of testing the store layer
- `internal/validation/validation_test.go` - Examples of testing input validation

You can run tests to verify your implementations:
```bash
go test ./...
```

## Implementation Strategy

Recommended order during the interview:

1. **Start with `internal/validation/validation.go`** (2 functions)
   - `NormalizeItemInput()` - Validate item name and tags
   - `ParseListFilter()` - Parse query parameters
   - These are pure functions with no dependencies

2. **Then `internal/httpjson/json.go`** (3 functions)
   - `WriteJSON()` - Encode responses
   - `WriteError()` - Encode errors
   - `DecodeJSON()` - Decode requests
   - These handle all request/response I/O

3. **Then `internal/store/memory/memory.go`** (5 methods)
   - Store constructor (`New()`)
   - Core CRUD: `Create()`, `Get()`, `Update()`, `Delete()`
   - List with filtering: `List()`

4. **Finally `internal/handlers/items.go`** (5 functions)
   - Route dispatch: `HandleItems()`, `HandleItem()`
   - CRUD handlers: `listItems()`, `createItem()`, `getItem()`, `updateItem()`, `deleteItem()`
   - These orchestrate everything together

**Test as you go:**
```bash
go test ./internal/validation
go test ./internal/httpjson
go test ./internal/store/memory
go test ./internal/handlers
go test ./...
```

## Quick Build & Test

```bash
# View what's in the starter template
cat internal/handlers/items.go
cat internal/httpjson/json.go
cat internal/validation/validation.go
cat internal/store/memory/memory.go

# Once you implement the functions, test with:
go test ./...
PORT=8080 STORAGE=memory go run ./cmd/server
```

## Endpoints

- `GET /health`
- `GET /items?name=alpha&tag=one&limit=10&offset=0`
- `POST /items`
- `GET /items/{id}`
- `PUT /items/{id}`
- `DELETE /items/{id}`

## Example Requests

### Running the Server

```bash
# macOS/Linux
PORT=8080 STORAGE=memory go run ./cmd/server
```

```powershell
# Windows PowerShell
$env:PORT=8080; $env:STORAGE=memory; & go run ./cmd/server
```

---

## cURL Requests (macOS/Linux/Git Bash)

### Health Check
```bash
# Simple health check
curl -s http://localhost:8080/health | jq
# Response: {"status":"ok"}
```

### Create Item (POST)
```bash
# Create a new item
curl -s -X POST http://localhost:8080/items \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Go Backend API",
    "tags": ["golang", "backend", "rest"]
  }' | jq

# Response (example):
# {
#   "id": "item-1",
#   "name": "Go Backend API",
#   "tags": ["golang", "backend", "rest"],
#   "createdAt": "2026-02-17T10:30:45.123Z",
#   "updatedAt": "2026-02-17T10:30:45.123Z"
# }
```

### List Items (GET)
```bash
# List all items (default limit 50)
curl -s http://localhost:8080/items | jq

# List with pagination
curl -s "http://localhost:8080/items?limit=5&offset=0" | jq

# Filter by name (substring, case-insensitive)
curl -s "http://localhost:8080/items?name=backend" | jq

# Filter by tag (exact match, case-insensitive)
curl -s "http://localhost:8080/items?tag=golang" | jq

# Combine filters
curl -s "http://localhost:8080/items?name=go&tag=backend&limit=10" | jq

# Response (example):
# {
#   "items": [
#     {
#       "id": "item-1",
#       "name": "Go Backend API",
#       "tags": ["golang", "backend", "rest"],
#       "createdAt": "2026-02-17T10:30:45.123Z",
#       "updatedAt": "2026-02-17T10:30:45.123Z"
#     }
#   ],
#   "limit": 5,
#   "offset": 0,
#   "count": 1
# }
```

### Get Single Item (GET)
```bash
# Replace item-1 with actual ID from create response
curl -s http://localhost:8080/items/item-1 | jq

# Response (same as create response above)
```

### Update Item (PUT)
```bash
# Update an item
curl -s -X PUT http://localhost:8080/items/item-1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Go Backend API",
    "tags": ["golang", "backend", "v2"]
  }' | jq

# Response (id and createdAt unchanged, updatedAt updated):
# {
#   "id": "item-1",
#   "name": "Updated Go Backend API",
#   "tags": ["golang", "backend", "v2"],
#   "createdAt": "2026-02-17T10:30:45.123Z",
#   "updatedAt": "2026-02-17T10:31:50.456Z"
# }
```

### Delete Item (DELETE)
```bash
# Delete an item
curl -s -X DELETE http://localhost:8080/items/item-1

# Response: Empty (204 No Content)
# Verify it's deleted:
curl -s http://localhost:8080/items/item-1 | jq
# Response: {"error":"not found"}
```

### Error Scenarios
```bash
# Missing required name field
curl -s -X POST http://localhost:8080/items \
  -H "Content-Type: application/json" \
  -d '{"tags":["test"]}' | jq
# Response 400: {"error":"name is required"}

# Too many tags (max 20)
curl -s -X POST http://localhost:8080/items \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Too many tags",
    "tags": ['tag1", "tag2", ... "tag21"]
  }' | jq
# Response 400: {"error":"tags limit exceeded"}

# Invalid method
curl -s -X PATCH http://localhost:8080/items | jq
# Response 405: {"error":"method not allowed"}

# Item not found
curl -s http://localhost:8080/items/nonexistent | jq
# Response 404: {"error":"not found"}
```

---

## PowerShell Requests (Windows)

### Health Check
```powershell
Invoke-WebRequest -Uri http://localhost:8080/health | ConvertTo-Json
```

### Create Item (POST)
```powershell
$body = @{
    name = "Go Backend API"
    tags = @("golang", "backend", "rest")
} | ConvertTo-Json

Invoke-WebRequest -Uri http://localhost:8080/items `
  -Method POST `
  -ContentType "application/json" `
  -Body $body
```

### List Items (GET)
```powershell
# List all items
Invoke-WebRequest -Uri http://localhost:8080/items

# With pagination
Invoke-WebRequest -Uri "http://localhost:8080/items?limit=5&offset=0"

# Filter by name
Invoke-WebRequest -Uri "http://localhost:8080/items?name=backend"

# Filter by tag
Invoke-WebRequest -Uri "http://localhost:8080/items?tag=golang"

# Combined filters
Invoke-WebRequest -Uri "http://localhost:8080/items?name=go&tag=backend&limit=10"
```

### Get Single Item (GET)
```powershell
Invoke-WebRequest -Uri http://localhost:8080/items/item-1
```

### Update Item (PUT)
```powershell
$body = @{
    name = "Updated Go Backend API"
    tags = @("golang", "backend", "v2")
} | ConvertTo-Json

Invoke-WebRequest -Uri http://localhost:8080/items/item-1 `
  -Method PUT `
  -ContentType "application/json" `
  -Body $body
```

### Delete Item (DELETE)
```powershell
Invoke-WebRequest -Uri http://localhost:8080/items/item-1 `
  -Method DELETE

# Verify deletion:
Invoke-WebRequest -Uri http://localhost:8080/items/item-1
# Should return 404 error
```

### Error Scenarios
```powershell
# Missing required name field
$body = @{ tags = @("test") } | ConvertTo-Json
Invoke-WebRequest -Uri http://localhost:8080/items `
  -Method POST `
  -ContentType "application/json" `
  -Body $body
# Response 400: {"error":"name is required"}

# Item not found
Invoke-WebRequest -Uri http://localhost:8080/items/nonexistent
# Response 404: {"error":"not found"}
```

---

## Notes

- The default store is in-memory. Set `STORAGE=sqlite` to start the sqlite stub.
- Replace the module path in go.mod as needed.
- **Windows users:** Use PowerShell commands above, or install Git for Windows to use bash/curl commands.

# Go Event Management API

A RESTful API for managing events and attendees built with Go, Gin framework, and SQLite database.

## 🚀 Features

- **User Authentication**: JWT-based authentication system
- **Event Management**: Create, read, update, and delete events
- **Attendee Management**: Add/remove attendees to/from events
- **User Registration**: Secure user registration with password hashing
- **Database Migrations**: Automated database schema management
- **Authorization**: Role-based access control for event operations

## 🛠️ Tech Stack

- **Language**: Go 1.25.1
- **Web Framework**: Gin
- **Database**: SQLite
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Database Migrations**: golang-migrate

## 📦 Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/KhaledEemam/go-warm-up.git
   cd go-warm-up
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables** (optional)
   ```bash
   export PORT_NUMBER=8080
   export JWT_SECRET=your-secret-key
   ```

4. **Run database migrations**
   ```bash
   go run cmd/migrate/main.go up
   ```

5. **Start the server**
   ```bash
   go run cmd/api/main.go
   ```

The server will start on `http://localhost:8080` (or your specified port).

## 📊 Database Schema

The application uses three main tables:

- **users**: Store user information (id, email, name, password)
- **events**: Store event details (id, owner_id, name, description, date, location)
- **attendees**: Junction table for user-event relationships (id, user_id, event_id)

## 🔐 API Endpoints

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/events` | Get all events |
| GET | `/api/v1/events/:id` | Get event by ID |
| GET | `/api/v1/events/:id/attendees` | Get attendees for an event |
| GET | `/api/v1/attendees/:id/events` | Get events for an attendee |
| POST | `/api/v1/register` | Register a new user |
| POST | `/api/v1/login` | User login |

### Protected Endpoints (Require Authentication)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/events` | Create a new event |
| PUT | `/api/v1/events/:id` | Update an event (owner only) |
| DELETE | `/api/v1/events/:id` | Delete an event (owner only) |
| POST | `/api/v1/events/:id/attendees/:userId` | Add attendee to event (owner only) |
| DELETE | `/api/v1/events/:id/attendees/:userId` | Remove attendee from event (owner only) |

## 📝 API Usage Examples

### User Registration
```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### User Login
```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Create Event (Authenticated)
```bash
curl -X POST http://localhost:8080/api/v1/events \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Tech Conference 2024",
    "description": "Annual technology conference",
    "date": "2024-12-15",
    "location": "Convention Center"
  }'
```

### Get All Events
```bash
curl -X GET http://localhost:8080/api/v1/events
```

## 🏗️ Project Structure

```
├── cmd/
│   ├── api/                 # API server
│   │   ├── auth.go          # Authentication handlers
│   │   ├── context.go       # Context utilities
│   │   ├── events.go        # Event handlers
│   │   ├── main.go          # Application entry point
│   │   ├── middleware.go    # JWT middleware
│   │   ├── routes.go        # Route definitions
│   │   ├── server.go        # HTTP server setup
│   │   └── users.go         # User handlers
│   └── migrate/             # Database migrations
│       ├── main.go          # Migration runner
│       └── migrations/      # SQL migration files
├── internal/
│   ├── database/            # Database models and operations
│   │   ├── Attendees.go     # Attendee model
│   │   ├── events.go        # Event model
│   │   ├── main.go          # Database initialization
│   │   └── users.go         # User model
│   └── env/                 # Environment configuration
│       └── env.go           # Environment utilities
├── data.db                  # SQLite database file
├── go.mod                   # Go module definition
└── go.sum                   # Go module checksums
```

## 🔧 Configuration

The application supports the following environment variables:

- `PORT_NUMBER`: Server port (default: 8080)
- `JWT_SECRET`: Secret key for JWT signing (default: "new-secret-jwt")

## 🗄️ Database Migrations

To run database migrations:

```bash
# Apply all up migrations
go run cmd/migrate/main.go up

# Rollback all migrations
go run cmd/migrate/main.go down
```

## 🔒 Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

Tokens expire after 72 hours and contain the user ID for authorization purposes.

## 📋 Request/Response Examples

### Event Object
```json
{
  "id": 1,
  "ownerId": 1,
  "name": "Tech Conference 2024",
  "description": "Annual technology conference",
  "date": "2024-12-15",
  "location": "Convention Center"
}
```

### User Object
```json
{
  "id": 1,
  "email": "john@example.com",
  "name": "John Doe"
}
```

### Login Response
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

## 🛡️ Security Features

- Password hashing with bcrypt
- JWT-based stateless authentication
- Input validation and sanitization
- SQL injection prevention with parameterized queries
# Chronotes Template

This starter provides a tiny Go backend using GORM and Chi plus a React frontend. Google OAuth is used for authentication and a protected API returns the current user's profile.

## Quick start

1. Start Postgres and create a database.
2. Copy `backend/env.development.template` to `backend/.env` and adjust values including Google OAuth credentials.
3. From `backend`, run `go run ./cmd/server`.

Example request:

```bash
curl -v --cookie "session=1" http://localhost:8080/v1/profile
```

The server expects a `session` cookie containing a numeric user ID. If the user exists, it returns:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "jane@example.com",
    "name": "Jane Doe",
    "created_at": "2025-07-28T12:34:56Z"
  }
}
```

Logout with:

```bash
curl -X POST -v --cookie "session=1" http://localhost:8080/v1/auth/logout
```

See [backend](backend/README.md) and [frontend](frontend/README.md) for more details.

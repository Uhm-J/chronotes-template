# Backend

This backend uses Go with Chi, GORM and Google OAuth to authenticate users.

## Setup

1. Install Go 1.22 or newer.
2. Copy `env.development.template` to `.env` and fill in your database and Google OAuth credentials.
3. Run `go run ./cmd/server`.

The server automatically runs migrations for the `users` table and listens on the port configured in `.env`.

## OAuth

Two endpoints handle authentication:

- `GET /v1/auth/google/login` starts the Google sign‑in flow.
- `GET /v1/auth/google/callback` processes the OAuth callback and creates the user if necessary.

After logging in a `session` cookie is set. Use it when calling the protected `GET /v1/profile` endpoint.

To log out, send a `POST /v1/auth/logout` request which clears the cookie.

A health check is available at `GET /v1/health`.

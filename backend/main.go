package main

// A simplified backend for the Chronotes template site.
//
// This server demonstrates how to implement Google OAuth2 login, a
// minimal SQLite-backed user table and a few HTTP endpoints.  It uses
// the pure Go SQLite driver (modernc.org/sqlite) so it can build
// without CGO and should run out of the box.  The code intentionally
// omits robust error handling and security features such as CSRF
// protection or state verification to keep the example concise.  See
// the original Chronotes repository for a production-ready
// implementation.

import (
    "context"
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    _ "modernc.org/sqlite"
)

// User represents a minimal user record stored in the database.
type User struct {
    ID    int64  `json:"id"`
    Email string `json:"email"`
    Name  string `json:"name"`
}

var (
    db          *sql.DB
    oauthConfig *oauth2.Config
)

// initDB opens the SQLite database file and ensures the users table exists.
func initDB() error {
    var err error
    // Connect to a local SQLite database. The query parameter "cache=shared"
    // allows multiple connections to share the same cache.
    db, err = sql.Open("sqlite", "file:users.db?cache=shared")
    if err != nil {
        return err
    }
    // Create the users table if it does not already exist.  In a
    // production system you would likely manage migrations via a
    // migration tool.
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT NOT NULL UNIQUE,
        name TEXT
    )`)
    return err
}

// handleGoogleLogin begins the OAuth2 flow by redirecting the user
// to Google's authorization page.  The AuthCodeURL call generates
// a URL with the appropriate client ID, scope and redirect URI.
func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
    url := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// handleGoogleCallback receives the authorization code from Google,
// exchanges it for an access token and retrieves the user's profile.
// It then inserts or updates the user in the database and sets a
// simple session cookie containing the user ID.
func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
    // In a real implementation you should verify the provided state
    // parameter to mitigate CSRF attacks.  This example omits state
    // verification for brevity.
    code := r.FormValue("code")
    if code == "" {
        http.Error(w, "code not provided", http.StatusBadRequest)
        return
    }
    // Exchange the authorization code for an access token.
    token, err := oauthConfig.Exchange(context.Background(), code)
    if err != nil {
        http.Error(w, "failed to exchange token: "+err.Error(), http.StatusInternalServerError)
        return
    }
    // Create an HTTP client that automatically attaches the bearer token.
    client := oauthConfig.Client(context.Background(), token)
    // Request the user's profile information from Google.  The
    // `userinfo` endpoint is part of the OpenID Connect standard.
    resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
    if err != nil {
        http.Error(w, "failed to get user info: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        http.Error(w, "unexpected status from userinfo endpoint", http.StatusInternalServerError)
        return
    }
    // Decode the JSON response into a temporary struct.  We only
    // extract the fields we need.
    var info struct {
        Email string `json:"email"`
        Name  string `json:"name"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
        http.Error(w, "failed to decode user info: "+err.Error(), http.StatusInternalServerError)
        return
    }
    // Insert or update the user in the database.  If the user does
    // not exist we insert a new row; otherwise we update the name.
    var userID int64
    err = db.QueryRow("SELECT id FROM users WHERE email = ?", info.Email).Scan(&userID)
    if err == sql.ErrNoRows {
        res, err := db.Exec("INSERT INTO users (email, name) VALUES (?, ?)", info.Email, info.Name)
        if err != nil {
            http.Error(w, "failed to insert user: "+err.Error(), http.StatusInternalServerError)
            return
        }
        userID, _ = res.LastInsertId()
    } else if err != nil {
        http.Error(w, "failed to query user: "+err.Error(), http.StatusInternalServerError)
        return
    } else {
        _, err = db.Exec("UPDATE users SET name = ? WHERE id = ?", info.Name, userID)
        if err != nil {
            http.Error(w, "failed to update user: "+err.Error(), http.StatusInternalServerError)
            return
        }
    }
    // Set a cookie with the user ID so the frontend can fetch the
    // current user on subsequent requests.  In a real application you
    // would issue a signed JWT or store a secure session token.
    http.SetCookie(w, &http.Cookie{
        Name:     "session",
        Value:    fmt.Sprintf("%d", userID),
        Path:     "/",
        HttpOnly: true,
        // For demonstration we leave Secure off; enable it in production when
        // serving over HTTPS.
    })
    // Redirect back to the root path.  The frontend will detect
    // whether the user is logged in by calling `/v1/auth/me`.
    http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// handleCurrentUser reads the session cookie, looks up the user in
// the database and returns their record as JSON.  If no valid
// session cookie is found the request returns 401 Unauthorized.
func handleCurrentUser(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session")
    if err != nil || cookie.Value == "" {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
    var user User
    err = db.QueryRow("SELECT id, email, name FROM users WHERE id = ?", cookie.Value).Scan(&user.ID, &user.Email, &user.Name)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func main() {
    // Initialise the database connection and schema.
    if err := initDB(); err != nil {
        log.Fatalf("failed to initialise database: %v", err)
    }
    // Read OAuth configuration from the environment.  These must be
    // provided by whoever deploys the application.  See
    // logbook/GOOGLE_OAUTH_SETUP.md in the original Chronotes repo for
    // instructions on obtaining client credentials.
    clientID := os.Getenv("GOOGLE_CLIENT_ID")
    clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
    redirectURL := os.Getenv("GOOGLE_REDIRECT_URL")
    if clientID == "" || clientSecret == "" || redirectURL == "" {
        log.Println("Warning: Google OAuth credentials not set. Login will fail.")
    }
    oauthConfig = &oauth2.Config{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        RedirectURL:  redirectURL,
        Scopes: []string{
            "https://www.googleapis.com/auth/userinfo.profile",
            "https://www.googleapis.com/auth/userinfo.email",
        },
        Endpoint: google.Endpoint,
    }
    // Register handlers for the authentication endpoints and current
    // user lookup.  We prefix auth endpoints with /v1 to match the
    // Chronotes API style.
    http.HandleFunc("/v1/auth/google/login", handleGoogleLogin)
    http.HandleFunc("/v1/auth/google/callback", handleGoogleCallback)
    http.HandleFunc("/v1/auth/me", handleCurrentUser)
    // Serve the frontend from the ./frontend/dist directory if it exists.
    // Running `npm run build` or `yarn build` in the frontend will
    // produce this directory via Vite.  If it does not exist we
    // simply return a 404 for static file requests.
    if _, err := os.Stat("frontend/dist"); err == nil {
        fs := http.FileServer(http.Dir("frontend/dist"))
        http.Handle("/", fs)
    } else {
        // Fallback handler to prompt the developer to build the frontend.
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprintln(w, "frontend not built – run `npm run build` in the frontend directory")
        })
    }
    // Determine the port to listen on.  Default to 8080 if PORT is
    // unset.
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("listening on :%s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

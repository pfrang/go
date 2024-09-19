package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Message struct {
	Content string `json:"content"`
}

func main() {
	log.Println("Starting server...")

	// Open the SQLite database file
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table if it doesn't exist
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS messages (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        content TEXT
    );
    `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	// Set up the router
	r := mux.NewRouter()

	// Serve static files
	// Handle API requests with CORS middleware
	r.HandleFunc("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		handleMessages(w, r, db)
	}).Methods("GET", "POST", "OPTIONS")

	// Apply CORS middleware
	r.Use(corsMiddleware)

	log.Println("Server is ready to accept connections on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Handle HTTP requests for messages
func handleMessages(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	log.Println("Handling request for messages")
	switch r.Method {
	case "GET":
		log.Println("Handling GET request")
		// Fetch messages from the database
		rows, err := db.Query("SELECT content FROM messages")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var messages []string
		for rows.Next() {
			var content string
			if err := rows.Scan(&content); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			messages = append(messages, content)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)

	case "POST":
		log.Println("Handling POST request")
		// Insert a new message into the database
		var msg Message
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Send the message to OpenAI and get the response
		reply, err := SendMessageToOpenAI(msg.Content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Insert the OpenAI reply into the database
		_, err = db.Exec("INSERT INTO messages(content) VALUES(?)", reply)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

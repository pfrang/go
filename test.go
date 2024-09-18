package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
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

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(w, r, db)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handle WebSocket connections
func handleWebSocket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		// Read a message from the client
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// Print the message to the server console
		fmt.Printf("Received: %s\n", message)

		// Insert the message into the database
		_, err = db.Exec("INSERT INTO messages(content) VALUES(?)", string(message))
		if err != nil {
			log.Println(err)
			return
		}

		// Send a confirmation message back to the client
		confirmationMessage := fmt.Sprintf("Message '%s' has been stored in the database.", message)
		err = conn.WriteMessage(websocket.TextMessage, []byte(confirmationMessage))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

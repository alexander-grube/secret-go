package main

import (
	"alexander-grube/secret-go/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
)

type Handlers struct {
	Queries *db.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	handlers := &Handlers{
		Queries: queries,
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.initRouter()))
}

func (h *Handlers) initRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", h.helloWorld)
	mux.HandleFunc("POST /secret", h.createSecret)
	mux.HandleFunc("GET /secret/{id}", h.getSecret)
	return mux
}

func (h *Handlers) helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func (h *Handlers) createSecret(w http.ResponseWriter, r *http.Request) {
	var secret PostSecretMessage
	err := json.NewDecoder(r.Body).Decode(&secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	message, err := h.Queries.CreateSecret(context.Background(), secret.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func (h *Handlers) getSecret(w http.ResponseWriter, r *http.Request) {
	uuid := pgtype.UUID{}
	err := uuid.Scan(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	message, err := h.Queries.GetSecret(context.Background(), uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.Queries.DeleteSecret(context.Background(), uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// secret message struct
type PostSecretMessage struct {
	Message string `json:"message"`
}

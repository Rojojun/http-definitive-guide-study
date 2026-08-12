package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	mu          sync.RWMutex
	idempotancy = make(map[string]string)
	order       = make(map[string]string)
)

func main() {
	mux := http.NewServeMux()

	start()

	mux.HandleFunc("GET /", mainMessage)
	mux.HandleFunc("POST /create", create)
	mux.HandleFunc("GET /find", findAll)

	http.ListenAndServe(":8080", mux)
}

func mainMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func create(w http.ResponseWriter, r *http.Request) {
	idempotancyKey := r.Header.Get("Idempotancy-Key")
	if idempotancyKey == "" {
		http.Error(w, "Idempotancy-Key is required", http.StatusBadRequest)
		return
	}

	mu.RLock()
	existing, exists := idempotancy[idempotancyKey]
	mu.RUnlock()

	if exists {
		fmt.Fprintln(w, existing)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	reqStr := string(body)

	mu.Lock()
	if existing, exists := idempotancy[idempotancyKey]; exists {
		mu.Unlock()
		fmt.Fprintln(w, existing)
		return
	}

	order[idempotancyKey] = reqStr
	response := fmt.Sprintf("CREATED : %s", reqStr)
	idempotancy[idempotancyKey] = response
	mu.Unlock()

	fmt.Fprintln(w, response)
}

func findAll(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	if len(idempotancy) == 0 {
		fmt.Fprintln(w, "empty")
		return
	}

	for k, v := range idempotancy {
		fmt.Fprintf(w, "key :%s -> value : %s\n", k, v)
	}
}

func idempotancyFind(key string) string {
	return idempotancy[key]
}

func start() {
	fmt.Println("================")
	fmt.Println("| Server Start |")
	fmt.Println("| port : 8080  |")
	fmt.Println("================")
}

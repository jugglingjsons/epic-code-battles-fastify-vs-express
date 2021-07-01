package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Answer struct {
	Hello string
}

func test(w http.ResponseWriter, req *http.Request) {
	data := Answer{Hello: "world"}
	w.Header().Set("Content-Type", "application/json")

	// timeoutString := os.Getenv("TIMEOUT")
	// timeout, _ := strconv.Atoi(timeoutString)
	// time.Sleep(time.Duration(timeout) * time.Millisecond)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println(os.Getenv("TIMEOUT"), os.Getenv("PORT"))

	http.HandleFunc("/", test)

	http.ListenAndServe(":"+os.Getenv("PORT"), logRequest(http.DefaultServeMux))
}

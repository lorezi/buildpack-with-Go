package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	addr := ":5001"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		res := map[string]interface{}{
			"hello": "world",
		}
		if err := enc.Encode(res); err != nil {
			panic("unable to encode response")

		}
	})
	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))

}

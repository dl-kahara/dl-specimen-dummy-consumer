package main

import (
	"crypto/sha256"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func consume(config Config) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Allow", "OPTIONS, POST")
			w.WriteHeader(204) // "No Content"
			return
		}
		if r.Method != "POST" {
			w.Header().Set("Allow", "OPTIONS, POST")
			w.WriteHeader(405) // "Method Not Allowed"
			return
		}
		if r.Header.Get("Content-Type") != "application/json" {
			w.Header().Set("Accept-Post", "application/json")
			w.WriteHeader(415) // "Unsupported Media Type"
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Err(err).Msg("Failed to read request body")
			w.WriteHeader(400) // "Bad Request"
			return
		}

		var payload interface{}
		if err := json.Unmarshal(body, &payload); err != nil {
			log.Err(err).Bytes("body", body).Msg("Failed to parse payload")
			w.WriteHeader(400) // "Bad Request"
			return
		}

		log.Debug().Any("payload", payload).Msg("Payload parsed")

		requests_consumed.WithLabelValues().Inc()
		bytes_consumed.WithLabelValues().Add(float64(len(body)))

		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(201) // "Created"
		checksum := sha256.Sum256(body)
		w.Write(checksum[:])
	})

	log.Fatal().Err(http.ListenAndServe(":8080", nil))
}

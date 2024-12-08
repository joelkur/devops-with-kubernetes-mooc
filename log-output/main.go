package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func getOutputWithTimestamp(s string) string {
	timestamp := time.Now().Format(time.RFC3339)
	return fmt.Sprintf("%s: %s", timestamp, s)
}

func main() {
	randomString := uuid.New().String()
	output := getOutputWithTimestamp(randomString)

	go func() {
		for {
			output = getOutputWithTimestamp(randomString)
			fmt.Println(output)
			time.Sleep(5 * time.Second)
		}
	}()

	port, portFromEnv := os.LookupEnv("PORT")
	if !portFromEnv {
		port = "3001"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", output)
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

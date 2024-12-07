package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	randomString := uuid.New()
	for {
		fmt.Printf("%s: %s\n", time.Now().Format(time.RFC3339), randomString.String())
		time.Sleep(5 * time.Second)
	}
}

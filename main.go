package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/morristai/rarbg-notifier/discord"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		http.HandleFunc("/", hello)
		http.ListenAndServe(":8090", nil)
	}()

	go func() {
		defer wg.Done()
		go discord.Run()
	}()

	wg.Wait()
}

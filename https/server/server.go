package main

import "net/http"
import "fmt"

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi this is golang server")
}
func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServeTLS(":8081", "../crt/server.crt",
		"../crt/server.key", nil)
}

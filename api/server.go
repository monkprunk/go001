package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
	hello := []byte("{\"name\": \"MonK\"}")
	w.Write(hello)
}

func mongHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("MonG")
	mong := []byte("{\"name\": \"MonG\"}")
	w.Write(mong)
}

func main() {
	fmt.Println("Start")

	http.HandleFunc("/mong", mongHandler)
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":1234", nil) //:port
	log.Fatal(err)
	fmt.Println("End")

}

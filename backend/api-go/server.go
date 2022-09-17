package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/configMap", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, %s! You are %s years old.", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	
	data, err := ioutil.ReadFile("/app/myfamily/family.txt");

	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	fmt.Fprintf(w, "My family: %s", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "User: %s | Password: %s", user, password)
}
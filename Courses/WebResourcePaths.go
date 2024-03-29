package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello Web !")
}

func spanishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hola web !")
}

func polishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "cześć")
}

func main() {

	http.HandleFunc("/eng", englishHandler)
	http.HandleFunc("/es", spanishHandler)
	http.HandleFunc("/pl", polishHandler)
	err := http.ListenAndServe("localhost:8080", nil) // "nil" to indicate that the requests will be handle by "HandleFunc"
	log.Fatal(err)

}

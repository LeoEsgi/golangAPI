package main

import "net/http"

type Attraction struct {
	id           int32
	name         string
	inPark       string
	manufacturer string
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "Hello world !"}`))
}

func handlerAttraction(w http.ResponseWriter, r *http.Request) {
	var attr = Attraction{1, "Montagne russe", "Miami", "Disney"}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "` + attr.name + ` !"}`))
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/attraction", handlerAttraction)
	http.ListenAndServe(":8001", nil)
}

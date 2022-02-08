package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type Attraction struct {
	Id           int
	Name         string
	InPark       string
	Manufacturer string
}

func delete(id string, attr []Attraction) {

	fmt.Println(id)
	var int_id, _ = strconv.Atoi(id)
	fmt.Println("Tableau de base :", attr)

	for i := 0; i < len(attr); i++ {
		if attr[i].Id == int_id {

			attr = append(attr[:i], attr[(i+1):]...)

		}

	}

	fmt.Println("Tableau delete : ", attr)

}

func patch(id string, attr []Attraction, name string, inPark string, manufacturer string) {
	fmt.Println(attr)
	var int_id, _ = strconv.Atoi(id)
	for i := 0; i < len(attr); i++ {
		fmt.Println(int_id)
		if attr[i].Id == int_id {

			attr[i].Name = name
			attr[i].InPark = inPark
			attr[i].Manufacturer = manufacturer
			fmt.Println("Attraction updated : ", attr[i])
			break
		}

	}
	fmt.Println(attr)
}

func get(attr []Attraction) {
	fmt.Print(attr)

}

func exist(attr []Attraction, id int) bool {
	for i := 0; i < len(attr); i++ {

		if attr[i].Id == id {

			return true

		}

	}
	return false
}

func create(attr []Attraction, name string, inPark string, manufacturer string) {
	var id int
	for {
		id = rand.Intn(100)
		if exist(attr, id) {
			continue

		}
		break
	}
	var attraction = Attraction{id, name, inPark, manufacturer}
	attr = append(attr, attraction)
	fmt.Println(attr)

}

func encode() {
	var id = rand.Intn(100)
	s1 := &Attraction{id, "lÉO", "AD", "be"}
	bytes_s1, err := json.Marshal(s1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes_s1))
	decode(string(bytes_s1))
}

func decode(s1 string) {
	var stu = &Attraction{}
	var err = json.Unmarshal([]byte(s1), stu)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal: Id : %d Name: %s, InPark: %s , Manufacturer : %s \n", stu.Id, stu.Name, stu.InPark, stu.Manufacturer)
}

func router() {

	var size = 2
	attractions := make([]Attraction, size)
	attractions[0] = Attraction{1, "Léo", "Blob", "GB"}
	attractions[1] = Attraction{2, "Nacer", "Blob", "GB"}

	http.HandleFunc("/", home)
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		get(attractions)
	})
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		var id = r.FormValue("id")
		delete(id, attractions)
	})
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		var name = r.FormValue("Name")
		var inPark = r.FormValue("InPark")
		var manufacturer = r.FormValue("Manufacturer")
		create(attractions, name, inPark, manufacturer)
	})

	http.HandleFunc("/patch", func(w http.ResponseWriter, r *http.Request) {
		var name = r.FormValue("Name")
		var inPark = r.FormValue("InPark")
		var manufacturer = r.FormValue("Manufacturer")
		var id = r.FormValue("Id")
		patch(id, attractions, name, inPark, manufacturer)

		//	patch(id, attractions, name, inPark, manufacturer)

	})
	http.ListenAndServe(":8001", nil)
}

func main() {

	router()

}

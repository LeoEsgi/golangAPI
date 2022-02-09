package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

type Attraction struct {
	Id           int    `json:"Id"`
	Name         string `json:"Name"`
	InPark       string `json:"InPark"`
	Place        string `json:"Place"`
	Manufacturer string `json:"Manufacturer"`
}

func saveInfile(attr []Attraction) {

	for i := 0; i < len(attr); i++ {
		fmt.Print("hello")
		data := Attraction{

			Id:           attr[i].Id,
			Name:         attr[i].Name,
			InPark:       attr[i].InPark,
			Place:        attr[i].Place,
			Manufacturer: attr[i].Manufacturer,
		}

		file, _ := json.MarshalIndent(data, "", " ")

		_ = ioutil.WriteFile("test.json", file, 0644)

	}
}

func delete(id string, attr []Attraction) {

	var int_id, _ = strconv.Atoi(id)
	fmt.Println("Tableau de base :", attr)

	for i := 0; i < len(attr); i++ {
		if attr[i].Id == int_id {

			attr = append(attr[:i], attr[(i+1):]...)

		}

	}

	fmt.Println("Tableau delete : ", attr)
}

func patch(id string, attr []Attraction, name string, inPark string, place string, manufacturer string) []Attraction {
	var int_id, _ = strconv.Atoi(id)
	for i := 0; i < len(attr); i++ {
		fmt.Println(int_id)
		if attr[i].Id == int_id {

			if len(name) != 0 {
				attr[i].Name = name
			}
			if len(inPark) != 0 {
				attr[i].InPark = inPark
			}
			if len(manufacturer) != 0 {
				attr[i].Manufacturer = manufacturer
			}
			if len(place) != 0 {
				attr[i].Place = place
			}
			fmt.Println("Attraction updated : ", attr[i])

			break
		}

	}

	return attr
}

func exist(attr []Attraction, id int) bool {
	for i := 0; i < len(attr); i++ {

		if attr[i].Id == id {

			return true

		}

	}
	return false
}

func create(attr []Attraction, name string, inPark string, place string, manufacturer string) []Attraction {
	var id int
	for {
		id = rand.Intn(100000000000)
		if exist(attr, id) {
			continue

		}
		break
	}
	var attraction = Attraction{id, name, inPark, place, manufacturer}
	json := encode(attraction)
	attraction = decode(json)
	attr = append(attr, attraction)

	return attr

}

func encode(s1 Attraction) []byte {

	bytes_s1, err := json.Marshal(s1)
	if err != nil {
		panic(err)
	}
	return bytes_s1
}

func decode(s1 []byte) Attraction {

	var stu = &Attraction{}
	var err = json.Unmarshal(s1, stu)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Unmarshal: Name: %s, InPark: %s ,Place : %s  ,  Manufacturer : %s \n", stu.Name, stu.InPark, stu.Place, stu.Manufacturer)
	return *stu
}

func router() {

	var size = 0
	attractions := make([]Attraction, size)
	http.Handle("/", http.FileServer(http.Dir("../static")))

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(attractions)

	})
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		var id = r.FormValue("Id")
		delete(id, attractions)
	})
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("creation en cours ... : ")
		name := r.FormValue("Name")
		place := r.FormValue("Place")
		manufacturer := r.FormValue("Manufacturer")
		inPark := r.FormValue("InPark")

		attractions = create(attractions, name, inPark, place, manufacturer)
		fmt.Println(attractions)

	})

	http.HandleFunc("/saveInfile", func(w http.ResponseWriter, r *http.Request) {
		saveInfile(attractions)

	})

	http.HandleFunc("/patch", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("update en cours ... : ")
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		place := r.FormValue("Place")
		manufacturer := r.FormValue("Manufacturer")
		inPark := r.FormValue("InPark")

		patch(id, attractions, name, inPark, place, manufacturer)
		fmt.Println(attractions)

		//	patch(id, attractions, name, inPark, manufacturer)

	})
	http.ListenAndServe(":8001", nil)
}

func main() {

	router()

}

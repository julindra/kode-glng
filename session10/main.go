package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"path"
	"runtime"
	"text/template"
	"time"

	"session10/models"
)

func autoUpdate(status *models.Status) {
	for range time.Tick(time.Second * 15) {
		status.Status.Water = rand.Intn(99) + 1
		status.Status.Wind = rand.Intn(99) + 1

		b, err := json.Marshal(status)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile("data.json", b, 0755)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	file, _ := ioutil.ReadFile("data.json")
	var status models.Status
	json.Unmarshal([]byte(file), &status)

	go autoUpdate(&status)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening on PORT 8080")
	http.ListenAndServe(":8080", nil)
}

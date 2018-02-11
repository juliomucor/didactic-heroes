package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Hero defines a hero
type Hero struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var heroes []Hero

// define handler to get all heroes
func getHeroHandler(w http.ResponseWriter, r *http.Request) {
	heroListBytes, err := json.Marshal(heroes)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(heroListBytes)
}

// define handler to create a new entry of heroes
func createHeroHandler(w http.ResponseWriter, r *http.Request) {
	// create a new instance of hero
	hero := Hero{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hero.ID, err = strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hero.Name = r.Form.Get("name")
	heroes = append(heroes, hero)

	http.Redirect(w, r, assets, http.StatusFound)
}

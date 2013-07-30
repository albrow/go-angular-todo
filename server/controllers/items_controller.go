package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stephenalexbrowne/go-angular-todo/server/models"
	"github.com/stephenalexbrowne/zoom"
	"net/http"
)

type ItemsController struct{}

func (*ItemsController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Called Index()")

	items, err := models.FindAllItems()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	itemsJson, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, string(itemsJson))
}

func (*ItemsController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Called Create()")

	decoder := json.NewDecoder(r.Body)
	var item *models.Item
	err := decoder.Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	i := models.NewItem(item.Content)
	err = zoom.Save(i)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	itemJson, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, string(itemJson))
}

func (*ItemsController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("called Update()")

	// get the Id from the url muxer
	vars := mux.Vars(r)
	itemId := vars["id"]
	if itemId == "" {
		http.Error(w, "Missing required paramater: id", 400)
		return
	}

	// get the model
	item, err := models.FindItemById(itemId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	content, ok := r.PostForm["content"]
	if ok && content[0] != "" && content[0] != "null" {
		item.Content = content[0]
	}

	err = zoom.Save(item)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	itemJson, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, string(itemJson))
}

func (*ItemsController) Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Called Show()")

	// get the Id from the url muxer
	vars := mux.Vars(r)
	itemId := vars["id"]
	if itemId == "" {
		http.Error(w, "Missing required paramater: Id", 400)
		return
	}

	// get the model
	i, err := models.FindItemById(itemId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	itemJson, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, string(itemJson))
}

func (*ItemsController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Called Delete()")

	// get the Id from the url muxer
	vars := mux.Vars(r)
	itemId := vars["id"]
	if itemId == "" {
		http.Error(w, "Missing required paramater: Id", 400)
		return
	}

	err := zoom.DeleteById("item", itemId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, `{"status":"OK"}`)
}

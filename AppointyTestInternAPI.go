package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

// The Objects In The list
type content struct{
	ID int64 `json:"id"`
	value string `json:"value"`
	completed string `json:"completed"`
}

// To DO List object
type toDoList struct{
	ID      int64  `json:"id"`
	name    string `json:"name"`
	items   content `json:"items"`
}


// this function handles the request and diverts to corresponding job/methods
func list(w http.ResponseWriter, r *http.Request) {
	 // If method is post create a new list as well as create new items if needed
	  if(r.Method=="POST"){
		id, _ := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
		itemsid, _ :=strconv.ParseInt(r.PostFormValue("items.id"), 10, 64)
		// items given
		itemsof :=content{
			ID : itemsid,
			value: r.PostFormValue("items.value"),
			completed: r.PostFormValue("items.completed"),
		}

		//list given
		todo := &toDoList{
			ID: id,
			name: r.PostFormValue("name"),
			items: itemsof,
		}

		// Insert the list into the LIST table
		query := "INSERT INTO public.list(id, name) VALUES($1,$2)"
		err := db.QueryRow(query, todo.ID, todo.name)
		if err!=nil{
			panic(err)
		}

		// Insert the items into the ITEM table
		query = "INSERT INTO public.item(listId,id, value,complete) VALUES($1,$2,$3,$4)"
		err = db.QueryRow(query,id,itemsof.ID, itemsof.value,itemsof.completed)
		if err!=nil{
			panic(err)
		}
		json.NewEncoder(w).Encode(todo)
		return 
	  }

	  // If method is delete then delete list with given list id
	  if(r.Method=="DELETE"){
			  // If you are going to delete list in List table first delete all the items corresponding to that list
			
			id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
			 query := "DELETE from public.item where listid=$1"
			_, err = db.Exec(query,id)
			if err != nil {
			panic(err)
			}

			
			query = "DELETE from public.list where id=$1"
			_, err = db.Exec(query,id)
			if err != nil {
				panic(err)
			}
			return

	  }

	  // if the method is patch update the name of the list with given list id
	  if(r.Method=="PATCH"){
			id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
			name:= r.PostFormValue("name")
			query := "UPDATE public.list set name=$2 where id=$1"
			_, err = db.Exec(query,name,id)
			if err != nil {
				panic(err)
			}
			return
	  }
	  http.Error(w, "404 not found.", http.StatusNotFound)
	  return
}

// Adds the item in the given List
func add(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	id, _ := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	itemsid, _ := strconv.ParseInt(r.PostFormValue("item.id"), 10, 64)

	// items given
	itemsof :=content{
		ID : itemsid,
		value: r.PostFormValue("items.value"),
		completed : r.PostFormValue("items.completed"),
	}

	//list given
	todo := &toDoList{
		ID: id,
		items: itemsof,
	}
	//fmt.Println(user)
	query := "INSERT INTO public.item(id, itemId,value,completed) VALUES($1, $2, $3, $4)"
	err = db.QueryRow(query, todo.ID, itemsof.ID, itemsof.value,itemsof.completed).Scan(&itemsof.ID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(todo)
}

// Delets the item from the given List
func delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	query := "DELETE from public.item where id=$1"
	_, err = db.Exec(query,id)
	if err != nil {
		panic(err)
	}
	//json.NewEncoder(w).Encode()	
}


// Gets the from the given List
func get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	id := r.URL.Query().Get("id")
	//fmt.Println(id)
	query := "SELECT id,value,completed FROM public.item WHERE id = $1"
	row := db.QueryRow(query, id)
	contents := content{}
	row.Scan(&contents.ID, &contents.value, &contents.completed)
	json.NewEncoder(w).Encode(contents)
}

// updates the item of the Given id
func update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	contents := &content{
		ID:      id,
		value:   r.PostFormValue("value"),
		completed: r.PostFormValue("completed"),
	}
	//fmt.Println(user)
	query := "UPDATE public.item SET completed=$3,value=$2 WHERE id = $1;"
	_, err = db.Exec(query, contents.ID, contents.value,contents.completed)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(contents)
}


func main() {
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=appointy dbname=test sslmode=disable")
	if err != nil {
		fmt.Println("database connection err")
		panic(err)
	}


	http.HandleFunc("/todolist", list)
	http.HandleFunc("/todolist:addItem", add)
	http.HandleFunc("/todolist:deletItem", delete)
	http.HandleFunc("/todolist:getItem", get)
	http.HandleFunc("/todolist:updateItem", update)
	http.ListenAndServe(":8080", nil)
}

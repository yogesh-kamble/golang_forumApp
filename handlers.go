package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Forum_Data struct {
    email string `json:"email"`
    title string `json:"title"`
    message string `json:"message"`
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    /*todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }*/
    c := conn.C("forum")
    error := c.Find(nil).All(&forums)
    if error != nil {
        log.Fatal(error)
    }
    if err := json.NewEncoder(w).Encode(forums); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}

func createForum(w http.ResponseWriter, r *http.Request){
    forum_data := Forum_Data{}
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&forum_data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
    } 
    //email := forum_data.email
    //title := forum_data.title
    //message := forum_data.message

    c := conn.C("forum")
    err = c.Insert(&Forum{forum_data.email, forum_data.title, forum_data.message})
    
    if err != nil {
        log.Fatal(err)
    }
    
    var m map[string]string
    m = make(map[string]string)
    m["message"] = "forum added succesfully"
    if err := json.NewEncoder(w).Encode(m); err != nil {
        panic(err)
    }
}


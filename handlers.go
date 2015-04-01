package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type Forum_Data struct {
    email string `json:"email"`
    title string `json:"title"`
    message string `json:"message"`
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func getForums(w http.ResponseWriter, r *http.Request) {
    /*todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }*/
    //w.Header().Set("Access-Control-Allow-Origin", "*")
    //w.Header().Set("Content-Type", "application/json")
    c := conn.C("forum")
    error := c.Find(nil).All(&forums)
    if error != nil {
        log.Fatal(error)
    }
    if err := json.NewEncoder(w).Encode(forums); err != nil {
        panic(err)
    }
}

func getForum(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    forumId := vars["forumId"]
    c := conn.C("forum")
    //forum := Forum{}
    err := c.FindId(bson.ObjectIdHex(forumId)).One(&forum)
    if err != nil{
       log.Fatal(err)
    }
    if err = json.NewEncoder(w).Encode(forum); err != nil {
        panic(err)
    }
}

func createForum(w http.ResponseWriter, r *http.Request){
    //w.Header().Add("Access-Control-Allow-Origin", "*")
    //w.Header().Add("Content-Type", "application/json")
    var forum Forum
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&forum)
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    c := conn.C("forum")
    now := time.Now()
    forum.CreatedTime = now.Unix()
    forum.UpdatedTime = now.Unix()
    err = c.Insert(&forum)
    fmt.Println("%v\n",forum)
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


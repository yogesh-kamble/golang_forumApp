package main

/*import (
        "fmt"
        "log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)*/

type Forum struct {
        Email string
        Title string
        Message string
}

var forums []Forum

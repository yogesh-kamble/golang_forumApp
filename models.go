package main


type Forum struct {
        Email string
        Title string
        Message string
        CreatedTime int64
        UpdatedTime int64
}

var forums []Forum
var forum Forum

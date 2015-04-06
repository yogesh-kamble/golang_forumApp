package main

import (
        "gopkg.in/mgo.v2"
)

func connect() *mgo.Database {
      session, err := mgo.Dial("mongodb://yogesh:yogesh@ds045031.mongolab.com:45031/forum_app")
        if err != nil {
                panic(err)
        }
        //defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("forum_app")
        return c
}

var conn = connect()

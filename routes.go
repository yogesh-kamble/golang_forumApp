package main

import (
    "strings"
    "net/http"
    "github.com/gorilla/mux"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    //http.Handle("/", &CorsWrapper{router})
    //router.Headers("Content-Type", "application/json")
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

var routes = Routes{
    Route{
        "GetForums",
        "GET",
        "/forums",
        getForums,
    },
    Route{
        "AddForum",
        "POST",
        "/forums",
        createForum,
    },
    Route{
        "getForum",
        "GET",
        "/forums/{forumId}",
        getForum,
    },
}

type CorsWrapper struct {
    router *mux.Router
}

func (this *CorsWrapper) ServeHTTP(
    response http.ResponseWriter,
    request *http.Request) {
    allowedMethods := []string{
        "POST",
        "GET",
        "OPTIONS",
        "PUT",
        "PATCH",
        "DELETE",
    }

    allowedHeaders := []string{
        "Accept",
        "Content-Type",
        "Content-Length",
        "Accept-Encoding",
        "Authorization",
        "X-CSRF-Token",
        "X-Auth-Token",
    }

    if origin := request.Header.Get("Origin"); origin != "" {
        response.Header().Set("Access-Control-Allow-Origin", origin)

        response.Header().Set(
            "Access-Control-Allow-Methods",
            strings.Join(allowedMethods, ", "))

        response.Header().Set(
            "Access-Control-Allow-Headers",
            strings.Join(allowedHeaders, ", "))
    }

    if request.Method == "OPTIONS" {
        return
    }

    this.router.ServeHTTP(response, request)
}

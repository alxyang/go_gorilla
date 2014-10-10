package main

import (
    "net/http"
    "time"
    "flag"
    "log"

    "git-go-websiteskeleton/app/common"
    "git-go-websiteskeleton/app/home"
    "git-go-websiteskeleton/app/user"

    "github.com/gorilla/mux"
)

var router *mux.Router

func main() {
    flag.Parse()

    router = mux.NewRouter()
    http.HandleFunc("/", httpInterceptor)

    router.HandleFunc("/", home.GetHomePage).Methods("GET")
    router.HandleFunc("/user{_:/?}", user.GetHomePage).Methods("GET")

    router.HandleFunc("/user/view/{id:[0-9]+}", user.GetViewPage).Methods("GET")
    router.HandleFunc("/user/{id:[0-9]+}", user.GetViewPage).Methods("GET")

    router.HandleFunc("/test", home.GetTestingPost).Methods("POST")

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    log.Println("server started on port 2014.")
    http.ListenAndServe(":2014", nil)
}

func httpInterceptor(w http.ResponseWriter, req *http.Request) {
    startTime := time.Now()

    router.ServeHTTP(w, req)

    finishTime := time.Now()
    elapsedTime := finishTime.Sub(startTime)

    switch req.Method {
    case "GET":
        // We may not always want to StatusOK
        common.LogAccess(w, req, elapsedTime)
    case "POST":
        // here we might use http.StatusCreated
    }
}
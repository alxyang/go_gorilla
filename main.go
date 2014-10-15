package main

import (
    "flag"
    "net/http"
    "time"
    "log"

    "cilantro/app/common"
    "cilantro/app/home"
    "cilantro/app/user"

    "github.com/gorilla/mux"
)

var addr = flag.String("addr", ":8080", "http service address")
var router *mux.Router

func main() {
    flag.Parse()

    go common.H.Run()


    router = mux.NewRouter()
    http.HandleFunc("/", httpInterceptor)

    router.HandleFunc("/ws", common.WsHandler)

    router.HandleFunc("/", home.GetHomePage).Methods("GET")
    router.HandleFunc("/user{_:/?}", user.GetHomePage).Methods("GET")

    router.HandleFunc("/user/view/{id:[0-9]+}", user.GetViewPage).Methods("GET")
    router.HandleFunc("/user/{id:[0-9]+}", user.GetViewPage).Methods("GET")

    router.HandleFunc("/test", home.GetTestingPost).Methods("POST")

    // fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    // http.Handle("/static/", fileServer)
    http.Handle("/static/", http.FileServer(http.Dir(".")))

    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
    log.Println("server started.")
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
         common.LogAccess(w, req, elapsedTime)
    }
}
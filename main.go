package main

import (
    "flag"
    "net/http"
    "time"
    "log"

    "go_gorilla/app/common"
    "go_gorilla/app/websockets"
    "go_gorilla/app/home"
    "go_gorilla/app/user"

    "github.com/gorilla/mux"
)

var addr = flag.String("addr", ":8080", "http service address")
var router *mux.Router

func main() {
    flag.Parse()

    go websockets.H.Run()


    router = mux.NewRouter()
    http.HandleFunc("/", httpInterceptor)

    router.HandleFunc("/ws", websockets.WsHandler)
    router.HandleFunc("/", home.GetHomePage).Methods("GET")
    router.HandleFunc("/user{_:/?}", user.GetHomePage).Methods("GET")
    router.HandleFunc("/user/view/{id:[0-9]+}", user.GetViewPage).Methods("GET")
    router.HandleFunc("/user/{id:[0-9]+}", user.GetViewPage).Methods("GET")
    router.HandleFunc("/test", home.GetTestingPost).Methods("POST")

    // fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    // http.Handle("/static/", fileServer)
    http.Handle("/static/", http.FileServer(http.Dir(".")))

    log.Println("Attempting to start web server.")
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }

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

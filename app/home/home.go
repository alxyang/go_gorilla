package home

import (
    "net/http"
    "html/template"
    "encoding/json"
    "log"


    "cilantro/app/common"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
    }
    
    p := Page{
        Title: "home",
    }

    common.Templates = template.Must(template.ParseFiles("templates/home/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}

type User struct {
    Name string `json:"name"`
    City string `json:"city"`
}

//test POST request sending data from front-end to back-end
func GetTestingPost(rw http.ResponseWriter, req *http.Request) {
    log.Println("request")
    decoder := json.NewDecoder(req.Body)
    var newUser User
    errc := decoder.Decode(&newUser)
    if errc != nil {
        //panic()
    }
    log.Println(newUser.City)
}
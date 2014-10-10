package home

import (
    "net/http"
    "html/template"
    
    "github.com/golang/glog"
    "git-go-websiteskeleton/app/common"
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

func GetTestingPost(rw http.ResponseWriter, req *http.Request) {
    glog.Infoln("asdfadsfadsfadsf")

    type Page struct {
        Title string
    }
    
    p := Page{
        Title: "test",
    }

    common.Templates = template.Must(template.ParseFiles("templates/home/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}
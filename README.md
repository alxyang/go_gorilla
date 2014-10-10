Go Template
===============

A basic web application template written in go that comes with the Gorilla Multiplexer for routing, glog for access and error logging, as well as jquery, bootstrap, and font-awesome.

### Prerequisites ###
1. A go environment http://golang.org/doc/install#install
2. Gorilla mux: go get github.com/gorilla/mux (http://www.gorillatoolkit.org/pkg/mux)
3. glog: go get github.com/golang/glog
4. Some basic knowledge of Go's httpd package. See the excellent gowiki tutorial at http://golang.org/doc/articles/wiki/

### Installation ###
1. cd $GOPATH/src
2. git clone https://github.com/aly006/cilantro
3. cd git-go-websiteskeleton
4. go run main.go
5. Navigate to http://localhost:2014

### TODO ###
Add database connectivity
Possibly convert to Angular, may be overkill.
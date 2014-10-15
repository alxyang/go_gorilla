Go Template
===============

A basic web application template written in go that comes with the Gorilla Multiplexer for routing, glog for access and error logging, as well as jquery, bootstrap, and font-awesome.

### Prerequisites ###
1. A go environment http://golang.org/doc/install#install
2. Gorilla mux: go get github.com/gorilla/mux (http://www.gorillatoolkit.org/pkg/mux)
3. Gorilla websocket: go get github.com/gorilla/websocket (http://www.gorillatoolkit.org/pkg/websocket)

### Installation ###
1. cd $GOPATH/src
2. git clone https://github.com/aly006/cilantro
3. cd cilantro
4. go run main.go
5. Navigate to http://localhost:8080

### TODO ###
Add database connectivity
Possibly convert to Angular, may be overkill.
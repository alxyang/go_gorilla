package hardware

import (
  "net/http"
  "log"

  "github.com/gorilla/websocket"
)

type connection struct {
  // The websocket connection.
  ws *websocket.Conn

  // Buffered channel of outbound messages.
  send chan []byte
}

func (c *connection) reader() {
  log.Println("Reader goroutine initialized.")
  for {
    _, message, err := c.ws.ReadMessage()
    if err != nil {
      break
    }
    userCommand := string(message)
    log.Printf("Recieved Message: " + userCommand);
    //keep track of state using a counter?

    //connect to motion controller cmd, have it go back to zero'd position
    if (userCommand == "motion") {
      initializeMotionController();
    } else {
      H.broadcast <- message
    }
    // //set up uv lamp
    // else if (s == "uv") {

    // }
    // //set up dmd
    // else if (s == "dmd") {

    // }    
    // //set up image directory
    // else if (s == "imgdir") {

    // }
    // //set lamp iris
    // else if (s == "iris") {

    // }
    // //set motion controller to focal - z
    // else if (s == "focalz") {

    // }
    // //fabrication calculations
    // else if (s == "fabCalc") {

    // }
    // //prepare to print scaffold
    // else if (s == "scaffoldPrepare") {

    // }
    // //enter fabricaton loop process
    // else if (s == "fabLoop") {

    // }
    // //exit loop
    // else if (s == "exit") {

    // }
    // //end
    // else if (s == "end") {

    // }
    //broadcast the message to all other clients
    // else {
    //   H.broadcast <- message
    // }
  }
  c.ws.Close()
}

func (c *connection) writer() {
  log.Println("Writer goroutine initialized.")
  for message := range c.send {
    err := c.ws.WriteMessage(websocket.TextMessage, message)
    if err != nil {
      break
    }
    s := string(message)
    log.Printf("Broadcasting Message: " + s);
  }
  c.ws.Close()
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func WsHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("entered websocket handler")
  ws, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Println(err)
    return
  }
  c := &connection{send: make(chan []byte, 256), ws: ws}
  H.register <- c
  defer func() { H.unregister <- c }()
  go c.writer()
  c.reader()
}

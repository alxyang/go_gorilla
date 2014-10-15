package hardware

import (
  "log"
)

func initializeMotionController() {
  log.Println("entered websocket handler")
  H.broadcast <- []byte("weeee")
}
package hardware

import (
  "log"

  "github.com/dmdcilantro/hardware/nivisa"
)

func initializeMotionController() {
  log.Println("entered websocket handler")
  H.broadcast <- []byte("weeee")

// COM Port Settings initialization ====================================

  COM_setting_ESP300 := "ASRL4::INSTR" //currently hardcoded.

//  Motion Controller setup ====================================
    // Establish connection
  var (
    esp300 *nivisa.Esp300
    errm    error
  )

  log.Println("Connecting to Motion Controller... ")
  if esp300, errm = nivisa.NewEsp300(COM_setting_ESP300); errm != nil {
    log.Printf("err: %v\n", errm)
    os.Exit(1)
  }
  log.Printf("        ... Done.\n\n")

// Initialize to home position
  log.Println("Zeroing stage to home position...")
  log.Printf("\n")
  
  init_speed := float32(3)
  z_curr_pos := float32(-10)
  z_prev_pos := z_curr_pos
  
  esp300.SendCmd(log.Sprintf("1MO")) 
  time.Sleep(1000*time.Millisecond)

  esp300.SetVelocity(1,init_speed)
  time.Sleep(1000*time.Millisecond)
  // time.Sleep(2000*time.Millisecond)
  // z_curr_pos, _ = esp300.Position(1)
  // time.Sleep(500*time.Millisecond)
  // z_curr_pos, _ = esp300.Position(1)
  // time.Sleep(500*time.Millisecond)
  // z_curr_pos, _ = esp300.Position(1)
  // time.Sleep(500*time.Millisecond)
  
  
  go esp300.SendCmd(log.Sprintf("1OR0"))
  time.Sleep(1000*time.Millisecond)
  
  loop_count := 0
  z_prev_pos, _ = esp300.Position(1)
  z_curr_pos = z_prev_pos

  for z_curr_pos > 0.0001 && loop_count < 100 {
    z_curr_pos, _ = esp300.Position(1)
    
    //broadcast this through the websocket
    log.Printf("Current z = %+8.4f",z_curr_pos)
    log.Printf("\n")

    H.broadcast <- []byte("looping real time")

    if z_curr_pos == z_prev_pos {
      loop_count++
      // if stuck, try sending a non-concurrent command to zero the stage
      if loop_count > 10 {
        esp300.SendCmd(log.Sprintf("1OR0"))
        time.Sleep(1000*time.Millisecond) 
      }
    }
    z_prev_pos = z_curr_pos 
  }

  if loop_count >= 100 {
    log.Println("Stage could not be zeroed.")
    os.Exit(1)
  }

  log.Printf("\n")
  log.Print("Stage zeroed to home position: ")
  log.Printf("z = %8.4f",z_curr_pos)
  log.Printf("\n\n\n")
  time.Sleep(1500*time.Millisecond)

}
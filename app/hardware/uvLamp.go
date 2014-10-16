package hardware

// import (
//   "log"

//   "github.com/dmdcilantro/hardware/omnicure"
// )

// func initializeUvLamp() {
//   //  Lamp Setup =================================================
//     log.Println("Connecting to UV Lamp... ")
//     var state_uv bool =false
//     s2000, err := omnicure.NewS2000(COM_setting_UV)
//     if err == nil{
//       state_uv = true
//     }
//     if err != nil {
//       log.Println("Could not connect uv lamp.")
//       os.Exit(1)
//     }

//     // connect
//     if state_uv{
//       if _, err := s2000.SendCommand("CONN"); err == nil{
//         //log.Printf("Result? %s\n", res)
//         log.Println("Done.")
//       } else {
//         log.Printf("Read Err: %v\n", err)
//         os.Exit(1)
//       }
//     }
// }
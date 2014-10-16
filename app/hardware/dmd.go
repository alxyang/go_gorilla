package hardware

// import (
//   "log"
//   "os"

//   "github.com/dmdcilantro/hardware/dmd"
// )

// func initializeUvLamp() {


//   //  DMD Setup ==================================================
//   var dev *dmd.Device
//   var errd error

//   log.Printf("Initializing DMD...")
//   if errd = dmd.Initialize(); errd != nil {
//     log.Printf("Error on Initialize: %v\n", err)
//     os.Exit(1)
//   }
    
//   if devCount := dmd.NumDev(); devCount != 1 {
//     log.Printf("Unexpected number of DMD devices detected (%v founded; expected 1)\n", devCount)
//     os.Exit(1)
//   }

//   if dev, errd = dmd.ConnectDevice(0); errd != nil {
//     log.Printf("Error on ConnectDevice: %v\n", errd)
//     os.Exit(1)
//   }

//   dev.ClearFifos()
//   dev.Clear(dmd.AllBlocks, true)
//   log.Println("Done.")

//   choose_new_settings := 1
//   file_dir_value := "nil"
//   prev_dir_value := "nil"
  
//   imageList := make([][]byte, 0)

//   for choose_new_settings == 1 {

//     //  Retrieve, parse, validate, and convert inputs ==============

//     if _, err := os.Stat("fab_params.txt"); err == nil {

//       fmt.Printf("\n")
//       fmt.Print("Fabrication parameters file ('fab_params.txt') exists. Processing... ")

      

//       contents,_ := ioutil.ReadFile("fab_params.txt")
//       text_reader := strings.NewReader(string(contents))
//       text_scanner := bufio.NewScanner(text_reader)
//       text_scanner.Split(bufio.ScanWords)

//       var contents_by_line []string
//       for text_scanner.Scan() {
//         contents_by_line = append(contents_by_line,text_scanner.Text())
//       }

//       last_line := "nil"
//       file_iris_value := "nil"
//       file_height_value := "nil"
//       file_focal_z_value := "nil"
//       //file_speed_value := "nil"
//       file_time_value := "nil"
      
//       prev_dir_value = file_dir_value
//       file_dir_value = "nil"
      
//       for _, each_line := range contents_by_line {
        
//         if last_line == "iris" {
//           file_iris_value = each_line
//         } else if last_line == "height" {
//           file_height_value = each_line
//         } else if last_line == "focal_z" {
//           file_focal_z_value = each_line
//         // } else if last_line == "speed" {
//         //  file_speed_value = each_line
//         } else if last_line == "time" {
//           file_time_value = each_line
//         } else if last_line == "directory" {
//           file_dir_value = each_line
//         }

//         last_line = each_line
//       }


//       *iris_strpnt = file_iris_value

//       iris_val, _ :=strconv.ParseFloat(*iris_strpnt, 64)
//       scaf_z, _ :=strconv.ParseFloat(file_height_value, 64)
//       focal_pos_val, _ :=strconv.ParseFloat(file_focal_z_value, 64)
//       // speed_z, _ :=strconv.ParseFloat(file_speed_value, 64)
//       timef, _ :=strconv.ParseFloat(file_time_value, 64)

//       iris_val_int = int16(iris_val)
//       scaf_z32 = float32(scaf_z)
//       focal_pos = float32(focal_pos_val)
//       // init_speed = float32(speed_z)
//       exp_time = float32(timef)

//       if len(file_dir_value)-1 != strings.LastIndex(file_dir_value, "\\"){
//         *directory = file_dir_value + "\\"
//       }

//       log.Println("Done.")
//       } else {
//         // Validate numeric inputs
//         for i:= 0; i<len(letters); i++ { 
//           if strings.ContainsAny(*height, letters[i] ){
//             log.Println("invalid input letter in height: "+ letters[i])
//             os.Exit(1)
//           }

//           // if strings.ContainsAny(*speed, letters[i]){
//           //  fmt.Println("invalid input letter in speed: "+ letters[i])
//           //  os.Exit(1)
//           // }

//           if strings.ContainsAny(*iris_strpnt, letters[i]) || *iris_strpnt == "nil"{
//             log.Println("invalid input letter in iris: " + letters[i])
//             os.Exit(1)
//           }
//         }

//         // Validate and format directory input
//         if *directory == "nil"{
//           log.Println("Not a valid directory")
//           os.Exit(1)
//         }

//         if len(*directory)-1 != strings.LastIndex(*directory, "\\"){
//           *directory = *directory + "\\"
//         }

//         // Convert values
//         scaf_z, _ :=strconv.ParseFloat(*height, 64)
//         timef, _ :=strconv.ParseFloat(*time_value, 64)
//         iris_val, _ :=strconv.ParseFloat(*iris_strpnt, 64)
//         // speed_z, _ :=strconv.ParseFloat(*speed, 64)
//         focal_pos_val, _ :=strconv.ParseFloat(*focal_pos_flag, 64)

//         iris_val_int = int16(iris_val)
//         scaf_z32 = float32(scaf_z)
//         focal_pos = float32(focal_pos_val)
//         // init_speed = float32(speed_z)
//         exp_time = float32(timef)

//         log.Println("Done.")
//       }

//     //  Image Directory Setup ======================================
//       if file_dir_value != prev_dir_value {

//         fmt.Println("Preparing Masks...")
//         fmt.Printf("\n")

//         imageList = nil
//         imgs, _ := ioutil.ReadDir(*directory)
//         mask_i := 0

//         for _, img := range imgs {

//           mask_i++
//           fmt.Printf("  Importing Mask %4v of %v",mask_i,len(imgs))
//           fmt.Println("")

//           f, _ := os.Open( *directory+img.Name())
//           imgData, _ := png.Decode(f)

//           b := imgData.Bounds()

//           dmdSize := 1024 * 768
//           imgBytes := make([]byte, dmdSize/8)

//           srcImg := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
//           draw.Draw(srcImg, srcImg.Bounds(), imgData, b.Min, draw.Src)

//           rgbaImg := image.NewRGBA(image.Rect(0, 0, 1024, 768))
//           draw.DrawMask(rgbaImg, image.Rect(0, 0, 1024, 768), srcImg, image.Pt(0,0), nil, image.Point{}, 1)

//           xMax := rgbaImg.Bounds().Dx()
//           yMax := rgbaImg.Bounds().Dy()

//           for y := 0; y < yMax; y++ {
//             for x := 0; x < xMax; x++ {
//               r, g, b, _ := rgbaImg.At(x, y).RGBA()

//               // idx := 1024*y + x
//               idx := xMax * y + x

//               posByte := uint(idx / 8)
//               posBit := uint(idx % 8)

//               current := ((0xFF7F >> posBit) & int(imgBytes[posByte])) & 0x00FF


//               if r+g+b < 384 {
//                 imgBytes[posByte] = ((0 << (8 - (posBit + 1))) | byte(current))
                
//               } else {
//                 imgBytes[posByte] = ((1 << (8 - (posBit + 1))) | byte(current))
                
//               }
//             }
//           }

//           imageList = append(imageList, imgBytes)
//         }

//         fmt.Printf("\n")
//         fmt.Println("Masks ready.")
//         fmt.Printf("\n\n")
      
//       } else if file_dir_value == "nil" {
//         fmt.Printf("\n")
//         fmt.Println("Directory for masks not specified.")
//         os.Exit(1)
//       } else {
//         fmt.Printf("\n")
//         fmt.Println("Using the same mask set.")
//         fmt.Printf("\n\n")        
//       }
// }


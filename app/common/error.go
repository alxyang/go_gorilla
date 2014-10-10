package common

import (
    "log"
    "runtime"
)

// 0 = info
// 1 = warning
// 2 = error - should be most common
// 3 = fatal
func CheckError(err error, level int) {
    if err != nil {
        var stack [4096]byte
        runtime.Stack(stack[:], false)
        log.Printf("%q\n%s\n", err, stack[:])

        switch level {
        case 0:
            log.Println("%q\n%s\n", err)
        case 1:
            log.Println("%q\n%s\n", err)
        case 2:
            log.Println("%q\n%s\n", err)
        case 3:
            log.Println("%q\n%s\n", err)
        }
    }
}
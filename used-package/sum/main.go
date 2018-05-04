package sum

import "fmt"

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L./src -lsum
// #include <src/sum.h>
import "C"

func Sum(a, b int) {
	fmt.Printf("Invoking C library...\n")
	fmt.Println("Sum: ", C.sum(C.int(a), C.int(b)))
}

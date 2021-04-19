package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("[")
	fmt.Println("matched files", files)
}

// package main

// import (
// 	"fmt"
// 	"path/filepath"
// )

// func main() {
// 	files, error := filepath.Glob("[")
// 	if error != nil && error == filepath.ErrBadPattern {
// 		fmt.Println(error)
// 		return
// 	}
// 	fmt.Println("matched files", files)
// }

// package main

// import (
// 	"fmt"
// 	"net"
// )

// func main() {
// 	addr, err := net.LookupHost("golangbot123.com")
// 	if err, ok := err.(*net.DNSError); ok {
// 		if err.Timeout() {
// 			fmt.Println("operation timed out")
// 		} else if err.Temporary() {
// 			fmt.Println("temporary error")
// 		} else {
// 			fmt.Println("generic error: ", err)
// 		}
// 		return
// 	}
// 	fmt.Println(addr)
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	f, err := os.Open("/test.txt")
// 	if err, ok := err.(*os.PathError); ok {
// 		fmt.Println("File at path", err.Path, "failed to open")
// 		return
// 	}
// 	fmt.Println(f.Name(), "opened successfully")
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	f, err := os.Open("/test.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(f.Name(), "opened successfully")
// }

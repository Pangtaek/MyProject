// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", helloWorld)

// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		fmt.Println("Failed to start server:", err)
// 	}
// }

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	// w.Write([]byte("Hello, World!"))
// 	fmt.Println("Hello, World!")
// }

package main

import "github.com/pangtaek/MyProject/init/cmd"

func main() {
	cmd.NewCmd("./config.toml")
}

package main

import (
	"flag"
	"fmt"
)

// var (
// 	errValidation = errors.New("validation error")
// 	errNotFound   = errors.New("not found error")
// )

func main() {
	// // 1. FMT
	// firstname := "jalal"
	// lastname := "akbar"
	// fmt.Printf("FMT: "+"Hello `%s %s`\n", firstname, lastname)
	// // 2. errors
	// err := getById("")
	// if err != nil {
	// 	if errors.Is(err, errValidation) {
	// 		fmt.Println("validation error")
	// 	} else if errors.Is(err, errNotFound) {
	// 		fmt.Println("not found error")
	// 	} else {
	// 		fmt.Println("unknown error")
	// 	}
	// }
	// fmt.Println("Success")
	// // 3. os
	// args := os.Args

	// for i, arg := range args {
	// 	fmt.Println("arg:", arg[i])
	// }
	// hostname, err := os.Hostname()
	// if err == nil {
	// 	fmt.Println("hostname", hostname)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// 4. flag
	var username = flag.String("username", "root", "database username")
	var password = flag.String("password", "root", "database pasword")
	var hostName = flag.String("hostname", "root", "database hostname")
	var port = flag.Int("port", 0, "database port")
	flag.Parse()

	fmt.Println("username: ", *username)
	fmt.Println("password: ", *password)
	fmt.Println("hostname: ", *hostName)
	fmt.Println("port: ", *port)
}

// func getById(id string) error {
// 	if id == "" {
// 		return errValidation
// 	} else if id != "jalal" {
// 		return errNotFound
// 	}
// 	return nil
// }

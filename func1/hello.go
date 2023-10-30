package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	// Create a route with the folder name: func1
	// and assign it to helloHandler
	mux := http.NewServeMux()
	route := fmt.Sprintf("/api/%s", getDirName())
	mux.HandleFunc(route, helloHandler)

	// Get Listen address from Env: FUNCTIONS_CUSTOMHANDLER_PORT
	listenAddr := getListenAddr()
	fmt.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)

	// Start Server
	log.Fatal(http.ListenAndServe(listenAddr, mux))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "This HTTP triggered function executed successfully. Pass a name in the query string for a personalized response.\n"
	name := r.URL.Query().Get("name")
	if name != "" {
		message = fmt.Sprintf("Hello, %s. This HTTP triggered function executed successfully.\n", name)
	}
	fmt.Fprint(w, message)
}

func getListenAddr() string {
	listenAddr := ":8080"
	port, _ := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	fmt.Println("FUNCTIONS_CUSTOMHANDLER_PORT is: " + port)
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	return listenAddr
}
func getDirName() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		fmt.Println("Exitting....")
		os.Exit(1)
	}
	dir := filepath.Base(cwd)

	return dir
}

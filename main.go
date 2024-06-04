package main

import (
	"net/http"
	 "os"
	 "fmt"
	 "errors"
	 "path/filepath"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("static", "home.html")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}
func cvHandler(w http.ResponseWriter, r *http.Request)   {
	filePath := filepath.Join("static" , "cv.html")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w , "404 not found.", http.StatusNotFound)
		return
	}
	
	http.ServeFile(w , r , filePath)
}
func blogHandler(w http.ResponseWriter, r *http.Request)   {
	filePath := filepath.Join("static" , "blog.html")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w , "404 not found.", http.StatusNotFound)
		return
	}
	
	http.ServeFile(w , r , filePath)
}


func main() {
	fmt.Println("Server starting...")

	// Corrected the directory path and function name for serving static files
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/cv", cvHandler)
	http.HandleFunc("/blog", blogHandler)

	// Starting the HTTP server
	err := http.ListenAndServe("127.0.0.1:3000", nil)
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server closed gracefully.")
		} else {
			fmt.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}
	
}

package internal

import (
	"fmt"
	"net/http"
)

func HandleRequest() {
	fmt.Println("click on the link", "'http://localhost:8080/'")
	http.HandleFunc("/", main_page)
	http.HandleFunc("/asciiArt", asciiArt_page)
	http.ListenAndServe(":8080", nil)
}

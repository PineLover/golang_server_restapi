package main

/*
import (
	"fmt"
	"net/http"

	"github.com/PineLover/golang_server_restapi/server/myapp"
)

func main() {
	fmt.Print("running")

	http.ListenAndServe(":3002", myapp.NewHttpHandler())
}
*/

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile, header, err := r.FormFile("upload_file")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Print(w, filepath)
}

func main() {
	http.HandleFunc("/uploads", uploadsHandler)
	http.Handle("/", http.FileServer((http.Dir("public"))))
	http.ListenAndServe(":3002", nil)
}

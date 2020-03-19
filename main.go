package main

import (
	"fmt"
	"goyun/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("failed to start server ,err: %s", err.Error())
	}
}

package main

import (
	"fmt"
	"net/http"
)

func initHTTP() {
	http.HandleFunc("/get_val", handleGetVal)
	fmt.Println("SERVING ... ")
	http.ListenAndServe(":8181", nil)
}

func handleGetVal(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprint("nilainya : ", getI())))
}

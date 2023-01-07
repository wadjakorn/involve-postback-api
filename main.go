package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte("Success"))
		return
	})
	http.HandleFunc("/postback", func(w http.ResponseWriter, r *http.Request) {
		clickid := r.URL.Query().Get("clickid")
		fmt.Println("clickid =>", clickid)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte(fmt.Sprintf("clickid =>%v", clickid)))
		return
	})
	http.ListenAndServe(":3000", nil)
}

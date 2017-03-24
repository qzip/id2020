package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"uuid"
)

// uuid, the web server
func UUIDServer(w http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()
	count := 1
	var err error
	if values != nil {
		cnt := values.Get("count")
		if cnt != "" {
			count, err = strconv.Atoi(cnt)
			if err != nil {
				fmt.Fprintf(w, "{\"Error\":\"%s\"}\n", err)
				return
			} 
		}
	}
	for ; count > 0; count -= 1 {
		uuid, err := uuid.GenUUID()
		if err != nil {
			fmt.Fprintf(w, "{\"Error\":\"%s\"}\n", err)
		} else {
			fmt.Fprintf(w, "{\"uuid\":[\"%s\"]}\n", uuid)
		}
	}
}

func main() {
	http.HandleFunc("/_uuid", UUIDServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

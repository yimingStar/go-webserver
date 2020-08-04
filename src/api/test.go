package api

import (
	"encoding/json"
	"net/http"
	utils "go-webservice/src/utils"
)

func TestAlive(w http.ResponseWriter, r *http.Request) {  
	res := setResponceBody("I'm alive", utils.ErrorCode["success"])
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}
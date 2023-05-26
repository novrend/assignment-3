package controllers

import (
	"assignment-3/models"
	"assignment-3/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var bodyJson = models.Data{}
	json.Unmarshal(body, &bodyJson)
	err := repository.UpdateData(bodyJson)
	if err != nil {
		http.Error(w, "Error update data", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Data updated successfully",
	})
}

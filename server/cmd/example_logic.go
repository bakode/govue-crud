package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aasumitro/govue/server/db/queries"
	"github.com/aasumitro/govue/server/entities"
	"github.com/gorilla/mux"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, PUT")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Content-Type", "application/json")

	var response entities.ResponseMany
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = queries.All()
	var result, err = json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
}

func FindOne(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, PUT")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	passID := vars["id"]
	id, _ := strconv.Atoi(passID)
	var successCallback entities.ResponseRow
	var errorCallback entities.Response
	var data = queries.Find(id)
	if data.Id != 0 {
		successCallback.Status = http.StatusOK
		successCallback.Message = http.StatusText(http.StatusOK)
		successCallback.Data = data
		var result, err = json.Marshal(successCallback)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
	} else {
		errorCallback.Status = http.StatusNotFound
		errorCallback.Message = http.StatusText(http.StatusNotFound)
		var callback, _ = json.Marshal(errorCallback)
		w.Write(callback)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, PUT")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Content-Type", "application/json")

	title := r.FormValue("title")
	description := r.FormValue("description")
	var result = queries.Store(title, description)
	var callback entities.Response
	if result == false {
		callback.Status = http.StatusBadRequest
		callback.Message = http.StatusText(http.StatusBadRequest)
		var result, _ = json.Marshal(callback)
		w.Write(result)
	} else {
		callback.Status = http.StatusCreated
		callback.Message = http.StatusText(http.StatusCreated)
		var result, _ = json.Marshal(callback)
		w.Write(result)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, PUT")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	passID := vars["id"]
	id, _ := strconv.Atoi(passID)
	title := r.FormValue("title")
	description := r.FormValue("description")
	rowsAffected, err := queries.Update(id, title, description)
	var callback entities.Response
	if rowsAffected < 1 && err == nil {
		fmt.Println(err)
		callback.Status = http.StatusNotModified
		callback.Message = http.StatusText(http.StatusNotModified)
		var result, _ = json.Marshal(callback)
		w.Write(result)
	} else {
		fmt.Println("Rows Affected:", rowsAffected)
		callback.Status = http.StatusCreated
		callback.Message = http.StatusText(http.StatusCreated)
		var result, _ = json.Marshal(callback)
		w.Write(result)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	passID := vars["id"]
	id, _ := strconv.Atoi(passID)
	rowsAffected, err := queries.Destroy(id)
	var callback entities.Response
	if rowsAffected < 1 && err == nil {
		fmt.Println(err)
		callback.Status = http.StatusNotModified
		callback.Message = http.StatusText(http.StatusNotModified)
		var result, _ = json.Marshal(callback)
		w.Write(result)
	} else {
		fmt.Println("Rows Affected:", rowsAffected)
		callback.Status = http.StatusCreated
		callback.Message = http.StatusText(http.StatusCreated)
		var result, _ = json.Marshal(callback)
		w.Write(result)
	}
}

package simpleHttpServer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is rest api for human data.")
}

func updateHumanData(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, "Parametar id does not correct."+err.Error(), http.StatusBadRequest)
		return
	}
	var h Human
	json.NewDecoder(r.Body).Decode(&h)

	err = updateHumanForDataFile(id, h)
	if err != nil {
		http.Error(w, fmt.Sprintf("Can not update human data.: >> %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("Human %v is updated.", h.Id))
	return
}

func deleteHumanData(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, "Parametar id does not correct."+err.Error(), http.StatusBadRequest)
		return
	}
	err = deleteHumnFromDataFile(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Can not delete human data.: >> %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("Human %v is deleted.", id))
	return
}

func createHumanData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var h Human
	json.NewDecoder(r.Body).Decode(&h)

	err := appendHumanToDataFile(h)
	if err != nil {
		http.Error(w, fmt.Sprintf("Can not create human data.: >> %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("Human %v is created.", h.Id))
	return
}

func getHumanData(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, "Parametar id does not correct."+err.Error(), http.StatusBadRequest)
		return
	}
	humans, err := readDataFile()
	if err != nil {
		http.Error(w, "Can not read data.", http.StatusInternalServerError)
		return
	}

	for _, h := range humans {
		if h.Id == id {
			returnJsonData, err := json.Marshal(h)
			if err != nil {
				http.Error(w, "Can not convert to json text.", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(returnJsonData)
			return
		}
	}
	http.Error(w, "Data not found.", http.StatusNotFound)
}

func getHumanDataList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	humans, err := readDataFile()
	if err != nil {
		http.Error(w, "Can not read data.", http.StatusInternalServerError)
		return
	}

	returnJsonData, err := json.Marshal(humans)
	if err != nil {
		http.Error(w, "Can not convert to json text.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(returnJsonData)
	return
}

func RunServer() {

	router := httprouter.New()
	router.GET("/about", about)
	router.GET("/human", getHumanDataList)
	router.GET("/human/:id", getHumanData)
	router.POST("/human", createHumanData)
	router.PUT("/human/:id", updateHumanData)
	router.DELETE("/human/:id", deleteHumanData)

	port := 8080
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	if err != nil {
		log.Fatal("Can not run server.")
	}

}

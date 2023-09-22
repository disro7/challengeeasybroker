package controller

import (
	mo "challengeeasybroker/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	ApiKey = "l7u502p8v46ba3ppgvj5y2aad50lb9"
)

func ContenidoContactShow(w http.ResponseWriter, r *http.Request) {

	limitNumber, limitErr := strconv.Atoi(r.URL.Query().Get("limit"))
	pageNumber, pageNumberErr := strconv.Atoi(r.URL.Query().Get("page"))
	if limitErr != nil || pageNumberErr != nil {
		w.WriteHeader(404)
	}
	page := strconv.Itoa(pageNumber)
	limit := strconv.Itoa(limitNumber)

	url := "https://api.stagingeb.com/v1/contact_requests?page="
	req, err := http.NewRequest("GET", url+page+"&limit="+limit, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("X-Authorization", "l7u502p8v46ba3ppgvj5y2aad50lb9")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	responseContenidoContacto(w, 200, bodyText)
}

func responseContenidoContacto(w http.ResponseWriter, status int, results []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var obj mo.Contact
	if err := json.Unmarshal(results, &obj); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(obj)
}

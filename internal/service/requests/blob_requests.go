package requests

import (
	"blob-service/internal/service/handlers"
	res "blob-service/resources"
	"encoding/json"
	"gitlab.com/distributed_lab/ape"
	"net/http"
	"strconv"
	"strings"
)

func retrieve_id(r *http.Request) (int, error) {
	str_repr := strings.TrimPrefix(r.URL.Path, "/blob-service/blob/")
	return strconv.Atoi(str_repr)
}

func Get_blob_by_id(w http.ResponseWriter, r *http.Request) {
	id, err := retrieve_id(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	blob, err := handlers.Get_blob_by_id(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	ape.Render(w, res.BlobResponse{Data: *blob})
}

func Create_new_blob(w http.ResponseWriter, r *http.Request) {
	new_blob := new(res.Blob)
	err := json.NewDecoder(r.Body).Decode(new_blob)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handlers.Save_blob(new_blob)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func Delete_blob(w http.ResponseWriter, r *http.Request) {
	id, err := retrieve_id(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handlers.Delete_Blob(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func Get_page_of_blobs(w http.ResponseWriter, r *http.Request) {
	query_params := r.URL.Query()
	params := make(map[string]string)
	page_limit := query_params.Get("page[limit]")
	if page_limit != "" {
		params["limit"] = page_limit
	}
	page_number := query_params.Get("page[number]")
	if page_number != "" {
		params["number"] = page_number
	}
	page_order := query_params.Get("page[order]")
	if page_order != "" {
		params["order"] = page_order
	}
	response, err := handlers.Get_blobs(params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	ape.Render(w, response)
}

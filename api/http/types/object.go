package types

import (
	"encoding/json"
	"errors"
	"http_server_arch/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GetObjectHandlerRequest struct {
	ID string `json:"id"`
}

func CreateGetObjectHandlerRequest(r *http.Request) (*GetObjectHandlerRequest, error) {
	id := r.URL.Query().Get("id")
	if id == "" {
		return nil, errors.New("missing id")
	}
	return &GetObjectHandlerRequest{ID: id}, nil
}

type PutObjectHandlerRequest struct {
	Code     string `json:"code"`
	Compiler string `json:"compiler"`
}

func CreatePutObjectHandlerRequest(r *http.Request) (*PutObjectHandlerRequest, error) {
	var request PutObjectHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.New("error while decoding json")
	}
	return &request, nil
}

type PostObjectHandlerRequest struct {
	Code     string `json:"code"`
	Compiler string `json:"compiler"`
}

func CreatePostObjectHandlerRequest(r *http.Request) (*PostObjectHandlerRequest, error) {
	var request PostObjectHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.New("error while decoding json")
	}
	return &request, nil
}

type DeleteObjectHandlerRequest struct {
	ID string `json:"id"`
}

func CreateDeleteObjectHandlerRequest(r *http.Request) (*DeleteObjectHandlerRequest, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, errors.New("missing id")
	}
	return &DeleteObjectHandlerRequest{ID: id}, nil
}

type GetStatusHandlerResponse struct {
	Status *string `json:"status"`
}

type GetResultHandlerResponse struct {
	Result *string `json:"result"`
}

type PutObjectHandlerResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type PostObjectHandlerResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func ProcessError(w http.ResponseWriter, err error, resp any) {
	if err == repository.ErrNotFound {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if resp != nil {
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}
	}
}

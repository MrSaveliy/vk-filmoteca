package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/MrSaveliy/vk-filmoteca/models"
)

func (h *Handler) createMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	var input models.Movie
	if err := json.Unmarshal(body, &input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	id, err := h.services.Movies.CreateMovie(input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Internal Server Error")
		return
	}

	response := map[string]interface{}{
		"id": id,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func (h *Handler) getMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	movieIdStr := pathParts[len(pathParts)-1]
	movieId, err := strconv.Atoi(movieIdStr)
	if err != nil {
		http.Error(w, "Invalid movie ID"+err.Error(), http.StatusBadRequest)
		log.Fatalf("Invalid movie ID %s", err.Error())
		return
	}
	movie, err := h.services.Movies.GetMovieById(movieId)
	if err != nil {
		http.Error(w, "Error retrieving movie: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"name":        movie.Name,
		"description": movie.Description,
		"releaseDate": movie.ReleaseDate,
		"rating":      movie.Rating,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

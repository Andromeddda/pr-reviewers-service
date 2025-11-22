package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"prs/internal/dto"
	"prs/internal/repository"
	"prs/internal/service"
)

type PRSHandler struct {
	service 	service.PRService
}

func NewPRSHandler(s service.PRService) *PRSHandler {
	return &PRSHandler{
		service: s,
	}
}

func (h* PRSHandler) AddTeam(w http.ResponseWriter, r *http.Request) {
	var request dto.Team

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("ERROR: Invalid request body for AddTeam: %s", err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return 
	}

	res, err := h.service.AddTeam(r.Context(), &request)
	if err != nil {
		// 400
		if errors.Is(err, repository.ErrTeamExist) {
			log.Printf("ERROR: Team already exist: %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 500
		log.Printf("ERROR: Error creating order: %s", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 201
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
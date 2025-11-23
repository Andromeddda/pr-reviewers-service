package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"prs/internal/dto"
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

func (h* PRSHandler) writeError(w http.ResponseWriter, status int, code dto.ErrorCode, message string) {
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    resp := dto.ErrorResponse{
        Error: dto.ErrorObj{
            Code:    code,
            Message: message,
        },
    }

    json.NewEncoder(w).Encode(resp)
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
		if errors.Is(err, service.ErrTeamExist) {
			message := fmt.Sprintf("AddTeam team_name=%s : %s", request.TeamName, err.Error())
			log.Println(message)
			h.writeError(w, http.StatusBadRequest, dto.ErrorTeamExist, message)
			return
		}

		// 500
		log.Printf("ERROR: Error creating team: %s", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 201
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h* PRSHandler) GetTeam(w http.ResponseWriter, r *http.Request) {
	team_name := r.URL.Query().Get("team_name")

	if team_name == "" {
		log.Printf("ERROR: Invalid request for GetTeam: %s", team_name)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return 
	}

	res, err := h.service.GetTeam(r.Context(), team_name)
	if err != nil {
		// 404
		if errors.Is(err, service.ErrTeamNotFound) {
			message := fmt.Sprintf("GetTeam team_name=%s : %s", team_name, err.Error())
			log.Println(message)
			h.writeError(w, http.StatusNotFound, dto.ErrorNotFound, message)
			return
		}

		// 500
		log.Printf("ERROR: Error getting team: %s", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
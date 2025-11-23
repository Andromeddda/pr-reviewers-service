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


func (h* PRSHandler) UserSetIsActive(w http.ResponseWriter, r *http.Request) {
	var body struct {
        UserID   string `json:"user_id"`
        IsActive bool   `json:"is_active"`
    }

    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, "invalid body", http.StatusBadRequest)
        return
    }

	res, err := h.service.UserSetIsActive(r.Context(), body.UserID, body.IsActive)
	if err != nil {
		// 404
		if errors.Is(err, service.ErrUserNotFound) {
			message := fmt.Sprintf("UserSetIsActive user_id=%s : %s", body.UserID, err.Error())
			log.Println(message)
			h.writeError(w, http.StatusNotFound, dto.ErrorNotFound, message)
			return
		}

		// 500
		log.Printf("ERROR: Error updating user: %s", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h* PRSHandler) CreatePullRequest(w http.ResponseWriter, r *http.Request) {
	var body struct {
        PullRequestID   string 	`json:"pull_request_id"`
        PullRequestName string  `json:"pull_request_name"`
		AuthorID		string	`json:"author_id"`	
    }

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("ERROR: Invalid request body for CreatePullRequest: %s", err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return 
	}

	res, err := h.service.CreatePullRequest(r.Context(), body.PullRequestID, body.PullRequestName, body.AuthorID)
	if err != nil {
		// 409
		if errors.Is(err, service.ErrPRExist) {
			message := fmt.Sprintf("CreatePullRequest pull_request_id=%s : %s", body.PullRequestID, err.Error())
			log.Println(message)
			h.writeError(w, http.StatusConflict, dto.ErrorPRExist, message)
			return
		}

		// 404
		if errors.Is(err, service.ErrAuthorNotFound) {
			message := fmt.Sprintf("CreatePullRequest pull_request_id=%s : %s", body.PullRequestID, err.Error())
			log.Println(message)
			h.writeError(w, http.StatusNotFound, dto.ErrorNotFound, message)
			return
		}

		// 500
		log.Printf("ERROR: Error creating pull request: %s", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 201
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h* PRSHandler) MergePullRequest(w http.ResponseWriter, r *http.Request) {
	var body struct {
        PullRequestID   string 	`json:"pull_request_id"`
    }

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("ERROR: Invalid request body for CreatePullRequest: %s", err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return 
	}

	res, err := h.service.MergePullRequest(r.Context(), body.PullRequestID)
	if err != nil {
		// 404
		if errors.Is(err, service.ErrPRNotFound) {
			message := fmt.Sprintf("MergePullRequest pull_request_id=%s : %s", body.PullRequestID, err.Error())
			log.Println(message)
			h.writeError(w, http.StatusNotFound, dto.ErrorNotFound, message)
			return
		}

		// 500
		log.Printf("ERROR: Error merging pull request: %s", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h* PRSHandler) ReassignPullRequest(w http.ResponseWriter, r *http.Request) {
	var body struct {
        PullRequestID   string 	`json:"pull_request_id"`
		OldUserId   	string 	`json:"old_user_id"`
    }

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("ERROR: Invalid request body for ReassignPullRequest: %s", err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return 
	}

	res, err := h.service.ReassignPullRequest(r.Context(), body.PullRequestID, body.OldUserId)
	if err != nil {
		// 404
		if errors.Is(err, service.ErrPRNotFound) {
			message := fmt.Sprintf("ReassignPullRequest pull_request_id=%s : %s", body.PullRequestID, err.Error())
			log.Println(message)
			h.writeError(w, http.StatusNotFound, dto.ErrorNotFound, message)
			return
		}

		// 409
		if errors.Is(err, service.ErrPRMerged) {
			message := fmt.Sprintf("ReassignPullRequest pull_request_id=%s : %s", body.PullRequestID, err.Error())
			log.Println(message)
			h.writeError(w, http.StatusConflict, dto.ErrorPRMerged, message)
			return
		}

		// 409
		if errors.Is(err, service.ErrNotAssigned) {
			message := fmt.Sprintf("ReassignPullRequest pull_request_id=%s : %s", body.PullRequestID, err.Error())
			log.Println(message)
			h.writeError(w, http.StatusConflict, dto.ErrorNotAssigned, message)
			return
		}

		// 409
		if errors.Is(err, service.ErrNoCandidate) {
			message := fmt.Sprintf("ReassignPullRequest pull_request_id=%s : %s", body.PullRequestID, err.Error())
			log.Println(message)
			h.writeError(w, http.StatusConflict, dto.ErrorNoCandidate, message)
			return
		}

		// 500
		log.Printf("ERROR: Error merging pull request: %s", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h* PRSHandler) UsersGetReview(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("user_id")

	res, err := h.service.UsersGetReview(r.Context(), user_id)
	if err != nil {
		// 404
		if errors.Is(err, service.ErrUserNotFound) {
			message := fmt.Sprintf("UsersGetReview user_id=%s : %s", user_id, err.Error())
			log.Println(message)
			h.writeError(w, http.StatusNotFound, dto.ErrorNotFound, message)
			return
		}

		// 500
		log.Printf("ERROR: Error getting user's reviews: %s", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

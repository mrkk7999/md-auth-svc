package controller

import (
	"context"
	"encoding/json"
	signin "md-auth-svc/request_response/sign_in"
	signup "md-auth-svc/request_response/sign_up"
	"net/http"
)

// UserSignUpHandler handles user registration
func (c *Controller) UserSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var req signup.UserSignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	response, err := c.service.UserSignUp(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(response)
}

// SysAdminSignUpHandler handles system admin registration
func (c *Controller) SysAdminSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var req signup.SysAdminSignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	response, err := c.service.SysAdminSignUp(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(response)
}

// UserSignInHandler handles user login
func (c *Controller) UserSignInHandler(w http.ResponseWriter, r *http.Request) {
	var req signin.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	response, err := c.service.SignIn(ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(response)
}

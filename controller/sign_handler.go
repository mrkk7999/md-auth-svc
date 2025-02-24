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
		c.log.WithError(err).Error("Failed to decode user sign-up request")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	ctx := context.Background()
	response, err := c.service.UserSignUp(ctx, req)
	if err != nil {
		c.log.WithError(err).Error("User sign-up failed")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	c.log.Info("User sign-up successful")
	json.NewEncoder(w).Encode(response)
}

// SysAdminSignUpHandler handles system admin registration
func (c *Controller) SysAdminSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var req signup.SysAdminSignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.log.WithError(err).Error("Failed to decode system admin sign-up request")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}
	ctx := context.Background()
	response, err := c.service.SysAdminSignUp(ctx, req)
	if err != nil {
		c.log.WithError(err).Error("System admin sign-up failed")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	c.log.Info("System admin sign-up successful")
	json.NewEncoder(w).Encode(response)
}

// UserSignInHandler handles user login
func (c *Controller) UserSignInHandler(w http.ResponseWriter, r *http.Request) {
	var req signin.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.log.WithError(err).Error("Failed to decode user sign-in request")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}
	ctx := context.Background()
	response, err := c.service.SignIn(ctx, &req)
	if err != nil {
		c.log.WithError(err).Error("User sign-in failed")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	c.log.Info("User sign-in successful")
	json.NewEncoder(w).Encode(response)
}

// SignOutHandler handles global signout
func (c *Controller) SignOutHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		c.log.Error("Missing Authorization header")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing Authorization header"})
		return
	}

	// Ensure the token is in "Bearer <token>" format
	if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		c.log.Error("Invalid Authorization format")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid Authorization format"})
		return
	}
	accessToken := authHeader[7:]

	ctx := context.Background()
	err := c.service.SignOut(ctx, accessToken)
	if err != nil {
		c.log.WithError(err).Error("Sign-out failed")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	c.log.Info("Sign-out successful")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sign-Out successful"})
}

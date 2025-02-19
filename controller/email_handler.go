package controller

import (
	"context"
	"encoding/json"
	emailverification "md-auth-svc/request_response/email_verification"
	"net/http"
)

// UserEmailVerificationHandler handles verifies user
func (c *Controller) UserEmailVerificationHandler(w http.ResponseWriter, r *http.Request) {
	var req emailverification.VerifyEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	response, err := c.service.EmailVerification(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(response)
}

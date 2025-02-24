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
		c.log.Error("Invalid request payload: ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	ctx := context.Background()
	response, err := c.service.EmailVerification(ctx, req)
	if err != nil {
		c.log.Error("Email verification failed: ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	c.log.Info(req.Username, "Email verified successfully!")
	json.NewEncoder(w).Encode(response)
}

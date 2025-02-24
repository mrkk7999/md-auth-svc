package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllTenantUsersHandler handles user registration
func (c *Controller) GetAllTenantUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	response, err := c.service.GetAllTenantUsers(ctx)
	if err != nil {
		c.log.WithError(err).Error("Failed to get all tenant users")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if len(response) == 0 {
		c.log.Info("No records found for all tenant users")
		json.NewEncoder(w).Encode(map[string]string{"message": "No records found"})
		return
	}
	c.log.Info("Successfully retrieved all tenant users")
	json.NewEncoder(w).Encode(response)
}

// GetTenantUserByIDHandler handles user retrieval by ID
func (c *Controller) GetTenantUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		c.log.Error("Empty ID provided for GetTenantUserByIDHandler")
		encodeJSONResponse(w, http.StatusBadRequest, nil, errors.New("empty id"))
		return
	}

	ctx := context.Background()
	response, err := c.service.GetTenantUserByID(ctx, id)
	if err != nil {
		c.log.WithError(err).Errorf("Failed to get tenant user by ID: %s", id)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	c.log.Infof("Successfully retrieved tenant user by ID: %s", id)
	json.NewEncoder(w).Encode(response)
}

// GetUsersByTenantHandler handles user retrieval by tenant ID
func (c *Controller) GetUsersByTenantHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		c.log.Error("Empty ID provided for GetUsersByTenantHandler")
		encodeJSONResponse(w, http.StatusBadRequest, nil, errors.New("empty id"))
		return
	}

	ctx := context.Background()
	response, err := c.service.GetUsersByTenantID(ctx, id)
	if err != nil {
		c.log.WithError(err).Errorf("Failed to get users by tenant ID: %s", id)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if len(response) == 0 {
		c.log.Infof("No records found for users by tenant ID: %s", id)
		json.NewEncoder(w).Encode(map[string]string{"message": "No records found"})
		return
	}

	c.log.Infof("Successfully retrieved users by tenant ID: %s", id)
	json.NewEncoder(w).Encode(response)
}

// GetAdminUsersByTenantHandler handles admin user retrieval by tenant ID
func (c *Controller) GetAdminUsersByTenantHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		c.log.Error("Empty ID provided for GetAdminUsersByTenantHandler")
		encodeJSONResponse(w, http.StatusBadRequest, nil, errors.New("empty id"))
		return
	}

	ctx := context.Background()
	response, err := c.service.GetAdminByTenantID(ctx, id)
	if err != nil {
		c.log.WithError(err).Errorf("Failed to get admin users by tenant ID: %s", id)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if len(response) == 0 {
		c.log.Infof("No records found for admin users by tenant ID: %s", id)
		json.NewEncoder(w).Encode(map[string]string{"message": "No records found"})
		return
	}

	c.log.Infof("Successfully retrieved admin users by tenant ID: %s", id)
	json.NewEncoder(w).Encode(response)
}

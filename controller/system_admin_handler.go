package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllSysAdminsHandler retrieves all system admins
func (c *Controller) GetAllSysAdminsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	admins, err := c.service.GetAllSysAdmins(ctx)
	if err != nil {
		c.log.WithError(err).Error("Failed to fetch system admins")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	c.log.Info("Fetched all system admins successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(admins)
}

// GetSysAdminByIDHandler retrieves a system admin by ID
func (c *Controller) GetSysAdminByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminID := vars["id"]

	if adminID == "" {
		c.log.Error("Missing admin ID in request")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Admin ID is required"})
		return
	}

	ctx := context.Background()
	admin, err := c.service.GetSysAdminByID(ctx, adminID)
	if err != nil {
		c.log.WithError(err).Error("Failed to fetch system admin")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if admin == nil {
		c.log.Info("System admin not found")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "System admin not found"})
		return
	}

	c.log.Info("Fetched system admin successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(admin)
}

// // UpdateSysAdminHandler updates a system admin
// func (c *Controller) UpdateSysAdminHandler(w http.ResponseWriter, r *http.Request) {
// 	var req sysadminuser.SysAdmin
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		c.log.WithError(err).Error("Failed to decode system admin update request")
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
// 		return
// 	}

// 	ctx := context.Background()
// 	err := c.service.UpdateSysAdmin(ctx, &req)
// 	if err != nil {
// 		c.log.WithError(err).Error("Failed to update system admin")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
// 		return
// 	}

// 	c.log.Info("System admin updated successfully")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "System admin updated successfully"})
// }

// // DeleteSysAdminHandler deletes a system admin
// func (c *Controller) DeleteSysAdminHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	adminID := vars["id"]

// 	if adminID == "" {
// 		c.log.Error("Missing admin ID in request")
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(map[string]string{"error": "Admin ID is required"})
// 		return
// 	}

// 	ctx := context.Background()
// 	err := c.service.DeleteSysAdmin(ctx, "", adminID) // Username is not required for deletion
// 	if err != nil {
// 		c.log.WithError(err).Error("Failed to delete system admin")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
// 		return
// 	}

// 	c.log.Info("System admin deleted successfully")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "System admin deleted successfully"})
// }

package controller

// func (c *Controller) MakeTenantUserAdminHandler(w http.ResponseWriter, r *http.Request) {
// 	req := usergroup.UserToGroupRequest{}
// 	// Parse the request body to fill the req object
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		c.log.Error("Failed to decode request body: ", err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
// 		return
// 	}
// 	if req.OldGroup == "admin" {
// 		c.log.Error("User is already an admin")
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(map[string]string{"error": "User is already an admin"})
// 		return
// 	}

// 	req.NewGroup = "admin"

// 	err = c.service.ChangeGroup(r.Context(), req)
// 	if err != nil {
// 		c.log.Error("Failed to change group: ", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
// }

// func (c *Controller) MakeTenantAdminUserHandler(w http.ResponseWriter, r *http.Request) {
// 	req := usergroup.UserToGroupRequest{}
// 	// Parse the request body to fill the req object
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		c.log.Error("Failed to decode request body: ", err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
// 		return
// 	}

// 	if req.OldGroup == "user" {
// 		c.log.Error("User is already a user")
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(map[string]string{"error": "User is already a user"})
// 		return
// 	}

// 	req.NewGroup = "user"

// 	err = c.service.ChangeGroup(r.Context(), req)
// 	if err != nil {
// 		c.log.Error("Failed to change group: ", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
// }

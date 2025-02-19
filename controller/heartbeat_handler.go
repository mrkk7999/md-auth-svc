package controller

import "net/http"

// HeartBeatHandler handles the heartbeat requests.
func (c *Controller) HeartBeatHandler(w http.ResponseWriter, r *http.Request) {
	response := c.service.HeartBeat()
	encodeJSONResponse(w, http.StatusOK, response, nil)
}

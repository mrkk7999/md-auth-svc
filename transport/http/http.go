package http

import (
	"md-auth-svc/controller"
	"md-auth-svc/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpRouter(controller *controller.Controller) http.Handler {
	var (
		router = mux.NewRouter()

		subRouter = router.PathPrefix("/auth/api/v1").Subrouter()
	)
	// Logging Middleware
	subRouter.Use(middleware.LoggingMiddleware)
	// Heartbeat Route
	subRouter.HandleFunc("/heartbeat", controller.HeartBeatHandler).Methods("GET")
	// User Routes
	subRouter.HandleFunc("/sign-up/user", controller.UserSignUpHandler).Methods("POST")
	subRouter.HandleFunc("/validate-email/user", controller.UserEmailVerificationHandler).Methods("POST")
	subRouter.HandleFunc("/sign-in/user", controller.UserSignInHandler).Methods("POST")

	// SysAdmin Routes
	subRouter.HandleFunc("/sign-up/sysadmin", controller.SysAdminSignUpHandler).Methods("POST")

	// subRouter.HandleFunc("/refresh-token", handlers.RefreshTokenHandler).Methods("POST")
	// subRouter.HandleFunc("/forgot-password", handlers.ForgotPasswordHandler).Methods("POST")
	// subRouter.HandleFunc("/reset-password", handlers.ResetPasswordHandler).Methods("POST")
	// subRouter.HandleFunc("/validate-token", handlers.ValidateTokenHandler).Methods("POST")

	return router
}

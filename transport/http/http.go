package http

import (
	"md-auth-svc/controller"
	"md-auth-svc/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func SetUpRouter(controller *controller.Controller, log *logrus.Logger) http.Handler {
	var (
		router = mux.NewRouter()

		subRouter = router.PathPrefix("/auth/api/v1").Subrouter()
	)
	// Logging Middleware

	subRouter.Use(func(next http.Handler) http.Handler {
		return middleware.LoggingMiddleware(next, log)
	})

	// Heartbeat Route
	subRouter.HandleFunc("/heartbeat", controller.HeartBeatHandler).Methods("GET")
	// User Routes
	subRouter.HandleFunc("/sign-up/user", controller.UserSignUpHandler).Methods("POST")
	subRouter.HandleFunc("/validate-email/user", controller.UserEmailVerificationHandler).Methods("POST")
	subRouter.HandleFunc("/sign-in/user", controller.UserSignInHandler).Methods("POST")
	subRouter.HandleFunc("/sign-out", controller.SignOutHandler).Methods("POST")
	subRouter.HandleFunc("/tenant/users/all", controller.GetAllTenantUsersHandler).Methods("GET")
	subRouter.HandleFunc("/tenant/users/{id}", controller.GetTenantUserByIDHandler).Methods("GET")
	subRouter.HandleFunc("/tenant/{id}/users", controller.GetUsersByTenantHandler).Methods("GET")
	subRouter.HandleFunc("/tenant/{id}/admins", controller.GetAdminUsersByTenantHandler).Methods("GET")
	// subRouter.HandleFunc("/tenant/user/make-admin", controller.MakeTenantUserAdminHandler).Methods("POST")
	// subRouter.HandleFunc("/tenant/admin/make-user", controller.MakeTenantAdminUserHandler).Methods("POST")

	// SysAdmin Routes
	subRouter.HandleFunc("/sign-up/sysadmin", controller.SysAdminSignUpHandler).Methods("POST")
	subRouter.HandleFunc("/sysadmins/all", controller.GetAllSysAdminsHandler).Methods("GET")
	subRouter.HandleFunc("/sysadmins/{id}/get", controller.GetSysAdminByIDHandler).Methods("GET")

	return router
}

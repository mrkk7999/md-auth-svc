package main

import (
	"fmt"
	"md-auth-svc/controller"
	"md-auth-svc/implementation"
	"md-auth-svc/repository"
	"md-auth-svc/services/cognito"
	httpTransport "md-auth-svc/transport/http"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// Logger
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		httpAddr = os.Getenv("HTTP_ADDR")

		dbHost     = os.Getenv("DB_HOST")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbName     = os.Getenv("DB_NAME")
		dbPort     = os.Getenv("DB_PORT")
		dbSSLMode  = os.Getenv("DB_SSLMODE")
		dbTimeZone = os.Getenv("DB_TIMEZONE")

		cogUserPoolID = os.Getenv("COG_USER_POOL_ID")
		cogClientID   = os.Getenv("COG_CLIENT_ID")
		cogRegion     = os.Getenv("COG_REGION")
		cogSecret     = os.Getenv("COG_SECRET")
	)

	// PostgreSQL DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbTimeZone)

	// Connect to Database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	repo := repository.New(db)

	cognitoClient, err := cognito.NewCognitoClient(cogUserPoolID, cogClientID, cogSecret, cogRegion)
	if err != nil {
		log.Warn("Failed to initialise cognito client", err)
	}
	log.Info("Cognito client initialized")

	cognito := cognito.NewAuthService(cognitoClient)

	svc := implementation.New(repo, cognito)

	controller := controller.New(svc, log)

	handler := httpTransport.SetUpRouter(controller, log)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Info("Server is running on ", httpAddr)

	go func() {
		server := &http.Server{
			Addr:    httpAddr,
			Handler: handler,
		}
		errs <- server.ListenAndServe()
	}()

	log.Error("exit", <-errs)
}

package main

import (
	"fmt"
	"log"
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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

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
		// export AWS_ACCESS_KEY_ID=AKIAXTKSIGC46ZDFJYPP
		// export AWS_SECRET_ACCESS_KEY=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
		// export AWS_REGION=eu-north-1
		// awsKeyID        = os.Getenv("AWS_ACCESS_KEY_ID")
		// awsSecretAccess = os.Getenv("AWS_SECRET_ACCESS_KEY")
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
		log.Fatal("Failed to initialise cognito client", err)
	}

	cognito := cognito.NewAuthService(cognitoClient)

	svc := implementation.New(repo, cognito)

	controller := controller.New(svc)

	handler := httpTransport.SetUpRouter(controller)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println("Server is running " + httpAddr)

	go func() {
		server := &http.Server{
			Addr:    httpAddr,
			Handler: handler,
		}
		errs <- server.ListenAndServe()
	}()

	log.Println("exit", <-errs)
}

package cognito

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// CognitoClient
type CognitoClient struct {
	Client       *cognitoidentityprovider.Client
	UserPoolID   string
	ClientID     string
	ClientSecret string
}

// NewCognitoClient initializes the Cognito client with IAM credentials
func NewCognitoClient(userPoolID, clientID, clientSecret, region string) (*CognitoClient, error) {
	// log.Println("aws key:", awsKeyID, "aws secret access:", awsSecretAccess)

	aws_access_key_id := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			aws_access_key_id, aws_secret_access_key, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %v", err)
	}

	client := cognitoidentityprovider.NewFromConfig(cfg)

	log.Println("Cognito client initialized with UserPoolID:", userPoolID, "Region:", region)

	return &CognitoClient{
		Client:       client,
		UserPoolID:   userPoolID,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}, nil
}

// // CognitoClient
// type CognitoClient struct {
// 	Client       *cognitoidentityprovider.Client
// 	UserPoolID   string
// 	ClientID     string
// 	ClientSecret string
// }

// // NewCognitoClient
// func NewCognitoClient(userPoolID, clientID, clientSecret, region string) (*CognitoClient, error) {
// 	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load AWS config: %v", err)
// 	}

// 	client := cognitoidentityprovider.NewFromConfig(cfg)

// 	log.Println("Cognito client initialized with UserPoolID:", userPoolID, "Region:", region)

// 	return &CognitoClient{
// 		Client:       client,
// 		UserPoolID:   userPoolID,
// 		ClientID:     clientID,
// 		ClientSecret: clientSecret,
// 	}, nil
// }

func (c *CognitoClient) GenerateSecretHash(username string) string {
	h := hmac.New(sha256.New, []byte(c.ClientSecret))
	h.Write([]byte(username + c.ClientID))
	secretHash := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return secretHash
}

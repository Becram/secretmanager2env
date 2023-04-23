package aws

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func CreateSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String("eu-west-1")},
	}))
	return sess
}

func GetSecret(secretName *string, secretVersion *string) {
	svc := secretsmanager.New(CreateSession())
	var versionID string
	if *secretVersion == "version" {
		versionID = "AWSCURRENT"
	} else {
		versionID = *secretVersion
	}
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(*secretName),
		VersionStage: aws.String(versionID),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeDecryptionFailure:
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	var secrets map[string]string
	error := json.Unmarshal([]byte(*result.SecretString), &secrets)
	if error != nil {
		panic(error)
	}

	for key, value := range secrets {
		// Each value is an `any` type, that is type asserted as a string
		os.Setenv(key, value)
		if os.Getenv("ENV_DEBUG") == "true" {
			fmt.Printf("Setting Variable :   %s = %s***\n", key, value[0:4])
		} else {
			fmt.Printf("Setting Variable :   %s = ******\n", key)
		}
	}

}

package secrets

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func ReadSecret(name string, region string) (secretValue string, err error) {
	Session, err := Connect(region)
	if err != nil {
		fmt.Println(err)
	}
	sm := secretsmanager.New(Session)

	value, err := sm.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	})

	if err != nil {
		fmt.Println(err)
	}
	return *value.SecretString, err
}

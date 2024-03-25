package secrets

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func Connect(region string) (Session *session.Session, err error) {
	Session, err = session.NewSession(
		&aws.Config{
			Region: aws.String(region),
		},
	)
	return Session, err
}

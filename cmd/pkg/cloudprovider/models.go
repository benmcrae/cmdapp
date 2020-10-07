package cloudprovider

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type Clients struct {
	s3 s3iface.S3API
}

func NewClients() *Clients {
	sess := session.Must(
		session.NewSession(
			&aws.Config{
				Region:      aws.String("us-west-2"),
				Credentials: credentials.NewSharedCredentials("", "hcom-tech-ops"),
			},
		),
	)

	s3client := s3.New(sess)

	return &Clients{
		s3: s3client,
	}
}

type S3Buckets struct {
	BucketName string
}

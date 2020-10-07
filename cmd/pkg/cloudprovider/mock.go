package cloudprovider

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type MockS3Client struct {
	s3iface.S3API
}

func NewMockClients() *Clients {
	return &Clients{
		s3: &MockS3Client{},
	}
}

func (s3mock *MockS3Client) ListBuckets(*s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {

	return &s3.ListBucketsOutput{
		Buckets: []*s3.Bucket{
			{
				Name: aws.String("S3 Bucket One"),
			},
			{
				Name: aws.String("S3 Bucket Two"),
			},
		},
	}, nil
}

package cloudprovider

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (c *Clients) ListS3Buckets() []S3Buckets {
	var buckets []S3Buckets

	results, err := c.s3.ListBuckets(&s3.ListBucketsInput{})

	if err != nil {
		log.Fatal("There was an error retrieving the buckets")
	}

	for _, b := range results.Buckets {
		buckets = append(buckets, S3Buckets{
			BucketName: aws.StringValue(b.Name),
		})
	}

	return buckets
}

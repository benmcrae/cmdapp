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

func (c *Clients) GetS3Objects(bucketName string) []S3Object {
	var s3objects []S3Object
	results, err := c.s3.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		log.Fatal("There was an error retrieving the bucket's objects")
	}

	for _, o := range results.Contents {
		s3objects = append(s3objects, S3Object{
			ObjectName:   aws.StringValue(o.Key),
			LastModified: aws.TimeValue(o.LastModified),
			SizeMB:       float32(aws.Int64Value(o.Size) / 1000000),
			StorageClass: aws.StringValue(o.StorageClass),
		})
	}

	return s3objects
}

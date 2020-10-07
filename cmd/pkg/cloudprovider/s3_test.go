package cloudprovider

import (
	"reflect"
	"testing"
)

func TestListBuckets(t *testing.T) {
	TestingClient := NewMockClients()

	buckets := TestingClient.ListS3Buckets()

	expectedB := []S3Buckets{
		{
			BucketName: "S3 Bucket One",
		},
		{
			BucketName: "S3 Bucket Two",
		},
	}

	if len(buckets) != 2 {
		t.Errorf("Expected 2 buckets, got %v", len(buckets))
	}

	if !reflect.DeepEqual(expectedB, buckets) {
		t.Errorf("Wrong buckets returned.")
	}
}

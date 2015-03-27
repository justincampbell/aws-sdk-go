// +build integration

package s3_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

var bucketName *string
var svc *s3.S3

func TestMain(m *testing.M) {
	setup()
	defer teardown() // only called if we panic
	result := m.Run()
	teardown()
	os.Exit(result)
}

// Create a bucket for testing
func setup() {
	svc = s3.New(nil)
	bucketName = aws.String(
		fmt.Sprintf("aws-sdk-go-integration-%d", time.Now().Unix()))

	for i := 0; i < 10; i++ {
		_, err := svc.CreateBucket(&s3.CreateBucketInput{Bucket: bucketName})
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
}

// Delete the bucket
func teardown() {
	resp, _ := svc.ListObjects(&s3.ListObjectsInput{Bucket: bucketName})
	for _, o := range resp.Contents {
		svc.DeleteObject(&s3.DeleteObjectInput{Bucket: bucketName, Key: o.Key})
	}
	svc.DeleteBucket(&s3.DeleteBucketInput{Bucket: bucketName})
}

func TestWriteToObject(t *testing.T) {
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: bucketName,
		Key:    aws.String("key name"),
		Body:   bytes.NewReader([]byte("hello world")),
	})
	assert.NoError(t, err)

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: bucketName,
		Key:    aws.String("key name"),
	})
	assert.NoError(t, err)

	b, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, []byte("hello world"), b)
}
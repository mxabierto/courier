package utils

import (
	"strings"
	"os"
  "io/ioutil"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

var s3BucketURL = "https://%s.s3.amazonaws.com%s"

// TestS3 tests whether the passed in s3 client is properly configured and the passed in bucket is accessible
func TestS3(s3Client s3iface.S3API, bucket string) error {
	params := &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	}
	_, err := s3Client.HeadBucket(params)
	if err != nil {
		return err
	}

	return nil
}

// PutS3File writes the passed in file to the bucket with the passed in content type
func PutS3File(s3Client s3iface.S3API, bucket string, path string, contentType string, contents []byte) (string, error) {
	result := strings.Split(path, "/");
  folder := strings.Join(result[:len(result)-1],"/");
  os.MkdirAll(folder, os.ModePerm);
	ioutil.WriteFile(path, contents, 0644);
	url := "https://rapidpro.datos.gob.mx"+ path;
	return url, nil
}

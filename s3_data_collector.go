package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
)

type S3DataCollector struct {
	Params *s3.GetObjectInput
	S3     *s3.S3
}

func NewS3DataCollector(s3Service *s3.S3, params *s3.GetObjectInput) (d DataCollector) {
	d = &S3DataCollector{
		S3:     s3Service,
		Params: params,
	}
	return d
}

func NewDefaultS3DataCollector() (d DataCollector) {
	d = &S3DataCollector{
		S3: s3.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")}),
		Params: &s3.GetObjectInput{
			Bucket: aws.String(s3Bucket),
			Key:    aws.String(lastUpdatedKey),
		},
	}
	return d
}

func (s *S3DataCollector) Collect() (data []byte, err error) {
	response, err := s.S3.GetObject(s.Params)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	data, err = ioutil.ReadAll(response.Body)
	return
}

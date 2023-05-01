package cloudpolysdk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Store struct {
	Aliyun *aliyun
	AWS    *_aws
	Err    error
}

type aliyun struct {
	s3     *s3.Client
	config *Config
	common
}

type _aws struct {
	s3     *s3.Client
	config *Config
	common
}

var (
	_ objecter = (*aliyun)(nil)
	_ objecter = (*_aws)(nil)
)

func newSession(config *Config) (*s3.Client, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               config.url(),
			SigningRegion:     config.Region,
			HostnameImmutable: config.hostnameImmutable(),
		}, nil
	})
	creds := credentials.NewStaticCredentialsProvider(config.AccessKeyId, config.AccessSecret, "")
	cfg, err := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithCredentialsProvider(creds),
		awsConfig.WithEndpointResolverWithOptions(customResolver),
		awsConfig.WithRegion(config.Region),
	)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}

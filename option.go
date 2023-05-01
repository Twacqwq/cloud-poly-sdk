package cloudpolysdk

import (
	"errors"
	"fmt"
	"reflect"
)

type Config struct {
	AccessKeyId       string `required:"true"`
	AccessSecret      string `required:"true"`
	Region            string `required:"true"`
	Bucket            string `required:"true"`
	EndPoint          string `required:"true"`
	UseSSL            bool
	HostnameImmutable *bool
}

func (cm *Config) validate() bool {
	typ := reflect.TypeOf(cm).Elem()
	val := reflect.ValueOf(cm).Elem()
	for i := 0; i < typ.NumField(); i++ {
		if val.Field(i).IsZero() && typ.Field(i).Tag.Get("required") == "true" {
			return false
		}
	}

	return true
}

func (cm *Config) url() string {
	if cm.UseSSL {
		return fmt.Sprintf("https://%s", cm.EndPoint)
	}

	return fmt.Sprintf("http://%s", cm.EndPoint)
}

func (cm *Config) hostnameImmutable() bool {
	if cm.HostnameImmutable != nil {
		return *cm.HostnameImmutable
	}

	return false
}

type Option func(conf *Store)

func WithProvide(provide int, config *Config) Option {
	return func(store *Store) {
		if !config.validate() {
			store.Err = errors.New("the configuration information is incomplete")
			return
		}

		switch provide {
		case AWS:
			store.AWS = &_aws{
				config: &Config{
					AccessKeyId:  config.AccessKeyId,
					AccessSecret: config.AccessSecret,
					Bucket:       config.Bucket,
					EndPoint:     config.Bucket,
					Region:       config.Region,
					UseSSL:       config.UseSSL,
				},
			}
			store.AWS.s3, store.Err = newSession(store.AWS.config)
			if store.Err != nil {
				store.Err = fmt.Errorf("aws s3 client initialization failure, %v", store.Err.Error())
			}
		case Aliyun:
			store.Aliyun = &aliyun{
				config: &Config{
					AccessKeyId:  config.AccessKeyId,
					AccessSecret: config.AccessSecret,
					Bucket:       config.Bucket,
					EndPoint:     config.Bucket,
					Region:       config.Region,
					UseSSL:       config.UseSSL,
				},
			}
			store.Aliyun.s3, store.Err = newSession(store.Aliyun.config)
			if store.Err != nil {
				store.Err = fmt.Errorf("aliyun s3 client initialization failure, %v", store.Err.Error())
			}
		}
	}
}

package option

import "reflect"

type Config struct {
	AccessKeyId  string
	AccessSecret string
	Region       string
	Bucket       string
	EndPoint     string
}

func (cm *Config) validate() bool {
	val := reflect.ValueOf(cm).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).IsZero() {
			return false
		}
	}

	return true
}

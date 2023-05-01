package cloudpolysdk

import "io"

type objecter interface {
	// Put upload object
	Put(key string, r io.Reader) error
	// Get get object information
	Get()
}

type common struct {
	objecter
}

func (c *common) Put(key string, r io.Reader) error {
	println("PUT")

	return nil
}

func (c *common) Get() {
	println("GET")
}

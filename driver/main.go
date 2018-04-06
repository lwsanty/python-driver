package main

import (
	_ "github.com/bblfsh/python-driver/driver/impl"
	"github.com/bblfsh/python-driver/driver/normalizer"

	"gopkg.in/bblfsh/sdk.v1/sdk/driver"
)

func main() {
	driver.Run(driver.Transforms{
		Native: normalizer.Native,
		Code:   normalizer.Code,
	})
}

package tests

import (
	"github.com/goravel/framework/testing"

	"rm/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}

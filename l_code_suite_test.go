package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LCode Suite")
}

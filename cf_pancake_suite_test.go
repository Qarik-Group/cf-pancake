package cf_pancake_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCfPancake(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CfPancake Suite")
}

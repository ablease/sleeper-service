package sleeperservice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSleeperService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SleeperService Suite")
}

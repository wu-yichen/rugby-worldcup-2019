package helpers_test

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/subosito/gotenv"
)

func TestHelpers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helpers Suite")
}

var _ = BeforeSuite(func() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
})

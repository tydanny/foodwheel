package cuisine_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2" //nolint:revive
	. "github.com/onsi/gomega"    //nolint:revive
)

func TestCuisine(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cuisine Suite")
}

package inmemory_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	cuisinesv1 "github.com/tydanny/foodwheel/gen/foodwheel/cuisines/v1"
	"github.com/tydanny/foodwheel/internal/app/foodwheel/cuisines"
	"github.com/tydanny/foodwheel/internal/app/foodwheel/cuisines/cuisineadapter/inmemory"
)

var _ = Describe("Inmemory", func() {
	var service cuisines.Service

	BeforeEach(func() {
		service = inmemory.NewStore()
	})

	Context("Create", func(ctx SpecContext) {
		It("should create a Cuisine", func() {
			cuisine := &cuisinesv1.Cuisine{
				Name:        "Italian",
				Description: "Food from Italy",
			}

			_, err := service.Create(ctx, cuisine)
			Expect(err).ToNot(HaveOccurred())
		})
	})
	Context("Get", func(ctx SpecContext) {})
	Context("List", func(ctx SpecContext) {})
	Context("Update", func(ctx SpecContext) {})
	Context("Delete", func(ctx SpecContext) {})
})

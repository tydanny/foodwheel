package cuisine_test

import (
	. "github.com/onsi/ginkgo/v2" //nolint:revive
	. "github.com/onsi/gomega"    //nolint:revive
	"github.com/tydanny/foodwheel/pkg/cuisine"
)

var _ = Describe("Dish", func() {
	var dish cuisine.Dish

	BeforeEach(func() {
		dish = cuisine.NewDish(
			"Burgers",
			"beef",
			"grilled",
		)
	})

	Describe("HasTag", func() {
		It("Should return true when the tag exists", func() {
			Expect(dish.HasTag("beef")).To(BeTrue())
		})

		It("Should return false when the tag doesn't exist", func() {
			Expect(dish.HasTag("chicken")).To(BeFalse())
		})
	})
})

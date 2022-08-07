package foodwheel_test

import (
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tydanny/foodwheel/pkg/foodwheel"
)

var _ = Describe("Foodwheel API", func() {

	Describe("cuisenes endpoint", func() {

		Context("get all cuisines", func() {

			It("should have 4 elements", func() {
				stream, err := client.GetCuisines(ctx, &foodwheel.Empty{})
				Expect(err).NotTo(HaveOccurred())
				var cuisines []*foodwheel.Cuisine
				for {
					c, recErr := stream.Recv()
					if recErr == io.EOF {
						break
					}
					Expect(recErr).ToNot(HaveOccurred())
					cuisines = append(cuisines, c)
				}
				Expect(cuisines).Should(HaveLen(4))
			})

		})

		Context("get cuisine by name", func() {

			It("should get the correct cuisine", func() {
				naReq := &foodwheel.CuisineRequest{
					Name: "North_American",
				}
				naCuisine, getErr := client.GetCuisineByName(ctx, naReq)
				Expect(getErr).NotTo(HaveOccurred())

				Expect(naCuisine.Name).To(Equal("North_American"))
				Expect(naCuisine.Dishes[0]).To(Equal("Burgers"))
				Expect(naCuisine.Dishes[1]).To(Equal("Fried Chicken"))
			})

			It("should fail to return a cuisine that doesn't exist", func() {
				wrongReq := &foodwheel.CuisineRequest{
					Name: "not_real",
				}
				wrongCuisine, getErr := client.GetCuisineByName(ctx, wrongReq)
				Expect(getErr).To(HaveOccurred())
				Expect(getErr.Error()).To(Equal("rpc error: code = Unknown desc = requested cuisine \"not_real\" not found"))
				Expect(wrongCuisine).To(BeNil())
			})

		})

		Context("post new cuisine and get it by name", func() {

			It("should create a new cuisine entry", func() {
				newCuisine := &foodwheel.Cuisine{Name: "Korean", Dishes: []string{"BBQ Pork", "Kimchi"}}
				respCuisine, addErr := client.AddCuisine(ctx, newCuisine)
				Expect(addErr).ShouldNot(HaveOccurred())
				Expect(respCuisine).To(HaveField("Name", Equal("Korean")))
				Expect(respCuisine.GetDishes()).Should(
					HaveLen(2),
					ContainElements("BBQ Pork", "Kimchi"))
			})

			It("should be available for query", func() {
				kReq := foodwheel.CuisineRequest{
					Name: "Korean",
				}
				kCuisine, getErr := client.GetCuisineByName(ctx, &kReq)
				Expect(getErr).ShouldNot(HaveOccurred())

				Expect(kCuisine.Name).To(Equal("Korean"))
				Expect(kCuisine.Dishes[0]).To(Equal("BBQ Pork"))
				Expect(kCuisine.Dishes[1]).To(Equal("Kimchi"))
			})

		})

	})

	Describe("spin endpoint", func() {

		Context("get random cuisine", func() {

			It("should return a random cuisine", func() {
				ranCuisine, spinErr := client.Spin(ctx, &foodwheel.Empty{})
				Expect(spinErr).ShouldNot(HaveOccurred())

				Expect(ranCuisine.Name).ToNot(BeEmpty())
				Expect(len(ranCuisine.Dishes)).To(BeNumerically(">", 0))
			})

		})

	})

})

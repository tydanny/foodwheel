package foodwheel_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	api "github.com/tydanny/foodwheel/pkg"
)

var _ = Describe("Foodwheel API", func() {

	Describe("cuisenes endpoint", func() {

		Context("get all cuisines", func() {

			It("should have 4 elements", func() {
				resp, err := getResponse("GET", "/cuisines", nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp).Should(HaveHTTPStatus(http.StatusOK))

				cuisines := api.Cuisines{}
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				Expect(err).ShouldNot(HaveOccurred())

				uMarshErr := json.Unmarshal(bodyBytes, &cuisines)
				Expect(uMarshErr).NotTo(HaveOccurred())
				Expect(cuisines).To(HaveLen(4))
			})

		})

		Context("get cuisine by name", func() {

			It("should get the correct cuisine", func() {
				resp, err := getResponse("GET", "/cuisines/North_American", nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp).Should(HaveHTTPStatus(http.StatusOK))

				naCuisine := api.Cuisine{}
				bodyBytes, readErr := ioutil.ReadAll(resp.Body)
				Expect(readErr).NotTo(HaveOccurred())
				uMarshErr := json.Unmarshal(bodyBytes, &naCuisine)
				Expect(uMarshErr).NotTo(HaveOccurred())

				Expect(naCuisine.Name).To(Equal("North_American"))
				Expect(naCuisine.Dishes[0]).To(Equal("Burgers"))
				Expect(naCuisine.Dishes[1]).To(Equal("Fried Chicken"))
			})

			It("should fail to return a cuisine that doesn't exist", func() {
				resp, err := getResponse("GET", "/cuisines/not_real", nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp).To(HaveHTTPStatus(http.StatusNotFound))
			})

		})

		Context("post new cuisine and get it by name", func() {

			It("should create a new cuisine entry", func() {
				newCuisine := api.Cuisine{Name: "Korean", Dishes: []string{"BBQ Pork", "Kimchi"}}
				data, err := json.Marshal(newCuisine)
				Expect(err).ShouldNot(HaveOccurred())
				resp, err := getResponse("POST", "/cuisines", strings.NewReader(string(data)))
				Expect(err).ShouldNot(HaveOccurred())

				Expect(resp).Should(HaveHTTPStatus(http.StatusCreated))
			})

			It("should be available for query", func() {
				resp, err := getResponse("GET", "/cuisines/Korean", nil)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).To(HaveHTTPStatus(http.StatusOK))

				kCuisine := api.Cuisine{}
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				Expect(err).ShouldNot(HaveOccurred())
				err = json.Unmarshal(bodyBytes, &kCuisine)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(kCuisine.Name).To(Equal("Korean"))
				Expect(kCuisine.Dishes[0]).To(Equal("BBQ Pork"))
				Expect(kCuisine.Dishes[1]).To(Equal("Kimchi"))
			})

		})

	})

	Describe("spin endpoint", func() {

		Context("get random cuisine", func() {

			It("should return a random cuisine", func() {
				resp, err := getResponse("GET", "/spin", nil)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).To(HaveHTTPStatus(http.StatusOK))

				ranCuisine := api.Cuisine{}
				bodyBytes, _ := ioutil.ReadAll(resp.Body)
				uMarshErr := json.Unmarshal(bodyBytes, &ranCuisine)
				Expect(uMarshErr).NotTo(HaveOccurred())

				Expect(ranCuisine.Name).ToNot(BeEmpty())
				Expect(len(ranCuisine.Dishes)).To(BeNumerically(">", 0))
			})

		})

	})

})

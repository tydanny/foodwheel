package v1alpha1_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	api "github.com/tydanny/foodwheel/v1alpha1"
)

var router *gin.Engine

func TestV1alpha1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "V1alpha1 Suite")
}

var _ = BeforeSuite(func() {
	// set gin to test mode
	gin.SetMode(gin.TestMode)
	router = gin.Default()
	Expect(router).ToNot(BeNil())

	api.InitializeRoutes(router)

	// TODO: initialize mongodb database
})

func getResponse(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to build new http request: %w", err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	return recorder.Result(), nil
}

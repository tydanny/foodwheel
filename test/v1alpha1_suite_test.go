package foodwheel_test

import (
	"context"
	"log"
	"net"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tydanny/foodwheel/pkg/foodwheel"
	"github.com/tydanny/foodwheel/pkg/fwServer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var (
	ctx        context.Context
	lis        *bufconn.Listener
	conn       *grpc.ClientConn
	client     foodwheel.FoodwheelClient
	testServer *grpc.Server
)

func TestV1alpha1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "V1alpha1 Suite")
}

var _ = BeforeSuite(func() {
	ctx = context.Background()
	var dialErr error
	conn, dialErr = grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	Expect(dialErr).ToNot(HaveOccurred())
	client = foodwheel.NewFoodwheelClient(conn)

	// TODO: initialize mongodb database
})

var _ = AfterSuite(func() {

})

func init() {
	lis = bufconn.Listen(bufSize)
	testServer = grpc.NewServer()
	foodwheel.RegisterFoodwheelServer(testServer, fwServer.ExampleServer())
	go func() {
		if servErr := testServer.Serve(lis); servErr != nil {
			log.Fatalf("server exited unexpectedly: %v", servErr)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

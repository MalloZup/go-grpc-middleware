package tokenbucket

import (
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	grpc_ratelimit "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/ratelimit"
	"github.com/juju/ratelimit"
	"google.golang.org/grpc"
)

// Here hard-coded for simplicity sake, but in your application you should make this configurable.
const (
	// Add 5 token per seconds.
	rate = 5
	// Capacity of bucket. allow only 40 requests.
	tokenCapacity = 40
)

// Simple example of server initialization code.
func Example() {

	limiter := TockenBucketInterceptor{}

	limiter.tokenBucket = ratelimit.NewBucket(rate, int64(tokenCapacity))

	_ = grpc.NewServer(
		middleware.WithUnaryServerChain(
			grpc_ratelimit.UnaryServerInterceptor(limiter),
		),
		middleware.WithStreamServerChain(
			grpc_ratelimit.UnaryServerInterceptor(limiter),
		),
	)
}

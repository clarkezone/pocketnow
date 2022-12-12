// Package geocacheservice is an implementation of the geocacheservice service.
package geocacheservice

import (
	context "context"

	clarkezoneLog "github.com/clarkezone/geocache/pkg/log"
)

// GeocacheServiceImpl is the server API for GreetingService service.
type GeocacheServiceImpl struct {
	UnimplementedGeoCacheServiceServer
}

// GetGreeting implements GreetingServer
func (s *GeocacheServiceImpl) SaveLocations(ctx context.Context, in *Locations) (*Empty, error) {
	//name := os.Getenv("MY_POD_NAME")
	clarkezoneLog.Debugf("SaveLocations called with ^v items", len(in.Locations))
	return &Empty{}, nil
}

// Package geocacheservice is an implementation of the geocacheservice service.
package geocacheservice

import (
	context "context"
	"fmt"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"
)

// GeocacheServiceImpl is the server API for GreetingService service.
type GeocacheServiceImpl struct {
	UnimplementedGeoCacheServiceServer
	last       *Location
	writeQueue *Queue
}

func NewGeoCacheServiceImpl() *GeocacheServiceImpl {
	clarkezoneLog.Debugf("NewGeoCacheServiceImpl")
	si := &GeocacheServiceImpl{}
	co := &CosmosDBWriter{}
	si.writeQueue = NewQueue(1000, co)
	si.writeQueue.Reader()
	return si
}

func (s *GeocacheServiceImpl) Done() {
	clarkezoneLog.Debugf("GeocacheServiceImpl: Done called")
	s.writeQueue.Close()
}

// GetGreeting implements GreetingServer
func (s *GeocacheServiceImpl) SaveLocations(ctx context.Context, in *Locations) (*Empty, error) {
	//name := os.Getenv("MY_POD_NAME")
	clarkezoneLog.Debugf("SaveLocations called with %v items", len(in.Locations))
	s.last = in.Locations[len(in.Locations)-1]
	return &Empty{}, nil
}

func (s *GeocacheServiceImpl) GetLastLocation(ctx context.Context, in *Empty) (*Location, error) {
	if s.last == nil {
		return nil, fmt.Errorf("no last location")
	}
	return s.last, nil
}

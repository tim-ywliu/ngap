package ngapType

import "github.com/free5gc/aper"

// Need to import "github.com/free5gc/aper" if it uses "aper"

const (
	RRCStatePresentInactive  aper.Enumerated = 0
	RRCStatePresentConnected aper.Enumerated = 1
)

type RRCState struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

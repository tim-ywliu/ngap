package ngapType

import "github.com/free5gc/aper"

// Need to import "github.com/free5gc/aper" if it uses "aper"

type PDUSessionResourceSetupItemCxtRes struct {
	PDUSessionID                            PDUSessionID
	PDUSessionResourceSetupResponseTransfer aper.OctetString
	IEExtensions                            *ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs `aper:"optional"`
}

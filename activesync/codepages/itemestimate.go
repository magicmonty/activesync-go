package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func ItemEstimate(protocolVersion byte) CodePage {
	result := NewCodePage("ItemEstimate", 6)

	result.AddTag("GetItemEstimate", 0x05)
	if protocolVersion <= PROTOCOL_VERSION_12_1 {
		result.AddTag("Version", 0x06) // only supported with MS-ASProtocolVersion 12.1
	}
	result.AddTag("Collections", 0x07)
	result.AddTag("Collection", 0x08)
	if protocolVersion <= PROTOCOL_VERSION_12_1 {
		result.AddTag("Class", 0x09) // only supported with MS-ASProtocolVersion 12.1
		// with 14.0 and 14.1 the Class tag from CodePage 0 (ActiveSync) is used
	}
	result.AddTag("CollectionId", 0x0A)
	if protocolVersion <= PROTOCOL_VERSION_12_1 {
		result.AddTag("DateTime", 0x0B) // only supported with MS-ASProtocolVersion 12.1
	}
	result.AddTag("Estimate", 0x0C)
	result.AddTag("Response", 0x0D)
	result.AddTag("Status", 0x0E)

	return result
}

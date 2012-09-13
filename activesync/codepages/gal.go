package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func GAL(protocolVersion byte) CodePage {
	result := NewCodePage("GAL", 16)

	result.AddTag("DisplayName", 0x05)
	result.AddTag("Phone", 0x06)
	result.AddTag("Office", 0x07)
	result.AddTag("Title", 0x08)
	result.AddTag("Company", 0x09)
	result.AddTag("Alias", 0x0A)
	result.AddTag("FirstName", 0x0B)
	result.AddTag("LastName", 0x0C)
	result.AddTag("HomePhone", 0x0D)
	result.AddTag("MobilePhone", 0x0E)
	result.AddTag("EmailAddress", 0x0F)
	if protocolVersion > PROTOCOL_VERSION_14_0 {
		result.AddTag("Picture", 0x10) // not supported with MS-ASProtocolVersion 12.1 or 14.0
		result.AddTag("Status", 0x11)  // not supported with MS-ASProtocolVersion 12.1 or 14.0
		result.AddTag("Data", 0x12)    // not supported with MS-ASProtocolVersion 12.1 or 14.0
	}

	return result
}

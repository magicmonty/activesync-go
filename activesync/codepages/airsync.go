package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func AirSync(protocolVersion byte) CodePage {
	result := NewCodePage("AirSync", 0)

	result.AddTag("Sync", 0x05)
	result.AddTag("Responses", 0x06)
	result.AddTag("Add", 0x07)
	result.AddTag("Change", 0x08)
	result.AddTag("Delete", 0x09)
	result.AddTag("Fetch", 0x0A)
	result.AddTag("SyncKey", 0x0B)
	result.AddTag("ClientId", 0x0C)
	result.AddTag("ServerId", 0x0D)
	result.AddTag("Status", 0x0E)
	result.AddTag("Collection", 0x0F)
	result.AddTag("Class", 0x10)
	result.AddTag("CollectionId", 0x12)
	result.AddTag("GetChanges", 0x13)
	result.AddTag("MoreAvailable", 0x14)
	result.AddTag("WindowSize", 0x15)
	result.AddTag("Commands", 0x16)
	result.AddTag("Options", 0x17)
	result.AddTag("FilterType", 0x18)
	result.AddTag("Conflict", 0x1B)
	result.AddTag("Collections", 0x1C)
	result.AddTag("ApplicationData", 0x1D)
	result.AddTag("DeletesAsMoves", 0x1E)
	result.AddTag("Supported", 0x20)
	result.AddTag("SoftDelete", 0x21)
	result.AddTag("MIMESupport", 0x22)
	result.AddTag("MIMETruncation", 0x23)
	result.AddTag("Wait", 0x24)
	result.AddTag("Limit", 0x25)
	result.AddTag("Partial", 0x26)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag("ConversationMode", 0x27)  // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("MaxItems", 0x28)          // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("HeartbeatInterval", 0x29) // not supported with MS-ASProtocolVersion 12.1
	}

	return result
}

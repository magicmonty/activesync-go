package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func Search(protocolVersion byte) CodePage {
	result := NewCodePage("Search", 15)

	result.AddTag("Search", 0x05)
	result.AddTag("Store", 0x07)
	result.AddTag("Name", 0x08)
	result.AddTag("Query", 0x09)
	result.AddTag("Options", 0x0A)
	result.AddTag("Range", 0x0B)
	result.AddTag("Status", 0x0C)
	result.AddTag("Response", 0x0D)
	result.AddTag("Result", 0x0E)
	result.AddTag("Properties", 0x0F)
	result.AddTag("Total", 0x10)
	result.AddTag("EqualTo", 0x11)
	result.AddTag("Value", 0x12)
	result.AddTag("And", 0x13)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag("Or", 0x14) // not supported with MS-ASProtocolVersion 12.1
	}
	result.AddTag("FreeText", 0x15)
	result.AddTag("DeepTraversal", 0x17)
	result.AddTag("LongId", 0x18)
	result.AddTag("RebuildResults", 0x19)
	result.AddTag("LessThan", 0x1A)
	result.AddTag("GreaterThan", 0x1B)
	result.AddTag("UserName", 0x1E)
	result.AddTag("Password", 0x1F)
	result.AddTag("ConversationId", 0x20)
	if protocolVersion > PROTOCOL_VERSION_14_0 {
		result.AddTag("Picture", 0x21)     // not supported with MS-ASProtocolVersion 12.1 or 14.0
		result.AddTag("MaxSize", 0x22)     // not supported with MS-ASProtocolVersion 12.1 or 14.0
		result.AddTag("MaxPictures", 0x23) // not supported with MS-ASProtocolVersion 12.1 or 14.0
	}

	return result
}

package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func ItemOperations(protocolVersion byte) CodePage {
	result := NewCodePage("ItemOperations", 20)

	result.AddTag("ItemOperations", 0x05)
	result.AddTag("Fetch", 0x06)
	result.AddTag("Store", 0x07)
	result.AddTag("Options", 0x08)
	result.AddTag("Range", 0x09)
	result.AddTag("Total", 0x0A)
	result.AddTag("Properties", 0x0B)
	result.AddTag("Data", 0x0C)
	result.AddTag("Status", 0x0D)
	result.AddTag("Response", 0x0E)
	result.AddTag("Version", 0x0F)
	result.AddTag("Schema", 0x10)
	result.AddTag("Part", 0x11)
	result.AddTag("EmptyFolderContents", 0x12)
	result.AddTag("DeleteSubFolders", 0x13)
	result.AddTag("UserName", 0x14)
	result.AddTag("Password", 0x15)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag("Move", 0x16)           // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("DstFldId", 0x17)       // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("ConversationId", 0x18) // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("MoveAlways", 0x19)     // not supported with MS-ASProtocolVersion 12.1
	}

	return result
}

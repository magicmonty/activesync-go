package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

// not supported with MS-ASProtocolVersion 12.1
func ComposeMail(protocolVersion byte) CodePage {
	result := NewCodePage("ComposeMail", 21)

	result.AddTag("SendMail", 0x05)
	result.AddTag("SmartForward", 0x06)
	result.AddTag("SmartReply", 0x07)
	result.AddTag("SaveInSentItems", 0x08)
	result.AddTag("ReplaceMime", 0x09)
	result.AddTag("Source", 0x0B)
	result.AddTag("FolderId", 0x0C)
	result.AddTag("ItemId", 0x0D)
	result.AddTag("LongId", 0x0E)
	result.AddTag("InstanceId", 0x0F)
	result.AddTag("Mime", 0x10)
	result.AddTag("ClientId", 0x11)
	result.AddTag("Status", 0x12)
	if protocolVersion > PROTOCOL_VERSION_14_0 {
		result.AddTag("AccountId", 0x13) // not supported with MS-ASProtocolVersion 12.1 or 14.0
	}

	return result
}

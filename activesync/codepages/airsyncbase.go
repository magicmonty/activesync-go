package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func AirSyncBase(protocolVersion byte) CodePage {
	result := NewCodePage("AirSyncBase", 17)

	result.AddTag("BodyPreference", 0x05)
	result.AddTag("Type", 0x06)
	result.AddTag("TruncationSize", 0x07)
	result.AddTag("AllOrNone", 0x08)
	result.AddTag("Body", 0x0A)
	result.AddTag("Data", 0x0B)
	result.AddTag("EstimatedDataSize", 0x0C)
	result.AddTag("Truncated", 0x0D)
	result.AddTag("Attachments", 0x0E)
	result.AddTag("Attachment", 0x0F)
	result.AddTag("DisplayName", 0x10)
	result.AddTag("FileReference", 0x11)
	result.AddTag("Method", 0x12)
	result.AddTag("ContentId", 0x13)
	result.AddTag("ContentLocation", 0x14) // not used
	result.AddTag("IsInline", 0x15)
	result.AddTag("NativeBodyType", 0x16)
	result.AddTag("ContentType", 0x17)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag("Preview", 0x18) // not supported with MS-ASProtocolVersion 12.1
		if protocolVersion > PROTOCOL_VERSION_14_0 {
			result.AddTag("BodyPartPreference", 0x19) // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("BodyPart", 0x1A)           // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("Status", 0x1B)             // not supported with MS-ASProtocolVersion 12.1 or 14.0
		}
	}

	return result
}

package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

// not supported with MS-ASProtocolVersion 12.1
func Email2(protocolVersion byte) CodePage {
	result := NewCodePage("Email2", 22)

	result.AddTag("UmCallerID", 0x05)
	result.AddTag("UmUserNotes", 0x06)
	result.AddTag("UmAttDuration", 0x07)
	result.AddTag("UmAttOrder", 0x08)
	result.AddTag("ConversationId", 0x09)
	result.AddTag("ConversationIndex", 0x0A)
	result.AddTag("LastVerbExecuted", 0x0B)
	result.AddTag("LastVerbExecutionTime", 0x0C)
	result.AddTag("ReceivedAsBcc", 0x0D)
	result.AddTag("Sender", 0x0E)
	result.AddTag("CalendarType", 0x0F)
	result.AddTag("IsLeapMonth", 0x10)
	if protocolVersion > PROTOCOL_VERSION_14_0 {
		result.AddTag("AccountId", 0x11)          // not supported with MS-ASProtocolVersion 12.1 or 14.0
		result.AddTag("FirstDayOfWeek", 0x12)     // not supported with MS-ASProtocolVersion 12.1 or 14.0
		result.AddTag("MeetingMessageType", 0x13) // not supported with MS-ASProtocolVersion 12.1 or 14.0
	}

	return result
}

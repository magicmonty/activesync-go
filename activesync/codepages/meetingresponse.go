package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func MeetingResponse() CodePage {
	result := NewCodePage("MeetingResponse", 8)

	result.AddTag("CalendarId", 0x05)
	result.AddTag("CollectionId", 0x06)
	result.AddTag("MeetingResponse", 0x07)
	result.AddTag("RequestId", 0x08)
	result.AddTag("Request", 0x09)
	result.AddTag("Result", 0x0A)
	result.AddTag("Status", 0x0B)
	result.AddTag("UserResponse", 0x0C)
	result.AddTag("InstanceId", 0x0E)

	return result
}

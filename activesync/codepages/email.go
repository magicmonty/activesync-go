package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func Email(protocolVersion byte) CodePage {
	result := NewCodePage("Email", 2)

	result.AddTag("DateReceived", 0x0F)
	result.AddTag("DisplayTo", 0x11)
	result.AddTag("Importance", 0x12)
	result.AddTag("MessageClass", 0x13)
	result.AddTag("Subject", 0x14)
	result.AddTag("Read", 0x15)
	result.AddTag("To", 0x16)
	result.AddTag("Cc", 0x17)
	result.AddTag("From", 0x18)
	result.AddTag("ReplyTo", 0x19)
	result.AddTag("AllDayEvent", 0x1A)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag("Categories", 0x1B) // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("Category", 0x1C)   // not supported with MS-ASProtocolVersion 12.1
	}
	result.AddTag("DTStamp", 0x1D)
	result.AddTag("EndTime", 0x1E)
	result.AddTag("InstanceType", 0x1F)
	result.AddTag("BusyStatus", 0x20)
	result.AddTag("Location", 0x21)
	result.AddTag("MeetingRequest", 0x22)
	result.AddTag("Organizer", 0x23)
	result.AddTag("RecurrenceId", 0x24)
	result.AddTag("Reminder", 0x25)
	result.AddTag("ResponseRequested", 0x26)
	result.AddTag("Recurrences", 0x27)
	result.AddTag("Recurrence", 0x28)
	result.AddTag("Recurrence_Type", 0x29)
	result.AddTag("Recurrence_Until", 0x2A)
	result.AddTag("Recurrence_Occurrences", 0x2B)
	result.AddTag("Recurrence_Interval", 0x2C)
	result.AddTag("Recurrence_DayOfWeek", 0x2D)
	result.AddTag("Recurrence_DayOfMonth", 0x2E)
	result.AddTag("Recurrence_WeekOfMonth", 0x2F)
	result.AddTag("Recurrence_MonthOfYear", 0x30)
	result.AddTag("StartTime", 0x31)
	result.AddTag("Sensitivity", 0x32)
	result.AddTag("TimeZone", 0x33)
	result.AddTag("GlobalObjId", 0x34)
	result.AddTag("ThreadTopic", 0x35)
	result.AddTag("InternetCPID", 0x39)
	result.AddTag("Flag", 0x3A)
	result.AddTag("Status", 0x3B)
	result.AddTag("ContentClass", 0x3C)
	result.AddTag("FlagType", 0x3D)
	result.AddTag("CompleteTime", 0x3E)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag("DisallowNewTimeProposal", 0x3F) // not supported with MS-ASProtocolVersion 12.1
	}

	return result
}

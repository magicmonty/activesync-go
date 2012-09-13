package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func Calendar(protocolVersion byte) CodePage {
	result := NewCodePage("Calendar", 4)

	result.AddTag("TimeZone", 0x05)
	result.AddTag("AllDayEvent", 0x06)
	result.AddTag("Attendees", 0x07)
	result.AddTag("Attendee", 0x08)
	result.AddTag("Email", 0x09)
	result.AddTag("Name", 0x0A)
	result.AddTag("BusyStatus", 0x0D)
	result.AddTag("Categories", 0x0E)
	result.AddTag("Category", 0x0F)
	result.AddTag("DtStamp", 0x11)
	result.AddTag("EndTime", 0x12)
	result.AddTag("Exception", 0x13)
	result.AddTag("Exceptions", 0x14)
	result.AddTag("Deleted", 0x15)
	result.AddTag("ExceptionStartTime", 0x16)
	result.AddTag("Location", 0x17)
	result.AddTag("MeetingStatus", 0x18)
	result.AddTag("OrganizerEmail", 0x19)
	result.AddTag("OrganizerName", 0x1A)
	result.AddTag("Recurrence", 0x1B)
	result.AddTag("Type", 0x1C)
	result.AddTag("Until", 0x1D)
	result.AddTag("Occurrences", 0x1E)
	result.AddTag("Interval", 0x1F)
	result.AddTag("DayOfWeek", 0x20)
	result.AddTag("DayOfMonth", 0x21)
	result.AddTag("WeekOfMonth", 0x22)
	result.AddTag("MonthOfYear", 0x23)
	result.AddTag("Reminder", 0x24)
	result.AddTag("Sensitivity", 0x25)
	result.AddTag("Subject", 0x26)
	result.AddTag("StartTime", 0x27)
	result.AddTag("UID", 0x28)
	result.AddTag("AttendeeStatus", 0x29)
	result.AddTag("AttendeeType", 0x2A)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag("DisallowNewTimeProposal", 0x33) // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("ResponseRequested", 0x34)       // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("AppointmentReplyTime", 0x35)    // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("ResponseType", 0x36)            // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("CalendarType", 0x37)            // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("IsLeapMonth", 0x38)             // not supported with MS-ASProtocolVersion 12.1
		if protocolVersion > PROTOCOL_VERSION_14_0 {
			result.AddTag("FirstDayOfWeek", 0x39)            // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("OnlineMeetingConfLink", 0x3A)     // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("OnlineMeetingExternalLink", 0x3B) // not supported with MS-ASProtocolVersion 12.1 or 14.0
		}
	}
	return result
}

package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func Tasks(protocolVersion byte) CodePage {
	result := NewCodePage("Tasks", 9)

	result.AddTag("Categories", 0x08)
	result.AddTag("Category", 0x09)
	result.AddTag("Complete", 0x0A)
	result.AddTag("DateCompleted", 0x0B)
	result.AddTag("DueDate", 0x0C)
	result.AddTag("UtcDueDate", 0x0D)
	result.AddTag("Importance", 0x0E)
	result.AddTag("Recurrence", 0x0F)
	result.AddTag("Recurrence_Type", 0x10)
	result.AddTag("Recurrence_Start", 0x11)
	result.AddTag("Recurrence_Until", 0x12)
	result.AddTag("Recurrence_Occurrences", 0x13)
	result.AddTag("Recurrence_Interval", 0x14)
	result.AddTag("Recurrence_DayOfMonth", 0x15)
	result.AddTag("Recurrence_DayOfWeek", 0x16)
	result.AddTag("Recurrence_WeekOfMonth", 0x17)
	result.AddTag("Recurrence_MonthOfYear", 0x18)
	result.AddTag("Recurrence_Regenerate", 0x19)
	result.AddTag("Recurrence_DeadOccur", 0x1A)
	result.AddTag("ReminderSet", 0x1B)
	result.AddTag("ReminderTime", 0x1C)
	result.AddTag("Sensitivity", 0x1D)
	result.AddTag("StartDate", 0x1E)
	result.AddTag("UtcStartDate", 0x1F)
	result.AddTag("Subject", 0x20)
	result.AddTag("OrdinalDate", 0x22)
	result.AddTag("SubOrdinalDate", 0x23)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag("CalendarType", 0x24) // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("IsLeapMonth", 0x25)  // not supported with MS-ASProtocolVersion 12.1
		if protocolVersion > PROTOCOL_VERSION_14_0 {
			result.AddTag("FirstDayOfWeek", 0x26) // not supported with MS-ASProtocolVersion 12.1 or 14.0
		}
	}

	return result
}

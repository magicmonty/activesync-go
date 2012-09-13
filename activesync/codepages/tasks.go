package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_TASKS                         string = "Tasks"
	TAG_TASKS_CATEGORIES             string = "Categories"
	TAG_TASKS_CATEGORY               string = "Category"
	TAG_TASKS_COMPLETE               string = "Complete"
	TAG_TASKS_DATECOMPLETED          string = "DateCompleted"
	TAG_TASKS_DUEDATE                string = "DueDate"
	TAG_TASKS_UTCDUEDATE             string = "UtcDueDate"
	TAG_TASKS_IMPORTANCE             string = "Importance"
	TAG_TASKS_RECURRENCE             string = "Recurrence"
	TAG_TASKS_RECURRENCE_TYPE        string = "Recurrence_Type"
	TAG_TASKS_RECURRENCE_START       string = "Recurrence_Start"
	TAG_TASKS_RECURRENCE_UNTIL       string = "Recurrence_Until"
	TAG_TASKS_RECURRENCE_OCCURRENCES string = "Recurrence_Occurrences"
	TAG_TASKS_RECURRENCE_INTERVAL    string = "Recurrence_Interval"
	TAG_TASKS_RECURRENCE_DAYOFMONTH  string = "Recurrence_DayOfMonth"
	TAG_TASKS_RECURRENCE_DAYOFWEEK   string = "Recurrence_DayOfWeek"
	TAG_TASKS_RECURRENCE_WEEKOFMONTH string = "Recurrence_WeekOfMonth"
	TAG_TASKS_RECURRENCE_MONTHOFYEAR string = "Recurrence_MonthOfYear"
	TAG_TASKS_RECURRENCE_REGENERATE  string = "Recurrence_Regenerate"
	TAG_TASKS_RECURRENCE_DEADOCCUR   string = "Recurrence_DeadOccur"
	TAG_TASKS_REMINDERSET            string = "ReminderSet"
	TAG_TASKS_REMINDERTIME           string = "ReminderTime"
	TAG_TASKS_SENSITIVITY            string = "Sensitivity"
	TAG_TASKS_STARTDATE              string = "StartDate"
	TAG_TASKS_UTCSTARTDATE           string = "UtcStartDate"
	TAG_TASKS_SUBJECT                string = "Subject"
	TAG_TASKS_ORDINALDATE            string = "OrdinalDate"
	TAG_TASKS_SUBORDINALDATE         string = "SubOrdinalDate"
	TAG_TASKS_CALENDARTYPE           string = "CalendarType"
	TAG_TASKS_ISLEAPMONTH            string = "IsLeapMonth"
	TAG_TASKS_FIRSTDAYOFWEEK         string = "FirstDayOfWeek"
)

const (
	CP_TASKS                        byte = 9
	ID_TASKS_CATEGORIES             byte = 0x08
	ID_TASKS_CATEGORY               byte = 0x09
	ID_TASKS_COMPLETE               byte = 0x0A
	ID_TASKS_DATECOMPLETED          byte = 0x0B
	ID_TASKS_DUEDATE                byte = 0x0C
	ID_TASKS_UTCDUEDATE             byte = 0x0D
	ID_TASKS_IMPORTANCE             byte = 0x0E
	ID_TASKS_RECURRENCE             byte = 0x0F
	ID_TASKS_RECURRENCE_TYPE        byte = 0x10
	ID_TASKS_RECURRENCE_START       byte = 0x11
	ID_TASKS_RECURRENCE_UNTIL       byte = 0x12
	ID_TASKS_RECURRENCE_OCCURRENCES byte = 0x13
	ID_TASKS_RECURRENCE_INTERVAL    byte = 0x14
	ID_TASKS_RECURRENCE_DAYOFMONTH  byte = 0x15
	ID_TASKS_RECURRENCE_DAYOFWEEK   byte = 0x16
	ID_TASKS_RECURRENCE_WEEKOFMONTH byte = 0x17
	ID_TASKS_RECURRENCE_MONTHOFYEAR byte = 0x18
	ID_TASKS_RECURRENCE_REGENERATE  byte = 0x19
	ID_TASKS_RECURRENCE_DEADOCCUR   byte = 0x1A
	ID_TASKS_REMINDERSET            byte = 0x1B
	ID_TASKS_REMINDERTIME           byte = 0x1C
	ID_TASKS_SENSITIVITY            byte = 0x1D
	ID_TASKS_STARTDATE              byte = 0x1E
	ID_TASKS_UTCSTARTDATE           byte = 0x1F
	ID_TASKS_SUBJECT                byte = 0x20
	ID_TASKS_ORDINALDATE            byte = 0x22
	ID_TASKS_SUBORDINALDATE         byte = 0x23
	ID_TASKS_CALENDARTYPE           byte = 0x24
	ID_TASKS_ISLEAPMONTH            byte = 0x25
	ID_TASKS_FIRSTDAYOFWEEK         byte = 0x26
)

func Tasks(protocolVersion byte) CodePage {
	result := NewCodePage(NS_TASKS, CP_TASKS)

	result.AddTag(TAG_TASKS_CATEGORIES, ID_TASKS_CATEGORIES)
	result.AddTag(TAG_TASKS_CATEGORY, ID_TASKS_CATEGORY)
	result.AddTag(TAG_TASKS_COMPLETE, ID_TASKS_COMPLETE)
	result.AddTag(TAG_TASKS_DATECOMPLETED, ID_TASKS_DATECOMPLETED)
	result.AddTag(TAG_TASKS_DUEDATE, ID_TASKS_DUEDATE)
	result.AddTag(TAG_TASKS_UTCDUEDATE, ID_TASKS_UTCDUEDATE)
	result.AddTag(TAG_TASKS_IMPORTANCE, ID_TASKS_IMPORTANCE)
	result.AddTag(TAG_TASKS_RECURRENCE, ID_TASKS_RECURRENCE)
	result.AddTag(TAG_TASKS_RECURRENCE_TYPE, ID_TASKS_RECURRENCE_TYPE)
	result.AddTag(TAG_TASKS_RECURRENCE_START, ID_TASKS_RECURRENCE_START)
	result.AddTag(TAG_TASKS_RECURRENCE_UNTIL, ID_TASKS_RECURRENCE_UNTIL)
	result.AddTag(TAG_TASKS_RECURRENCE_OCCURRENCES, ID_TASKS_RECURRENCE_OCCURRENCES)
	result.AddTag(TAG_TASKS_RECURRENCE_INTERVAL, ID_TASKS_RECURRENCE_INTERVAL)
	result.AddTag(TAG_TASKS_RECURRENCE_DAYOFMONTH, ID_TASKS_RECURRENCE_DAYOFMONTH)
	result.AddTag(TAG_TASKS_RECURRENCE_DAYOFWEEK, ID_TASKS_RECURRENCE_DAYOFWEEK)
	result.AddTag(TAG_TASKS_RECURRENCE_WEEKOFMONTH, ID_TASKS_RECURRENCE_WEEKOFMONTH)
	result.AddTag(TAG_TASKS_RECURRENCE_MONTHOFYEAR, ID_TASKS_RECURRENCE_MONTHOFYEAR)
	result.AddTag(TAG_TASKS_RECURRENCE_REGENERATE, ID_TASKS_RECURRENCE_REGENERATE)
	result.AddTag(TAG_TASKS_RECURRENCE_DEADOCCUR, ID_TASKS_RECURRENCE_DEADOCCUR)
	result.AddTag(TAG_TASKS_REMINDERSET, ID_TASKS_REMINDERSET)
	result.AddTag(TAG_TASKS_REMINDERTIME, ID_TASKS_REMINDERTIME)
	result.AddTag(TAG_TASKS_SENSITIVITY, ID_TASKS_SENSITIVITY)
	result.AddTag(TAG_TASKS_STARTDATE, ID_TASKS_STARTDATE)
	result.AddTag(TAG_TASKS_UTCSTARTDATE, ID_TASKS_UTCSTARTDATE)
	result.AddTag(TAG_TASKS_SUBJECT, ID_TASKS_SUBJECT)
	result.AddTag(TAG_TASKS_ORDINALDATE, ID_TASKS_ORDINALDATE)
	result.AddTag(TAG_TASKS_SUBORDINALDATE, ID_TASKS_SUBORDINALDATE)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_TASKS_CALENDARTYPE, ID_TASKS_CALENDARTYPE) // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_TASKS_ISLEAPMONTH, ID_TASKS_ISLEAPMONTH)   // not supported with MS-ASProtocolVersion 12.1
		if protocolVersion > PROTOCOL_VERSION_14_0 {
			result.AddTag(TAG_TASKS_FIRSTDAYOFWEEK, ID_TASKS_FIRSTDAYOFWEEK) // not supported with MS-ASProtocolVersion 12.1 or 14.0
		}
	}

	return result
}

package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_CALENDAR                            string = "Calendar"
	TAG_CALENDAR_TIMEZONE                  string = "TimeZone"
	TAG_CALENDAR_ALLDAYEVENT               string = "AllDayEvent"
	TAG_CALENDAR_ATTENDEES                 string = "Attendees"
	TAG_CALENDAR_ATTENDEE                  string = "Attendee"
	TAG_CALENDAR_EMAIL                     string = "Email"
	TAG_CALENDAR_NAME                      string = "Name"
	TAG_CALENDAR_BUSYSTATUS                string = "BusyStatus"
	TAG_CALENDAR_CATEGORIES                string = "Categories"
	TAG_CALENDAR_CATEGORY                  string = "Category"
	TAG_CALENDAR_DTSTAMP                   string = "DtStamp"
	TAG_CALENDAR_ENDTIME                   string = "EndTime"
	TAG_CALENDAR_EXCEPTION                 string = "Exception"
	TAG_CALENDAR_EXCEPTIONS                string = "Exceptions"
	TAG_CALENDAR_DELETED                   string = "Deleted"
	TAG_CALENDAR_EXCEPTIONSTARTTIME        string = "ExceptionStartTime"
	TAG_CALENDAR_LOCATION                  string = "Location"
	TAG_CALENDAR_MEETINGSTATUS             string = "MeetingStatus"
	TAG_CALENDAR_ORGANIZEREMAIL            string = "OrganizerEmail"
	TAG_CALENDAR_ORGANIZERNAME             string = "OrganizerName"
	TAG_CALENDAR_RECURRENCE                string = "Recurrence"
	TAG_CALENDAR_TYPE                      string = "Type"
	TAG_CALENDAR_UNTIL                     string = "Until"
	TAG_CALENDAR_OCCURRENCES               string = "Occurrences"
	TAG_CALENDAR_INTERVAL                  string = "Interval"
	TAG_CALENDAR_DAYOFWEEK                 string = "DayOfWeek"
	TAG_CALENDAR_DAYOFMONTH                string = "DayOfMonth"
	TAG_CALENDAR_WEEKOFMONTH               string = "WeekOfMonth"
	TAG_CALENDAR_MONTHOFYEAR               string = "MonthOfYear"
	TAG_CALENDAR_REMINDER                  string = "Reminder"
	TAG_CALENDAR_SENSITIVITY               string = "Sensitivity"
	TAG_CALENDAR_SUBJECT                   string = "Subject"
	TAG_CALENDAR_STARTTIME                 string = "StartTime"
	TAG_CALENDAR_UID                       string = "UID"
	TAG_CALENDAR_ATTENDEESTATUS            string = "AttendeeStatus"
	TAG_CALENDAR_ATTENDEETYPE              string = "AttendeeType"
	TAG_CALENDAR_DISALLOWNEWTIMEPROPOSAL   string = "DisallowNewTimeProposal"
	TAG_CALENDAR_RESPONSEREQUESTED         string = "ResponseRequested"
	TAG_CALENDAR_APPOINTMENTREPLYTIME      string = "AppointmentReplyTime"
	TAG_CALENDAR_RESPONSETYPE              string = "ResponseType"
	TAG_CALENDAR_CALENDARTYPE              string = "CalendarType"
	TAG_CALENDAR_ISLEAPMONTH               string = "IsLeapMonth"
	TAG_CALENDAR_FIRSTDAYOFWEEK            string = "FirstDayOfWeek"
	TAG_CALENDAR_ONLINEMEETINGCONFLINK     string = "OnlineMeetingConfLink"
	TAG_CALENDAR_ONLINEMEETINGEXTERNALLINK string = "OnlineMeetingExternalLink"
)

const (
	CP_CALENDAR                           byte = 4
	ID_CALENDAR_TIMEZONE                  byte = 0x05
	ID_CALENDAR_ALLDAYEVENT               byte = 0x06
	ID_CALENDAR_ATTENDEES                 byte = 0x07
	ID_CALENDAR_ATTENDEE                  byte = 0x08
	ID_CALENDAR_EMAIL                     byte = 0x09
	ID_CALENDAR_NAME                      byte = 0x0A
	ID_CALENDAR_BUSYSTATUS                byte = 0x0D
	ID_CALENDAR_CATEGORIES                byte = 0x0E
	ID_CALENDAR_CATEGORY                  byte = 0x0F
	ID_CALENDAR_DTSTAMP                   byte = 0x11
	ID_CALENDAR_ENDTIME                   byte = 0x12
	ID_CALENDAR_EXCEPTION                 byte = 0x13
	ID_CALENDAR_EXCEPTIONS                byte = 0x14
	ID_CALENDAR_DELETED                   byte = 0x15
	ID_CALENDAR_EXCEPTIONSTARTTIME        byte = 0x16
	ID_CALENDAR_LOCATION                  byte = 0x17
	ID_CALENDAR_MEETINGSTATUS             byte = 0x18
	ID_CALENDAR_ORGANIZEREMAIL            byte = 0x19
	ID_CALENDAR_ORGANIZERNAME             byte = 0x1A
	ID_CALENDAR_RECURRENCE                byte = 0x1B
	ID_CALENDAR_TYPE                      byte = 0x1C
	ID_CALENDAR_UNTIL                     byte = 0x1D
	ID_CALENDAR_OCCURRENCES               byte = 0x1E
	ID_CALENDAR_INTERVAL                  byte = 0x1F
	ID_CALENDAR_DAYOFWEEK                 byte = 0x20
	ID_CALENDAR_DAYOFMONTH                byte = 0x21
	ID_CALENDAR_WEEKOFMONTH               byte = 0x22
	ID_CALENDAR_MONTHOFYEAR               byte = 0x23
	ID_CALENDAR_REMINDER                  byte = 0x24
	ID_CALENDAR_SENSITIVITY               byte = 0x25
	ID_CALENDAR_SUBJECT                   byte = 0x26
	ID_CALENDAR_STARTTIME                 byte = 0x27
	ID_CALENDAR_UID                       byte = 0x28
	ID_CALENDAR_ATTENDEESTATUS            byte = 0x29
	ID_CALENDAR_ATTENDEETYPE              byte = 0x2A
	ID_CALENDAR_DISALLOWNEWTIMEPROPOSAL   byte = 0x33
	ID_CALENDAR_RESPONSEREQUESTED         byte = 0x34
	ID_CALENDAR_APPOINTMENTREPLYTIME      byte = 0x35
	ID_CALENDAR_RESPONSETYPE              byte = 0x36
	ID_CALENDAR_CALENDARTYPE              byte = 0x37
	ID_CALENDAR_ISLEAPMONTH               byte = 0x38
	ID_CALENDAR_FIRSTDAYOFWEEK            byte = 0x39
	ID_CALENDAR_ONLINEMEETINGCONFLINK     byte = 0x3A
	ID_CALENDAR_ONLINEMEETINGEXTERNALLINK byte = 0x3B
)

func Calendar(protocolVersion byte) CodePage {
	result := NewCodePage(NS_CALENDAR, CP_CALENDAR)

	result.AddTag(TAG_CALENDAR_TIMEZONE, ID_CALENDAR_TIMEZONE)
	result.AddTag(TAG_CALENDAR_ALLDAYEVENT, ID_CALENDAR_ALLDAYEVENT)
	result.AddTag(TAG_CALENDAR_ATTENDEES, ID_CALENDAR_ATTENDEES)
	result.AddTag(TAG_CALENDAR_ATTENDEE, ID_CALENDAR_ATTENDEE)
	result.AddTag(TAG_CALENDAR_EMAIL, ID_CALENDAR_EMAIL)
	result.AddTag(TAG_CALENDAR_NAME, ID_CALENDAR_NAME)
	result.AddTag(TAG_CALENDAR_BUSYSTATUS, ID_CALENDAR_BUSYSTATUS)
	result.AddTag(TAG_CALENDAR_CATEGORIES, ID_CALENDAR_CATEGORIES)
	result.AddTag(TAG_CALENDAR_CATEGORY, ID_CALENDAR_CATEGORY)
	result.AddTag(TAG_CALENDAR_DTSTAMP, ID_CALENDAR_DTSTAMP)
	result.AddTag(TAG_CALENDAR_ENDTIME, ID_CALENDAR_ENDTIME)
	result.AddTag(TAG_CALENDAR_EXCEPTION, ID_CALENDAR_EXCEPTION)
	result.AddTag(TAG_CALENDAR_EXCEPTIONS, ID_CALENDAR_EXCEPTIONS)
	result.AddTag(TAG_CALENDAR_DELETED, ID_CALENDAR_DELETED)
	result.AddTag(TAG_CALENDAR_EXCEPTIONSTARTTIME, ID_CALENDAR_EXCEPTIONSTARTTIME)
	result.AddTag(TAG_CALENDAR_LOCATION, ID_CALENDAR_LOCATION)
	result.AddTag(TAG_CALENDAR_MEETINGSTATUS, ID_CALENDAR_MEETINGSTATUS)
	result.AddTag(TAG_CALENDAR_ORGANIZEREMAIL, ID_CALENDAR_ORGANIZEREMAIL)
	result.AddTag(TAG_CALENDAR_ORGANIZERNAME, ID_CALENDAR_ORGANIZERNAME)
	result.AddTag(TAG_CALENDAR_RECURRENCE, ID_CALENDAR_RECURRENCE)
	result.AddTag(TAG_CALENDAR_TYPE, ID_CALENDAR_TYPE)
	result.AddTag(TAG_CALENDAR_UNTIL, ID_CALENDAR_UNTIL)
	result.AddTag(TAG_CALENDAR_OCCURRENCES, ID_CALENDAR_OCCURRENCES)
	result.AddTag(TAG_CALENDAR_INTERVAL, ID_CALENDAR_INTERVAL)
	result.AddTag(TAG_CALENDAR_DAYOFWEEK, ID_CALENDAR_DAYOFWEEK)
	result.AddTag(TAG_CALENDAR_DAYOFMONTH, ID_CALENDAR_DAYOFMONTH)
	result.AddTag(TAG_CALENDAR_WEEKOFMONTH, ID_CALENDAR_WEEKOFMONTH)
	result.AddTag(TAG_CALENDAR_MONTHOFYEAR, ID_CALENDAR_MONTHOFYEAR)
	result.AddTag(TAG_CALENDAR_REMINDER, ID_CALENDAR_REMINDER)
	result.AddTag(TAG_CALENDAR_SENSITIVITY, ID_CALENDAR_SENSITIVITY)
	result.AddTag(TAG_CALENDAR_SUBJECT, ID_CALENDAR_SUBJECT)
	result.AddTag(TAG_CALENDAR_STARTTIME, ID_CALENDAR_STARTTIME)
	result.AddTag(TAG_CALENDAR_UID, ID_CALENDAR_UID)
	result.AddTag(TAG_CALENDAR_ATTENDEESTATUS, ID_CALENDAR_ATTENDEESTATUS)
	result.AddTag(TAG_CALENDAR_ATTENDEETYPE, ID_CALENDAR_ATTENDEETYPE)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_CALENDAR_DISALLOWNEWTIMEPROPOSAL, ID_CALENDAR_DISALLOWNEWTIMEPROPOSAL) // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_CALENDAR_RESPONSEREQUESTED, ID_CALENDAR_RESPONSEREQUESTED)             // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_CALENDAR_APPOINTMENTREPLYTIME, ID_CALENDAR_APPOINTMENTREPLYTIME)       // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_CALENDAR_RESPONSETYPE, ID_CALENDAR_RESPONSETYPE)                       // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_CALENDAR_CALENDARTYPE, ID_CALENDAR_CALENDARTYPE)                       // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_CALENDAR_ISLEAPMONTH, ID_CALENDAR_ISLEAPMONTH)                         // not supported with MS-ASProtocolVersion 12.1
		if protocolVersion > PROTOCOL_VERSION_14_0 {
			result.AddTag(TAG_CALENDAR_FIRSTDAYOFWEEK, ID_CALENDAR_FIRSTDAYOFWEEK)                       // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_CALENDAR_ONLINEMEETINGCONFLINK, ID_CALENDAR_ONLINEMEETINGCONFLINK)         // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_CALENDAR_ONLINEMEETINGEXTERNALLINK, ID_CALENDAR_ONLINEMEETINGEXTERNALLINK) // not supported with MS-ASProtocolVersion 12.1 or 14.0
		}
	}
	return result
}

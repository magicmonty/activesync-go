package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_EMAIL                          string = "Email"
	TAG_EMAIL_DATERECEIVED            string = "DateReceived"
	TAG_EMAIL_DISPLAYTO               string = "DisplayTo"
	TAG_EMAIL_IMPORTANCE              string = "Importance"
	TAG_EMAIL_MESSAGECLASS            string = "MessageClass"
	TAG_EMAIL_SUBJECT                 string = "Subject"
	TAG_EMAIL_READ                    string = "Read"
	TAG_EMAIL_TO                      string = "To"
	TAG_EMAIL_CC                      string = "Cc"
	TAG_EMAIL_FROM                    string = "From"
	TAG_EMAIL_REPLYTO                 string = "ReplyTo"
	TAG_EMAIL_ALLDAYEVENT             string = "AllDayEvent"
	TAG_EMAIL_CATEGORIES              string = "Categories"
	TAG_EMAIL_CATEGORY                string = "Category"
	TAG_EMAIL_DTSTAMP                 string = "DTStamp"
	TAG_EMAIL_ENDTIME                 string = "EndTime"
	TAG_EMAIL_INSTANCETYPE            string = "InstanceType"
	TAG_EMAIL_BUSYSTATUS              string = "BusyStatus"
	TAG_EMAIL_LOCATION                string = "Location"
	TAG_EMAIL_MEETINGREQUEST          string = "MeetingRequest"
	TAG_EMAIL_ORGANIZER               string = "Organizer"
	TAG_EMAIL_RECURRENCEID            string = "RecurrenceId"
	TAG_EMAIL_REMINDER                string = "Reminder"
	TAG_EMAIL_RESPONSEREQUESTED       string = "ResponseRequested"
	TAG_EMAIL_RECURRENCES             string = "Recurrences"
	TAG_EMAIL_RECURRENCE              string = "Recurrence"
	TAG_EMAIL_RECURRENCE_TYPE         string = "Recurrence_Type"
	TAG_EMAIL_RECURRENCE_UNTIL        string = "Recurrence_Until"
	TAG_EMAIL_RECURRENCE_OCCURRENCES  string = "Recurrence_Occurrences"
	TAG_EMAIL_RECURRENCE_INTERVAL     string = "Recurrence_Interval"
	TAG_EMAIL_RECURRENCE_DAYOFWEEK    string = "Recurrence_DayOfWeek"
	TAG_EMAIL_RECURRENCE_DAYOFMONTH   string = "Recurrence_DayOfMonth"
	TAG_EMAIL_RECURRENCE_WEEKOFMONTH  string = "Recurrence_WeekOfMonth"
	TAG_EMAIL_RECURRENCE_MONTHOFYEAR  string = "Recurrence_MonthOfYear"
	TAG_EMAIL_STARTTIME               string = "StartTime"
	TAG_EMAIL_SENSITIVITY             string = "Sensitivity"
	TAG_EMAIL_TIMEZONE                string = "TimeZone"
	TAG_EMAIL_GLOBALOBJID             string = "GlobalObjId"
	TAG_EMAIL_THREADTOPIC             string = "ThreadTopic"
	TAG_EMAIL_INTERNETCPID            string = "InternetCPID"
	TAG_EMAIL_FLAG                    string = "Flag"
	TAG_EMAIL_STATUS                  string = "Status"
	TAG_EMAIL_CONTENTCLASS            string = "ContentClass"
	TAG_EMAIL_FLAGTYPE                string = "FlagType"
	TAG_EMAIL_COMPLETETIME            string = "CompleteTime"
	TAG_EMAIL_DISALLOWNEWTIMEPROPOSAL string = "DisallowNewTimeProposal"
)

const (
	CP_EMAIL                         byte = 2
	ID_EMAIL_DATERECEIVED            byte = 0x0F
	ID_EMAIL_DISPLAYTO               byte = 0x11
	ID_EMAIL_IMPORTANCE              byte = 0x12
	ID_EMAIL_MESSAGECLASS            byte = 0x13
	ID_EMAIL_SUBJECT                 byte = 0x14
	ID_EMAIL_READ                    byte = 0x15
	ID_EMAIL_TO                      byte = 0x16
	ID_EMAIL_CC                      byte = 0x17
	ID_EMAIL_FROM                    byte = 0x18
	ID_EMAIL_REPLYTO                 byte = 0x19
	ID_EMAIL_ALLDAYEVENT             byte = 0x1A
	ID_EMAIL_CATEGORIES              byte = 0x1B
	ID_EMAIL_CATEGORY                byte = 0x1C
	ID_EMAIL_DTSTAMP                 byte = 0x1D
	ID_EMAIL_ENDTIME                 byte = 0x1E
	ID_EMAIL_INSTANCETYPE            byte = 0x1F
	ID_EMAIL_BUSYSTATUS              byte = 0x20
	ID_EMAIL_LOCATION                byte = 0x21
	ID_EMAIL_MEETINGREQUEST          byte = 0x22
	ID_EMAIL_ORGANIZER               byte = 0x23
	ID_EMAIL_RECURRENCEID            byte = 0x24
	ID_EMAIL_REMINDER                byte = 0x25
	ID_EMAIL_RESPONSEREQUESTED       byte = 0x26
	ID_EMAIL_RECURRENCES             byte = 0x27
	ID_EMAIL_RECURRENCE              byte = 0x28
	ID_EMAIL_RECURRENCE_TYPE         byte = 0x29
	ID_EMAIL_RECURRENCE_UNTIL        byte = 0x2A
	ID_EMAIL_RECURRENCE_OCCURRENCES  byte = 0x2B
	ID_EMAIL_RECURRENCE_INTERVAL     byte = 0x2C
	ID_EMAIL_RECURRENCE_DAYOFWEEK    byte = 0x2D
	ID_EMAIL_RECURRENCE_DAYOFMONTH   byte = 0x2E
	ID_EMAIL_RECURRENCE_WEEKOFMONTH  byte = 0x2F
	ID_EMAIL_RECURRENCE_MONTHOFYEAR  byte = 0x30
	ID_EMAIL_STARTTIME               byte = 0x31
	ID_EMAIL_SENSITIVITY             byte = 0x32
	ID_EMAIL_TIMEZONE                byte = 0x33
	ID_EMAIL_GLOBALOBJID             byte = 0x34
	ID_EMAIL_THREADTOPIC             byte = 0x35
	ID_EMAIL_INTERNETCPID            byte = 0x39
	ID_EMAIL_FLAG                    byte = 0x3A
	ID_EMAIL_STATUS                  byte = 0x3B
	ID_EMAIL_CONTENTCLASS            byte = 0x3C
	ID_EMAIL_FLAGTYPE                byte = 0x3D
	ID_EMAIL_COMPLETETIME            byte = 0x3E
	ID_EMAIL_DISALLOWNEWTIMEPROPOSAL byte = 0x3F
)

func Email(protocolVersion byte) CodePage {
	result := NewCodePage(NS_EMAIL, CP_EMAIL)

	result.AddTag(TAG_EMAIL_DATERECEIVED, ID_EMAIL_DATERECEIVED)
	result.AddTag(TAG_EMAIL_DISPLAYTO, ID_EMAIL_DISPLAYTO)
	result.AddTag(TAG_EMAIL_IMPORTANCE, ID_EMAIL_IMPORTANCE)
	result.AddTag(TAG_EMAIL_MESSAGECLASS, ID_EMAIL_MESSAGECLASS)
	result.AddTag(TAG_EMAIL_SUBJECT, ID_EMAIL_SUBJECT)
	result.AddTag(TAG_EMAIL_READ, ID_EMAIL_READ)
	result.AddTag(TAG_EMAIL_TO, ID_EMAIL_TO)
	result.AddTag(TAG_EMAIL_CC, ID_EMAIL_CC)
	result.AddTag(TAG_EMAIL_FROM, ID_EMAIL_FROM)
	result.AddTag(TAG_EMAIL_REPLYTO, ID_EMAIL_REPLYTO)
	result.AddTag(TAG_EMAIL_ALLDAYEVENT, ID_EMAIL_ALLDAYEVENT)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_EMAIL_CATEGORIES, ID_EMAIL_CATEGORIES) // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_EMAIL_CATEGORY, ID_EMAIL_CATEGORY)     // not supported with MS-ASProtocolVersion 12.1
	}
	result.AddTag(TAG_EMAIL_DTSTAMP, ID_EMAIL_DTSTAMP)
	result.AddTag(TAG_EMAIL_ENDTIME, ID_EMAIL_ENDTIME)
	result.AddTag(TAG_EMAIL_INSTANCETYPE, ID_EMAIL_INSTANCETYPE)
	result.AddTag(TAG_EMAIL_BUSYSTATUS, ID_EMAIL_BUSYSTATUS)
	result.AddTag(TAG_EMAIL_LOCATION, ID_EMAIL_LOCATION)
	result.AddTag(TAG_EMAIL_MEETINGREQUEST, ID_EMAIL_MEETINGREQUEST)
	result.AddTag(TAG_EMAIL_ORGANIZER, ID_EMAIL_ORGANIZER)
	result.AddTag(TAG_EMAIL_RECURRENCEID, ID_EMAIL_RECURRENCEID)
	result.AddTag(TAG_EMAIL_REMINDER, ID_EMAIL_REMINDER)
	result.AddTag(TAG_EMAIL_RESPONSEREQUESTED, ID_EMAIL_RESPONSEREQUESTED)
	result.AddTag(TAG_EMAIL_RECURRENCES, ID_EMAIL_RECURRENCES)
	result.AddTag(TAG_EMAIL_RECURRENCE, ID_EMAIL_RECURRENCE)
	result.AddTag(TAG_EMAIL_RECURRENCE_TYPE, ID_EMAIL_RECURRENCE_TYPE)
	result.AddTag(TAG_EMAIL_RECURRENCE_UNTIL, ID_EMAIL_RECURRENCE_UNTIL)
	result.AddTag(TAG_EMAIL_RECURRENCE_OCCURRENCES, ID_EMAIL_RECURRENCE_OCCURRENCES)
	result.AddTag(TAG_EMAIL_RECURRENCE_INTERVAL, ID_EMAIL_RECURRENCE_INTERVAL)
	result.AddTag(TAG_EMAIL_RECURRENCE_DAYOFWEEK, ID_EMAIL_RECURRENCE_DAYOFWEEK)
	result.AddTag(TAG_EMAIL_RECURRENCE_DAYOFMONTH, ID_EMAIL_RECURRENCE_DAYOFMONTH)
	result.AddTag(TAG_EMAIL_RECURRENCE_WEEKOFMONTH, ID_EMAIL_RECURRENCE_WEEKOFMONTH)
	result.AddTag(TAG_EMAIL_RECURRENCE_MONTHOFYEAR, ID_EMAIL_RECURRENCE_MONTHOFYEAR)
	result.AddTag(TAG_EMAIL_STARTTIME, ID_EMAIL_STARTTIME)
	result.AddTag(TAG_EMAIL_SENSITIVITY, ID_EMAIL_SENSITIVITY)
	result.AddTag(TAG_EMAIL_TIMEZONE, ID_EMAIL_TIMEZONE)
	result.AddTag(TAG_EMAIL_GLOBALOBJID, ID_EMAIL_GLOBALOBJID)
	result.AddTag(TAG_EMAIL_THREADTOPIC, ID_EMAIL_THREADTOPIC)
	result.AddTag(TAG_EMAIL_INTERNETCPID, ID_EMAIL_INTERNETCPID)
	result.AddTag(TAG_EMAIL_FLAG, ID_EMAIL_FLAG)
	result.AddTag(TAG_EMAIL_STATUS, ID_EMAIL_STATUS)
	result.AddTag(TAG_EMAIL_CONTENTCLASS, ID_EMAIL_CONTENTCLASS)
	result.AddTag(TAG_EMAIL_FLAGTYPE, ID_EMAIL_FLAGTYPE)
	result.AddTag(TAG_EMAIL_COMPLETETIME, ID_EMAIL_COMPLETETIME)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_EMAIL_DISALLOWNEWTIMEPROPOSAL, ID_EMAIL_DISALLOWNEWTIMEPROPOSAL) // not supported with MS-ASProtocolVersion 12.1
	}

	return result
}

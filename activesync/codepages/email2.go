package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_EMAIL2                        string = "Email2"
	TAG_EMAIL2_UMCALLERID            string = "UmCallerID"
	TAG_EMAIL2_UMUSERNOTES           string = "UmUserNotes"
	TAG_EMAIL2_UMATTDURATION         string = "UmAttDuration"
	TAG_EMAIL2_UMATTORDER            string = "UmAttOrder"
	TAG_EMAIL2_CONVERSATIONID        string = "ConversationId"
	TAG_EMAIL2_CONVERSATIONINDEX     string = "ConversationIndex"
	TAG_EMAIL2_LASTVERBEXECUTED      string = "LastVerbExecuted"
	TAG_EMAIL2_LASTVERBEXECUTIONTIME string = "LastVerbExecutionTime"
	TAG_EMAIL2_RECEIVEDASBCC         string = "ReceivedAsBcc"
	TAG_EMAIL2_SENDER                string = "Sender"
	TAG_EMAIL2_CALENDARTYPE          string = "CalendarType"
	TAG_EMAIL2_ISLEAPMONTH           string = "IsLeapMonth"
	TAG_EMAIL2_ACCOUNTID             string = "AccountId"
	TAG_EMAIL2_FIRSTDAYOFWEEK        string = "FirstDayOfWeek"
	TAG_EMAIL2_MEETINGMESSAGETYPE    string = "MeetingMessageType"
)

const (
	CP_EMAIL2                       byte = 22
	ID_EMAIL2_UMCALLERID            byte = 0x05
	ID_EMAIL2_UMUSERNOTES           byte = 0x06
	ID_EMAIL2_UMATTDURATION         byte = 0x07
	ID_EMAIL2_UMATTORDER            byte = 0x08
	ID_EMAIL2_CONVERSATIONID        byte = 0x09
	ID_EMAIL2_CONVERSATIONINDEX     byte = 0x0A
	ID_EMAIL2_LASTVERBEXECUTED      byte = 0x0B
	ID_EMAIL2_LASTVERBEXECUTIONTIME byte = 0x0C
	ID_EMAIL2_RECEIVEDASBCC         byte = 0x0D
	ID_EMAIL2_SENDER                byte = 0x0E
	ID_EMAIL2_CALENDARTYPE          byte = 0x0F
	ID_EMAIL2_ISLEAPMONTH           byte = 0x10
	ID_EMAIL2_ACCOUNTID             byte = 0x11
	ID_EMAIL2_FIRSTDAYOFWEEK        byte = 0x12
	ID_EMAIL2_MEETINGMESSAGETYPE    byte = 0x13
)

// not supported with MS-ASProtocolVersion 12.1
func Email2(protocolVersion byte) CodePage {
	result := NewCodePage(NS_EMAIL2, CP_EMAIL2)

	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_EMAIL2_UMCALLERID, ID_EMAIL2_UMCALLERID)
		result.AddTag(TAG_EMAIL2_UMUSERNOTES, ID_EMAIL2_UMUSERNOTES)
		result.AddTag(TAG_EMAIL2_UMATTDURATION, ID_EMAIL2_UMATTDURATION)
		result.AddTag(TAG_EMAIL2_UMATTORDER, ID_EMAIL2_UMATTORDER)
		result.AddTag(TAG_EMAIL2_CONVERSATIONID, ID_EMAIL2_CONVERSATIONID)
		result.AddTag(TAG_EMAIL2_CONVERSATIONINDEX, ID_EMAIL2_CONVERSATIONINDEX)
		result.AddTag(TAG_EMAIL2_LASTVERBEXECUTED, ID_EMAIL2_LASTVERBEXECUTED)
		result.AddTag(TAG_EMAIL2_LASTVERBEXECUTIONTIME, ID_EMAIL2_LASTVERBEXECUTIONTIME)
		result.AddTag(TAG_EMAIL2_RECEIVEDASBCC, ID_EMAIL2_RECEIVEDASBCC)
		result.AddTag(TAG_EMAIL2_SENDER, ID_EMAIL2_SENDER)
		result.AddTag(TAG_EMAIL2_CALENDARTYPE, ID_EMAIL2_CALENDARTYPE)
		result.AddTag(TAG_EMAIL2_ISLEAPMONTH, ID_EMAIL2_ISLEAPMONTH)
		if protocolVersion > PROTOCOL_VERSION_14_0 {
			result.AddTag(TAG_EMAIL2_ACCOUNTID, ID_EMAIL2_ACCOUNTID)                   // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_EMAIL2_FIRSTDAYOFWEEK, ID_EMAIL2_FIRSTDAYOFWEEK)         // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_EMAIL2_MEETINGMESSAGETYPE, ID_EMAIL2_MEETINGMESSAGETYPE) // not supported with MS-ASProtocolVersion 12.1 or 14.0
		}
	}

	return result
}

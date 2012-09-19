package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
	"testing"
)

func checkEmail2CommonTags(t *testing.T, cp CodePage) {
	checkTagExists(t, TAG_EMAIL2_UMCALLERID, ID_EMAIL2_UMCALLERID, cp)
	checkTagExists(t, TAG_EMAIL2_UMUSERNOTES, ID_EMAIL2_UMUSERNOTES, cp)
	checkTagExists(t, TAG_EMAIL2_UMATTDURATION, ID_EMAIL2_UMATTDURATION, cp)
	checkTagExists(t, TAG_EMAIL2_UMATTORDER, ID_EMAIL2_UMATTORDER, cp)
	checkTagExists(t, TAG_EMAIL2_CONVERSATIONID, ID_EMAIL2_CONVERSATIONID, cp)
	checkTagExists(t, TAG_EMAIL2_CONVERSATIONINDEX, ID_EMAIL2_CONVERSATIONINDEX, cp)
	checkTagExists(t, TAG_EMAIL2_LASTVERBEXECUTED, ID_EMAIL2_LASTVERBEXECUTED, cp)
	checkTagExists(t, TAG_EMAIL2_LASTVERBEXECUTIONTIME, ID_EMAIL2_LASTVERBEXECUTIONTIME, cp)
	checkTagExists(t, TAG_EMAIL2_RECEIVEDASBCC, ID_EMAIL2_RECEIVEDASBCC, cp)
	checkTagExists(t, TAG_EMAIL2_SENDER, ID_EMAIL2_SENDER, cp)
	checkTagExists(t, TAG_EMAIL2_CALENDARTYPE, ID_EMAIL2_CALENDARTYPE, cp)
	checkTagExists(t, TAG_EMAIL2_ISLEAPMONTH, ID_EMAIL2_ISLEAPMONTH, cp)
}

func Test_Email2_120(t *testing.T) {
	cp := Email2(PROTOCOL_VERSION_12_0)
	checkTagCount(t, 0, cp)
}

func Test_Email2_121(t *testing.T) {
	cp := Email2(PROTOCOL_VERSION_12_1)
	checkTagCount(t, 0, cp)
}

func Test_Email2_140(t *testing.T) {
	cp := Email2(PROTOCOL_VERSION_14_0)
	checkEmail2CommonTags(t, cp)
	checkTagCount(t, 12, cp)
	checkTagNotExists(t, TAG_EMAIL2_ACCOUNTID, ID_EMAIL2_ACCOUNTID, cp)
	checkTagNotExists(t, TAG_EMAIL2_FIRSTDAYOFWEEK, ID_EMAIL2_FIRSTDAYOFWEEK, cp)
	checkTagNotExists(t, TAG_EMAIL2_MEETINGMESSAGETYPE, ID_EMAIL2_MEETINGMESSAGETYPE, cp)
}

func Test_Email2_141(t *testing.T) {
	cp := Email2(PROTOCOL_VERSION_14_1)
	checkEmail2CommonTags(t, cp)
	checkTagCount(t, 15, cp)
	checkTagExists(t, TAG_EMAIL2_ACCOUNTID, ID_EMAIL2_ACCOUNTID, cp)
	checkTagExists(t, TAG_EMAIL2_FIRSTDAYOFWEEK, ID_EMAIL2_FIRSTDAYOFWEEK, cp)
	checkTagExists(t, TAG_EMAIL2_MEETINGMESSAGETYPE, ID_EMAIL2_MEETINGMESSAGETYPE, cp)
}

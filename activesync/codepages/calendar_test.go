package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
	"testing"
)

func checkCalendarCommonTags(t *testing.T, cp CodePage) {
	checkTagExists(t, TAG_CALENDAR_TIMEZONE, ID_CALENDAR_TIMEZONE, cp)
	checkTagExists(t, TAG_CALENDAR_ALLDAYEVENT, ID_CALENDAR_ALLDAYEVENT, cp)
	checkTagExists(t, TAG_CALENDAR_ATTENDEES, ID_CALENDAR_ATTENDEES, cp)
	checkTagExists(t, TAG_CALENDAR_ATTENDEE, ID_CALENDAR_ATTENDEE, cp)
	checkTagExists(t, TAG_CALENDAR_EMAIL, ID_CALENDAR_EMAIL, cp)
	checkTagExists(t, TAG_CALENDAR_NAME, ID_CALENDAR_NAME, cp)
	checkTagExists(t, TAG_CALENDAR_BUSYSTATUS, ID_CALENDAR_BUSYSTATUS, cp)
	checkTagExists(t, TAG_CALENDAR_CATEGORIES, ID_CALENDAR_CATEGORIES, cp)
	checkTagExists(t, TAG_CALENDAR_CATEGORY, ID_CALENDAR_CATEGORY, cp)
	checkTagExists(t, TAG_CALENDAR_DTSTAMP, ID_CALENDAR_DTSTAMP, cp)
	checkTagExists(t, TAG_CALENDAR_ENDTIME, ID_CALENDAR_ENDTIME, cp)
	checkTagExists(t, TAG_CALENDAR_EXCEPTION, ID_CALENDAR_EXCEPTION, cp)
	checkTagExists(t, TAG_CALENDAR_EXCEPTIONS, ID_CALENDAR_EXCEPTIONS, cp)
	checkTagExists(t, TAG_CALENDAR_DELETED, ID_CALENDAR_DELETED, cp)
	checkTagExists(t, TAG_CALENDAR_EXCEPTIONSTARTTIME, ID_CALENDAR_EXCEPTIONSTARTTIME, cp)
	checkTagExists(t, TAG_CALENDAR_LOCATION, ID_CALENDAR_LOCATION, cp)
	checkTagExists(t, TAG_CALENDAR_MEETINGSTATUS, ID_CALENDAR_MEETINGSTATUS, cp)
	checkTagExists(t, TAG_CALENDAR_ORGANIZEREMAIL, ID_CALENDAR_ORGANIZEREMAIL, cp)
	checkTagExists(t, TAG_CALENDAR_ORGANIZERNAME, ID_CALENDAR_ORGANIZERNAME, cp)
	checkTagExists(t, TAG_CALENDAR_RECURRENCE, ID_CALENDAR_RECURRENCE, cp)
	checkTagExists(t, TAG_CALENDAR_TYPE, ID_CALENDAR_TYPE, cp)
	checkTagExists(t, TAG_CALENDAR_UNTIL, ID_CALENDAR_UNTIL, cp)
	checkTagExists(t, TAG_CALENDAR_OCCURRENCES, ID_CALENDAR_OCCURRENCES, cp)
	checkTagExists(t, TAG_CALENDAR_INTERVAL, ID_CALENDAR_INTERVAL, cp)
	checkTagExists(t, TAG_CALENDAR_DAYOFWEEK, ID_CALENDAR_DAYOFWEEK, cp)
	checkTagExists(t, TAG_CALENDAR_DAYOFMONTH, ID_CALENDAR_DAYOFMONTH, cp)
	checkTagExists(t, TAG_CALENDAR_WEEKOFMONTH, ID_CALENDAR_WEEKOFMONTH, cp)
	checkTagExists(t, TAG_CALENDAR_MONTHOFYEAR, ID_CALENDAR_MONTHOFYEAR, cp)
	checkTagExists(t, TAG_CALENDAR_REMINDER, ID_CALENDAR_REMINDER, cp)
	checkTagExists(t, TAG_CALENDAR_SENSITIVITY, ID_CALENDAR_SENSITIVITY, cp)
	checkTagExists(t, TAG_CALENDAR_SUBJECT, ID_CALENDAR_SUBJECT, cp)
	checkTagExists(t, TAG_CALENDAR_STARTTIME, ID_CALENDAR_STARTTIME, cp)
	checkTagExists(t, TAG_CALENDAR_UID, ID_CALENDAR_UID, cp)
	checkTagExists(t, TAG_CALENDAR_ATTENDEESTATUS, ID_CALENDAR_ATTENDEESTATUS, cp)
	checkTagExists(t, TAG_CALENDAR_ATTENDEETYPE, ID_CALENDAR_ATTENDEETYPE, cp)
}

func Test_Calendar_120(t *testing.T) {
	cp := Calendar(PROTOCOL_VERSION_12_0)
	checkCalendarCommonTags(t, cp)
	checkTagCount(t, 35, cp)
	checkTagNotExists(t, TAG_CALENDAR_DISALLOWNEWTIMEPROPOSAL, ID_CALENDAR_DISALLOWNEWTIMEPROPOSAL, cp)
	checkTagNotExists(t, TAG_CALENDAR_RESPONSEREQUESTED, ID_CALENDAR_RESPONSEREQUESTED, cp)
	checkTagNotExists(t, TAG_CALENDAR_APPOINTMENTREPLYTIME, ID_CALENDAR_APPOINTMENTREPLYTIME, cp)
	checkTagNotExists(t, TAG_CALENDAR_RESPONSETYPE, ID_CALENDAR_RESPONSETYPE, cp)
	checkTagNotExists(t, TAG_CALENDAR_CALENDARTYPE, ID_CALENDAR_CALENDARTYPE, cp)
	checkTagNotExists(t, TAG_CALENDAR_ISLEAPMONTH, ID_CALENDAR_ISLEAPMONTH, cp)
	checkTagNotExists(t, TAG_CALENDAR_FIRSTDAYOFWEEK, ID_CALENDAR_FIRSTDAYOFWEEK, cp)
	checkTagNotExists(t, TAG_CALENDAR_ONLINEMEETINGCONFLINK, ID_CALENDAR_ONLINEMEETINGCONFLINK, cp)
	checkTagNotExists(t, TAG_CALENDAR_ONLINEMEETINGEXTERNALLINK, ID_CALENDAR_ONLINEMEETINGEXTERNALLINK, cp)
}

func Test_Calendar_121(t *testing.T) {
	cp := Calendar(PROTOCOL_VERSION_12_1)
	checkCalendarCommonTags(t, cp)
	checkTagCount(t, 35, cp)
	checkTagNotExists(t, TAG_CALENDAR_DISALLOWNEWTIMEPROPOSAL, ID_CALENDAR_DISALLOWNEWTIMEPROPOSAL, cp)
	checkTagNotExists(t, TAG_CALENDAR_RESPONSEREQUESTED, ID_CALENDAR_RESPONSEREQUESTED, cp)
	checkTagNotExists(t, TAG_CALENDAR_APPOINTMENTREPLYTIME, ID_CALENDAR_APPOINTMENTREPLYTIME, cp)
	checkTagNotExists(t, TAG_CALENDAR_RESPONSETYPE, ID_CALENDAR_RESPONSETYPE, cp)
	checkTagNotExists(t, TAG_CALENDAR_CALENDARTYPE, ID_CALENDAR_CALENDARTYPE, cp)
	checkTagNotExists(t, TAG_CALENDAR_ISLEAPMONTH, ID_CALENDAR_ISLEAPMONTH, cp)
	checkTagNotExists(t, TAG_CALENDAR_FIRSTDAYOFWEEK, ID_CALENDAR_FIRSTDAYOFWEEK, cp)
	checkTagNotExists(t, TAG_CALENDAR_ONLINEMEETINGCONFLINK, ID_CALENDAR_ONLINEMEETINGCONFLINK, cp)
	checkTagNotExists(t, TAG_CALENDAR_ONLINEMEETINGEXTERNALLINK, ID_CALENDAR_ONLINEMEETINGEXTERNALLINK, cp)
}

func Test_Calendar_140(t *testing.T) {
	cp := Calendar(PROTOCOL_VERSION_14_0)
	checkCalendarCommonTags(t, cp)
	checkTagCount(t, 41, cp)
	checkTagExists(t, TAG_CALENDAR_DISALLOWNEWTIMEPROPOSAL, ID_CALENDAR_DISALLOWNEWTIMEPROPOSAL, cp)
	checkTagExists(t, TAG_CALENDAR_RESPONSEREQUESTED, ID_CALENDAR_RESPONSEREQUESTED, cp)
	checkTagExists(t, TAG_CALENDAR_APPOINTMENTREPLYTIME, ID_CALENDAR_APPOINTMENTREPLYTIME, cp)
	checkTagExists(t, TAG_CALENDAR_RESPONSETYPE, ID_CALENDAR_RESPONSETYPE, cp)
	checkTagExists(t, TAG_CALENDAR_CALENDARTYPE, ID_CALENDAR_CALENDARTYPE, cp)
	checkTagExists(t, TAG_CALENDAR_ISLEAPMONTH, ID_CALENDAR_ISLEAPMONTH, cp)
	checkTagNotExists(t, TAG_CALENDAR_FIRSTDAYOFWEEK, ID_CALENDAR_FIRSTDAYOFWEEK, cp)
	checkTagNotExists(t, TAG_CALENDAR_ONLINEMEETINGCONFLINK, ID_CALENDAR_ONLINEMEETINGCONFLINK, cp)
	checkTagNotExists(t, TAG_CALENDAR_ONLINEMEETINGEXTERNALLINK, ID_CALENDAR_ONLINEMEETINGEXTERNALLINK, cp)
}

func Test_Calendar_141(t *testing.T) {
	cp := Calendar(PROTOCOL_VERSION_14_1)
	checkCalendarCommonTags(t, cp)
	checkTagCount(t, 44, cp)
	checkTagExists(t, TAG_CALENDAR_DISALLOWNEWTIMEPROPOSAL, ID_CALENDAR_DISALLOWNEWTIMEPROPOSAL, cp)
	checkTagExists(t, TAG_CALENDAR_RESPONSEREQUESTED, ID_CALENDAR_RESPONSEREQUESTED, cp)
	checkTagExists(t, TAG_CALENDAR_APPOINTMENTREPLYTIME, ID_CALENDAR_APPOINTMENTREPLYTIME, cp)
	checkTagExists(t, TAG_CALENDAR_RESPONSETYPE, ID_CALENDAR_RESPONSETYPE, cp)
	checkTagExists(t, TAG_CALENDAR_CALENDARTYPE, ID_CALENDAR_CALENDARTYPE, cp)
	checkTagExists(t, TAG_CALENDAR_ISLEAPMONTH, ID_CALENDAR_ISLEAPMONTH, cp)
	checkTagExists(t, TAG_CALENDAR_FIRSTDAYOFWEEK, ID_CALENDAR_FIRSTDAYOFWEEK, cp)
	checkTagExists(t, TAG_CALENDAR_ONLINEMEETINGCONFLINK, ID_CALENDAR_ONLINEMEETINGCONFLINK, cp)
	checkTagExists(t, TAG_CALENDAR_ONLINEMEETINGEXTERNALLINK, ID_CALENDAR_ONLINEMEETINGEXTERNALLINK, cp)
}

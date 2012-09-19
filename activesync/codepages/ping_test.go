package codepages

import (
	"testing"
)

func Test_Ping(t *testing.T) {
	cp := Ping()
	checkTagCount(t, 9, cp)
	checkTagExists(t, TAG_PING_PING, ID_PING_PING, cp)
	checkTagExists(t, TAG_PING_AUTDSTATE, ID_PING_AUTDSTATE, cp)
	checkTagExists(t, TAG_PING_STATUS, ID_PING_STATUS, cp)
	checkTagExists(t, TAG_PING_HEARTBEATINTERVAL, ID_PING_HEARTBEATINTERVAL, cp)
	checkTagExists(t, TAG_PING_FOLDERS, ID_PING_FOLDERS, cp)
	checkTagExists(t, TAG_PING_FOLDER, ID_PING_FOLDER, cp)
	checkTagExists(t, TAG_PING_ID, ID_PING_ID, cp)
	checkTagExists(t, TAG_PING_CLASS, ID_PING_CLASS, cp)
	checkTagExists(t, TAG_PING_MAXFOLDERS, ID_PING_MAXFOLDERS, cp)
}

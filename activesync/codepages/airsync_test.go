package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
	"testing"
)

func checkAirSyncCommonTags(t *testing.T, cp CodePage) {
	checkTagExists(t, TAG_AIRSYNC_SYNC, ID_AIRSYNC_SYNC, cp)
	checkTagExists(t, TAG_AIRSYNC_RESPONSES, ID_AIRSYNC_RESPONSES, cp)
	checkTagExists(t, TAG_AIRSYNC_ADD, ID_AIRSYNC_ADD, cp)
	checkTagExists(t, TAG_AIRSYNC_CHANGE, ID_AIRSYNC_CHANGE, cp)
	checkTagExists(t, TAG_AIRSYNC_DELETE, ID_AIRSYNC_DELETE, cp)
	checkTagExists(t, TAG_AIRSYNC_FETCH, ID_AIRSYNC_FETCH, cp)
	checkTagExists(t, TAG_AIRSYNC_SYNCKEY, ID_AIRSYNC_SYNCKEY, cp)
	checkTagExists(t, TAG_AIRSYNC_CLIENTID, ID_AIRSYNC_CLIENTID, cp)
	checkTagExists(t, TAG_AIRSYNC_SERVERID, ID_AIRSYNC_SERVERID, cp)
	checkTagExists(t, TAG_AIRSYNC_STATUS, ID_AIRSYNC_STATUS, cp)
	checkTagExists(t, TAG_AIRSYNC_COLLECTION, ID_AIRSYNC_COLLECTION, cp)
	checkTagExists(t, TAG_AIRSYNC_CLASS, ID_AIRSYNC_CLASS, cp)
	checkTagExists(t, TAG_AIRSYNC_COLLECTIONID, ID_AIRSYNC_COLLECTIONID, cp)
	checkTagExists(t, TAG_AIRSYNC_GETCHANGES, ID_AIRSYNC_GETCHANGES, cp)
	checkTagExists(t, TAG_AIRSYNC_MOREAVAILABLE, ID_AIRSYNC_MOREAVAILABLE, cp)
	checkTagExists(t, TAG_AIRSYNC_WINDOWSIZE, ID_AIRSYNC_WINDOWSIZE, cp)
	checkTagExists(t, TAG_AIRSYNC_COMMANDS, ID_AIRSYNC_COMMANDS, cp)
	checkTagExists(t, TAG_AIRSYNC_OPTIONS, ID_AIRSYNC_OPTIONS, cp)
	checkTagExists(t, TAG_AIRSYNC_FILTERTYPE, ID_AIRSYNC_FILTERTYPE, cp)
	checkTagExists(t, TAG_AIRSYNC_CONFLICT, ID_AIRSYNC_CONFLICT, cp)
	checkTagExists(t, TAG_AIRSYNC_COLLECTIONS, ID_AIRSYNC_COLLECTIONS, cp)
	checkTagExists(t, TAG_AIRSYNC_APPLICATIONDATA, ID_AIRSYNC_APPLICATIONDATA, cp)
	checkTagExists(t, TAG_AIRSYNC_DELETESASMOVES, ID_AIRSYNC_DELETESASMOVES, cp)
	checkTagExists(t, TAG_AIRSYNC_SUPPORTED, ID_AIRSYNC_SUPPORTED, cp)
	checkTagExists(t, TAG_AIRSYNC_SOFTDELETE, ID_AIRSYNC_SOFTDELETE, cp)
	checkTagExists(t, TAG_AIRSYNC_MIMESUPPORT, ID_AIRSYNC_MIMESUPPORT, cp)
	checkTagExists(t, TAG_AIRSYNC_MIMETRUNCATION, ID_AIRSYNC_MIMETRUNCATION, cp)
	checkTagExists(t, TAG_AIRSYNC_WAIT, ID_AIRSYNC_WAIT, cp)
	checkTagExists(t, TAG_AIRSYNC_LIMIT, ID_AIRSYNC_LIMIT, cp)
	checkTagExists(t, TAG_AIRSYNC_PARTIAL, ID_AIRSYNC_PARTIAL, cp)
}

func Test_AirSync_120(t *testing.T) {
	cp := AirSync(PROTOCOL_VERSION_12_0)
	checkAirSyncCommonTags(t, cp)
	checkTagCount(t, 30, cp)
	checkTagNotExists(t, TAG_AIRSYNC_CONVERSATIONMODE, ID_AIRSYNC_CONVERSATIONMODE, cp)
	checkTagNotExists(t, TAG_AIRSYNC_MAXITEMS, ID_AIRSYNC_MAXITEMS, cp)
	checkTagNotExists(t, TAG_AIRSYNC_HEARTBEATINTERVAL, ID_AIRSYNC_HEARTBEATINTERVAL, cp)
}

func Test_AirSync_121(t *testing.T) {
	cp := AirSync(PROTOCOL_VERSION_12_1)
	checkAirSyncCommonTags(t, cp)
	checkTagCount(t, 30, cp)
	checkTagNotExists(t, TAG_AIRSYNC_CONVERSATIONMODE, ID_AIRSYNC_CONVERSATIONMODE, cp)
	checkTagNotExists(t, TAG_AIRSYNC_MAXITEMS, ID_AIRSYNC_MAXITEMS, cp)
	checkTagNotExists(t, TAG_AIRSYNC_HEARTBEATINTERVAL, ID_AIRSYNC_HEARTBEATINTERVAL, cp)
}

func Test_AirSync_140(t *testing.T) {
	cp := AirSync(PROTOCOL_VERSION_14_0)
	checkAirSyncCommonTags(t, cp)
	checkTagCount(t, 33, cp)
	checkTagExists(t, TAG_AIRSYNC_CONVERSATIONMODE, ID_AIRSYNC_CONVERSATIONMODE, cp)
	checkTagExists(t, TAG_AIRSYNC_MAXITEMS, ID_AIRSYNC_MAXITEMS, cp)
	checkTagExists(t, TAG_AIRSYNC_HEARTBEATINTERVAL, ID_AIRSYNC_HEARTBEATINTERVAL, cp)
}

func Test_AirSync_141(t *testing.T) {
	cp := AirSync(PROTOCOL_VERSION_14_1)
	checkAirSyncCommonTags(t, cp)
	checkTagCount(t, 33, cp)
	checkTagExists(t, TAG_AIRSYNC_CONVERSATIONMODE, ID_AIRSYNC_CONVERSATIONMODE, cp)
	checkTagExists(t, TAG_AIRSYNC_MAXITEMS, ID_AIRSYNC_MAXITEMS, cp)
	checkTagExists(t, TAG_AIRSYNC_HEARTBEATINTERVAL, ID_AIRSYNC_HEARTBEATINTERVAL, cp)
}

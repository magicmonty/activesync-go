package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
	"testing"
)

func checkItemEstimateCommonTags(t *testing.T, cp CodePage) {
	checkTagExists(t, TAG_ITEM_ESTIMATE_GETITEMESTIMATE, ID_ITEM_ESTIMATE_GETITEMESTIMATE, cp)
	checkTagExists(t, TAG_ITEM_ESTIMATE_COLLECTIONS, ID_ITEM_ESTIMATE_COLLECTIONS, cp)
	checkTagExists(t, TAG_ITEM_ESTIMATE_COLLECTION, ID_ITEM_ESTIMATE_COLLECTION, cp)
	checkTagExists(t, TAG_ITEM_ESTIMATE_COLLECTIONID, ID_ITEM_ESTIMATE_COLLECTIONID, cp)
	checkTagExists(t, TAG_ITEM_ESTIMATE_ESTIMATE, ID_ITEM_ESTIMATE_ESTIMATE, cp)
	checkTagExists(t, TAG_ITEM_ESTIMATE_RESPONSE, ID_ITEM_ESTIMATE_RESPONSE, cp)
	checkTagExists(t, TAG_ITEM_ESTIMATE_STATUS, ID_ITEM_ESTIMATE_STATUS, cp)
}

func Test_ItemEstimate_120(t *testing.T) {
	cp := ItemEstimate(PROTOCOL_VERSION_12_0)
	checkItemEstimateCommonTags(t, cp)
	checkTagCount(t, 10, cp)
	checkTagExists(t, TAG_ITEM_ESTIMATE_VERSION, ID_ITEM_ESTIMATE_VERSION, cp)
	checkTagExists(t, TAG_ITEM_ESTIMATE_CLASS, ID_ITEM_ESTIMATE_CLASS, cp)
	checkTagExists(t, TAG_ITEM_ESTIMATE_DATETIME, ID_ITEM_ESTIMATE_DATETIME, cp)
}

func Test_ItemEstimate_121(t *testing.T) {
	cp := ItemEstimate(PROTOCOL_VERSION_12_1)
	checkItemEstimateCommonTags(t, cp)
	checkTagCount(t, 10, cp)
	checkTagExists(t, TAG_ITEM_ESTIMATE_VERSION, ID_ITEM_ESTIMATE_VERSION, cp)
	checkTagExists(t, TAG_ITEM_ESTIMATE_CLASS, ID_ITEM_ESTIMATE_CLASS, cp)
	checkTagExists(t, TAG_ITEM_ESTIMATE_DATETIME, ID_ITEM_ESTIMATE_DATETIME, cp)
}

func Test_ItemEstimate_140(t *testing.T) {
	cp := ItemEstimate(PROTOCOL_VERSION_14_0)
	checkItemEstimateCommonTags(t, cp)
	checkTagCount(t, 7, cp)
	checkTagNotExists(t, TAG_ITEM_ESTIMATE_VERSION, ID_ITEM_ESTIMATE_VERSION, cp)
	checkTagNotExists(t, TAG_ITEM_ESTIMATE_CLASS, ID_ITEM_ESTIMATE_CLASS, cp)
	checkTagNotExists(t, TAG_ITEM_ESTIMATE_DATETIME, ID_ITEM_ESTIMATE_DATETIME, cp)
}

func Test_ItemEstimate_141(t *testing.T) {
	cp := ItemEstimate(PROTOCOL_VERSION_14_1)
	checkItemEstimateCommonTags(t, cp)
	checkTagCount(t, 7, cp)
	checkTagNotExists(t, TAG_ITEM_ESTIMATE_VERSION, ID_ITEM_ESTIMATE_VERSION, cp)
	checkTagNotExists(t, TAG_ITEM_ESTIMATE_CLASS, ID_ITEM_ESTIMATE_CLASS, cp)
	checkTagNotExists(t, TAG_ITEM_ESTIMATE_DATETIME, ID_ITEM_ESTIMATE_DATETIME, cp)
}

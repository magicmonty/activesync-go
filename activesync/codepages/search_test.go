package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
	"testing"
)

func checkSearchCommonTags(t *testing.T, cp CodePage) {
	checkTagExists(t, TAG_SEARCH_SEARCH, ID_SEARCH_SEARCH, cp)
	checkTagExists(t, TAG_SEARCH_STORE, ID_SEARCH_STORE, cp)
	checkTagExists(t, TAG_SEARCH_NAME, ID_SEARCH_NAME, cp)
	checkTagExists(t, TAG_SEARCH_QUERY, ID_SEARCH_QUERY, cp)
	checkTagExists(t, TAG_SEARCH_OPTIONS, ID_SEARCH_OPTIONS, cp)
	checkTagExists(t, TAG_SEARCH_RANGE, ID_SEARCH_RANGE, cp)
	checkTagExists(t, TAG_SEARCH_STATUS, ID_SEARCH_STATUS, cp)
	checkTagExists(t, TAG_SEARCH_RESPONSE, ID_SEARCH_RESPONSE, cp)
	checkTagExists(t, TAG_SEARCH_RESULT, ID_SEARCH_RESULT, cp)
	checkTagExists(t, TAG_SEARCH_PROPERTIES, ID_SEARCH_PROPERTIES, cp)
	checkTagExists(t, TAG_SEARCH_TOTAL, ID_SEARCH_TOTAL, cp)
	checkTagExists(t, TAG_SEARCH_EQUALTO, ID_SEARCH_EQUALTO, cp)
	checkTagExists(t, TAG_SEARCH_VALUE, ID_SEARCH_VALUE, cp)
	checkTagExists(t, TAG_SEARCH_AND, ID_SEARCH_AND, cp)
	checkTagExists(t, TAG_SEARCH_FREETEXT, ID_SEARCH_FREETEXT, cp)
	checkTagExists(t, TAG_SEARCH_DEEPTRAVERSAL, ID_SEARCH_DEEPTRAVERSAL, cp)
	checkTagExists(t, TAG_SEARCH_LONGID, ID_SEARCH_LONGID, cp)
	checkTagExists(t, TAG_SEARCH_REBUILDRESULTS, ID_SEARCH_REBUILDRESULTS, cp)
	checkTagExists(t, TAG_SEARCH_LESSTHAN, ID_SEARCH_LESSTHAN, cp)
	checkTagExists(t, TAG_SEARCH_GREATERTHAN, ID_SEARCH_GREATERTHAN, cp)
	checkTagExists(t, TAG_SEARCH_USERNAME, ID_SEARCH_USERNAME, cp)
	checkTagExists(t, TAG_SEARCH_PASSWORD, ID_SEARCH_PASSWORD, cp)
	checkTagExists(t, TAG_SEARCH_CONVERSATIONID, ID_SEARCH_CONVERSATIONID, cp)
}

func Test_Search_120(t *testing.T) {
	cp := Search(PROTOCOL_VERSION_12_0)
	checkSearchCommonTags(t, cp)
	checkTagCount(t, 23, cp)
	checkTagNotExists(t, TAG_SEARCH_OR, ID_SEARCH_OR, cp)
	checkTagNotExists(t, TAG_SEARCH_PICTURE, ID_SEARCH_PICTURE, cp)
	checkTagNotExists(t, TAG_SEARCH_MAXSIZE, ID_SEARCH_MAXSIZE, cp)
	checkTagNotExists(t, TAG_SEARCH_MAXPICTURES, ID_SEARCH_MAXPICTURES, cp)
}

func Test_Search_121(t *testing.T) {
	cp := Search(PROTOCOL_VERSION_12_1)
	checkSearchCommonTags(t, cp)
	checkTagCount(t, 23, cp)
	checkTagNotExists(t, TAG_SEARCH_OR, ID_SEARCH_OR, cp)
	checkTagNotExists(t, TAG_SEARCH_PICTURE, ID_SEARCH_PICTURE, cp)
	checkTagNotExists(t, TAG_SEARCH_MAXSIZE, ID_SEARCH_MAXSIZE, cp)
	checkTagNotExists(t, TAG_SEARCH_MAXPICTURES, ID_SEARCH_MAXPICTURES, cp)
}

func Test_Search_140(t *testing.T) {
	cp := Search(PROTOCOL_VERSION_14_0)
	checkSearchCommonTags(t, cp)
	checkTagCount(t, 24, cp)
	checkTagExists(t, TAG_SEARCH_OR, ID_SEARCH_OR, cp)
	checkTagNotExists(t, TAG_SEARCH_PICTURE, ID_SEARCH_PICTURE, cp)
	checkTagNotExists(t, TAG_SEARCH_MAXSIZE, ID_SEARCH_MAXSIZE, cp)
	checkTagNotExists(t, TAG_SEARCH_MAXPICTURES, ID_SEARCH_MAXPICTURES, cp)
}

func Test_Search_141(t *testing.T) {
	cp := Search(PROTOCOL_VERSION_14_1)
	checkSearchCommonTags(t, cp)
	checkTagCount(t, 27, cp)
	checkTagExists(t, TAG_SEARCH_OR, ID_SEARCH_OR, cp)
	checkTagExists(t, TAG_SEARCH_PICTURE, ID_SEARCH_PICTURE, cp)
	checkTagExists(t, TAG_SEARCH_MAXSIZE, ID_SEARCH_MAXSIZE, cp)
	checkTagExists(t, TAG_SEARCH_MAXPICTURES, ID_SEARCH_MAXPICTURES, cp)
}

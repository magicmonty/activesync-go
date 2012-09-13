package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_SEARCH                 string = "Search"
	TAG_SEARCH_SEARCH         string = "Search"
	TAG_SEARCH_STORE          string = "Store"
	TAG_SEARCH_NAME           string = "Name"
	TAG_SEARCH_QUERY          string = "Query"
	TAG_SEARCH_OPTIONS        string = "Options"
	TAG_SEARCH_RANGE          string = "Range"
	TAG_SEARCH_STATUS         string = "Status"
	TAG_SEARCH_RESPONSE       string = "Response"
	TAG_SEARCH_RESULT         string = "Result"
	TAG_SEARCH_PROPERTIES     string = "Properties"
	TAG_SEARCH_TOTAL          string = "Total"
	TAG_SEARCH_EQUALTO        string = "EqualTo"
	TAG_SEARCH_VALUE          string = "Value"
	TAG_SEARCH_AND            string = "And"
	TAG_SEARCH_OR             string = "Or"
	TAG_SEARCH_FREETEXT       string = "FreeText"
	TAG_SEARCH_DEEPTRAVERSAL  string = "DeepTraversal"
	TAG_SEARCH_LONGID         string = "LongId"
	TAG_SEARCH_REBUILDRESULTS string = "RebuildResults"
	TAG_SEARCH_LESSTHAN       string = "LessThan"
	TAG_SEARCH_GREATERTHAN    string = "GreaterThan"
	TAG_SEARCH_USERNAME       string = "UserName"
	TAG_SEARCH_PASSWORD       string = "Password"
	TAG_SEARCH_CONVERSATIONID string = "ConversationId"
	TAG_SEARCH_PICTURE        string = "Picture"
	TAG_SEARCH_MAXSIZE        string = "MaxSize"
	TAG_SEARCH_MAXPICTURES    string = "MaxPictures"
)

const (
	CP_SEARCH                byte = 15
	ID_SEARCH_SEARCH         byte = 0x05
	ID_SEARCH_STORE          byte = 0x07
	ID_SEARCH_NAME           byte = 0x08
	ID_SEARCH_QUERY          byte = 0x09
	ID_SEARCH_OPTIONS        byte = 0x0A
	ID_SEARCH_RANGE          byte = 0x0B
	ID_SEARCH_STATUS         byte = 0x0C
	ID_SEARCH_RESPONSE       byte = 0x0D
	ID_SEARCH_RESULT         byte = 0x0E
	ID_SEARCH_PROPERTIES     byte = 0x0F
	ID_SEARCH_TOTAL          byte = 0x10
	ID_SEARCH_EQUALTO        byte = 0x11
	ID_SEARCH_VALUE          byte = 0x12
	ID_SEARCH_AND            byte = 0x13
	ID_SEARCH_OR             byte = 0x14
	ID_SEARCH_FREETEXT       byte = 0x15
	ID_SEARCH_DEEPTRAVERSAL  byte = 0x17
	ID_SEARCH_LONGID         byte = 0x18
	ID_SEARCH_REBUILDRESULTS byte = 0x19
	ID_SEARCH_LESSTHAN       byte = 0x1A
	ID_SEARCH_GREATERTHAN    byte = 0x1B
	ID_SEARCH_USERNAME       byte = 0x1E
	ID_SEARCH_PASSWORD       byte = 0x1F
	ID_SEARCH_CONVERSATIONID byte = 0x20
	ID_SEARCH_PICTURE        byte = 0x21
	ID_SEARCH_MAXSIZE        byte = 0x22
	ID_SEARCH_MAXPICTURES    byte = 0x23
)

func Search(protocolVersion byte) CodePage {
	result := NewCodePage(NS_SEARCH, CP_SEARCH)

	result.AddTag(TAG_SEARCH_SEARCH, ID_SEARCH_SEARCH)
	result.AddTag(TAG_SEARCH_STORE, ID_SEARCH_STORE)
	result.AddTag(TAG_SEARCH_NAME, ID_SEARCH_NAME)
	result.AddTag(TAG_SEARCH_QUERY, ID_SEARCH_QUERY)
	result.AddTag(TAG_SEARCH_OPTIONS, ID_SEARCH_OPTIONS)
	result.AddTag(TAG_SEARCH_RANGE, ID_SEARCH_RANGE)
	result.AddTag(TAG_SEARCH_STATUS, ID_SEARCH_STATUS)
	result.AddTag(TAG_SEARCH_RESPONSE, ID_SEARCH_RESPONSE)
	result.AddTag(TAG_SEARCH_RESULT, ID_SEARCH_RESULT)
	result.AddTag(TAG_SEARCH_PROPERTIES, ID_SEARCH_PROPERTIES)
	result.AddTag(TAG_SEARCH_TOTAL, ID_SEARCH_TOTAL)
	result.AddTag(TAG_SEARCH_EQUALTO, ID_SEARCH_EQUALTO)
	result.AddTag(TAG_SEARCH_VALUE, ID_SEARCH_VALUE)
	result.AddTag(TAG_SEARCH_AND, ID_SEARCH_AND)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_SEARCH_OR, ID_SEARCH_OR) // not supported with MS-ASProtocolVersion 12.1
	}
	result.AddTag(TAG_SEARCH_FREETEXT, ID_SEARCH_FREETEXT)
	result.AddTag(TAG_SEARCH_DEEPTRAVERSAL, ID_SEARCH_DEEPTRAVERSAL)
	result.AddTag(TAG_SEARCH_LONGID, ID_SEARCH_LONGID)
	result.AddTag(TAG_SEARCH_REBUILDRESULTS, ID_SEARCH_REBUILDRESULTS)
	result.AddTag(TAG_SEARCH_LESSTHAN, ID_SEARCH_LESSTHAN)
	result.AddTag(TAG_SEARCH_GREATERTHAN, ID_SEARCH_GREATERTHAN)
	result.AddTag(TAG_SEARCH_USERNAME, ID_SEARCH_USERNAME)
	result.AddTag(TAG_SEARCH_PASSWORD, ID_SEARCH_PASSWORD)
	result.AddTag(TAG_SEARCH_CONVERSATIONID, ID_SEARCH_CONVERSATIONID)
	if protocolVersion > PROTOCOL_VERSION_14_0 {
		result.AddTag(TAG_SEARCH_PICTURE, ID_SEARCH_PICTURE)         // not supported with MS-ASProtocolVersion 12.1 or 14.0
		result.AddTag(TAG_SEARCH_MAXSIZE, ID_SEARCH_MAXSIZE)         // not supported with MS-ASProtocolVersion 12.1 or 14.0
		result.AddTag(TAG_SEARCH_MAXPICTURES, ID_SEARCH_MAXPICTURES) // not supported with MS-ASProtocolVersion 12.1 or 14.0
	}

	return result
}

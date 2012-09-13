package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_ITEM_OPERATIONS                      string = "ItemOperations"
	TAG_ITEM_OPERATIONS_ITEMOPERATIONS      string = "ItemOperations"
	TAG_ITEM_OPERATIONS_FETCH               string = "Fetch"
	TAG_ITEM_OPERATIONS_STORE               string = "Store"
	TAG_ITEM_OPERATIONS_OPTIONS             string = "Options"
	TAG_ITEM_OPERATIONS_RANGE               string = "Range"
	TAG_ITEM_OPERATIONS_TOTAL               string = "Total"
	TAG_ITEM_OPERATIONS_PROPERTIES          string = "Properties"
	TAG_ITEM_OPERATIONS_DATA                string = "Data"
	TAG_ITEM_OPERATIONS_STATUS              string = "Status"
	TAG_ITEM_OPERATIONS_RESPONSE            string = "Response"
	TAG_ITEM_OPERATIONS_VERSION             string = "Version"
	TAG_ITEM_OPERATIONS_SCHEMA              string = "Schema"
	TAG_ITEM_OPERATIONS_PART                string = "Part"
	TAG_ITEM_OPERATIONS_EMPTYFOLDERCONTENTS string = "EmptyFolderContents"
	TAG_ITEM_OPERATIONS_DELETESUBFOLDERS    string = "DeleteSubFolders"
	TAG_ITEM_OPERATIONS_USERNAME            string = "UserName"
	TAG_ITEM_OPERATIONS_PASSWORD            string = "Password"
	TAG_ITEM_OPERATIONS_MOVE                string = "Move"
	TAG_ITEM_OPERATIONS_DSTFLDID            string = "DstFldId"
	TAG_ITEM_OPERATIONS_CONVERSATIONID      string = "ConversationId"
	TAG_ITEM_OPERATIONS_MOVEALWAYS          string = "MoveAlways"
)

const (
	CP_ITEM_OPERATIONS                     byte = 20
	ID_ITEM_OPERATIONS_ITEMOPERATIONS      byte = 0x05
	ID_ITEM_OPERATIONS_FETCH               byte = 0x06
	ID_ITEM_OPERATIONS_STORE               byte = 0x07
	ID_ITEM_OPERATIONS_OPTIONS             byte = 0x08
	ID_ITEM_OPERATIONS_RANGE               byte = 0x09
	ID_ITEM_OPERATIONS_TOTAL               byte = 0x0A
	ID_ITEM_OPERATIONS_PROPERTIES          byte = 0x0B
	ID_ITEM_OPERATIONS_DATA                byte = 0x0C
	ID_ITEM_OPERATIONS_STATUS              byte = 0x0D
	ID_ITEM_OPERATIONS_RESPONSE            byte = 0x0E
	ID_ITEM_OPERATIONS_VERSION             byte = 0x0F
	ID_ITEM_OPERATIONS_SCHEMA              byte = 0x10
	ID_ITEM_OPERATIONS_PART                byte = 0x11
	ID_ITEM_OPERATIONS_EMPTYFOLDERCONTENTS byte = 0x12
	ID_ITEM_OPERATIONS_DELETESUBFOLDERS    byte = 0x13
	ID_ITEM_OPERATIONS_USERNAME            byte = 0x14
	ID_ITEM_OPERATIONS_PASSWORD            byte = 0x15
	ID_ITEM_OPERATIONS_MOVE                byte = 0x16
	ID_ITEM_OPERATIONS_DSTFLDID            byte = 0x17
	ID_ITEM_OPERATIONS_CONVERSATIONID      byte = 0x18
	ID_ITEM_OPERATIONS_MOVEALWAYS          byte = 0x19
)

func ItemOperations(protocolVersion byte) CodePage {
	result := NewCodePage(NS_ITEM_OPERATIONS, CP_ITEM_OPERATIONS)

	result.AddTag(TAG_ITEM_OPERATIONS_ITEMOPERATIONS, ID_ITEM_OPERATIONS_ITEMOPERATIONS)
	result.AddTag(TAG_ITEM_OPERATIONS_FETCH, ID_ITEM_OPERATIONS_FETCH)
	result.AddTag(TAG_ITEM_OPERATIONS_STORE, ID_ITEM_OPERATIONS_STORE)
	result.AddTag(TAG_ITEM_OPERATIONS_OPTIONS, ID_ITEM_OPERATIONS_OPTIONS)
	result.AddTag(TAG_ITEM_OPERATIONS_RANGE, ID_ITEM_OPERATIONS_RANGE)
	result.AddTag(TAG_ITEM_OPERATIONS_TOTAL, ID_ITEM_OPERATIONS_TOTAL)
	result.AddTag(TAG_ITEM_OPERATIONS_PROPERTIES, ID_ITEM_OPERATIONS_PROPERTIES)
	result.AddTag(TAG_ITEM_OPERATIONS_DATA, ID_ITEM_OPERATIONS_DATA)
	result.AddTag(TAG_ITEM_OPERATIONS_STATUS, ID_ITEM_OPERATIONS_STATUS)
	result.AddTag(TAG_ITEM_OPERATIONS_RESPONSE, ID_ITEM_OPERATIONS_RESPONSE)
	result.AddTag(TAG_ITEM_OPERATIONS_VERSION, ID_ITEM_OPERATIONS_VERSION)
	result.AddTag(TAG_ITEM_OPERATIONS_SCHEMA, ID_ITEM_OPERATIONS_SCHEMA)
	result.AddTag(TAG_ITEM_OPERATIONS_PART, ID_ITEM_OPERATIONS_PART)
	result.AddTag(TAG_ITEM_OPERATIONS_EMPTYFOLDERCONTENTS, ID_ITEM_OPERATIONS_EMPTYFOLDERCONTENTS)
	result.AddTag(TAG_ITEM_OPERATIONS_DELETESUBFOLDERS, ID_ITEM_OPERATIONS_DELETESUBFOLDERS)
	result.AddTag(TAG_ITEM_OPERATIONS_USERNAME, ID_ITEM_OPERATIONS_USERNAME)
	result.AddTag(TAG_ITEM_OPERATIONS_PASSWORD, ID_ITEM_OPERATIONS_PASSWORD)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_ITEM_OPERATIONS_MOVE, ID_ITEM_OPERATIONS_MOVE)                     // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_ITEM_OPERATIONS_DSTFLDID, ID_ITEM_OPERATIONS_DSTFLDID)             // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_ITEM_OPERATIONS_CONVERSATIONID, ID_ITEM_OPERATIONS_CONVERSATIONID) // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_ITEM_OPERATIONS_MOVEALWAYS, ID_ITEM_OPERATIONS_MOVEALWAYS)         // not supported with MS-ASProtocolVersion 12.1
	}

	return result
}

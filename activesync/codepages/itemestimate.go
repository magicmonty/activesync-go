package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_ITEM_ESTIMATE                  string = "ItemEstimate"
	TAG_ITEM_ESTIMATE_GETITEMESTIMATE string = "GetItemEstimate"
	TAG_ITEM_ESTIMATE_VERSION         string = "Version"
	TAG_ITEM_ESTIMATE_COLLECTIONS     string = "Collections"
	TAG_ITEM_ESTIMATE_COLLECTION      string = "Collection"
	TAG_ITEM_ESTIMATE_CLASS           string = "Class"
	TAG_ITEM_ESTIMATE_COLLECTIONID    string = "CollectionId"
	TAG_ITEM_ESTIMATE_DATETIME        string = "DateTime"
	TAG_ITEM_ESTIMATE_ESTIMATE        string = "Estimate"
	TAG_ITEM_ESTIMATE_RESPONSE        string = "Response"
	TAG_ITEM_ESTIMATE_STATUS          string = "Status"
)

const (
	CP_ITEM_ESTIMATE                 byte = 6
	ID_ITEM_ESTIMATE_GETITEMESTIMATE byte = 0x05
	ID_ITEM_ESTIMATE_VERSION         byte = 0x06
	ID_ITEM_ESTIMATE_COLLECTIONS     byte = 0x07
	ID_ITEM_ESTIMATE_COLLECTION      byte = 0x08
	ID_ITEM_ESTIMATE_CLASS           byte = 0x09
	ID_ITEM_ESTIMATE_COLLECTIONID    byte = 0x0A
	ID_ITEM_ESTIMATE_DATETIME        byte = 0x0B
	ID_ITEM_ESTIMATE_ESTIMATE        byte = 0x0C
	ID_ITEM_ESTIMATE_RESPONSE        byte = 0x0D
	ID_ITEM_ESTIMATE_STATUS          byte = 0x0E
)

func ItemEstimate(protocolVersion byte) CodePage {
	result := NewCodePage(NS_ITEM_ESTIMATE, CP_ITEM_ESTIMATE)

	result.AddTag(TAG_ITEM_ESTIMATE_GETITEMESTIMATE, ID_ITEM_ESTIMATE_GETITEMESTIMATE)
	if protocolVersion <= PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_ITEM_ESTIMATE_VERSION, ID_ITEM_ESTIMATE_VERSION) // only supported with MS-ASProtocolVersion 12.1
	}
	result.AddTag(TAG_ITEM_ESTIMATE_COLLECTIONS, ID_ITEM_ESTIMATE_COLLECTIONS)
	result.AddTag(TAG_ITEM_ESTIMATE_COLLECTION, ID_ITEM_ESTIMATE_COLLECTION)
	if protocolVersion <= PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_ITEM_ESTIMATE_CLASS, ID_ITEM_ESTIMATE_CLASS) // only supported with MS-ASProtocolVersion 12.1
		// with 14.0 and 14.1 the Class tag from CodePage 0 (ActiveSync) is used
	}
	result.AddTag(TAG_ITEM_ESTIMATE_COLLECTIONID, ID_ITEM_ESTIMATE_COLLECTIONID)
	if protocolVersion <= PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_ITEM_ESTIMATE_DATETIME, ID_ITEM_ESTIMATE_DATETIME) // only supported with MS-ASProtocolVersion 12.1
	}
	result.AddTag(TAG_ITEM_ESTIMATE_ESTIMATE, ID_ITEM_ESTIMATE_ESTIMATE)
	result.AddTag(TAG_ITEM_ESTIMATE_RESPONSE, ID_ITEM_ESTIMATE_RESPONSE)
	result.AddTag(TAG_ITEM_ESTIMATE_STATUS, ID_ITEM_ESTIMATE_STATUS)

	return result
}

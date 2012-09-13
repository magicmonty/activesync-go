package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_AIRSYNC                    string = "AirSync"
	TAG_AIRSYNC_SYNC              string = "Sync"
	TAG_AIRSYNC_RESPONSES         string = "Responses"
	TAG_AIRSYNC_ADD               string = "Add"
	TAG_AIRSYNC_CHANGE            string = "Change"
	TAG_AIRSYNC_DELETE            string = "Delete"
	TAG_AIRSYNC_FETCH             string = "Fetch"
	TAG_AIRSYNC_SYNCKEY           string = "SyncKey"
	TAG_AIRSYNC_CLIENTID          string = "ClientId"
	TAG_AIRSYNC_SERVERID          string = "ServerId"
	TAG_AIRSYNC_STATUS            string = "Status"
	TAG_AIRSYNC_COLLECTION        string = "Collection"
	TAG_AIRSYNC_CLASS             string = "Class"
	TAG_AIRSYNC_COLLECTIONID      string = "CollectionId"
	TAG_AIRSYNC_GETCHANGES        string = "GetChanges"
	TAG_AIRSYNC_MOREAVAILABLE     string = "MoreAvailable"
	TAG_AIRSYNC_WINDOWSIZE        string = "WindowSize"
	TAG_AIRSYNC_COMMANDS          string = "Commands"
	TAG_AIRSYNC_OPTIONS           string = "Options"
	TAG_AIRSYNC_FILTERTYPE        string = "FilterType"
	TAG_AIRSYNC_CONFLICT          string = "Conflict"
	TAG_AIRSYNC_COLLECTIONS       string = "Collections"
	TAG_AIRSYNC_APPLICATIONDATA   string = "ApplicationData"
	TAG_AIRSYNC_DELETESASMOVES    string = "DeletesAsMoves"
	TAG_AIRSYNC_SUPPORTED         string = "Supported"
	TAG_AIRSYNC_SOFTDELETE        string = "SoftDelete"
	TAG_AIRSYNC_MIMESUPPORT       string = "MIMESupport"
	TAG_AIRSYNC_MIMETRUNCATION    string = "MIMETruncation"
	TAG_AIRSYNC_WAIT              string = "Wait"
	TAG_AIRSYNC_LIMIT             string = "Limit"
	TAG_AIRSYNC_PARTIAL           string = "Partial"
	TAG_AIRSYNC_CONVERSATIONMODE  string = "ConversationMode"
	TAG_AIRSYNC_MAXITEMS          string = "MaxItems"
	TAG_AIRSYNC_HEARTBEATINTERVAL string = "HeartbeatInterval"
)

const (
	CP_AIRSYNC                   byte = 0
	ID_AIRSYNC_SYNC              byte = 0x05
	ID_AIRSYNC_RESPONSES         byte = 0x06
	ID_AIRSYNC_ADD               byte = 0x07
	ID_AIRSYNC_CHANGE            byte = 0x08
	ID_AIRSYNC_DELETE            byte = 0x09
	ID_AIRSYNC_FETCH             byte = 0x0A
	ID_AIRSYNC_SYNCKEY           byte = 0x0B
	ID_AIRSYNC_CLIENTID          byte = 0x0C
	ID_AIRSYNC_SERVERID          byte = 0x0D
	ID_AIRSYNC_STATUS            byte = 0x0E
	ID_AIRSYNC_COLLECTION        byte = 0x0F
	ID_AIRSYNC_CLASS             byte = 0x10
	ID_AIRSYNC_COLLECTIONID      byte = 0x12
	ID_AIRSYNC_GETCHANGES        byte = 0x13
	ID_AIRSYNC_MOREAVAILABLE     byte = 0x14
	ID_AIRSYNC_WINDOWSIZE        byte = 0x15
	ID_AIRSYNC_COMMANDS          byte = 0x16
	ID_AIRSYNC_OPTIONS           byte = 0x17
	ID_AIRSYNC_FILTERTYPE        byte = 0x18
	ID_AIRSYNC_CONFLICT          byte = 0x1B
	ID_AIRSYNC_COLLECTIONS       byte = 0x1C
	ID_AIRSYNC_APPLICATIONDATA   byte = 0x1D
	ID_AIRSYNC_DELETESASMOVES    byte = 0x1E
	ID_AIRSYNC_SUPPORTED         byte = 0x20
	ID_AIRSYNC_SOFTDELETE        byte = 0x21
	ID_AIRSYNC_MIMESUPPORT       byte = 0x22
	ID_AIRSYNC_MIMETRUNCATION    byte = 0x23
	ID_AIRSYNC_WAIT              byte = 0x24
	ID_AIRSYNC_LIMIT             byte = 0x25
	ID_AIRSYNC_PARTIAL           byte = 0x26
	ID_AIRSYNC_CONVERSATIONMODE  byte = 0x27
	ID_AIRSYNC_MAXITEMS          byte = 0x28
	ID_AIRSYNC_HEARTBEATINTERVAL byte = 0x29
)

func AirSync(protocolVersion byte) CodePage {
	result := NewCodePage(NS_AIRSYNC, CP_AIRSYNC)

	result.AddTag(TAG_AIRSYNC_SYNC, ID_AIRSYNC_SYNC)
	result.AddTag(TAG_AIRSYNC_RESPONSES, ID_AIRSYNC_RESPONSES)
	result.AddTag(TAG_AIRSYNC_ADD, ID_AIRSYNC_ADD)
	result.AddTag(TAG_AIRSYNC_CHANGE, ID_AIRSYNC_CHANGE)
	result.AddTag(TAG_AIRSYNC_DELETE, ID_AIRSYNC_DELETE)
	result.AddTag(TAG_AIRSYNC_FETCH, ID_AIRSYNC_FETCH)
	result.AddTag(TAG_AIRSYNC_SYNCKEY, ID_AIRSYNC_SYNCKEY)
	result.AddTag(TAG_AIRSYNC_CLIENTID, ID_AIRSYNC_CLIENTID)
	result.AddTag(TAG_AIRSYNC_SERVERID, ID_AIRSYNC_SERVERID)
	result.AddTag(TAG_AIRSYNC_STATUS, ID_AIRSYNC_STATUS)
	result.AddTag(TAG_AIRSYNC_COLLECTION, ID_AIRSYNC_COLLECTION)
	result.AddTag(TAG_AIRSYNC_CLASS, ID_AIRSYNC_CLASS)
	result.AddTag(TAG_AIRSYNC_COLLECTIONID, ID_AIRSYNC_COLLECTIONID)
	result.AddTag(TAG_AIRSYNC_GETCHANGES, ID_AIRSYNC_GETCHANGES)
	result.AddTag(TAG_AIRSYNC_MOREAVAILABLE, ID_AIRSYNC_MOREAVAILABLE)
	result.AddTag(TAG_AIRSYNC_WINDOWSIZE, ID_AIRSYNC_WINDOWSIZE)
	result.AddTag(TAG_AIRSYNC_COMMANDS, ID_AIRSYNC_COMMANDS)
	result.AddTag(TAG_AIRSYNC_OPTIONS, ID_AIRSYNC_OPTIONS)
	result.AddTag(TAG_AIRSYNC_FILTERTYPE, ID_AIRSYNC_FILTERTYPE)
	result.AddTag(TAG_AIRSYNC_CONFLICT, ID_AIRSYNC_CONFLICT)
	result.AddTag(TAG_AIRSYNC_COLLECTIONS, ID_AIRSYNC_COLLECTIONS)
	result.AddTag(TAG_AIRSYNC_APPLICATIONDATA, ID_AIRSYNC_APPLICATIONDATA)
	result.AddTag(TAG_AIRSYNC_DELETESASMOVES, ID_AIRSYNC_DELETESASMOVES)
	result.AddTag(TAG_AIRSYNC_SUPPORTED, ID_AIRSYNC_SUPPORTED)
	result.AddTag(TAG_AIRSYNC_SOFTDELETE, ID_AIRSYNC_SOFTDELETE)
	result.AddTag(TAG_AIRSYNC_MIMESUPPORT, ID_AIRSYNC_MIMESUPPORT)
	result.AddTag(TAG_AIRSYNC_MIMETRUNCATION, ID_AIRSYNC_MIMETRUNCATION)
	result.AddTag(TAG_AIRSYNC_WAIT, ID_AIRSYNC_WAIT)
	result.AddTag(TAG_AIRSYNC_LIMIT, ID_AIRSYNC_LIMIT)
	result.AddTag(TAG_AIRSYNC_PARTIAL, ID_AIRSYNC_PARTIAL)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_AIRSYNC_CONVERSATIONMODE, ID_AIRSYNC_CONVERSATIONMODE)   // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_AIRSYNC_MAXITEMS, ID_AIRSYNC_MAXITEMS)                   // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_AIRSYNC_HEARTBEATINTERVAL, ID_AIRSYNC_HEARTBEATINTERVAL) // not supported with MS-ASProtocolVersion 12.1
	}

	return result
}

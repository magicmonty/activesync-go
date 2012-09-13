package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_MOVE                   = "Move"
	TAG_MOVE_MOVEITEMS string = "MoveItems"
	TAG_MOVE_MOVE      string = "Move"
	TAG_MOVE_SRCMSGID  string = "SrcMsgId"
	TAG_MOVE_SRCFLDID  string = "SrcFldId"
	TAG_MOVE_DSTFLDID  string = "DstFldId"
	TAG_MOVE_RESPONSE  string = "Response"
	TAG_MOVE_STATUS    string = "Status"
	TAG_MOVE_DSTMSGID  string = "DstMsgId"
)

const (
	CP_MOVE           byte = 5
	ID_MOVE_MOVEITEMS byte = 0x05
	ID_MOVE_MOVE      byte = 0x06
	ID_MOVE_SRCMSGID  byte = 0x07
	ID_MOVE_SRCFLDID  byte = 0x08
	ID_MOVE_DSTFLDID  byte = 0x09
	ID_MOVE_RESPONSE  byte = 0x0A
	ID_MOVE_STATUS    byte = 0x0B
	ID_MOVE_DSTMSGID  byte = 0x0C
)

func Move() CodePage {
	result := NewCodePage(NS_MOVE, CP_MOVE)

	result.AddTag(TAG_MOVE_MOVEITEMS, ID_MOVE_MOVEITEMS)
	result.AddTag(TAG_MOVE_MOVE, ID_MOVE_MOVE)
	result.AddTag(TAG_MOVE_SRCMSGID, ID_MOVE_SRCMSGID)
	result.AddTag(TAG_MOVE_SRCFLDID, ID_MOVE_SRCFLDID)
	result.AddTag(TAG_MOVE_DSTFLDID, ID_MOVE_DSTFLDID)
	result.AddTag(TAG_MOVE_RESPONSE, ID_MOVE_RESPONSE)
	result.AddTag(TAG_MOVE_STATUS, ID_MOVE_STATUS)
	result.AddTag(TAG_MOVE_DSTMSGID, ID_MOVE_DSTMSGID)

	return result
}

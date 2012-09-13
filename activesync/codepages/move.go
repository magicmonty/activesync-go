package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func Move() CodePage {
	result := NewCodePage("Move", 5)

	result.AddTag("MoveItems", 0x05)
	result.AddTag("Move", 0x06)
	result.AddTag("SrcMsgId", 0x07)
	result.AddTag("SrcFldId", 0x08)
	result.AddTag("DstFldId", 0x09)
	result.AddTag("Response", 0x0A)
	result.AddTag("Status", 0x0B)
	result.AddTag("DstMsgId", 0x0C)

	return result
}

package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func FolderHierarchy() CodePage {
	result := NewCodePage("FolderHierarchy", 7)

	result.AddTag("DisplayName", 0x07)
	result.AddTag("ServerId", 0x08)
	result.AddTag("ParentId", 0x09)
	result.AddTag("Type", 0x0A)
	result.AddTag("Status", 0x0C)
	result.AddTag("Changes", 0x0E)
	result.AddTag("Add", 0x0F)
	result.AddTag("Delete", 0x10)
	result.AddTag("Update", 0x11)
	result.AddTag("SyncKey", 0x12)
	result.AddTag("FolderCreate", 0x13)
	result.AddTag("FolderDelete", 0x14)
	result.AddTag("FolderUpdate", 0x15)
	result.AddTag("FolderSync", 0x16)
	result.AddTag("Count", 0x17)

	return result
}

package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func DocumentLibrary() CodePage {
	result := NewCodePage("DocumentLibrary", 19)

	result.AddTag("LinkId", 0x05)
	result.AddTag("DisplayName", 0x06)
	result.AddTag("IsFolder", 0x07)
	result.AddTag("CreationDate", 0x08)
	result.AddTag("LastModifiedDate", 0x09)
	result.AddTag("IsHidden", 0x0A)
	result.AddTag("ContentLength", 0x0B)
	result.AddTag("ContentType", 0x0C)

	return result
}

package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

// not supported with MS-ASProtocolVersion 12.1
func Notes() CodePage {
	result := NewCodePage("Notes", 23)

	result.AddTag("Subject", 0x05)
	result.AddTag("MessageClass", 0x06)
	result.AddTag("LastModifiedDate", 0x07)
	result.AddTag("Categories", 0x08)
	result.AddTag("Category", 0x09)

	return result
}

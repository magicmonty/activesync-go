package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func Ping() CodePage {
	result := NewCodePage("Ping", 13)

	result.AddTag("Ping", 0x05)
	result.AddTag("AutdState", 0x06) // not used
	result.AddTag("Status", 0x07)
	result.AddTag("HeartbeatInterval", 0x08)
	result.AddTag("Folders", 0x09)
	result.AddTag("Folder", 0x0A)
	result.AddTag("Id", 0x0B)
	result.AddTag("Class", 0x0C)
	result.AddTag("MaxFolders", 0x0D)

	return result
}

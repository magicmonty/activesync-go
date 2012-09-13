package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_PING                    string = "Ping"
	TAG_PING_PING              string = "Ping"
	TAG_PING_AUTDSTATE         string = "AutdState"
	TAG_PING_STATUS            string = "Status"
	TAG_PING_HEARTBEATINTERVAL string = "HeartbeatInterval"
	TAG_PING_FOLDERS           string = "Folders"
	TAG_PING_FOLDER            string = "Folder"
	TAG_PING_ID                string = "Id"
	TAG_PING_CLASS             string = "Class"
	TAG_PING_MAXFOLDERS        string = "MaxFolders"
)

const (
	CP_PING                   byte = 13
	ID_PING_PING              byte = 0x05
	ID_PING_AUTDSTATE         byte = 0x06
	ID_PING_STATUS            byte = 0x07
	ID_PING_HEARTBEATINTERVAL byte = 0x08
	ID_PING_FOLDERS           byte = 0x09
	ID_PING_FOLDER            byte = 0x0A
	ID_PING_ID                byte = 0x0B
	ID_PING_CLASS             byte = 0x0C
	ID_PING_MAXFOLDERS        byte = 0x0D
)

func Ping() CodePage {
	result := NewCodePage(NS_PING, CP_PING)

	result.AddTag(TAG_PING_PING, ID_PING_PING)
	result.AddTag(TAG_PING_AUTDSTATE, ID_PING_AUTDSTATE) // not used
	result.AddTag(TAG_PING_STATUS, ID_PING_STATUS)
	result.AddTag(TAG_PING_HEARTBEATINTERVAL, ID_PING_HEARTBEATINTERVAL)
	result.AddTag(TAG_PING_FOLDERS, ID_PING_FOLDERS)
	result.AddTag(TAG_PING_FOLDER, ID_PING_FOLDER)
	result.AddTag(TAG_PING_ID, ID_PING_ID)
	result.AddTag(TAG_PING_CLASS, ID_PING_CLASS)
	result.AddTag(TAG_PING_MAXFOLDERS, ID_PING_MAXFOLDERS)

	return result
}

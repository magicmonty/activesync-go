package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_GAL               string = "GAL"
	TAG_GAL_DISPLAYNAME  string = "DisplayName"
	TAG_GAL_PHONE        string = "Phone"
	TAG_GAL_OFFICE       string = "Office"
	TAG_GAL_TITLE        string = "Title"
	TAG_GAL_COMPANY      string = "Company"
	TAG_GAL_ALIAS        string = "Alias"
	TAG_GAL_FIRSTNAME    string = "FirstName"
	TAG_GAL_LASTNAME     string = "LastName"
	TAG_GAL_HOMEPHONE    string = "HomePhone"
	TAG_GAL_MOBILEPHONE  string = "MobilePhone"
	TAG_GAL_EMAILADDRESS string = "EmailAddress"
	TAG_GAL_PICTURE      string = "Picture"
	TAG_GAL_STATUS       string = "Status"
	TAG_GAL_DATA         string = "Data"
)

const (
	CP_GAL              byte = 16
	ID_GAL_DISPLAYNAME  byte = 0x05
	ID_GAL_PHONE        byte = 0x06
	ID_GAL_OFFICE       byte = 0x07
	ID_GAL_TITLE        byte = 0x08
	ID_GAL_COMPANY      byte = 0x09
	ID_GAL_ALIAS        byte = 0x0A
	ID_GAL_FIRSTNAME    byte = 0x0B
	ID_GAL_LASTNAME     byte = 0x0C
	ID_GAL_HOMEPHONE    byte = 0x0D
	ID_GAL_MOBILEPHONE  byte = 0x0E
	ID_GAL_EMAILADDRESS byte = 0x0F
	ID_GAL_PICTURE      byte = 0x10
	ID_GAL_STATUS       byte = 0x11
	ID_GAL_DATA         byte = 0x12
)

func GAL(protocolVersion byte) CodePage {
	result := NewCodePage(NS_GAL, CP_GAL)

	result.AddTag(TAG_GAL_DISPLAYNAME, ID_GAL_DISPLAYNAME)
	result.AddTag(TAG_GAL_PHONE, ID_GAL_PHONE)
	result.AddTag(TAG_GAL_OFFICE, ID_GAL_OFFICE)
	result.AddTag(TAG_GAL_TITLE, ID_GAL_TITLE)
	result.AddTag(TAG_GAL_COMPANY, ID_GAL_COMPANY)
	result.AddTag(TAG_GAL_ALIAS, ID_GAL_ALIAS)
	result.AddTag(TAG_GAL_FIRSTNAME, ID_GAL_FIRSTNAME)
	result.AddTag(TAG_GAL_LASTNAME, ID_GAL_LASTNAME)
	result.AddTag(TAG_GAL_HOMEPHONE, ID_GAL_HOMEPHONE)
	result.AddTag(TAG_GAL_MOBILEPHONE, ID_GAL_MOBILEPHONE)
	result.AddTag(TAG_GAL_EMAILADDRESS, ID_GAL_EMAILADDRESS)
	if protocolVersion > PROTOCOL_VERSION_14_0 {
		result.AddTag(TAG_GAL_PICTURE, ID_GAL_PICTURE) // not supported with MS-ASProtocolVersion 12.1 or 14.0
		result.AddTag(TAG_GAL_STATUS, ID_GAL_STATUS)   // not supported with MS-ASProtocolVersion 12.1 or 14.0
		result.AddTag(TAG_GAL_DATA, ID_GAL_DATA)       // not supported with MS-ASProtocolVersion 12.1 or 14.0
	}

	return result
}

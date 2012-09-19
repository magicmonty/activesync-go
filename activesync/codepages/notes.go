package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_NOTES                   string = "Notes"
	TAG_NOTES_SUBJECT          string = "Subject"
	TAG_NOTES_MESSAGECLASS     string = "MessageClass"
	TAG_NOTES_LASTMODIFIEDDATE string = "LastModifiedDate"
	TAG_NOTES_CATEGORIES       string = "Categories"
	TAG_NOTES_CATEGORY         string = "Category"
)

const (
	CP_NOTES                  byte = 23
	ID_NOTES_SUBJECT          byte = 0x05
	ID_NOTES_MESSAGECLASS     byte = 0x06
	ID_NOTES_LASTMODIFIEDDATE byte = 0x07
	ID_NOTES_CATEGORIES       byte = 0x08
	ID_NOTES_CATEGORY         byte = 0x09
)

// not supported with MS-ASProtocolVersion 12.1
func Notes(protocolVersion byte) CodePage {
	result := NewCodePage(NS_NOTES, CP_NOTES)

	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_NOTES_SUBJECT, ID_NOTES_SUBJECT)
		result.AddTag(TAG_NOTES_MESSAGECLASS, ID_NOTES_MESSAGECLASS)
		result.AddTag(TAG_NOTES_LASTMODIFIEDDATE, ID_NOTES_LASTMODIFIEDDATE)
		result.AddTag(TAG_NOTES_CATEGORIES, ID_NOTES_CATEGORIES)
		result.AddTag(TAG_NOTES_CATEGORY, ID_NOTES_CATEGORY)
	}

	return result
}

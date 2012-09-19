package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
	"testing"
)

func checkNotesCommonTags(t *testing.T, cp CodePage) {
	checkTagExists(t, TAG_NOTES_SUBJECT, ID_NOTES_SUBJECT, cp)
	checkTagExists(t, TAG_NOTES_MESSAGECLASS, ID_NOTES_MESSAGECLASS, cp)
	checkTagExists(t, TAG_NOTES_LASTMODIFIEDDATE, ID_NOTES_LASTMODIFIEDDATE, cp)
	checkTagExists(t, TAG_NOTES_CATEGORIES, ID_NOTES_CATEGORIES, cp)
	checkTagExists(t, TAG_NOTES_CATEGORY, ID_NOTES_CATEGORY, cp)
}

func Test_Notes_120(t *testing.T) {
	cp := Notes(PROTOCOL_VERSION_12_0)
	checkTagCount(t, 0, cp)
}

func Test_Notes_121(t *testing.T) {
	cp := Notes(PROTOCOL_VERSION_12_1)
	checkTagCount(t, 0, cp)
}

func Test_Notes_140(t *testing.T) {
	cp := Notes(PROTOCOL_VERSION_14_0)
	checkNotesCommonTags(t, cp)
	checkTagCount(t, 5, cp)
}

func Test_Notes_141(t *testing.T) {
	cp := Notes(PROTOCOL_VERSION_14_1)
	checkNotesCommonTags(t, cp)
	checkTagCount(t, 5, cp)
}

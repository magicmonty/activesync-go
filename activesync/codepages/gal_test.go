package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
	"testing"
)

func checkGALCommonTags(t *testing.T, cp CodePage) {
	checkTagExists(t, TAG_GAL_DISPLAYNAME, ID_GAL_DISPLAYNAME, cp)
	checkTagExists(t, TAG_GAL_PHONE, ID_GAL_PHONE, cp)
	checkTagExists(t, TAG_GAL_OFFICE, ID_GAL_OFFICE, cp)
	checkTagExists(t, TAG_GAL_TITLE, ID_GAL_TITLE, cp)
	checkTagExists(t, TAG_GAL_COMPANY, ID_GAL_COMPANY, cp)
	checkTagExists(t, TAG_GAL_ALIAS, ID_GAL_ALIAS, cp)
	checkTagExists(t, TAG_GAL_FIRSTNAME, ID_GAL_FIRSTNAME, cp)
	checkTagExists(t, TAG_GAL_LASTNAME, ID_GAL_LASTNAME, cp)
	checkTagExists(t, TAG_GAL_HOMEPHONE, ID_GAL_HOMEPHONE, cp)
	checkTagExists(t, TAG_GAL_MOBILEPHONE, ID_GAL_MOBILEPHONE, cp)
	checkTagExists(t, TAG_GAL_EMAILADDRESS, ID_GAL_EMAILADDRESS, cp)
}

func Test_GAL_120(t *testing.T) {
	cp := GAL(PROTOCOL_VERSION_12_0)
	checkGALCommonTags(t, cp)
	checkTagCount(t, 11, cp)
	checkTagNotExists(t, TAG_GAL_PICTURE, ID_GAL_PICTURE, cp)
	checkTagNotExists(t, TAG_GAL_STATUS, ID_GAL_STATUS, cp)
	checkTagNotExists(t, TAG_GAL_DATA, ID_GAL_DATA, cp)
}

func Test_GAL_121(t *testing.T) {
	cp := GAL(PROTOCOL_VERSION_12_1)
	checkGALCommonTags(t, cp)
	checkTagCount(t, 11, cp)
	checkTagNotExists(t, TAG_GAL_PICTURE, ID_GAL_PICTURE, cp)
	checkTagNotExists(t, TAG_GAL_STATUS, ID_GAL_STATUS, cp)
	checkTagNotExists(t, TAG_GAL_DATA, ID_GAL_DATA, cp)
}

func Test_GAL_140(t *testing.T) {
	cp := GAL(PROTOCOL_VERSION_14_0)
	checkGALCommonTags(t, cp)
	checkTagCount(t, 11, cp)
	checkTagNotExists(t, TAG_GAL_PICTURE, ID_GAL_PICTURE, cp)
	checkTagNotExists(t, TAG_GAL_STATUS, ID_GAL_STATUS, cp)
	checkTagNotExists(t, TAG_GAL_DATA, ID_GAL_DATA, cp)
}

func Test_GAL_141(t *testing.T) {
	cp := GAL(PROTOCOL_VERSION_14_1)
	checkGALCommonTags(t, cp)
	checkTagCount(t, 14, cp)
	checkTagExists(t, TAG_GAL_PICTURE, ID_GAL_PICTURE, cp)
	checkTagExists(t, TAG_GAL_STATUS, ID_GAL_STATUS, cp)
	checkTagExists(t, TAG_GAL_DATA, ID_GAL_DATA, cp)
}

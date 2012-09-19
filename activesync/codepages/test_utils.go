package codepages

import (
	"github.com/magicmonty/wbxml-go/wbxml"
	"testing"
)

func checkTagCount(t *testing.T, expectedCount int, cp wbxml.CodePage) {
	if len(cp.TagCodes) != expectedCount {
		t.Errorf("Length of tag codes should be %d but was %d", expectedCount, len(cp.TagCodes))
	}
}

func checkTagExists(t *testing.T, tagName string, tagCode byte, cp wbxml.CodePage) {
	if !cp.HasTag(tagName) {
		t.Errorf("codepage %s should have tag %s", cp.Name, tagName)
	}

	if !cp.HasTagCode(tagCode) {
		t.Errorf("codepage %s should have tag code 0x%0.2X", cp.Name, tagCode)
	}
}

func checkTagNotExists(t *testing.T, tagName string, tagCode byte, cp wbxml.CodePage) {
	if cp.HasTag(tagName) {
		t.Errorf("codepage %s should NOT have tag %s", cp.Name, tagName)
	}

	if cp.HasTagCode(tagCode) {
		t.Errorf("codepage %s should have NOT tag code 0x%0.2X", cp.Name, tagCode)
	}
}

package activesync

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	"github.com/magicmonty/wbxml-go/wbxml"
	"testing"
)

func Test_MakeCodeBook_12_0(t *testing.T) {
	checkCodeBook(
		t,
		PROTOCOL_VERSION_12_0,
		21,
		getExpectedCodePageCounts_12_0())
}

func getExpectedCodePageCounts_12_0() map[byte]int {
	var expectedEntryCount map[byte]int = make(map[byte]int)

	expectedEntryCount[0] = 30
	expectedEntryCount[1] = 52
	expectedEntryCount[2] = 42
	expectedEntryCount[4] = 35
	expectedEntryCount[5] = 8
	expectedEntryCount[6] = 10
	expectedEntryCount[7] = 15
	expectedEntryCount[8] = 9
	expectedEntryCount[9] = 27
	expectedEntryCount[10] = 17
	expectedEntryCount[11] = 6
	expectedEntryCount[12] = 10
	expectedEntryCount[13] = 9
	expectedEntryCount[14] = 53
	expectedEntryCount[15] = 23
	expectedEntryCount[16] = 11
	expectedEntryCount[17] = 18
	expectedEntryCount[18] = 28
	expectedEntryCount[19] = 8
	expectedEntryCount[20] = 17

	return expectedEntryCount
}

func Test_MakeCodeBook_12_1(t *testing.T) {
	checkCodeBook(
		t,
		PROTOCOL_VERSION_12_1,
		21,
		getExpectedCodePageCounts_12_1())
}

func getExpectedCodePageCounts_12_1() map[byte]int {
	var expectedEntryCount map[byte]int = make(map[byte]int)

	expectedEntryCount[0] = 30
	expectedEntryCount[1] = 52
	expectedEntryCount[2] = 42
	expectedEntryCount[4] = 35
	expectedEntryCount[5] = 8
	expectedEntryCount[6] = 10
	expectedEntryCount[7] = 15
	expectedEntryCount[8] = 9
	expectedEntryCount[9] = 27
	expectedEntryCount[10] = 17
	expectedEntryCount[11] = 6
	expectedEntryCount[12] = 10
	expectedEntryCount[13] = 9
	expectedEntryCount[14] = 53
	expectedEntryCount[15] = 23
	expectedEntryCount[16] = 11
	expectedEntryCount[17] = 18
	expectedEntryCount[18] = 28
	expectedEntryCount[19] = 8
	expectedEntryCount[20] = 17

	return expectedEntryCount
}

func Test_MakeCodeBook_14_0(t *testing.T) {
	checkCodeBook(
		t,
		PROTOCOL_VERSION_14_0,
		24,
		getExpectedCodePageCounts_14_0())
}

func getExpectedCodePageCounts_14_0() map[byte]int {
	var expectedEntryCount map[byte]int = make(map[byte]int)

	expectedEntryCount[0] = 33
	expectedEntryCount[1] = 54
	expectedEntryCount[2] = 45
	expectedEntryCount[4] = 41
	expectedEntryCount[5] = 8
	expectedEntryCount[6] = 7
	expectedEntryCount[7] = 15
	expectedEntryCount[8] = 9
	expectedEntryCount[9] = 29
	expectedEntryCount[10] = 21
	expectedEntryCount[11] = 6
	expectedEntryCount[12] = 10
	expectedEntryCount[13] = 9
	expectedEntryCount[14] = 53
	expectedEntryCount[15] = 24
	expectedEntryCount[16] = 11
	expectedEntryCount[17] = 19
	expectedEntryCount[18] = 30
	expectedEntryCount[19] = 8
	expectedEntryCount[20] = 21
	expectedEntryCount[21] = 13
	expectedEntryCount[22] = 12
	expectedEntryCount[23] = 5

	return expectedEntryCount
}

func Test_MakeCodeBook_14_1(t *testing.T) {
	checkCodeBook(
		t,
		PROTOCOL_VERSION_14_1,
		25,
		getExpectedCodePageCounts_14_1())
}

func getExpectedCodePageCounts_14_1() map[byte]int {
	var expectedEntryCount map[byte]int = make(map[byte]int)

	expectedEntryCount[0] = 33
	expectedEntryCount[1] = 54
	expectedEntryCount[2] = 45
	expectedEntryCount[4] = 44
	expectedEntryCount[5] = 8
	expectedEntryCount[6] = 7
	expectedEntryCount[7] = 15
	expectedEntryCount[8] = 9
	expectedEntryCount[9] = 30
	expectedEntryCount[10] = 25
	expectedEntryCount[11] = 6
	expectedEntryCount[12] = 10
	expectedEntryCount[13] = 9
	expectedEntryCount[14] = 53
	expectedEntryCount[15] = 27
	expectedEntryCount[16] = 14
	expectedEntryCount[17] = 22
	expectedEntryCount[18] = 38
	expectedEntryCount[19] = 8
	expectedEntryCount[20] = 21
	expectedEntryCount[21] = 14
	expectedEntryCount[22] = 15
	expectedEntryCount[23] = 5
	expectedEntryCount[24] = 20

	return expectedEntryCount
}

func checkCodeBook(t *testing.T, protocolVersion byte, maxCodePage byte, expectedCodePageCounts map[byte]int) {
	var (
		cb *wbxml.CodeBook = MakeCodeBook(protocolVersion)
		i  byte
	)

	if !cb.IsReady() {
		t.Error("CodeBook should be ready")
	}

	for i = 0; i < 25; i++ {
		if i != 3 && i < maxCodePage {
			checkCodePage(t, cb, i, expectedCodePageCounts)
		} else {
			checkHasNotCodePage(t, cb, i)
		}
	}
}

func checkCodePage(t *testing.T, cb *wbxml.CodeBook, code byte, expectedCount map[byte]int) {
	if !cb.HasTagCode(code) {
		t.Errorf("codebook should have tag code page %d!", code)
	}
	checkCodePageCount(t, cb.TagCodePages[code], expectedCount)
}

func checkCodePageCount(t *testing.T, cp wbxml.CodePage, expectedCount map[byte]int) {
	cpSize := len(cp.TagCodes)
	expectedSize := expectedCount[cp.Code]
	if cpSize != expectedSize {
		t.Errorf("codepage %d should have %d entries but has %d!", cp.Code, expectedSize, cpSize)
	}
}

func checkHasNotCodePage(t *testing.T, cb *wbxml.CodeBook, code byte) {
	if cb.HasTagCode(code) {
		t.Errorf("codebook should NOT have tag code page %d!", code)
	}
}

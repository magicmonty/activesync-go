package codepages

import (
	"testing"
)

func Test_Move(t *testing.T) {
	cp := Move()
	checkTagCount(t, 8, cp)
	checkTagExists(t, TAG_MOVE_MOVEITEMS, ID_MOVE_MOVEITEMS, cp)
	checkTagExists(t, TAG_MOVE_MOVE, ID_MOVE_MOVE, cp)
	checkTagExists(t, TAG_MOVE_SRCMSGID, ID_MOVE_SRCMSGID, cp)
	checkTagExists(t, TAG_MOVE_SRCFLDID, ID_MOVE_SRCFLDID, cp)
	checkTagExists(t, TAG_MOVE_DSTFLDID, ID_MOVE_DSTFLDID, cp)
	checkTagExists(t, TAG_MOVE_RESPONSE, ID_MOVE_RESPONSE, cp)
	checkTagExists(t, TAG_MOVE_STATUS, ID_MOVE_STATUS, cp)
	checkTagExists(t, TAG_MOVE_DSTMSGID, ID_MOVE_DSTMSGID, cp)
}

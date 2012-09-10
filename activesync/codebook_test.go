package activesync

import (
	"github.com/magicmonty/wbxml-go/wbxml"
	"testing"
)

func Test_MakeCodeBook(t *testing.T) {
	var (
		cb *wbxml.CodeBook = MakeCodeBook()
		i  byte
	)
	if !cb.IsReady() {
		t.Error("CodeBook should be ready")
	}

	for i = 0; i < 25; i++ {
		if !cb.HasTagCode(i) {
			t.Errorf("codebook should have tag code page %d!", i)
		}
	}
}

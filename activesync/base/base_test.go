package base

import (
	"testing"
)

func Test_CommandCodes(t *testing.T) {
	if CCSync != 0 {
		t.Errorf("CCSync should be 0 but was %d", CCSync)
	}
	if CCSendMail != 1 {
		t.Errorf("CCSendMail should be 1 but was %d", CCSendMail)
	}
	if CCGetAttachment != 4 {
		t.Errorf("CCGetAttachment should be 4 but was %d", CCGetAttachment)
	}
	if CCFolderSync != 9 {
		t.Errorf("CCFolderSync should be 9 but was %d", CCFolderSync)
	}
	if CCFolderCreate != 10 {
		t.Errorf("CCFolderCreate should be 10 but was %d", CCFolderCreate)
	}
	if CCValidateCert != 22 {
		t.Errorf("CCValidateCert should be 22 but was %d", CCValidateCert)
	}
}

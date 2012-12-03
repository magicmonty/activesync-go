package activesync

import (
	"bytes"
	. "github.com/magicmonty/activesync-go/activesync/base"
	"net/http"
	"testing"
)

func Test_getRequestHeader(t *testing.T) {
	body := bytes.NewReader([]byte{})

	// URL: 8C 00 09 04 0A 76 31 34 30 44 65 76 69 63 65 00 0A 53 6D 61 72 74 50 68 6F 6E 65
	r, _ := http.NewRequest("POST", "/Microsoft-Server-ActiveSync?jAAJBAp2MTQwRGV2aWNlAApTbWFydFBob25l", body)

	rh, err := getRequestHeader(r)

	if err == nil {
		if rh.ProtocolVersion != PROTOCOL_VERSION_14_0 {
			t.Errorf("Protocol Version should be %d but was %d", PROTOCOL_VERSION_14_0, rh.ProtocolVersion)
		}

		if rh.Command != CCSync {
			t.Errorf("Command should be %d but was %d", CCSync, rh.Command)
		}

		if rh.Locale != "en-US" {
			t.Errorf("Locale should be 'en-US' but was %s", rh.Locale)
		}

		if rh.DeviceId != "v140Device" {
			t.Errorf("DeviceId should be \"v140Device\" but was \"%s\"", rh.DeviceId)
		}

		if rh.PolicyKey != "0" {
			t.Errorf("PolicyKey should be \"0\" but was \"%s\"", rh.PolicyKey)
		}

		if rh.DeviceType != "SmartPhone" {
			t.Errorf("DeviceType should be \"SmartPhone\" but was \"%s\"", rh.DeviceType)
		}
	} else {
		t.Errorf("Got error: %s", err.Error())
	}
}

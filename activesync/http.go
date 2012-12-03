package activesync

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/magicmonty/activesync-go/activesync/base"
	"net/http"
	"strings"
)

type RequestHeader struct {
	ProtocolVersion   byte
	Command           base.CommandCode
	Locale            string
	DeviceId          string
	PolicyKey         string
	DeviceType        string
	CommandParameters map[string]string
}

func NewRequestHeader() *RequestHeader {
	r := new(RequestHeader)
	r.ProtocolVersion = 0
	r.Command = base.CCUnknown
	r.Locale = base.GetLanguageByCode(uint16(0x0409)) // en-us
	r.DeviceId = ""
	r.PolicyKey = "0"
	r.DeviceType = ""
	r.CommandParameters = make(map[string]string)
	return r
}

func getRequestHeader(r *http.Request) (*RequestHeader, error) {
	var (
		result *RequestHeader
		err    error
	)

	if strings.Contains(r.URL.RequestURI(), "Cmd=") {
		result = NewRequestHeader()
	} else {
		result, err = decodeBase64Url(r.URL.RequestURI())
	}

	return result, err
}

func decodeBase64Url(url string) (*RequestHeader, error) {
	var (
		result *RequestHeader = NewRequestHeader()
		err    error
	)

	parts := strings.Split(url, "?")
	if len(parts) == 2 {
		decoded, err := base64.StdEncoding.DecodeString(parts[1])
		if err == nil {
			br := bytes.NewReader(decoded)

			result.ProtocolVersion, err = br.ReadByte()
			if err != nil {
				return nil, err
			}

			b, err := br.ReadByte()
			if err != nil {
				return nil, err
			}
			result.Command = base.CommandCode(b)

			b, err = br.ReadByte()
			if err != nil {
				return nil, err
			}
			b2, err := br.ReadByte()
			if err != nil {
				return nil, err
			}
			result.Locale = base.GetLanguageByCode(uint16(b2)<<8 + uint16(b))

			b, err = br.ReadByte()
			if err != nil {
				return nil, err
			}
			var l byte
			for l = 0; l < b; l++ {
				b2, err = br.ReadByte()
				if err != nil {
					return nil, err
				}
				result.DeviceId += fmt.Sprintf("%c", b2)
			}
			b, err = br.ReadByte()
			if err != nil {
				return nil, err
			}
			if b > 0 {
				result.PolicyKey = ""
				for l = 0; l < b; l++ {
					b2, err = br.ReadByte()
					if err != nil {
						return nil, err
					}
					result.PolicyKey += fmt.Sprintf("%c", b2)
				}
			}
			b, err = br.ReadByte()
			if err != nil {
				return nil, err
			}
			for l = 0; l < b; l++ {
				b2, err = br.ReadByte()
				if err != nil {
					return nil, err
				}
				result.DeviceType += fmt.Sprintf("%c", b2)
			}

			
		}
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

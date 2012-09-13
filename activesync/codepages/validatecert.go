package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func ValidateCert() CodePage {
	result := NewCodePage("ValidateCert", 11)

	result.AddTag("ValidateCert", 0x05)
	result.AddTag("Certificates", 0x06)
	result.AddTag("Certificate", 0x07)
	result.AddTag("CertificateChain", 0x08)
	result.AddTag("CheckCRL", 0x09)
	result.AddTag("Status", 0x0A)

	return result
}

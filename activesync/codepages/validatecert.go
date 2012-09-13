package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_VALIDATECERT                    string = "ValidateCert"
	TAG_VALIDATE_CERT_VALIDATECERT     string = "ValidateCert"
	TAG_VALIDATE_CERT_CERTIFICATES     string = "Certificates"
	TAG_VALIDATE_CERT_CERTIFICATE      string = "Certificate"
	TAG_VALIDATE_CERT_CERTIFICATECHAIN string = "CertificateChain"
	TAG_VALIDATE_CERT_CHECKCRL         string = "CheckCRL"
	TAG_VALIDATE_CERT_STATUS           string = "Status"
)

const (
	CP_VALIDATECERT                   byte = 11
	ID_VALIDATE_CERT_VALIDATECERT     byte = 0x05
	ID_VALIDATE_CERT_CERTIFICATES     byte = 0x06
	ID_VALIDATE_CERT_CERTIFICATE      byte = 0x07
	ID_VALIDATE_CERT_CERTIFICATECHAIN byte = 0x08
	ID_VALIDATE_CERT_CHECKCRL         byte = 0x09
	ID_VALIDATE_CERT_STATUS           byte = 0x0A
)

func ValidateCert() CodePage {
	result := NewCodePage(NS_VALIDATECERT, CP_VALIDATECERT)

	result.AddTag(TAG_VALIDATE_CERT_VALIDATECERT, ID_VALIDATE_CERT_VALIDATECERT)
	result.AddTag(TAG_VALIDATE_CERT_CERTIFICATES, ID_VALIDATE_CERT_CERTIFICATES)
	result.AddTag(TAG_VALIDATE_CERT_CERTIFICATE, ID_VALIDATE_CERT_CERTIFICATE)
	result.AddTag(TAG_VALIDATE_CERT_CERTIFICATECHAIN, ID_VALIDATE_CERT_CERTIFICATECHAIN)
	result.AddTag(TAG_VALIDATE_CERT_CHECKCRL, ID_VALIDATE_CERT_CHECKCRL)
	result.AddTag(TAG_VALIDATE_CERT_STATUS, ID_VALIDATE_CERT_STATUS)

	return result
}

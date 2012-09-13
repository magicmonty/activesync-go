package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func ResolveRecipients(protocolVersion byte) CodePage {
	result := NewCodePage("ResolveRecipients", 10)

	result.AddTag("ResolveRecipients", 0x05)
	result.AddTag("Response", 0x06)
	result.AddTag("Status", 0x07)
	result.AddTag("Type", 0x08)
	result.AddTag("Recipient", 0x09)
	result.AddTag("DisplayName", 0x0A)
	result.AddTag("EmailAddress", 0x0B)
	result.AddTag("Certificates", 0x0C)
	result.AddTag("Certificate", 0x0D)
	result.AddTag("MiniCertificate", 0x0E)
	result.AddTag("Options", 0x0F)
	result.AddTag("To", 0x10)
	result.AddTag("CertificateRetrieval", 0x11)
	result.AddTag("RecipientCount", 0x12)
	result.AddTag("MaxCertificates", 0x13)
	result.AddTag("MaxAmbiguousRecipients", 0x14)
	result.AddTag("CertificateCount", 0x15)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag("Availability", 0x16)   // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("StartTime", 0x17)      // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("EndTime", 0x18)        // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("MergedFreeBusy", 0x19) // not supported with MS-ASProtocolVersion 12.1
		if protocolVersion > PROTOCOL_VERSION_14_0 {
			result.AddTag("Picture", 0x1A)     // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("MaxSize", 0x1B)     // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("Data", 0x1C)        // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("MaxPictures", 0x1D) // not supported with MS-ASProtocolVersion 12.1 or 14.0
		}
	}

	return result
}

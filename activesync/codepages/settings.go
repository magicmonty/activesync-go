package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func Settings(protocolVersion byte) CodePage {
	result := NewCodePage("Settings", 18)

	result.AddTag("Settings", 0x05)
	result.AddTag("Status", 0x06)
	result.AddTag("Get", 0x07)
	result.AddTag("Set", 0x08)
	result.AddTag("Oof", 0x09)
	result.AddTag("OofState", 0x0A)
	result.AddTag("StartTime", 0x0B)
	result.AddTag("EndTime", 0x0C)
	result.AddTag("OofMessage", 0x0D)
	result.AddTag("AppliesToInternal", 0x0E)
	result.AddTag("AppliesToExternalKnown", 0x0F)
	result.AddTag("AppliesToExternalUnknown", 0x10)
	result.AddTag("Enabled", 0x11)
	result.AddTag("ReplyMessage", 0x12)
	result.AddTag("BodyType", 0x13)
	result.AddTag("DevicePassword", 0x14)
	result.AddTag("Password", 0x15)
	result.AddTag("DeviceInformation", 0x16)
	result.AddTag("Model", 0x17)
	result.AddTag("IMEI", 0x18)
	result.AddTag("FriendlyName", 0x19)
	result.AddTag("OS", 0x1A)
	result.AddTag("OSLanguage", 0x1B)
	result.AddTag("PhoneNumber", 0x1C)
	result.AddTag("UserInformation", 0x1D)
	result.AddTag("EmailAddresses", 0x1E)
	result.AddTag("SmtpAddress", 0x1F)
	result.AddTag("UserAgent", 0x20)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag("EnableOutboundSMS", 0x21) // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("MobileOperator", 0x22)    // not supported with MS-ASProtocolVersion 12.1
		if protocolVersion > PROTOCOL_VERSION_14_0 {
			result.AddTag("PrimarySmtpAddress", 0x23)          // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("Accounts", 0x24)                    // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("Account", 0x25)                     // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("AccountId", 0x26)                   // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("AccountName", 0x27)                 // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("UserDisplayName", 0x28)             // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("SendDisabled", 0x29)                // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag("RightsManagementInformation", 0x2B) // not supported with MS-ASProtocolVersion 12.1 or 14.0
		}
	}

	return result
}

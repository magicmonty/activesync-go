package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func Contacts2() CodePage {
	result := NewCodePage("Contacts2", 12)

	result.AddTag("CustomerId", 0x05)
	result.AddTag("GovernmentId", 0x06)
	result.AddTag("IMAddress", 0x07)
	result.AddTag("IMAddress2", 0x08)
	result.AddTag("IMAddress3", 0x09)
	result.AddTag("ManagerName", 0x0A)
	result.AddTag("CompanyMainPhone", 0x0B)
	result.AddTag("AccountName", 0x0C)
	result.AddTag("NickName", 0x0D)
	result.AddTag("MMS", 0x0E)

	return result
}

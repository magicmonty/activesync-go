package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func Contacts(protocolVersion byte) CodePage {
	result := NewCodePage("Contacts", 1)

	result.AddTag("Anniversary", 0x05)
	result.AddTag("AssistantName", 0x06)
	result.AddTag("AssistantPhoneNumber", 0x07)
	result.AddTag("Birthday", 0x08)
	result.AddTag("Business2PhoneNumber", 0x0C)
	result.AddTag("BusinessAddressCity", 0x0D)
	result.AddTag("BusinessAddressCountry", 0x0E)
	result.AddTag("BusinessAddressPostalCode", 0x0F)
	result.AddTag("BusinessAddressState", 0x10)
	result.AddTag("BusinessAddressStreet", 0x11)
	result.AddTag("BusinessFaxNumber", 0x12)
	result.AddTag("BusinessPhoneNumber", 0x13)
	result.AddTag("CarPhoneNumber", 0x14)
	result.AddTag("Categories", 0x15)
	result.AddTag("Category", 0x16)
	result.AddTag("Children", 0x17)
	result.AddTag("Child", 0x18)
	result.AddTag("CompanyName", 0x19)
	result.AddTag("Department", 0x1A)
	result.AddTag("Email1Address", 0x1B)
	result.AddTag("Email2Address", 0x1C)
	result.AddTag("Email3Address", 0x1D)
	result.AddTag("FileAs", 0x1E)
	result.AddTag("FirstName", 0x1F)
	result.AddTag("Home2PhoneNumber", 0x20)
	result.AddTag("HomeAddressCity", 0x21)
	result.AddTag("HomeAddressCountry", 0x22)
	result.AddTag("HomeAddressPostalCode", 0x23)
	result.AddTag("HomeAddressState", 0x24)
	result.AddTag("HomeAddressStreet", 0x25)
	result.AddTag("HomeFaxNumber", 0x26)
	result.AddTag("HomePhoneNumber", 0x27)
	result.AddTag("JobTitle", 0x28)
	result.AddTag("LastName", 0x29)
	result.AddTag("MiddleName", 0x2A)
	result.AddTag("MobilePhoneNumber", 0x2B)
	result.AddTag("OfficeLocation", 0x2C)
	result.AddTag("OtherAddressCity", 0x2D)
	result.AddTag("OtherAddressCountry", 0x2E)
	result.AddTag("OtherAddressPostalCode", 0x2F)
	result.AddTag("OtherAddressState", 0x30)
	result.AddTag("OtherAddressStreet", 0x31)
	result.AddTag("PagerNumber", 0x32)
	result.AddTag("RadioPhoneNumber", 0x33)
	result.AddTag("Spouse", 0x34)
	result.AddTag("Suffix", 0x35)
	result.AddTag("Title", 0x36)
	result.AddTag("WebPage", 0x37)
	result.AddTag("YomiCompanyName", 0x38)
	result.AddTag("YomiFirstName", 0x39)
	result.AddTag("YomiLastName", 0x3A)
	result.AddTag("Picture", 0x3C)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag("Alias", 0x3D)        // not supported with MS-ASProtocolVersion 12.1
		result.AddTag("WeightedRank", 0x3E) // not supported with MS-ASProtocolVersion 12.1
	}

	return result
}

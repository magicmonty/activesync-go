package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_CONTACTS                            string = "Contacts"
	TAG_CONTACTS_ANNIVERSARY               string = "Anniversary"
	TAG_CONTACTS_ASSISTANTNAME             string = "AssistantName"
	TAG_CONTACTS_ASSISTANTPHONENUMBER      string = "AssistantPhoneNumber"
	TAG_CONTACTS_BIRTHDAY                  string = "Birthday"
	TAG_CONTACTS_BUSINESS2PHONENUMBER      string = "Business2PhoneNumber"
	TAG_CONTACTS_BUSINESSADDRESSCITY       string = "BusinessAddressCity"
	TAG_CONTACTS_BUSINESSADDRESSCOUNTRY    string = "BusinessAddressCountry"
	TAG_CONTACTS_BUSINESSADDRESSPOSTALCODE string = "BusinessAddressPostalCode"
	TAG_CONTACTS_BUSINESSADDRESSSTATE      string = "BusinessAddressState"
	TAG_CONTACTS_BUSINESSADDRESSSTREET     string = "BusinessAddressStreet"
	TAG_CONTACTS_BUSINESSFAXNUMBER         string = "BusinessFaxNumber"
	TAG_CONTACTS_BUSINESSPHONENUMBER       string = "BusinessPhoneNumber"
	TAG_CONTACTS_CARPHONENUMBER            string = "CarPhoneNumber"
	TAG_CONTACTS_CATEGORIES                string = "Categories"
	TAG_CONTACTS_CATEGORY                  string = "Category"
	TAG_CONTACTS_CHILDREN                  string = "Children"
	TAG_CONTACTS_CHILD                     string = "Child"
	TAG_CONTACTS_COMPANYNAME               string = "CompanyName"
	TAG_CONTACTS_DEPARTMENT                string = "Department"
	TAG_CONTACTS_EMAIL1ADDRESS             string = "Email1Address"
	TAG_CONTACTS_EMAIL2ADDRESS             string = "Email2Address"
	TAG_CONTACTS_EMAIL3ADDRESS             string = "Email3Address"
	TAG_CONTACTS_FILEAS                    string = "FileAs"
	TAG_CONTACTS_FIRSTNAME                 string = "FirstName"
	TAG_CONTACTS_HOME2PHONENUMBER          string = "Home2PhoneNumber"
	TAG_CONTACTS_HOMEADDRESSCITY           string = "HomeAddressCity"
	TAG_CONTACTS_HOMEADDRESSCOUNTRY        string = "HomeAddressCountry"
	TAG_CONTACTS_HOMEADDRESSPOSTALCODE     string = "HomeAddressPostalCode"
	TAG_CONTACTS_HOMEADDRESSSTATE          string = "HomeAddressState"
	TAG_CONTACTS_HOMEADDRESSSTREET         string = "HomeAddressStreet"
	TAG_CONTACTS_HOMEFAXNUMBER             string = "HomeFaxNumber"
	TAG_CONTACTS_HOMEPHONENUMBER           string = "HomePhoneNumber"
	TAG_CONTACTS_JOBTITLE                  string = "JobTitle"
	TAG_CONTACTS_LASTNAME                  string = "LastName"
	TAG_CONTACTS_MIDDLENAME                string = "MiddleName"
	TAG_CONTACTS_MOBILEPHONENUMBER         string = "MobilePhoneNumber"
	TAG_CONTACTS_OFFICELOCATION            string = "OfficeLocation"
	TAG_CONTACTS_OTHERADDRESSCITY          string = "OtherAddressCity"
	TAG_CONTACTS_OTHERADDRESSCOUNTRY       string = "OtherAddressCountry"
	TAG_CONTACTS_OTHERADDRESSPOSTALCODE    string = "OtherAddressPostalCode"
	TAG_CONTACTS_OTHERADDRESSSTATE         string = "OtherAddressState"
	TAG_CONTACTS_OTHERADDRESSSTREET        string = "OtherAddressStreet"
	TAG_CONTACTS_PAGERNUMBER               string = "PagerNumber"
	TAG_CONTACTS_RADIOPHONENUMBER          string = "RadioPhoneNumber"
	TAG_CONTACTS_SPOUSE                    string = "Spouse"
	TAG_CONTACTS_SUFFIX                    string = "Suffix"
	TAG_CONTACTS_TITLE                     string = "Title"
	TAG_CONTACTS_WEBPAGE                   string = "WebPage"
	TAG_CONTACTS_YOMICOMPANYNAME           string = "YomiCompanyName"
	TAG_CONTACTS_YOMIFIRSTNAME             string = "YomiFirstName"
	TAG_CONTACTS_YOMILASTNAME              string = "YomiLastName"
	TAG_CONTACTS_PICTURE                   string = "Picture"
	TAG_CONTACTS_ALIAS                     string = "Alias"
	TAG_CONTACTS_WEIGHTEDRANK              string = "WeightedRank"
)

const (
	CP_CONTACTS                           byte = 1
	ID_CONTACTS_ANNIVERSARY               byte = 0x05
	ID_CONTACTS_ASSISTANTNAME             byte = 0x06
	ID_CONTACTS_ASSISTANTPHONENUMBER      byte = 0x07
	ID_CONTACTS_BIRTHDAY                  byte = 0x08
	ID_CONTACTS_BUSINESS2PHONENUMBER      byte = 0x0C
	ID_CONTACTS_BUSINESSADDRESSCITY       byte = 0x0D
	ID_CONTACTS_BUSINESSADDRESSCOUNTRY    byte = 0x0E
	ID_CONTACTS_BUSINESSADDRESSPOSTALCODE byte = 0x0F
	ID_CONTACTS_BUSINESSADDRESSSTATE      byte = 0x10
	ID_CONTACTS_BUSINESSADDRESSSTREET     byte = 0x11
	ID_CONTACTS_BUSINESSFAXNUMBER         byte = 0x12
	ID_CONTACTS_BUSINESSPHONENUMBER       byte = 0x13
	ID_CONTACTS_CARPHONENUMBER            byte = 0x14
	ID_CONTACTS_CATEGORIES                byte = 0x15
	ID_CONTACTS_CATEGORY                  byte = 0x16
	ID_CONTACTS_CHILDREN                  byte = 0x17
	ID_CONTACTS_CHILD                     byte = 0x18
	ID_CONTACTS_COMPANYNAME               byte = 0x19
	ID_CONTACTS_DEPARTMENT                byte = 0x1A
	ID_CONTACTS_EMAIL1ADDRESS             byte = 0x1B
	ID_CONTACTS_EMAIL2ADDRESS             byte = 0x1C
	ID_CONTACTS_EMAIL3ADDRESS             byte = 0x1D
	ID_CONTACTS_FILEAS                    byte = 0x1E
	ID_CONTACTS_FIRSTNAME                 byte = 0x1F
	ID_CONTACTS_HOME2PHONENUMBER          byte = 0x20
	ID_CONTACTS_HOMEADDRESSCITY           byte = 0x21
	ID_CONTACTS_HOMEADDRESSCOUNTRY        byte = 0x22
	ID_CONTACTS_HOMEADDRESSPOSTALCODE     byte = 0x23
	ID_CONTACTS_HOMEADDRESSSTATE          byte = 0x24
	ID_CONTACTS_HOMEADDRESSSTREET         byte = 0x25
	ID_CONTACTS_HOMEFAXNUMBER             byte = 0x26
	ID_CONTACTS_HOMEPHONENUMBER           byte = 0x27
	ID_CONTACTS_JOBTITLE                  byte = 0x28
	ID_CONTACTS_LASTNAME                  byte = 0x29
	ID_CONTACTS_MIDDLENAME                byte = 0x2A
	ID_CONTACTS_MOBILEPHONENUMBER         byte = 0x2B
	ID_CONTACTS_OFFICELOCATION            byte = 0x2C
	ID_CONTACTS_OTHERADDRESSCITY          byte = 0x2D
	ID_CONTACTS_OTHERADDRESSCOUNTRY       byte = 0x2E
	ID_CONTACTS_OTHERADDRESSPOSTALCODE    byte = 0x2F
	ID_CONTACTS_OTHERADDRESSSTATE         byte = 0x30
	ID_CONTACTS_OTHERADDRESSSTREET        byte = 0x31
	ID_CONTACTS_PAGERNUMBER               byte = 0x32
	ID_CONTACTS_RADIOPHONENUMBER          byte = 0x33
	ID_CONTACTS_SPOUSE                    byte = 0x34
	ID_CONTACTS_SUFFIX                    byte = 0x35
	ID_CONTACTS_TITLE                     byte = 0x36
	ID_CONTACTS_WEBPAGE                   byte = 0x37
	ID_CONTACTS_YOMICOMPANYNAME           byte = 0x38
	ID_CONTACTS_YOMIFIRSTNAME             byte = 0x39
	ID_CONTACTS_YOMILASTNAME              byte = 0x3A
	ID_CONTACTS_PICTURE                   byte = 0x3C
	ID_CONTACTS_ALIAS                     byte = 0x3D
	ID_CONTACTS_WEIGHTEDRANK              byte = 0x3E
)

func Contacts(protocolVersion byte) CodePage {
	result := NewCodePage(NS_CONTACTS, CP_CONTACTS)

	result.AddTag(TAG_CONTACTS_ANNIVERSARY, ID_CONTACTS_ANNIVERSARY)
	result.AddTag(TAG_CONTACTS_ASSISTANTNAME, ID_CONTACTS_ASSISTANTNAME)
	result.AddTag(TAG_CONTACTS_ASSISTANTPHONENUMBER, ID_CONTACTS_ASSISTANTPHONENUMBER)
	result.AddTag(TAG_CONTACTS_BIRTHDAY, ID_CONTACTS_BIRTHDAY)
	result.AddTag(TAG_CONTACTS_BUSINESS2PHONENUMBER, ID_CONTACTS_BUSINESS2PHONENUMBER)
	result.AddTag(TAG_CONTACTS_BUSINESSADDRESSCITY, ID_CONTACTS_BUSINESSADDRESSCITY)
	result.AddTag(TAG_CONTACTS_BUSINESSADDRESSCOUNTRY, ID_CONTACTS_BUSINESSADDRESSCOUNTRY)
	result.AddTag(TAG_CONTACTS_BUSINESSADDRESSPOSTALCODE, ID_CONTACTS_BUSINESSADDRESSPOSTALCODE)
	result.AddTag(TAG_CONTACTS_BUSINESSADDRESSSTATE, ID_CONTACTS_BUSINESSADDRESSSTATE)
	result.AddTag(TAG_CONTACTS_BUSINESSADDRESSSTREET, ID_CONTACTS_BUSINESSADDRESSSTREET)
	result.AddTag(TAG_CONTACTS_BUSINESSFAXNUMBER, ID_CONTACTS_BUSINESSFAXNUMBER)
	result.AddTag(TAG_CONTACTS_BUSINESSPHONENUMBER, ID_CONTACTS_BUSINESSPHONENUMBER)
	result.AddTag(TAG_CONTACTS_CARPHONENUMBER, ID_CONTACTS_CARPHONENUMBER)
	result.AddTag(TAG_CONTACTS_CATEGORIES, ID_CONTACTS_CATEGORIES)
	result.AddTag(TAG_CONTACTS_CATEGORY, ID_CONTACTS_CATEGORY)
	result.AddTag(TAG_CONTACTS_CHILDREN, ID_CONTACTS_CHILDREN)
	result.AddTag(TAG_CONTACTS_CHILD, ID_CONTACTS_CHILD)
	result.AddTag(TAG_CONTACTS_COMPANYNAME, ID_CONTACTS_COMPANYNAME)
	result.AddTag(TAG_CONTACTS_DEPARTMENT, ID_CONTACTS_DEPARTMENT)
	result.AddTag(TAG_CONTACTS_EMAIL1ADDRESS, ID_CONTACTS_EMAIL1ADDRESS)
	result.AddTag(TAG_CONTACTS_EMAIL2ADDRESS, ID_CONTACTS_EMAIL2ADDRESS)
	result.AddTag(TAG_CONTACTS_EMAIL3ADDRESS, ID_CONTACTS_EMAIL3ADDRESS)
	result.AddTag(TAG_CONTACTS_FILEAS, ID_CONTACTS_FILEAS)
	result.AddTag(TAG_CONTACTS_FIRSTNAME, ID_CONTACTS_FIRSTNAME)
	result.AddTag(TAG_CONTACTS_HOME2PHONENUMBER, ID_CONTACTS_HOME2PHONENUMBER)
	result.AddTag(TAG_CONTACTS_HOMEADDRESSCITY, ID_CONTACTS_HOMEADDRESSCITY)
	result.AddTag(TAG_CONTACTS_HOMEADDRESSCOUNTRY, ID_CONTACTS_HOMEADDRESSCOUNTRY)
	result.AddTag(TAG_CONTACTS_HOMEADDRESSPOSTALCODE, ID_CONTACTS_HOMEADDRESSPOSTALCODE)
	result.AddTag(TAG_CONTACTS_HOMEADDRESSSTATE, ID_CONTACTS_HOMEADDRESSSTATE)
	result.AddTag(TAG_CONTACTS_HOMEADDRESSSTREET, ID_CONTACTS_HOMEADDRESSSTREET)
	result.AddTag(TAG_CONTACTS_HOMEFAXNUMBER, ID_CONTACTS_HOMEFAXNUMBER)
	result.AddTag(TAG_CONTACTS_HOMEPHONENUMBER, ID_CONTACTS_HOMEPHONENUMBER)
	result.AddTag(TAG_CONTACTS_JOBTITLE, ID_CONTACTS_JOBTITLE)
	result.AddTag(TAG_CONTACTS_LASTNAME, ID_CONTACTS_LASTNAME)
	result.AddTag(TAG_CONTACTS_MIDDLENAME, ID_CONTACTS_MIDDLENAME)
	result.AddTag(TAG_CONTACTS_MOBILEPHONENUMBER, ID_CONTACTS_MOBILEPHONENUMBER)
	result.AddTag(TAG_CONTACTS_OFFICELOCATION, ID_CONTACTS_OFFICELOCATION)
	result.AddTag(TAG_CONTACTS_OTHERADDRESSCITY, ID_CONTACTS_OTHERADDRESSCITY)
	result.AddTag(TAG_CONTACTS_OTHERADDRESSCOUNTRY, ID_CONTACTS_OTHERADDRESSCOUNTRY)
	result.AddTag(TAG_CONTACTS_OTHERADDRESSPOSTALCODE, ID_CONTACTS_OTHERADDRESSPOSTALCODE)
	result.AddTag(TAG_CONTACTS_OTHERADDRESSSTATE, ID_CONTACTS_OTHERADDRESSSTATE)
	result.AddTag(TAG_CONTACTS_OTHERADDRESSSTREET, ID_CONTACTS_OTHERADDRESSSTREET)
	result.AddTag(TAG_CONTACTS_PAGERNUMBER, ID_CONTACTS_PAGERNUMBER)
	result.AddTag(TAG_CONTACTS_RADIOPHONENUMBER, ID_CONTACTS_RADIOPHONENUMBER)
	result.AddTag(TAG_CONTACTS_SPOUSE, ID_CONTACTS_SPOUSE)
	result.AddTag(TAG_CONTACTS_SUFFIX, ID_CONTACTS_SUFFIX)
	result.AddTag(TAG_CONTACTS_TITLE, ID_CONTACTS_TITLE)
	result.AddTag(TAG_CONTACTS_WEBPAGE, ID_CONTACTS_WEBPAGE)
	result.AddTag(TAG_CONTACTS_YOMICOMPANYNAME, ID_CONTACTS_YOMICOMPANYNAME)
	result.AddTag(TAG_CONTACTS_YOMIFIRSTNAME, ID_CONTACTS_YOMIFIRSTNAME)
	result.AddTag(TAG_CONTACTS_YOMILASTNAME, ID_CONTACTS_YOMILASTNAME)
	result.AddTag(TAG_CONTACTS_PICTURE, ID_CONTACTS_PICTURE)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_CONTACTS_ALIAS, ID_CONTACTS_ALIAS)               // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_CONTACTS_WEIGHTEDRANK, ID_CONTACTS_WEIGHTEDRANK) // not supported with MS-ASProtocolVersion 12.1
	}

	return result
}

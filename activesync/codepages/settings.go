package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_SETTINGS                              string = "Settings"
	TAG_SETTINGS_SETTINGS                    string = "Settings"
	TAG_SETTINGS_STATUS                      string = "Status"
	TAG_SETTINGS_GET                         string = "Get"
	TAG_SETTINGS_SET                         string = "Set"
	TAG_SETTINGS_OOF                         string = "Oof"
	TAG_SETTINGS_OOFSTATE                    string = "OofState"
	TAG_SETTINGS_STARTTIME                   string = "StartTime"
	TAG_SETTINGS_ENDTIME                     string = "EndTime"
	TAG_SETTINGS_OOFMESSAGE                  string = "OofMessage"
	TAG_SETTINGS_APPLIESTOINTERNAL           string = "AppliesToInternal"
	TAG_SETTINGS_APPLIESTOEXTERNALKNOWN      string = "AppliesToExternalKnown"
	TAG_SETTINGS_APPLIESTOEXTERNALUNKNOWN    string = "AppliesToExternalUnknown"
	TAG_SETTINGS_ENABLED                     string = "Enabled"
	TAG_SETTINGS_REPLYMESSAGE                string = "ReplyMessage"
	TAG_SETTINGS_BODYTYPE                    string = "BodyType"
	TAG_SETTINGS_DEVICEPASSWORD              string = "DevicePassword"
	TAG_SETTINGS_PASSWORD                    string = "Password"
	TAG_SETTINGS_DEVICEINFORMATION           string = "DeviceInformation"
	TAG_SETTINGS_MODEL                       string = "Model"
	TAG_SETTINGS_IMEI                        string = "IMEI"
	TAG_SETTINGS_FRIENDLYNAME                string = "FriendlyName"
	TAG_SETTINGS_OS                          string = "OS"
	TAG_SETTINGS_OSLANGUAGE                  string = "OSLanguage"
	TAG_SETTINGS_PHONENUMBER                 string = "PhoneNumber"
	TAG_SETTINGS_USERINFORMATION             string = "UserInformation"
	TAG_SETTINGS_EMAILADDRESSES              string = "EmailAddresses"
	TAG_SETTINGS_SMTPADDRESS                 string = "SmtpAddress"
	TAG_SETTINGS_USERAGENT                   string = "UserAgent"
	TAG_SETTINGS_ENABLEOUTBOUNDSMS           string = "EnableOutboundSMS"
	TAG_SETTINGS_MOBILEOPERATOR              string = "MobileOperator"
	TAG_SETTINGS_PRIMARYSMTPADDRESS          string = "PrimarySmtpAddress"
	TAG_SETTINGS_ACCOUNTS                    string = "Accounts"
	TAG_SETTINGS_ACCOUNT                     string = "Account"
	TAG_SETTINGS_ACCOUNTID                   string = "AccountId"
	TAG_SETTINGS_ACCOUNTNAME                 string = "AccountName"
	TAG_SETTINGS_USERDISPLAYNAME             string = "UserDisplayName"
	TAG_SETTINGS_SENDDISABLED                string = "SendDisabled"
	TAG_SETTINGS_RIGHTSMANAGEMENTINFORMATION string = "RightsManagementInformation"
)

const (
	CP_SETTINGS                             byte = 18
	ID_SETTINGS_SETTINGS                    byte = 0x05
	ID_SETTINGS_STATUS                      byte = 0x06
	ID_SETTINGS_GET                         byte = 0x07
	ID_SETTINGS_SET                         byte = 0x08
	ID_SETTINGS_OOF                         byte = 0x09
	ID_SETTINGS_OOFSTATE                    byte = 0x0A
	ID_SETTINGS_STARTTIME                   byte = 0x0B
	ID_SETTINGS_ENDTIME                     byte = 0x0C
	ID_SETTINGS_OOFMESSAGE                  byte = 0x0D
	ID_SETTINGS_APPLIESTOINTERNAL           byte = 0x0E
	ID_SETTINGS_APPLIESTOEXTERNALKNOWN      byte = 0x0F
	ID_SETTINGS_APPLIESTOEXTERNALUNKNOWN    byte = 0x10
	ID_SETTINGS_ENABLED                     byte = 0x11
	ID_SETTINGS_REPLYMESSAGE                byte = 0x12
	ID_SETTINGS_BODYTYPE                    byte = 0x13
	ID_SETTINGS_DEVICEPASSWORD              byte = 0x14
	ID_SETTINGS_PASSWORD                    byte = 0x15
	ID_SETTINGS_DEVICEINFORMATION           byte = 0x16
	ID_SETTINGS_MODEL                       byte = 0x17
	ID_SETTINGS_IMEI                        byte = 0x18
	ID_SETTINGS_FRIENDLYNAME                byte = 0x19
	ID_SETTINGS_OS                          byte = 0x1A
	ID_SETTINGS_OSLANGUAGE                  byte = 0x1B
	ID_SETTINGS_PHONENUMBER                 byte = 0x1C
	ID_SETTINGS_USERINFORMATION             byte = 0x1D
	ID_SETTINGS_EMAILADDRESSES              byte = 0x1E
	ID_SETTINGS_SMTPADDRESS                 byte = 0x1F
	ID_SETTINGS_USERAGENT                   byte = 0x20
	ID_SETTINGS_ENABLEOUTBOUNDSMS           byte = 0x21
	ID_SETTINGS_MOBILEOPERATOR              byte = 0x22
	ID_SETTINGS_PRIMARYSMTPADDRESS          byte = 0x23
	ID_SETTINGS_ACCOUNTS                    byte = 0x24
	ID_SETTINGS_ACCOUNT                     byte = 0x25
	ID_SETTINGS_ACCOUNTID                   byte = 0x26
	ID_SETTINGS_ACCOUNTNAME                 byte = 0x27
	ID_SETTINGS_USERDISPLAYNAME             byte = 0x28
	ID_SETTINGS_SENDDISABLED                byte = 0x29
	ID_SETTINGS_RIGHTSMANAGEMENTINFORMATION byte = 0x2B
)

func Settings(protocolVersion byte) CodePage {
	result := NewCodePage(NS_SETTINGS, CP_SETTINGS)

	result.AddTag(TAG_SETTINGS_SETTINGS, ID_SETTINGS_SETTINGS)
	result.AddTag(TAG_SETTINGS_STATUS, ID_SETTINGS_STATUS)
	result.AddTag(TAG_SETTINGS_GET, ID_SETTINGS_GET)
	result.AddTag(TAG_SETTINGS_SET, ID_SETTINGS_SET)
	result.AddTag(TAG_SETTINGS_OOF, ID_SETTINGS_OOF)
	result.AddTag(TAG_SETTINGS_OOFSTATE, ID_SETTINGS_OOFSTATE)
	result.AddTag(TAG_SETTINGS_STARTTIME, ID_SETTINGS_STARTTIME)
	result.AddTag(TAG_SETTINGS_ENDTIME, ID_SETTINGS_ENDTIME)
	result.AddTag(TAG_SETTINGS_OOFMESSAGE, ID_SETTINGS_OOFMESSAGE)
	result.AddTag(TAG_SETTINGS_APPLIESTOINTERNAL, ID_SETTINGS_APPLIESTOINTERNAL)
	result.AddTag(TAG_SETTINGS_APPLIESTOEXTERNALKNOWN, ID_SETTINGS_APPLIESTOEXTERNALKNOWN)
	result.AddTag(TAG_SETTINGS_APPLIESTOEXTERNALUNKNOWN, ID_SETTINGS_APPLIESTOEXTERNALUNKNOWN)
	result.AddTag(TAG_SETTINGS_ENABLED, ID_SETTINGS_ENABLED)
	result.AddTag(TAG_SETTINGS_REPLYMESSAGE, ID_SETTINGS_REPLYMESSAGE)
	result.AddTag(TAG_SETTINGS_BODYTYPE, ID_SETTINGS_BODYTYPE)
	result.AddTag(TAG_SETTINGS_DEVICEPASSWORD, ID_SETTINGS_DEVICEPASSWORD)
	result.AddTag(TAG_SETTINGS_PASSWORD, ID_SETTINGS_PASSWORD)
	result.AddTag(TAG_SETTINGS_DEVICEINFORMATION, ID_SETTINGS_DEVICEINFORMATION)
	result.AddTag(TAG_SETTINGS_MODEL, ID_SETTINGS_MODEL)
	result.AddTag(TAG_SETTINGS_IMEI, ID_SETTINGS_IMEI)
	result.AddTag(TAG_SETTINGS_FRIENDLYNAME, ID_SETTINGS_FRIENDLYNAME)
	result.AddTag(TAG_SETTINGS_OS, ID_SETTINGS_OS)
	result.AddTag(TAG_SETTINGS_OSLANGUAGE, ID_SETTINGS_OSLANGUAGE)
	result.AddTag(TAG_SETTINGS_PHONENUMBER, ID_SETTINGS_PHONENUMBER)
	result.AddTag(TAG_SETTINGS_USERINFORMATION, ID_SETTINGS_USERINFORMATION)
	result.AddTag(TAG_SETTINGS_EMAILADDRESSES, ID_SETTINGS_EMAILADDRESSES)
	result.AddTag(TAG_SETTINGS_SMTPADDRESS, ID_SETTINGS_SMTPADDRESS)
	result.AddTag(TAG_SETTINGS_USERAGENT, ID_SETTINGS_USERAGENT)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_SETTINGS_ENABLEOUTBOUNDSMS, ID_SETTINGS_ENABLEOUTBOUNDSMS) // not supported with MS-ASProtocolVersion 12.1
		result.AddTag(TAG_SETTINGS_MOBILEOPERATOR, ID_SETTINGS_MOBILEOPERATOR)       // not supported with MS-ASProtocolVersion 12.1
		if protocolVersion > PROTOCOL_VERSION_14_0 {
			result.AddTag(TAG_SETTINGS_PRIMARYSMTPADDRESS, ID_SETTINGS_PRIMARYSMTPADDRESS)                   // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_SETTINGS_ACCOUNTS, ID_SETTINGS_ACCOUNTS)                                       // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_SETTINGS_ACCOUNT, ID_SETTINGS_ACCOUNT)                                         // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_SETTINGS_ACCOUNTID, ID_SETTINGS_ACCOUNTID)                                     // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_SETTINGS_ACCOUNTNAME, ID_SETTINGS_ACCOUNTNAME)                                 // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_SETTINGS_USERDISPLAYNAME, ID_SETTINGS_USERDISPLAYNAME)                         // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_SETTINGS_SENDDISABLED, ID_SETTINGS_SENDDISABLED)                               // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_SETTINGS_RIGHTSMANAGEMENTINFORMATION, ID_SETTINGS_RIGHTSMANAGEMENTINFORMATION) // not supported with MS-ASProtocolVersion 12.1 or 14.0
		}
	}

	return result
}

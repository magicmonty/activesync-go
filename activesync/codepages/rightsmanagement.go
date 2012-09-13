package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

// not supported with MS-ASProtocolVersion 12.1 or 14.0
func RightsManagement() CodePage {
	result := NewCodePage("RightsManagement", 24)

	result.AddTag("RightsManagementSupport", 0x05)
	result.AddTag("RightsManagementTemplates", 0x06)
	result.AddTag("RightsManagementTemplate", 0x07)
	result.AddTag("RightsManagementLicense", 0x08)
	result.AddTag("EditAllowed", 0x09)
	result.AddTag("ReplyAllowed", 0x0A)
	result.AddTag("ReplyAllAllowed", 0x0B)
	result.AddTag("ForwardAllowed", 0x0C)
	result.AddTag("ModifyRecipientsAllowed", 0x0D)
	result.AddTag("ExtractAllowed", 0x0E)
	result.AddTag("PrintAllowed", 0x0F)
	result.AddTag("ExportAllowed", 0x10)
	result.AddTag("ProgrammaticAccessAllowed", 0x11)
	result.AddTag("Owner", 0x12)
	result.AddTag("ContentExpiryDate", 0x13)
	result.AddTag("TemplateID", 0x14)
	result.AddTag("TemplateName", 0x15)
	result.AddTag("TemplateDescription", 0x16)
	result.AddTag("ContentOwner", 0x17)
	result.AddTag("RemoveRightsManagementDistribution", 0x18)

	return result
}

package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_AIRSYNC_BASE                     string = "AirSyncBase"
	TAG_AIRSYNC_BASE_BODYPREFERENCE     string = "BodyPreference"
	TAG_AIRSYNC_BASE_TYPE               string = "Type"
	TAG_AIRSYNC_BASE_TRUNCATIONSIZE     string = "TruncationSize"
	TAG_AIRSYNC_BASE_ALLORNONE          string = "AllOrNone"
	TAG_AIRSYNC_BASE_BODY               string = "Body"
	TAG_AIRSYNC_BASE_DATA               string = "Data"
	TAG_AIRSYNC_BASE_ESTIMATEDDATASIZE  string = "EstimatedDataSize"
	TAG_AIRSYNC_BASE_TRUNCATED          string = "Truncated"
	TAG_AIRSYNC_BASE_ATTACHMENTS        string = "Attachments"
	TAG_AIRSYNC_BASE_ATTACHMENT         string = "Attachment"
	TAG_AIRSYNC_BASE_DISPLAYNAME        string = "DisplayName"
	TAG_AIRSYNC_BASE_FILEREFERENCE      string = "FileReference"
	TAG_AIRSYNC_BASE_METHOD             string = "Method"
	TAG_AIRSYNC_BASE_CONTENTID          string = "ContentId"
	TAG_AIRSYNC_BASE_CONTENTLOCATION    string = "ContentLocation"
	TAG_AIRSYNC_BASE_ISINLINE           string = "IsInline"
	TAG_AIRSYNC_BASE_NATIVEBODYTYPE     string = "NativeBodyType"
	TAG_AIRSYNC_BASE_CONTENTTYPE        string = "ContentType"
	TAG_AIRSYNC_BASE_PREVIEW            string = "Preview"
	TAG_AIRSYNC_BASE_BODYPARTPREFERENCE string = "BodyPartPreference"
	TAG_AIRSYNC_BASE_BODYPART           string = "BodyPart"
	TAG_AIRSYNC_BASE_STATUS             string = "Status"
)

const (
	CP_AIRSYNC_BASE                    byte = 17
	ID_AIRSYNC_BASE_BODYPREFERENCE     byte = 0x05
	ID_AIRSYNC_BASE_TYPE               byte = 0x06
	ID_AIRSYNC_BASE_TRUNCATIONSIZE     byte = 0x07
	ID_AIRSYNC_BASE_ALLORNONE          byte = 0x08
	ID_AIRSYNC_BASE_BODY               byte = 0x0A
	ID_AIRSYNC_BASE_DATA               byte = 0x0B
	ID_AIRSYNC_BASE_ESTIMATEDDATASIZE  byte = 0x0C
	ID_AIRSYNC_BASE_TRUNCATED          byte = 0x0D
	ID_AIRSYNC_BASE_ATTACHMENTS        byte = 0x0E
	ID_AIRSYNC_BASE_ATTACHMENT         byte = 0x0F
	ID_AIRSYNC_BASE_DISPLAYNAME        byte = 0x10
	ID_AIRSYNC_BASE_FILEREFERENCE      byte = 0x11
	ID_AIRSYNC_BASE_METHOD             byte = 0x12
	ID_AIRSYNC_BASE_CONTENTID          byte = 0x13
	ID_AIRSYNC_BASE_CONTENTLOCATION    byte = 0x14
	ID_AIRSYNC_BASE_ISINLINE           byte = 0x15
	ID_AIRSYNC_BASE_NATIVEBODYTYPE     byte = 0x16
	ID_AIRSYNC_BASE_CONTENTTYPE        byte = 0x17
	ID_AIRSYNC_BASE_PREVIEW            byte = 0x18
	ID_AIRSYNC_BASE_BODYPARTPREFERENCE byte = 0x19
	ID_AIRSYNC_BASE_BODYPART           byte = 0x1A
	ID_AIRSYNC_BASE_STATUS             byte = 0x1B
)

func AirSyncBase(protocolVersion byte) CodePage {
	result := NewCodePage(NS_AIRSYNC_BASE, CP_AIRSYNC_BASE)

	result.AddTag(TAG_AIRSYNC_BASE_BODYPREFERENCE, ID_AIRSYNC_BASE_BODYPREFERENCE)
	result.AddTag(TAG_AIRSYNC_BASE_TYPE, ID_AIRSYNC_BASE_TYPE)
	result.AddTag(TAG_AIRSYNC_BASE_TRUNCATIONSIZE, ID_AIRSYNC_BASE_TRUNCATIONSIZE)
	result.AddTag(TAG_AIRSYNC_BASE_ALLORNONE, ID_AIRSYNC_BASE_ALLORNONE)
	result.AddTag(TAG_AIRSYNC_BASE_BODY, ID_AIRSYNC_BASE_BODY)
	result.AddTag(TAG_AIRSYNC_BASE_DATA, ID_AIRSYNC_BASE_DATA)
	result.AddTag(TAG_AIRSYNC_BASE_ESTIMATEDDATASIZE, ID_AIRSYNC_BASE_ESTIMATEDDATASIZE)
	result.AddTag(TAG_AIRSYNC_BASE_TRUNCATED, ID_AIRSYNC_BASE_TRUNCATED)
	result.AddTag(TAG_AIRSYNC_BASE_ATTACHMENTS, ID_AIRSYNC_BASE_ATTACHMENTS)
	result.AddTag(TAG_AIRSYNC_BASE_ATTACHMENT, ID_AIRSYNC_BASE_ATTACHMENT)
	result.AddTag(TAG_AIRSYNC_BASE_DISPLAYNAME, ID_AIRSYNC_BASE_DISPLAYNAME)
	result.AddTag(TAG_AIRSYNC_BASE_FILEREFERENCE, ID_AIRSYNC_BASE_FILEREFERENCE)
	result.AddTag(TAG_AIRSYNC_BASE_METHOD, ID_AIRSYNC_BASE_METHOD)
	result.AddTag(TAG_AIRSYNC_BASE_CONTENTID, ID_AIRSYNC_BASE_CONTENTID)
	result.AddTag(TAG_AIRSYNC_BASE_CONTENTLOCATION, ID_AIRSYNC_BASE_CONTENTLOCATION) // not used
	result.AddTag(TAG_AIRSYNC_BASE_ISINLINE, ID_AIRSYNC_BASE_ISINLINE)
	result.AddTag(TAG_AIRSYNC_BASE_NATIVEBODYTYPE, ID_AIRSYNC_BASE_NATIVEBODYTYPE)
	result.AddTag(TAG_AIRSYNC_BASE_CONTENTTYPE, ID_AIRSYNC_BASE_CONTENTTYPE)
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTag(TAG_AIRSYNC_BASE_PREVIEW, ID_AIRSYNC_BASE_PREVIEW) // not supported with MS-ASProtocolVersion 12.1
		if protocolVersion > PROTOCOL_VERSION_14_0 {
			result.AddTag(TAG_AIRSYNC_BASE_BODYPARTPREFERENCE, ID_AIRSYNC_BASE_BODYPARTPREFERENCE) // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_AIRSYNC_BASE_BODYPART, ID_AIRSYNC_BASE_BODYPART)                     // not supported with MS-ASProtocolVersion 12.1 or 14.0
			result.AddTag(TAG_AIRSYNC_BASE_STATUS, ID_AIRSYNC_BASE_STATUS)                         // not supported with MS-ASProtocolVersion 12.1 or 14.0
		}
	}

	return result
}

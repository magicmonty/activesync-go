package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_COMPOSE_MAIL                  string = "ComposeMail"
	TAG_COMPOSE_MAIL_SENDMAIL        string = "SendMail"
	TAG_COMPOSE_MAIL_SMARTFORWARD    string = "SmartForward"
	TAG_COMPOSE_MAIL_SMARTREPLY      string = "SmartReply"
	TAG_COMPOSE_MAIL_SAVEINSENTITEMS string = "SaveInSentItems"
	TAG_COMPOSE_MAIL_REPLACEMIME     string = "ReplaceMime"
	TAG_COMPOSE_MAIL_SOURCE          string = "Source"
	TAG_COMPOSE_MAIL_FOLDERID        string = "FolderId"
	TAG_COMPOSE_MAIL_ITEMID          string = "ItemId"
	TAG_COMPOSE_MAIL_LONGID          string = "LongId"
	TAG_COMPOSE_MAIL_INSTANCEID      string = "InstanceId"
	TAG_COMPOSE_MAIL_MIME            string = "Mime"
	TAG_COMPOSE_MAIL_CLIENTID        string = "ClientId"
	TAG_COMPOSE_MAIL_STATUS          string = "Status"
	TAG_COMPOSE_MAIL_ACCOUNTID       string = "AccountId"
)

const (
	CP_COMPOSE_MAIL                 byte = 21
	ID_COMPOSE_MAIL_SENDMAIL        byte = 0x05
	ID_COMPOSE_MAIL_SMARTFORWARD    byte = 0x06
	ID_COMPOSE_MAIL_SMARTREPLY      byte = 0x07
	ID_COMPOSE_MAIL_SAVEINSENTITEMS byte = 0x08
	ID_COMPOSE_MAIL_REPLACEMIME     byte = 0x09
	ID_COMPOSE_MAIL_SOURCE          byte = 0x0B
	ID_COMPOSE_MAIL_FOLDERID        byte = 0x0C
	ID_COMPOSE_MAIL_ITEMID          byte = 0x0D
	ID_COMPOSE_MAIL_LONGID          byte = 0x0E
	ID_COMPOSE_MAIL_INSTANCEID      byte = 0x0F
	ID_COMPOSE_MAIL_MIME            byte = 0x10
	ID_COMPOSE_MAIL_CLIENTID        byte = 0x11
	ID_COMPOSE_MAIL_STATUS          byte = 0x12
	ID_COMPOSE_MAIL_ACCOUNTID       byte = 0x13
)

// not supported with MS-ASProtocolVersion 12.1
func ComposeMail(protocolVersion byte) CodePage {
	result := NewCodePage(NS_COMPOSE_MAIL, CP_COMPOSE_MAIL)

	result.AddTag(TAG_COMPOSE_MAIL_SENDMAIL, ID_COMPOSE_MAIL_SENDMAIL)
	result.AddTag(TAG_COMPOSE_MAIL_SMARTFORWARD, ID_COMPOSE_MAIL_SMARTFORWARD)
	result.AddTag(TAG_COMPOSE_MAIL_SMARTREPLY, ID_COMPOSE_MAIL_SMARTREPLY)
	result.AddTag(TAG_COMPOSE_MAIL_SAVEINSENTITEMS, ID_COMPOSE_MAIL_SAVEINSENTITEMS)
	result.AddTag(TAG_COMPOSE_MAIL_REPLACEMIME, ID_COMPOSE_MAIL_REPLACEMIME)
	result.AddTag(TAG_COMPOSE_MAIL_SOURCE, ID_COMPOSE_MAIL_SOURCE)
	result.AddTag(TAG_COMPOSE_MAIL_FOLDERID, ID_COMPOSE_MAIL_FOLDERID)
	result.AddTag(TAG_COMPOSE_MAIL_ITEMID, ID_COMPOSE_MAIL_ITEMID)
	result.AddTag(TAG_COMPOSE_MAIL_LONGID, ID_COMPOSE_MAIL_LONGID)
	result.AddTag(TAG_COMPOSE_MAIL_INSTANCEID, ID_COMPOSE_MAIL_INSTANCEID)
	result.AddTag(TAG_COMPOSE_MAIL_MIME, ID_COMPOSE_MAIL_MIME)
	result.AddTag(TAG_COMPOSE_MAIL_CLIENTID, ID_COMPOSE_MAIL_CLIENTID)
	result.AddTag(TAG_COMPOSE_MAIL_STATUS, ID_COMPOSE_MAIL_STATUS)
	if protocolVersion > PROTOCOL_VERSION_14_0 {
		result.AddTag(TAG_COMPOSE_MAIL_ACCOUNTID, ID_COMPOSE_MAIL_ACCOUNTID) // not supported with MS-ASProtocolVersion 12.1 or 14.0
	}

	return result
}

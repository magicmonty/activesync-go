package codepages

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	. "github.com/magicmonty/wbxml-go/wbxml"
	"testing"
)

func checkComposeMailCommonTags(t *testing.T, cp CodePage) {
	checkTagExists(t, TAG_COMPOSE_MAIL_SENDMAIL, ID_COMPOSE_MAIL_SENDMAIL, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_SMARTFORWARD, ID_COMPOSE_MAIL_SMARTFORWARD, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_SMARTREPLY, ID_COMPOSE_MAIL_SMARTREPLY, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_SAVEINSENTITEMS, ID_COMPOSE_MAIL_SAVEINSENTITEMS, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_REPLACEMIME, ID_COMPOSE_MAIL_REPLACEMIME, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_SOURCE, ID_COMPOSE_MAIL_SOURCE, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_FOLDERID, ID_COMPOSE_MAIL_FOLDERID, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_ITEMID, ID_COMPOSE_MAIL_ITEMID, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_LONGID, ID_COMPOSE_MAIL_LONGID, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_INSTANCEID, ID_COMPOSE_MAIL_INSTANCEID, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_MIME, ID_COMPOSE_MAIL_MIME, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_CLIENTID, ID_COMPOSE_MAIL_CLIENTID, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_STATUS, ID_COMPOSE_MAIL_STATUS, cp)
}

func Test_ComposeMail_120(t *testing.T) {
	cp := ComposeMail(PROTOCOL_VERSION_12_0)
	checkTagCount(t, 0, cp)
}

func Test_ComposeMail_121(t *testing.T) {
	cp := ComposeMail(PROTOCOL_VERSION_12_1)
	checkTagCount(t, 0, cp)
}

func Test_ComposeMail_140(t *testing.T) {
	cp := ComposeMail(PROTOCOL_VERSION_14_0)
	checkComposeMailCommonTags(t, cp)
	checkTagCount(t, 13, cp)
	checkTagNotExists(t, TAG_COMPOSE_MAIL_ACCOUNTID, ID_COMPOSE_MAIL_ACCOUNTID, cp)
}

func Test_ComposeMail_141(t *testing.T) {
	cp := ComposeMail(PROTOCOL_VERSION_14_1)
	checkComposeMailCommonTags(t, cp)
	checkTagCount(t, 14, cp)
	checkTagExists(t, TAG_COMPOSE_MAIL_ACCOUNTID, ID_COMPOSE_MAIL_ACCOUNTID, cp)
}

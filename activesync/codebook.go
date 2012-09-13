package activesync

import (
	. "github.com/magicmonty/activesync-go/activesync/base"
	"github.com/magicmonty/activesync-go/activesync/codepages"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func MakeCodeBook(protocolVersion byte) *CodeBook {
	result := NewCodeBook()

	result.AddTagCodePage(codepages.AirSync(protocolVersion))
	result.AddTagCodePage(codepages.Contacts(protocolVersion))
	result.AddTagCodePage(codepages.Email(protocolVersion))
	// CodePage 3 (AirNotify) is not used any longer
	result.AddTagCodePage(codepages.Calendar(protocolVersion))
	result.AddTagCodePage(codepages.Move())
	result.AddTagCodePage(codepages.ItemEstimate(protocolVersion))
	result.AddTagCodePage(codepages.FolderHierarchy())
	result.AddTagCodePage(codepages.MeetingResponse())
	result.AddTagCodePage(codepages.Tasks(protocolVersion))
	result.AddTagCodePage(codepages.ResolveRecipients(protocolVersion))
	result.AddTagCodePage(codepages.ValidateCert())
	result.AddTagCodePage(codepages.Contacts2())
	result.AddTagCodePage(codepages.Ping())
	result.AddTagCodePage(codepages.Provision())
	result.AddTagCodePage(codepages.Search(protocolVersion))
	result.AddTagCodePage(codepages.GAL(protocolVersion))
	result.AddTagCodePage(codepages.AirSyncBase(protocolVersion))
	result.AddTagCodePage(codepages.Settings(protocolVersion))
	result.AddTagCodePage(codepages.DocumentLibrary())
	result.AddTagCodePage(codepages.ItemOperations(protocolVersion))
	if protocolVersion > PROTOCOL_VERSION_12_1 {
		result.AddTagCodePage(codepages.ComposeMail(protocolVersion))
		result.AddTagCodePage(codepages.Email2(protocolVersion))
		result.AddTagCodePage(codepages.Notes())
		if protocolVersion > PROTOCOL_VERSION_14_0 {
			result.AddTagCodePage(codepages.RightsManagement())
		}
	}

	return result
}

package base

const (
	PROTOCOL_VERSION_12_0 byte = 120
	PROTOCOL_VERSION_12_1 byte = 121
	PROTOCOL_VERSION_14_0 byte = 140
	PROTOCOL_VERSION_14_1 byte = 141
)

type CommandCode byte

const (
	CCSync              CommandCode = iota // 0
	CCSendMail                             // 1
	CCSmartForward                         // 2
	CCSmartReply                           // 3
	CCGetAttachment                        // 4
	_                                      // skip 5
	_                                      // skip 6
	_                                      // skip 7
	_                                      // skip 8
	CCFolderSync                           // 9
	CCFolderCreate                         // 10
	CCFolderDelete                         // 11
	CCFolderUpdate                         // 12
	CCMoveItems                            // 13
	CCGetItemEstimate                      // 14
	CCMeetingResponse                      // 15
	CCSearch                               // 16
	CCSettings                             // 17
	CCPing                                 // 18
	CCItemOperations                       // 19
	CCProvision                            // 20
	CCResolveRecipients                    // 21
	CCValidateCert                         // 22
	CCUnknown           = 255
)

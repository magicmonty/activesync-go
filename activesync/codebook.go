package activesync

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func MakeCodeBook(ProtocolVersion string) *CodeBook {
	result := NewCodeBook()

	result.AddTagCodePage(makeAirSyncCodePage())
	result.AddTagCodePage(makeEmailCodePage())
	result.AddTagCodePage(makeContactsCodePage())
	result.AddTagCodePage(makeAirNotifyCodePage())
	result.AddTagCodePage(makeCalendarCodePage())
	result.AddTagCodePage(makeMoveCodePage())
	result.AddTagCodePage(makeItemEstimateCodePage())
	result.AddTagCodePage(makeFolderHierarchyCodePage())
	result.AddTagCodePage(makeMeetingResponseCodePage())
	result.AddTagCodePage(makeTasksCodePage())
	result.AddTagCodePage(makeResolveRecipientsCodePage())
	result.AddTagCodePage(makeValidateCertCodePage())
	result.AddTagCodePage(makeContacts2CodePage())
	result.AddTagCodePage(makePingCodePage())
	result.AddTagCodePage(makeProvisionCodePage())
	result.AddTagCodePage(makeSearchCodePage())
	result.AddTagCodePage(makeGALCodePage())
	result.AddTagCodePage(makeAirSyncBaseCodePage())
	result.AddTagCodePage(makeSettingsCodePage())
	result.AddTagCodePage(makeDocumentLibraryCodePage())
	result.AddTagCodePage(makeItemOperationsCodePage())
	result.AddTagCodePage(makeComposeMailCodePage())
	result.AddTagCodePage(makeEmail2CodePage())
	result.AddTagCodePage(makeNotesCodePage())
	result.AddTagCodePage(makeRightsManagementCodePage())

	return result
}

func makeAirSyncCodePage() CodePage {
	result := NewCodePage("AirSync", 0)

	result.AddTag("Sync", 0x05)
	result.AddTag("Responses", 0x06)
	result.AddTag("Add", 0x07)
	result.AddTag("Change", 0x08)
	result.AddTag("Delete", 0x09)
	result.AddTag("Fetch", 0x0A)
	result.AddTag("SyncKey", 0x0B)
	result.AddTag("ClientId", 0x0C)
	result.AddTag("ServerId", 0x0D)
	result.AddTag("Status", 0x0E)
	result.AddTag("Collection", 0x0F)
	result.AddTag("Class", 0x10)
	result.AddTag("CollectionId", 0x12)
	result.AddTag("GetChanges", 0x13)
	result.AddTag("MoreAvailable", 0x14)
	result.AddTag("WindowSize", 0x15)
	result.AddTag("Commands", 0x16)
	result.AddTag("Options", 0x17)
	result.AddTag("FilterType", 0x18)
	result.AddTag("Conflict", 0x1B)
	result.AddTag("Collections", 0x1C)
	result.AddTag("ApplicationData", 0x1D)
	result.AddTag("DeletesAsMoves", 0x1E)
	result.AddTag("Supported", 0x20)
	result.AddTag("SoftDelete", 0x21)
	result.AddTag("MIMESupport", 0x22)
	result.AddTag("MIMETruncation", 0x23)
	result.AddTag("Wait", 0x24)
	result.AddTag("Limit", 0x25)
	result.AddTag("Partial", 0x26)
	result.AddTag("ConversationMode", 0x27)  // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("MaxItems", 0x28)          // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("HeartbeatInterval", 0x29) // not supported with MS-ASProtocolVersion 12.1

	return result
}

func makeContactsCodePage() CodePage {
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
	result.AddTag("Alias", 0x3D)        // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("WeightedRank", 0x3E) // not supported with MS-ASProtocolVersion 12.1

	return result
}

func makeEmailCodePage() CodePage {
	result := NewCodePage("Email", 2)

	result.AddTag("DateReceived", 0x0F)
	result.AddTag("DisplayTo", 0x11)
	result.AddTag("Importance", 0x12)
	result.AddTag("MessageClass", 0x13)
	result.AddTag("Subject", 0x14)
	result.AddTag("Read", 0x15)
	result.AddTag("To", 0x16)
	result.AddTag("Cc", 0x17)
	result.AddTag("From", 0x18)
	result.AddTag("ReplyTo", 0x19)
	result.AddTag("AllDayEvent", 0x1A)
	result.AddTag("Categories", 0x1B) // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("Category", 0x1C)   // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("DTStamp", 0x1D)
	result.AddTag("EndTime", 0x1E)
	result.AddTag("InstanceType", 0x1F)
	result.AddTag("BusyStatus", 0x20)
	result.AddTag("Location", 0x21)
	result.AddTag("MeetingRequest", 0x22)
	result.AddTag("Organizer", 0x23)
	result.AddTag("RecurrenceId", 0x24)
	result.AddTag("Reminder", 0x25)
	result.AddTag("ResponseRequested", 0x26)
	result.AddTag("Recurrences", 0x27)
	result.AddTag("Recurrence", 0x28)
	result.AddTag("Recurrence_Type", 0x29)
	result.AddTag("Recurrence_Until", 0x2A)
	result.AddTag("Recurrence_Occurrences", 0x2B)
	result.AddTag("Recurrence_Interval", 0x2C)
	result.AddTag("Recurrence_DayOfWeek", 0x2D)
	result.AddTag("Recurrence_DayOfMonth", 0x2E)
	result.AddTag("Recurrence_WeekOfMonth", 0x2F)
	result.AddTag("Recurrence_MonthOfYear", 0x30)
	result.AddTag("StartTime", 0x31)
	result.AddTag("Sensitivity", 0x32)
	result.AddTag("TimeZone", 0x33)
	result.AddTag("GlobalObjId", 0x34)
	result.AddTag("ThreadTopic", 0x35)
	result.AddTag("InternetCPID", 0x39)
	result.AddTag("Flag", 0x3A)
	result.AddTag("Status", 0x3B)
	result.AddTag("ContentClass", 0x3C)
	result.AddTag("FlagType", 0x3D)
	result.AddTag("CompleteTime", 0x3E)
	result.AddTag("DisallowNewTimeProposal", 0x3F) // not supported with MS-ASProtocolVersion 12.1

	return result
}

// not used anymore
func makeAirNotifyCodePage() CodePage {
	result := NewCodePage("AirNotify", 3)
	return result
}

func makeCalendarCodePage() CodePage {
	result := NewCodePage("Calendar", 4)

	result.AddTag("TimeZone", 0x05)
	result.AddTag("AllDayEvent", 0x06)
	result.AddTag("Attendees", 0x07)
	result.AddTag("Attendee", 0x08)
	result.AddTag("Email", 0x09)
	result.AddTag("Name", 0x0A)
	result.AddTag("BusyStatus", 0x0D)
	result.AddTag("Categories", 0x0E)
	result.AddTag("Category", 0x0F)
	result.AddTag("DtStamp", 0x11)
	result.AddTag("EndTime", 0x12)
	result.AddTag("Exception", 0x13)
	result.AddTag("Exceptions", 0x14)
	result.AddTag("Deleted", 0x15)
	result.AddTag("ExceptionStartTime", 0x16)
	result.AddTag("Location", 0x17)
	result.AddTag("MeetingStatus", 0x18)
	result.AddTag("OrganizerEmail", 0x19)
	result.AddTag("OrganizerName", 0x1A)
	result.AddTag("Recurrence", 0x1B)
	result.AddTag("Type", 0x1C)
	result.AddTag("Until", 0x1D)
	result.AddTag("Occurrences", 0x1E)
	result.AddTag("Interval", 0x1F)
	result.AddTag("DayOfWeek", 0x20)
	result.AddTag("DayOfMonth", 0x21)
	result.AddTag("WeekOfMonth", 0x22)
	result.AddTag("MonthOfYear", 0x23)
	result.AddTag("Reminder", 0x24)
	result.AddTag("Sensitivity", 0x25)
	result.AddTag("Subject", 0x26)
	result.AddTag("StartTime", 0x27)
	result.AddTag("UID", 0x28)
	result.AddTag("AttendeeStatus", 0x29)
	result.AddTag("AttendeeType", 0x2A)
	result.AddTag("DisallowNewTimeProposal", 0x33)   // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("ResponseRequested", 0x34)         // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("AppointmentReplyTime", 0x35)      // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("ResponseType", 0x36)              // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("CalendarType", 0x37)              // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("IsLeapMonth", 0x38)               // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("FirstDayOfWeek", 0x39)            // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("OnlineMeetingConfLink", 0x3A)     // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("OnlineMeetingExternalLink", 0x3B) // not supported with MS-ASProtocolVersion 12.1 or 14.0

	return result
}

func makeMoveCodePage() CodePage {
	result := NewCodePage("Move", 5)

	result.AddTag("MoveItems", 0x05)
	result.AddTag("Move", 0x06)
	result.AddTag("SrcMsgId", 0x07)
	result.AddTag("SrcFldId", 0x08)
	result.AddTag("DstFldId", 0x09)
	result.AddTag("Response", 0x0A)
	result.AddTag("Status", 0x0B)
	result.AddTag("DstMsgId", 0x0C)

	return result
}

func makeItemEstimateCodePage() CodePage {
	result := NewCodePage("ItemEstimate", 6)

	result.AddTag("GetItemEstimate", 0x05)
	result.AddTag("Version", 0x06) // only supported with MS-ASProtocolVersion 12.1
	result.AddTag("Collections", 0x07)
	result.AddTag("Collection", 0x08)
	result.AddTag("Class", 0x09) // only supported with MS-ASProtocolVersion 12.1
	// with 14.0 and 14.1 Tha Class tag from CodePage 0 (ActiveSync) is used
	result.AddTag("CollectionId", 0x0A)
	result.AddTag("DateTime", 0x0B) // only supported with MS-ASProtocolVersion 12.1
	result.AddTag("Estimate", 0x0C)
	result.AddTag("Response", 0x0D)
	result.AddTag("Status", 0x0E)

	return result
}

func makeFolderHierarchyCodePage() CodePage {
	result := NewCodePage("FolderHierarchy", 7)

	result.AddTag("DisplayName", 0x07)
	result.AddTag("ServerId", 0x08)
	result.AddTag("ParentId", 0x09)
	result.AddTag("Type", 0x0A)
	result.AddTag("Status", 0x0C)
	result.AddTag("Changes", 0x0E)
	result.AddTag("Add", 0x0F)
	result.AddTag("Delete", 0x10)
	result.AddTag("Update", 0x11)
	result.AddTag("SyncKey", 0x12)
	result.AddTag("FolderCreate", 0x13)
	result.AddTag("FolderDelete", 0x14)
	result.AddTag("FolderUpdate", 0x15)
	result.AddTag("FolderSync", 0x16)
	result.AddTag("Count", 0x17)

	return result
}

func makeMeetingResponseCodePage() CodePage {
	result := NewCodePage("MeetingResponse", 8)

	result.AddTag("CalendarId", 0x05)
	result.AddTag("CollectionId", 0x06)
	result.AddTag("MeetingResponse", 0x07)
	result.AddTag("RequestId", 0x08)
	result.AddTag("Request", 0x09)
	result.AddTag("Result", 0x0A)
	result.AddTag("Status", 0x0B)
	result.AddTag("UserResponse", 0x0C)
	result.AddTag("InstanceId", 0x0E)

	return result
}

func makeTasksCodePage() CodePage {
	result := NewCodePage("Tasks", 9)

	result.AddTag("Categories", 0x08)
	result.AddTag("Category", 0x09)
	result.AddTag("Complete", 0x0A)
	result.AddTag("DateCompleted", 0x0B)
	result.AddTag("DueDate", 0x0C)
	result.AddTag("UtcDueDate", 0x0D)
	result.AddTag("Importance", 0x0E)
	result.AddTag("Recurrence", 0x0F)
	result.AddTag("Recurrence_Type", 0x10)
	result.AddTag("Recurrence_Start", 0x11)
	result.AddTag("Recurrence_Until", 0x12)
	result.AddTag("Recurrence_Occurrences", 0x13)
	result.AddTag("Recurrence_Interval", 0x14)
	result.AddTag("Recurrence_DayOfMonth", 0x15)
	result.AddTag("Recurrence_DayOfWeek", 0x16)
	result.AddTag("Recurrence_WeekOfMonth", 0x17)
	result.AddTag("Recurrence_MonthOfYear", 0x18)
	result.AddTag("Recurrence_Regenerate", 0x19)
	result.AddTag("Recurrence_DeadOccur", 0x1A)
	result.AddTag("ReminderSet", 0x1B)
	result.AddTag("ReminderTime", 0x1C)
	result.AddTag("Sensitivity", 0x1D)
	result.AddTag("StartDate", 0x1E)
	result.AddTag("UtcStartDate", 0x1F)
	result.AddTag("Subject", 0x20)
	result.AddTag("OrdinalDate", 0x22)
	result.AddTag("SubOrdinalDate", 0x23)
	result.AddTag("CalendarType", 0x24)   // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("IsLeapMonth", 0x25)    // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("FirstDayOfWeek", 0x26) // not supported with MS-ASProtocolVersion 12.1 or 14.0

	return result
}

func makeResolveRecipientsCodePage() CodePage {
	result := NewCodePage("ResolveRecipients", 10)

	result.AddTag("ResolveRecipients", 0x05)
	result.AddTag("Response", 0x06)
	result.AddTag("Status", 0x07)
	result.AddTag("Type", 0x08)
	result.AddTag("Recipient", 0x09)
	result.AddTag("DisplayName", 0x0A)
	result.AddTag("EmailAddress", 0x0B)
	result.AddTag("Certificates", 0x0C)
	result.AddTag("Certificate", 0x0D)
	result.AddTag("MiniCertificate", 0x0E)
	result.AddTag("Options", 0x0F)
	result.AddTag("To", 0x10)
	result.AddTag("CertificateRetrieval", 0x11)
	result.AddTag("RecipientCount", 0x12)
	result.AddTag("MaxCertificates", 0x13)
	result.AddTag("MaxAmbiguousRecipients", 0x14)
	result.AddTag("CertificateCount", 0x15)
	result.AddTag("Availability", 0x16)   // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("StartTime", 0x17)      // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("EndTime", 0x18)        // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("MergedFreeBusy", 0x19) // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("Picture", 0x1A)        // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("MaxSize", 0x1B)        // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("Data", 0x1C)           // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("MaxPictures", 0x1D)    // not supported with MS-ASProtocolVersion 12.1 or 14.0

	return result
}

func makeValidateCertCodePage() CodePage {
	result := NewCodePage("ValidateCert", 11)

	result.AddTag("ValidateCert", 0x05)
	result.AddTag("Certificates", 0x06)
	result.AddTag("Certificate", 0x07)
	result.AddTag("CertificateChain", 0x08)
	result.AddTag("CheckCRL", 0x09)
	result.AddTag("Status", 0x0A)

	return result
}

func makeContacts2CodePage() CodePage {
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

func makePingCodePage() CodePage {
	result := NewCodePage("Ping", 13)

	result.AddTag("Ping", 0x05)
	result.AddTag("AutdState", 0x06) // not used
	result.AddTag("Status", 0x07)
	result.AddTag("HeartbeatInterval", 0x08)
	result.AddTag("Folders", 0x09)
	result.AddTag("Folder", 0x0A)
	result.AddTag("Id", 0x0B)
	result.AddTag("Class", 0x0C)
	result.AddTag("MaxFolders", 0x0D)

	return result
}

func makeProvisionCodePage() CodePage {
	result := NewCodePage("Provision", 14)

	result.AddTag("Provision", 0x05)
	result.AddTag("Policies", 0x06)
	result.AddTag("Policy", 0x07)
	result.AddTag("PolicyType", 0x08)
	result.AddTag("PolicyKey", 0x09)
	result.AddTag("Data", 0x0A)
	result.AddTag("Status", 0x0B)
	result.AddTag("RemoteWipe", 0x0C)
	result.AddTag("EASProvisionDoc", 0x0D)
	result.AddTag("DevicePasswordEnabled", 0x0E)
	result.AddTag("AlphanumericDevicePasswordRequired", 0x0F)
	result.AddTag("DeviceEncryptionEnabled", 0x10)
	result.AddTag("RequireStorageCardEncryption", 0x10) // equivalent to DeviceEncryptionEnabled
	result.AddTag("PasswordRecoveryEnabled", 0x11)
	result.AddTag("AttachmentsEnabled", 0x13)
	result.AddTag("MinDevicePasswordLength", 0x14)
	result.AddTag("MaxInactivityTimeDeviceLock", 0x15)
	result.AddTag("MaxDevicePasswordFailedAttempts", 0x16)
	result.AddTag("MaxAttachmentSize", 0x17)
	result.AddTag("AllowSimpleDevicePassword", 0x18)
	result.AddTag("DevicePasswordExpiration", 0x19)
	result.AddTag("DevicePasswordHistory", 0x1A)
	result.AddTag("AllowStorageCard", 0x1B)
	result.AddTag("AllowCamera", 0x1C)
	result.AddTag("RequireDeviceEncryption", 0x1D)
	result.AddTag("AllowUnsignedApplications", 0x1E)
	result.AddTag("AllowUnsignedInstallationPackages", 0x1F)
	result.AddTag("MinDevicePasswordComplexCharacters", 0x20)
	result.AddTag("AllowWiFi", 0x21)
	result.AddTag("AllowTextMessaging", 0x22)
	result.AddTag("AllowPOPIMAPEmail", 0x23)
	result.AddTag("AllowBluetooth", 0x24)
	result.AddTag("AllowIrDA", 0x25)
	result.AddTag("RequireManualSyncWhenRoaming", 0x26)
	result.AddTag("AllowDesktopSync", 0x27)
	result.AddTag("MaxCalendarAgeFilter", 0x28)
	result.AddTag("AllowHTMLEmail", 0x29)
	result.AddTag("MaxEmailAgeFilter", 0x2A)
	result.AddTag("MaxEmailBodyTruncationSize", 0x2B)
	result.AddTag("MaxEmailHTMLBodyTruncationSize", 0x2C)
	result.AddTag("RequireSignedSMIMEMessages", 0x2D)
	result.AddTag("RequireEncryptedSMIMEMessages", 0x2E)
	result.AddTag("RequireSignedSMIMEAlgorithm", 0x2F)
	result.AddTag("RequireEncryptionSMIMEAlgorithm", 0x30)
	result.AddTag("AllowSMIMEEncryptionAlgorithmNegotiation", 0x31)
	result.AddTag("AllowSMIMESoftCerts", 0x32)
	result.AddTag("AllowBrowser", 0x33)
	result.AddTag("AllowConsumerEmail", 0x34)
	result.AddTag("AllowRemoteDesktop", 0x35)
	result.AddTag("AllowInternetSharing", 0x36)
	result.AddTag("UnapprovedInROMApplicationList", 0x37)
	result.AddTag("ApplicationName", 0x38)
	result.AddTag("ApprovedApplicationList", 0x39)
	result.AddTag("Hash", 0x3A)

	return result
}

func makeSearchCodePage() CodePage {
	result := NewCodePage("Search", 15)

	result.AddTag("Search", 0x05)
	result.AddTag("Store", 0x07)
	result.AddTag("Name", 0x08)
	result.AddTag("Query", 0x09)
	result.AddTag("Options", 0x0A)
	result.AddTag("Range", 0x0B)
	result.AddTag("Status", 0x0C)
	result.AddTag("Response", 0x0D)
	result.AddTag("Result", 0x0E)
	result.AddTag("Properties", 0x0F)
	result.AddTag("Total", 0x10)
	result.AddTag("EqualTo", 0x11)
	result.AddTag("Value", 0x12)
	result.AddTag("And", 0x13)
	result.AddTag("Or", 0x14) // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("FreeText", 0x15)
	result.AddTag("DeepTraversal", 0x17)
	result.AddTag("LongId", 0x18)
	result.AddTag("RebuildResults", 0x19)
	result.AddTag("LessThan", 0x1A)
	result.AddTag("GreaterThan", 0x1B)
	result.AddTag("UserName", 0x1E)
	result.AddTag("Password", 0x1F)
	result.AddTag("ConversationId", 0x20)
	result.AddTag("Picture", 0x21)     // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("MaxSize", 0x22)     // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("MaxPictures", 0x23) // not supported with MS-ASProtocolVersion 12.1 or 14.0

	return result
}

func makeGALCodePage() CodePage {
	result := NewCodePage("GAL", 16)

	result.AddTag("DisplayName", 0x05)
	result.AddTag("Phone", 0x06)
	result.AddTag("Office", 0x07)
	result.AddTag("Title", 0x08)
	result.AddTag("Company", 0x09)
	result.AddTag("Alias", 0x0A)
	result.AddTag("FirstName", 0x0B)
	result.AddTag("LastName", 0x0C)
	result.AddTag("HomePhone", 0x0D)
	result.AddTag("MobilePhone", 0x0E)
	result.AddTag("EmailAddress", 0x0F)
	result.AddTag("Picture", 0x10) // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("Status", 0x11)  // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("Data", 0x12)    // not supported with MS-ASProtocolVersion 12.1 or 14.0

	return result
}

func makeAirSyncBaseCodePage() CodePage {
	result := NewCodePage("AirSyncBase", 17)

	result.AddTag("BodyPreference", 0x05)
	result.AddTag("Type", 0x06)
	result.AddTag("TruncationSize", 0x07)
	result.AddTag("AllOrNone", 0x08)
	result.AddTag("Body", 0x0A)
	result.AddTag("Data", 0x0B)
	result.AddTag("EstimatedDataSize", 0x0C)
	result.AddTag("Truncated", 0x0D)
	result.AddTag("Attachments", 0x0E)
	result.AddTag("Attachment", 0x0F)
	result.AddTag("DisplayName", 0x10)
	result.AddTag("FileReference", 0x11)
	result.AddTag("Method", 0x12)
	result.AddTag("ContentId", 0x13)
	result.AddTag("ContentLocation", 0x14) // not used
	result.AddTag("IsInline", 0x15)
	result.AddTag("NativeBodyType", 0x16)
	result.AddTag("ContentType", 0x17)
	result.AddTag("Preview", 0x18)            // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("BodyPartPreference", 0x19) // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("BodyPart", 0x1A)           // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("Status", 0x1B)             // not supported with MS-ASProtocolVersion 12.1 or 14.0

	return result
}

func makeSettingsCodePage() CodePage {
	result := NewCodePage("Settings", 18)

	result.AddTag("Settings", 0x05)
	result.AddTag("Status", 0x06)
	result.AddTag("Get", 0x07)
	result.AddTag("Set", 0x08)
	result.AddTag("Oof", 0x09)
	result.AddTag("OofState", 0x0A)
	result.AddTag("StartTime", 0x0B)
	result.AddTag("EndTime", 0x0C)
	result.AddTag("OofMessage", 0x0D)
	result.AddTag("AppliesToInternal", 0x0E)
	result.AddTag("AppliesToExternalKnown", 0x0F)
	result.AddTag("AppliesToExternalUnknown", 0x10)
	result.AddTag("Enabled", 0x11)
	result.AddTag("ReplyMessage", 0x12)
	result.AddTag("BodyType", 0x13)
	result.AddTag("DevicePassword", 0x14)
	result.AddTag("Password", 0x15)
	result.AddTag("DeviceInformation", 0x16)
	result.AddTag("Model", 0x17)
	result.AddTag("IMEI", 0x18)
	result.AddTag("FriendlyName", 0x19)
	result.AddTag("OS", 0x1A)
	result.AddTag("OSLanguage", 0x1B)
	result.AddTag("PhoneNumber", 0x1C)
	result.AddTag("UserInformation", 0x1D)
	result.AddTag("EmailAddresses", 0x1E)
	result.AddTag("SmtpAddress", 0x1F)
	result.AddTag("UserAgent", 0x20)
	result.AddTag("EnableOutboundSMS", 0x21)           // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("MobileOperator", 0x22)              // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("PrimarySmtpAddress", 0x23)          // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("Accounts", 0x24)                    // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("Account", 0x25)                     // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("AccountId", 0x26)                   // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("AccountName", 0x27)                 // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("UserDisplayName", 0x28)             // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("SendDisabled", 0x29)                // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("RightsManagementInformation", 0x2B) // not supported with MS-ASProtocolVersion 12.1 or 14.0

	return result
}

func makeDocumentLibraryCodePage() CodePage {
	result := NewCodePage("DocumentLibrary", 19)

	result.AddTag("LinkId", 0x05)
	result.AddTag("DisplayName", 0x06)
	result.AddTag("IsFolder", 0x07)
	result.AddTag("CreationDate", 0x08)
	result.AddTag("LastModifiedDate", 0x09)
	result.AddTag("IsHidden", 0x0A)
	result.AddTag("ContentLength", 0x0B)
	result.AddTag("ContentType", 0x0C)

	return result
}

func makeItemOperationsCodePage() CodePage {
	result := NewCodePage("ItemOperations", 20)

	result.AddTag("ItemOperations", 0x05)
	result.AddTag("Fetch", 0x06)
	result.AddTag("Store", 0x07)
	result.AddTag("Options", 0x08)
	result.AddTag("Range", 0x09)
	result.AddTag("Total", 0x0A)
	result.AddTag("Properties", 0x0B)
	result.AddTag("Data", 0x0C)
	result.AddTag("Status", 0x0D)
	result.AddTag("Response", 0x0E)
	result.AddTag("Version", 0x0F)
	result.AddTag("Schema", 0x10)
	result.AddTag("Part", 0x11)
	result.AddTag("EmptyFolderContents", 0x12)
	result.AddTag("DeleteSubFolders", 0x13)
	result.AddTag("UserName", 0x14)
	result.AddTag("Password", 0x15)
	result.AddTag("Move", 0x16)           // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("DstFldId", 0x17)       // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("ConversationId", 0x18) // not supported with MS-ASProtocolVersion 12.1
	result.AddTag("MoveAlways", 0x19)     // not supported with MS-ASProtocolVersion 12.1

	return result
}

// not supported with MS-ASProtocolVersion 12.1
func makeComposeMailCodePage() CodePage {
	result := NewCodePage("ComposeMail", 21)

	result.AddTag("SendMail", 0x05)
	result.AddTag("SmartForward", 0x06)
	result.AddTag("SmartReply", 0x07)
	result.AddTag("SaveInSentItems", 0x08)
	result.AddTag("ReplaceMime", 0x09)
	result.AddTag("Source", 0x0B)
	result.AddTag("FolderId", 0x0C)
	result.AddTag("ItemId", 0x0D)
	result.AddTag("LongId", 0x0E)
	result.AddTag("InstanceId", 0x0F)
	result.AddTag("Mime", 0x10)
	result.AddTag("ClientId", 0x11)
	result.AddTag("Status", 0x12)
	result.AddTag("AccountId", 0x13) // not supported with MS-ASProtocolVersion 12.1 or 14.0

	return result
}

// not supported with MS-ASProtocolVersion 12.1
func makeEmail2CodePage() CodePage {
	result := NewCodePage("Email2", 22)

	result.AddTag("UmCallerID", 0x05)
	result.AddTag("UmUserNotes", 0x06)
	result.AddTag("UmAttDuration", 0x07)
	result.AddTag("UmAttOrder", 0x08)
	result.AddTag("ConversationId", 0x09)
	result.AddTag("ConversationIndex", 0x0A)
	result.AddTag("LastVerbExecuted", 0x0B)
	result.AddTag("LastVerbExecutionTime", 0x0C)
	result.AddTag("ReceivedAsBcc", 0x0D)
	result.AddTag("Sender", 0x0E)
	result.AddTag("CalendarType", 0x0F)
	result.AddTag("IsLeapMonth", 0x10)
	result.AddTag("AccountId", 0x11)          // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("FirstDayOfWeek", 0x12)     // not supported with MS-ASProtocolVersion 12.1 or 14.0
	result.AddTag("MeetingMessageType", 0x13) // not supported with MS-ASProtocolVersion 12.1 or 14.0

	return result
}

// not supported with MS-ASProtocolVersion 12.1
func makeNotesCodePage() CodePage {
	result := NewCodePage("Notes", 23)

	result.AddTag("Subject", 0x05)
	result.AddTag("MessageClass", 0x06)
	result.AddTag("LastModifiedDate", 0x07)
	result.AddTag("Categories", 0x08)
	result.AddTag("Category", 0x09)

	return result
}

// not supported with MS-ASProtocolVersion 12.1 or 14.1
func makeRightsManagementCodePage() CodePage {
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

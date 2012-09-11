package activesync

import (
	"bytes"
	"fmt"
	. "github.com/magicmonty/wbxml-go/wbxml"
)

func makeDataBuffer(data ...byte) *bytes.Buffer {
	return bytes.NewBuffer(data)
}

func getDecodeResult(data ...byte) string {
	var result string
	result, _ = Decode(bytes.NewBuffer(data), MakeCodeBook())
	return result
}

// See http://msdn.microsoft.com/en-us/library/ee237245(v=exchg.80)
func ExampleMSAsWbxml() {
	fmt.Println(
		getDecodeResult(
			WBXML_1_3, UNKNOWN_PI, CHARSET_UTF8, 0x00,
			0x05|TAG_HAS_CONTENT,                                                           // <AirSync:Sync>
			0x1C|TAG_HAS_CONTENT,                                                           // <AirSync:Collections>
			0x0F|TAG_HAS_CONTENT,                                                           // <AirSync:Collection>
			0x10|TAG_HAS_CONTENT, STR_I, 'c', 'o', 'n', 't', 'a', 'c', 't', 's', 0x00, END, // <AirSync:Class>contacts</AirSync:Class>
			0x0B|TAG_HAS_CONTENT, STR_I, '2', 0x00, END, // <AirSync:SyncKey>2</AirSync:SyncKey>
			0x12|TAG_HAS_CONTENT, STR_I, '2', 0x00, END, // <AirSync:CollectionId>2</AirSync:CollectionId>
			0x0E|TAG_HAS_CONTENT, STR_I, '1', 0x00, END, // <AirSync:Status>1</AirSync:Status>
			0x16|TAG_HAS_CONTENT,                                  // <AirSync:Commands>
			0x07|TAG_HAS_CONTENT,                                  // <AirSync:Add>
			0x0D|TAG_HAS_CONTENT, STR_I, '2', ':', '1', 0x00, END, // <AirSync:ServerId>2:1</AirSync:ServerId>
			0x1D|TAG_HAS_CONTENT, // <AirSync:ApplicationData>
			SWITCH_PAGE, 0x11,    // Select CodePage AirSyncBase
			0x0A|TAG_HAS_CONTENT,                        // <AirSyncBase:Body>
			0x06|TAG_HAS_CONTENT, STR_I, '1', 0x00, END, // <AirSyncBase:Type>1</AirSyncBase:Type>
			0x0C|TAG_HAS_CONTENT, STR_I, '0', 0x00, END, // <AirSyncBase:EstimatedDataSize>0</AirSyncBase:EstimatedDataSize>
			0x0D|TAG_HAS_CONTENT, STR_I, '1', 0x00, END, // <AirSyncBase:Truncated>1</AirSyncBase:Truncated>
			END,               // </AirSyncBase:Body>
			SWITCH_PAGE, 0x01, // Select CodePage Contacts
			0x1E|TAG_HAS_CONTENT, STR_I, 'F', 'u', 'n', 'k', ',', ' ', 'D', 'o', 'n', 0x00, END, // <Contacts:FileAs>Funk, Don</Contacts:FileAs>
			0x1F|TAG_HAS_CONTENT, STR_I, 'D', 'o', 'n', 0x00, END, // <Contacts:FirstName>Don</Contacts:FirstName>
			0x29|TAG_HAS_CONTENT, STR_I, 'F', 'u', 'n', 'k', 0x00, END, // <Contacts:LastName>Funk</Contacts:LastName>
			SWITCH_PAGE, 0x11, // Select CodePage AirSyncBase
			0x16|TAG_HAS_CONTENT, STR_I, '1', 0x00, END, // <AirSyncBase:NativeBodyType>1</AirSyncBase:NativeBodyType>
			END, // </AirSync:ApplicationData>
			END, // </AirSync:Add>
			END, // </AirSync:Commands>
			END, // </AirSync:Collection>
			END, // </AirSync:Collections>
			END, // </AirSync:Sync>
		))
	// OUTPUT: <?xml version="1.0" encoding="utf-8"?>
	// <Sync xmlns="AirSync" xmlns:B="Contacts" xmlns:R="AirSyncBase"><Collections><Collection><Class>contacts</Class><SyncKey>2</SyncKey><CollectionId>2</CollectionId><Status>1</Status><Commands><Add><ServerId>2:1</ServerId><ApplicationData><R:Body><R:Type>1</R:Type><R:EstimatedDataSize>0</R:EstimatedDataSize><R:Truncated>1</R:Truncated></R:Body><B:FileAs>Funk, Don</B:FileAs><B:FirstName>Don</B:FirstName><B:LastName>Funk</B:LastName><R:NativeBodyType>1</R:NativeBodyType></ApplicationData></Add></Commands></Collection></Collections></Sync>
}

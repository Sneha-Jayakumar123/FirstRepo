// package main

// import (
// 	// "bytes"
// 	// "encoding/xml"
// 	"encoding/xml"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	xmlFilePath := "mule.xml" // Replace with your XML file path

// 	xmlFile, err := os.Open(xmlFilePath)
// 	if err != nil {
// 		fmt.Println("Error opening XML file:", err)
// 		return
// 	}
// 	defer xmlFile.Close()

// 	// Create a buffer to read the file content
// 	buf := make([]byte, 1024) // You can adjust the buffer size as needed

// 	var xmlString string

// 	for {
// 		n, err := xmlFile.Read(buf)
// 		if err != nil {
// 			if err.Error() == "EOF" {
// 				// Reached the end of the file, break the loop
// 				break
// 			}
// 			fmt.Println("Error reading XML file:", err)
// 			break
// 		}

// 		// Print the content read in this iteration
// 		fmt.Print(string(buf[:n]))
// 		xmlString = string(buf[:n])
// 	}

// 	// Your XML string
// 	// xmlString := `<root><element1>Value1</element1><element2>Value2</element2></root>`

// 	// Define a struct for parsing the XML
// 	type Root struct {
// 		XMLName xml.Name
// 	}

// 	var data Root
// 	err = xml.Unmarshal([]byte(xmlString), &data)
// 	if err != nil {
// 		fmt.Printf("Error unmarshaling XML: %v\n", err)
// 		return
// 	}

// 	// Print the XML tags
// 	printXMLTags(data.XMLName, xmlString)
// 	//fmt.Println(`<?xml profile><test>` + Xml(`test '123'`) + `</test>`)
// }

// func Xml(in string) string {
// 	var b bytes.Buffer
// 	xml.EscapeText(&b, []byte(in))
// 	return b.String()
// }

// ------------------------------------------------------------------------------------------------------
// package main

// import (
// 	"encoding/json"
// 	"encoding/xml"
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// Your XML string (the structure is unknown)
// 	xmlString := `<root><element1>Value1</element1><element2>Value2</element2></root>`

// 	// Unmarshal the XML into an interface{} to handle dynamic content
// 	// var data interface{}
// 	type Root struct {
// 		XMLName xml.Name `xml:"root"`
// 	}

// 	var data Root
// 	err := xml.Unmarshal([]byte(xmlString), &data)
// 	if err != nil {
// 		fmt.Printf("Error unmarshaling XML: %v\n", err)
// 		return
// 	}

// 	// Marshal the interface into JSON
// 	_, err = json.Marshal(data)
// 	if err != nil {
// 		fmt.Printf("Error marshaling JSON: %v\n", err)
// 		return
// 	}

// 	// Print the JSON string
// 	// fmt.Println(string(jsonData))
// 	printXMLTags(data.XMLName, xmlString)
// }

//-------------------------------------------------------------------------------------------------
// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// )

// func main() {
// 	// Your XML string
// 	xmlString := `<root><element1>Value1</element1><element2>Value2</element2></root>`

// 	// Define a struct that matches the structure of the XML
// 	type Root struct {
// 		Element1 string `xml:"element1"`
// 		Element2 string `xml:"element2"`
// 	}

// 	var data Root
// 	err := xml.Unmarshal([]byte(xmlString), &data)
// 	if err != nil {
// 		fmt.Printf("Error unmarshaling XML: %v\n", err)
// 		return
// 	}

// 	// Now you can access the values based on the tags
// 	fmt.Printf("Element1: %s\n", data.Element1)
// 	fmt.Printf("Element2: %s\n", data.Element2)
// }
//--------------------------------------------------------------------------------------------

// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// Your XML string
// 	xmlString := `<root><element1>Value1</element1><element2>Value2</element2></root>`

// 	// Define a struct for parsing the XML
// 	type Root struct {
// 		XMLName xml.Name `xml:"root"`
// 	}

// 	var data Root
// 	err := xml.Unmarshal([]byte(xmlString), &data)
// 	if err != nil {
// 		fmt.Printf("Error unmarshaling XML: %v\n", err)
// 		return
// 	}

// 	// Print the XML tags
// 	printXMLTags(data.XMLName, xmlString)
// }

// func printXMLTags(name xml.Name, xmlString string) {
// 	fmt.Println("Root Element:", name.Local)
// 	startArray :=[]string{}
// 	endArray:=[]string{}
// 	decoder := xml.NewDecoder(strings.NewReader(xmlString))
// 	for {
// 		token, err := decoder.Token()
// 		if err != nil {
// 			break
// 		}
// 		switch t := token.(type) {
// 		case xml.StartElement:{
// 			startArray=append(startArray,t.Name.Local)
// 			fmt.Println("Start Tag:",t.Name.Local )
// 		}
// 		case xml.EndElement:
// 			endArray=append(endArray,t.Name.Local)
// 			fmt.Println("End Tag:",t.Name.Local)
// 		}
// 	}
// 	fmt.Printf("Start Tag:%+v \n", startArray)
// 	fmt.Printf("End Tag:%+v ",endArray)
// }

//----------------------------------------------------------------------
//----this code works in the only by writing the content of the one logic
// package main

// import (
//     "encoding/xml"
//     "fmt"
//     "strings"
// )

// func main() {
//     xmlData := `<mule>This is the code <mule>`

//     decoder := xml.NewDecoder(strings.NewReader(xmlData))
//     var result string
//     var inMuleTag bool

//     for {
//         token, err := decoder.Token()
//         if err != nil {
//             break
//         }

//         switch t:=token.(type) {
//         case xml.StartElement:
//             element := token.(xml.StartElement)
// 			fmt.Println("The T",t)
//             if element.Name.Local == "mule" {
//                 inMuleTag = true
//             }
//         case xml.EndElement:
//             element := token.(xml.EndElement)
//             if element.Name.Local == "mule" {
//                 inMuleTag = false
//             }
//         case xml.CharData:
//             if inMuleTag {
//                 result += string(token.(xml.CharData))
//             }
//         }
//     }

//     fmt.Println(result) // This is the code
// }
//-------------------------------------End of the code ------------------------------------------

// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	xmlData := `<root>
//     <mule>This is the mule code</mule>
//     <cat>This is the cat code</cat>
//     <dog>This is the dog code</dog>
// </root>`

// 	decoder := xml.NewDecoder(strings.NewReader(xmlData))
// 	tagContent := make(map[string]string)
// 	var currentTag string

// 	for {
// 		token, err := decoder.Token()
// 		if err != nil {
// 			break
// 		}

// 		switch t := token.(type) {
// 		case xml.StartElement:
// 			currentTag = t.Name.Local
// 		case xml.EndElement:
// 			currentTag = ""
// 		case xml.CharData:
// 			if currentTag != "" {
// 				tagContent[currentTag] += string(t)
// 			}
// 		}
// 	}

// 	for tag, content := range tagContent {
// 		fmt.Printf("%s: %s\n", tag, content)
// 	}
// }
//-------------This code works for the many tags but not with the attributes--------------------

// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	xmlData := `<?xml version="1.0" encoding="UTF-8"?>

// 	<mule xmlns="http://www.mulesoft.org/schema/mule/core"
// 		  xmlns:http="http://www.mulesoft.org/schema/mule/http"
// 		  xmlns:doc="http://www.mulesoft.org/schema/mule/documentation"
// 		  xmlns:spring="http://www.springframework.org/schema/beans"
// 		  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
// 		  xsi:schemaLocation="
// 			http://www.mulesoft.org/schema/mule/core http://www.mulesoft.org/schema/mule/core/current/mule.xsd
// 			http://www.mulesoft.org/schema/mule/http http://www.mulesoft.org/schema/mule/http/current/mule-http.xsd
// 			http://www.mulesoft.org/schema/mule/documentation http://www.mulesoft.org/schema/mule/documentation/current/mule-documentation.xsd
// 			http://www.springframework.org/schema/beans http://www.springframework.org/schema/beans/spring-beans-current.xsd">

// 		<http:listener-config name="HTTP_Listener_Config" host="0.0.0.0" port="8081" doc:name="HTTP Listener Configuration" />

// 		<flow name="simpleFlow">
// 			<http:listener config-ref="HTTP_Listener_Config" path="/example" doc:name="HTTP" />
// 			<set-payload value="Hello, this is a MuleSoft example!" doc:name="Set Payload" />
// 		</flow>
// 	</mule>.`

// 	decoder := xml.NewDecoder(strings.NewReader(xmlData))
// 	var result string
// 	var inElement bool
// 	var currentElement string

// 	for {
// 		token, err := decoder.Token()
// 		if err != nil {
// 			break
// 		}

// 		switch t := token.(type) {
// 		case xml.StartElement:
// 			element := token.(xml.StartElement)
// 			currentElement = element.Name.Local
// 			if len(element.Attr) > 0 {
// 				for _, attr := range element.Attr {
// 					fmt.Printf("%s: %s\n", attr.Name.Local, attr.Value)
// 				}
// 			}
// 			inElement = true
// 		case xml.EndElement:
// 			inElement = false
// 		case xml.CharData:
// 			if inElement {
// 				result += string(t)
// 			}
// 		}
// 	}

// 	fmt.Printf("%s: %s\n", currentElement, result)
// }
//-------------------------it is not working as expected----------------------

// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	xmlData := `<?xml version="1.0" encoding="UTF-8"?>

// 	<mule xmlns:db="http://www.mulesoft.org/schema/mule/db" xmlns:ee="http://www.mulesoft.org/schema/mule/ee/core"
// 		xmlns:compression="http://www.mulesoft.org/schema/mule/compression"
// 		xmlns:http="http://www.mulesoft.org/schema/mule/http" xmlns="http://www.mulesoft.org/schema/mule/core" xmlns:doc="http://www.mulesoft.org/schema/mule/documentation" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.mulesoft.org/schema/mule/core http://www.mulesoft.org/schema/mule/core/current/mule.xsd
// 	http://www.mulesoft.org/schema/mule/http http://www.mulesoft.org/schema/mule/http/current/mule-http.xsd
// 	http://www.mulesoft.org/schema/mule/compression http://www.mulesoft.org/schema/mule/compression/current/mule-compression.xsd
// 	http://www.mulesoft.org/schema/mule/ee/core http://www.mulesoft.org/schema/mule/ee/core/current/mule-ee.xsd
// 	http://www.mulesoft.org/schema/mule/db http://www.mulesoft.org/schema/mule/db/current/mule-db.xsd">
// 		<flow name="animeFlow" doc:id="8934cc76-3405-49bf-a626-da9ae53b4547" >
// 			<http:listener doc:name="Triggers Flow" doc:id="09e321e1-2d69-483b-8654-c05b12b8e67d" config-ref="HTTP_Listener_config" path="${HL.path}">
// 				<http:error-response statusCode="#[vars.status]" >
// 					<http:body ><![CDATA[#[output application/json --- payload]]]></http:body>
// 				</http:error-response>
// 			</http:listener>
// 			<set-variable value="#[payload.customerID]" doc:name="Set Variable" doc:id="af73ad48-b949-48e1-b64c-29031316e35b" variableName="customer_Id"/>
// 			<until-successful maxRetries="5" doc:name="Until Successful" doc:id="b1283b69-0067-4f23-bb99-f6036b334efe" millisBetweenRetries="50000">
// 				<try doc:name="Try" doc:id="2a8aaa75-fe2e-4728-a557-af78930f6be1" >
// 					<http:request method="GET" doc:name="HTTP: Get Anime API" doc:id="56950f58-6094-457c-b1d5-4aa5b094b7a0" url="${Request.URL}">
// 						<ee:repeatable-file-store-stream bufferUnit="MB" />
// 					</http:request>
// 				</try>
// 			</until-successful>
// 			<compression:decompress doc:name="Decompress:Data" doc:id="17e91826-ff45-45bc-8deb-12ffa227a9dd" outputMimeType="application/xml" outputEncoding="UTF-8">
// 				<compression:decompressor >
// 					<compression:gzip-decompressor />
// 				</compression:decompressor>
// 			</compression:decompress>
// 			<ee:transform doc:name="Build:converting XML to JSON" doc:id="0e294c24-5609-41a9-8ef3-7d5a9a939a7e" >
// 				<ee:message >
// 					<ee:set-payload ><![CDATA[%dw 2.0
// 	output application/json
// 	---
// 	payload]]></ee:set-payload>
// 				</ee:message>
// 				<ee:variables >
// 					<ee:set-variable variableName="response" ><![CDATA[%dw 2.0
// 	output application/json
// 	---
// 	payload]]></ee:set-variable>
// 				</ee:variables>
// 			</ee:transform>
// 			<logger level="INFO" doc:name="Logs payload" doc:id="90a11369-1dd2-4b4d-9587-39f61c239e88" message="#[payload]"/>
// 			<choice doc:name="Choice" doc:id="51d66274-de56-4cdd-b1a8-6bc14e9f03d2" >
// 				<when expression="#[vars.response.error=='banned' or vars.response.error=='Missing']">
// 					<raise-error doc:name="Raise error" doc:id="0f18a6f7-71cb-4b86-9731-b17926a20119" type="CUSTOM:UNKNOWN_ERROR"/>
// 				</when>
// 				<when expression="#[vars.response.error=='Not Found']">
// 					<raise-error doc:name="Raise error" doc:id="2a870867-5d26-4c43-8ec7-f0fdfb49cdcb" type="CUSTOM:NOT_FOUND"/>
// 				</when>
// 				<otherwise >
// 					<try doc:name="Try" doc:id="d9d85582-4115-48bd-a2b3-18dfca4d7d0a" >
// 						<db:select doc:name="Retrieve:Recommendation" doc:id="e3f22f19-a4e7-4993-a0af-8e2d7e26e40b" config-ref="Database_Config">
// 							<db:sql ><![CDATA[select MAX(recommendation) as recommendation from customer_detail where customerID=:customer_Id]]></db:sql>
// 							<db:input-parameters ><![CDATA[#[customer_Id:vars.customer_Id]]]></db:input-parameters>
// 						</db:select>
// 					</try>
// 					<set-variable value="#[payload[0].recommendation]" doc:name="Store:recommendation" doc:id="0c1eb81d-dc56-4321-83f1-affeaaa525d6" variableName="count"/>
// 					<ee:transform doc:name="Build:Manipulation On Data" doc:id="181aee64-3617-410c-944e-6755461c6237" >
// 						<ee:message >
// 							<ee:set-payload resource="Anime.dwl" />
// 						</ee:message>
// 					</ee:transform>
// 					<logger level="INFO" doc:name="Logger" doc:id="8a4a2614-0ecf-4489-96c0-65ac9b7a94f4" message="#[payload]"/>
// 					<try doc:name="Try" doc:id="b77fd3d6-0113-4328-99e4-21691434556a" >
// 						<db:insert doc:name="Insert:Data" doc:id="cfc334e6-ba4c-43cb-8791-d9f1f1acc2fa" config-ref="Database_Config">
// 							<db:sql ><![CDATA[insert into customer_detail(customerID,episodecount,startdate,enddate,Title,review_rating,tags,creators,recommendation,created_date)
// 	  values (:customerID,:episodecount,:startdate,:enddate,:Title,:review_rating,:tags,:creators,:recommendation,:created_date)]]></db:sql>
// 							<db:input-parameters ><![CDATA[#[{
// 		customerID:vars.customer_Id,
// 		episodecount:payload.episodecount,
// 		startdate:payload.startdate,
// 		enddate:payload.enddate,
// 		Title:payload.Title,
// 		review_rating:payload.review_rating,
// 		tags:payload.tags,
// 		creators:payload.creators,
// 		recommendation:payload.recommendation,
// 		created_date:payload.created_date
// 	}]]]></db:input-parameters>
// 						</db:insert>
// 					</try>
// 					<set-payload value='#[{"output":"The data are inserted successfully",&#10;	"customerID":vars.customer_Id&#10;}]' doc:name="Success Response" doc:id="51931b48-e022-4e00-b55f-e6570014ff35" />
// 				</otherwise>
// 			</choice>
// 		</flow>
// 	</mule>`

// 	decoder := xml.NewDecoder(strings.NewReader(xmlData))
// 	var currentElement string
// 	var elementContent string
// 	var insideElement bool

// 	for {
// 		token, err := decoder.Token()
// 		if err != nil {
// 			break
// 		}

// 		switch t := token.(type) {
// 		case xml.StartElement:
// 			element := token.(xml.StartElement)
// 			currentElement = element.Name.Local
// 			insideElement = true
// 			// Initialize elementContent when starting to process the element
// 			elementContent = ""
// 			// Print element attributes
// 			for _, attr := range element.Attr {
// 				elementContent += attr.Name.Local + `="` + attr.Value + `" `
// 			}
// 		case xml.EndElement:
// 			if insideElement {
// 				// Print element content
// 				fmt.Printf("%s: %s\n", currentElement, elementContent)
// 				insideElement = false
// 			}
// 		case xml.CharData:
// 			if insideElement {
// 				elementContent += string(t)
// 			}
// 		}
// 	}
// }

//-------------------The code That works perfectly -------------------------------------------------------------------

// package main

// import (
//     "encoding/json"
//     "encoding/xml"
//     "fmt"
//     //"reflect"
//     "strings"
// )

// // Define a struct that can handle various XML structures
// type GenericXML struct {
//     XMLName xml.Name
//     Content interface{} `xml:",innerxml"`
// }

// func xmlToJSON(xmlData string) (string, error) {
//     decoder := xml.NewDecoder(strings.NewReader(xmlData))
//     var data GenericXML
//     if err := decoder.Decode(&data); err != nil {
//         return "", err
//     }

//     jsonData, err := json.Marshal(data.Content)
//     if err != nil {
//         return "", err
//     }

//     return string(jsonData), nil
// }

// func main() {
//     xmlData := `
//     <root>
//         <person>
//             <name>John</name>
//             <age>30</age>
//         </person>
//         <product>
//             <name>Laptop</name>
//             <price>1000</price>
//         </product>
//     </root>`

//     jsonString, err := xmlToJSON(xmlData)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     fmt.Println(jsonString)
// }

//-------------------------------------------------------------------------------------

// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"strings"

// 	xj "github.com/basgys/goxml2json"
// )

// const prefix = "veryuniqueattrprefix-"

// func main() {
// 	xml := strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?>

// 	<mule xmlns:db="http://www.mulesoft.org/schema/mule/db" xmlns:ee="http://www.mulesoft.org/schema/mule/ee/core"
// 		xmlns:compression="http://www.mulesoft.org/schema/mule/compression"
// 		xmlns:http="http://www.mulesoft.org/schema/mule/http" xmlns="http://www.mulesoft.org/schema/mule/core" xmlns:doc="http://www.mulesoft.org/schema/mule/documentation" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.mulesoft.org/schema/mule/core http://www.mulesoft.org/schema/mule/core/current/mule.xsd
// 	http://www.mulesoft.org/schema/mule/http http://www.mulesoft.org/schema/mule/http/current/mule-http.xsd
// 	http://www.mulesoft.org/schema/mule/compression http://www.mulesoft.org/schema/mule/compression/current/mule-compression.xsd
// 	http://www.mulesoft.org/schema/mule/ee/core http://www.mulesoft.org/schema/mule/ee/core/current/mule-ee.xsd
// 	http://www.mulesoft.org/schema/mule/db http://www.mulesoft.org/schema/mule/db/current/mule-db.xsd">
// 		<flow name="animeFlow" doc:id="8934cc76-3405-49bf-a626-da9ae53b4547" >
// 			<http:listener doc:name="Triggers Flow" doc:id="09e321e1-2d69-483b-8654-c05b12b8e67d" config-ref="HTTP_Listener_config" path="${HL.path}">
// 				<http:error-response statusCode="#[vars.status]" >
// 					<http:body ><![CDATA[#[output application/json --- payload]]]></http:body>
// 				</http:error-response>
// 			</http:listener>
// 			<set-variable value="#[payload.customerID]" doc:name="Set Variable" doc:id="af73ad48-b949-48e1-b64c-29031316e35b" variableName="customer_Id"/>
// 			<until-successful maxRetries="5" doc:name="Until Successful" doc:id="b1283b69-0067-4f23-bb99-f6036b334efe" millisBetweenRetries="50000">
// 				<try doc:name="Try" doc:id="2a8aaa75-fe2e-4728-a557-af78930f6be1" >
// 					<http:request method="GET" doc:name="HTTP: Get Anime API" doc:id="56950f58-6094-457c-b1d5-4aa5b094b7a0" url="${Request.URL}">
// 						<ee:repeatable-file-store-stream bufferUnit="MB" />
// 					</http:request>
// 				</try>
// 			</until-successful>
// 			<compression:decompress doc:name="Decompress:Data" doc:id="17e91826-ff45-45bc-8deb-12ffa227a9dd" outputMimeType="application/xml" outputEncoding="UTF-8">
// 				<compression:decompressor >
// 					<compression:gzip-decompressor />
// 				</compression:decompressor>
// 			</compression:decompress>
// 			<ee:transform doc:name="Build:converting XML to JSON" doc:id="0e294c24-5609-41a9-8ef3-7d5a9a939a7e" >
// 				<ee:message >
// 					<ee:set-payload ><![CDATA[%dw 2.0
// 	output application/json
// 	---
// 	payload]]></ee:set-payload>
// 				</ee:message>
// 				<ee:variables >
// 					<ee:set-variable variableName="response" ><![CDATA[%dw 2.0
// 	output application/json
// 	---
// 	payload]]></ee:set-variable>
// 				</ee:variables>
// 			</ee:transform>
// 			<logger level="INFO" doc:name="Logs payload" doc:id="90a11369-1dd2-4b4d-9587-39f61c239e88" message="#[payload]"/>
// 			<choice doc:name="Choice" doc:id="51d66274-de56-4cdd-b1a8-6bc14e9f03d2" >
// 				<when expression="#[vars.response.error=='banned' or vars.response.error=='Missing']">
// 					<raise-error doc:name="Raise error" doc:id="0f18a6f7-71cb-4b86-9731-b17926a20119" type="CUSTOM:UNKNOWN_ERROR"/>
// 				</when>
// 				<when expression="#[vars.response.error=='Not Found']">
// 					<raise-error doc:name="Raise error" doc:id="2a870867-5d26-4c43-8ec7-f0fdfb49cdcb" type="CUSTOM:NOT_FOUND"/>
// 				</when>
// 				<otherwise >
// 					<try doc:name="Try" doc:id="d9d85582-4115-48bd-a2b3-18dfca4d7d0a" >
// 						<db:select doc:name="Retrieve:Recommendation" doc:id="e3f22f19-a4e7-4993-a0af-8e2d7e26e40b" config-ref="Database_Config">
// 							<db:sql ><![CDATA[select MAX(recommendation) as recommendation from customer_detail where customerID=:customer_Id]]></db:sql>
// 							<db:input-parameters ><![CDATA[#[customer_Id:vars.customer_Id]]]></db:input-parameters>
// 						</db:select>
// 					</try>
// 					<set-variable value="#[payload[0].recommendation]" doc:name="Store:recommendation" doc:id="0c1eb81d-dc56-4321-83f1-affeaaa525d6" variableName="count"/>
// 					<ee:transform doc:name="Build:Manipulation On Data" doc:id="181aee64-3617-410c-944e-6755461c6237" >
// 						<ee:message >
// 							<ee:set-payload resource="Anime.dwl" />
// 						</ee:message>
// 					</ee:transform>
// 					<logger level="INFO" doc:name="Logger" doc:id="8a4a2614-0ecf-4489-96c0-65ac9b7a94f4" message="#[payload]"/>
// 					<try doc:name="Try" doc:id="b77fd3d6-0113-4328-99e4-21691434556a" >
// 						<db:insert doc:name="Insert:Data" doc:id="cfc334e6-ba4c-43cb-8791-d9f1f1acc2fa" config-ref="Database_Config">
// 							<db:sql ><![CDATA[insert into customer_detail(customerID,episodecount,startdate,enddate,Title,review_rating,tags,creators,recommendation,created_date)
// 	  values (:customerID,:episodecount,:startdate,:enddate,:Title,:review_rating,:tags,:creators,:recommendation,:created_date)]]></db:sql>
// 							<db:input-parameters ><![CDATA[#[{
// 		customerID:vars.customer_Id,
// 		episodecount:payload.episodecount,
// 	startdate:payload.startdate,
// 	enddate:payload.enddate,
// 	Title:payload.Title,
// 	review_rating:payload.review_rating,
// 	tags:payload.tags,
// 	creators:payload.creators,
// 	recommendation:payload.recommendation,
// 	created_date:payload.created_date
// 	}]]]></db:input-parameters>
// 						</db:insert>
// 					</try>
// 					<set-payload value='#[{"output":"The data are inserted successfully",&#10;	"customerID":vars.customer_Id&#10;}]' doc:name="Success Response" doc:id="51931b48-e022-4e00-b55f-e6570014ff35" />
// 				</otherwise>
// 			</choice>
// 		</flow>
// 	</mule>
// 	`)

// 	// Decode XML document
// 	root := &xj.Node{}
// 	err := xj.NewDecoder(xml).Decode(root)
// 	if err != nil {
// 		panic(err)
// 	}

// 	RemoveAttr(root)

// 	// Then encode it in JSON
// 	buf := new(bytes.Buffer)
// 	e := xj.NewEncoder(buf)
// 	err = e.Encode(root)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(buf.String())

// 	jsonString := `{"key": "value", "nested": {"number": 42}}`

// 	// Define an empty interface to store the parsed JSON
// 	var data map[string]interface{}

// 	// Unmarshal the JSON string into the map
// 	err = json.Unmarshal([]byte(jsonString), &data)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Now, 'data' contains the parsed JSON data
// 	// You can access values dynamically
// 	if value, ok := data["key"]; ok {
// 		fmt.Printf("Key: %v\n", value)
// 	}

// 	if nested, ok := data["nested"].(map[string]interface{}); ok {
// 		if number, ok := nested["number"]; ok {
// 			fmt.Printf("Number: %v\n", number)
// 		}
// 	}
// }

// func RemoveAttr(n *xj.Node) {
// 	for k, v := range n.Children {
// 		if strings.HasPrefix(k, prefix) {
// 			delete(n.Children, k)
// 		} else {
// 			for _, n := range v {
// 				RemoveAttr(n)
// 			}
// 		}
// 	}
// }

//-----------code works perfectly and produces the desired result------------------------------------

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	// "reflect"
// )

// // Recursive function to find a key in a nested JSON structure
// func findKey(data map[string]interface{}, targetKey string) interface{} {
// 	for key, value := range data {
// 		if key == targetKey {
// 			return value
// 		}

// 		if nestedData, ok := value.(map[string]interface{}); ok {
// 			if result := findKey(nestedData, targetKey); result != nil {
// 				return result
// 			}
// 		}

// 		if nestedDataArray, ok := value.([]interface{}); ok {
// 			for _, nestedData := range nestedDataArray {
// 				if nestedDataMap, isMap := nestedData.(map[string]interface{}); isMap {
// 					if result := findKey(nestedDataMap, targetKey); result != nil {
// 						return result
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	// JSON-formatted string
// 	jsonString := `{
// 		"mule": {
// 			"-compression": "http://www.mulesoft.org/schema/mule/compression",
// 			"-db": "http://www.mulesoft.org/schema/mule/db",
// 			"-doc": "http://www.mulesoft.org/schema/mule/documentation",
// 			"-ee": "http://www.mulesoft.org/schema/mule/ee/core",
// 			"-http": "http://www.mulesoft.org/schema/mule/http",
// 			"-schemaLocation": "http://www.mulesoft.org/schema/mule/core http://www.mulesoft.org/schema/mule/core/current/mule.xsd\nhttp://www.mulesoft.org/schema/mule/http http://www.mulesoft.org/schema/mule/http/current/mule-http.xsd\nhttp://www.mulesoft.org/schema/mule/compression http://www.mulesoft.org/schema/mule/compression/current/mule-compression.xsd\nhttp://www.mulesoft.org/schema/mule/ee/core http://www.mulesoft.org/schema/mule/ee/core/current/mule-ee.xsd\nhttp://www.mulesoft.org/schema/mule/db http://www.mulesoft.org/schema/mule/db/current/mule-db.xsd",
// 			"-xmlns": "http://www.mulesoft.org/schema/mule/core",
// 			"-xsi": "http://www.w3.org/2001/XMLSchema-instance",
// 			"flow": {
// 				"-id": "8934cc76-3405-49bf-a626-da9ae53b4547",
// 				"-name": "animeFlow",
// 				"choice": {
// 					"-id": "51d66274-de56-4cdd-b1a8-6bc14e9f03d2",
// 					"-name": "Choice",
// 					"otherwise": {
// 						"logger": {
// 							"-id": "8a4a2614-0ecf-4489-96c0-65ac9b7a94f4",
// 							"-level": "INFO",
// 							"-message": "#[payload]",
// 							"-name": "Logger"
// 						},
// 						"set-payload": {
// 							"-id": "51931b48-e022-4e00-b55f-e6570014ff35",
// 							"-name": "Success Response",
// 							"-value": "#[{\"output\":\"The data are inserted successfully\",\n\t\"customerID\":vars.customer_Id\n}]"
// 						},
// 						"set-variable": {
// 							"-id": "0c1eb81d-dc56-4321-83f1-affeaaa525d6",
// 							"-name": "Store:recommendation",
// 							"-value": "#[payload[0].recommendation]",
// 							"-variableName": "count"
// 						},
// 						"transform": {
// 							"-id": "181aee64-3617-410c-944e-6755461c6237",
// 							"-name": "Build:Manipulation On Data",
// 							"message": {
// 								"set-payload": {
// 									"-resource": "Anime.dwl"
// 								}
// 							}
// 						},
// 						"try": [
// 							{
// 								"-id": "d9d85582-4115-48bd-a2b3-18dfca4d7d0a",
// 								"-name": "Try",
// 								"select": {
// 									"-config-ref": "Database_Config",
// 									"-id": "e3f22f19-a4e7-4993-a0af-8e2d7e26e40b",
// 									"-name": "Retrieve:Recommendation",
// 									"input-parameters": "#[customer_Id:vars.customer_Id]",
// 									"sql": "select MAX(recommendation) as recommendation from customer_detail where customerID=:customer_Id"
// 								}
// 							},
// 							{
// 								"-id": "b77fd3d6-0113-4328-99e4-21691434556a",
// 								"-name": "Try",
// 								"insert": {
// 									"-config-ref": "Database_Config",
// 									"-id": "cfc334e6-ba4c-43cb-8791-d9f1f1acc2fa",
// 									"-name": "Insert:Data",
// 									"input-parameters": "#[{\n\tcustomerID:vars.customer_Id,\n\tepisodecount:payload.episodecount,\nstartdate:payload.startdate,\nenddate:payload.enddate,\nTitle:payload.Title,\nreview_rating:payload.review_rating,\ntags:payload.tags,\ncreators:payload.creators,\nrecommendation:payload.recommendation,\ncreated_date:payload.created_date\n}]",
// 									"sql": "insert into customer_detail(customerID,episodecount,startdate,enddate,Title,review_rating,tags,creators,recommendation,created_date)\n  values (:customerID,:episodecount,:startdate,:enddate,:Title,:review_rating,:tags,:creators,:recommendation,:created_date)"
// 								}
// 							}
// 						]
// 					},
// 					"when": [
// 						{
// 							"-expression": "#[vars.response.error=='banned' or vars.response.error=='Missing']",
// 							"raise-error": {
// 								"-id": "0f18a6f7-71cb-4b86-9731-b17926a20119",
// 								"-name": "Raise error",
// 								"-type": "CUSTOM:UNKNOWN_ERROR"
// 							}
// 						},
// 						{
// 							"-expression": "#[vars.response.error=='Not Found']",
// 							"raise-error": {
// 								"-id": "2a870867-5d26-4c43-8ec7-f0fdfb49cdcb",
// 								"-name": "Raise error",
// 								"-type": "CUSTOM:NOT_FOUND"
// 							}
// 						}
// 					]
// 				},
// 				"decompress": {
// 					"-id": "17e91826-ff45-45bc-8deb-12ffa227a9dd",
// 					"-name": "Decompress:Data",
// 					"-outputEncoding": "UTF-8",
// 					"-outputMimeType": "application/xml",
// 					"decompressor": {
// 						"gzip-decompressor": ""
// 					}
// 				},
// 				"listener": {
// 					"-config-ref": "HTTP_Listener_config",
// 					"-id": "09e321e1-2d69-483b-8654-c05b12b8e67d",
// 					"-name": "Triggers Flow",
// 					"-path": "${HL.path}",
// 					"error-response": {
// 						"-statusCode": "#[vars.status]",
// 						"body": "#[output application/json --- payload]"
// 					}
// 				},
// 				"logger": {
// 					"-id": "90a11369-1dd2-4b4d-9587-39f61c239e88",
// 					"-level": "INFO",
// 					"-message": "#[payload]",
// 					"-name": "Logs payload"
// 				},
// 				"set-variable": {
// 					"-id": "af73ad48-b949-48e1-b64c-29031316e35b",
// 					"-name": "Set Variable",
// 					"-value": "#[payload.customerID]",
// 					"-variableName": "customer_Id"
// 				},
// 				"transform": {
// 					"-id": "0e294c24-5609-41a9-8ef3-7d5a9a939a7e",
// 					"-name": "Build:converting XML to JSON",
// 					"message": {
// 						"set-payload": "%dw 2.0\noutput application/json\n---\npayload"
// 					},
// 					"variables": {
// 						"set-variable": {
// 							"#content": "%dw 2.0\noutput application/json\n---\npayload",
// 							"-variableName": "response"
// 						}
// 					}
// 				},
// 				"until-successful": {
// 					"-id": "b1283b69-0067-4f23-bb99-f6036b334efe",
// 					"-maxRetries": "5",
// 					"-millisBetweenRetries": "50000",
// 					"-name": "Until Successful",
// 					"try": {
// 						"-id": "2a8aaa75-fe2e-4728-a557-af78930f6be1",
// 						"-name": "Try",
// 						"request": {
// 							"-id": "56950f58-6094-457c-b1d5-4aa5b094b7a0",
// 							"-method": "GET",
// 							"-name": "HTTP: Get Anime API",
// 							"-url": "${Request.URL}",
// 							"repeatable-file-store-stream": {
// 								"-bufferUnit": "MB"
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}`

// 	// Define a map to store the parsed JSON data
// 	var data map[string]interface{}

// 	// Unmarshal the JSON string into the map
// 	err := json.Unmarshal([]byte(jsonString), &data)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Key you want to search for
// 	targetKey := "logger"

// 	// Call the recursive function to find the key in the JSON data
// 	loggerData := findKey(data, targetKey)

// 	if loggerData != nil {
// 		// You can work with the "logger" data here
// 		fmt.Printf("Found '%s' data: %v\n", targetKey, loggerData)
// 	} else {
// 		fmt.Printf("'%s' data not found\n", targetKey)
// 	}
// }
//------------------code works perfectly but finds only one instance---------

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"reflect"
// )

// // Recursive function to find all instances of a key in a nested JSON structure
// func findAllInstances(data interface{}, targetKey string) []interface{} {
// 	var results []interface{}

// 	switch reflect.TypeOf(data).Kind() {
// 	case reflect.Map:
// 		for key, value := range data.(map[string]interface{}) {
// 			if key == targetKey {
// 				results = append(results, value)
// 			}
// 			results = append(results, findAllInstances(value, targetKey)...)
// 		}
// 	case reflect.Slice:
// 		for _, value := range data.([]interface{}) {
// 			results = append(results, findAllInstances(value, targetKey)...)
// 		}
// 	}

// 	return results
// }

// func main() {
// 	// JSON-formatted string
// 	jsonString := `{
// 				"mule": {
// 					"-compression": "http://www.mulesoft.org/schema/mule/compression",
// 					"-db": "http://www.mulesoft.org/schema/mule/db",
// 					"-doc": "http://www.mulesoft.org/schema/mule/documentation",
// 					"-ee": "http://www.mulesoft.org/schema/mule/ee/core",
// 					"-http": "http://www.mulesoft.org/schema/mule/http",
// 					"-schemaLocation": "http://www.mulesoft.org/schema/mule/core http://www.mulesoft.org/schema/mule/core/current/mule.xsd\nhttp://www.mulesoft.org/schema/mule/http http://www.mulesoft.org/schema/mule/http/current/mule-http.xsd\nhttp://www.mulesoft.org/schema/mule/compression http://www.mulesoft.org/schema/mule/compression/current/mule-compression.xsd\nhttp://www.mulesoft.org/schema/mule/ee/core http://www.mulesoft.org/schema/mule/ee/core/current/mule-ee.xsd\nhttp://www.mulesoft.org/schema/mule/db http://www.mulesoft.org/schema/mule/db/current/mule-db.xsd",
// 					"-xmlns": "http://www.mulesoft.org/schema/mule/core",
// 					"-xsi": "http://www.w3.org/2001/XMLSchema-instance",
// 					"flow": {
// 						"-id": "8934cc76-3405-49bf-a626-da9ae53b4547",
// 						"-name": "animeFlow",
// 						"choice": {
// 							"-id": "51d66274-de56-4cdd-b1a8-6bc14e9f03d2",
// 							"-name": "Choice",
// 							"otherwise": {
// 								"logger": {
// 									"-id": "8a4a2614-0ecf-4489-96c0-65ac9b7a94f4",
// 									"-level": "INFO",
// 									"-message": "#[payload]",
// 									"-name": "Logger"
// 								},
// 								"set-payload": {
// 									"-id": "51931b48-e022-4e00-b55f-e6570014ff35",
// 									"-name": "Success Response",
// 									"-value": "#[{\"output\":\"The data are inserted successfully\",\n\t\"customerID\":vars.customer_Id\n}]"
// 								},
// 								"set-variable": {
// 									"-id": "0c1eb81d-dc56-4321-83f1-affeaaa525d6",
// 									"-name": "Store:recommendation",
// 									"-value": "#[payload[0].recommendation]",
// 									"-variableName": "count"
// 								},
// 								"transform": {
// 									"-id": "181aee64-3617-410c-944e-6755461c6237",
// 									"-name": "Build:Manipulation On Data",
// 									"message": {
// 										"set-payload": {
// 											"-resource": "Anime.dwl"
// 										}
// 									}
// 								},
// 								"try": [
// 									{
// 										"-id": "d9d85582-4115-48bd-a2b3-18dfca4d7d0a",
// 										"-name": "Try",
// 										"select": {
// 											"-config-ref": "Database_Config",
// 											"-id": "e3f22f19-a4e7-4993-a0af-8e2d7e26e40b",
// 											"-name": "Retrieve:Recommendation",
// 											"input-parameters": "#[customer_Id:vars.customer_Id]",
// 											"sql": "select MAX(recommendation) as recommendation from customer_detail where customerID=:customer_Id"
// 										}
// 									},
// 									{
// 										"-id": "b77fd3d6-0113-4328-99e4-21691434556a",
// 										"-name": "Try",
// 										"insert": {
// 											"-config-ref": "Database_Config",
// 											"-id": "cfc334e6-ba4c-43cb-8791-d9f1f1acc2fa",
// 											"-name": "Insert:Data",
// 											"input-parameters": "#[{\n\tcustomerID:vars.customer_Id,\n\tepisodecount:payload.episodecount,\nstartdate:payload.startdate,\nenddate:payload.enddate,\nTitle:payload.Title,\nreview_rating:payload.review_rating,\ntags:payload.tags,\ncreators:payload.creators,\nrecommendation:payload.recommendation,\ncreated_date:payload.created_date\n}]",
// 											"sql": "insert into customer_detail(customerID,episodecount,startdate,enddate,Title,review_rating,tags,creators,recommendation,created_date)\n  values (:customerID,:episodecount,:startdate,:enddate,:Title,:review_rating,:tags,:creators,:recommendation,:created_date)"
// 										}
// 									}
// 								]
// 							},
// 							"when": [
// 								{
// 									"-expression": "#[vars.response.error=='banned' or vars.response.error=='Missing']",
// 									"raise-error": {
// 										"-id": "0f18a6f7-71cb-4b86-9731-b17926a20119",
// 										"-name": "Raise error",
// 										"-type": "CUSTOM:UNKNOWN_ERROR"
// 									}
// 								},
// 								{
// 									"-expression": "#[vars.response.error=='Not Found']",
// 									"raise-error": {
// 										"-id": "2a870867-5d26-4c43-8ec7-f0fdfb49cdcb",
// 										"-name": "Raise error",
// 										"-type": "CUSTOM:NOT_FOUND"
// 									}
// 								}
// 							]
// 						},
// 						"decompress": {
// 							"-id": "17e91826-ff45-45bc-8deb-12ffa227a9dd",
// 							"-name": "Decompress:Data",
// 							"-outputEncoding": "UTF-8",
// 							"-outputMimeType": "application/xml",
// 							"decompressor": {
// 								"gzip-decompressor": ""
// 							}
// 						},
// 						"listener": {
// 							"-config-ref": "HTTP_Listener_config",
// 							"-id": "09e321e1-2d69-483b-8654-c05b12b8e67d",
// 							"-name": "Triggers Flow",
// 							"-path": "${HL.path}",
// 							"error-response": {
// 								"-statusCode": "#[vars.status]",
// 								"body": "#[output application/json --- payload]"
// 							}
// 						},
// 						"logger": {
// 							"-id": "90a11369-1dd2-4b4d-9587-39f61c239e88",
// 							"-level": "INFO",
// 							"-message": "#[payload]",
// 							"-name": "Logs payload"
// 						},
// 						"set-variable": {
// 							"-id": "af73ad48-b949-48e1-b64c-29031316e35b",
// 							"-name": "Set Variable",
// 							"-value": "#[payload.customerID]",
// 							"-variableName": "customer_Id"
// 						},
// 						"transform": {
// 							"-id": "0e294c24-5609-41a9-8ef3-7d5a9a939a7e",
// 							"-name": "Build:converting XML to JSON",
// 							"message": {
// 								"set-payload": "%dw 2.0\noutput application/json\n---\npayload"
// 							},
// 							"variables": {
// 								"set-variable": {
// 									"#content": "%dw 2.0\noutput application/json\n---\npayload",
// 									"-variableName": "response"
// 								}
// 							}
// 						},
// 						"until-successful": {
// 							"-id": "b1283b69-0067-4f23-bb99-f6036b334efe",
// 							"-maxRetries": "5",
// 							"-millisBetweenRetries": "50000",
// 							"-name": "Until Successful",
// 							"try": {
// 								"-id": "2a8aaa75-fe2e-4728-a557-af78930f6be1",
// 								"-name": "Try",
// 								"request": {
// 									"-id": "56950f58-6094-457c-b1d5-4aa5b094b7a0",
// 									"-method": "GET",
// 									"-name": "HTTP: Get Anime API",
// 									"-url": "${Request.URL}",
// 									"repeatable-file-store-stream": {
// 										"-bufferUnit": "MB"
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}`

// 	// Define an empty interface to store the parsed JSON data
// 	var data map[string]interface{}

// 	// Unmarshal the JSON string into the interface
// 	err := json.Unmarshal([]byte(jsonString), &data)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Key you want to search for
// 	targetKey := "logger"

// 	// Call the recursive function to find all instances of the key in the JSON data
// 	instances := findAllInstances(data, targetKey)

// 	if len(instances) > 0 {
// 		// You can work with the found instances of the key here
// 		fmt.Printf("Found %d instances of '%s' data:\n", len(instances), targetKey)
// 		for i, instance := range instances {
// 			fmt.Printf("Instance %d: %v\n", i+1, instance)
// 		}
// 	} else {
// 		fmt.Printf("No instances of '%s' data found\n", targetKey)
// 	}
// }

// //--------------------------------code works------------------------

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"reflect"
// 	"strings"
// )

// // Recursive function to find all instances of a key with the "-name" format in a nested JSON structure
// func findAllNameKeys(data interface{}, targetKey string) []interface{} {
// 	var results []interface{}

// 	switch reflect.TypeOf(data).Kind() {
// 	case reflect.Map:
// 		for key, value := range data.(map[string]interface{}) {
// 			if strings.HasSuffix(key, targetKey) {
// 				results = append(results, value)
// 			}
// 			results = append(results, findAllNameKeys(value, targetKey)...)
// 		}
// 	case reflect.Slice:
// 		for _, value := range data.([]interface{}) {
// 			results = append(results, findAllNameKeys(value, targetKey)...)
// 		}
// 	}

// 	return results
// }

// func main() {
// 	// JSON-formatted string
// 	jsonString := `{
// 						"mule": {
// 							"-compression": "http://www.mulesoft.org/schema/mule/compression",
// 							"-db": "http://www.mulesoft.org/schema/mule/db",
// 							"-doc": "http://www.mulesoft.org/schema/mule/documentation",
// 							"-ee": "http://www.mulesoft.org/schema/mule/ee/core",
// 							"-http": "http://www.mulesoft.org/schema/mule/http",
// 							"-schemaLocation": "http://www.mulesoft.org/schema/mule/core http://www.mulesoft.org/schema/mule/core/current/mule.xsd\nhttp://www.mulesoft.org/schema/mule/http http://www.mulesoft.org/schema/mule/http/current/mule-http.xsd\nhttp://www.mulesoft.org/schema/mule/compression http://www.mulesoft.org/schema/mule/compression/current/mule-compression.xsd\nhttp://www.mulesoft.org/schema/mule/ee/core http://www.mulesoft.org/schema/mule/ee/core/current/mule-ee.xsd\nhttp://www.mulesoft.org/schema/mule/db http://www.mulesoft.org/schema/mule/db/current/mule-db.xsd",
// 							"-xmlns": "http://www.mulesoft.org/schema/mule/core",
// 							"-xsi": "http://www.w3.org/2001/XMLSchema-instance",
// 							"flow": {
// 								"-id": "8934cc76-3405-49bf-a626-da9ae53b4547",
// 								"-name": "animeFlow",
// 								"choice": {
// 									"-id": "51d66274-de56-4cdd-b1a8-6bc14e9f03d2",
// 									"-name": "Choice",
// 									"otherwise": {
// 										"logger": {
// 											"-id": "8a4a2614-0ecf-4489-96c0-65ac9b7a94f4",
// 											"-level": "INFO",
// 											"-message": "#[payload]",
// 											"-name": "Logger"
// 										},
// 										"set-payload": {
// 											"-id": "51931b48-e022-4e00-b55f-e6570014ff35",
// 											"-name": "Success Response",
// 											"-value": "#[{\"output\":\"The data are inserted successfully\",\n\t\"customerID\":vars.customer_Id\n}]"
// 										},
// 										"set-variable": {
// 											"-id": "0c1eb81d-dc56-4321-83f1-affeaaa525d6",
// 											"-name": "Store:recommendation",
// 											"-value": "#[payload[0].recommendation]",
// 											"-variableName": "count"
// 										},
// 										"transform": {
// 											"-id": "181aee64-3617-410c-944e-6755461c6237",
// 											"-name": "Build:Manipulation On Data",
// 											"message": {
// 												"set-payload": {
// 													"-resource": "Anime.dwl"
// 												}
// 											}
// 										},
// 										"try": [
// 											{
// 												"-id": "d9d85582-4115-48bd-a2b3-18dfca4d7d0a",
// 												"-name": "Try",
// 												"select": {
// 													"-config-ref": "Database_Config",
// 													"-id": "e3f22f19-a4e7-4993-a0af-8e2d7e26e40b",
// 													"-name": "Retrieve:Recommendation",
// 													"input-parameters": "#[customer_Id:vars.customer_Id]",
// 													"sql": "select MAX(recommendation) as recommendation from customer_detail where customerID=:customer_Id"
// 												}
// 											},
// 											{
// 												"-id": "b77fd3d6-0113-4328-99e4-21691434556a",
// 												"-name": "Try",
// 												"insert": {
// 													"-config-ref": "Database_Config",
// 													"-id": "cfc334e6-ba4c-43cb-8791-d9f1f1acc2fa",
// 													"-name": "Insert:Data",
// 													"input-parameters": "#[{\n\tcustomerID:vars.customer_Id,\n\tepisodecount:payload.episodecount,\nstartdate:payload.startdate,\nenddate:payload.enddate,\nTitle:payload.Title,\nreview_rating:payload.review_rating,\ntags:payload.tags,\ncreators:payload.creators,\nrecommendation:payload.recommendation,\ncreated_date:payload.created_date\n}]",
// 													"sql": "insert into customer_detail(customerID,episodecount,startdate,enddate,Title,review_rating,tags,creators,recommendation,created_date)\n  values (:customerID,:episodecount,:startdate,:enddate,:Title,:review_rating,:tags,:creators,:recommendation,:created_date)"
// 												}
// 											}
// 										]
// 									},
// 									"when": [
// 										{
// 											"-expression": "#[vars.response.error=='banned' or vars.response.error=='Missing']",
// 											"raise-error": {
// 												"-id": "0f18a6f7-71cb-4b86-9731-b17926a20119",
// 												"-name": "Raise error",
// 												"-type": "CUSTOM:UNKNOWN_ERROR"
// 											}
// 										},
// 										{
// 											"-expression": "#[vars.response.error=='Not Found']",
// 											"raise-error": {
// 												"-id": "2a870867-5d26-4c43-8ec7-f0fdfb49cdcb",
// 												"-name": "Raise error",
// 												"-type": "CUSTOM:NOT_FOUND"
// 											}
// 										}
// 									]
// 								},
// 								"decompress": {
// 									"-id": "17e91826-ff45-45bc-8deb-12ffa227a9dd",
// 									"-name": "Decompress:Data",
// 									"-outputEncoding": "UTF-8",
// 									"-outputMimeType": "application/xml",
// 									"decompressor": {
// 										"gzip-decompressor": ""
// 									}
// 								},
// 								"listener": {
// 									"-config-ref": "HTTP_Listener_config",
// 									"-id": "09e321e1-2d69-483b-8654-c05b12b8e67d",
// 									"-name": "Triggers Flow",
// 									"-path": "${HL.path}",
// 									"error-response": {
// 										"-statusCode": "#[vars.status]",
// 										"body": "#[output application/json --- payload]"
// 									}
// 								},
// 								"logger": {
// 									"-id": "90a11369-1dd2-4b4d-9587-39f61c239e88",
// 									"-level": "INFO",
// 									"-message": "#[payload]",
// 									"-name": "Logs payload"
// 								},
// 								"set-variable": {
// 									"-id": "af73ad48-b949-48e1-b64c-29031316e35b",
// 									"-name": "Set Variable",
// 									"-value": "#[payload.customerID]",
// 									"-variableName": "customer_Id"
// 								},
// 								"transform": {
// 									"-id": "0e294c24-5609-41a9-8ef3-7d5a9a939a7e",
// 									"-name": "Build:converting XML to JSON",
// 									"message": {
// 										"set-payload": "%dw 2.0\noutput application/json\n---\npayload"
// 									},
// 									"variables": {
// 										"set-variable": {
// 											"#content": "%dw 2.0\noutput application/json\n---\npayload",
// 											"-variableName": "response"
// 										}
// 									}
// 								},
// 								"until-successful": {
// 									"-id": "b1283b69-0067-4f23-bb99-f6036b334efe",
// 									"-maxRetries": "5",
// 									"-millisBetweenRetries": "50000",
// 									"-name": "Until Successful",
// 									"try": {
// 										"-id": "2a8aaa75-fe2e-4728-a557-af78930f6be1",
// 										"-name": "Try",
// 										"request": {
// 											"-id": "56950f58-6094-457c-b1d5-4aa5b094b7a0",
// 											"-method": "GET",
// 											"-name": "HTTP: Get Anime API",
// 											"-url": "${Request.URL}",
// 											"repeatable-file-store-stream": {
// 												"-bufferUnit": "MB"
// 											}
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}`

// 	// Define an empty interface to store the parsed JSON data
// 	var data map[string]interface{}

// 	// Unmarshal the JSON string into the interface
// 	err := json.Unmarshal([]byte(jsonString), &data)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Key you want to search for (in this case, "-name")
// 	targetKey := "-name"

// 	// Call the recursive function to find all instances of keys with the "-name" format in the JSON data
// 	instances := findAllNameKeys(data, targetKey)

// 	if len(instances) > 0 {
// 		// You can work with the found instances of the "-name" keys here
// 		fmt.Printf("Found %d instances of keys with the format '%s' data:\n", len(instances), targetKey)
// 		for i, instance := range instances {
// 			fmt.Printf("Instance %d: %v\n", i+1, instance)
// 		}
// 	} else {
// 		fmt.Printf("No instances of keys with the format '%s' data found\n", targetKey)
// 	}
// }
//------------------------------------we can extract the code for the name alone-------

package main

import (
	"encoding/json"
	"fmt"
	// "reflect"
)

// Recursive function to find the "-name" field in "logger" objects
func findLoggerNames(data interface{}) []string {
	names := []string{}

	switch data := data.(type) {
	case map[string]interface{}:
		// Check if this object is a "logger"
		if value, exists := data["-name"]; exists {
			if loggerName, ok := value.(string); ok {
				names = append(names, loggerName)
			}
		}

		// Recursively search within the "logger" object
		for _, value := range data {
			names = append(names, findLoggerNames(value)...)
		}
	case []interface{}:
		// Recursively search within the slice
		for _, value := range data {
			names = append(names, findLoggerNames(value)...)
		}
	}

	return names
}

func main() {
	// JSON-formatted string
	jsonString := `{
								"mule": {
									"-compression": "http://www.mulesoft.org/schema/mule/compression",
									"-db": "http://www.mulesoft.org/schema/mule/db",
									"-doc": "http://www.mulesoft.org/schema/mule/documentation",
									"-ee": "http://www.mulesoft.org/schema/mule/ee/core",
									"-http": "http://www.mulesoft.org/schema/mule/http",
									"-schemaLocation": "http://www.mulesoft.org/schema/mule/core http://www.mulesoft.org/schema/mule/core/current/mule.xsd\nhttp://www.mulesoft.org/schema/mule/http http://www.mulesoft.org/schema/mule/http/current/mule-http.xsd\nhttp://www.mulesoft.org/schema/mule/compression http://www.mulesoft.org/schema/mule/compression/current/mule-compression.xsd\nhttp://www.mulesoft.org/schema/mule/ee/core http://www.mulesoft.org/schema/mule/ee/core/current/mule-ee.xsd\nhttp://www.mulesoft.org/schema/mule/db http://www.mulesoft.org/schema/mule/db/current/mule-db.xsd",
									"-xmlns": "http://www.mulesoft.org/schema/mule/core",
									"-xsi": "http://www.w3.org/2001/XMLSchema-instance",
									"flow": {
										"-id": "8934cc76-3405-49bf-a626-da9ae53b4547",
										"-name": "animeFlow",
										"choice": {
											"-id": "51d66274-de56-4cdd-b1a8-6bc14e9f03d2",
											"-name": "Choice",
											"otherwise": {
												"logger": {
													"-id": "8a4a2614-0ecf-4489-96c0-65ac9b7a94f4",
													"-level": "INFO",
													"-message": "#[payload]",
													"-name": "Logger"
												},
												"set-payload": {
													"-id": "51931b48-e022-4e00-b55f-e6570014ff35",
													"-name": "Success Response",
													"-value": "#[{\"output\":\"The data are inserted successfully\",\n\t\"customerID\":vars.customer_Id\n}]"
												},
												"set-variable": {
													"-id": "0c1eb81d-dc56-4321-83f1-affeaaa525d6",
													"-name": "Store:recommendation",
													"-value": "#[payload[0].recommendation]",
													"-variableName": "count"
												},
												"transform": {
													"-id": "181aee64-3617-410c-944e-6755461c6237",
													"-name": "Build:Manipulation On Data",
													"message": {
														"set-payload": {
															"-resource": "Anime.dwl"
														}
													}
												},
												"try": [
													{
														"-id": "d9d85582-4115-48bd-a2b3-18dfca4d7d0a",
														"-name": "Try",
														"select": {
															"-config-ref": "Database_Config",
															"-id": "e3f22f19-a4e7-4993-a0af-8e2d7e26e40b",
															"-name": "Retrieve:Recommendation",
															"input-parameters": "#[customer_Id:vars.customer_Id]",
															"sql": "select MAX(recommendation) as recommendation from customer_detail where customerID=:customer_Id"
														}
													},
													{
														"-id": "b77fd3d6-0113-4328-99e4-21691434556a",
														"-name": "Try",
														"insert": {
															"-config-ref": "Database_Config",
															"-id": "cfc334e6-ba4c-43cb-8791-d9f1f1acc2fa",
															"-name": "Insert:Data",
															"input-parameters": "#[{\n\tcustomerID:vars.customer_Id,\n\tepisodecount:payload.episodecount,\nstartdate:payload.startdate,\nenddate:payload.enddate,\nTitle:payload.Title,\nreview_rating:payload.review_rating,\ntags:payload.tags,\ncreators:payload.creators,\nrecommendation:payload.recommendation,\ncreated_date:payload.created_date\n}]",
															"sql": "insert into customer_detail(customerID,episodecount,startdate,enddate,Title,review_rating,tags,creators,recommendation,created_date)\n  values (:customerID,:episodecount,:startdate,:enddate,:Title,:review_rating,:tags,:creators,:recommendation,:created_date)"
														}
													}
												]
											},
											"when": [
												{
													"-expression": "#[vars.response.error=='banned' or vars.response.error=='Missing']",
													"raise-error": {
														"-id": "0f18a6f7-71cb-4b86-9731-b17926a20119",
														"-name": "Raise error",
														"-type": "CUSTOM:UNKNOWN_ERROR"
													}
												},
												{
													"-expression": "#[vars.response.error=='Not Found']",
													"raise-error": {
														"-id": "2a870867-5d26-4c43-8ec7-f0fdfb49cdcb",
														"-name": "Raise error",
														"-type": "CUSTOM:NOT_FOUND"
													}
												}
											]
										},
										"decompress": {
											"-id": "17e91826-ff45-45bc-8deb-12ffa227a9dd",
											"-name": "Decompress:Data",
											"-outputEncoding": "UTF-8",
											"-outputMimeType": "application/xml",
											"decompressor": {
												"gzip-decompressor": ""
											}
										},
										"listener": {
											"-config-ref": "HTTP_Listener_config",
											"-id": "09e321e1-2d69-483b-8654-c05b12b8e67d",
											"-name": "Triggers Flow",
											"-path": "${HL.path}",
											"error-response": {
												"-statusCode": "#[vars.status]",
												"body": "#[output application/json --- payload]"
											}
										},
										"logger": {
											"-id": "90a11369-1dd2-4b4d-9587-39f61c239e88",
											"-level": "INFO",
											"-message": "#[payload]",
											"-name": "Logs payload"
										},
										"set-variable": {
											"-id": "af73ad48-b949-48e1-b64c-29031316e35b",
											"-name": "Set Variable",
											"-value": "#[payload.customerID]",
											"-variableName": "customer_Id"
										},
										"transform": {
											"-id": "0e294c24-5609-41a9-8ef3-7d5a9a939a7e",
											"-name": "Build:converting XML to JSON",
											"message": {
												"set-payload": "%dw 2.0\noutput application/json\n---\npayload"
											},
											"variables": {
												"set-variable": {
													"#content": "%dw 2.0\noutput application/json\n---\npayload",
													"-variableName": "response"
												}
											}
										},
										"until-successful": {
											"-id": "b1283b69-0067-4f23-bb99-f6036b334efe",
											"-maxRetries": "5",
											"-millisBetweenRetries": "50000",
											"-name": "Until Successful",
											"try": {
												"-id": "2a8aaa75-fe2e-4728-a557-af78930f6be1",
												"-name": "Try",
												"request": {
													"-id": "56950f58-6094-457c-b1d5-4aa5b094b7a0",
													"-method": "GET",
													"-name": "HTTP: Get Anime API",
													"-url": "${Request.URL}",
													"repeatable-file-store-stream": {
														"-bufferUnit": "MB"
													}
												}
											}
										}
									}
								}
							}`

	// Define an empty interface to store the parsed JSON data
	var data map[string]interface{}

	// Unmarshal the JSON string into the interface
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Call the recursive function to find "-name" fields within "logger" objects
	loggerNames := findLoggerNames(data)

	if len(loggerNames) > 0 {
		// You can work with the "-name" fields of "logger" objects here
		fmt.Println("Found the following logger names:")
		for i, name := range loggerNames {
			fmt.Printf("Logger %d: %s\n", i+1, name)
		}
	} else {
		fmt.Println("No logger names found.")
	}
}

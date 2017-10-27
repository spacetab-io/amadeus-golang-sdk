# Amadeus WS connector

This package contains structures, forms, functions and SOAP handler for Amadeus WS.

## SOAP Methods implemented

This package was made as proof of concept and for now have only Sessions handling and Command Cryptic methods.
It will be developed if it will be interested for people...

## Using

```golang
package main

import (
	"fmt"
	"log"

	//TODO: fix these path after the merge
	"github.com/kidem/amadeus-ws-go"
	"github.com/kidem/amadeus-ws-go/formats"
	sa "github.com/kidem/amadeus-ws-go/reqstructs/security_authenticate"
)

func main() {
	sess := authenticate("OfficeId")

	fmt.Println(sess)
}

func authenticate(officeId string) *amadeus.Session {
	service := amadeus.NewAmadeusWebServicesPT("https://test.webservices.amadeus.com", true, "WSAPId", &amadeus.BasicAuth{})
	response, session, err := service.SecurityAuthenticate(&sa.SecurityAuthenticate{
		UserIdentifier: &sa.UserIdentificationType{
			OriginIdentification: &sa.OriginatorIdentificationDetailsTypeI{
				SourceOffice: formats.AlphaNumericString_Length1To9(officeId),
			},
			Originator:         formats.AlphaNumericString_Length1To30("Originator"),
			OriginatorTypeCode: formats.AlphaNumericString_Length1To1("U"),
		},
		DutyCode: &sa.ReferenceInformationTypeI{
			DutyCodeDetails: &sa.ReferencingDetailsTypeI{
				ReferenceQualifier:  formats.AlphaNumericString_Length1To3("DUT"),
				ReferenceIdentifier: formats.AlphaNumericString_Length1To35("SU"),
			},
		},
		SystemDetails: &sa.SystemDetailsInfoType{
			OrganizationDetails: &sa.SystemDetailsTypeI{
				OrganizationId: formats.AlphaNumericString_Length1To35("OrganizationId"),
			},
		},
		PasswordInfo: &sa.BinaryDataType{
			DataLength: formats.NumericInteger_Length1To15(8),
			DataType:   formats.AlphaNumericString_Length1To1("E"),
			BinaryData: formats.AlphaNumericString_Length1To99999("Base64EncodedPassword"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if response.ErrorSection != nil {
		log.Fatal(string(*response.ErrorSection.InteractiveFreeText.FreeText))
	}

	return session
}
 ```

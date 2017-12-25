# Amadeus WBS SDK

This package contains structures, forms, functions and SOAP handler for Amadeus WS.

## SOAP Methods implemented

The following versions of services are currently implemented:

* Air_FlightInfo (05.1)
* Air_SellFromRecommendation (05.2)
* Command_Cryptic (07.3)
* DocIssuance_IssueTicket (09.1)
* Fare_CheckRules (07.1)
* Fare_ConvertCurrency (08.1)
* Fare_InformativeBestPricingWithoutPNR (12.4)
* Fare_InformativePricingWithoutPNR (12.4)
* Fare_MasterPricerCalendar (14.3, 12.2)
* Fare_MasterPricerTravelBoardSearch (14.3, 12.3)
* Fare_PricePNRWithBookingClass (14.1, 12.4)
* PNR_AddMultiElements (11.3)
* PNR_Cancel (11.3)
* PNR_Retrieve (11.3)
* Queue_CountPlanner (03.1)
* Queue_CountTotal (03.1)
* Queue_List (11.1)
* Queue_MoveItem (03.1)
* Queue_PlacePNR (03.1)
* Queue_RemoveItem (03.1)
* Security_Authenticate (06.1)
* Security_SignOut (04.1)
* Ticket_CreateTSTFromPricing (04.1)
* Ticket_CreditCardCheck (06.1)
* Ticket_DeleteTST (04.1)
* Ticket_DisplayTST (07.1)

### Example

```golang
package main

import (
	"fmt"
	"log"

	"github.com/tmconsulting/amadeus-golang-sdk"
	"github.com/tmconsulting/amadeus-golang-sdk/formats"
	sa "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/security_authenticate"
)

const (
	url  = "https://test.webservices.amadeus.com"
	wsap = "WSAPId"
)

func main() {
	sess := authenticate("OfficeId", "Originator", "U", "DUT", "SU", "OrganizationId", 8, "E", "Base64EncodedPassword")

	fmt.Println(sess)
}

func authenticate(officeId string, originator string, originatorTypeCode string,
                  referenceQualifier string, referenceIdentifier string, organizationId string,
                  passwordLength int, passwordType string, passwordData string) *amadeus.Session {
	service := amadeus.NewAmadeusWebServicesPT(url, true, wsap, &amadeus.BasicAuth{})
	response, session, err := service.SecurityAuthenticate(&sa.SecurityAuthenticate{
		UserIdentifier: []*sa.UserIdentificationType{
			{
				OriginIdentification: &sa.OriginatorIdentificationDetailsTypeI{
					SourceOffice: formats.AlphaNumericString_Length1To9(officeId),
				},
				Originator:         formats.AlphaNumericString_Length1To30(originator),
				OriginatorTypeCode: formats.AlphaNumericString_Length1To1(originatorTypeCode),
			},
		},
		DutyCode: &sa.ReferenceInformationTypeI{
			DutyCodeDetails: &sa.ReferencingDetailsTypeI{
				ReferenceQualifier:  formats.AlphaNumericString_Length1To3(referenceQualifier),
				ReferenceIdentifier: formats.AlphaNumericString_Length1To35(referenceIdentifier),
			},
		},
		SystemDetails: &sa.SystemDetailsInfoType{
			OrganizationDetails: &sa.SystemDetailsTypeI{
				OrganizationId: formats.AlphaNumericString_Length1To35(organizationId),
			},
		},
		PasswordInfo: []*sa.BinaryDataType{
			{
				DataLength: formats.NumericInteger_Length1To15(passwordLength),
				DataType:   formats.AlphaNumericString_Length1To1(passwordType),
				BinaryData: formats.AlphaNumericString_Length1To99999(passwordData),
			},
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

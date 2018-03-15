package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"strings"

	"github.com/tmconsulting/amadeus-golang-sdk/formats"
	sa "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/security_authenticate"
	soap2 "github.com/tmconsulting/amadeus-golang-sdk/soap2.0"
)

const (
	url  = "https://test.webservices.amadeus.com"
	wsap = "WSAPId"
)

func main() {
	var officeId = "OfficeId"
	var originator = "Originator"
	var organizationId = "OrganizationId"
	var clearPassword = "AMADEUS"
	s := sha1.New()
	s.Write([]byte(clearPassword))
	sess := authenticate(officeId, originator, "U", "DUT", "SU", organizationId, len(clearPassword), "E", string(s.Sum(nil)))

	fmt.Println(sess)
}

func authenticate(officeId, originator, originatorTypeCode, referenceQualifier, referenceIdentifier, organizationId string, passwordLength int, passwordType, passwordData string) *soap2.Session {
	service := soap2.NewAmadeusWebServicesPT(url, true, wsap)
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
		log.Fatalln("Error message:", strings.Join(response.ErrorSection.InteractiveFreeText.FreeText, "\n"))
	}

	return session
}

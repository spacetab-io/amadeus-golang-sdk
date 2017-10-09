# Amadeus WS connector

This package contains structures, forms, functions and SOAP handler for Amadeus WS.

## SOAP Methods implemented

This package was made as proof of concept and for now have only Sessions handling and Command Cryptic methods.
It will be developed if it will be interested for people...

## Using

    package main
    
    import (
        "fmt"
    	"log"
    
    	"github.com/smgladkovskiy/amadeus-ws-go"
    )
    
    func main() {
    	sess := authenticate("OfficeId")
    	
    	fmt.Println(sess)
    }
    
    func authenticate(officeId string) *amadeus.Session {
    	service := amadeus.NewAmadeusWebServicesPT("https://test.webservices.amadeus.com", true, "WSAPId", &amadeus.BasicAuth{})
    	response, session, err := service.SecurityAuthenticate(&amadeus.SecurityAuthenticate{
    		UserIdentifier: amadeus.UserIdentificationType{
    			OriginIdentification: &amadeus.OriginatorIdentificationDetailsTypeI{
    				SourceOffice: amadeus.AlphaNumericStringLength1To9(officeId),
    			},
    			Originator:         amadeus.AlphaNumericStringLength1To30("Originator"),
    			OriginatorTypeCode: amadeus.AlphaNumericStringLength1To1("U"),
    		},
    		DutyCode: amadeus.ReferenceInformationTypeI{
    			DutyCodeDetails: amadeus.ReferencingDetailsTypeI{
    				ReferenceQualifier:  amadeus.AlphaNumericStringLength1To3("DUT"),
    				ReferenceIdentifier: amadeus.AlphaNumericStringLength1To35("SU"),
    			},
    		},
    		SystemDetails: amadeus.SystemDetailsInfoType{
    			OrganizationDetails: amadeus.SystemDetailsTypeI{
    				OrganizationId: amadeus.AlphaNumericStringLength1To35("OrganizationId"),
    			},
    		},
    		PasswordInfo: amadeus.BinaryDataType{
    			DataLength: amadeus.NumericIntegerLength1To15(8),
    			DataType:   amadeus.AlphaNumericStringLength1To1("E"),
    			BinaryData: amadeus.AlphaNumericStringLength1To99999("Base64EncodedPassword"),
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
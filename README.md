# Amadeus WBS SDK

[![CircleCI](https://circleci.com/gh/tmconsulting/amadeus-golang-sdk/tree/temass-refactoringsts.svg?style=shield)](https://circleci.com/gh/tmconsulting/amadeus-golang-sdk/tree/tests) [![codecov](https://codecov.io/gh/tmconsulting/amadeus-golang-sdk/branch/mass-refactoring/graph/badge.svg)](https://codecov.io/gh/tmconsulting/amadeus-golang-sdk)

This package contains structures, forms, functions and SOAP handler for Amadeus WS.

## Methods implementation progress

- [ ] Air_FlightInfo (05.1)
- [x] Air_SellFromRecommendation (05.2)
- [x] Command_Cryptic (07.3)
- [x] DocIssuance_IssueTicket (09.1)
- [x] Fare_CheckRules (07.1)
- [ ] Fare_ConvertCurrency (08.1)
- [x] Fare_InformativeBestPricingWithoutPNR (12.4)
- [x] Fare_InformativePricingWithoutPNR (12.4)
- [ ] Fare_MasterPricerCalendar (14.3, 12.2)
- [x] Fare_MasterPricerTravelBoardSearch (16.3, 14.3)
- [x] Fare_PricePNRWithBookingClass (14.1)
- [x] PNR_AddMultiElements (11.3)
- [x] PNR_Cancel (11.3)
- [x] PNR_Retrieve (11.3)
- [ ] Queue_CountPlanner (03.1)
- [ ] Queue_CountTotal (03.1)
- [ ] Queue_List (11.1)
- [ ] Queue_MoveItem (03.1)
- [ ] Queue_PlacePNR (03.1)
- [ ] Queue_RemoveItem (03.1)
- [x] Security_Authenticate (06.1)
- [x] Security_SignOut (04.1)
- [ ] Service_BookPriceService
- [x] Ticket_CreateTSTFromPricing (04.1)
- [ ] Ticket_CreditCardCheck (06.1)
- [x] Ticket_DeleteTST (04.1)
- [x] Ticket_DisplayTST (07.1)
- [x] Ticket_ProcessEDoc (15.2)

## Installation

It is go gettable and go.mod powered

    $ go get github.com/tmconsulting/amadeus-golang-sdk@mass-refactoring

## Usage

Prepare log writer realisation if you need to see outgoing and incoming xmls (null logging used if nil is passed). 
 Check methods version that will be used in SDK (use `GetLatestMethodsMap()` to use latest methods versions that are 
 implemented). Ыуе credentials to connect: url, originator, password (not in base64!). Initiate SDK and service and use 
 service methods:

```go
package main

import (
    "log"
    
    "github.com/tmconsulting/amadeus-golang-sdk/client"
    "github.com/tmconsulting/amadeus-golang-sdk/service"
    "github.com/tmconsulting/amadeus-golang-sdk/structs/commandCryptic"
)

func main() {
    url := "https://nodeD1.test.webservices.amadeus.com/1ASIWXXXXXX"
    originator := "WSBENXXX"
    passwordRaw := "dGhlIHBhc3N3b3Jk"
    officeID := "BRUXX1111"
    
    cl := client.New(client.SetURL(url), client.SetUser(originator), client.SetPassword(passwordRaw), client.SetAgent(officeID))
    s := service.New(cl)
    
    resp, err := s.CommandCryptic("AN20DECMOWLED/ALH")
    if err != nil {
        panic(err)
    }
    
    log.Printf("responses: %v\n", resp)
}
```

## Testing

Run tests (`make test` or `go test ./... -v`) with ENV variabless:
* URL
* ORIGINATOR
* PASSWORD_RAW
* OFFICE_ID

## Contribution

Contribution, in any kind of way, is highly welcome!
It doesn't matter if you are not able to write code.
Creating issues or holding talks and help other people to use 
[amadeus-golang-sdk](https://github.com/tmconsulting/amadeus-golang-sdk) is contribution, too!

A few examples:

* Correct typos in the README / documentation
* Reporting bugs
* Implement a new feature or endpoint
* Sharing the love if like to use [amadeus-golang-sdk](https://github.com/tmconsulting/amadeus-golang-sdk) and help people 
to get use to it

If you are new to pull requests, checkout [Collaborating on projects using issues and pull requests / Creating a pull request](https://help.github.com/articles/creating-a-pull-request/).

## License

SDK is released under the [MIT License](./LICENSE).

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
* Fare_MasterPricerTravelBoardSearch (16.3, 14.3, 12.3)
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

## Installation

It is go gettable

    $ go get github.com/tmconsulting/amadeus-golang-sdk

```go
package main

import (
	amadeusSdk "github.com/tmconsulting/amadeus-golang-sdk"
)
...
```

## Usage examples

There are several usage examples in `./example` folder. Try it out.

## Tests

There are no tests yet. :( Feel free to help us to change this situation!

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

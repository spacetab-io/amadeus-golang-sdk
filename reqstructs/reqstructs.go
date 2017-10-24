package reqstructs

import "encoding/xml"

type QueueList struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A Queue_List"`

	// presence implies that this is a follow up scrolling entry to a previous entry. Absence implies start of a new search
	Scroll *ActionDetailsTypeI `xml:"scroll,omitempty"`

	// used to specify the target office for which the queue count is to be displayed
	TargetOffice *AdditionalBusinessSourceInformationTypeI `xml:"targetOffice,omitempty"`

	// used to specify the queue if required
	QueueNumber *QueueInformationTypeI `xml:"queueNumber,omitempty"`

	// used to select the category
	CategoryDetails *SubQueueInformationTypeI `xml:"categoryDetails,omitempty"`

	// date range as system defined
	Date *StructuredDateTimeInformationType `xml:"date,omitempty"`

	// defines the start point for the search and may also define the end point of the search
	ScanRange *RangeDetailsTypeI `xml:"scanRange,omitempty"`

	SearchCriteria *SearchCriteria `xml:"searchCriteria,omitempty"`

	// Passenger name list (all the names in the PNR).
	PassengerName *TravellerInformationTypeI `xml:"passengerName,omitempty"`

	// The last 2 characters of the sine of the agent who placed the PNR on queue.
	AgentSine *UserIdentificationType `xml:"agentSine,omitempty"`

	// Account number issue from AIAN entry in the PNR.
	AccountNumber *AccountingInformationElementType `xml:"accountNumber,omitempty"`

	FlightInformation *FlightInformation `xml:"flightInformation,omitempty"`

	// This is the point of sale of segments in PNRs: - 9 char Amadeus Office ID. - OR 2 char GDS code for OA PNRs  PNRs containing a segment sold in any Amadeus Office ID matching pattern NCE6X*** or ***BA0*** or sold in Sabre (1S) or Gallileo (1G).
	Pos *PointOfSaleInformationType `xml:"pos,omitempty"`

	// The repetition is 10 because we can transport: - until 5 tierLevel - until 5 customerValue, including possibly range of customerValue.  If we have tierLevel in the FTI, the customerValue must not be present. If we have customerValue in the FTI, the tierLevel must not be present.
	TierLevelAndCustomerValue *FrequentTravellerIdentificationCodeType `xml:"tierLevelAndCustomerValue,omitempty"`

	SortCriteria *SearchCriteria `xml:"sortCriteria,omitempty"`
}

package service

import (
	"errors"

	"github.com/tmconsulting/amadeus-golang-sdk/sdk"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/air/sellFromRecommendation/v05.2/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/air/sellFromRecommendation/v05.2/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ama/ticketIgnoreRefund/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ama/ticketInitRefund/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ama/ticketProcessRefund/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/commandCryptic"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/commandCryptic/v07.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/docIssuance/issueTicket/v09.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/checkRules/v07.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/checkRules/v07.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativeBestPricingWithoutPNR/v12.4/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativeBestPricingWithoutPNR/v12.4/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativePricingWithoutPNR/v12.4/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativePricingWithoutPNR/v12.4/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v14.3/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v14.3/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v16.3/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v16.3/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/pricePNRWithBookingClass/v14.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/pricePNRWithBookingClass/v14.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/addMultiElements/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/cancel/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/ignore/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/reply/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/retrieve/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/salesReports/displayQueryReport/v10.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/salesReports/displayQueryReport/v10.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/salesReports/displayTransactionReport/v13.2"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/security/signOut/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/session/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/cancelDocument/v11.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/createTSTFromPricing/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/deleteTST/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/displayTST/v07.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/displayTST/v07.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/processEDoc/v15.2/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/processEDoc/v15.2/response"
)

// Methods versions sdk realisation
const (
	PNRRetrieveV113 = iota
	TicketDisplayTSTV071
	FareInformativePricingWithoutPNRV124
	FareCheckRulesV071
	CommandCrypticV073
	SecuritySignOutV041
	UpdateSessionV030
	CloseSessionV041
	FareMasterPricerTravelBoardSearchV143
	FareMasterPricerTravelBoardSearchV163
	FareInformativeBestPricingWithoutPNRV124
	AirSellFromRecommendationV052
	PNRAddMultiElementsV113
	FarePricePNRWithBookingClassV141
	TicketCreateTSTFromPricingV041
	PNRCancelV113
	DocIssuanceIssueTicketV091
	SalesReportsDisplayQueryReportV101
	TicketCancelDocumentV111
	TicketDeleteTSTV041
	AMATicketInitRefundV030
	AMATicketIgnoreRefundV030
	AMATicketProcessRefundV030
	TicketProcessEDocV152
	SalesReportsDisplayTransactionReportV132
	PNRIgnoreV041
)

//MethodsMap Represents methods that have realisation in current version of sdk
type MethodsMap struct {
	PNRRetrieve                          int
	TicketDisplayTST                     int
	FareInformativePricingWithoutPNR     int
	FareCheckRules                       int
	CommandCryptic                       int
	SecuritySignOut                      int
	UpdateSession                        int
	CloseSession                         int
	FareMasterPricerTravelBoardSearch    int
	FareInformativeBestPricingWithoutPNR int
	AirSellFromRecommendation            int
	PNRAddMultiElements                  int
	FarePricePNRWithBookingClass         int
	TicketCreateTSTFromPricing           int
	PNRCancel                            int
	DocIssuanceIssueTicket               int
	SalesReportsDisplayQueryReport       int
	TicketCancelDocument                 int
	TicketDeleteTST                      int
	AMATicketInitRefund                  int
	AMATicketIgnoreRefund                int
	AMATicketProcessRefund               int
	TicketProcessEDoc                    int
	SalesReportsDisplayTransactionReport int
	PNRIgnore                            int
}

var ErrNoRealisation = errors.New("sorry, method has no realisation")

//GetLatest Generates latest actual methods map
func GetLatestMethodsMap() *MethodsMap {
	return &MethodsMap{
		PNRRetrieve:                          PNRRetrieveV113,
		TicketDisplayTST:                     TicketDisplayTSTV071,
		FareInformativePricingWithoutPNR:     FareInformativePricingWithoutPNRV124,
		FareCheckRules:                       FareCheckRulesV071,
		CommandCryptic:                       CommandCrypticV073,
		SecuritySignOut:                      SecuritySignOutV041,
		UpdateSession:                        UpdateSessionV030,
		CloseSession:                         CloseSessionV041,
		FareMasterPricerTravelBoardSearch:    FareMasterPricerTravelBoardSearchV143,
		FareInformativeBestPricingWithoutPNR: FareInformativeBestPricingWithoutPNRV124,
		AirSellFromRecommendation:            AirSellFromRecommendationV052,
		PNRAddMultiElements:                  PNRAddMultiElementsV113,
		FarePricePNRWithBookingClass:         FarePricePNRWithBookingClassV141,
		TicketCreateTSTFromPricing:           TicketCreateTSTFromPricingV041,
		PNRCancel:                            PNRCancelV113,
		DocIssuanceIssueTicket:               DocIssuanceIssueTicketV091,
		SalesReportsDisplayQueryReport:       SalesReportsDisplayQueryReportV101,
		TicketCancelDocument:                 TicketCancelDocumentV111,
		TicketDeleteTST:                      TicketDeleteTSTV041,
		AMATicketInitRefund:                  AMATicketInitRefundV030,
		AMATicketIgnoreRefund:                AMATicketIgnoreRefundV030,
		AMATicketProcessRefund:               AMATicketProcessRefundV030,
		TicketProcessEDoc:                    TicketProcessEDocV152,
		SalesReportsDisplayTransactionReport: SalesReportsDisplayTransactionReportV132,
		PNRIgnore:                            PNRIgnoreV041,
	}
}

func NewSKD(sdk AmadeusSDK, methodsMap *MethodsMap) Service {
	return &service{sdk: sdk, mm: methodsMap}
}

type Client struct {
	url     string
	user    string
	pass    string
	agent   string
	tls     bool
	headers []interface{}
}

type service struct {
	mm  *MethodsMap
	sdk AmadeusSDK
}

type AmadeusSDK interface {
	// Information
	PNRRetrieveV113(query *PNR_Retrieve_v11_3.Request) (*PNR_Reply_v11_3.Response, *sdk.ResponseSOAPHeader, error)
	TicketDisplayTSTV071(query *Ticket_DisplayTSTRequest_v07_1.Request) (*Ticket_DisplayTSTResponse_v07_1.Response, *sdk.ResponseSOAPHeader, error)
	FareInformativePricingWithoutPNRV124(query *Fare_InformativePricingWithoutPNR_v12_4.Request) (*Fare_InformativePricingWithoutPNRReply_v12_4.Response, *sdk.ResponseSOAPHeader, error)
	FareCheckRulesV071(query *Fare_CheckRulesRequest_v07_1.Request) (*Fare_CheckRulesResponse_v07_1.Response, *sdk.ResponseSOAPHeader, error)
	CommandCrypticV073(query *CommandCryptic_v07_3.Request) (*CommandCryptic_v07_3.Response, error)

	// Session
	SecuritySignOutV041() (*SecuritySignOut_v04_1.Response, *sdk.ResponseSOAPHeader, error)
	GetSession() *Session_v03_0.Session
	IncSessionSequenceNumber(header *sdk.ResponseSOAPHeader)
	CheckIfSessionIsClosed() bool
	SetSessionEndTransaction() bool
	UpdateSessionV030(session *Session_v03_0.Session) bool
	CloseSessionV041() (reply *SecuritySignOut_v04_1.Response, err error)

	// Search
	FareMasterPricerTravelBoardSearchV143(query *Fare_MasterPricerTravelBoardSearchRequest_v14_3.Request) (*Fare_MasterPricerTravelBoardSearchResponse_v14_3.Response, *sdk.ResponseSOAPHeader, error)
	FareMasterPricerTravelBoardSearchV163(query *Fare_MasterPricerTravelBoardSearchRequest_v16_3.Request) (*Fare_MasterPricerTravelBoardSearchResponse_v16_3.Response, *sdk.ResponseSOAPHeader, error)
	FareInformativeBestPricingWithoutPNRV124(query *Fare_InformativeBestPricingWithoutPNRRequest_v12_4.Request) (*Fare_InformativeBestPricingWithoutPNRResponse_v12_4.Response, *sdk.ResponseSOAPHeader, error)

	// Book
	AirSellFromRecommendationV052(query *Air_SellFromRecommendationRequest_v05_2.Request) (*Air_SellFromRecommendationResponse_v05_2.Response, *sdk.ResponseSOAPHeader, error)
	PNRAddMultiElementsV113(query *PNR_AddMultiElementsRequest_v11_3.Request) (*PNR_Reply_v11_3.Response, *sdk.ResponseSOAPHeader, error)
	FarePricePNRWithBookingClassV141(query *Fare_PricePNRWithBookingClassRequest_v14_1.Request) (*Fare_PricePNRWithBookingClassResponse_v14_1.Response, *sdk.ResponseSOAPHeader, error)
	TicketCreateTSTFromPricingV041(query *Ticket_CreateTSTFromPricing_v04_1.Request) (*Ticket_CreateTSTFromPricing_v04_1.Response, *sdk.ResponseSOAPHeader, error)
	PNRCancelV113(query *PNR_Cancel_v11_3.Request) (*PNR_Reply_v11_3.Response, *sdk.ResponseSOAPHeader, error)

	// Issue
	DocIssuanceIssueTicketV091(query *DocIssuance_IssueTicket_v09_1.Request) (*DocIssuance_IssueTicket_v09_1.Response, *sdk.ResponseSOAPHeader, error)

	// Void
	SalesReportsDisplayQueryReportV101(query *SalesReports_QueryReportRequest_v10_1.Request) (*SalesReports_QueryReportReply_v10_1.Response, *sdk.ResponseSOAPHeader, error)
	TicketCancelDocumentV111(query *Ticket_CancelDocument_v11_1.Request) (*Ticket_CancelDocument_v11_1.Response, *sdk.ResponseSOAPHeader, error)
	TicketDeleteTSTV041(query *Ticket_DeleteTST_v04_1.Request) (*Ticket_DeleteTST_v04_1.Response, *sdk.ResponseSOAPHeader, error)

	// Refund
	AMATicketInitRefundV030(query *AMA_TicketInitRefund_v03_0.Request) (*AMA_TicketInitRefund_v03_0.Response, *sdk.ResponseSOAPHeader, error)
	AMATicketIgnoreRefundV030(query *AMA_TicketIgnoreRefund_v03_0.Request) (*AMA_TicketIgnoreRefund_v03_0.Response, *sdk.ResponseSOAPHeader, error)
	AMATicketProcessRefundV030(query *AMA_TicketProcessRefund_v03_0.Request) (*AMA_TicketProcessRefund_v03_0.Response, *sdk.ResponseSOAPHeader, error)
	TicketProcessEDocV152(query *Ticket_ProcessEDocRequest_v15_2.Request) (*Ticket_ProcessEDocResponse_v15_2.Response, *sdk.ResponseSOAPHeader, error)
	SalesReportsDisplayTransactionReportV132(query *SalesReports_DisplayTransactionReport_v13_2.Request) (*SalesReports_DisplayTransactionReport_v13_2.Response, *sdk.ResponseSOAPHeader, error)
	PNRIgnoreV041(query *PNR_Ignore_v04_1.Request) (*PNR_Ignore_v04_1.Response, *sdk.ResponseSOAPHeader, error)
}

type Service interface {
	// Information
	PNRRetrieve(query *PNR_Retrieve_v11_3.Request) (*PNR_Reply_v11_3.Response, *sdk.ResponseSOAPHeader, error)
	TicketDisplayTST(query *Ticket_DisplayTSTRequest_v07_1.Request) (*Ticket_DisplayTSTResponse_v07_1.Response, *sdk.ResponseSOAPHeader, error)
	FareCheckRules(query *Fare_CheckRulesRequest_v07_1.Request) (*Fare_CheckRulesResponse_v07_1.Response, *sdk.ResponseSOAPHeader, error)
	CommandCryptic(msg string) (*commandCryptic.Response, error)

	// Session
	SecuritySignOut() (*SecuritySignOut_v04_1.Response, *sdk.ResponseSOAPHeader, error)
	Session() *Session_v03_0.Session
	IncSequenceNumber(header *sdk.ResponseSOAPHeader)
	IfSessionIsClosed() bool
	SessionEndTransaction() bool
	UpdateSession(session *Session_v03_0.Session) bool
	CloseSession() (reply *SecuritySignOut_v04_1.Response, err error)

	// Search
	FareMasterPricerTravelBoardSearch(query *Fare_MasterPricerTravelBoardSearchRequest_v14_3.Request) (*Fare_MasterPricerTravelBoardSearchResponse_v14_3.Response, *sdk.ResponseSOAPHeader, error)
	FareInformativeBestPricingWithout(query *Fare_InformativeBestPricingWithoutPNRRequest_v12_4.Request) (*Fare_InformativeBestPricingWithoutPNRResponse_v12_4.Response, *sdk.ResponseSOAPHeader, error)
	FareInformativePricingWithoutPNR(query *Fare_InformativePricingWithoutPNR_v12_4.Request) (*Fare_InformativePricingWithoutPNRReply_v12_4.Response, *sdk.ResponseSOAPHeader, error)

	// Book
	AirSellFromRecommendation(query *Air_SellFromRecommendationRequest_v05_2.Request) (*Air_SellFromRecommendationResponse_v05_2.Response, *sdk.ResponseSOAPHeader, error)
	PNRAddMultiElements(query *PNR_AddMultiElementsRequest_v11_3.Request) (*PNR_Reply_v11_3.Response, *sdk.ResponseSOAPHeader, error)
	FarePricePNRWithBookingClass(query *Fare_PricePNRWithBookingClassRequest_v14_1.Request) (*Fare_PricePNRWithBookingClassResponse_v14_1.Response, *sdk.ResponseSOAPHeader, error)
	TicketCreateTSTFromPricing(query *Ticket_CreateTSTFromPricing_v04_1.Request) (*Ticket_CreateTSTFromPricing_v04_1.Response, *sdk.ResponseSOAPHeader, error)

	// Cancellation
	PNRCancel(query *PNR_Cancel_v11_3.Request) (*PNR_Reply_v11_3.Response, *sdk.ResponseSOAPHeader, error)

	// Issue
	DocIssuanceIssueTicket(query *DocIssuance_IssueTicket_v09_1.Request) (*DocIssuance_IssueTicket_v09_1.Response, *sdk.ResponseSOAPHeader, error)

	// Void
	SalesReportsDisplayQueryReport(query *SalesReports_QueryReportRequest_v10_1.Request) (*SalesReports_QueryReportReply_v10_1.Response, *sdk.ResponseSOAPHeader, error)
	TicketCancelDocument(query *Ticket_CancelDocument_v11_1.Request) (*Ticket_CancelDocument_v11_1.Response, *sdk.ResponseSOAPHeader, error)
	TicketDeleteTST(query *Ticket_DeleteTST_v04_1.Request) (*Ticket_DeleteTST_v04_1.Response, *sdk.ResponseSOAPHeader, error)

	// Refund
	RefundIgnore(query *AMA_TicketIgnoreRefund_v03_0.Request) (*AMA_TicketIgnoreRefund_v03_0.Response, *sdk.ResponseSOAPHeader, error)
	RefundInit(query *AMA_TicketInitRefund_v03_0.Request) (*AMA_TicketInitRefund_v03_0.Response, *sdk.ResponseSOAPHeader, error)
	RefundProcess(query *AMA_TicketProcessRefund_v03_0.Request) (*AMA_TicketProcessRefund_v03_0.Response, *sdk.ResponseSOAPHeader, error)
	TicketProcessEDoc(query *Ticket_ProcessEDocRequest_v15_2.Request) (*Ticket_ProcessEDocResponse_v15_2.Response, *sdk.ResponseSOAPHeader, error)
	SalesReportsDisplayTransactionReport(query *SalesReports_DisplayTransactionReport_v13_2.Request) (*SalesReports_DisplayTransactionReport_v13_2.Response, *sdk.ResponseSOAPHeader, error)
	PNRIgnore(query *PNR_Ignore_v04_1.Request) (*PNR_Ignore_v04_1.Response, *sdk.ResponseSOAPHeader, error)
}

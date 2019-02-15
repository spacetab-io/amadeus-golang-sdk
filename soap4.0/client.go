package soap4_0

import (
	"bytes"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	soap "github.com/tmconsulting/amadeus-golang-sdk/soap2.0"
	"github.com/tmconsulting/amadeus-golang-sdk/utils"
)

var timeout = time.Duration(20 * time.Second)

const (
	WasNs    = "http://www.w3.org/2005/08/addressing"

	// Predefined WSS namespaces to be used in
	WssNsWSSE         = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU          = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsEncodingType = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary"
	WssNsPasswordType = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest"
)

type RequestSOAP4Envelope struct {
	XMLName        xml.Name                     `xml:"soap:Envelope"`

	SOAPAttr       string                       `xml:"xmlns:soap,attr"`
	XSIAttr        string                       `xml:"xmlns:xsi,attr"`
	XSDAttr        string                       `xml:"xmlns:xsd,attr"`

	Header         *RequsetSOAP4Header
	Body           soap.SOAPBody
}

type ResponseSOAP4Envelope struct {
	//XMLName        xml.Name                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header         ResponseSOAP4Header
	Body           soap.ResponseSOAPBody
}

type RequsetSOAP4Header struct {
	XMLName        xml.Name                     `xml:"soap:Header"`

	WSAAttr        string                       `xml:"xmlns:wsa,attr"`

	To             string                       `xml:"wsa:To"`
	Action         string                       `xml:"wsa:Action"`
	MessageId      string                       `xml:"wsa:MessageID"`

	Security       *WSSSecurityHeader           //`xml:"wsse:Security"`
	AMASecurity    *AMASecurityHostedUser       //`xml:"AMA_SecurityHostedUser"`
	Session        *RequestSession_v3           //`xml:"Session"`
}

type ResponseSOAP4Header struct {
	//XMLName        xml.Name                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ soap:Header"`

	To             string                       `xml:"To"`
	From           AddressingFrom               `xml:"From"`
	Action         string                       `xml:"Action"`
	MessageId      string                       `xml:"MessageID"`
	//RelatesTo      RelatesTo                    `xml:"RelatesTo"`
	RelatesTo      string                       `xml:"RelatesTo"`
	Session        Session_v3                   `xml:"Session"`

	Items   []interface{} `xml:",omitempty"`
}

type WSSSecurityHeader struct {
	XMLName        xml.Name                     `xml:"wsse:Security"`

	XmlNSWsse      string                       `xml:"xmlns:wsse,attr"`
	XmlNSWsu       string                       `xml:"xmlns:wsu,attr"`

	MustUnderstand string                       `xml:"mustUnderstand,attr,omitempty"`

	Token          *WSSUsernameToken            `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName        xml.Name                     `xml:"wsse:UsernameToken"`

	Id             string                       `xml:"wsu:Id,attr,omitempty"`

	Username       string                       `xml:"wsse:Username"`
	Nonce          *WSSNonce                    //`xml:"wsse:Nonce"`
	Password       *WSSPassword                 //`xml:"wsse:Password"`
	Created        string                       `xml:"wsu:Created"`
}

type WSSNonce struct {
	XMLName        xml.Name                     `xml:"wsse:Nonce"`

	EncodingType   string                       `xml:"EncodingType,attr"`
	Nonce          string                       `xml:",chardata"`
}

type WSSPassword struct {
	XMLName        xml.Name                     `xml:"wsse:Password"`

	Type           string                       `xml:"Type,attr"`
	Password       string                       `xml:",chardata"`
}

type AMASecurityHostedUser struct {
	XMLName        xml.Name                     `xml:"http://xml.amadeus.com/2010/06/Security_v1 AMA_SecurityHostedUser"`
	UserId         AMASecurityHostedUserUserID  //`xml:"UserID"`
}

type AMASecurityHostedUserUserID struct {
	XMLName        xml.Name                     `xml:"UserID"`

	AgentDutyCode  string                       `xml:"AgentDutyCode,attr"`
	RequestorType  string                       `xml:"RequestorType,attr"`
	PseudoCityCode string                       `xml:"PseudoCityCode,attr"`
	POSType        string                       `xml:"POS_Type,attr"`
}
type SOAPHeader struct {
	XMLName xml.Name      `xml:"soap:Header"`
	Items   []interface{} `xml:",omitempty"`
}

type AddressingFrom struct {
	//XMLName          xml.Name `xml:"wsa:From"`
	Address          string   `xml:"Address"`
}

//type RelatesTo struct {
//	XMLName          xml.Name `xml:"wsa:RelatesTo"`
//	RelationshipType string   `xml:"RelationshipType,attr"`			// "http://www.w3.org/2005/08/addressing/reply"
//	MessageId        string   `xml:",chardata"`
//}

func HashedPassword(password string, nonceRaw string, timestamp string) string {
	// HashedPassword = Base64 ( SHA-1 ( nonce + created + SHA-1 ( password ) ) )
	s := sha1.New()
	s.Write([]byte(password))
	pwd := string(s.Sum(nil))
	s.Reset()
	s.Write([]byte(nonceRaw + timestamp + pwd))
	return base64.StdEncoding.EncodeToString(s.Sum(nil))
}

func NewWSSSecurityHeader(user, pass, mustUnderstand string) *WSSSecurityHeader {
	nonce := utils.RandStringBytesMaskImprSrc(16)
	nonceBase64 := base64.StdEncoding.EncodeToString([]byte(nonce))
	timestamp := time.Now().UTC().Format(time.RFC3339)

	hdr := &WSSSecurityHeader{XmlNSWsse: WssNsWSSE, XmlNSWsu: WssNsWSU, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{Id: "UsernameToken-" + utils.RandStringBytesMaskImprSrc(9), Username: user}
	hdr.Token.Nonce = &WSSNonce{EncodingType: WssNsEncodingType, Nonce: nonceBase64}
	hdr.Token.Password = &WSSPassword{Type: WssNsPasswordType, Password: HashedPassword(pass, nonce, timestamp)}
	hdr.Token.Created = timestamp
	return hdr
}

func NewAMASecurityHostedUser(agent string) *AMASecurityHostedUser {
	return &AMASecurityHostedUser{UserId: AMASecurityHostedUserUserID{
		AgentDutyCode:  "SU",
		RequestorType:  "U",
		PseudoCityCode: agent,
		POSType:        "1",
	}}
}

type SOAP4Client struct {
	url     string
	user    string
	pass    string
	agent   string
	tls     bool
	headers []interface{}
}

type WebServicesPTSOAP4Header struct {
	client *SOAP4Client
	//wsap   string
}

func NewAmadeusWebServicesPTSOAP4Header(url, user, pass, agent string, tls bool) *WebServicesPTSOAP4Header {

	client := NewSOAP4Client(url, user, pass, agent, tls)

	return &WebServicesPTSOAP4Header{
		client: client,
		//wsap:   wsap,
	}
}

func (service *WebServicesPTSOAP4Header) Call(soapAction, messageId string, query, reply interface{}, session *Session_v3) (*ResponseSOAP4Header, error) {

	header, err := service.client.Call(soapAction, messageId, query, reply, session)
	if err != nil {
		return header, err
	}

	return header, nil

}

func (service *WebServicesPTSOAP4Header) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *WebServicesPTSOAP4Header) SetHeader(header interface{}) {
	service.client.UpdateHeader(header)
}

func (service *WebServicesPTSOAP4Header) UpdateHeader(header interface{}) {
	service.client.UpdateHeader(header)
}


func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

func NewSOAP4Client(url, user, pass, agent string, tls bool) *SOAP4Client {
	return &SOAP4Client{
		url:   url,
		user:  user,
		pass:  pass,
		agent: agent,
		tls:   tls,
	}
}

func (s *SOAP4Client) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

func (s *SOAP4Client) UpdateHeader(header interface{}) {
	s.headers = []interface{}{header}
}

func (s *SOAP4Client) Call(soapAction, messageId string, query, reply interface{}, session *Session_v3) (*ResponseSOAP4Header, error) {
	envelope := RequestSOAP4Envelope{SOAPAttr: soap.SoapNs, XSIAttr: soap.XsiNs, XSDAttr: soap.XsdNs}
	envelope.Header = &RequsetSOAP4Header{WSAAttr: WasNs, To: s.url, Action: soapAction, MessageId: messageId}
	if session == nil || session.TransactionStatusCode == TransactionStatusCode[Start] {
		envelope.Header.Security = NewWSSSecurityHeader(s.user, s.pass, "")
		envelope.Header.AMASecurity = NewAMASecurityHostedUser(s.agent)
	}
	if session != nil {
		envelope.Header.Session = &RequestSession_v3{Session_v3: *session}
	}

	envelope.Body.Content = query
	buffer := new(bytes.Buffer)
	buffer.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>"))

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return nil, err
	}

	if err := encoder.Flush(); err != nil {
		return nil, err
	}

	savebuf := buffer.Bytes()
	log.Println(string(savebuf))

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

	req.Header.Set("User-Agent", "connector_amadeus-go/0.2")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	numberOfAttempts := 3

	var res *http.Response
	client := &http.Client{
		Transport: tr,
		Timeout: timeout,
	}
	res, err = client.Do(req)
	for err != nil {
		time.Sleep(1 * time.Second)

		buffer := new(bytes.Buffer)
		buffer.Write(savebuf)
		rc := ioutil.NopCloser(buffer)
		req.Body = rc  //.(io.ReadCloser)

		client = &http.Client{
			Transport: tr,
			Timeout: timeout,
		}
		res, err = client.Do(req)

		numberOfAttempts--
		if numberOfAttempts == 0 {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil, nil
	}

	//log.Println(string(rawbody))
	var str = ""
	for _, ch := range rawbody {
		if ch != 0 && ch != '\r' {
			str += string(ch)
		}
	}
	log.Println(str)

	respEnvelope := new(ResponseSOAP4Envelope)
	respEnvelope.Body = soap.ResponseSOAPBody{Content: reply}

	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return nil, err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return &respEnvelope.Header, fault
	}

	return &respEnvelope.Header, nil
}

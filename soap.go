package amadeus

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

var timeout = time.Duration(30 * time.Second)

const (
	SoapNs    = "http://schemas.xmlsoap.org/soap/envelope/"
	XsiNs     = "http://www.w3.org/2001/XMLSchema-instance"
	XsdNs     = "http://www.w3.org/2001/XMLSchema"
)

type ResponseSOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  ResponseSOAPHeaderWithSession
	Body    ResponseSOAPBody
}
type SOAPEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`

	SOAPAttr string `xml:"xmlns:soap,attr"`
	XSIAttr  string `xml:"xmlns:xsi,attr"`
	XSDAttr  string `xml:"xmlns:xsd,attr"`

	Header SOAPHeaderWithSession
	Body   SOAPBody
}

type SOAPHeaderWithSession struct {
	XMLName xml.Name      `xml:"http://xml.amadeus.com/ws/2009/01/WBS_Session-2.0.xsd soap:Header"`
	Session *Session      `xml:"Session"`
	Items   []interface{} `xml:",omitempty"`
}
type ResponseSOAPHeaderWithSession struct {
	XMLName xml.Name      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Session *Session      `xml:",omitempty"`
	Items   []interface{} `xml:",omitempty"`
}
type SOAPHeader struct {
	XMLName xml.Name      `xml:"soap:Header"`
	Items   []interface{} `xml:",omitempty"`
}

type SOAPBody struct {
	XMLName xml.Name `xml:"soap:Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}
type ResponseSOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url     string
	tls     bool
	auth    *BasicAuth
	headers []interface{}
}

type WebServicesPT struct {
	client *SOAPClient
	wsap   string
}

func NewAmadeusWebServicesPT(url string, tls bool, WSAP string, auth *BasicAuth) *WebServicesPT {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &WebServicesPT{
		client: client,
		wsap:   WSAP,
	}
}

func (service *WebServicesPT) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *WebServicesPT) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

func (b *ResponseSOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}, session *Session) error {
	envelope := SOAPEnvelope{SOAPAttr: SoapNs, XSIAttr: XsiNs, XSDAttr: XsdNs}

	if s.headers != nil && len(s.headers) > 0 {
		soapHeader := SOAPHeaderWithSession{Items: make([]interface{}, len(s.headers))}
		copy(soapHeader.Items, s.headers)
		envelope.Header = soapHeader
	} else {
		soapHeader := SOAPHeaderWithSession{Session: &Session{}}
		envelope.Header = soapHeader
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)
	buffer.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>"))

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

	req.Header.Set("User-Agent", "connector_amadeus-go/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(ResponseSOAPEnvelope)
	respEnvelope.Header = ResponseSOAPHeaderWithSession{Session: session}
	respEnvelope.Body = ResponseSOAPBody{Content: response}

	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}

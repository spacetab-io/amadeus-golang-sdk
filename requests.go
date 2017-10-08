package amadeus

var url string = "http://webservices.amadeus.com"

func (service *WebServicesPT) SecurityAuthenticate(request *SecurityAuthenticate) (*SecurityAuthenticateReply, *Session, error) {
	soapAction := "VLSSLQ_06_1_1A"

	response := new(SecurityAuthenticateReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, request, response, session)
	if err != nil {
		return nil, nil, err
	}

	return response, session, nil
}

func (service *WebServicesPT) SecuritySignOut(request *SecuritySignOut) (*SecuritySignOutReply, error) {
	soapAction := "VLSSOQ_04_1_1A"

	response := new(SecuritySignOutReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, request, response, session)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *WebServicesPT) CommandCryptic(request *CommandCryptic) (*CommandCrypticReply, *Session, error) {
	soapAction := "HSFREQ_07_3_1A"

	response := new(CommandCrypticReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, request, response, session)
	if err != nil {
		return nil, nil, err
	}

	return response, session, nil
}

package wooba

import (
	"bytes"
	"encoding/xml"
	"errors"
	. "github.com/Panda-CRM/helpers-go"
	. "github.com/Panda-CRM/wooba-go/models"
	"golang.org/x/net/html/charset"
	"net/http"
	"time"
)

const (
	DATE_FORMAT = "2006-01-02T15:04:05"
)

var HEADERS = map[string]string{"Content-Type": "text/xml; charset=utf-8"}

type SoapEnvelope struct {
	Body SoapBody
}

type SoapBody struct {
	SoapResponse SoapResponse `xml:"VendasStringResponse"`
}

type SoapResponse struct {
	Result string `xml:"VendasStringResult"`
}

type Wooba struct {
	Url          string
	UserLogin    string
	UserPassword string
	soapAction   string
}

func (w Wooba) sendRequest(xmlRequest string) (string, error) {
	request, errNewRequest := http.NewRequest("POST", w.Url, bytes.NewBuffer([]byte(xmlRequest)))
	if errNewRequest != nil {
		LoggerErr(errNewRequest)
		return "", errNewRequest
	}
	//adiciona headers
	request.Header.Set("SOAPAction", w.soapAction)
	for hKey, hValue := range HEADERS {
		request.Header.Set(hKey, hValue)
	}
	//realiza conexão com o server
	client := &http.Client{}
	response, errClientDo := client.Do(request)
	if errClientDo != nil {
		LoggerErr(errClientDo)
		return "", errClientDo
	}
	//fecha conexão
	defer response.Body.Close()
	//faz o parse das informações para o objeto envelope
	var envelop SoapEnvelope
	errUnmarshalEnvelop := xml.NewDecoder(response.Body).Decode(&envelop)
	if errUnmarshalEnvelop != nil {
		LoggerErr(errUnmarshalEnvelop)
		return "", errUnmarshalEnvelop
	}
	//verifica se retornou erros no XML
	var salesErrors VendasErros
	reader := bytes.NewReader([]byte(envelop.Body.SoapResponse.Result))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoder.Decode(&salesErrors)
	if len(salesErrors.Erros) > 0 {
		return "", errors.New(salesErrors.Erros[0].Erro)
	}
	return envelop.Body.SoapResponse.Result, nil
}

func (w Wooba) VendasString(startDate, endDate time.Time) (VendasList, error) {
	w.soapAction = "http://tempuri.org/VendasString"
	xmlRequest := `<?xml version="1.0" encoding="utf-8"?>
		<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
			<soap:Body>
				<VendasString xmlns="http://tempuri.org/">
					<sUsuario>` + w.UserLogin + `</sUsuario>
					<sPassword>` + w.UserPassword + `</sPassword>
					<dInicio>` + startDate.Format(DATE_FORMAT) + `</dInicio>
					<dFinal>` + endDate.Format(DATE_FORMAT) + `</dFinal>
				</VendasString>
			</soap:Body>
		</soap:Envelope>`

	var sales VendasList
	//faz requisição para o webservice
	response, errRequest := w.sendRequest(xmlRequest)
	//verifica se houve erros na requisição
	if errRequest != nil {
		LoggerErr(errRequest)
		return sales, errRequest
	}
	//faz o parse das informações para o objeto de vendas
	reader := bytes.NewReader([]byte(response))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	errUnmarshal := decoder.Decode(&sales)
	if errUnmarshal != nil {
		LoggerErr(errUnmarshal)
		return sales, errUnmarshal
	}
	return sales, nil
}

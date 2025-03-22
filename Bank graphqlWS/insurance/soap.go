package insurance

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// SOAP API endpoint
	soapAPI = "http://localhost:8081/ws"
)

type CalculateInsuranceResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		CalculateInsuranceResponse struct {
			InsuranceAmount  float64 `xml:"insuranceAmount"`
			InsuranceCompany string  `xml:"insuranceCompany"`
		} `xml:"CalculateInsuranceResponse"`
	} `xml:"Body"`
}

func InsuranceDetails(purchaseID int32) (float64, string, error) {
	// Construct the SOAP request
	soapRequest := fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://insurancesoap.web">
	   <soapenv:Header/>
	   <soapenv:Body>
	      <web:CalculateInsuranceRequest>
	         <web:purchaseId>%d</web:purchaseId>
	      </web:CalculateInsuranceRequest>
	   </soapenv:Body>
	</soapenv:Envelope>`, purchaseID)

	// Send the SOAP request
	resp, err := http.Post(soapAPI, "text/xml", bytes.NewBufferString(soapRequest))
	if err != nil {
		return 0, "", fmt.Errorf("failed to call SOAP API: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", fmt.Errorf("failed to read SOAP response: %v", err)
	}

	// Log the raw SOAP response for debugging
	fmt.Println("Raw SOAP Response:", string(body))

	var soapResponse CalculateInsuranceResponse
	if err := xml.Unmarshal(body, &soapResponse); err != nil {
		return 0, "", fmt.Errorf("failed to parse SOAP response: %v", err)
	}

	insuranceAmount := soapResponse.Body.CalculateInsuranceResponse.InsuranceAmount
	insuranceCompany := soapResponse.Body.CalculateInsuranceResponse.InsuranceCompany

	return insuranceAmount, insuranceCompany, nil
}

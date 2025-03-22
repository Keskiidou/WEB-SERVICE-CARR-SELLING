package model

import "encoding/xml"

type CalculateInsuranceResponse struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    struct {
		CalculateInsuranceResponse struct {
			InsuranceAmount  float64 `xml:"insuranceAmount"`
			InsuranceCompany string  `xml:"insuranceCompany"`
		} `xml:"CalculateInsuranceResponse"`
	} `xml:"Body"`
}

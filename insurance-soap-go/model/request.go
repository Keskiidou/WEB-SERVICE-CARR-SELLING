package model

import "encoding/xml"

type CalculateInsuranceRequest struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    struct {
		CalculateInsurance struct {
			PurchaseID int32 `xml:"purchaseId"`
		} `xml:"CalculateInsuranceRequest"`
	} `xml:"Body"`
}

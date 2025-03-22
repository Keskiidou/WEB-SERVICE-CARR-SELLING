package handlers

import (
	"encoding/xml"
	"insurance-soap-go/model"
	"insurance-soap-go/services"
	"net/http"
)

type InsuranceHandler struct {
	insuranceService *services.InsuranceService
}

func NewInsuranceHandler(insuranceService *services.InsuranceService) *InsuranceHandler {
	return &InsuranceHandler{insuranceService: insuranceService}
}

func (h *InsuranceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming XML request
	var request model.CalculateInsuranceRequest
	err := xml.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid XML request", http.StatusBadRequest)
		return
	}

	// Process the request
	purchaseID := request.Body.CalculateInsurance.PurchaseID
	insuranceAmount, insuranceCompany, err := h.insuranceService.CalculateInsurance(purchaseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate the XML response
	response := model.CalculateInsuranceResponse{}
	response.Body.CalculateInsuranceResponse.InsuranceAmount = insuranceAmount
	response.Body.CalculateInsuranceResponse.InsuranceCompany = insuranceCompany

	// Write the XML response
	w.Header().Set("Content-Type", "text/xml")
	xml.NewEncoder(w).Encode(response)
}

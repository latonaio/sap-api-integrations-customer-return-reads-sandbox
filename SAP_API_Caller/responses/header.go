package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			CustomerReturn                string `json:"CustomerReturn"`
			CustomerReturnType            string `json:"CustomerReturnType"`
			SalesOrganization             string `json:"SalesOrganization"`
			DistributionChannel           string `json:"DistributionChannel"`
			OrganizationDivision          string `json:"OrganizationDivision"`
			SalesGroup                    string `json:"SalesGroup"`
			SalesOffice                   string `json:"SalesOffice"`
			SalesDistrict                 string `json:"SalesDistrict"`
			SoldToParty                   string `json:"SoldToParty"`
			CreationDate                  string `json:"CreationDate"`
			CreatedByUser                 string `json:"CreatedByUser"`
			LastChangeDate                string `json:"LastChangeDate"`
			SenderBusinessSystemName      string `json:"SenderBusinessSystemName"`
			LastChangeDateTime            string `json:"LastChangeDateTime"`
			PurchaseOrderByCustomer       string `json:"PurchaseOrderByCustomer"`
			CustomerPurchaseOrderType     string `json:"CustomerPurchaseOrderType"`
			CustomerPurchaseOrderDate     string `json:"CustomerPurchaseOrderDate"`
			CustomerReturnDate            string `json:"CustomerReturnDate"`
			TotalNetAmount                string `json:"TotalNetAmount"`
			TransactionCurrency           string `json:"TransactionCurrency"`
			SDDocumentReason              string `json:"SDDocumentReason"`
			PricingDate                   string `json:"PricingDate"`
			RequestedDeliveryDate         string `json:"RequestedDeliveryDate"`
			ShippingType                  string `json:"ShippingType"`
			HeaderBillingBlockReason      string `json:"HeaderBillingBlockReason"`
			DeliveryBlockReason           string `json:"DeliveryBlockReason"`
			IncotermsClassification       string `json:"IncotermsClassification"`
			IncotermsTransferLocation     string `json:"IncotermsTransferLocation"`
			IncotermsLocation1            string `json:"IncotermsLocation1"`
			IncotermsLocation2            string `json:"IncotermsLocation2"`
			IncotermsVersion              string `json:"IncotermsVersion"`
			CustomerPaymentTerms          string `json:"CustomerPaymentTerms"`
			PaymentMethod                 string `json:"PaymentMethod"`
			CustomerTaxClassification1    string `json:"CustomerTaxClassification1"`
			CustomerTaxClassification2    string `json:"CustomerTaxClassification2"`
			CustomerTaxClassification3    string `json:"CustomerTaxClassification3"`
			CustomerTaxClassification4    string `json:"CustomerTaxClassification4"`
			CustomerTaxClassification5    string `json:"CustomerTaxClassification5"`
			CustomerTaxClassification6    string `json:"CustomerTaxClassification6"`
			CustomerTaxClassification7    string `json:"CustomerTaxClassification7"`
			CustomerTaxClassification8    string `json:"CustomerTaxClassification8"`
			CustomerTaxClassification9    string `json:"CustomerTaxClassification9"`
			RetsMgmtProcess               string `json:"RetsMgmtProcess"`
			ReferenceSDDocument           string `json:"ReferenceSDDocument"`
			ReferenceSDDocumentCategory   string `json:"ReferenceSDDocumentCategory"`
			CustomerReturnApprovalReason  string `json:"CustomerReturnApprovalReason"`
			SalesDocApprovalStatus        string `json:"SalesDocApprovalStatus"`
			RetsMgmtLogProcgStatus        string `json:"RetsMgmtLogProcgStatus"`
			RetsMgmtCompnProcgStatus      string `json:"RetsMgmtCompnProcgStatus"`
			RetsMgmtProcessingStatus      string `json:"RetsMgmtProcessingStatus"`
			OverallSDProcessStatus        string `json:"OverallSDProcessStatus"`
			TotalCreditCheckStatus        string `json:"TotalCreditCheckStatus"`
			OverallSDDocumentRejectionSts string `json:"OverallSDDocumentRejectionSts"`
			ToItem                        struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Item"`
			ToHeaderPartner struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Partner"`
		} `json:"results"`
	} `json:"d"`
}

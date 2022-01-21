package main

import (
	sap_api_caller "sap-api-integrations-customer-return-reads/SAP_API_Caller"
	"sap-api-integrations-customer-return-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Customer_Return_Item_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "Item",
		}
	}

	caller.AsyncGetCustomerReturn(
		inoutSDC.CustomerReturn.CustomerReturn,
		inoutSDC.CustomerReturn.CustomerReturnItem.CustomerReturnItem,
		accepter,
	)
}
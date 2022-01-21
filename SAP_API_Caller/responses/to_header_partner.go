package responses

type ToHeaderPartner struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			CustomerReturn               string `json:"CustomerReturn"`
			PartnerFunction              string `json:"PartnerFunction"`
			Customer                     string `json:"Customer"`
			Supplier                     string `json:"Supplier"`
			Personnel                    string `json:"Personnel"`
			ContactPerson                string `json:"ContactPerson"`
			BusinessPartnerAddressUUID   string `json:"BusinessPartnerAddressUUID"`
			BPRefAddrForDocSpcfcAddrUUID string `json:"BPRefAddrForDocSpcfcAddrUUID"`
		} `json:"results"`
	} `json:"d"`
}
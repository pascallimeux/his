package ocms

import "fmt"

// swagger:model consent
type Consent struct {
	// required: true
	AppID 		string     `json:"appid"`
	// required: true
	ConsentID      	string     `json:"consentid"`
	// required: true
	OwnerID       	string     `json:"ownerid"`
	// required: true
	ConsumerID      string     `json:"consumerid"`
	DataType      	string     `json:"datatype"`
	DataAccess      string     `json:"dataaccess"`
	Dt_begin      	string     `json:"dtbegin"`
	Dt_end       	string     `json:"dtend"`
}

func (ch *Consent) ToString() string {
	consentStr := fmt.Sprintf("ConsentID:%s ConsumerID:%s OwnerID:%s Datatype:%s Dataaccess:%s Dt_begin:%s Dt_end:%s", ch.ConsentID, ch.ConsumerID, ch.OwnerID, ch.DataType, ch.DataAccess, ch.Dt_begin, ch.Dt_end)
	return consentStr
}

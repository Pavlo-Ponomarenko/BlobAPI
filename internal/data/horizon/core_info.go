package horizon

import (
	"encoding/json"
	"net/http"
)

type coreInfoResponse struct {
	Data coreInfo `json:"data"`
}

type coreInfo struct {
	Id         string             `json:"id"`
	Type       string             `json:"type"`
	Attributes coreInfoAttributes `json:"attributes"`
}

type coreInfoAttributes struct {
	Core               core          `json:"core"`
	CoreVersion        string        `json:"core_version"`
	CurrentTime        string        `json:"current_time"`
	CurrentTimeUnix    int64         `json:"current_time_unix"`
	EnvironmentName    string        `json:"environment_name"`
	History            coreHistory   `json:"history"`
	HistoryV2          coreHistoryV2 `json:"history_v2"`
	MasterAccountId    string        `json:"master_account_id"`
	NetworkPassphrase  string        `json:"network_passphrase"`
	Precision          int64         `json:"precision"`
	TxExpirationPeriod int64         `json:"tx_expiration_period"`
	XdrRevision        string        `json:"xdr_revision"`
}

type core struct {
	LastLedgerIncreaseTime string `json:"last_ledger_increase_time"`
	Latest                 int64  `json:"latest"`
	OldestOnStart          int64  `json:"oldest_on_start"`
}

type coreHistory struct {
	LastLedgerIncreaseTime string `json:"last_ledger_increase_time"`
	Latest                 int64  `json:"latest"`
	OldestOnStart          int64  `json:"oldest_on_start"`
}

type coreHistoryV2 struct {
	LastLedgerIncreaseTime string `json:"last_ledger_increase_time"`
	Latest                 int64  `json:"latest"`
	OldestOnStart          int64  `json:"oldest_on_start"`
}

type coreInfoCash struct {
	isLoaded bool
	coreInfo *coreInfoResponse
}

func (c *coreInfoCash) getInfo(coreInfoURL string) (*coreInfo, error) {
	if c.isLoaded {
		return &c.coreInfo.Data, nil
	}
	response, err := http.Get(coreInfoURL)
	if err != nil {
		return nil, err
	}
	info := new(coreInfoResponse)
	err = json.NewDecoder(response.Body).Decode(info)
	if err != nil {
		return nil, err
	}
	c.isLoaded = true
	c.coreInfo = info
	return &c.coreInfo.Data, nil
}

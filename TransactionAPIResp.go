package main

type TransactionTraceAPIResponse struct {
	Status  string             `json:"status" xml:"status"`
	Message string             `json:"message" xml:"message"`
	Result  []DataAPIRessponse `json:"result" xml:"result"`
}

//blockHash:0x7f18e6751c463f173f7308a4031afe3c434e2befefcdef250b114e22722340d0
//blockNumber:12063778
//confirmations:89
//contractAddress:
//cumulativeGasUsed:12465306
//from:0xa9066ba976c0cdb256b6ceb3f2bc5f826f3bb9e6
//gas:21000
//gasPrice:146400000000
//gasUsed:21000
//hash:0x0eff40dcb10f70104b92fdf200518db20e241a4ba6847d74e069d403122be8e6
//input:0x
//isError:0
//nonce:0
//timeStamp:1616085215
//to:0xd57aee8bf32d0e4c0a7e64ab4d2cb81e96a96f2c
//transactionIndex:181
//txreceipt_status:1
//value:32206196042394786

type DataAPIRessponse struct {
	BlockHash         string `json:"blockHash" xml:"blockHash"`
	BlockNumber       string `json:"blockNumber" xml:"blockNumber"`
	Confirmations     string `json:"confirmations" xml:"confirmations"`
	ContractAddress   string `json:"contractAddress" xml:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed" xml:"cumulativeGasUsed"`
	From              string `json:"from" xml:"from"`
	Gas               string `json:"gas" xml:"gas"`
	GasPrice          string `json:"gasPrice" xml:"gasPrice"`
	GasUsed           string `json:"gasUsed" xml:"gasUsed"`
	Hash              string `json:"hash" xml:"hash"`
	Input             string `json:"input" xml:"input"`
	IsError           string `json:"isError" xml:"isError"`
	Nonce             string `json:"nonce" xml:"nonce"`
	TimeStamp         string `json:"timeStamp" xml:"timeStamp"`
	To                string `json:"to" xml:"to"`
	TransactionIndex  string `json:"transactionIndex" xml:"transactionIndex"`
	Txreceipt_status  string `json:"txreceipt_status" xml:"txreceipt_status"`
	Value             string `json:"value" xml:"value"`
}
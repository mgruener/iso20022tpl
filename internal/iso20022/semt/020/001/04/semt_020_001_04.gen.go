// Code generated by main. DO NOT EDIT.

package semt_020_001_04

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesMsgCxlAdvc SecuritiesMessageCancellationAdviceV04 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SctiesMsgCxlAdvc"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SchmeNm,omitempty"`
}

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 PrtryId"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type References37Choice struct {
	SctiesSttlmTxConfId           SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SctiesSttlmTxConfId"`
	IntraPosMvmntConfId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 IntraPosMvmntConfId"`
	SctiesBalAcctgRptId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SctiesBalAcctgRptId"`
	SctiesBalCtdyRptId            Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SctiesBalCtdyRptId"`
	IntraPosMvmntPstngRptId       Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 IntraPosMvmntPstngRptId"`
	SctiesFincgConfId             SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SctiesFincgConfId"`
	SctiesTxPdgRptId              Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SctiesTxPdgRptId"`
	SctiesTxPstngRptId            Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SctiesTxPstngRptId"`
	SctiesSttlmTxAllgmtRptId      Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SctiesSttlmTxAllgmtRptId"`
	SctiesSttlmTxAllgmtNtfctnTxId SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SctiesSttlmTxAllgmtNtfctnTxId"`
	PrtflTrfNtfctnId              Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 PrtflTrfNtfctnId"`
	SctiesSttlmTxGnrtnNtfctnId    SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SctiesSttlmTxGnrtnNtfctnId"`
	OthrMsgId                     Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 OthrMsgId"`
	TtlPrtflValtnRptId            Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 TtlPrtflValtnRptId"`
}

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 Nm,omitempty"`
}

type SecuritiesMessageCancellationAdviceV04 struct {
	Ref         References37Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 Ref"`
	AcctOwnr    PartyIdentification36Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SfkpgAcct"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SplmtryData,omitempty"`
}

type SettlementTypeAndIdentification13 struct {
	TxId          Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 TxId"`
	SctiesMvmntTp ReceiveDelivery1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 SctiesMvmntTp"`
	Pmt           DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 Pmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

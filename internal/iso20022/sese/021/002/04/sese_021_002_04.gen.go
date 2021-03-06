// Code generated by main. DO NOT EDIT.

package sese_021_002_04

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type Document struct {
	SctiesTxStsQry SecuritiesTransactionStatusQuery002V04 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 SctiesTxStsQry"`
}

type DocumentNumber15 struct {
	Nb   DocumentNumber6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Nb"`
	Refs []Identification24    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Refs"`
}

type DocumentNumber6Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 LngNb"`
	PrtryNb GenericIdentification86           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 SchmeNm,omitempty"`
}

type GenericIdentification84 struct {
	Id      RestrictedFINXMax34Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 SchmeNm,omitempty"`
}

type GenericIdentification86 struct {
	Id      RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 SchmeNm,omitempty"`
}

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

type Identification24 struct {
	AcctOwnrTxId      RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 AcctOwnrTxId"`
	AcctSvcrTxId      RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 PrcrTxId,omitempty"`
	CmonId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 CmonId,omitempty"`
	TradId            []RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 TradId,omitempty"`
	MstrId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 MstrId,omitempty"`
	BsktId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 BsktId,omitempty"`
	IndxId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 IndxId,omitempty"`
	ListId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 ListId,omitempty"`
	PrgmId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 PrgmId,omitempty"`
	PoolId            RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 PoolId,omitempty"`
	CorpActnEvtId     RestrictedFINXMax16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 CorpActnEvtId,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max350Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

type PartyIdentification109 struct {
	Id  PartyIdentification114Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 LEI,omitempty"`
}

type PartyIdentification114Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 AnyBIC"`
	PrtryId GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 PrtryId"`
}

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax16Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax30Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax34Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,35}
type RestrictedFINXMax35Text string

type SecuritiesAccount27 struct {
	Id RestrictedFINXMax35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Id"`
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Nm,omitempty"`
}

type SecuritiesTransactionStatusQuery002V04 struct {
	StsAdvcReqd DocumentNumber15       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 StsAdvcReqd"`
	AcctOwnr    PartyIdentification109 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount27    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 SfkpgAcct"`
	SplmtryData []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 SplmtryData,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

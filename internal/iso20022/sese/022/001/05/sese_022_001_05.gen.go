// Code generated by main. DO NOT EDIT.

package sese_022_001_05

type AcknowledgedAcceptedStatus24Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason12 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Rsn"`
}

type AcknowledgementReason12 struct {
	Cd          AcknowledgementReason15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Cd"`
	AddtlRsnInf Max210Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 AddtlRsnInf,omitempty"`
}

type AcknowledgementReason15Choice struct {
	Cd    AcknowledgementReason3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Prtry"`
}

// May be one of ADEA, SMPG, OTHR
type AcknowledgementReason3Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type Document struct {
	SctiesStsOrStmtQryStsAdvc SecuritiesStatusOrStatementQueryStatusAdviceV05 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 SctiesStsOrStmtQryStsAdvc"`
}

type DocumentIdentification30 struct {
	MsgNb DocumentNumber5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 MsgNb,omitempty"`
	Ref   Identification14      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Ref"`
}

type DocumentNumber13 struct {
	Nb DocumentNumber5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Nb"`
}

type DocumentNumber16 struct {
	Nb   DocumentNumber5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Nb"`
	Refs []Identification26    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Refs"`
}

type DocumentNumber5Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 LngNb"`
	PrtryNb GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 SchmeNm,omitempty"`
}

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

type Identification14 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Id"`
}

type Identification26 struct {
	AcctOwnrTxId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 AcctOwnrTxId"`
	AcctSvcrTxId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 PrcrTxId,omitempty"`
	CmonId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 CmonId,omitempty"`
	TradId            []Max52Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 TradId,omitempty"`
	MstrId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 MstrId,omitempty"`
	BsktId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 BsktId,omitempty"`
	IndxId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 IndxId,omitempty"`
	ListId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 ListId,omitempty"`
	PrgmId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 PrgmId,omitempty"`
	PoolId            Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 PoolId,omitempty"`
	CorpActnEvtId     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 CorpActnEvtId,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max52Text string

// Must be at least 1 items long
type Max70Text string

// May be one of NORE
type NoReasonCode string

type PartyIdentification127Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 PrtryId"`
}

type PartyIdentification144 struct {
	Id  PartyIdentification127Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 LEI,omitempty"`
}

type ProcessingStatus55Choice struct {
	AckdAccptd AcknowledgedAcceptedStatus24Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 AckdAccptd"`
	Rjctd      RejectionOrRepairStatus32Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Rjctd"`
	Prtry      ProprietaryStatusAndReason6        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Prtry"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 PrtryRsn,omitempty"`
}

type RejectionAndRepairReason26Choice struct {
	Cd    RejectionReason24Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Prtry"`
}

type RejectionOrRepairReason26 struct {
	Cd          RejectionAndRepairReason26Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Cd"`
	AddtlRsnInf Max210Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairStatus32Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason26 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Rsn"`
}

// May be one of SAFE, DSEC, LATE, REFE, ADEA, OTHR, MISM
type RejectionReason24Code string

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Nm,omitempty"`
}

type SecuritiesStatusOrStatementQueryStatusAdviceV05 struct {
	QryDtls       DocumentIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 QryDtls"`
	AcctOwnr      PartyIdentification144   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 AcctOwnr,omitempty"`
	SfkpgAcct     SecuritiesAccount19      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 SfkpgAcct,omitempty"`
	StsOrStmtReqd StatusOrStatement9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 StsOrStmtReqd,omitempty"`
	PrcgSts       ProcessingStatus55Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 PrcgSts"`
	SplmtryData   []SupplementaryData1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 SplmtryData,omitempty"`
}

type StatusOrStatement9Choice struct {
	StsAdvc DocumentNumber16 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 StsAdvc"`
	Stmt    DocumentNumber13 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Stmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
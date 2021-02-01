// Code generated by main. DO NOT EDIT.

package sese_027_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AcknowledgedAcceptedStatus12Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Rsn"`
}

type AcknowledgementReason1 struct {
	Cd          AcknowledgementReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	AddtlRsnInf Max210Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AddtlRsnInf,omitempty"`
}

type AcknowledgementReason1Choice struct {
	Cd    AcknowledgementReason3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Prtry"`
}

// May be one of ADEA, SMPG, OTHR
type AcknowledgementReason3Code string

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AmountAndDirection8 struct {
	Amt                 ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 CdtDbtInd"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 OrgnlCcyAndOrdrdAmt,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CancellationReason2 struct {
	Cd          CancellationReason3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	AddtlRsnInf Max210Text                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AddtlRsnInf,omitempty"`
}

type CancellationReason3Choice struct {
	Cd    CancelledStatusReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Prtry"`
}

type CancellationStatus9Choice struct {
	NoSpcfdRsn NoReasonCode          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 NoSpcfdRsn"`
	Rsn        []CancellationReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Rsn"`
}

// May be one of CANI, OTHR
type CancelledStatusReason5Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 DtTm"`
}

// May be one of VARI
type DateType3Code string

// May be one of OPEN, UKWN
type DateType4Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type DeniedReason2 struct {
	Cd          DeniedReason2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	AddtlRsnInf Max210Text          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AddtlRsnInf,omitempty"`
}

type DeniedReason2Choice struct {
	Cd    DeniedReason4Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Prtry"`
}

// May be one of ADEA, DCAN, DPRG, DREP, DSET, LATE, OTHR, CDRG, CDCY, CDRE
type DeniedReason4Code string

type DeniedStatus6Choice struct {
	NoSpcfdRsn NoReasonCode    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 NoSpcfdRsn"`
	Rsn        []DeniedReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Rsn"`
}

type Document struct {
	SctiesTxCxlReqStsAdvc SecuritiesTransactionCancellationRequestStatusAdviceV03 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 SctiesTxCxlReqStsAdvc"`
}

type DocumentNumber1Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 LngNb"`
	PrtryNb GenericIdentification19           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AmtsdVal"`
}

type GenericDocumentIdentification1 struct {
	MsgNb DocumentNumber1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 MsgNb,omitempty"`
	Id    Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Id"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type Identification1 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Id"`
}

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Prtry"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Adr,omitempty"`
}

// May be one of NORE
type NoReasonCode string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Tp"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PrtryId"`
}

type PartyIdentification37Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 NmAndAdr"`
	Ctry     CountryCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Ctry"`
}

type PartyIdentification44Choice struct {
	AnyBIC   AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AnyBIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Ctry"`
}

type PartyIdentification45Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 NmAndAdr"`
}

type PartyIdentification46 struct {
	Id     PartyIdentification44Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Id"`
	PrcgId Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PrcgId,omitempty"`
}

type PartyIdentificationAndAccount44 struct {
	Id        PartyIdentification45Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Id"`
	SfkpgAcct SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 SfkpgAcct,omitempty"`
	PrcgId    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PrcgId,omitempty"`
}

type PendingReason15Choice struct {
	Cd    PendingReason9Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Prtry"`
}

type PendingReason7 struct {
	Cd          PendingReason15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, CONF, OTHR, CDRG, CDCY, CDRE, CDAC, INBC
type PendingReason9Code string

type PendingStatus11Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 NoSpcfdRsn"`
	Rsn        []PendingReason7 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Rsn"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Ctry"`
}

type ProcessingStatus20Choice struct {
	PdgCxl     PendingStatus11Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PdgCxl"`
	Rjctd      RejectionOrRepairStatus25Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Rjctd"`
	Rpr        RejectionOrRepairStatus14Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Rpr"`
	AckdAccptd AcknowledgedAcceptedStatus12Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AckdAccptd"`
	Prtry      ProprietaryStatusAndReason1        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Prtry"`
	Dnd        DeniedStatus6Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Dnd"`
	Canc       CancellationStatus9Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Canc"`
}

type ProprietaryReason1 struct {
	Rsn         GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason1 struct {
	PrtrySts GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PrtrySts"`
	PrtryRsn []ProprietaryReason1    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PrtryRsn,omitempty"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 OrgnlAndCurFace"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type References22Choice struct {
	OthrTxId        GenericDocumentIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 OthrTxId"`
	SctiesFincgTxId SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 SctiesFincgTxId"`
	SctiesSttlmTxId SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 SctiesSttlmTxId"`
	IntraPosMvmntId Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 IntraPosMvmntId"`
}

type RejectionAndRepairReason13Choice struct {
	Cd    RejectionReason27Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Prtry"`
}

type RejectionAndRepairReason18Choice struct {
	Cd    RejectionReason32Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Prtry"`
}

type RejectionOrRepairReason13 struct {
	Cd          RejectionAndRepairReason13Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	AddtlRsnInf Max210Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairReason18 struct {
	Cd          RejectionAndRepairReason18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	AddtlRsnInf Max210Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairStatus14Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason13 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Rsn"`
}

type RejectionOrRepairStatus25Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason18 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Rsn"`
}

// May be one of ADEA, LATE, SAFE, NRGM, NRGN, OTHR, REFE, INVM, INVL
type RejectionReason27Code string

// May be one of SAFE, ADEA, LATE, NRGN, REFE, NRGM, OTHR
type RejectionReason32Code string

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Nm,omitempty"`
}

type SecuritiesTransactionCancellationRequestStatusAdviceV03 struct {
	CxlReqRef   Identification1              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 CxlReqRef"`
	TxId        TransactionIdentifications17 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 TxId,omitempty"`
	PrcgSts     ProcessingStatus20Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PrcgSts"`
	TxDtls      TransactionDetails30         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 TxDtls,omitempty"`
	SplmtryData []SupplementaryData1         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 SplmtryData,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Desc,omitempty"`
}

type SettlementDate2Choice struct {
	Dt   DateAndDateTimeChoice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Dt"`
	DtCd SettlementDateCode2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 DtCd"`
}

type SettlementDateCode2Choice struct {
	Cd    DateType4Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Prtry"`
}

type SettlementParties13 struct {
	Dpstry PartyIdentification46           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Pty5,omitempty"`
}

type SettlementTypeAndIdentification13 struct {
	TxId          Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 TxId"`
	SctiesMvmntTp ReceiveDelivery1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 SctiesMvmntTp"`
	Pmt           DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Pmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeDate1Choice struct {
	Dt   DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Dt"`
	DtCd TradeDateCode1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 DtCd"`
}

type TradeDateCode1Choice struct {
	Cd    DateType3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Prtry"`
}

type TransactionDetails30 struct {
	AcctOwnr        PartyIdentification36Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AcctOwnr,omitempty"`
	SfkpgAcct       SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 SfkpgAcct"`
	FinInstrmId     SecurityIdentification14    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 FinInstrmId"`
	SttlmQty        Quantity6Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 SttlmQty"`
	SttlmAmt        AmountAndDirection8         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 SttlmAmt,omitempty"`
	TradDt          TradeDate1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 TradDt,omitempty"`
	SttlmDt         SettlementDate2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 SttlmDt"`
	DlvrgSttlmPties SettlementParties13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties  SettlementParties13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 RcvgSttlmPties,omitempty"`
	Invstr          PartyIdentification37Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Invstr,omitempty"`
}

type TransactionIdentifications17 struct {
	AcctSvcrTxId      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 PrcrTxId,omitempty"`
	AcctOwnrTxId      References22Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 AcctOwnrTxId"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
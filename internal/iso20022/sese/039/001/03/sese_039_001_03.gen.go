// Code generated by main. DO NOT EDIT.

package sese_039_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AcknowledgedAcceptedStatus14Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason7 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Rsn"`
}

// May be one of ADEA, SMPG, OTHR, NSTP, LATE
type AcknowledgementReason6Code string

type AcknowledgementReason7 struct {
	Cd          AcknowledgementReason9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	AddtlRsnInf Max210Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AddtlRsnInf,omitempty"`
}

type AcknowledgementReason9Choice struct {
	Cd    AcknowledgementReason6Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Prtry"`
}

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
	Amt                 ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 CdtDbtInd"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 OrgnlCcyAndOrdrdAmt,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 DtTm"`
}

// May be one of VARI
type DateType3Code string

// May be one of OPEN, UKWN
type DateType4Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type DeniedReason5 struct {
	Cd          DeniedReason7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	AddtlRsnInf Max210Text          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, CDCY, CDRE, DCAN, DSET, DPRG, DREP, LATE, OTHR, CDRG
type DeniedReason6Code string

type DeniedReason7Choice struct {
	Cd    DeniedReason6Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	Prtry GenericIdentification40 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Prtry"`
}

type DeniedStatus10Choice struct {
	NoSpcfdRsn NoReasonCode    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 NoSpcfdRsn"`
	Rsn        []DeniedReason5 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Rsn"`
}

type Document struct {
	SctiesSttlmTxModReqStsAdvc SecuritiesSettlementTransactionModificationRequestStatusAdviceV03 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 SctiesSttlmTxModReqStsAdvc"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AmtsdVal"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 SchmeNm,omitempty"`
}

type GenericIdentification40 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

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
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Id"`
}

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Prtry"`
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

type ModificationProcessingStatus4Choice struct {
	AckdAccptd AcknowledgedAcceptedStatus14Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AckdAccptd"`
	PdgPrcg    PendingProcessingStatus7Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 PdgPrcg"`
	Dnd        DeniedStatus10Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Dnd"`
	Rjctd      RejectionStatus8Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Rjctd"`
	Rprd       RepairStatus8Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Rprd"`
	Modfd      ModificationStatus2Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Modfd"`
	Prtry      ProprietaryStatusAndReason1        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Prtry"`
}

type ModificationReason2 struct {
	Cd          ModificationReason2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	AddtlRsnInf Max210Text                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AddtlRsnInf,omitempty"`
}

type ModificationReason2Choice struct {
	Cd    ModifiedStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	Prtry GenericIdentification20   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Prtry"`
}

type ModificationStatus2Choice struct {
	NoSpcfdRsn NoReasonCode          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 NoSpcfdRsn"`
	Rsn        []ModificationReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Rsn,omitempty"`
}

// May be one of MDBY, OTHR
type ModifiedStatusReason1Code string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Adr,omitempty"`
}

// May be one of NORE
type NoReasonCode string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Tp"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 PrtryId"`
}

type PartyIdentification37Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 NmAndAdr"`
	Ctry     CountryCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Ctry"`
}

type PartyIdentification44Choice struct {
	AnyBIC   AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AnyBIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Ctry"`
}

type PartyIdentification45Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 NmAndAdr"`
}

type PartyIdentification46 struct {
	Id     PartyIdentification44Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Id"`
	PrcgId Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 PrcgId,omitempty"`
}

type PartyIdentificationAndAccount44 struct {
	Id        PartyIdentification45Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Id"`
	SfkpgAcct SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 SfkpgAcct,omitempty"`
	PrcgId    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 PrcgId,omitempty"`
}

// May be one of ADEA, BLOC, MUNO, NEXT, MINO, OTHR, DENO, CERT
type PendingProcessingReason3Code string

type PendingProcessingReason5 struct {
	Cd          PendingProcessingReason5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	AddtlRsnInf Max210Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AddtlRsnInf,omitempty"`
}

type PendingProcessingReason5Choice struct {
	Cd    PendingProcessingReason3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	Prtry GenericIdentification20      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Prtry"`
}

type PendingProcessingStatus7Choice struct {
	NoSpcfdRsn NoReasonCode               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 NoSpcfdRsn"`
	Rsn        []PendingProcessingReason5 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Rsn"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Ctry"`
}

type ProprietaryReason1 struct {
	Rsn         GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason1 struct {
	PrtrySts GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 PrtrySts"`
	PrtryRsn []ProprietaryReason1    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 PrtryRsn,omitempty"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 OrgnlAndCurFace"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type RejectionReason11 struct {
	Cd          RejectionReason11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AddtlRsnInf,omitempty"`
}

type RejectionReason11Choice struct {
	Cd    RejectionReason31Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Prtry"`
}

// May be one of SAFE, DQUA, ADEA, DSEC, LATE, CASH, DDEA, DTRD, PLCE, RTGS, NCRR, PHYS, REFE, DMON, MINO, BATC, MUNO, TXST, SETS, IIND, CAEV, CASY, DDAT, SETR, SDUT, INPS, OTHR, ICUS, ICAG, DEPT, IEXE, INVL, INVB, INVN, VALR
type RejectionReason31Code string

type RejectionStatus8Choice struct {
	NoSpcfdRsn NoReasonCode        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 NoSpcfdRsn"`
	Rsn        []RejectionReason11 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Rsn"`
}

type RepairReason1Choice struct {
	Cd    RepairReason4Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Prtry"`
}

// May be one of BATC, CAEV, CASH, CASY, DDAT, DDEA, DMON, DQUA, DSEC, DTRD, IIND, MINO, MUNO, NCRR, PHYS, PLCE, REFE, RTGS, SAFE, SETR, SETS, TXST, INPS, SDUT, OTHR, IEXE, ICAG, DEPT, ICUS
type RepairReason4Code string

type RepairReason6 struct {
	Cd          RepairReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	AddtlRsnInf Max210Text          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AddtlRsnInf,omitempty"`
}

type RepairStatus8Choice struct {
	NoSpcfdRsn NoReasonCode    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 NoSpcfdRsn"`
	Rsn        []RepairReason6 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Rsn"`
}

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Nm,omitempty"`
}

type SecuritiesSettlementTransactionModificationRequestStatusAdviceV03 struct {
	ModReqRef   Identification1                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 ModReqRef"`
	AcctOwnr    PartyIdentification36Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount13                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 SfkpgAcct"`
	TxId        TransactionIdentifications25        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 TxId,omitempty"`
	ModPrcgSts  ModificationProcessingStatus4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 ModPrcgSts"`
	TxDtls      TransactionDetails45                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 TxDtls,omitempty"`
	SplmtryData []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 SplmtryData,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Desc,omitempty"`
}

type SettlementDate2Choice struct {
	Dt   DateAndDateTimeChoice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Dt"`
	DtCd SettlementDateCode2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 DtCd"`
}

type SettlementDateCode2Choice struct {
	Cd    DateType4Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Prtry"`
}

type SettlementParties13 struct {
	Dpstry PartyIdentification46           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount44 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Pty5,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeDate1Choice struct {
	Dt   DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Dt"`
	DtCd TradeDateCode1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 DtCd"`
}

type TradeDateCode1Choice struct {
	Cd    DateType3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Prtry"`
}

type TransactionDetails45 struct {
	FinInstrmId     SecurityIdentification14    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 FinInstrmId"`
	SctiesMvmntTp   ReceiveDelivery1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 SctiesMvmntTp"`
	Pmt             DeliveryReceiptType2Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Pmt"`
	SttlmQty        Quantity6Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 SttlmQty"`
	SttlmAmt        AmountAndDirection8         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 SttlmAmt,omitempty"`
	SttlmDt         SettlementDate2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 SttlmDt"`
	TradDt          TradeDate1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 TradDt,omitempty"`
	DlvrgSttlmPties SettlementParties13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties  SettlementParties13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 RcvgSttlmPties,omitempty"`
	Invstr          PartyIdentification37Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Invstr,omitempty"`
}

type TransactionIdentifications25 struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AcctOwnrTxId"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 PrcrTxId,omitempty"`
	OthrId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 OthrId,omitempty"`
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

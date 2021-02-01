// Code generated by main. DO NOT EDIT.

package semt_004_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Account7 struct {
	Id       AccountIdentification1     `xml:"urn:swift:xsd:semt.004.001.02 Id"`
	AcctSvcr PartyIdentification2Choice `xml:"urn:swift:xsd:semt.004.001.02 AcctSvcr,omitempty"`
}

type AccountIdentification1 struct {
	Prtry SimpleIdentificationInformation `xml:"urn:swift:xsd:semt.004.001.02 Prtry"`
}

type AccountIdentification3 struct {
	Id   AccountIdentification1 `xml:"urn:swift:xsd:semt.004.001.02 Id"`
	Issr Max8Text               `xml:"urn:swift:xsd:semt.004.001.02 Issr"`
	Inf  Exact4AlphaNumericText `xml:"urn:swift:xsd:semt.004.001.02 Inf"`
}

type AccountIdentificationAndPurpose struct {
	Id   AccountIdentification1            `xml:"urn:swift:xsd:semt.004.001.02 Id"`
	Purp SecuritiesAccountPurposeType1Code `xml:"urn:swift:xsd:semt.004.001.02 Purp"`
}

type AccountIdentificationFormatChoice struct {
	SmplId    AccountIdentification1          `xml:"urn:swift:xsd:semt.004.001.02 SmplId"`
	IdAndPurp AccountIdentificationAndPurpose `xml:"urn:swift:xsd:semt.004.001.02 IdAndPurp"`
	IdAsDSS   AccountIdentification3          `xml:"urn:swift:xsd:semt.004.001.02 IdAsDSS"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAnd13DecimalAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AdditionalBalanceInformation2 struct {
	Qty            SubBalanceQuantity1Choice  `xml:"urn:swift:xsd:semt.004.001.02 Qty"`
	SubBalTp       SecuritiesBalanceType2Code `xml:"urn:swift:xsd:semt.004.001.02 SubBalTp"`
	XtndedSubBalTp Extended350Code            `xml:"urn:swift:xsd:semt.004.001.02 XtndedSubBalTp"`
}

type AdditionalReference2 struct {
	Ref     Max35Text                  `xml:"urn:swift:xsd:semt.004.001.02 Ref"`
	RefIssr PartyIdentification1Choice `xml:"urn:swift:xsd:semt.004.001.02 RefIssr,omitempty"`
	MsgNm   Max35Text                  `xml:"urn:swift:xsd:semt.004.001.02 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AggregateBalanceInformation4 struct {
	AggtQty            BalanceQuantity1Choice                 `xml:"urn:swift:xsd:semt.004.001.02 AggtQty"`
	AvlblQty           BalanceQuantity1Choice                 `xml:"urn:swift:xsd:semt.004.001.02 AvlblQty,omitempty"`
	NotAvlblQty        BalanceQuantity1Choice                 `xml:"urn:swift:xsd:semt.004.001.02 NotAvlblQty,omitempty"`
	DaysAcrd           float64                                `xml:"urn:swift:xsd:semt.004.001.02 DaysAcrd,omitempty"`
	HldgVal            []ActiveOrHistoricCurrencyAndAmount    `xml:"urn:swift:xsd:semt.004.001.02 HldgVal,omitempty"`
	PrvsHldgVal        ActiveOrHistoricCurrencyAndAmount      `xml:"urn:swift:xsd:semt.004.001.02 PrvsHldgVal,omitempty"`
	AcrdIntrstAmt      ActiveOrHistoricCurrencyAndAmount      `xml:"urn:swift:xsd:semt.004.001.02 AcrdIntrstAmt,omitempty"`
	AcrdIntrstAmtSgn   bool                                   `xml:"urn:swift:xsd:semt.004.001.02 AcrdIntrstAmtSgn,omitempty"`
	BookVal            ActiveOrHistoricCurrencyAndAmount      `xml:"urn:swift:xsd:semt.004.001.02 BookVal,omitempty"`
	SfkpgPlc           SafekeepingPlaceFormatChoice           `xml:"urn:swift:xsd:semt.004.001.02 SfkpgPlc,omitempty"`
	FinInstrmDtls      FinancialInstrument13                  `xml:"urn:swift:xsd:semt.004.001.02 FinInstrmDtls"`
	PricDtls           []PriceInformation2                    `xml:"urn:swift:xsd:semt.004.001.02 PricDtls,omitempty"`
	FXDtls             ForeignExchangeTerms6                  `xml:"urn:swift:xsd:semt.004.001.02 FXDtls,omitempty"`
	BalBrkdwnDtls      []SubBalanceInformation2               `xml:"urn:swift:xsd:semt.004.001.02 BalBrkdwnDtls,omitempty"`
	AddtlBalBrkdwnDtls []AdditionalBalanceInformation2        `xml:"urn:swift:xsd:semt.004.001.02 AddtlBalBrkdwnDtls,omitempty"`
	BalAtSfkpgPlc      []AggregateBalancePerSafekeepingPlace3 `xml:"urn:swift:xsd:semt.004.001.02 BalAtSfkpgPlc,omitempty"`
}

type AggregateBalancePerSafekeepingPlace3 struct {
	AggtQty            BalanceQuantity1Choice              `xml:"urn:swift:xsd:semt.004.001.02 AggtQty"`
	AvlblQty           BalanceQuantity1Choice              `xml:"urn:swift:xsd:semt.004.001.02 AvlblQty,omitempty"`
	NotAvlblQty        BalanceQuantity1Choice              `xml:"urn:swift:xsd:semt.004.001.02 NotAvlblQty,omitempty"`
	DaysAcrd           float64                             `xml:"urn:swift:xsd:semt.004.001.02 DaysAcrd,omitempty"`
	HldgVal            []ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:semt.004.001.02 HldgVal,omitempty"`
	PrvsHldgVal        ActiveOrHistoricCurrencyAndAmount   `xml:"urn:swift:xsd:semt.004.001.02 PrvsHldgVal,omitempty"`
	AcrdIntrstAmt      ActiveOrHistoricCurrencyAndAmount   `xml:"urn:swift:xsd:semt.004.001.02 AcrdIntrstAmt,omitempty"`
	AcrdIntrstAmtSgn   bool                                `xml:"urn:swift:xsd:semt.004.001.02 AcrdIntrstAmtSgn,omitempty"`
	BookVal            ActiveOrHistoricCurrencyAndAmount   `xml:"urn:swift:xsd:semt.004.001.02 BookVal,omitempty"`
	SfkpgPlc           SafekeepingPlaceFormatChoice        `xml:"urn:swift:xsd:semt.004.001.02 SfkpgPlc"`
	PricDtls           []PriceInformation2                 `xml:"urn:swift:xsd:semt.004.001.02 PricDtls,omitempty"`
	FXDtls             ForeignExchangeTerms6               `xml:"urn:swift:xsd:semt.004.001.02 FXDtls,omitempty"`
	BalBrkdwnDtls      []SubBalanceInformation2            `xml:"urn:swift:xsd:semt.004.001.02 BalBrkdwnDtls,omitempty"`
	AddtlBalBrkdwnDtls []AdditionalBalanceInformation2     `xml:"urn:swift:xsd:semt.004.001.02 AddtlBalBrkdwnDtls,omitempty"`
}

type AlternateSecurityIdentification1 struct {
	Id         Max35Text   `xml:"urn:swift:xsd:semt.004.001.02 Id"`
	DmstIdSrc  CountryCode `xml:"urn:swift:xsd:semt.004.001.02 DmstIdSrc"`
	PrtryIdSrc Max35Text   `xml:"urn:swift:xsd:semt.004.001.02 PrtryIdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type BalanceQuantity1Choice struct {
	Qty      FinancialInstrumentQuantityChoice `xml:"urn:swift:xsd:semt.004.001.02 Qty"`
	QtyAsDSS GenericIdentification6            `xml:"urn:swift:xsd:semt.004.001.02 QtyAsDSS"`
}

// Must be at least 1 items long
type BloombergIdentifier string

// Must be at least 1 items long
type ConsolidatedTapeAssociationIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CustodyStatementOfHoldings2 struct {
	StmtGnlDtls Statement7                     `xml:"urn:swift:xsd:semt.004.001.02 StmtGnlDtls,omitempty"`
	AcctDtls    SafekeepingAccount2            `xml:"urn:swift:xsd:semt.004.001.02 AcctDtls,omitempty"`
	BalForAcct  []AggregateBalanceInformation4 `xml:"urn:swift:xsd:semt.004.001.02 BalForAcct,omitempty"`
	SubAcctDtls []SubAccountIdentification5    `xml:"urn:swift:xsd:semt.004.001.02 SubAcctDtls,omitempty"`
	TtlVals     TotalValueInPageAndStatement   `xml:"urn:swift:xsd:semt.004.001.02 TtlVals,omitempty"`
	Xtnsn       []Extension1                   `xml:"urn:swift:xsd:semt.004.001.02 Xtnsn,omitempty"`
}

type CustodyStatementOfHoldingsCancellationV02 struct {
	MsgId        MessageIdentification1      `xml:"urn:swift:xsd:semt.004.001.02 MsgId"`
	PrvsRef      AdditionalReference2        `xml:"urn:swift:xsd:semt.004.001.02 PrvsRef,omitempty"`
	RltdRef      AdditionalReference2        `xml:"urn:swift:xsd:semt.004.001.02 RltdRef,omitempty"`
	MsgPgntn     Pagination                  `xml:"urn:swift:xsd:semt.004.001.02 MsgPgntn"`
	StmtToBeCanc CustodyStatementOfHoldings2 `xml:"urn:swift:xsd:semt.004.001.02 StmtToBeCanc,omitempty"`
}

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:swift:xsd:semt.004.001.02 Dt"`
	DtTm ISODateTime `xml:"urn:swift:xsd:semt.004.001.02 DtTm"`
}

// May be one of DIST, ACCU
type DistributionPolicy1Code string

type Document struct {
	CtdyStmtOfHldgsCxlV02 CustodyStatementOfHoldingsCancellationV02 `xml:"urn:swift:xsd:semt.004.001.02 CtdyStmtOfHldgsCxlV02"`
}

// Must be at least 1 items long
type EuroclearClearstreamIdentifier string

// May be one of YEAR, SEMI, QUTR, TOMN, MNTH, TWMN, TOWK, WEEK, DAIL, ADHO, INDA, OVNG, ONDE
type EventFrequency1Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type Extended350Code string

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:swift:xsd:semt.004.001.02 PlcAndNm"`
	Txt      Max350Text `xml:"urn:swift:xsd:semt.004.001.02 Txt"`
}

type FinancialInstrument13 struct {
	Id          SecurityIdentification3Choice `xml:"urn:swift:xsd:semt.004.001.02 Id"`
	Nm          Max350Text                    `xml:"urn:swift:xsd:semt.004.001.02 Nm,omitempty"`
	SplmtryId   Max35Text                     `xml:"urn:swift:xsd:semt.004.001.02 SplmtryId,omitempty"`
	ClssTp      Max35Text                     `xml:"urn:swift:xsd:semt.004.001.02 ClssTp,omitempty"`
	SctiesForm  FormOfSecurity1Code           `xml:"urn:swift:xsd:semt.004.001.02 SctiesForm,omitempty"`
	DstrbtnPlcy DistributionPolicy1Code       `xml:"urn:swift:xsd:semt.004.001.02 DstrbtnPlcy,omitempty"`
}

type FinancialInstrumentQuantityChoice struct {
	Unit     float64 `xml:"urn:swift:xsd:semt.004.001.02 Unit"`
	FaceAmt  float64 `xml:"urn:swift:xsd:semt.004.001.02 FaceAmt"`
	AmtsdVal float64 `xml:"urn:swift:xsd:semt.004.001.02 AmtsdVal"`
}

type ForeignExchangeTerms6 struct {
	UnitCcy  ActiveOrHistoricCurrencyCode `xml:"urn:swift:xsd:semt.004.001.02 UnitCcy"`
	QtdCcy   ActiveOrHistoricCurrencyCode `xml:"urn:swift:xsd:semt.004.001.02 QtdCcy"`
	XchgRate float64                      `xml:"urn:swift:xsd:semt.004.001.02 XchgRate"`
	QtnDt    ISODateTime                  `xml:"urn:swift:xsd:semt.004.001.02 QtnDt,omitempty"`
	QtgInstn PartyIdentification2Choice   `xml:"urn:swift:xsd:semt.004.001.02 QtgInstn,omitempty"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type FrequencyCodeAndDSSCode1Choice struct {
	FrqcyAsCd  EventFrequency1Code    `xml:"urn:swift:xsd:semt.004.001.02 FrqcyAsCd"`
	FrqcyAsDSS GenericIdentification7 `xml:"urn:swift:xsd:semt.004.001.02 FrqcyAsDSS"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:swift:xsd:semt.004.001.02 Id"`
	SchmeNm Max35Text `xml:"urn:swift:xsd:semt.004.001.02 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:swift:xsd:semt.004.001.02 Issr,omitempty"`
}

type GenericIdentification5 struct {
	Issr  Max8Text               `xml:"urn:swift:xsd:semt.004.001.02 Issr"`
	Inf   Exact4AlphaNumericText `xml:"urn:swift:xsd:semt.004.001.02 Inf"`
	Nrrtv Max35Text              `xml:"urn:swift:xsd:semt.004.001.02 Nrrtv,omitempty"`
}

type GenericIdentification6 struct {
	Issr Max8Text               `xml:"urn:swift:xsd:semt.004.001.02 Issr"`
	Inf  Exact4AlphaNumericText `xml:"urn:swift:xsd:semt.004.001.02 Inf"`
	Bal  float64                `xml:"urn:swift:xsd:semt.004.001.02 Bal"`
}

type GenericIdentification7 struct {
	Issr Max8Text  `xml:"urn:swift:xsd:semt.004.001.02 Issr"`
	Inf  Max35Text `xml:"urn:swift:xsd:semt.004.001.02 Inf"`
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

type Intermediary11 struct {
	Id         PartyIdentification2Choice `xml:"urn:swift:xsd:semt.004.001.02 Id"`
	Acct       Account7                   `xml:"urn:swift:xsd:semt.004.001.02 Acct,omitempty"`
	Role       InvestmentFundRole2Code    `xml:"urn:swift:xsd:semt.004.001.02 Role,omitempty"`
	XtndedRole Extended350Code            `xml:"urn:swift:xsd:semt.004.001.02 XtndedRole,omitempty"`
}

// May be one of FMCO, REGI, TRAG, INTR, DIST, CONC, UCL1, UCL2, TRAN
type InvestmentFundRole2Code string

type LongPostalAddress1Choice struct {
	Ustrd Max140Text                   `xml:"urn:swift:xsd:semt.004.001.02 Ustrd"`
	Strd  StructuredLongPostalAddress1 `xml:"urn:swift:xsd:semt.004.001.02 Strd"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max70Text string

// Must be at least 1 items long
type Max8Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:swift:xsd:semt.004.001.02 Id"`
	CreDtTm ISODateTime `xml:"urn:swift:xsd:semt.004.001.02 CreDtTm"`
}

type NameAndAddress2 struct {
	Nm  Max35Text                `xml:"urn:swift:xsd:semt.004.001.02 Nm"`
	Adr LongPostalAddress1Choice `xml:"urn:swift:xsd:semt.004.001.02 Adr,omitempty"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:swift:xsd:semt.004.001.02 Nm"`
	Adr PostalAddress1 `xml:"urn:swift:xsd:semt.004.001.02 Adr,omitempty"`
}

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:swift:xsd:semt.004.001.02 PgNb"`
	LastPgInd bool            `xml:"urn:swift:xsd:semt.004.001.02 LastPgInd"`
}

type PartyIdentification1Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:swift:xsd:semt.004.001.02 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:swift:xsd:semt.004.001.02 PrtryId"`
	NmAndAdr NameAndAddress2        `xml:"urn:swift:xsd:semt.004.001.02 NmAndAdr"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:swift:xsd:semt.004.001.02 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:swift:xsd:semt.004.001.02 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:swift:xsd:semt.004.001.02 NmAndAdr"`
}

type PartyIdentification3 struct {
	BICOrBEI AnyBICIdentifier `xml:"urn:swift:xsd:semt.004.001.02 BICOrBEI"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:swift:xsd:semt.004.001.02 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:swift:xsd:semt.004.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:swift:xsd:semt.004.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:swift:xsd:semt.004.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:swift:xsd:semt.004.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:swift:xsd:semt.004.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:swift:xsd:semt.004.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:swift:xsd:semt.004.001.02 Ctry"`
}

type PriceInformation2 struct {
	Val       PriceRateOrAmountOrUnknownChoice `xml:"urn:swift:xsd:semt.004.001.02 Val"`
	ValTp     PriceValueType2Code              `xml:"urn:swift:xsd:semt.004.001.02 ValTp,omitempty"`
	Tp        TypeOfPrice11Code                `xml:"urn:swift:xsd:semt.004.001.02 Tp"`
	XtndedTp  Extended350Code                  `xml:"urn:swift:xsd:semt.004.001.02 XtndedTp"`
	SrcOfPric PriceSourceFormatChoice          `xml:"urn:swift:xsd:semt.004.001.02 SrcOfPric,omitempty"`
	QtnDt     DateAndDateTimeChoice            `xml:"urn:swift:xsd:semt.004.001.02 QtnDt,omitempty"`
	Yldd      bool                             `xml:"urn:swift:xsd:semt.004.001.02 Yldd,omitempty"`
}

type PriceRateOrAmountOrUnknownChoice struct {
	Rate     float64                                    `xml:"urn:swift:xsd:semt.004.001.02 Rate"`
	Amt      ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:swift:xsd:semt.004.001.02 Amt"`
	UknwnInd bool                                       `xml:"urn:swift:xsd:semt.004.001.02 UknwnInd"`
}

type PriceSource struct {
	PricSrc PriceSource1Code `xml:"urn:swift:xsd:semt.004.001.02 PricSrc"`
	Nrrtv   Max35Text        `xml:"urn:swift:xsd:semt.004.001.02 Nrrtv,omitempty"`
}

// May be one of FUND, THEO, VEND
type PriceSource1Code string

type PriceSourceFormatChoice struct {
	LclMktPlc    MICIdentifier          `xml:"urn:swift:xsd:semt.004.001.02 LclMktPlc"`
	NonLclMktPlc PriceSource            `xml:"urn:swift:xsd:semt.004.001.02 NonLclMktPlc"`
	PlcAsDSS     GenericIdentification5 `xml:"urn:swift:xsd:semt.004.001.02 PlcAsDSS"`
}

// May be one of DISC, PREM
type PriceValueType2Code string

type QuantityAndAvailability struct {
	Qty       FinancialInstrumentQuantityChoice `xml:"urn:swift:xsd:semt.004.001.02 Qty"`
	AvlbtyInd bool                              `xml:"urn:swift:xsd:semt.004.001.02 AvlbtyInd"`
}

// Must be at least 1 items long
type RICIdentifier string

type SafekeepingAccount2 struct {
	Id        AccountIdentificationFormatChoice `xml:"urn:swift:xsd:semt.004.001.02 Id"`
	FngbInd   bool                              `xml:"urn:swift:xsd:semt.004.001.02 FngbInd"`
	Nm        Max35Text                         `xml:"urn:swift:xsd:semt.004.001.02 Nm,omitempty"`
	Dsgnt     Max35Text                         `xml:"urn:swift:xsd:semt.004.001.02 Dsgnt,omitempty"`
	IntrmyInf []Intermediary11                  `xml:"urn:swift:xsd:semt.004.001.02 IntrmyInf,omitempty"`
	AcctOwnr  PartyIdentification2Choice        `xml:"urn:swift:xsd:semt.004.001.02 AcctOwnr,omitempty"`
	AcctSvcr  PartyIdentification2Choice        `xml:"urn:swift:xsd:semt.004.001.02 AcctSvcr,omitempty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

type SafekeepingPlaceAsCodeAndPartyIdentification struct {
	PlcSfkpg SafekeepingPlace1Code `xml:"urn:swift:xsd:semt.004.001.02 PlcSfkpg"`
	Nrrtv    Max35Text             `xml:"urn:swift:xsd:semt.004.001.02 Nrrtv,omitempty"`
	Pty      PartyIdentification3  `xml:"urn:swift:xsd:semt.004.001.02 Pty,omitempty"`
}

type SafekeepingPlaceFormatChoice struct {
	Id       SafekeepingPlaceAsCodeAndPartyIdentification `xml:"urn:swift:xsd:semt.004.001.02 Id"`
	IdAsDSS  GenericIdentification5                       `xml:"urn:swift:xsd:semt.004.001.02 IdAsDSS"`
	IdAsCtry CountryCode                                  `xml:"urn:swift:xsd:semt.004.001.02 IdAsCtry"`
}

// May be one of MARG, SHOR, ABRD, CEND, DVPA, PHYS
type SecuritiesAccountPurposeType1Code string

// May be one of BLOK, BORR, COLI, COLO, LOAN, MARG, PDMT, PRMT, PRUM, PECA, PEND, PENR, PLED, PDUM, REGO, RSTR, OTHR, TRAN, DRAW, WDOC, BTRA
type SecuritiesBalanceType1Code string

// May be one of CLEN, DIRT, NOMI, OTHR, SPOS, UNRG
type SecuritiesBalanceType2Code string

type SecurityIdentification3Choice struct {
	ISIN        ISINIdentifier                        `xml:"urn:swift:xsd:semt.004.001.02 ISIN"`
	SEDOL       string                                `xml:"urn:swift:xsd:semt.004.001.02 SEDOL"`
	CUSIP       string                                `xml:"urn:swift:xsd:semt.004.001.02 CUSIP"`
	RIC         RICIdentifier                         `xml:"urn:swift:xsd:semt.004.001.02 RIC"`
	TckrSymb    TickerIdentifier                      `xml:"urn:swift:xsd:semt.004.001.02 TckrSymb"`
	Blmbrg      BloombergIdentifier                   `xml:"urn:swift:xsd:semt.004.001.02 Blmbrg"`
	CTA         ConsolidatedTapeAssociationIdentifier `xml:"urn:swift:xsd:semt.004.001.02 CTA"`
	QUICK       string                                `xml:"urn:swift:xsd:semt.004.001.02 QUICK"`
	Wrtppr      string                                `xml:"urn:swift:xsd:semt.004.001.02 Wrtppr"`
	Dtch        string                                `xml:"urn:swift:xsd:semt.004.001.02 Dtch"`
	Vlrn        string                                `xml:"urn:swift:xsd:semt.004.001.02 Vlrn"`
	SCVM        string                                `xml:"urn:swift:xsd:semt.004.001.02 SCVM"`
	Belgn       string                                `xml:"urn:swift:xsd:semt.004.001.02 Belgn"`
	Cmon        EuroclearClearstreamIdentifier        `xml:"urn:swift:xsd:semt.004.001.02 Cmon"`
	OthrPrtryId AlternateSecurityIdentification1      `xml:"urn:swift:xsd:semt.004.001.02 OthrPrtryId"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:swift:xsd:semt.004.001.02 Id"`
}

type Statement7 struct {
	Ref       Max35Text                               `xml:"urn:swift:xsd:semt.004.001.02 Ref"`
	StmtDtTm  DateAndDateTimeChoice                   `xml:"urn:swift:xsd:semt.004.001.02 StmtDtTm"`
	CreDtTm   DateAndDateTimeChoice                   `xml:"urn:swift:xsd:semt.004.001.02 CreDtTm,omitempty"`
	Frqcy     FrequencyCodeAndDSSCode1Choice          `xml:"urn:swift:xsd:semt.004.001.02 Frqcy"`
	UpdTp     StatementUpdateTypeCodeAndDSSCodeChoice `xml:"urn:swift:xsd:semt.004.001.02 UpdTp"`
	ActvtyInd bool                                    `xml:"urn:swift:xsd:semt.004.001.02 ActvtyInd"`
	StmtBsis  StatementBasisCodeAndDSSCodeChoice      `xml:"urn:swift:xsd:semt.004.001.02 StmtBsis"`
	RptNb     Max5NumericText                         `xml:"urn:swift:xsd:semt.004.001.02 RptNb,omitempty"`
}

// May be one of CONT, SETT, TRAD
type StatementBasis1Code string

type StatementBasisCodeAndDSSCodeChoice struct {
	StmtBsisAsCd  StatementBasis1Code    `xml:"urn:swift:xsd:semt.004.001.02 StmtBsisAsCd"`
	StmtBsisAsDSS GenericIdentification7 `xml:"urn:swift:xsd:semt.004.001.02 StmtBsisAsDSS"`
}

// May be one of COMP, DELT
type StatementUpdateTypeCode string

type StatementUpdateTypeCodeAndDSSCodeChoice struct {
	StmtUpdTpAsCd  StatementUpdateTypeCode `xml:"urn:swift:xsd:semt.004.001.02 StmtUpdTpAsCd"`
	StmtUpdTpAsDSS GenericIdentification7  `xml:"urn:swift:xsd:semt.004.001.02 StmtUpdTpAsDSS"`
}

type StructuredLongPostalAddress1 struct {
	BldgNm     Max35Text   `xml:"urn:swift:xsd:semt.004.001.02 BldgNm,omitempty"`
	StrtNm     Max35Text   `xml:"urn:swift:xsd:semt.004.001.02 StrtNm,omitempty"`
	StrtBldgId Max35Text   `xml:"urn:swift:xsd:semt.004.001.02 StrtBldgId,omitempty"`
	Flr        Max16Text   `xml:"urn:swift:xsd:semt.004.001.02 Flr,omitempty"`
	TwnNm      Max35Text   `xml:"urn:swift:xsd:semt.004.001.02 TwnNm"`
	DstrctNm   Max35Text   `xml:"urn:swift:xsd:semt.004.001.02 DstrctNm,omitempty"`
	RgnId      Max35Text   `xml:"urn:swift:xsd:semt.004.001.02 RgnId,omitempty"`
	Stat       Max35Text   `xml:"urn:swift:xsd:semt.004.001.02 Stat,omitempty"`
	CtyId      Max35Text   `xml:"urn:swift:xsd:semt.004.001.02 CtyId,omitempty"`
	Ctry       CountryCode `xml:"urn:swift:xsd:semt.004.001.02 Ctry"`
	PstCdId    Max16Text   `xml:"urn:swift:xsd:semt.004.001.02 PstCdId"`
	POB        Max16Text   `xml:"urn:swift:xsd:semt.004.001.02 POB,omitempty"`
}

type SubAccountIdentification5 struct {
	Id            AccountIdentificationFormatChoice `xml:"urn:swift:xsd:semt.004.001.02 Id"`
	FngbInd       bool                              `xml:"urn:swift:xsd:semt.004.001.02 FngbInd"`
	ActvtyInd     bool                              `xml:"urn:swift:xsd:semt.004.001.02 ActvtyInd"`
	BalForSubAcct []AggregateBalanceInformation4    `xml:"urn:swift:xsd:semt.004.001.02 BalForSubAcct,omitempty"`
}

type SubBalanceInformation2 struct {
	Qty                SubBalanceQuantity1Choice       `xml:"urn:swift:xsd:semt.004.001.02 Qty"`
	SubBalTp           SecuritiesBalanceType1Code      `xml:"urn:swift:xsd:semt.004.001.02 SubBalTp"`
	XtndedSubBalTp     Extended350Code                 `xml:"urn:swift:xsd:semt.004.001.02 XtndedSubBalTp"`
	AddtlBalBrkdwnDtls []AdditionalBalanceInformation2 `xml:"urn:swift:xsd:semt.004.001.02 AddtlBalBrkdwnDtls,omitempty"`
}

type SubBalanceQuantity1Choice struct {
	Qty          FinancialInstrumentQuantityChoice `xml:"urn:swift:xsd:semt.004.001.02 Qty"`
	QtyAsDSS     GenericIdentification6            `xml:"urn:swift:xsd:semt.004.001.02 QtyAsDSS"`
	QtyAndAvlbty QuantityAndAvailability           `xml:"urn:swift:xsd:semt.004.001.02 QtyAndAvlbty"`
}

// Must be at least 1 items long
type TickerIdentifier string

type TotalValueInPageAndStatement struct {
	TtlHldgsValOfPg   ActiveCurrencyAndAmount `xml:"urn:swift:xsd:semt.004.001.02 TtlHldgsValOfPg,omitempty"`
	TtlHldgsValOfStmt ActiveCurrencyAndAmount `xml:"urn:swift:xsd:semt.004.001.02 TtlHldgsValOfStmt"`
}

// May be one of BIDE, OFFR, NAVL, CREA, CANC, INTE, SWNG, MIDD, RINV, SWIC, MRKT, INDC
type TypeOfPrice11Code string

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

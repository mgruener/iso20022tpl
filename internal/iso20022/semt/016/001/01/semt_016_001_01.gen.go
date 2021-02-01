// Code generated by main. DO NOT EDIT.

package semt_016_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveOrHistoricCurrencyAnd13DecimalAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateIdentification1 struct {
	Id    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id"`
	IdSrc IdentificationSource1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 IdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{1,6}
type CFIIdentifier string

type ClassificationType2Choice struct {
	ClssfctnFinInstrm CFIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 AltrnClssfctn"`
}

// May be one of CODU, COPY, DUPL
type CopyDuplicate1Code string

type CorporateActionEventType3Choice struct {
	Cd    CorporateActionEventType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Cd"`
	Prtry GenericIdentification20       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

// May be one of ACTV, ATTI, BIDS, BONU, BPUT, BRUP, CAPG, CAPI, CERT, CHAN, CLSA, CONS, CONV, COOP, DECR, DETI, DFLT, DLST, DRAW, DRIP, DSCL, DTCH, DVCA, DVOP, DVSC, DVSE, EXOF, EXRI, EXTM, EXWA, CAPD, INCR, INTR, LIQU, MCAL, MRGR, ODLT, OTHR, PARI, PCAL, PDEF, PINK, PLAC, PPMT, PRED, PRII, PRIO, REDM, REDO, REMK, RHDI, RHTS, SHPR, SMAL, SOFF, SPLF, SPLR, SUSP, TEND, TREC, WRTH, WTRC, CREV
type CorporateActionEventType6Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 DtTm"`
}

type DateTimePeriodDetails struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 ToDtTm"`
}

type Document struct {
	IntraPosMvmntPstngRpt IntraPositionMovementPostingReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 IntraPosMvmntPstngRpt"`
}

type DocumentIdentification11 struct {
	Id       Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id"`
	CreDtTm  DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 CreDtTm,omitempty"`
	CpyDplct CopyDuplicate1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 CpyDplct,omitempty"`
}

// May be one of YEAR, MNTH, QUTR, SEMI, WEEK
type EventFrequency3Code string

// May be one of YEAR, ADHO, MNTH, DAIL, INDA, WEEK
type EventFrequency4Code string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{5}
type Exact5NumericText string

type Extension2 struct {
	PlcAndNm   Max350Text         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PlcAndNm,omitempty"`
	XtnsnEnvlp ExtensionEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 XtnsnEnvlp"`
}

type ExtensionEnvelope1 struct {
	Item string `xml:",any"`
}

type FinancialInstrumentAttributes4 struct {
	PlcOfListg             MarketIdentification5                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PlcOfListg,omitempty"`
	DayCntBsis             InterestComputationMethodFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 DayCntBsis,omitempty"`
	RegnForm               FormOfSecurity2Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 RegnForm,omitempty"`
	PmtFrqcy               Frequency3Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PmtFrqcy,omitempty"`
	PmtSts                 SecuritiesPaymentStatus2Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PmtSts,omitempty"`
	PmtDrctn               PaymentDirection2Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PmtDrctn,omitempty"`
	VarblRateChngFrqcy     Frequency3Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 VarblRateChngFrqcy,omitempty"`
	PrefToIncm             PreferenceToIncome2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PrefToIncm,omitempty"`
	ClssfctnTp             ClassificationType2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 ClssfctnTp,omitempty"`
	OptnStyle              OptionStyle4Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 OptnStyle,omitempty"`
	OptnTp                 OptionType2Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 OptnTp,omitempty"`
	DnmtnCcy               ActiveOrHistoricCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 DnmtnCcy,omitempty"`
	CpnDt                  ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 CpnDt,omitempty"`
	XpryDt                 ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 XpryDt,omitempty"`
	FltgRateFxgDt          ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 FltgRateFxgDt,omitempty"`
	MtrtyDt                ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 MtrtyDt,omitempty"`
	IsseDt                 ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 IsseDt,omitempty"`
	NxtCllblDt             ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 NxtCllblDt,omitempty"`
	PutblDt                ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PutblDt,omitempty"`
	DtdDt                  ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 DtdDt,omitempty"`
	FrstPmtDt              ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 FrstPmtDt,omitempty"`
	PrvsFctr               float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PrvsFctr,omitempty"`
	CurFctr                float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 CurFctr,omitempty"`
	NxtFctr                float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 NxtFctr,omitempty"`
	IntrstRate             float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 IntrstRate,omitempty"`
	NxtIntrstRate          float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 NxtIntrstRate,omitempty"`
	IndxRateBsis           float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 IndxRateBsis,omitempty"`
	CpnAttchdNb            Number2Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 CpnAttchdNb,omitempty"`
	PoolNb                 Number2Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PoolNb,omitempty"`
	QtyBrkdwn              []QuantityBreakdown5                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 QtyBrkdwn,omitempty"`
	VarblRateInd           bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 VarblRateInd,omitempty"`
	CllblInd               bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 CllblInd,omitempty"`
	PutblInd               bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PutblInd,omitempty"`
	MktOrIndctvPric        PriceType1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 MktOrIndctvPric,omitempty"`
	ExrcPric               Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 ExrcPric,omitempty"`
	SbcptPric              Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 SbcptPric,omitempty"`
	ConvsPric              Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 ConvsPric,omitempty"`
	StrkPric               Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 StrkPric,omitempty"`
	MinNmnlQty             FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 MinNmnlQty,omitempty"`
	CtrctSz                FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 CtrctSz,omitempty"`
	UndrlygFinInstrmId     []SecurityIdentification11             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 UndrlygFinInstrmId,omitempty"`
	FinInstrmAttrAddtlDtls Max350Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 FinInstrmAttrAddtlDtls,omitempty"`
}

type FinancialInstrumentDetails1 struct {
	FinInstrmId      SecurityIdentification11       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 FinInstrmId"`
	FinInstrmAttrbts FinancialInstrumentAttributes4 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 FinInstrmAttrbts,omitempty"`
	SubBal           []IntraPositionDetails3        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 SubBal"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 AmtsdVal"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type FormOfSecurity2Choice struct {
	Cd    FormOfSecurity1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

type Frequency3Choice struct {
	Cd    EventFrequency3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

type Frequency4Choice struct {
	Cd    EventFrequency4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Issr,omitempty"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id,omitempty"`
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

type IdentificationSource1Choice struct {
	Dmst  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Dmst"`
	Prtry Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat1Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

type IntraPositionDetails3 struct {
	SfkpgPlc      SafekeepingPlaceFormat3Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 SfkpgPlc,omitempty"`
	BalFr         SecuritiesBalanceType3Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 BalFr"`
	IntraPosMvmnt []IntraPositionMovementDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 IntraPosMvmnt"`
}

type IntraPositionMovementDetails1 struct {
	Id                 References5Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id,omitempty"`
	SttldQty           FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 SttldQty"`
	PrevslySttldQty    FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PrevslySttldQty,omitempty"`
	RmngToBeSttldQty   FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 RmngToBeSttldQty,omitempty"`
	BalTo              SecuritiesBalanceType3Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 BalTo"`
	SttlmDt            DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 SttlmDt"`
	AvlblDt            DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 AvlblDt,omitempty"`
	CorpActnEvtTp      CorporateActionEventType3Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 CorpActnEvtTp,omitempty"`
	InstrPrcgAddtlDtls Max350Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 InstrPrcgAddtlDtls,omitempty"`
	Xtnsn              []Extension2                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Xtnsn,omitempty"`
}

type IntraPositionMovementPostingReportV01 struct {
	Id          DocumentIdentification11      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id"`
	Pgntn       Pagination                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Pgntn"`
	StmtGnlDtls Statement15                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 StmtGnlDtls"`
	AcctOwnr    PartyIdentification13Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount13           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 SfkpgAcct"`
	FinInstrm   []FinancialInstrumentDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 FinInstrm,omitempty"`
	MsgOrgtr    PartyIdentification10Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 MsgOrgtr,omitempty"`
	MsgRcpt     PartyIdentification10Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 MsgRcpt,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MarketIdentification1Choice struct {
	MktIdrCd MICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 MktIdrCd"`
	Desc     Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Desc"`
}

type MarketIdentification5 struct {
	Id MarketIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id,omitempty"`
	Tp MarketType2Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Tp"`
}

type MarketType2Choice struct {
	Cd    MarketType5Code         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

// May be one of OTCO, EXCH
type MarketType5Code string

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

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Adr,omitempty"`
}

type Number2Choice struct {
	Shrt Exact3NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Shrt"`
	Lng  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Lng"`
}

type Number3Choice struct {
	Shrt Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Shrt"`
	Lng  Exact5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Lng"`
}

// May be one of AMER, EURO
type OptionStyle2Code string

type OptionStyle4Choice struct {
	Cd    OptionStyle2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

// May be one of CALL, PUTO
type OptionType1Code string

type OptionType2Choice struct {
	Cd    OptionType1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 LastPgInd"`
}

type PartyIdentification10Choice struct {
	BICOrBEI AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 BICOrBEI"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 NmAndAdr"`
}

type PartyIdentification13Choice struct {
	BICOrBEI AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 BICOrBEI"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PrtryId"`
}

type PaymentDirection2Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

type Period2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 ToDt"`
}

type Period2Choice struct {
	FrDtTmToDtTm DateTimePeriodDetails `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 FrDtTmToDtTm"`
	FrDtToDt     Period2               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 FrDtToDt"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Ctry"`
}

// May be one of ORDN, PFRD
type PreferenceToIncome1Code string

type PreferenceToIncome2Choice struct {
	Cd    PreferenceToIncome1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

type Price2 struct {
	Tp  YieldedOrValueType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Tp"`
	Val PriceRateOrAmountChoice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Val"`
}

type PriceRateOrAmountChoice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Rate"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Amt"`
}

type PriceType1Choice struct {
	Mkt    Price2 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Mkt"`
	Indctv Price2 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Indctv"`
}

// May be one of DISC, PREM, PARV
type PriceValueType1Code string

type QuantityBreakdown5 struct {
	LotNb  Number2Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 LotNb"`
	LotQty FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 LotQty,omitempty"`
}

type References5Choice struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 AcctOwnrTxId"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 AcctSvcrTxId"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 PoolId"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 MktInfrstrctrTxId"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat3Choice struct {
	Id      SafekeepingPlaceTypeAndText3             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id"`
}

type SafekeepingPlaceTypeAndText3 struct {
	SfkpgPlcTp SafekeepingPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id,omitempty"`
}

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Nm,omitempty"`
}

// May be one of BLOK, AWAS, AVAI, BLCA, BLOT, BLOV, BORR, COLI, COLO, COLA, LOAN, MARG, PECA, PEDA, PLED, REGO, RSTR, OTHR, TRAN, DRAW, CLEN, DIRT, NOMI, SPOS, UNRG, ISSU, QUAS, LODE
type SecuritiesBalanceType11Code string

type SecuritiesBalanceType3Choice struct {
	Cd    SecuritiesBalanceType11Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Cd"`
	Prtry GenericIdentification20     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

// May be one of FULL, NILL, PART
type SecuritiesPaymentStatus1Code string

type SecuritiesPaymentStatus2Choice struct {
	Cd    SecuritiesPaymentStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Cd"`
	Prtry GenericIdentification20      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

type SecurityIdentification11 struct {
	Id   SecurityIdentification11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Id"`
	Desc Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Desc,omitempty"`
}

type SecurityIdentification11Choice struct {
	ISIN   ISINIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 ISIN"`
	OthrId AlternateIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 OthrId"`
}

type Statement15 struct {
	RptNb     Number3Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 RptNb,omitempty"`
	QryRef    Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 QryRef,omitempty"`
	StmtId    Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 StmtId,omitempty"`
	StmtPrd   Period2Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 StmtPrd"`
	Frqcy     Frequency4Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Frqcy,omitempty"`
	UpdTp     UpdateType2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 UpdTp,omitempty"`
	ActvtyInd bool              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 ActvtyInd"`
}

// May be one of COMP, DELT
type StatementUpdateType1Code string

type UpdateType2Choice struct {
	Cd    StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Cd"`
	Prtry GenericIdentification20  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Prtry"`
}

type YieldedOrValueType1Choice struct {
	Yldd  bool                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Yldd"`
	ValTp PriceValueType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 ValTp"`
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

// Code generated by main. DO NOT EDIT.

package semt_016_002_08

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AmountAndDirection55 struct {
	Amt                 RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Amt"`
	CdtDbtInd           CreditDebitCode                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 CdtDbtInd,omitempty"`
	OrgnlCcyAndOrdrdAmt RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 OrgnlCcyAndOrdrdAmt,omitempty"`
	FXDtls              ForeignExchangeTerms23                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 FXDtls,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

type ClassificationType33Choice struct {
	ClssfctnFinInstrm CFIOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification86 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 AltrnClssfctn"`
}

type CorporateActionEventType101Choice struct {
	Cd    CorporateActionEventType33Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Cd"`
	Prtry GenericIdentification47        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH, ACCU, MTNG, INFO, TNDP
type CorporateActionEventType33Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 DtTm"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 ToDtTm"`
}

type Document struct {
	IntraPosMvmntPstngRpt IntraPositionMovementPostingReport002V08 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 IntraPosMvmntPstngRpt"`
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

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentAttributes95 struct {
	PlcOfListg             MarketIdentification4Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PlcOfListg,omitempty"`
	DayCntBsis             InterestComputationMethodFormat5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 DayCntBsis,omitempty"`
	RegnForm               FormOfSecurity7Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 RegnForm,omitempty"`
	PmtFrqcy               Frequency27Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PmtFrqcy,omitempty"`
	PmtSts                 SecuritiesPaymentStatus6Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PmtSts,omitempty"`
	VarblRateChngFrqcy     Frequency27Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 VarblRateChngFrqcy,omitempty"`
	ClssfctnTp             ClassificationType33Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 ClssfctnTp,omitempty"`
	OptnStyle              OptionStyle9Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 OptnStyle,omitempty"`
	OptnTp                 OptionType7Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 OptnTp,omitempty"`
	DnmtnCcy               ActiveOrHistoricCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 DnmtnCcy,omitempty"`
	CpnDt                  ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 CpnDt,omitempty"`
	XpryDt                 ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 XpryDt,omitempty"`
	FltgRateFxgDt          ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 FltgRateFxgDt,omitempty"`
	MtrtyDt                ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 MtrtyDt,omitempty"`
	IsseDt                 ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 IsseDt,omitempty"`
	NxtCllblDt             ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 NxtCllblDt,omitempty"`
	PutblDt                ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PutblDt,omitempty"`
	DtdDt                  ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 DtdDt,omitempty"`
	FrstPmtDt              ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 FrstPmtDt,omitempty"`
	PrvsFctr               float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PrvsFctr,omitempty"`
	CurFctr                float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 CurFctr,omitempty"`
	NxtFctr                float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 NxtFctr,omitempty"`
	IntrstRate             float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 IntrstRate,omitempty"`
	YldToMtrtyRate         float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 YldToMtrtyRate,omitempty"`
	NxtIntrstRate          float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 NxtIntrstRate,omitempty"`
	IndxRateBsis           float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 IndxRateBsis,omitempty"`
	CpnAttchdNb            Number23Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 CpnAttchdNb,omitempty"`
	PoolNb                 GenericIdentification39                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PoolNb,omitempty"`
	QtyBrkdwn              []QuantityBreakdown33                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 QtyBrkdwn,omitempty"`
	VarblRateInd           bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 VarblRateInd,omitempty"`
	CllblInd               bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 CllblInd,omitempty"`
	PutblInd               bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PutblInd,omitempty"`
	MktOrIndctvPric        PriceType5Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 MktOrIndctvPric,omitempty"`
	ExrcPric               Price3                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 ExrcPric,omitempty"`
	SbcptPric              Price3                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SbcptPric,omitempty"`
	ConvsPric              Price3                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 ConvsPric,omitempty"`
	StrkPric               Price3                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 StrkPric,omitempty"`
	MinNmnlQty             FinancialInstrumentQuantity15Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 MinNmnlQty,omitempty"`
	CtrctSz                FinancialInstrumentQuantity15Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 CtrctSz,omitempty"`
	UndrlygFinInstrmId     []SecurityIdentification32             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 UndrlygFinInstrmId,omitempty"`
	FinInstrmAttrAddtlDtls RestrictedFINXMax350Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 FinInstrmAttrAddtlDtls,omitempty"`
}

type FinancialInstrumentDetails36 struct {
	FinInstrmId      SecurityIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 FinInstrmId"`
	FinInstrmAttrbts FinancialInstrumentAttributes95 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 FinInstrmAttrbts,omitempty"`
	SubBal           []IntraPositionDetails54        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SubBal"`
}

type FinancialInstrumentQuantity15Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 AmtsdVal"`
}

type ForeignExchangeTerms23 struct {
	UnitCcy  ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 UnitCcy"`
	QtdCcy   ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 QtdCcy"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 XchgRate"`
	RsltgAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 RsltgAmt"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type FormOfSecurity7Choice struct {
	Cd    FormOfSecurity1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

type Frequency26Choice struct {
	Cd    EventFrequency4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

type Frequency27Choice struct {
	Cd    EventFrequency3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

type GenericIdentification18 struct {
	Id      RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SchmeNm,omitempty"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Issr,omitempty"`
}

type GenericIdentification39 struct {
	Id   RestrictedFINMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id"`
	Issr RestrictedFINMax8Text  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Issr,omitempty"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SchmeNm,omitempty"`
}

type GenericIdentification84 struct {
	Id      RestrictedFINXMax34Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SchmeNm,omitempty"`
}

type GenericIdentification85 struct {
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Tp"`
	Id RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id,omitempty"`
}

type GenericIdentification86 struct {
	Id      RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

type IdentificationSource4Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Cd"`
	Prtry RestrictedFINExact2Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat5Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Cd"`
	Prtry GenericIdentification47        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

type IntraPositionDetails54 struct {
	SfkpgPlc      SafekeepingPlaceFormat39Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SfkpgPlc,omitempty"`
	BalFr         SecuritiesBalanceType8Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 BalFr"`
	IntraPosMvmnt []IntraPositionMovementDetails18 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 IntraPosMvmnt"`
}

type IntraPositionMovementDetails18 struct {
	Id                 References51Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id,omitempty"`
	SttldQty           FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SttldQty"`
	PrevslySttldQty    FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PrevslySttldQty,omitempty"`
	RmngToBeSttldQty   FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 RmngToBeSttldQty,omitempty"`
	SctiesSubBalId     GenericIdentification39             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SctiesSubBalId,omitempty"`
	BalTo              SecuritiesBalanceType8Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 BalTo"`
	SttlmDt            DateAndDateTime2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SttlmDt"`
	AvlblDt            DateAndDateTime2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 AvlblDt,omitempty"`
	AckdStsTmStmp      ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 AckdStsTmStmp,omitempty"`
	CorpActnEvtTp      CorporateActionEventType101Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 CorpActnEvtTp,omitempty"`
	CollMntrAmt        AmountAndDirection55                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 CollMntrAmt,omitempty"`
	InstrPrcgAddtlDtls RestrictedFINXMax350Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 InstrPrcgAddtlDtls,omitempty"`
	SplmtryData        []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SplmtryData,omitempty"`
}

type IntraPositionMovementPostingReport002V08 struct {
	Pgntn       Pagination1                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Pgntn"`
	StmtGnlDtls Statement81                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 StmtGnlDtls"`
	AcctOwnr    PartyIdentification136Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount30            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SfkpgAcct"`
	FinInstrm   []FinancialInstrumentDetails36 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 FinInstrm,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MarketIdentification4Choice struct {
	MktIdrCd MICIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 MktIdrCd"`
	Desc     RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Desc"`
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max70Text string

type Number23Choice struct {
	Shrt Exact3NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Shrt"`
	Lng  GenericIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Lng"`
}

type Number3Choice struct {
	Shrt Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Shrt"`
	Lng  Exact5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Lng"`
}

// May be one of AMER, EURO
type OptionStyle2Code string

type OptionStyle9Choice struct {
	Cd    OptionStyle2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

// May be one of CALL, PUTO
type OptionType1Code string

type OptionType7Choice struct {
	Cd    OptionType1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

type OtherIdentification2 struct {
	Id  RestrictedFINXMax31Text     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Sfx,omitempty"`
	Tp  IdentificationSource4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Tp"`
}

type OtherIdentification3 struct {
	Id  RestrictedFINXMax31Text     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Tp"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 LastPgInd"`
}

type PartyIdentification136Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 AnyBIC"`
	PrtryId GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PrtryId"`
}

type Period2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 ToDt"`
}

type Period7Choice struct {
	FrDtTmToDtTm DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 FrDtTmToDtTm"`
	FrDtToDt     Period2         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 FrDtToDt"`
}

type Price3 struct {
	Tp  YieldedOrValueType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Tp"`
	Val PriceRateOrAmount1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Val"`
}

type PriceRateOrAmount1Choice struct {
	Rate float64                                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Rate"`
	Amt  RestrictedFINActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Amt"`
}

type PriceType5Choice struct {
	Mkt    Price3 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Mkt"`
	Indctv Price3 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Indctv"`
}

// May be one of DISC, PREM, PARV
type PriceValueType1Code string

type QuantityBreakdown33 struct {
	LotNb  GenericIdentification39             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 LotNb"`
	LotQty FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 LotQty,omitempty"`
}

type References51Choice struct {
	AcctOwnrTxId      RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 AcctOwnrTxId"`
	AcctSvcrTxId      RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 AcctSvcrTxId"`
	PoolId            RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PoolId"`
	MktInfrstrctrTxId RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 MktInfrstrctrTxId"`
	PrcrTxId          RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PrcrTxId"`
}

type RestrictedFINActiveOrHistoricCurrencyAnd13DecimalAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type RestrictedFINActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern XX|TS
type RestrictedFINExact2Text string

// Must match the pattern ([^/]+/)+([^/]+)|([^/]*)
type RestrictedFINMax30Text string

// Must match the pattern ([^/]+/)+([^/]+)|([^/]*)
type RestrictedFINMax8Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,140}
type RestrictedFINXMax140Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax16Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax30Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,31}
type RestrictedFINXMax31Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax34Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,350}
type RestrictedFINXMax350Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,35}
type RestrictedFINXMax35Text string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat39Choice struct {
	Id      SafekeepingPlaceTypeAndText15          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id"`
	Ctry    CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Ctry"`
	TpAndId SafekeepingPlaceTypeAndIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 TpAndId"`
	Prtry   GenericIdentification85                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

type SafekeepingPlaceTypeAndIdentification1 struct {
	SfkpgPlcTp SafekeepingPlace1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SfkpgPlcTp"`
	Id         AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id"`
}

type SafekeepingPlaceTypeAndText15 struct {
	SfkpgPlcTp SafekeepingPlace3Code   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 SfkpgPlcTp"`
	Id         RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id,omitempty"`
}

type SecuritiesAccount30 struct {
	Id RestrictedFINXMax35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Id"`
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Nm,omitempty"`
}

// May be one of BLOK, AWAS, AVAI, BLCA, BLOT, BLOV, BORR, COLI, COLO, COLA, LOAN, MARG, PECA, PEDA, PLED, REGO, RSTR, OTHR, TRAN, DRAW, CLEN, DIRT, NOMI, SPOS, UNRG, ISSU, QUAS, LODE
type SecuritiesBalanceType11Code string

type SecuritiesBalanceType8Choice struct {
	Cd    SecuritiesBalanceType11Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Cd"`
	Prtry GenericIdentification47     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

// May be one of FULL, NILL, PART
type SecuritiesPaymentStatus1Code string

type SecuritiesPaymentStatus6Choice struct {
	Cd    SecuritiesPaymentStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Cd"`
	Prtry GenericIdentification47      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

type SecurityIdentification20 struct {
	ISIN   ISINOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 ISIN,omitempty"`
	OthrId []OtherIdentification2   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 OthrId,omitempty"`
	Desc   RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Desc,omitempty"`
}

type SecurityIdentification32 struct {
	ISIN   ISINOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 ISIN,omitempty"`
	OthrId []OtherIdentification3   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 OthrId,omitempty"`
	Desc   RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Desc,omitempty"`
}

type Statement81 struct {
	RptNb     Number3Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 RptNb,omitempty"`
	QryRef    RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 QryRef,omitempty"`
	StmtId    RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 StmtId,omitempty"`
	StmtPrd   Period7Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 StmtPrd"`
	Frqcy     Frequency26Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Frqcy,omitempty"`
	UpdTp     UpdateType16Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 UpdTp,omitempty"`
	ActvtyInd bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 ActvtyInd"`
}

// May be one of COMP, DELT
type StatementUpdateType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type UpdateType16Choice struct {
	Cd    StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Cd"`
	Prtry GenericIdentification47  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Prtry"`
}

type YieldedOrValueType1Choice struct {
	Yldd  bool                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 Yldd"`
	ValTp PriceValueType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.08 ValTp"`
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
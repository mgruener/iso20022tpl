// Code generated by main. DO NOT EDIT.

package semt_013_001_04

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

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

type ClassificationType32Choice struct {
	ClssfctnFinInstrm CFIOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 AltrnClssfctn"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 DtTm"`
}

type Document struct {
	IntraPosMvmntInstr IntraPositionMovementInstructionV04 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 IntraPosMvmntInstr"`
}

type DocumentNumber5Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 LngNb"`
	PrtryNb GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 PrtryNb"`
}

// May be one of YEAR, MNTH, QUTR, SEMI, WEEK
type EventFrequency3Code string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentAttributes63 struct {
	PlcOfListg             MarketIdentification3Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 PlcOfListg,omitempty"`
	DayCntBsis             InterestComputationMethodFormat4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 DayCntBsis,omitempty"`
	RegnForm               FormOfSecurity6Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 RegnForm,omitempty"`
	PmtFrqcy               Frequency23Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 PmtFrqcy,omitempty"`
	PmtSts                 SecuritiesPaymentStatus5Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 PmtSts,omitempty"`
	VarblRateChngFrqcy     Frequency23Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 VarblRateChngFrqcy,omitempty"`
	ClssfctnTp             ClassificationType32Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 ClssfctnTp,omitempty"`
	OptnStyle              OptionStyle8Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 OptnStyle,omitempty"`
	OptnTp                 OptionType6Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 OptnTp,omitempty"`
	DnmtnCcy               ActiveOrHistoricCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 DnmtnCcy,omitempty"`
	CpnDt                  ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 CpnDt,omitempty"`
	XpryDt                 ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 XpryDt,omitempty"`
	FltgRateFxgDt          ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 FltgRateFxgDt,omitempty"`
	MtrtyDt                ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 MtrtyDt,omitempty"`
	IsseDt                 ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 IsseDt,omitempty"`
	NxtCllblDt             ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 NxtCllblDt,omitempty"`
	PutblDt                ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 PutblDt,omitempty"`
	DtdDt                  ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 DtdDt,omitempty"`
	FrstPmtDt              ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 FrstPmtDt,omitempty"`
	PrvsFctr               float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 PrvsFctr,omitempty"`
	CurFctr                float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 CurFctr,omitempty"`
	NxtFctr                float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 NxtFctr,omitempty"`
	IntrstRate             float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 IntrstRate,omitempty"`
	YldToMtrtyRate         float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 YldToMtrtyRate,omitempty"`
	NxtIntrstRate          float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 NxtIntrstRate,omitempty"`
	IndxRateBsis           float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 IndxRateBsis,omitempty"`
	CpnAttchdNb            Number22Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 CpnAttchdNb,omitempty"`
	PoolNb                 GenericIdentification37                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 PoolNb,omitempty"`
	QtyBrkdwn              []QuantityBreakdown31                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 QtyBrkdwn,omitempty"`
	VarblRateInd           bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 VarblRateInd,omitempty"`
	CllblInd               bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 CllblInd,omitempty"`
	PutblInd               bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 PutblInd,omitempty"`
	MktOrIndctvPric        PriceType1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 MktOrIndctvPric,omitempty"`
	ExrcPric               Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 ExrcPric,omitempty"`
	SbcptPric              Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SbcptPric,omitempty"`
	ConvsPric              Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 ConvsPric,omitempty"`
	StrkPric               Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 StrkPric,omitempty"`
	MinNmnlQty             FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 MinNmnlQty,omitempty"`
	CtrctSz                FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 CtrctSz,omitempty"`
	UndrlygFinInstrmId     []SecurityIdentification19             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 UndrlygFinInstrmId,omitempty"`
	FinInstrmAttrAddtlDtls Max350Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 FinInstrmAttrAddtlDtls,omitempty"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 AmtsdVal"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type FormOfSecurity6Choice struct {
	Cd    FormOfSecurity1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Prtry"`
}

type Frequency23Choice struct {
	Cd    EventFrequency3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Prtry"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SchmeNm,omitempty"`
}

type GenericIdentification37 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Id"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Issr,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Id,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type Identification14 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Id"`
}

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Prtry"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat4Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Prtry"`
}

type IntraPositionDetails33 struct {
	Prty               PriorityNumeric4Choice                        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Prty,omitempty"`
	SttlmQty           FinancialInstrumentQuantity1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SttlmQty"`
	SctiesSubBalId     GenericIdentification37                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SctiesSubBalId,omitempty"`
	SttlmDt            DateAndDateTimeChoice                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SttlmDt"`
	BalFr              SecuritiesSubBalanceTypeAndQuantityBreakdown3 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 BalFr"`
	BalTo              SecuritiesSubBalanceTypeAndQuantityBreakdown3 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 BalTo"`
	InstrPrcgAddtlDtls Max350Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 InstrPrcgAddtlDtls,omitempty"`
}

type IntraPositionMovementInstructionV04 struct {
	TxId             Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 TxId"`
	CorpActnEvtId    Identification14                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 CorpActnEvtId,omitempty"`
	NbCounts         NumberCount1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 NbCounts,omitempty"`
	Lnkgs            []Linkages36                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Lnkgs,omitempty"`
	AcctOwnr         PartyIdentification92Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 AcctOwnr,omitempty"`
	SfkpgAcct        SecuritiesAccount24             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SfkpgAcct"`
	SfkpgPlc         SafekeepingPlaceFormat10Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SfkpgPlc,omitempty"`
	FinInstrmId      SecurityIdentification19        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 FinInstrmId"`
	FinInstrmAttrbts FinancialInstrumentAttributes63 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 FinInstrmAttrbts,omitempty"`
	IntraPosDtls     IntraPositionDetails33          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 IntraPosDtls"`
	SplmtryData      []SupplementaryData1            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SplmtryData,omitempty"`
}

type Linkages36 struct {
	PrcgPos ProcessingPosition7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 PrcgPos,omitempty"`
	MsgNb   DocumentNumber5Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 MsgNb,omitempty"`
	Ref     References41Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Ref"`
	RefOwnr PartyIdentification92Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 RefOwnr,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MarketIdentification3Choice struct {
	MktIdrCd MICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 MktIdrCd"`
	Desc     Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Desc"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type Number22Choice struct {
	Shrt Exact3NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Shrt"`
	Lng  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Lng"`
}

type NumberCount1Choice struct {
	CurInstrNb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 CurInstrNb"`
	TtlNb      TotalNumber1      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 TtlNb"`
}

// May be one of AMER, EURO
type OptionStyle2Code string

type OptionStyle8Choice struct {
	Cd    OptionStyle2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Prtry"`
}

// May be one of CALL, PUTO
type OptionType1Code string

type OptionType6Choice struct {
	Cd    OptionType1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Prtry"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Tp"`
}

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 PrtryId"`
}

type Price2 struct {
	Tp  YieldedOrValueType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Tp"`
	Val PriceRateOrAmountChoice   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Val"`
}

type PriceRateOrAmountChoice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Rate"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Amt"`
}

type PriceType1Choice struct {
	Mkt    Price2 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Mkt"`
	Indctv Price2 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Indctv"`
}

// May be one of DISC, PREM, PARV
type PriceValueType1Code string

type PriorityNumeric4Choice struct {
	Nmrc  Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Nmrc"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Prtry"`
}

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

type ProcessingPosition7Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Prtry"`
}

type QuantityBreakdown31 struct {
	LotNb  GenericIdentification37            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 LotNb"`
	LotQty FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 LotQty,omitempty"`
}

type QuantityBreakdown32 struct {
	LotNb          GenericIdentification37            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 LotNb,omitempty"`
	LotQty         FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 LotQty,omitempty"`
	SctiesSubBalTp GenericIdentification30            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SctiesSubBalTp,omitempty"`
}

type References41Choice struct {
	SctiesSttlmTxId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SctiesSttlmTxId"`
	IntraPosMvmntId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 IntraPosMvmntId"`
	IntraBalMvmntId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 IntraBalMvmntId"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 AcctSvcrTxId"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 MktInfrstrctrTxId"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 PoolId"`
	OthrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 OthrTxId"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat10Choice struct {
	Id      SafekeepingPlaceTypeAndText8             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 TpAndId"`
	Prtry   GenericIdentification78                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Id"`
}

type SafekeepingPlaceTypeAndText8 struct {
	SfkpgPlcTp SafekeepingPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Id,omitempty"`
}

type SecuritiesAccount24 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Nm,omitempty"`
}

// May be one of BLOK, AWAS, AVAI, BLCA, BLOT, BLOV, BORR, COLI, COLO, COLA, LOAN, MARG, PECA, PEDA, PLED, REGO, RSTR, OTHR, TRAN, DRAW, CLEN, DIRT, NOMI, SPOS, UNRG, ISSU, QUAS, LODE
type SecuritiesBalanceType11Code string

type SecuritiesBalanceType6Choice struct {
	Cd    SecuritiesBalanceType11Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Prtry"`
}

// May be one of FULL, NILL, PART
type SecuritiesPaymentStatus1Code string

type SecuritiesPaymentStatus5Choice struct {
	Cd    SecuritiesPaymentStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Cd"`
	Prtry GenericIdentification30      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Prtry"`
}

type SecuritiesSubBalanceTypeAndQuantityBreakdown3 struct {
	Tp        SecuritiesBalanceType6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Tp"`
	QtyBrkdwn []QuantityBreakdown32        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 QtyBrkdwn,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Desc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TotalNumber1 struct {
	CurInstrNb     Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 CurInstrNb"`
	TtlOfLkdInstrs Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 TtlOfLkdInstrs"`
}

type YieldedOrValueType1Choice struct {
	Yldd  bool                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Yldd"`
	ValTp PriceValueType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 ValTp"`
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

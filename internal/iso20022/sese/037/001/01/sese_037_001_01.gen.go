// Code generated by main. DO NOT EDIT.

package sese_037_001_01

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

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateIdentification1 struct {
	Id    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Id"`
	IdSrc IdentificationSource1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 IdSrc"`
}

type AlternatePartyIdentification2 struct {
	IdTp    IdentificationType4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 IdTp"`
	Ctry    CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Ctry"`
	AltrnId Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 AltrnId"`
}

type AmountAndDirection7 struct {
	Amt       ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Amt"`
	CdtDbtInd CreditDebitCode         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CdtDbtInd"`
}

type AmountAndDirection9 struct {
	Amt                 ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CdtDbtInd,omitempty"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 OrgnlCcyAndOrdrdAmt,omitempty"`
	FXDtls              ForeignExchangeTerms11            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 FXDtls,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{1,6}
type CFIIdentifier string

type ClassificationType2Choice struct {
	ClssfctnFinInstrm CFIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 AltrnClssfctn"`
}

// May be one of CODU, COPY, DUPL
type CopyDuplicate1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 DtTm"`
}

// May be one of VARI
type DateType3Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	PrtflTrfNtfctn PortfolioTransferNotificationV01 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PrtflTrfNtfctn"`
}

type DocumentIdentification11 struct {
	Id       Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Id"`
	CreDtTm  DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CreDtTm,omitempty"`
	CpyDplct CopyDuplicate1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CpyDplct,omitempty"`
}

// May be one of YEAR, MNTH, QUTR, SEMI, WEEK
type EventFrequency3Code string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{5}
type Exact5NumericText string

type Extension2 struct {
	PlcAndNm   Max350Text         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PlcAndNm,omitempty"`
	XtnsnEnvlp ExtensionEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 XtnsnEnvlp"`
}

type ExtensionEnvelope1 struct {
	Item string `xml:",any"`
}

type FinancialInstrumentAttributes8 struct {
	PlcOfListg             MarketIdentification5                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PlcOfListg,omitempty"`
	DayCntBsis             InterestComputationMethodFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 DayCntBsis,omitempty"`
	RegnForm               FormOfSecurity2Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 RegnForm,omitempty"`
	PmtFrqcy               Frequency3Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PmtFrqcy,omitempty"`
	PmtSts                 SecuritiesPaymentStatus2Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PmtSts,omitempty"`
	PmtDrctn               PaymentDirection2Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PmtDrctn,omitempty"`
	VarblRateChngFrqcy     Frequency3Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 VarblRateChngFrqcy,omitempty"`
	PrefToIncm             PreferenceToIncome2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PrefToIncm,omitempty"`
	ClssfctnTp             ClassificationType2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 ClssfctnTp,omitempty"`
	OptnStyle              OptionStyle4Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 OptnStyle,omitempty"`
	OptnTp                 OptionType2Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 OptnTp,omitempty"`
	DnmtnCcy               ActiveOrHistoricCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 DnmtnCcy,omitempty"`
	CpnDt                  ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CpnDt,omitempty"`
	XpryDt                 ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 XpryDt,omitempty"`
	FltgRateFxgDt          ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 FltgRateFxgDt,omitempty"`
	MtrtyDt                ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 MtrtyDt,omitempty"`
	IsseDt                 ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 IsseDt,omitempty"`
	NxtCllblDt             ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 NxtCllblDt,omitempty"`
	PutblDt                ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PutblDt,omitempty"`
	DtdDt                  ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 DtdDt,omitempty"`
	FrstPmtDt              ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 FrstPmtDt,omitempty"`
	PrvsFctr               float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PrvsFctr,omitempty"`
	CurFctr                float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CurFctr,omitempty"`
	NxtFctr                float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 NxtFctr,omitempty"`
	IntrstRate             float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 IntrstRate,omitempty"`
	NxtIntrstRate          float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 NxtIntrstRate,omitempty"`
	IndxRateBsis           float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 IndxRateBsis,omitempty"`
	CpnAttchdNb            Number2Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CpnAttchdNb,omitempty"`
	PoolNb                 Number2Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PoolNb,omitempty"`
	VarblRateInd           bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 VarblRateInd,omitempty"`
	CllblInd               bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CllblInd,omitempty"`
	PutblInd               bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PutblInd,omitempty"`
	MktOrIndctvPric        PriceType1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 MktOrIndctvPric,omitempty"`
	ExrcPric               Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 ExrcPric,omitempty"`
	SbcptPric              Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SbcptPric,omitempty"`
	ConvsPric              Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 ConvsPric,omitempty"`
	StrkPric               Price2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 StrkPric,omitempty"`
	MinNmnlQty             FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 MinNmnlQty,omitempty"`
	CtrctSz                FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CtrctSz,omitempty"`
	UndrlygFinInstrmId     []SecurityIdentification11             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 UndrlygFinInstrmId,omitempty"`
	FinInstrmAttrAddtlDtls Max350Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 FinInstrmAttrAddtlDtls,omitempty"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 AmtsdVal"`
}

type ForeignExchangeTerms11 struct {
	UnitCcy  ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 UnitCcy"`
	QtdCcy   ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 QtdCcy"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 XchgRate"`
	RsltgAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 RsltgAmt"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type FormOfSecurity2Choice struct {
	Cd    FormOfSecurity1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

type Frequency3Choice struct {
	Cd    EventFrequency3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Issr,omitempty"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SchmeNm,omitempty"`
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
	Dmst  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Dmst"`
	Prtry Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

type IdentificationType4Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat1Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MarketIdentification1Choice struct {
	MktIdrCd MICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 MktIdrCd"`
	Desc     Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Desc"`
}

type MarketIdentification5 struct {
	Id MarketIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Id,omitempty"`
	Tp MarketType2Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Tp"`
}

type MarketType2Choice struct {
	Cd    MarketType5Code         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

// May be one of OTCO, EXCH
type MarketType5Code string

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

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Adr,omitempty"`
}

type Number2Choice struct {
	Shrt Exact3NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Shrt"`
	Lng  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Lng"`
}

type Number3Choice struct {
	Shrt Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Shrt"`
	Lng  Exact5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Lng"`
}

// May be one of AMER, EURO
type OptionStyle2Code string

type OptionStyle4Choice struct {
	Cd    OptionStyle2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

// May be one of CALL, PUTO
type OptionType1Code string

type OptionType2Choice struct {
	Cd    OptionType1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

type OtherAmounts2 struct {
	AcrdIntrstAmt  AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 AcrdIntrstAmt,omitempty"`
	ChrgsFees      AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 ChrgsFees,omitempty"`
	CtryNtlFdrlTax AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CtryNtlFdrlTax,omitempty"`
	PmtLevyTax     AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PmtLevyTax,omitempty"`
	LclTax         AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 LclTax,omitempty"`
	Othr           AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Othr,omitempty"`
	PstgAmt        AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PstgAmt,omitempty"`
	RgltryAmt      AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 RgltryAmt,omitempty"`
	ShppgAmt       AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 ShppgAmt,omitempty"`
	StmpDty        AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 StmpDty,omitempty"`
	StockXchgTax   AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 StockXchgTax,omitempty"`
	TrfTax         AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 TrfTax,omitempty"`
	TxTax          AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 TxTax,omitempty"`
	ValAddedTax    AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 ValAddedTax,omitempty"`
	WhldgTax       AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 WhldgTax,omitempty"`
	CsmptnTax      AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CsmptnTax,omitempty"`
	AcrdCptlstnAmt AmountAndDirection9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 AcrdCptlstnAmt,omitempty"`
}

type OtherParties4 struct {
	Invstr    PartyIdentification14Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Invstr,omitempty"`
	StockXchg PartyIdentification10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 StockXchg,omitempty"`
	TradRgltr PartyIdentification10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 TradRgltr,omitempty"`
}

// May be one of A144, NRST, RSTR
type OwnershipLegalRestrictions1Code string

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 LastPgInd"`
}

type PartyIdentification10Choice struct {
	BICOrBEI AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 BICOrBEI"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 NmAndAdr"`
}

type PartyIdentification12Choice struct {
	BICOrBEI AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 BICOrBEI"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Ctry"`
}

type PartyIdentification13Choice struct {
	BICOrBEI AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 BICOrBEI"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PrtryId"`
}

type PartyIdentification14Choice struct {
	BICOrBEI AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 BICOrBEI"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 NmAndAdr"`
	Ctry     CountryCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Ctry"`
}

type PartyIdentification2 struct {
	Id       PartyIdentification12Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Id"`
	AltrnId  AlternatePartyIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 AltrnId,omitempty"`
	PrcgDt   DateAndDateTimeChoice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PrcgDt,omitempty"`
	PrcgId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PrcgId,omitempty"`
	AddtlInf PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 AddtlInf,omitempty"`
}

type PartyIdentificationAndAccount1 struct {
	Id        PartyIdentification10Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Id"`
	AltrnId   AlternatePartyIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 AltrnId,omitempty"`
	SfkpgAcct SecuritiesAccount13           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SfkpgAcct,omitempty"`
	PrcgDt    DateAndDateTimeChoice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PrcgDt,omitempty"`
	PrcgId    Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PrcgId,omitempty"`
	AddtlInf  PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 AddtlInf,omitempty"`
}

type PartyTextInformation1 struct {
	DclrtnDtls  Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 DclrtnDtls,omitempty"`
	PtyCtctDtls Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PtyCtctDtls,omitempty"`
	RegnDtls    Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 RegnDtls,omitempty"`
}

type PaymentDirection2Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

type PortfolioTransferNotificationV01 struct {
	Id            DocumentIdentification11    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Id"`
	Pgntn         Pagination                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Pgntn"`
	StmtGnlDtls   Statement19                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 StmtGnlDtls"`
	AcctOwnr      PartyIdentification13Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 AcctOwnr,omitempty"`
	SfkpgAcct     SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SfkpgAcct"`
	TrfNtfctnDtls []SecuritiesTradeDetails7   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 TrfNtfctnDtls,omitempty"`
	MsgOrgtr      PartyIdentification10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 MsgOrgtr,omitempty"`
	MsgRcpt       PartyIdentification10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 MsgRcpt,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Ctry"`
}

// May be one of ORDN, PFRD
type PreferenceToIncome1Code string

type PreferenceToIncome2Choice struct {
	Cd    PreferenceToIncome1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

type Price2 struct {
	Tp  YieldedOrValueType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Tp"`
	Val PriceRateOrAmountChoice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Val"`
}

type PriceRateOrAmountChoice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Rate"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Amt"`
}

type PriceType1Choice struct {
	Mkt    Price2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Mkt"`
	Indctv Price2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Indctv"`
}

// May be one of DISC, PREM, PARV
type PriceValueType1Code string

type Quantity5 struct {
	SttlmQty  FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SttlmQty"`
	DnmtnChc  Max210Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 DnmtnChc,omitempty"`
	CertNb    []SecuritiesCertificate1           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CertNb,omitempty"`
	QtyBrkdwn []QuantityBreakdown3               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 QtyBrkdwn,omitempty"`
}

type QuantityBreakdown3 struct {
	LotNb    Number2Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 LotNb,omitempty"`
	LotQty   FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 LotQty,omitempty"`
	LotDtTm  DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 LotDtTm,omitempty"`
	LotPric  Price2                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 LotPric,omitempty"`
	TpOfPric TypeOfPrice3Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 TpOfPric,omitempty"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type Registration1Choice struct {
	Cd    Registration1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

// May be one of NREG, YREG
type Registration1Code string

type Reporting1Choice struct {
	Cd    Reporting1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

// May be one of STEX, REGU
type Reporting1Code string

type Restriction1Choice struct {
	Cd    OwnershipLegalRestrictions1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Nm,omitempty"`
}

type SecuritiesCertificate1 struct {
	Nb      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Nb"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Issr,omitempty"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SchmeNm,omitempty"`
}

// May be one of FULL, NILL, PART
type SecuritiesPaymentStatus1Code string

type SecuritiesPaymentStatus2Choice struct {
	Cd    SecuritiesPaymentStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

type SecuritiesRTGS1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

type SecuritiesTradeDetails7 struct {
	NtfctnSndrTxId   Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 NtfctnSndrTxId,omitempty"`
	NtfctnRcvrTxId   Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 NtfctnRcvrTxId,omitempty"`
	CmonId           Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CmonId,omitempty"`
	SctiesMvmntTp    ReceiveDelivery1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SctiesMvmntTp"`
	Pmt              DeliveryReceiptType2Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Pmt"`
	TradDt           TradeDate1Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 TradDt,omitempty"`
	SttlmDt          SettlementDate1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SttlmDt"`
	NbOfDaysAcrd     float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 NbOfDaysAcrd,omitempty"`
	FinInstrmId      SecurityIdentification11       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 FinInstrmId"`
	FinInstrmAttrbts FinancialInstrumentAttributes8 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 FinInstrmAttrbts,omitempty"`
	Rptg             []Reporting1Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Rptg,omitempty"`
	QtyDtls          Quantity5                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 QtyDtls"`
	SttlmParams      SettlementDetails4             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SttlmParams,omitempty"`
	DlvrgSttlmPties  SettlementParties5             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties   SettlementParties5             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 RcvgSttlmPties,omitempty"`
	SttlmAmt         AmountAndDirection7            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SttlmAmt,omitempty"`
	OthrAmts         OtherAmounts2                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 OthrAmts,omitempty"`
	OthrBizPties     OtherParties4                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 OthrBizPties,omitempty"`
	Xtnsn            []Extension2                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Xtnsn,omitempty"`
}

type SecurityIdentification11 struct {
	Id   SecurityIdentification11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Id"`
	Desc Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Desc,omitempty"`
}

type SecurityIdentification11Choice struct {
	ISIN   ISINIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 ISIN"`
	OthrId AlternateIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 OthrId"`
}

type SettlementDate1Choice struct {
	Dt   DateAndDateTimeChoice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Dt"`
	DtCd SettlementDateCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 DtCd"`
}

// May be one of WISS
type SettlementDate4Code string

type SettlementDateCode1Choice struct {
	Cd    SettlementDate4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

type SettlementDetails4 struct {
	SttlmTxCond    []SettlementTransactionCondition2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SttlmTxCond,omitempty"`
	Regn           Registration1Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Regn,omitempty"`
	LglRstrctns    Restriction1Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 LglRstrctns,omitempty"`
	SctiesRTGS     SecuritiesRTGS1Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SctiesRTGS,omitempty"`
	SttlmSysMtd    SettlementSystemMethod1Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 SttlmSysMtd,omitempty"`
	TaxCpcty       TaxCapacityParty1Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 TaxCpcty,omitempty"`
	StmpDtyTaxBsis GenericIdentification20                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 StmpDtyTaxBsis,omitempty"`
}

type SettlementParties5 struct {
	Dpstry PartyIdentification2           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Pty5,omitempty"`
}

type SettlementSystemMethod1Choice struct {
	Cd    SettlementSystemMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

// May be one of NSET, YSET
type SettlementSystemMethod1Code string

type SettlementTransactionCondition2Choice struct {
	Cd    SettlementTransactionCondition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

// May be one of ASGN, CLEN, DIRT, DLWM, DRAW, EXER, FRCL, KNOC, PHYS, RESI, SPDL, SPST, UNEX
type SettlementTransactionCondition3Code string

type Statement19 struct {
	CtrPtyPrtflTrfNtfctnRef Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 CtrPtyPrtflTrfNtfctnRef,omitempty"`
	RptNb                   Number3Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 RptNb,omitempty"`
	StmtId                  Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 StmtId,omitempty"`
	StmtDtTm                DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 StmtDtTm"`
	UpdTp                   UpdateType2Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 UpdTp,omitempty"`
	ActvtyInd               bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 ActvtyInd"`
}

// May be one of COMP, DELT
type StatementUpdateType1Code string

type TaxCapacityParty1Choice struct {
	Cd    TaxLiability1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

// May be one of PRIN, AGEN
type TaxLiability1Code string

type TradeDate1Choice struct {
	Dt   DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Dt"`
	DtCd TradeDateCode1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 DtCd"`
}

type TradeDateCode1Choice struct {
	Cd    DateType3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

// May be one of ARNU, CCPT, CHTY, CORP, DRLC, FIIN, TXID
type TypeOfIdentification1Code string

// May be one of AVER
type TypeOfPrice14Code string

type TypeOfPrice3Choice struct {
	Cd    TypeOfPrice14Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

type UpdateType2Choice struct {
	Cd    StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Cd"`
	Prtry GenericIdentification20  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Prtry"`
}

type YieldedOrValueType1Choice struct {
	Yldd  bool                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Yldd"`
	ValTp PriceValueType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 ValTp"`
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
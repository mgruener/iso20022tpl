// Code generated by main. DO NOT EDIT.

package secl_004_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

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

type AlternatePartyIdentification4 struct {
	IdTp    IdentificationType6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 IdTp"`
	Ctry    CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Ctry"`
	AltrnId Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AltrnId"`
}

type AlternatePartyIdentification5 struct {
	IdTp    IdentificationType40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 IdTp"`
	Ctry    CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Ctry"`
	AltrnId Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AltrnId"`
}

type AmountAndDirection21 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 CdtDbtInd,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of HOUS, CLIE, LIPR
type ClearingAccountType1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 DtTm"`
}

type DateCode3Choice struct {
	Cd    DateType1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Prtry"`
}

type DateFormat15Choice struct {
	Dt   ISODate         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Dt"`
	DtCd DateCode3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 DtCd"`
}

// May be one of UKWN
type DateType1Code string

type Document struct {
	NetPos NetPositionV03 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 NetPos"`
}

// May be one of DAIL, INDA, ONDE
type EventFrequency6Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{5}
type Exact5NumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AmtsdVal"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SchmeNm,omitempty"`
}

type GenericIdentification29 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SchmeNm,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SchmeNm,omitempty"`
}

type GenericIdentification40 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SchmeNm,omitempty"`
}

type GenericIdentification58 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id,omitempty"`
	Tp GenericIdentification40 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Tp"`
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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Prtry"`
}

type IdentificationType40Choice struct {
	Cd    TypeOfIdentification2Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Cd"`
	Prtry GenericIdentification29   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Prtry"`
}

type IdentificationType6Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Cd"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MarketIdentification1Choice struct {
	MktIdrCd MICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 MktIdrCd"`
	Desc     Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Desc"`
}

type MarketIdentification20 struct {
	Id MarketIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id,omitempty"`
	Tp MarketType8Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Tp"`
}

type MarketIdentification84 struct {
	Id MarketIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id,omitempty"`
	Tp MarketType8Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Tp"`
}

type MarketIdentification85 struct {
	Id MarketIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id,omitempty"`
	Tp MarketType9Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Tp"`
}

// May be one of PRIM, SECM, OTCO, VARI, EXCH
type MarketType2Code string

// May be one of OTCO, EXCH
type MarketType5Code string

type MarketType8Choice struct {
	Cd    MarketType2Code         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Prtry"`
}

type MarketType9Choice struct {
	Cd    MarketType5Code         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Prtry"`
}

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

type NameAndAddress13 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Nm"`
	Adr PostalAddress8 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Adr,omitempty"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Adr,omitempty"`
}

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Adr"`
}

type NetPosition3 struct {
	ClrAcct       SecuritiesAccount18                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 ClrAcct"`
	NonClrMmb     PartyIdentificationAndAccount31    `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 NonClrMmb,omitempty"`
	DlvryAcct     SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 DlvryAcct,omitempty"`
	FinInstrmId   SecurityIdentification14           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 FinInstrmId"`
	InitlPosAmt   AmountAndDirection21               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 InitlPosAmt,omitempty"`
	NetPosAmt     AmountAndDirection21               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 NetPosAmt"`
	AcrdIntrstAmt AmountAndDirection21               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AcrdIntrstAmt,omitempty"`
	AvrgDealPric  Price4                             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AvrgDealPric,omitempty"`
	NetQty        FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 NetQty"`
	SctiesMvmntTp ReceiveDelivery1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SctiesMvmntTp"`
	Dpstry        PartyIdentification34Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Dpstry"`
	TradgCpcty    TradingCapacity5Code               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradgCpcty,omitempty"`
	PlcOfTrad     MarketIdentification20             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PlcOfTrad,omitempty"`
	TradDt        ISODate                            `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradDt,omitempty"`
	SttlmDt       DateFormat15Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SttlmDt,omitempty"`
	TradLegDtls   []TradeLeg10                       `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradLegDtls,omitempty"`
}

type NetPositionV03 struct {
	RptParams   ReportParameters1           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 RptParams"`
	Pgntn       Pagination                  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Pgntn"`
	ClrMmb      PartyIdentification35Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 ClrMmb"`
	ClrSgmt     PartyIdentification35Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 ClrSgmt,omitempty"`
	NetPosRpt   []NetPosition3              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 NetPosRpt"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SplmtryData,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Tp"`
}

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 LastPgInd"`
}

type PartyIdentification33Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AnyBIC"`
	PrtryId  GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PrtryId"`
	NmAndAdr NameAndAddress6         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 NmAndAdr"`
}

type PartyIdentification34Choice struct {
	BIC      AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 BIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Ctry"`
}

type PartyIdentification35Choice struct {
	BIC     AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 BIC"`
	PrtryId GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PrtryId"`
}

type PartyIdentification83Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AnyBIC"`
	PrtryId  GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PrtryId"`
	NmAndAdr NameAndAddress13        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 NmAndAdr"`
}

type PartyIdentificationAndAccount100 struct {
	Id        PartyIdentification83Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id"`
	AltrnId   AlternatePartyIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AltrnId,omitempty"`
	SfkpgAcct Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SfkpgAcct,omitempty"`
	PrcgId    Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PrcgId,omitempty"`
	AddtlInf  PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AddtlInf,omitempty"`
}

type PartyIdentificationAndAccount31 struct {
	Id       PartyIdentification33Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id"`
	AltrnId  AlternatePartyIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AltrnId,omitempty"`
	AddtlInf PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AddtlInf,omitempty"`
	ClrAcct  SecuritiesAccount18           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 ClrAcct,omitempty"`
}

type PartyTextInformation1 struct {
	DclrtnDtls  Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 DclrtnDtls,omitempty"`
	PtyCtctDtls Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PtyCtctDtls,omitempty"`
	RegnDtls    Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 RegnDtls,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Ctry"`
}

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Ctry"`
}

type PostalAddress8 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Ctry"`
}

type Price4 struct {
	Val PriceRateOrAmountChoice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Val"`
	Tp  PriceValueType7Code     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Tp,omitempty"`
}

type PriceRateOrAmountChoice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Rate"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Amt"`
}

// May be one of DISC, PREM, PARV, YIEL, SPRE, PEUN, ABSO, TEDP, TEDY, FICT, VACT, PRCT, ACTU
type PriceValueType7Code string

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type ReportParameters1 struct {
	NetPosId   Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 NetPosId"`
	RptDtAndTm DateAndDateTimeChoice    `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 RptDtAndTm"`
	UpdTp      StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 UpdTp"`
	Frqcy      EventFrequency6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Frqcy"`
	RptNb      Exact5NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 RptNb,omitempty"`
	ActvtyInd  bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 ActvtyInd"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat7Choice struct {
	Id      SafekeepingPlaceTypeAndText1             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TpAndId"`
	Prtry   GenericIdentification58                  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id"`
}

type SafekeepingPlaceTypeAndText1 struct {
	SfkpgPlcTp SafekeepingPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id,omitempty"`
}

type SecuritiesAccount18 struct {
	Id Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id"`
	Tp ClearingAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Tp"`
	Nm Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Nm,omitempty"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Nm,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Desc,omitempty"`
}

// May be one of BUYI, SELL, TWOS, BUMI, SEPL, SESH, SSEX, CROS, CRSH, CSHE, DEFI, OPPO, UNDI
type Side1Code string

// May be one of COMP, DELT
type StatementUpdateType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeLeg10 struct {
	TradLegId     Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradLegId"`
	TradId        Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradId,omitempty"`
	TradExctnId   Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradExctnId"`
	OrdrId        Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 OrdrId,omitempty"`
	AllcnId       Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 AllcnId,omitempty"`
	TradDt        ISODate                            `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradDt"`
	TxDtAndTm     ISODateTime                        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TxDtAndTm,omitempty"`
	SttlmDt       DateFormat15Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SttlmDt"`
	TradgCcy      CurrencyCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradgCcy,omitempty"`
	BuySellInd    Side1Code                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 BuySellInd"`
	TradQty       FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradQty"`
	DealPric      Price4                             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 DealPric"`
	GrssAmt       AmountAndDirection21               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 GrssAmt,omitempty"`
	PlcOfTrad     MarketIdentification84             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PlcOfTrad"`
	PlcOfListg    MarketIdentification85             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 PlcOfListg,omitempty"`
	TradTp        TradeType1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradTp"`
	DerivRltdTrad bool                               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 DerivRltdTrad,omitempty"`
	Brkr          PartyIdentificationAndAccount100   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 Brkr,omitempty"`
	TradgPty      PartyIdentification35Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradgPty"`
	TradRegnOrgn  Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradRegnOrgn,omitempty"`
	TradgPtyAcct  SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradgPtyAcct,omitempty"`
	TradgCpcty    TradingCapacity5Code               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradgCpcty"`
	TradPstngCd   TradePosting1Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 TradPstngCd,omitempty"`
	SfkpgPlc      SafekeepingPlaceFormat7Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SfkpgPlc,omitempty"`
	SfkpgAcct     SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.004.001.03 SfkpgAcct,omitempty"`
}

// May be one of GROS, NETT
type TradePosting1Code string

// May be one of OOBK, OFBK, BKTR, COTR, GUTR, LKTR
type TradeType1Code string

// May be one of PRIN, RISP, AGEN
type TradingCapacity5Code string

// May be one of ARNU, CCPT, CHTY, CORP, DRLC, FIIN, TXID
type TypeOfIdentification1Code string

// May be one of ARNU, CHTY, CORP, FIIN, TXID
type TypeOfIdentification2Code string

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

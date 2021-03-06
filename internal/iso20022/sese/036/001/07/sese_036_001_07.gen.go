// Code generated by main. DO NOT EDIT.

package sese_036_001_07

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

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternatePartyIdentification7 struct {
	IdTp    IdentificationType42Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 IdTp"`
	Ctry    CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Ctry"`
	AltrnId Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AltrnId"`
}

type AmountAndDirection21 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 CdtDbtInd,omitempty"`
}

type AmountAndDirection49 struct {
	Amt                 ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 CdtDbtInd,omitempty"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 OrgnlCcyAndOrdrdAmt,omitempty"`
	FXDtls              ForeignExchangeTerms23            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 FXDtls,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// May be one of LAMI, NBOR, YBOR
type AutoBorrowing1Code string

type AutomaticBorrowing6Choice struct {
	Cd    AutoBorrowing1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type BeneficialOwnership4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

// May be one of BLPA, BLCH
type BlockTrade1Code string

type BlockTrade4Choice struct {
	Cd    BlockTrade1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type CashAccountIdentification5Choice struct {
	IBAN  IBAN2007Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 IBAN"`
	Prtry Max34Text          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

// May be one of GROS, NETS
type CashSettlementSystem2Code string

type CashSettlementSystem4Choice struct {
	Cd    CashSettlementSystem2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type CentralCounterPartyEligibility4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 DtTm"`
}

type DateCode18Choice struct {
	Cd    DateType5Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

// May be one of OPEN
type DateType5Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesFincgModInstr SecuritiesFinancingModificationInstructionV07 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SctiesFincgModInstr"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FXStandingInstruction4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AmtsdVal"`
}

type ForeignExchangeTerms23 struct {
	UnitCcy  ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 UnitCcy"`
	QtdCcy   ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 QtdCcy"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 XchgRate"`
	RsltgAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 RsltgAmt"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Id,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type IdentificationType42Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat4Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be one of FRAN
type LegalFramework1Code string

type LegalFramework3Choice struct {
	Cd    LegalFramework1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

// May be one of CLNT, MAKT
type MarketClientSide1Code string

type MarketClientSide6Choice struct {
	Cd    MarketClientSide1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max52Text string

// Must be at least 1 items long
type Max70Text string

// Must be at least 1 items long
type Max8Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Adr,omitempty"`
}

type NettingEligibility4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Tp"`
}

// May be one of A144, NRST, RSTR
type OwnershipLegalRestrictions1Code string

type PartyIdentification120Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 NmAndAdr"`
}

type PartyIdentification122Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AnyBIC"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 NmAndAdr"`
	Ctry     CountryCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Ctry"`
}

type PartyIdentification127Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PrtryId"`
}

type PartyIdentification144 struct {
	Id  PartyIdentification127Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 LEI,omitempty"`
}

type PartyIdentification146 struct {
	Id       PartyIdentification122Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Id"`
	LEI      LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 LEI,omitempty"`
	AltrnId  AlternatePartyIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AltrnId,omitempty"`
	PrcgDt   DateAndDateTime2Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PrcgDt,omitempty"`
	PrcgId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PrcgId,omitempty"`
	AddtlInf PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AddtlInf,omitempty"`
}

type PartyIdentificationAndAccount169 struct {
	Id        PartyIdentification120Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Id"`
	LEI       LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 LEI,omitempty"`
	AltrnId   AlternatePartyIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AltrnId,omitempty"`
	SfkpgAcct SecuritiesAccount22           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SfkpgAcct,omitempty"`
	PrcgDt    DateAndDateTime2Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PrcgDt,omitempty"`
	PrcgId    Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PrcgId,omitempty"`
	AddtlInf  PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AddtlInf,omitempty"`
}

type PartyTextInformation1 struct {
	DclrtnDtls  Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 DclrtnDtls,omitempty"`
	PtyCtctDtls Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PtyCtctDtls,omitempty"`
	RegnDtls    Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 RegnDtls,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Ctry"`
}

type PriorityNumeric4Choice struct {
	Nmrc  Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Nmrc"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type QuantityAndAccount78 struct {
	SttlmQty  FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SttlmQty"`
	AcctOwnr  PartyIdentification144             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AcctOwnr,omitempty"`
	SfkpgAcct SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SfkpgAcct"`
	CshAcct   CashAccountIdentification5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 CshAcct,omitempty"`
	SfkpgPlc  SafeKeepingPlace3                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SfkpgPlc,omitempty"`
}

type Rate2 struct {
	Sgn  bool    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Sgn,omitempty"`
	Rate float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Rate"`
}

type RateName1 struct {
	Issr   Max8Text  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Issr,omitempty"`
	RateNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 RateNm"`
}

type RateOrName1Choice struct {
	Rate   Rate2     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Rate"`
	RateNm RateName1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 RateNm"`
}

// May be one of FIXE, FORF, VARI
type RateType1Code string

type RateType35Choice struct {
	Cd    RateType1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type RepurchaseType21Choice struct {
	Cd    RepurchaseType8Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

// May be one of PADJ, ROLP, RATE, CALL
type RepurchaseType8Code string

type Restriction5Choice struct {
	Cd    OwnershipLegalRestrictions1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type RevaluationIndicator3Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type SafeKeepingPlace3 struct {
	SfkpgPlcFrmt SafekeepingPlaceFormat29Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SfkpgPlcFrmt,omitempty"`
	LEI          LEIIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 LEI,omitempty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat29Choice struct {
	Id      SafekeepingPlaceTypeAndText8           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Id"`
	Ctry    CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Ctry"`
	TpAndId SafekeepingPlaceTypeAndIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 TpAndId"`
	Prtry   GenericIdentification78                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type SafekeepingPlaceTypeAndIdentification1 struct {
	SfkpgPlcTp SafekeepingPlace1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SfkpgPlcTp"`
	Id         AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Id"`
}

type SafekeepingPlaceTypeAndText8 struct {
	SfkpgPlcTp SafekeepingPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Id,omitempty"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Nm,omitempty"`
}

type SecuritiesAccount22 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Nm,omitempty"`
}

type SecuritiesFinancingModificationInstructionV07 struct {
	TxTpAndModAddtlParams TransactionTypeAndAdditionalParameters17 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 TxTpAndModAddtlParams"`
	TradDtls              SecuritiesTradeDetails100                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 TradDtls"`
	FinInstrmId           SecurityIdentification19                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 FinInstrmId"`
	QtyAndAcctDtls        QuantityAndAccount78                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 QtyAndAcctDtls"`
	SctiesFincgAddtlDtls  SecuritiesFinancingTransactionDetails42  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SctiesFincgAddtlDtls"`
	SttlmParams           SettlementDetails148                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SttlmParams,omitempty"`
	DlvrgSttlmPties       SettlementParties77                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties        SettlementParties77                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 RcvgSttlmPties,omitempty"`
	OpngSttlmAmt          AmountAndDirection49                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 OpngSttlmAmt,omitempty"`
	SplmtryData           []SupplementaryData1                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SplmtryData,omitempty"`
}

type SecuritiesFinancingTransactionDetails42 struct {
	SctiesFincgTradId     Max52Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SctiesFincgTradId,omitempty"`
	ClsgLegId             Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 ClsgLegId,omitempty"`
	TermntnDt             TerminationDate6Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 TermntnDt,omitempty"`
	RateChngDt            DateAndDateTime2Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 RateChngDt,omitempty"`
	EarlstCallBckDt       DateAndDateTime2Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 EarlstCallBckDt,omitempty"`
	ComssnClctnDt         DateAndDateTime2Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 ComssnClctnDt,omitempty"`
	RateTp                RateType35Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 RateTp,omitempty"`
	Rvaltn                RevaluationIndicator3Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Rvaltn,omitempty"`
	LglFrmwk              LegalFramework3Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 LglFrmwk,omitempty"`
	IntrstCmptnMtd        InterestComputationMethodFormat4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 IntrstCmptnMtd,omitempty"`
	MtrtyDtMod            bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 MtrtyDtMod,omitempty"`
	IntrstPmt             bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 IntrstPmt,omitempty"`
	VarblRateSpprt        RateName1                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 VarblRateSpprt,omitempty"`
	RpRate                Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 RpRate,omitempty"`
	StockLnMrgn           Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 StockLnMrgn,omitempty"`
	SctiesHrcut           Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SctiesHrcut,omitempty"`
	ChrgsRate             Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 ChrgsRate,omitempty"`
	PricgRate             RateOrName1Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PricgRate,omitempty"`
	Sprd                  Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Sprd,omitempty"`
	TxCallDely            Exact3NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 TxCallDely,omitempty"`
	TtlNbOfCollInstrs     Exact3NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 TtlNbOfCollInstrs,omitempty"`
	LclBrkrComssn         AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 LclBrkrComssn,omitempty"`
	DealAmt               AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 DealAmt,omitempty"`
	AcrdIntrstAmt         AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AcrdIntrstAmt,omitempty"`
	FrftAmt               AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 FrftAmt,omitempty"`
	PrmAmt                AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PrmAmt,omitempty"`
	TermntnAmtPerPcOfColl AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 TermntnAmtPerPcOfColl,omitempty"`
	TermntnTxAmt          AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 TermntnTxAmt,omitempty"`
	ScndLegNrrtv          Max140Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 ScndLegNrrtv,omitempty"`
}

// May be one of REPU, RVPO, SECB, SECL, BSBK, SBBK
type SecuritiesFinancingTransactionType2Code string

type SecuritiesRTGS4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type SecuritiesTradeDetails100 struct {
	TradDt             DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 TradDt,omitempty"`
	OpngSttlmDt        DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 OpngSttlmDt"`
	NbOfDaysAcrd       float64                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 NbOfDaysAcrd,omitempty"`
	InstrPrcgAddtlDtls Max350Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 InstrPrcgAddtlDtls,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Desc,omitempty"`
}

type SettlementDetails148 struct {
	HldInd         bool                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 HldInd,omitempty"`
	Prty           PriorityNumeric4Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prty,omitempty"`
	SttlmTxCond    []SettlementTransactionCondition18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SttlmTxCond,omitempty"`
	SttlgCpcty     SettlingCapacity7Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SttlgCpcty,omitempty"`
	StmpDtyTaxBsis GenericIdentification30                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 StmpDtyTaxBsis,omitempty"`
	SctiesRTGS     SecuritiesRTGS4Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SctiesRTGS,omitempty"`
	BnfclOwnrsh    BeneficialOwnership4Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 BnfclOwnrsh,omitempty"`
	CshClrSys      CashSettlementSystem4Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 CshClrSys,omitempty"`
	TaxCpcty       TaxCapacityParty4Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 TaxCpcty,omitempty"`
	MktClntSd      MarketClientSide6Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 MktClntSd,omitempty"`
	FxStgInstr     FXStandingInstruction4Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 FxStgInstr,omitempty"`
	BlckTrad       BlockTrade4Choice                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 BlckTrad,omitempty"`
	LglRstrctns    Restriction5Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 LglRstrctns,omitempty"`
	SttlmSysMtd    SettlementSystemMethod4Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SttlmSysMtd,omitempty"`
	NetgElgblty    NettingEligibility4Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 NetgElgblty,omitempty"`
	CCPElgblty     CentralCounterPartyEligibility4Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 CCPElgblty,omitempty"`
	Trckg          Tracking4Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Trckg,omitempty"`
	AutomtcBrrwg   AutomaticBorrowing6Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AutomtcBrrwg,omitempty"`
	PrtlSttlmInd   SettlementTransactionCondition5Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PrtlSttlmInd,omitempty"`
	ElgblForColl   bool                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 ElgblForColl,omitempty"`
}

type SettlementParties77 struct {
	Dpstry PartyIdentification146           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount169 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount169 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount169 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount169 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount169 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Pty5,omitempty"`
}

// May be one of NSET, YSET
type SettlementSystemMethod1Code string

type SettlementSystemMethod4Choice struct {
	Cd    SettlementSystemMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type SettlementTransactionCondition18Choice struct {
	Cd    SettlementTransactionCondition6Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

// May be one of PART, NPAR, PARC, PARQ
type SettlementTransactionCondition5Code string

// May be one of ASGN, BUTC, CLEN, DIRT, DLWM, DRAW, EXER, FRCL, KNOC, PHYS, RESI, SHOR, SPDL, SPST, EXPI, PENS, UNEX, TRIP, NOMC, TRAN, RHYP, ADEA
type SettlementTransactionCondition6Code string

// May be one of SAGE, CUST, SPRI, RISP
type SettlingCapacity2Code string

type SettlingCapacity7Choice struct {
	Cd    SettlingCapacity2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxCapacityParty4Choice struct {
	Cd    TaxLiability1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

// May be one of PRIN, AGEN
type TaxLiability1Code string

type TerminationDate6Choice struct {
	Dt DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Dt"`
	Cd DateCode18Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Cd"`
}

type Tracking4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Prtry"`
}

type TransactionTypeAndAdditionalParameters17 struct {
	AcctOwnrTxId    Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AcctOwnrTxId"`
	AcctSvcrTxId    Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 AcctSvcrTxId,omitempty"`
	SctiesFincgTxTp SecuritiesFinancingTransactionType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 SctiesFincgTxTp"`
	Pmt             DeliveryReceiptType2Code                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 Pmt"`
	ModTp           RepurchaseType21Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 ModTp,omitempty"`
	CmonId          Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 CmonId,omitempty"`
	PoolId          Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.07 PoolId,omitempty"`
}

// May be one of ARNU, CCPT, CHTY, CORP, DRLC, FIIN, TXID
type TypeOfIdentification1Code string

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

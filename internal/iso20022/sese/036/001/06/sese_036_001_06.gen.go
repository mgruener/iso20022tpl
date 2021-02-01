// Code generated by main. DO NOT EDIT.

package sese_036_001_06

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
	IdTp    IdentificationType42Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 IdTp"`
	Ctry    CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Ctry"`
	AltrnId Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AltrnId"`
}

type AmountAndDirection21 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 CdtDbtInd,omitempty"`
}

type AmountAndDirection49 struct {
	Amt                 ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 CdtDbtInd,omitempty"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 OrgnlCcyAndOrdrdAmt,omitempty"`
	FXDtls              ForeignExchangeTerms23            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 FXDtls,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of LAMI, NBOR, YBOR
type AutoBorrowing1Code string

type AutomaticBorrowing6Choice struct {
	Cd    AutoBorrowing1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type BeneficialOwnership4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

// May be one of BLPA, BLCH
type BlockTrade1Code string

type BlockTrade4Choice struct {
	Cd    BlockTrade1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type CashAccountIdentification5Choice struct {
	IBAN  IBAN2007Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 IBAN"`
	Prtry Max34Text          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

// May be one of GROS, NETS
type CashSettlementSystem2Code string

type CashSettlementSystem4Choice struct {
	Cd    CashSettlementSystem2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type CentralCounterPartyEligibility4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 DtTm"`
}

type DateCode18Choice struct {
	Cd    DateType5Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

// May be one of OPEN
type DateType5Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesFincgModInstr SecuritiesFinancingModificationInstructionV06 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SctiesFincgModInstr"`
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
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AmtsdVal"`
}

type ForeignExchangeTerms23 struct {
	UnitCcy  ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 UnitCcy"`
	QtdCcy   ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 QtdCcy"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 XchgRate"`
	RsltgAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 RsltgAmt"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Id,omitempty"`
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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type IdentificationType42Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat4Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be one of FRAN
type LegalFramework1Code string

type LegalFramework3Choice struct {
	Cd    LegalFramework1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type MarketClientSide4Choice struct {
	Cd    MarketClientSideCode    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

// May be one of MAKT, CLNT
type MarketClientSideCode string

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
type Max70Text string

// Must be at least 1 items long
type Max8Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Adr,omitempty"`
}

type NettingEligibility4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Tp"`
}

// May be one of A144, NRST, RSTR
type OwnershipLegalRestrictions1Code string

type PartyIdentification44Choice struct {
	AnyBIC   AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AnyBIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Ctry"`
}

type PartyIdentification71Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 NmAndAdr"`
}

type PartyIdentification75 struct {
	Id       PartyIdentification44Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Id"`
	LEI      LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 LEI,omitempty"`
	AltrnId  AlternatePartyIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AltrnId,omitempty"`
	PrcgDt   DateAndDateTimeChoice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PrcgDt,omitempty"`
	PrcgId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PrcgId,omitempty"`
	AddtlInf PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AddtlInf,omitempty"`
}

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PrtryId"`
}

type PartyIdentification98 struct {
	Id  PartyIdentification92Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 LEI,omitempty"`
}

type PartyIdentificationAndAccount106 struct {
	Id        PartyIdentification71Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Id"`
	LEI       LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 LEI,omitempty"`
	AltrnId   AlternatePartyIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AltrnId,omitempty"`
	SfkpgAcct SecuritiesAccount24           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SfkpgAcct,omitempty"`
	PrcgDt    DateAndDateTimeChoice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PrcgDt,omitempty"`
	PrcgId    Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PrcgId,omitempty"`
	AddtlInf  PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AddtlInf,omitempty"`
}

type PartyTextInformation1 struct {
	DclrtnDtls  Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 DclrtnDtls,omitempty"`
	PtyCtctDtls Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PtyCtctDtls,omitempty"`
	RegnDtls    Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 RegnDtls,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Ctry"`
}

type PriorityNumeric4Choice struct {
	Nmrc  Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Nmrc"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type QuantityAndAccount38 struct {
	SttlmQty  FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SttlmQty"`
	AcctOwnr  PartyIdentification98              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AcctOwnr,omitempty"`
	SfkpgAcct SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SfkpgAcct"`
	CshAcct   CashAccountIdentification5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 CshAcct,omitempty"`
	SfkpgPlc  SafeKeepingPlace1                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SfkpgPlc,omitempty"`
}

type Rate2 struct {
	Sgn  bool    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Sgn,omitempty"`
	Rate float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Rate"`
}

type RateName1 struct {
	Issr   Max8Text  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Issr,omitempty"`
	RateNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 RateNm"`
}

type RateOrName1Choice struct {
	Rate   Rate2     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Rate"`
	RateNm RateName1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 RateNm"`
}

// May be one of FIXE, FORF, VARI
type RateType1Code string

type RateType35Choice struct {
	Cd    RateType1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type RepurchaseType21Choice struct {
	Cd    RepurchaseType8Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

// May be one of PADJ, ROLP, RATE, CALL
type RepurchaseType8Code string

type Restriction5Choice struct {
	Cd    OwnershipLegalRestrictions1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type RevaluationIndicator3Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type SafeKeepingPlace1 struct {
	SfkpgPlcFrmt SafekeepingPlaceFormat10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SfkpgPlcFrmt,omitempty"`
	LEI          LEIIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 LEI,omitempty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat10Choice struct {
	Id      SafekeepingPlaceTypeAndText8             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 TpAndId"`
	Prtry   GenericIdentification78                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Id"`
}

type SafekeepingPlaceTypeAndText8 struct {
	SfkpgPlcTp SafekeepingPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Id,omitempty"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Nm,omitempty"`
}

type SecuritiesAccount24 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Nm,omitempty"`
}

type SecuritiesFinancingModificationInstructionV06 struct {
	TxTpAndModAddtlParams TransactionTypeAndAdditionalParameters17 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 TxTpAndModAddtlParams"`
	TradDtls              SecuritiesTradeDetails5                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 TradDtls"`
	FinInstrmId           SecurityIdentification19                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 FinInstrmId"`
	QtyAndAcctDtls        QuantityAndAccount38                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 QtyAndAcctDtls"`
	SctiesFincgAddtlDtls  SecuritiesFinancingTransactionDetails27  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SctiesFincgAddtlDtls"`
	SttlmParams           SettlementDetails97                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SttlmParams,omitempty"`
	DlvrgSttlmPties       SettlementParties36                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties        SettlementParties36                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 RcvgSttlmPties,omitempty"`
	OpngSttlmAmt          AmountAndDirection49                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 OpngSttlmAmt,omitempty"`
	SplmtryData           []SupplementaryData1                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SplmtryData,omitempty"`
}

type SecuritiesFinancingTransactionDetails27 struct {
	SctiesFincgTradId     Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SctiesFincgTradId,omitempty"`
	ClsgLegId             Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 ClsgLegId,omitempty"`
	TermntnDt             TerminationDate4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 TermntnDt,omitempty"`
	RateChngDt            DateAndDateTimeChoice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 RateChngDt,omitempty"`
	EarlstCallBckDt       DateAndDateTimeChoice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 EarlstCallBckDt,omitempty"`
	ComssnClctnDt         DateAndDateTimeChoice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 ComssnClctnDt,omitempty"`
	RateTp                RateType35Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 RateTp,omitempty"`
	Rvaltn                RevaluationIndicator3Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Rvaltn,omitempty"`
	LglFrmwk              LegalFramework3Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 LglFrmwk,omitempty"`
	IntrstCmptnMtd        InterestComputationMethodFormat4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 IntrstCmptnMtd,omitempty"`
	MtrtyDtMod            bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 MtrtyDtMod,omitempty"`
	IntrstPmt             bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 IntrstPmt,omitempty"`
	VarblRateSpprt        RateName1                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 VarblRateSpprt,omitempty"`
	RpRate                Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 RpRate,omitempty"`
	StockLnMrgn           Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 StockLnMrgn,omitempty"`
	SctiesHrcut           Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SctiesHrcut,omitempty"`
	ChrgsRate             Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 ChrgsRate,omitempty"`
	PricgRate             RateOrName1Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PricgRate,omitempty"`
	Sprd                  Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Sprd,omitempty"`
	TxCallDely            Exact3NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 TxCallDely,omitempty"`
	TtlNbOfCollInstrs     Exact3NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 TtlNbOfCollInstrs,omitempty"`
	LclBrkrComssn         AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 LclBrkrComssn,omitempty"`
	DealAmt               AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 DealAmt,omitempty"`
	AcrdIntrstAmt         AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AcrdIntrstAmt,omitempty"`
	FrftAmt               AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 FrftAmt,omitempty"`
	PrmAmt                AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PrmAmt,omitempty"`
	TermntnAmtPerPcOfColl AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 TermntnAmtPerPcOfColl,omitempty"`
	TermntnTxAmt          AmountAndDirection21                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 TermntnTxAmt,omitempty"`
	ScndLegNrrtv          Max140Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 ScndLegNrrtv,omitempty"`
}

// May be one of REPU, RVPO, SECB, SECL, BSBK, SBBK
type SecuritiesFinancingTransactionType2Code string

type SecuritiesRTGS4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type SecuritiesTradeDetails5 struct {
	TradDt             DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 TradDt,omitempty"`
	OpngSttlmDt        DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 OpngSttlmDt"`
	NbOfDaysAcrd       float64               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 NbOfDaysAcrd,omitempty"`
	InstrPrcgAddtlDtls Max350Text            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 InstrPrcgAddtlDtls,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Desc,omitempty"`
}

type SettlementDetails97 struct {
	HldInd         bool                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 HldInd,omitempty"`
	Prty           PriorityNumeric4Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prty,omitempty"`
	SttlmTxCond    []SettlementTransactionCondition18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SttlmTxCond,omitempty"`
	SttlgCpcty     SettlingCapacity7Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SttlgCpcty,omitempty"`
	StmpDtyTaxBsis GenericIdentification30                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 StmpDtyTaxBsis,omitempty"`
	SctiesRTGS     SecuritiesRTGS4Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SctiesRTGS,omitempty"`
	BnfclOwnrsh    BeneficialOwnership4Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 BnfclOwnrsh,omitempty"`
	CshClrSys      CashSettlementSystem4Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 CshClrSys,omitempty"`
	TaxCpcty       TaxCapacityParty4Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 TaxCpcty,omitempty"`
	MktClntSd      MarketClientSide4Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 MktClntSd,omitempty"`
	FxStgInstr     FXStandingInstruction4Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 FxStgInstr,omitempty"`
	BlckTrad       BlockTrade4Choice                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 BlckTrad,omitempty"`
	LglRstrctns    Restriction5Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 LglRstrctns,omitempty"`
	SttlmSysMtd    SettlementSystemMethod4Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SttlmSysMtd,omitempty"`
	NetgElgblty    NettingEligibility4Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 NetgElgblty,omitempty"`
	CCPElgblty     CentralCounterPartyEligibility4Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 CCPElgblty,omitempty"`
	Trckg          Tracking4Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Trckg,omitempty"`
	AutomtcBrrwg   AutomaticBorrowing6Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AutomtcBrrwg,omitempty"`
	PrtlSttlmInd   SettlementTransactionCondition5Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PrtlSttlmInd,omitempty"`
	ElgblForColl   bool                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 ElgblForColl,omitempty"`
}

type SettlementParties36 struct {
	Dpstry PartyIdentification75            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount106 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount106 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount106 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount106 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount106 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Pty5,omitempty"`
}

// May be one of NSET, YSET
type SettlementSystemMethod1Code string

type SettlementSystemMethod4Choice struct {
	Cd    SettlementSystemMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type SettlementTransactionCondition18Choice struct {
	Cd    SettlementTransactionCondition6Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

// May be one of PART, NPAR, PARC, PARQ
type SettlementTransactionCondition5Code string

// May be one of ASGN, BUTC, CLEN, DIRT, DLWM, DRAW, EXER, FRCL, KNOC, PHYS, RESI, SHOR, SPDL, SPST, EXPI, PENS, UNEX, TRIP, NOMC, TRAN, RHYP, ADEA
type SettlementTransactionCondition6Code string

// May be one of SAGE, CUST, SPRI, RISP
type SettlingCapacity2Code string

type SettlingCapacity7Choice struct {
	Cd    SettlingCapacity2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxCapacityParty4Choice struct {
	Cd    TaxLiability1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

// May be one of PRIN, AGEN
type TaxLiability1Code string

type TerminationDate4Choice struct {
	Dt DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Dt"`
	Cd DateCode18Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Cd"`
}

type Tracking4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Prtry"`
}

type TransactionTypeAndAdditionalParameters17 struct {
	AcctOwnrTxId    Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AcctOwnrTxId"`
	AcctSvcrTxId    Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 AcctSvcrTxId,omitempty"`
	SctiesFincgTxTp SecuritiesFinancingTransactionType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 SctiesFincgTxTp"`
	Pmt             DeliveryReceiptType2Code                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 Pmt"`
	ModTp           RepurchaseType21Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 ModTp,omitempty"`
	CmonId          Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 CmonId,omitempty"`
	PoolId          Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.06 PoolId,omitempty"`
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

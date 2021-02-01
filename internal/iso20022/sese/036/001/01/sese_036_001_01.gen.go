// Code generated by main. DO NOT EDIT.

package sese_036_001_01

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

type AlternateIdentification1 struct {
	Id    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id"`
	IdSrc IdentificationSource1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 IdSrc"`
}

type AlternatePartyIdentification2 struct {
	IdTp    IdentificationType4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 IdTp"`
	Ctry    CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Ctry"`
	AltrnId Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AltrnId"`
}

type AmountAndDirection10 struct {
	Amt                 ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 CdtDbtInd,omitempty"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 OrgnlCcyAndOrdrdAmt,omitempty"`
	FXDtls              ForeignExchangeTerms11            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 FXDtls,omitempty"`
}

type AmountAndDirection4 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 CdtDbtInd,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of LAMI, NBOR, YBOR
type AutoBorrowing1Code string

type AutomaticBorrowing1Choice struct {
	Cd    AutoBorrowing1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

type BeneficialOwnership1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

type BlockTrade1Choice struct {
	Cd    BlockTrade1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of BLPA, BLCH
type BlockTrade1Code string

type CashAccountIdentification5Choice struct {
	IBAN  IBAN2007Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 IBAN"`
	Prtry Max34Text          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

type CashSettlementSystem1Choice struct {
	Cd    CashSettlementSystem2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of GROS, NETS
type CashSettlementSystem2Code string

type CentralCounterPartyEligibility1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of CODU, COPY, DUPL
type CopyDuplicate1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 DtTm"`
}

type DateCode1Choice struct {
	Cd    DateType5Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of OPEN
type DateType5Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesFincgModInstr SecuritiesFinancingModificationInstructionV01 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SctiesFincgModInstr"`
}

type DocumentIdentification11 struct {
	Id       Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id"`
	CreDtTm  DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 CreDtTm,omitempty"`
	CpyDplct CopyDuplicate1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 CpyDplct,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

type Extension2 struct {
	PlcAndNm   Max350Text         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PlcAndNm,omitempty"`
	XtnsnEnvlp ExtensionEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 XtnsnEnvlp"`
}

type ExtensionEnvelope1 struct {
	Item string `xml:",any"`
}

type FXStandingInstruction1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AmtsdVal"`
}

type ForeignExchangeTerms11 struct {
	UnitCcy  ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 UnitCcy"`
	QtdCcy   ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 QtdCcy"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 XchgRate"`
	RsltgAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 RsltgAmt"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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
	Dmst  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Dmst"`
	Prtry Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

type IdentificationType4Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat1Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

type LegalFramework1Choice struct {
	Cd    LegalFramework1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of FRAN
type LegalFramework1Code string

type MarketClientSide1Choice struct {
	Cd    MarketClientSideCode    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
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
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Adr,omitempty"`
}

type NettingEligibility1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of A144, NRST, RSTR
type OwnershipLegalRestrictions1Code string

type PartyIdentification10Choice struct {
	BICOrBEI AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 BICOrBEI"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 NmAndAdr"`
}

type PartyIdentification12Choice struct {
	BICOrBEI AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 BICOrBEI"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Ctry"`
}

type PartyIdentification13Choice struct {
	BICOrBEI AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 BICOrBEI"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PrtryId"`
}

type PartyIdentification2 struct {
	Id       PartyIdentification12Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id"`
	AltrnId  AlternatePartyIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AltrnId,omitempty"`
	PrcgDt   DateAndDateTimeChoice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PrcgDt,omitempty"`
	PrcgId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PrcgId,omitempty"`
	AddtlInf PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AddtlInf,omitempty"`
}

type PartyIdentificationAndAccount1 struct {
	Id        PartyIdentification10Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id"`
	AltrnId   AlternatePartyIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AltrnId,omitempty"`
	SfkpgAcct SecuritiesAccount13           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SfkpgAcct,omitempty"`
	PrcgDt    DateAndDateTimeChoice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PrcgDt,omitempty"`
	PrcgId    Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PrcgId,omitempty"`
	AddtlInf  PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AddtlInf,omitempty"`
}

type PartyTextInformation1 struct {
	DclrtnDtls  Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 DclrtnDtls,omitempty"`
	PtyCtctDtls Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PtyCtctDtls,omitempty"`
	RegnDtls    Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 RegnDtls,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Ctry"`
}

type PriorityNumeric1Choice struct {
	Nmrc  Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Nmrc"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

type QuantityAndAccount7 struct {
	SttlmQty  FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SttlmQty"`
	AcctOwnr  PartyIdentification13Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AcctOwnr,omitempty"`
	SfkpgAcct SecuritiesAccount13                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SfkpgAcct"`
	CshAcct   CashAccountIdentification5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 CshAcct,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat3Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SfkpgPlc,omitempty"`
}

type Rate2 struct {
	Sgn  bool    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Sgn,omitempty"`
	Rate float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Rate"`
}

type RateName1 struct {
	Issr   Max8Text  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Issr,omitempty"`
	RateNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 RateNm"`
}

type RateOrName1Choice struct {
	Rate   Rate2     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Rate"`
	RateNm RateName1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 RateNm"`
}

// May be one of FIXE, FORF, VARI
type RateType1Code string

type RateType5Choice struct {
	Cd    RateType1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

type RepurchaseType2Choice struct {
	Cd    RepurchaseType4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of CALL, RATE, ROLP
type RepurchaseType4Code string

type Restriction1Choice struct {
	Cd    OwnershipLegalRestrictions1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

type RevaluationIndicator1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat3Choice struct {
	Id      SafekeepingPlaceTypeAndText3             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id"`
}

type SafekeepingPlaceTypeAndText3 struct {
	SfkpgPlcTp SafekeepingPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id,omitempty"`
}

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Nm,omitempty"`
}

type SecuritiesFinancingModificationInstructionV01 struct {
	Id                    DocumentIdentification11                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id"`
	TxTpAndModAddtlParams TransactionTypeAndAdditionalParameters2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 TxTpAndModAddtlParams"`
	TradDtls              SecuritiesTradeDetails5                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 TradDtls"`
	FinInstrmId           SecurityIdentification11                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 FinInstrmId"`
	QtyAndAcctDtls        QuantityAndAccount7                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 QtyAndAcctDtls"`
	SctiesFincgAddtlDtls  SecuritiesFinancingTransactionDetails1  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SctiesFincgAddtlDtls"`
	SttlmParams           SettlementDetails3                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SttlmParams,omitempty"`
	DlvrgSttlmPties       SettlementParties5                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties        SettlementParties5                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 RcvgSttlmPties,omitempty"`
	OpngSttlmAmt          AmountAndDirection10                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 OpngSttlmAmt,omitempty"`
	MsgOrgtr              PartyIdentification10Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 MsgOrgtr,omitempty"`
	MsgRcpt               PartyIdentification10Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 MsgRcpt,omitempty"`
	Xtnsn                 []Extension2                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Xtnsn,omitempty"`
}

type SecuritiesFinancingTransactionDetails1 struct {
	SctiesFincgTradId     Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SctiesFincgTradId,omitempty"`
	ClsgLegId             Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 ClsgLegId,omitempty"`
	TermntnDt             TerminationDate2Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 TermntnDt,omitempty"`
	RateChngDt            DateAndDateTimeChoice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 RateChngDt,omitempty"`
	RateTp                RateType5Choice                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 RateTp,omitempty"`
	Rvaltn                RevaluationIndicator1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Rvaltn,omitempty"`
	LglFrmwk              LegalFramework1Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 LglFrmwk,omitempty"`
	IntrstCmptnMtd        InterestComputationMethodFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 IntrstCmptnMtd,omitempty"`
	MtrtyDtMod            bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 MtrtyDtMod,omitempty"`
	IntrstPmt             bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 IntrstPmt,omitempty"`
	VarblRateSpprt        RateName1                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 VarblRateSpprt,omitempty"`
	RpRate                Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 RpRate,omitempty"`
	StockLnMrgn           Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 StockLnMrgn,omitempty"`
	SctiesHrcut           Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SctiesHrcut,omitempty"`
	PricgRate             RateOrName1Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PricgRate,omitempty"`
	Sprd                  Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Sprd,omitempty"`
	TxCallDely            Exact3NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 TxCallDely,omitempty"`
	TtlNbOfCollInstrs     Exact3NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 TtlNbOfCollInstrs,omitempty"`
	DealAmt               AmountAndDirection4                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 DealAmt,omitempty"`
	AcrdIntrstAmt         AmountAndDirection4                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AcrdIntrstAmt,omitempty"`
	FrftAmt               AmountAndDirection4                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 FrftAmt,omitempty"`
	PrmAmt                AmountAndDirection4                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PrmAmt,omitempty"`
	TermntnAmtPerPcOfColl AmountAndDirection4                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 TermntnAmtPerPcOfColl,omitempty"`
	TermntnTxAmt          AmountAndDirection4                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 TermntnTxAmt,omitempty"`
	ScndLegNrrtv          Max140Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 ScndLegNrrtv,omitempty"`
}

// May be one of REPU, RVPO, SECB, SECL
type SecuritiesFinancingTransactionType1Code string

type SecuritiesRTGS1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

type SecuritiesTradeDetails5 struct {
	TradDt             DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 TradDt,omitempty"`
	OpngSttlmDt        DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 OpngSttlmDt"`
	NbOfDaysAcrd       float64               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 NbOfDaysAcrd,omitempty"`
	InstrPrcgAddtlDtls Max350Text            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 InstrPrcgAddtlDtls,omitempty"`
}

type SecurityIdentification11 struct {
	Id   SecurityIdentification11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Id"`
	Desc Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Desc,omitempty"`
}

type SecurityIdentification11Choice struct {
	ISIN   ISINIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 ISIN"`
	OthrId AlternateIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 OthrId"`
}

type SettlementDetails3 struct {
	HldInd         bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 HldInd,omitempty"`
	Prty           PriorityNumeric1Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prty,omitempty"`
	SttlmTxCond    []SettlementTransactionCondition1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SttlmTxCond,omitempty"`
	SttlgCpcty     SettlingCapacity1Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SttlgCpcty,omitempty"`
	StmpDtyTaxBsis GenericIdentification20                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 StmpDtyTaxBsis,omitempty"`
	SctiesRTGS     SecuritiesRTGS1Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SctiesRTGS,omitempty"`
	BnfclOwnrsh    BeneficialOwnership1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 BnfclOwnrsh,omitempty"`
	CshClrSys      CashSettlementSystem1Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 CshClrSys,omitempty"`
	TaxCpcty       TaxCapacityParty1Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 TaxCpcty,omitempty"`
	MktClntSd      MarketClientSide1Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 MktClntSd,omitempty"`
	FxStgInstr     FXStandingInstruction1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 FxStgInstr,omitempty"`
	BlckTrad       BlockTrade1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 BlckTrad,omitempty"`
	LglRstrctns    Restriction1Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 LglRstrctns,omitempty"`
	SttlmSysMtd    SettlementSystemMethod1Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SttlmSysMtd,omitempty"`
	NetgElgblty    NettingEligibility1Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 NetgElgblty,omitempty"`
	CCPElgblty     CentralCounterPartyEligibility1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 CCPElgblty,omitempty"`
	Trckg          Tracking1Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Trckg,omitempty"`
	AutomtcBrrwg   AutomaticBorrowing1Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AutomtcBrrwg,omitempty"`
	PrtlSttlmInd   bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PrtlSttlmInd,omitempty"`
	ElgblForColl   bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 ElgblForColl,omitempty"`
}

type SettlementParties5 struct {
	Dpstry PartyIdentification2           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Pty5,omitempty"`
}

type SettlementSystemMethod1Choice struct {
	Cd    SettlementSystemMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of NSET, YSET
type SettlementSystemMethod1Code string

type SettlementTransactionCondition1Choice struct {
	Cd    SettlementTransactionCondition2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of ASGN, BUTC, CLEN, DIRT, DLWM, DRAW, EXER, FRCL, KNOC, PHYS, RESI, SHOR, SPDL, SPST, EXPI, PENS, UNEX, TRIP, NOMC
type SettlementTransactionCondition2Code string

type SettlingCapacity1Choice struct {
	Cd    SettlingCapacity1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of CUST, SAGE, SPRI
type SettlingCapacity1Code string

type TaxCapacityParty1Choice struct {
	Cd    TaxLiability1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

// May be one of PRIN, AGEN
type TaxLiability1Code string

type TerminationDate2Choice struct {
	Dt DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Dt"`
	Cd DateCode1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Cd"`
}

type Tracking1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Prtry"`
}

type TransactionTypeAndAdditionalParameters2 struct {
	AcctOwnrTxId    Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AcctOwnrTxId"`
	AcctSvcrTxId    Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 AcctSvcrTxId,omitempty"`
	SctiesFincgTxTp SecuritiesFinancingTransactionType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 SctiesFincgTxTp"`
	Pmt             DeliveryReceiptType2Code                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Pmt"`
	ModTp           RepurchaseType2Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 ModTp,omitempty"`
	CmonId          Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 CmonId,omitempty"`
	PoolId          Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 PoolId,omitempty"`
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
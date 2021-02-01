// Code generated by main. DO NOT EDIT.

package sese_036_002_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AlternatePartyIdentification9 struct {
	IdTp    IdentificationType44Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 IdTp"`
	Ctry    CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Ctry"`
	AltrnId RestrictedFINXMax30Text    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AltrnId"`
}

type AmountAndDirection59 struct {
	Amt       RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Amt"`
	CdtDbtInd CreditDebitCode                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 CdtDbtInd,omitempty"`
}

type AmountAndDirection66 struct {
	Amt                 RestrictedFINActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Amt"`
	CdtDbtInd           CreditDebitCode                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 CdtDbtInd,omitempty"`
	OrgnlCcyAndOrdrdAmt RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 OrgnlCcyAndOrdrdAmt,omitempty"`
	FXDtls              ForeignExchangeTerms27                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 FXDtls,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of LAMI, NBOR, YBOR
type AutoBorrowing1Code string

type AutomaticBorrowing8Choice struct {
	Cd    AutoBorrowing1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type BeneficialOwnership5Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Ind"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

// May be one of BLPA, BLCH
type BlockTrade1Code string

type BlockTrade5Choice struct {
	Cd    BlockTrade1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type CashAccountIdentification6Choice struct {
	IBAN  IBAN2007Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 IBAN"`
	Prtry RestrictedFINX2Max34Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

// May be one of GROS, NETS
type CashSettlementSystem2Code string

type CashSettlementSystem5Choice struct {
	Cd    CashSettlementSystem2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type CentralCounterPartyEligibility5Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Ind"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 DtTm"`
}

type DateCode32Choice struct {
	Cd    DateType5Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

// May be one of OPEN
type DateType5Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesFincgModInstr SecuritiesFinancingModificationInstruction002V06 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SctiesFincgModInstr"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FXStandingInstruction5Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Ind"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type FinancialInstrumentQuantity15Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AmtsdVal"`
}

type ForeignExchangeTerms27 struct {
	UnitCcy  ActiveCurrencyCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 UnitCcy"`
	QtdCcy   ActiveCurrencyCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 QtdCcy"`
	XchgRate float64                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 XchgRate"`
	RsltgAmt RestrictedFINActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 RsltgAmt"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SchmeNm,omitempty"`
}

type GenericIdentification84 struct {
	Id      RestrictedFINXMax34Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SchmeNm,omitempty"`
}

type GenericIdentification85 struct {
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Tp"`
	Id RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Id,omitempty"`
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

type IdentificationSource4Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry RestrictedFINExact2Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type IdentificationType44Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat5Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be one of FRAN
type LegalFramework1Code string

type LegalFramework4Choice struct {
	Cd    LegalFramework1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type MarketClientSide5Choice struct {
	Cd    MarketClientSideCode    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

// May be one of MAKT, CLNT
type MarketClientSideCode string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress12 struct {
	Nm RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Nm"`
}

type NettingEligibility5Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Ind"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type OtherIdentification2 struct {
	Id  RestrictedFINXMax31Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Sfx,omitempty"`
	Tp  IdentificationSource4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Tp"`
}

// May be one of A144, NRST, RSTR
type OwnershipLegalRestrictions1Code string

type PartyIdentification103 struct {
	Id       PartyIdentification58Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Id"`
	LEI      LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 LEI,omitempty"`
	AltrnId  AlternatePartyIdentification9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AltrnId,omitempty"`
	PrcgDt   DateAndDateTimeChoice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 PrcgDt,omitempty"`
	PrcgId   RestrictedFINXMax16Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 PrcgId,omitempty"`
	AddtlInf PartyTextInformation3         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AddtlInf,omitempty"`
}

type PartyIdentification103Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AnyBIC"`
	PrtryId GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 PrtryId"`
}

type PartyIdentification104Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AnyBIC"`
	PrtryId  GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 PrtryId"`
	NmAndAdr NameAndAddress12        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 NmAndAdr"`
}

type PartyIdentification119 struct {
	Id  PartyIdentification103Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 LEI,omitempty"`
}

type PartyIdentification58Choice struct {
	AnyBIC   AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AnyBIC"`
	NmAndAdr NameAndAddress12 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Ctry"`
}

type PartyIdentificationAndAccount131 struct {
	Id        PartyIdentification104Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Id"`
	LEI       LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 LEI,omitempty"`
	AltrnId   AlternatePartyIdentification9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AltrnId,omitempty"`
	SfkpgAcct SecuritiesAccount27           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SfkpgAcct,omitempty"`
	PrcgDt    DateAndDateTimeChoice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 PrcgDt,omitempty"`
	PrcgId    RestrictedFINXMax16Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 PrcgId,omitempty"`
	AddtlInf  PartyTextInformation3         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AddtlInf,omitempty"`
}

type PartyTextInformation3 struct {
	DclrtnDtls  RestrictedFINXMax350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 DclrtnDtls,omitempty"`
	PtyCtctDtls RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 PtyCtctDtls,omitempty"`
	RegnDtls    RestrictedFINXMax350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 RegnDtls,omitempty"`
}

type PriorityNumeric5Choice struct {
	Nmrc  Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Nmrc"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type QuantityAndAccount61 struct {
	SttlmQty  FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SttlmQty"`
	AcctOwnr  PartyIdentification119              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AcctOwnr,omitempty"`
	SfkpgAcct SecuritiesAccount30                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SfkpgAcct"`
	CshAcct   CashAccountIdentification6Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 CshAcct,omitempty"`
	SfkpgPlc  SafeKeepingPlace2                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SfkpgPlc,omitempty"`
}

type Rate2 struct {
	Sgn  bool    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Sgn,omitempty"`
	Rate float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Rate"`
}

type RateName2 struct {
	Issr   RestrictedFINXMax8Text  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Issr,omitempty"`
	RateNm RestrictedFINXMax24Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 RateNm"`
}

type RateOrName2Choice struct {
	Rate   Rate2     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Rate"`
	RateNm RateName2 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 RateNm"`
}

// May be one of FIXE, FORF, VARI
type RateType1Code string

type RateType67Choice struct {
	Cd    RateType1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type RepurchaseType31Choice struct {
	Cd    RepurchaseType8Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

// May be one of PADJ, ROLP, RATE, CALL
type RepurchaseType8Code string

type RestrictedFINActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

type RestrictedFINActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern XX|TS
type RestrictedFINExact2Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,34}
type RestrictedFINX2Max34Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,140}
type RestrictedFINXMax140Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax16Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax24Text string

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

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,8}
type RestrictedFINXMax8Text string

type Restriction6Choice struct {
	Cd    OwnershipLegalRestrictions1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type RevaluationIndicator4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Ind"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type SafeKeepingPlace2 struct {
	SfkpgPlcFrmt SafekeepingPlaceFormat17Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SfkpgPlcFrmt,omitempty"`
	LEI          LEIIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 LEI,omitempty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat17Choice struct {
	Id      SafekeepingPlaceTypeAndText15            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 TpAndId"`
	Prtry   GenericIdentification85                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Id"`
}

type SafekeepingPlaceTypeAndText15 struct {
	SfkpgPlcTp SafekeepingPlace3Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SfkpgPlcTp"`
	Id         RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Id,omitempty"`
}

type SecuritiesAccount27 struct {
	Id RestrictedFINXMax35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Id"`
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Nm,omitempty"`
}

type SecuritiesAccount30 struct {
	Id RestrictedFINXMax35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Id"`
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Nm,omitempty"`
}

type SecuritiesFinancingModificationInstruction002V06 struct {
	TxTpAndModAddtlParams TransactionTypeAndAdditionalParameters20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 TxTpAndModAddtlParams"`
	TradDtls              SecuritiesTradeDetails12                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 TradDtls"`
	FinInstrmId           SecurityIdentification20                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 FinInstrmId"`
	QtyAndAcctDtls        QuantityAndAccount61                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 QtyAndAcctDtls"`
	SctiesFincgAddtlDtls  SecuritiesFinancingTransactionDetails32  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SctiesFincgAddtlDtls"`
	SttlmParams           SettlementDetails106                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SttlmParams,omitempty"`
	DlvrgSttlmPties       SettlementParties44                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties        SettlementParties44                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 RcvgSttlmPties,omitempty"`
	OpngSttlmAmt          AmountAndDirection66                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 OpngSttlmAmt,omitempty"`
	SplmtryData           []SupplementaryData1                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SplmtryData,omitempty"`
}

type SecuritiesFinancingTransactionDetails32 struct {
	SctiesFincgTradId     RestrictedFINXMax16Text                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SctiesFincgTradId,omitempty"`
	ClsgLegId             RestrictedFINXMax16Text                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 ClsgLegId,omitempty"`
	TermntnDt             TerminationDate5Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 TermntnDt,omitempty"`
	RateChngDt            DateAndDateTimeChoice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 RateChngDt,omitempty"`
	EarlstCallBckDt       DateAndDateTimeChoice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 EarlstCallBckDt,omitempty"`
	ComssnClctnDt         DateAndDateTimeChoice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 ComssnClctnDt,omitempty"`
	RateTp                RateType67Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 RateTp,omitempty"`
	Rvaltn                RevaluationIndicator4Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Rvaltn,omitempty"`
	LglFrmwk              LegalFramework4Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 LglFrmwk,omitempty"`
	IntrstCmptnMtd        InterestComputationMethodFormat5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 IntrstCmptnMtd,omitempty"`
	MtrtyDtMod            bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 MtrtyDtMod,omitempty"`
	IntrstPmt             bool                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 IntrstPmt,omitempty"`
	VarblRateSpprt        RateName2                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 VarblRateSpprt,omitempty"`
	RpRate                Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 RpRate,omitempty"`
	StockLnMrgn           Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 StockLnMrgn,omitempty"`
	SctiesHrcut           Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SctiesHrcut,omitempty"`
	ChrgsRate             Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 ChrgsRate,omitempty"`
	PricgRate             RateOrName2Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 PricgRate,omitempty"`
	Sprd                  Rate2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Sprd,omitempty"`
	TxCallDely            Exact3NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 TxCallDely,omitempty"`
	TtlNbOfCollInstrs     Exact3NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 TtlNbOfCollInstrs,omitempty"`
	LclBrkrComssn         AmountAndDirection59                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 LclBrkrComssn,omitempty"`
	DealAmt               AmountAndDirection59                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 DealAmt,omitempty"`
	AcrdIntrstAmt         AmountAndDirection59                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AcrdIntrstAmt,omitempty"`
	FrftAmt               AmountAndDirection59                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 FrftAmt,omitempty"`
	PrmAmt                AmountAndDirection59                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 PrmAmt,omitempty"`
	TermntnAmtPerPcOfColl AmountAndDirection59                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 TermntnAmtPerPcOfColl,omitempty"`
	TermntnTxAmt          AmountAndDirection59                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 TermntnTxAmt,omitempty"`
	ScndLegNrrtv          RestrictedFINXMax140Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 ScndLegNrrtv,omitempty"`
}

// May be one of REPU, RVPO, SECB, SECL, BSBK, SBBK
type SecuritiesFinancingTransactionType2Code string

type SecuritiesRTGS5Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Ind"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type SecuritiesTradeDetails12 struct {
	TradDt             DateAndDateTimeChoice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 TradDt,omitempty"`
	OpngSttlmDt        DateAndDateTimeChoice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 OpngSttlmDt"`
	NbOfDaysAcrd       float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 NbOfDaysAcrd,omitempty"`
	InstrPrcgAddtlDtls RestrictedFINXMax350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 InstrPrcgAddtlDtls,omitempty"`
}

type SecurityIdentification20 struct {
	ISIN   ISINOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 ISIN,omitempty"`
	OthrId []OtherIdentification2   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 OthrId,omitempty"`
	Desc   RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Desc,omitempty"`
}

type SettlementDetails106 struct {
	HldInd         bool                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 HldInd,omitempty"`
	Prty           PriorityNumeric5Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prty,omitempty"`
	SttlmTxCond    []SettlementTransactionCondition22Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SttlmTxCond,omitempty"`
	SttlgCpcty     SettlingCapacity8Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SttlgCpcty,omitempty"`
	StmpDtyTaxBsis GenericIdentification47                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 StmpDtyTaxBsis,omitempty"`
	SctiesRTGS     SecuritiesRTGS5Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SctiesRTGS,omitempty"`
	BnfclOwnrsh    BeneficialOwnership5Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 BnfclOwnrsh,omitempty"`
	CshClrSys      CashSettlementSystem5Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 CshClrSys,omitempty"`
	TaxCpcty       TaxCapacityParty5Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 TaxCpcty,omitempty"`
	MktClntSd      MarketClientSide5Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 MktClntSd,omitempty"`
	FxStgInstr     FXStandingInstruction5Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 FxStgInstr,omitempty"`
	BlckTrad       BlockTrade5Choice                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 BlckTrad,omitempty"`
	LglRstrctns    Restriction6Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 LglRstrctns,omitempty"`
	SttlmSysMtd    SettlementSystemMethod5Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SttlmSysMtd,omitempty"`
	NetgElgblty    NettingEligibility5Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 NetgElgblty,omitempty"`
	CCPElgblty     CentralCounterPartyEligibility5Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 CCPElgblty,omitempty"`
	Trckg          Tracking5Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Trckg,omitempty"`
	AutomtcBrrwg   AutomaticBorrowing8Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AutomtcBrrwg,omitempty"`
	PrtlSttlmInd   SettlementTransactionCondition5Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 PrtlSttlmInd,omitempty"`
	ElgblForColl   bool                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 ElgblForColl,omitempty"`
}

type SettlementParties44 struct {
	Dpstry PartyIdentification103           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount131 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount131 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount131 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount131 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount131 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Pty5,omitempty"`
}

// May be one of NSET, YSET
type SettlementSystemMethod1Code string

type SettlementSystemMethod5Choice struct {
	Cd    SettlementSystemMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type SettlementTransactionCondition22Choice struct {
	Cd    SettlementTransactionCondition6Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

// May be one of PART, NPAR, PARC, PARQ
type SettlementTransactionCondition5Code string

// May be one of ASGN, BUTC, CLEN, DIRT, DLWM, DRAW, EXER, FRCL, KNOC, PHYS, RESI, SHOR, SPDL, SPST, EXPI, PENS, UNEX, TRIP, NOMC, TRAN, RHYP, ADEA
type SettlementTransactionCondition6Code string

// May be one of SAGE, CUST, SPRI, RISP
type SettlingCapacity2Code string

type SettlingCapacity8Choice struct {
	Cd    SettlingCapacity2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxCapacityParty5Choice struct {
	Cd    TaxLiability1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

// May be one of PRIN, AGEN
type TaxLiability1Code string

type TerminationDate5Choice struct {
	Dt DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Dt"`
	Cd DateCode32Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Cd"`
}

type Tracking5Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Ind"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Prtry"`
}

type TransactionTypeAndAdditionalParameters20 struct {
	AcctOwnrTxId    RestrictedFINXMax16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AcctOwnrTxId"`
	AcctSvcrTxId    RestrictedFINXMax16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 AcctSvcrTxId,omitempty"`
	SctiesFincgTxTp SecuritiesFinancingTransactionType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 SctiesFincgTxTp"`
	Pmt             DeliveryReceiptType2Code                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Pmt"`
	ModTp           RepurchaseType31Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 ModTp,omitempty"`
	CmonId          RestrictedFINXMax16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 CmonId,omitempty"`
	PoolId          RestrictedFINXMax16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 PoolId,omitempty"`
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

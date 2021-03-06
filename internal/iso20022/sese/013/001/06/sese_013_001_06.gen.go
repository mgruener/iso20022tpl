// Code generated by main. DO NOT EDIT.

package sese_013_001_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Account15 struct {
	Id    Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Id,omitempty"`
	Dsgnt Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Dsgnt,omitempty"`
	Svcr  PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Svcr"`
}

type Account16 struct {
	Id    Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Id"`
	Dsgnt Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Dsgnt,omitempty"`
	Svcr  PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Svcr,omitempty"`
}

type AccountIdentification1 struct {
	Prtry SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Prtry"`
}

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

type AdditionalReference3 struct {
	Ref     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Ref"`
	RefIssr PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 RefIssr,omitempty"`
	MsgNm   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateSecurityIdentification1 struct {
	Id         Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Id"`
	DmstIdSrc  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 DmstIdSrc"`
	PrtryIdSrc Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PrtryIdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern AT[0-9]{5,5}
type AustrianBankleitzahlIdentifier string

// Must match the pattern [a-zA-Z0-9]{1,30}
type BBANIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

// Must be at least 1 items long
type BloombergIdentifier string

// Must match the pattern CP[0-9]{4,4}
type CHIPSParticipantIdentifier string

// Must match the pattern CH[0-9]{6,6}
type CHIPSUniversalIdentifier string

// Must match the pattern CA[0-9]{9,9}
type CanadianPaymentsARNIdentifier string

type CashAccount29 struct {
	Id       CashAccountIdentification1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Id"`
	AcctSvcr FinancialInstitutionIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 AcctSvcr,omitempty"`
}

type CashAccountIdentification1Choice struct {
	IBAN     IBANIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 IBAN"`
	BBAN     BBANIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 BBAN"`
	UPIC     UPICIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 UPIC"`
	DmstAcct SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 DmstAcct"`
}

type ClearingSystemMemberIdentificationChoice struct {
	USCHU  CHIPSUniversalIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 USCHU"`
	NZNCC  NewZealandNCCIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 NZNCC"`
	IENSC  IrishNSCIdentifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 IENSC"`
	GBSC   UKDomesticSortCodeIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 GBSC"`
	USCH   CHIPSParticipantIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 USCH"`
	CHBC   SwissBCIdentifier                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CHBC"`
	USFW   FedwireRoutingNumberIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 USFW"`
	PTNCC  PortugueseNCCIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PTNCC"`
	RUCB   RussianCentralBankIdentificationCodeIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 RUCB"`
	ITNCC  ItalianDomesticIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 ITNCC"`
	ATBLZ  AustrianBankleitzahlIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 ATBLZ"`
	CACPA  CanadianPaymentsARNIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CACPA"`
	CHSIC  SwissSICIdentifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CHSIC"`
	DEBLZ  GermanBankleitzahlIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 DEBLZ"`
	ESNCC  SpanishDomesticInterbankingIdentifier          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 ESNCC"`
	ZANCC  SouthAfricanNCCIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 ZANCC"`
	HKNCC  HongKongBankIdentifier                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 HKNCC"`
	AUBSBx ExtensiveBranchNetworkIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 AUBSBx"`
	AUBSBs SmallNetworkIdentifier                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 AUBSBs"`
}

// Must be at least 1 items long
type ConsolidatedTapeAssociationIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type CurrentYearType1Choice struct {
	CurYrTp       ISAType1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CurYrTp"`
	XtndedCurYrTp Extended350Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 XtndedCurYrTp"`
}

type Document struct {
	PrtflTrfConf PortfolioTransferConfirmationV06 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PrtflTrfConf"`
}

// Must be at least 1 items long
type EuroclearClearstreamIdentifier string

// Must be at least 1 items long
type Extended350Code string

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PlcAndNm"`
	Txt      Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Txt"`
}

// Must match the pattern AU[0-9]{6,6}
type ExtensiveBranchNetworkIdentifier string

// Must match the pattern FW[0-9]{9,9}
type FedwireRoutingNumberIdentifier string

type FinancialInstitutionIdentification3Choice struct {
	NmAndAdr    NameAndAddress5                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 NmAndAdr"`
	BIC         BICIdentifier                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 BIC"`
	ClrSysMmbId ClearingSystemMemberIdentificationChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 ClrSysMmbId"`
	PrtryId     SimpleIdentificationInformation          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PrtryId"`
}

type FinancialInstrument40 struct {
	Id             SecurityIdentification3Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Id"`
	Nm             Max350Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Nm,omitempty"`
	TrfTp          TransferType1Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TrfTp"`
	Qty            Quantity14Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Qty"`
	AvrgAcqstnPric ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 AvrgAcqstnPric,omitempty"`
	TrfCcy         CurrencyCode                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TrfCcy,omitempty"`
	TtlBookVal     ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TtlBookVal,omitempty"`
	TrfeeAcct      Account16                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TrfeeAcct,omitempty"`
	SubAcctDtls    SubAccount1                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 SubAcctDtls,omitempty"`
	RcvgAgtDtls    PartyIdentificationAndAccount93   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 RcvgAgtDtls,omitempty"`
	DlvrgAgtDtls   PartyIdentificationAndAccount93   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 DlvrgAgtDtls,omitempty"`
	ReqdSttlmDt    ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 ReqdSttlmDt,omitempty"`
}

type FinancialInstrumentQuantity1 struct {
	Unit float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Unit"`
}

// May be one of MALE, FEMA
type GenderCode string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Issr,omitempty"`
}

// Must match the pattern BL[0-9]{8,8}
type GermanBankleitzahlIdentifier string

// Must match the pattern HK[0-9]{3,3}
type HongKongBankIdentifier string

// Must match the pattern [a-zA-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBANIdentifier string

type ISAPortfolio2Choice struct {
	ISA   ISAYearsOfIssue5 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 ISA"`
	Prtfl Portfolio1       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Prtfl"`
}

type ISATransfer21 struct {
	MstrRef             Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 MstrRef,omitempty"`
	TrfConfId           Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TrfConfId"`
	TrfInstrRef         Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TrfInstrRef"`
	ActlTrfDt           ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 ActlTrfDt,omitempty"`
	RsdlCsh             ResidualCash1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 RsdlCsh,omitempty"`
	Prtfl               ISAPortfolio2Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Prtfl,omitempty"`
	AllOthrCsh          bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 AllOthrCsh"`
	FinInstrmAsstForTrf []FinancialInstrument40 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 FinInstrmAsstForTrf,omitempty"`
}

// May be one of MINE, MAXI, MINC
type ISAType1Code string

type ISAYearsOfIssue5 struct {
	CurYr          CurrentYearType1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CurYr,omitempty"`
	CshCmpntInd    bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CshCmpntInd"`
	PrvsYrs        PreviousYear3            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PrvsYrs,omitempty"`
	CurYrSbcptDtls SubscriptionInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CurYrSbcptDtls"`
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

type ISOYear time.Time

func (t *ISOYear) UnmarshalText(text []byte) error {
	return (*xsdGYear)(t).UnmarshalText(text)
}
func (t ISOYear) MarshalText() ([]byte, error) {
	return xsdGYear(t).MarshalText()
}

type ISOYearMonth time.Time

func (t *ISOYearMonth) UnmarshalText(text []byte) error {
	return (*xsdGYearMonth)(t).UnmarshalText(text)
}
func (t ISOYearMonth) MarshalText() ([]byte, error) {
	return xsdGYearMonth(t).MarshalText()
}

type IndividualPerson8 struct {
	Nm            Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Nm"`
	GvnNm         Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 GvnNm"`
	NmPrfx        NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 NmPrfx,omitempty"`
	NmSfx         Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 NmSfx,omitempty"`
	Gndr          GenderCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Gndr,omitempty"`
	BirthDt       ISODate         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 BirthDt,omitempty"`
	SclSctyNb     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 SclSctyNb,omitempty"`
	IndvInvstrAdr PostalAddress1  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 IndvInvstrAdr"`
}

// Must match the pattern IE[0-9]{6,6}
type IrishNSCIdentifier string

// Must match the pattern IT[0-9]{10,10}
type ItalianDomesticIdentifier string

type MarketPracticeVersion1 struct {
	Nm Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Nm"`
	Dt ISOYearMonth `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Dt,omitempty"`
	Nb Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Nb,omitempty"`
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

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Adr,omitempty"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

// Must match the pattern NZ[0-9]{6,6}
type NewZealandNCCIdentifier string

type Organisation4 struct {
	Nm            Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Nm"`
	Id            PartyIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Id,omitempty"`
	Purp          Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Purp,omitempty"`
	TaxtnCtry     CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TaxtnCtry,omitempty"`
	RegnCtry      CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 RegnCtry,omitempty"`
	RegnDt        ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 RegnDt,omitempty"`
	TaxIdNb       Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TaxIdNb,omitempty"`
	NtlRegnNb     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 NtlRegnNb,omitempty"`
	CorpInvstrAdr PostalAddress1             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CorpInvstrAdr"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 NmAndAdr"`
}

type PartyIdentification4Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PrtryId"`
}

type PartyIdentificationAndAccount93 struct {
	PtyId      PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PtyId,omitempty"`
	AcctId     AccountIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 AcctId,omitempty"`
	PlcOfSttlm PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PlcOfSttlm"`
}

type Portfolio1 struct {
	PrtflInf []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PrtflInf,omitempty"`
}

type PortfolioTransferConfirmationV06 struct {
	MsgRef           MessageIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 MsgRef"`
	PoolRef          AdditionalReference3       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PoolRef,omitempty"`
	PrvsRef          AdditionalReference3       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PrvsRef,omitempty"`
	RltdRef          AdditionalReference3       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 RltdRef,omitempty"`
	PmryIndvInvstr   IndividualPerson8          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PmryIndvInvstr,omitempty"`
	ScndryIndvInvstr IndividualPerson8          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 ScndryIndvInvstr,omitempty"`
	OthrIndvInvstr   []IndividualPerson8        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 OthrIndvInvstr,omitempty"`
	PmryCorpInvstr   Organisation4              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PmryCorpInvstr,omitempty"`
	ScndryCorpInvstr Organisation4              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 ScndryCorpInvstr,omitempty"`
	OthrCorpInvstr   []Organisation4            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 OthrCorpInvstr,omitempty"`
	TrfrAcct         Account15                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TrfrAcct"`
	NmneeAcct        Account16                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 NmneeAcct,omitempty"`
	Trfee            PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Trfee"`
	CshAcct          CashAccount29              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CshAcct,omitempty"`
	PdctTrf          []ISATransfer21            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PdctTrf"`
	MktPrctcVrsn     MarketPracticeVersion1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 MktPrctcVrsn,omitempty"`
	Xtnsn            []Extension1               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Xtnsn,omitempty"`
}

// Must match the pattern PT[0-9]{8,8}
type PortugueseNCCIdentifier string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Ctry"`
}

// Must match the pattern ALL
type PreviousAll string

type PreviousYear1Choice struct {
	AllPrvsYrs   PreviousAll `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 AllPrvsYrs"`
	SpcfcPrvsYrs []ISOYear   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 SpcfcPrvsYrs"`
}

type PreviousYear3 struct {
	PrvsYr      PreviousYear1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PrvsYr"`
	CshCmpntInd bool                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CshCmpntInd"`
}

type Quantity14Choice struct {
	Unit     Unit4   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Unit"`
	PctgRate float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 PctgRate"`
}

// Must be at least 1 items long
type RICIdentifier string

// May be one of NRCT, RCTR
type ResidualCash1Code string

// Must match the pattern RU[0-9]{9,9}
type RussianCentralBankIdentificationCodeIdentifier string

type SecurityIdentification3Choice struct {
	ISIN        ISINIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 ISIN"`
	SEDOL       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 SEDOL"`
	CUSIP       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CUSIP"`
	RIC         RICIdentifier                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 RIC"`
	TckrSymb    TickerIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TckrSymb"`
	Blmbrg      BloombergIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Blmbrg"`
	CTA         ConsolidatedTapeAssociationIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CTA"`
	QUICK       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 QUICK"`
	Wrtppr      string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Wrtppr"`
	Dtch        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Dtch"`
	Vlrn        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Vlrn"`
	SCVM        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 SCVM"`
	Belgn       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Belgn"`
	Cmon        EuroclearClearstreamIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Cmon"`
	OthrPrtryId AlternateSecurityIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 OthrPrtryId"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Id"`
}

// Must match the pattern AU[0-9]{6,6}
type SmallNetworkIdentifier string

// Must match the pattern ZA[0-9]{6,6}
type SouthAfricanNCCIdentifier string

// Must match the pattern ES[0-9]{8,9}
type SpanishDomesticInterbankingIdentifier string

type SubAccount1 struct {
	Id    AccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Id"`
	Nm    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Nm,omitempty"`
	Chrtc Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Chrtc,omitempty"`
}

type SubscriptionInformation1 struct {
	DtOfFrstSbcpt ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 DtOfFrstSbcpt"`
	EqtyCmpnt     ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 EqtyCmpnt,omitempty"`
	CshCmpnt      ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 CshCmpnt,omitempty"`
	TtlAmtYrToDt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TtlAmtYrToDt"`
}

// Must match the pattern SW[0-9]{3,5}
type SwissBCIdentifier string

// Must match the pattern SW[0-9]{6,6}
type SwissSICIdentifier string

// Must be at least 1 items long
type TickerIdentifier string

// May be one of SECU, CASH
type TransferType1Code string

// Must match the pattern SC[0-9]{6,6}
type UKDomesticSortCodeIdentifier string

// May be one of GRP1, GRP2
type UKTaxGroupUnitCode string

// Must match the pattern [0-9]{8,17}
type UPICIdentifier string

type Unit4 struct {
	TtlUnitsNb FinancialInstrumentQuantity1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 TtlUnitsNb"`
	UnitDtls   []Unit5                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 UnitDtls,omitempty"`
}

type Unit5 struct {
	UnitsNb      FinancialInstrumentQuantity1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 UnitsNb"`
	Grp1Or2Units UKTaxGroupUnitCode           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Grp1Or2Units"`
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

type xsdGYear time.Time

func (t *xsdGYear) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006")
}
func (t xsdGYear) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006")
}
func (t xsdGYear) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYear) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdGYearMonth time.Time

func (t *xsdGYearMonth) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYearMonth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

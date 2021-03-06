// Code generated by main. DO NOT EDIT.

package seev_001_001_02

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

// May be one of PRIN, SUBA
type AgentRole1Code string

type AlternateFinancialInstrumentIdentification1 struct {
	DmstIdSrc  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 DmstIdSrc"`
	PrtryIdSrc Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PrtryIdSrc"`
	Id         Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Id"`
}

type AmendInformation1 struct {
	PrvsRef      MessageIdentification `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PrvsRef"`
	RcnfrmInstrs bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RcnfrmInstrs"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CommunicationAddress4 struct {
	EmailAdr Max256Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 EmailAdr,omitempty"`
	URLAdr   Max256Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 URLAdr,omitempty"`
}

type ContactIdentification1 struct {
	Nm       Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Nm"`
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 GvnNm,omitempty"`
	Role     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Role,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PhneNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 EmailAdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyAndAmount struct {
	Value float64      `xml:",chardata"`
	Ccy   CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DateFormat2Choice struct {
	Dt   ISODateTime   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Dt"`
	DtCd DateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 DtCd"`
}

type DateFormat3Choice struct {
	Dt   ISODate       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Dt"`
	DtCd DateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 DtCd"`
}

// May be one of UKWN
type DateType1Code string

type Document struct {
	MtgNtfctn MeetingNotificationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 MtgNtfctn"`
}

type EligiblePosition2 struct {
	AcctId    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AcctId,omitempty"`
	AcctOwnr  PartyIdentification9Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AcctOwnr,omitempty"`
	HldgBal   []HoldingBalance3            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 HldgBal,omitempty"`
	RghtsHldr []PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RghtsHldr,omitempty"`
}

type EntitlementAssessment1 struct {
	SctiesBlckgDdln     DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 SctiesBlckgDdln,omitempty"`
	SctiesBlckgSTPDdln  DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 SctiesBlckgSTPDdln,omitempty"`
	SctiesBlckgMktDdln  DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 SctiesBlckgMktDdln,omitempty"`
	SctiesBlckgPrdEndDt ISODateTime       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 SctiesBlckgPrdEndDt,omitempty"`
	EntitlmntFxgDt      DateFormat3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 EntitlmntFxgDt,omitempty"`
	RegnSctiesDdln      DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RegnSctiesDdln,omitempty"`
	RegnSctiesSTPDdln   DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RegnSctiesSTPDdln,omitempty"`
	RegnSctiesMktDdln   DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RegnSctiesMktDdln,omitempty"`
	RegnPrtcptnDdln     DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RegnPrtcptnDdln,omitempty"`
	RegnPrtcptnSTPDdln  DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RegnPrtcptnSTPDdln,omitempty"`
	RegnPrtcptnMktDdln  DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RegnPrtcptnMktDdln,omitempty"`
	EntitlmntDesc       Max350Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 EntitlmntDesc,omitempty"`
	EntitlmntRatio      float64           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 EntitlmntRatio,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type Extended350Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Issr,omitempty"`
}

type GenericIdentification4 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Id"`
	IdTp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 IdTp"`
}

type GenericIdentification5 struct {
	Issr  Max8Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Issr"`
	Inf   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Inf"`
	Nrrtv Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Nrrtv,omitempty"`
}

type HoldingBalance3 struct {
	Bal      UnitOrFaceAmountChoice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Bal,omitempty"`
	BalTp    SecuritiesEntryType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 BalTp,omitempty"`
	SfkpgPlc SafekeepingPlaceFormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 SfkpgPlc,omitempty"`
	Dt       ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Dt,omitempty"`
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

type IncentivePremium2 struct {
	Desc       Max350Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Desc,omitempty"`
	Amt        PriceRateOrAmountChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Amt"`
	PerScty    float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PerScty"`
	PerVote    float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PerVote"`
	PerAttndee bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PerAttndee"`
	PmtDt      DateFormat3Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PmtDt,omitempty"`
}

type IndividualPerson14 struct {
	BirthNm   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 BirthNm"`
	GvnNm     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 GvnNm,omitempty"`
	Id        PersonIdentification2      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Id,omitempty"`
	Adr       PostalAddress1             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Adr,omitempty"`
	EmplngPty PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 EmplngPty,omitempty"`
}

type IssuerAgent1 struct {
	Id   PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Id"`
	Role AgentRole1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Role,omitempty"`
}

type LocationFormat1Choice struct {
	Adr    PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Adr"`
	LctnCd PlaceType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 LctnCd"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

// Must be at least 1 items long
type Max1025Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

// Must be at least 1 items long
type Max8Text string

type Meeting2 struct {
	DtAndTm    DateFormat2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 DtAndTm"`
	DtSts      MeetingDateStatus1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 DtSts,omitempty"`
	QrmReqrd   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 QrmReqrd"`
	Lctn       []LocationFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Lctn"`
	QrmQty     Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 QrmQty,omitempty"`
	QrmQtyPctg float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 QrmQtyPctg,omitempty"`
}

type MeetingContactPerson1 struct {
	CtctPrsn   ContactIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 CtctPrsn,omitempty"`
	EmplngPty  PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 EmplngPty,omitempty"`
	PlcOfListg MICIdentifier              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PlcOfListg,omitempty"`
}

// May be one of TNTA, CNFR, CANC, NOQO
type MeetingDateStatus1Code string

type MeetingNotice2 struct {
	MtgId                 Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 MtgId,omitempty"`
	IssrMtgId             Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 IssrMtgId,omitempty"`
	Tp                    MeetingType2Code               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Tp"`
	Clssfctn              MeetingTypeClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Clssfctn,omitempty"`
	XtndedClssfctn        Extended350Code                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 XtndedClssfctn,omitempty"`
	AnncmntDt             ISODate                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AnncmntDt,omitempty"`
	AttndncReqrd          bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AttndncReqrd"`
	AttndncConfInf        Max350Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AttndncConfInf,omitempty"`
	AttndncConfDdln       DateFormat2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AttndncConfDdln,omitempty"`
	AttndncConfSTPDdln    DateFormat2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AttndncConfSTPDdln,omitempty"`
	AttndncConfMktDdln    DateFormat2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AttndncConfMktDdln,omitempty"`
	AddtlDcmnttnURLAdr    Max256Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AddtlDcmnttnURLAdr,omitempty"`
	RsltnPrpslDdln        DateFormat2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RsltnPrpslDdln,omitempty"`
	RsltnPrpslMktDdln     DateFormat2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RsltnPrpslMktDdln,omitempty"`
	RsltnPrpslThrshld     Max350Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RsltnPrpslThrshld,omitempty"`
	RsltnPrpslThrshldPctg float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RsltnPrpslThrshldPctg,omitempty"`
	TtlNbOfSctiesOutsdng  CurrencyAndAmount              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 TtlNbOfSctiesOutsdng,omitempty"`
	TtlNbOfVtngRghts      float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 TtlNbOfVtngRghts,omitempty"`
	PrxyAppntmntNtfctnAdr PostalAddress1                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PrxyAppntmntNtfctnAdr,omitempty"`
	PrxyNotAllwd          ProxyNotAllowedCode            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PrxyNotAllwd,omitempty"`
	Prxy                  ProxyAppointmentInformation1   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Prxy,omitempty"`
	CtctPrsnDtls          []MeetingContactPerson1        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 CtctPrsnDtls,omitempty"`
	RsltPblctnDt          DateFormat3Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RsltPblctnDt,omitempty"`
}

type MeetingNotificationV02 struct {
	MtgNtfctnId       MessageIdentification1       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 MtgNtfctnId"`
	Amdmnt            AmendInformation1            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Amdmnt,omitempty"`
	NtfctnSts         NotificationStatus1          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 NtfctnSts"`
	Mtg               MeetingNotice2               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Mtg"`
	MtgDtls           []Meeting2                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 MtgDtls"`
	NtifngPty         PartyIdentification9Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 NtifngPty"`
	Issr              PartyIdentification9Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Issr"`
	IssrAgt           []IssuerAgent1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 IssrAgt,omitempty"`
	Scty              []SecurityPosition5          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Scty"`
	Rsltn             []Resolution2                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Rsltn,omitempty"`
	Vote              VoteParameters1              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Vote,omitempty"`
	EntitlmntSpcfctn  EntitlementAssessment1       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 EntitlmntSpcfctn"`
	PwrOfAttnyRqrmnts PowerOfAttorneyRequirements2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PwrOfAttnyRqrmnts,omitempty"`
}

// May be one of XMET, GMET, MIXD, SPCL
type MeetingType2Code string

// May be one of AMET, OMET, CLAS, ISSU, VRHI, CORT
type MeetingTypeClassification1Code string

type MessageIdentification struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Id"`
}

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Adr,omitempty"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type NotificationStatus1 struct {
	Sts NotificationStatus2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Sts"`
}

// May be one of ECON, EUNC
type NotificationStatus2Code string

type PartyIdentification3 struct {
	BICOrBEI AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 BICOrBEI"`
}

type PartyIdentification9Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 NmAndAdr"`
}

type PersonIdentification2 struct {
	DrvrsLicNb  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 DrvrsLicNb"`
	SclSctyNb   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 SclSctyNb"`
	AlnRegnNb   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AlnRegnNb"`
	PsptNb      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PsptNb"`
	TaxIdNb     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 TaxIdNb"`
	IdntyCardNb Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 IdntyCardNb"`
	MplyrIdNb   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 MplyrIdNb"`
	OthrId      GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 OthrId"`
	Issr        Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Issr,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

// May be one of UKWN
type PlaceType1Code string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Ctry"`
}

// May be one of NOTA, LOCA, APOS, COUN
type PowerOfAttorneyLegalisation1Code string

type PowerOfAttorneyRequirements2 struct {
	LglRqrmnt   []PowerOfAttorneyLegalisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 LglRqrmnt,omitempty"`
	OthrDcmnttn Max350Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 OthrDcmnttn,omitempty"`
}

type PriceRateOrAmountChoice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Rate"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Amt"`
}

type Proxy1 struct {
	PrxyTp      []ProxyType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PrxyTp"`
	PrssgndPrxy IndividualPerson14 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PrssgndPrxy,omitempty"`
}

type ProxyAppointmentInformation1 struct {
	RegnMtd     Max350Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RegnMtd,omitempty"`
	Ddln        DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Ddln,omitempty"`
	STPDdln     DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 STPDdln,omitempty"`
	MktDdln     DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 MktDdln,omitempty"`
	AuthrsdPrxy []Proxy1          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AuthrsdPrxy,omitempty"`
}

// May be one of NPRO
type ProxyNotAllowedCode string

// May be one of CHRM, DISC, HLDR
type ProxyType2Code string

type Resolution2 struct {
	IssrLabl           Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 IssrLabl"`
	Desc               Max1025Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Desc,omitempty"`
	Titl               Max350Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Titl,omitempty"`
	Tp                 ResolutionType1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Tp"`
	ForInfOnly         bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 ForInfOnly"`
	Sts                ResolutionStatus1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Sts"`
	SubmittdBySctyHldr bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 SubmittdBySctyHldr"`
	VoteInstrTp        []VoteInstruction2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VoteInstrTp,omitempty"`
	MgmtRcmmndtn       VoteInstruction1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 MgmtRcmmndtn,omitempty"`
	NtifngPtyRcmmndtn  VoteInstruction1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 NtifngPtyRcmmndtn,omitempty"`
}

// May be one of ACTV, WDRA
type ResolutionStatus1Code string

// May be one of EXTR, ORDI, SPCL
type ResolutionType1Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

type SafekeepingPlaceAsCodeAndPartyIdentification struct {
	PlcSfkpg SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PlcSfkpg"`
	Nrrtv    Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Nrrtv,omitempty"`
	Pty      PartyIdentification3  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Pty,omitempty"`
}

type SafekeepingPlaceFormatChoice struct {
	Id       SafekeepingPlaceAsCodeAndPartyIdentification `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Id"`
	IdAsDSS  GenericIdentification5                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 IdAsDSS"`
	IdAsCtry CountryCode                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 IdAsCtry"`
}

// May be one of BLOK, ELIG, PEND, PENR, NOMI, SETD, BORR, LOAN, SPOS, TRAD, COLI, COLO, UNBA, INBA, REGO
type SecuritiesEntryType2Code string

type SecurityIdentification3 struct {
	ISIN     ISINIdentifier                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 ISIN"`
	TckrSymb TickerIdentifier                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 TckrSymb,omitempty"`
	CUSIP    string                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 CUSIP,omitempty"`
	SEDOL    string                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 SEDOL,omitempty"`
	QUICK    string                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 QUICK,omitempty"`
	OthrId   AlternateFinancialInstrumentIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 OthrId,omitempty"`
}

type SecurityPosition5 struct {
	Id  SecurityIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Id"`
	Pos []EligiblePosition2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Pos,omitempty"`
}

// Must be at least 1 items long
type TickerIdentifier string

type UnitOrFaceAmountChoice struct {
	Unit    float64           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Unit"`
	FaceAmt CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 FaceAmt"`
}

// May be one of CFOR, CAGS, ABST, WTHH, NOAC
type VoteInstruction1Code string

// May be one of CFOR, CAGS, ABST, WTHH, WMGT, AMGT, NOAC, DISC
type VoteInstruction2Code string

type VoteMethods struct {
	VoteThrghNtwk AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VoteThrghNtwk,omitempty"`
	VoteByMail    PostalAddress1        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VoteByMail,omitempty"`
	ElctrncVote   CommunicationAddress4 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 ElctrncVote,omitempty"`
	VoteByTel     Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VoteByTel,omitempty"`
}

type VoteParameters1 struct {
	SctiesQtyReqrdToVote float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 SctiesQtyReqrdToVote,omitempty"`
	PrtlVoteAllwd        bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 PrtlVoteAllwd"`
	SpltVoteAllwd        bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 SpltVoteAllwd"`
	VoteDdln             DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VoteDdln,omitempty"`
	VoteSTPDdln          DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VoteSTPDdln,omitempty"`
	VoteMktDdln          DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VoteMktDdln,omitempty"`
	VoteMthds            VoteMethods            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VoteMthds,omitempty"`
	VtngBlltElctrncAdr   CommunicationAddress4  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VtngBlltElctrncAdr,omitempty"`
	VtngBlltReqAdr       PostalAddress1         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VtngBlltReqAdr,omitempty"`
	RvcbltyDdln          DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RvcbltyDdln,omitempty"`
	RvcbltySTPDdln       DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RvcbltySTPDdln,omitempty"`
	RvcbltyMktDdln       DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 RvcbltyMktDdln,omitempty"`
	BnfclOwnrDsclsr      bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 BnfclOwnrDsclsr"`
	VoteInstrTp          []VoteInstruction2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VoteInstrTp,omitempty"`
	IncntivPrm           IncentivePremium2      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 IncntivPrm,omitempty"`
	VoteWthPrmDdln       DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VoteWthPrmDdln,omitempty"`
	VoteWthPrmSTPDdln    DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VoteWthPrmSTPDdln,omitempty"`
	VoteWthPrmMktDdln    DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 VoteWthPrmMktDdln,omitempty"`
	AddtlVtngRqrmnts     Max350Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 AddtlVtngRqrmnts,omitempty"`
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

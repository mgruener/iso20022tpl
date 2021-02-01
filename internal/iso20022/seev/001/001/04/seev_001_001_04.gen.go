// Code generated by main. DO NOT EDIT.

package seev_001_001_04

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

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of WQPS, RSPS, AIPS
type AdditionalRight1Code string

type AdditionalRightCode1Choice struct {
	Cd    AdditionalRight1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Prtry"`
}

type AdditionalRightThreshold1Choice struct {
	AddtlRghtThrshld     Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AddtlRghtThrshld"`
	AddtlRghtThrshldPctg float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AddtlRghtThrshldPctg"`
}

type AdditionalRights1 struct {
	AddtlRght          AdditionalRightCode1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AddtlRght"`
	AddtlRghtInfURLAdr Max256Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AddtlRghtInfURLAdr,omitempty"`
	AddtlRghtDdln      DateFormat2Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AddtlRghtDdln,omitempty"`
	AddtlRghtMktDdln   DateFormat2Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AddtlRghtMktDdln,omitempty"`
	AddtlRghtThrshld   AdditionalRightThreshold1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AddtlRghtThrshld,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// May be one of PRIN, SUBA
type AgentRole1Code string

type AlternateIdentification1 struct {
	Id    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id"`
	IdSrc IdentificationSource1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 IdSrc"`
}

type AmendInformation1 struct {
	PrvsRef      MessageIdentification `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PrvsRef"`
	RcnfrmInstrs bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RcnfrmInstrs"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CommunicationAddress4 struct {
	EmailAdr Max256Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 EmailAdr,omitempty"`
	URLAdr   Max256Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 URLAdr,omitempty"`
}

type ContactIdentification1 struct {
	Nm       Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Nm"`
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 GvnNm,omitempty"`
	Role     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Role,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PhneNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 EmailAdr,omitempty"`
}

type CorporateEventNarrative2 struct {
	Dsclmr []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Dsclmr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateFormat2Choice struct {
	Dt   ISODateTime   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Dt"`
	DtCd DateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 DtCd"`
}

type DateFormat3Choice struct {
	Dt   ISODate       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Dt"`
	DtCd DateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 DtCd"`
}

// May be one of UKWN
type DateType1Code string

type Document struct {
	MtgNtfctn MeetingNotificationV04 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 MtgNtfctn"`
}

type EligiblePosition3 struct {
	AcctId    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AcctId,omitempty"`
	AcctOwnr  PartyIdentification9Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AcctOwnr,omitempty"`
	HldgBal   []HoldingBalance6            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 HldgBal,omitempty"`
	RghtsHldr []PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RghtsHldr,omitempty"`
}

type Entitlement1Choice struct {
	EntitlmntRatio float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 EntitlmntRatio"`
	EntitlmntDesc  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 EntitlmntDesc"`
}

type EntitlementAssessment2 struct {
	SctiesBlckgDdln     DateFormat2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 SctiesBlckgDdln,omitempty"`
	SctiesBlckgSTPDdln  DateFormat2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 SctiesBlckgSTPDdln,omitempty"`
	SctiesBlckgMktDdln  DateFormat2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 SctiesBlckgMktDdln,omitempty"`
	SctiesBlckgPrdEndDt ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 SctiesBlckgPrdEndDt,omitempty"`
	EntitlmntFxgDt      DateFormat3Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 EntitlmntFxgDt,omitempty"`
	RegnSctiesDdln      DateFormat2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RegnSctiesDdln,omitempty"`
	RegnSctiesSTPDdln   DateFormat2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RegnSctiesSTPDdln,omitempty"`
	RegnSctiesMktDdln   DateFormat2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RegnSctiesMktDdln,omitempty"`
	RegnPrtcptnDdln     DateFormat2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RegnPrtcptnDdln,omitempty"`
	RegnPrtcptnSTPDdln  DateFormat2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RegnPrtcptnSTPDdln,omitempty"`
	RegnPrtcptnMktDdln  DateFormat2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RegnPrtcptnMktDdln,omitempty"`
	Entitlmnt           Entitlement1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Entitlmnt,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type Extension2 struct {
	PlcAndNm   Max350Text         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PlcAndNm,omitempty"`
	XtnsnEnvlp ExtensionEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 XtnsnEnvlp"`
}

type ExtensionEnvelope1 struct {
	Item string `xml:",any"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Issr"`
}

type GenericIdentification4 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id"`
	IdTp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 IdTp"`
}

type GenericIdentification5 struct {
	Issr  Max8Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Issr"`
	Inf   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Inf"`
	Nrrtv Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Nrrtv,omitempty"`
}

type HoldingBalance6 struct {
	Bal      UnitOrFaceAmount1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Bal,omitempty"`
	BalTp    SecuritiesEntryType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 BalTp,omitempty"`
	SfkpgPlc SafekeepingPlaceFormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 SfkpgPlc,omitempty"`
	Dt       ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Dt,omitempty"`
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
	Dmst  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Dmst"`
	Prtry Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Prtry"`
}

type IncentivePremium3 struct {
	Desc  Max350Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Desc,omitempty"`
	Amt   PriceRateOrAmountChoice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Amt"`
	Tp    IncentivePremiumType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Tp"`
	PmtDt DateFormat3Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PmtDt,omitempty"`
}

type IncentivePremiumType1Choice struct {
	PerScty    float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PerScty"`
	PerVote    float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PerVote"`
	PerAttndee bool    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PerAttndee"`
}

type IndividualPerson16 struct {
	BirthNm   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 BirthNm"`
	GvnNm     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 GvnNm,omitempty"`
	Id        PersonIdentification6      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id,omitempty"`
	Adr       LongPostalAddress2Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Adr,omitempty"`
	EmplngPty PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 EmplngPty,omitempty"`
}

type IssuerAgent1 struct {
	Id   PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id"`
	Role AgentRole1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Role,omitempty"`
}

type IssuerInformation1 struct {
	Id     PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id"`
	URLAdr Max256Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 URLAdr,omitempty"`
}

type LocationFormat1Choice struct {
	Adr    PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Adr"`
	LctnCd PlaceType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 LctnCd"`
}

type LongPostalAddress2Choice struct {
	Ustrd Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Ustrd"`
	Strd  PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Strd"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

// Must be at least 1 items long
type Max1025Text string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

// Must be at least 1 items long
type Max8Text string

type Meeting3 struct {
	DtAndTm  DateFormat2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 DtAndTm"`
	DtSts    MeetingDateStatus1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 DtSts,omitempty"`
	QrmReqrd bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 QrmReqrd"`
	Lctn     []LocationFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Lctn"`
	QrmQty   QuorumQuantity1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 QrmQty,omitempty"`
}

type MeetingContactPerson1 struct {
	CtctPrsn   ContactIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 CtctPrsn,omitempty"`
	EmplngPty  PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 EmplngPty,omitempty"`
	PlcOfListg MICIdentifier              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PlcOfListg,omitempty"`
}

// May be one of TNTA, CNFR, CANC, NOQO
type MeetingDateStatus1Code string

type MeetingNotice3 struct {
	MtgId                 Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 MtgId,omitempty"`
	IssrMtgId             Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 IssrMtgId,omitempty"`
	Tp                    MeetingType2Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Tp"`
	Clssfctn              MeetingTypeClassification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Clssfctn,omitempty"`
	AnncmntDt             ISODate                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AnncmntDt,omitempty"`
	AttndncReqrd          bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AttndncReqrd"`
	AttndncConfInf        Max350Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AttndncConfInf,omitempty"`
	AttndncConfDdln       DateFormat2Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AttndncConfDdln,omitempty"`
	AttndncConfSTPDdln    DateFormat2Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AttndncConfSTPDdln,omitempty"`
	AttndncConfMktDdln    DateFormat2Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AttndncConfMktDdln,omitempty"`
	AddtlDcmnttnURLAdr    Max256Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AddtlDcmnttnURLAdr,omitempty"`
	AddtlPrcdrDtls        []AdditionalRights1              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AddtlPrcdrDtls,omitempty"`
	TtlNbOfSctiesOutsdng  ActiveCurrencyAndAmount          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 TtlNbOfSctiesOutsdng,omitempty"`
	TtlNbOfVtngRghts      float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 TtlNbOfVtngRghts,omitempty"`
	PrxyAppntmntNtfctnAdr PostalAddress1                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PrxyAppntmntNtfctnAdr,omitempty"`
	PrxyChc               Proxy1Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PrxyChc,omitempty"`
	CtctPrsnDtls          []MeetingContactPerson1          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 CtctPrsnDtls,omitempty"`
	RsltPblctnDt          DateFormat3Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RsltPblctnDt,omitempty"`
}

type MeetingNotificationV04 struct {
	Id                MessageIdentification1       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id"`
	Amdmnt            AmendInformation1            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Amdmnt,omitempty"`
	NtfctnSts         NotificationStatus1          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 NtfctnSts"`
	Mtg               MeetingNotice3               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Mtg"`
	MtgDtls           []Meeting3                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 MtgDtls"`
	NtifngPty         PartyIdentification9Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 NtifngPty"`
	Issr              IssuerInformation1           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Issr"`
	IssrAgt           []IssuerAgent1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 IssrAgt,omitempty"`
	Scty              []SecurityPosition6          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Scty"`
	Rsltn             []Resolution2                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Rsltn,omitempty"`
	Vote              VoteParameters3              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Vote,omitempty"`
	EntitlmntSpcfctn  EntitlementAssessment2       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 EntitlmntSpcfctn"`
	PwrOfAttnyRqrmnts PowerOfAttorneyRequirements2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PwrOfAttnyRqrmnts,omitempty"`
	AddtlInf          CorporateEventNarrative2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AddtlInf,omitempty"`
	Xtnsn             []Extension2                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Xtnsn,omitempty"`
}

// May be one of XMET, GMET, MIXD, SPCL
type MeetingType2Code string

type MeetingTypeClassification1Choice struct {
	Cd    MeetingTypeClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Cd"`
	Prtry GenericIdentification13        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Prtry"`
}

// May be one of AMET, OMET, CLAS, ISSU, VRHI, CORT
type MeetingTypeClassification1Code string

type MessageIdentification struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id"`
}

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Adr,omitempty"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type NotificationStatus1 struct {
	Sts NotificationStatus2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Sts"`
}

// May be one of ECON, EUNC
type NotificationStatus2Code string

type PartyIdentification3 struct {
	BICOrBEI AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 BICOrBEI"`
}

type PartyIdentification9Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 NmAndAdr"`
}

type PersonIdentification6 struct {
	Issr     Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Issr,omitempty"`
	PrsnIdTp PersonIdentificationType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PrsnIdTp"`
}

type PersonIdentificationType1Choice struct {
	PsptNb      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PsptNb"`
	TaxIdNb     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 TaxIdNb"`
	SclSctyNb   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 SclSctyNb"`
	MplyrIdNb   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 MplyrIdNb"`
	DrvrsLicNb  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 DrvrsLicNb"`
	AlnRegnNb   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AlnRegnNb"`
	IdntyCardNb Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 IdntyCardNb"`
	OthrId      GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 OthrId"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

// May be one of UKWN
type PlaceType1Code string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Ctry"`
}

// May be one of NOTA, LOCA, APOS, COUN
type PowerOfAttorneyLegalisation1Code string

type PowerOfAttorneyRequirements2 struct {
	LglRqrmnt   []PowerOfAttorneyLegalisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 LglRqrmnt,omitempty"`
	OthrDcmnttn Max350Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 OthrDcmnttn,omitempty"`
}

type PriceRateOrAmountChoice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Rate"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Amt"`
}

type Proxy1Choice struct {
	Prxy         ProxyAppointmentInformation2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Prxy"`
	PrxyNotAllwd ProxyNotAllowedCode          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PrxyNotAllwd"`
}

type Proxy3 struct {
	PrxyTp      []ProxyType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PrxyTp"`
	PrssgndPrxy IndividualPerson16 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PrssgndPrxy,omitempty"`
}

type ProxyAppointmentInformation2 struct {
	RegnMtd     Max350Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RegnMtd,omitempty"`
	Ddln        DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Ddln,omitempty"`
	STPDdln     DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 STPDdln,omitempty"`
	MktDdln     DateFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 MktDdln,omitempty"`
	AuthrsdPrxy []Proxy3          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AuthrsdPrxy,omitempty"`
}

// May be one of NPRO
type ProxyNotAllowedCode string

// May be one of CHRM, DISC, HLDR
type ProxyType2Code string

type QuorumQuantity1Choice struct {
	QrmQty     Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 QrmQty"`
	QrmQtyPctg float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 QrmQtyPctg"`
}

type Resolution2 struct {
	IssrLabl           Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 IssrLabl"`
	Desc               Max1025Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Desc,omitempty"`
	Titl               Max350Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Titl,omitempty"`
	Tp                 ResolutionType1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Tp"`
	ForInfOnly         bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 ForInfOnly"`
	Sts                ResolutionStatus1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Sts"`
	SubmittdBySctyHldr bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 SubmittdBySctyHldr"`
	VoteInstrTp        []VoteInstruction2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VoteInstrTp,omitempty"`
	MgmtRcmmndtn       VoteInstruction1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 MgmtRcmmndtn,omitempty"`
	NtifngPtyRcmmndtn  VoteInstruction1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 NtifngPtyRcmmndtn,omitempty"`
}

// May be one of ACTV, WDRA
type ResolutionStatus1Code string

// May be one of EXTR, ORDI, SPCL
type ResolutionType1Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

type SafekeepingPlaceAsCodeAndPartyIdentification struct {
	PlcSfkpg SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PlcSfkpg"`
	Nrrtv    Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Nrrtv,omitempty"`
	Pty      PartyIdentification3  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Pty,omitempty"`
}

type SafekeepingPlaceFormatChoice struct {
	Id       SafekeepingPlaceAsCodeAndPartyIdentification `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id"`
	IdAsDSS  GenericIdentification5                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 IdAsDSS"`
	IdAsCtry CountryCode                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 IdAsCtry"`
}

// May be one of BLOK, ELIG, PEND, PENR, NOMI, SETD, BORR, LOAN, SPOS, TRAD, COLI, COLO, UNBA, INBA, REGO
type SecuritiesEntryType2Code string

type SecurityIdentification11 struct {
	Id   SecurityIdentification11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id"`
	Desc Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Desc,omitempty"`
}

type SecurityIdentification11Choice struct {
	ISIN   ISINIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 ISIN"`
	OthrId AlternateIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 OthrId"`
}

type SecurityPosition6 struct {
	Id  SecurityIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Id"`
	Pos []EligiblePosition3      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Pos,omitempty"`
}

type UnitOrFaceAmount1Choice struct {
	Unit    float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Unit"`
	FaceAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 FaceAmt"`
}

// May be one of CFOR, CAGS, ABST, WTHH, NOAC
type VoteInstruction1Code string

// May be one of CFOR, CAGS, ABST, WTHH, WMGT, AMGT, NOAC, DISC
type VoteInstruction2Code string

type VoteMethods2 struct {
	VoteThrghNtwk []AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VoteThrghNtwk,omitempty"`
	VoteByMail    []PostalAddress1        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VoteByMail,omitempty"`
	ElctrncVote   []CommunicationAddress4 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 ElctrncVote,omitempty"`
	VoteByTel     []Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VoteByTel,omitempty"`
}

type VoteParameters3 struct {
	SctiesQtyReqrdToVote float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 SctiesQtyReqrdToVote,omitempty"`
	PrtlVoteAllwd        bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 PrtlVoteAllwd"`
	SpltVoteAllwd        bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 SpltVoteAllwd"`
	VoteDdln             DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VoteDdln,omitempty"`
	VoteSTPDdln          DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VoteSTPDdln,omitempty"`
	VoteMktDdln          DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VoteMktDdln,omitempty"`
	VoteMthds            VoteMethods2           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VoteMthds,omitempty"`
	VtngBlltElctrncAdr   CommunicationAddress4  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VtngBlltElctrncAdr,omitempty"`
	VtngBlltReqAdr       PostalAddress1         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VtngBlltReqAdr,omitempty"`
	RvcbltyDdln          DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RvcbltyDdln,omitempty"`
	RvcbltySTPDdln       DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RvcbltySTPDdln,omitempty"`
	RvcbltyMktDdln       DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 RvcbltyMktDdln,omitempty"`
	BnfclOwnrDsclsr      bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 BnfclOwnrDsclsr"`
	VoteInstrTp          []VoteInstruction2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VoteInstrTp,omitempty"`
	IncntivPrm           IncentivePremium3      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 IncntivPrm,omitempty"`
	VoteWthPrmDdln       DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VoteWthPrmDdln,omitempty"`
	VoteWthPrmSTPDdln    DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VoteWthPrmSTPDdln,omitempty"`
	VoteWthPrmMktDdln    DateFormat2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 VoteWthPrmMktDdln,omitempty"`
	AddtlVtngRqrmnts     Max350Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 AddtlVtngRqrmnts,omitempty"`
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

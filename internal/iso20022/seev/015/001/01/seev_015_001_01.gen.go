// Code generated by main. DO NOT EDIT.

package seev_015_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AgentCAElectionStatusAdviceV01 struct {
	Id                    DocumentIdentification8           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Id"`
	AgtCAElctnAdvcId      DocumentIdentification8           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AgtCAElctnAdvcId"`
	AgtCAElctnCxlReqId    DocumentIdentification8           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AgtCAElctnCxlReqId"`
	AgtCAElctnAmdmntReqId DocumentIdentification8           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AgtCAElctnAmdmntReqId"`
	CorpActnGnlInf        CorporateActionInformation1       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 CorpActnGnlInf"`
	ElctnAdvcSts          ElectionAdviceStatus1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 ElctnAdvcSts"`
	ElctnCxlReqSts        ElectionCancellationStatus1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 ElctnCxlReqSts"`
	ElctnAmdmntReqSts     ElectionAmendmentStatus1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 ElctnAmdmntReqSts"`
}

type AlternateSecurityIdentification3 struct {
	Id         Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Id"`
	DmstIdSrc  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 DmstIdSrc"`
	PrtryIdSrc Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 PrtryIdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CorporateActionAmendmentProcessingStatus1 struct {
	Sts      ProcessedStatus5FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Sts"`
	AddtlInf Max350Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AddtlInf,omitempty"`
}

type CorporateActionAmendmentRejectionStatus1 struct {
	Rsn      []RejectionReason8FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Rsn"`
	AddtlInf Max350Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AddtlInf,omitempty"`
}

type CorporateActionCancellationProcessingStatus1 struct {
	Sts      ProcessedStatus5FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Sts"`
	AddtlInf Max350Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AddtlInf,omitempty"`
}

type CorporateActionCancellationRejectionStatus1 struct {
	Rsn      []RejectionReason9FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Rsn"`
	AddtlInf Max350Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AddtlInf,omitempty"`
}

// May be one of GENL, DISN, REOR
type CorporateActionEventProcessingType1Code string

type CorporateActionEventProcessingType1FormatChoice struct {
	Cd    CorporateActionEventProcessingType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Cd"`
	Prtry GenericIdentification13                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Prtry"`
}

// May be one of ACTV, ATTI, BIDS, BONU, BPUT, BRUP, CAPG, CAPI, CERT, CHAN, CLSA, CONS, CONV, COOP, DECR, DETI, DFLT, DLST, DRAW, DRIP, DSCL, DTCH, DVCA, DVOP, DVSC, DVSE, EXOF, EXRI, EXTM, EXWA, INCR, INTR, LIQU, MCAL, MRGR, ODLT, PARI, PCAL, PDEF, PINK, PLAC, PPMT, PRED, PRII, PRIO, REDM, REDO, REMK, RHDI, RHTS, SHPR, SMAL, SOFF, SPLF, SPLR, SUSP, TEND, TREC, WRTH, WTRC, OTHR
type CorporateActionEventType2Code string

type CorporateActionEventType2FormatChoice struct {
	Cd    CorporateActionEventType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Cd"`
	Prtry GenericIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Prtry"`
}

type CorporateActionInformation1 struct {
	AgtId             PartyIdentification2Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AgtId"`
	IssrCorpActnId    Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 IssrCorpActnId,omitempty"`
	CorpActnPrcgId    Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 CorpActnPrcgId,omitempty"`
	EvtTp             CorporateActionEventType2FormatChoice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 EvtTp"`
	MndtryVlntryEvtTp CorporateActionMandatoryVoluntary1FormatChoice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 MndtryVlntryEvtTp"`
	EvtPrcgTp         CorporateActionEventProcessingType1FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 EvtPrcgTp,omitempty"`
	UndrlygScty       FinancialInstrumentDescription3                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 UndrlygScty"`
}

type CorporateActionInstructionProcessingStatus1 struct {
	Sts      ProcessedStatus3FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Sts"`
	AddtlInf Max350Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AddtlInf,omitempty"`
}

type CorporateActionInstructionRejectionStatus1 struct {
	Rsn      []RejectionReason18FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Rsn"`
	AddtlInf Max350Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AddtlInf,omitempty"`
}

// May be one of MAND, CHOS, VOLU
type CorporateActionMandatoryVoluntary1Code string

type CorporateActionMandatoryVoluntary1FormatChoice struct {
	Cd    CorporateActionMandatoryVoluntary1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Cd"`
	Prtry GenericIdentification13                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	AgtCAElctnStsAdvc AgentCAElectionStatusAdviceV01 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AgtCAElctnStsAdvc"`
}

type DocumentIdentification8 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 CreDtTm,omitempty"`
}

type ElectionAdviceStatus1Choice struct {
	PrcdSts  CorporateActionInstructionProcessingStatus1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 PrcdSts"`
	RjctdSts CorporateActionInstructionRejectionStatus1  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 RjctdSts"`
}

type ElectionAmendmentStatus1Choice struct {
	PrcdSts  CorporateActionAmendmentProcessingStatus1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 PrcdSts"`
	RjctdSts CorporateActionAmendmentRejectionStatus1  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 RjctdSts"`
}

type ElectionCancellationStatus1Choice struct {
	PrcdSts  CorporateActionCancellationProcessingStatus1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 PrcdSts"`
	RjctdSts CorporateActionCancellationRejectionStatus1  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 RjctdSts"`
}

type FinancialInstrumentDescription3 struct {
	SctyId     SecurityIdentification7    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 SctyId"`
	PlcOfListg MICIdentifier              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 PlcOfListg,omitempty"`
	SfkpgPlc   PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 SfkpgPlc,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Issr"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Adr,omitempty"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Ctry"`
}

// May be one of RECE, PEND, PACK
type ProcessedStatus3Code string

type ProcessedStatus3FormatChoice struct {
	Cd    ProcessedStatus3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Prtry"`
}

// May be one of RECE, PACK
type ProcessedStatus5Code string

type ProcessedStatus5FormatChoice struct {
	Cd    ProcessedStatus5Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Prtry"`
}

// May be one of FAIL, INHO, LATT
type RejectionReason18Code string

type RejectionReason18FormatChoice struct {
	Cd    RejectionReason18Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Prtry"`
}

// May be one of NAMD, LATT, ELEC, FAIL
type RejectionReason8Code string

type RejectionReason8FormatChoice struct {
	Cd    RejectionReason8Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Prtry"`
}

// May be one of NCAN, LATT, ELEC, FAIL
type RejectionReason9Code string

type RejectionReason9FormatChoice struct {
	Cd    RejectionReason9Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Prtry"`
}

type SecurityIdentification7 struct {
	ISIN   ISINIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 ISIN"`
	OthrId AlternateSecurityIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 OthrId"`
	Desc   Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Desc,omitempty"`
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

// Code generated by main. DO NOT EDIT.

package acmt_005_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification1 struct {
	Prtry SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Prtry"`
}

type AccountManagementMessageReference1 struct {
	OthrRef     AdditionalReference3       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 OthrRef,omitempty"`
	PrvsRef     AdditionalReference3       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 PrvsRef,omitempty"`
	StsReqTp    AccountManagementType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 StsReqTp"`
	AcctApplId  Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 AcctApplId,omitempty"`
	InvstmtAcct InvestmentAccount14        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 InvstmtAcct,omitempty"`
}

// May be one of ACCO, ACCM
type AccountManagementType1Code string

type AdditionalReference3 struct {
	Ref     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Ref"`
	RefIssr PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 RefIssr,omitempty"`
	MsgNm   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	ReqForAcctMgmtStsRptV02 RequestForAccountManagementStatusReportV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 ReqForAcctMgmtStsRptV02"`
}

// Must be at least 1 items long
type Extended350Code string

// May be one of MALE, FEMA
type GenderCode string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Issr,omitempty"`
}

type GenericIdentification10 struct {
	Id         Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Id"`
	IdTp       PersonIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 IdTp"`
	XtndedIdTp Extended350Code               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 XtndedIdTp"`
}

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

type IndividualPerson4 struct {
	GvnNm   Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 GvnNm"`
	MddlNm  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 MddlNm,omitempty"`
	Nm      Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Nm"`
	Gndr    GenderCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Gndr,omitempty"`
	BirthDt ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 BirthDt,omitempty"`
}

type IndividualPersonIdentificationChoice struct {
	IdNb   GenericIdentification10 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 IdNb"`
	PrsnNm IndividualPerson4       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 PrsnNm"`
}

type InvestmentAccount14 struct {
	AcctId     AccountIdentification1               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 AcctId"`
	AcctNm     Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 AcctNm,omitempty"`
	AcctDsgnt  Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 AcctDsgnt,omitempty"`
	IndvOwnrId IndividualPersonIdentificationChoice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 IndvOwnrId,omitempty"`
	OrgOwnrId  PartyIdentification2Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 OrgOwnrId,omitempty"`
	AcctSvcr   PartyIdentification2Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 AcctSvcr,omitempty"`
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Adr,omitempty"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 NmAndAdr"`
}

// May be one of PASS, CPFA, SRSA, NRIN, OTHR, DRLC, SOCS, AREG, IDCD, EMID
type PersonIdentificationType1Code string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Ctry"`
}

type RequestForAccountManagementStatusReportV02 struct {
	MsgId   MessageIdentification1             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 MsgId"`
	ReqDtls AccountManagementMessageReference1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 ReqDtls"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Id"`
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
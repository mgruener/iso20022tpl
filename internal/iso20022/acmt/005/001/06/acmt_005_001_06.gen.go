// Code generated by main. DO NOT EDIT.

package acmt_005_001_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Account23 struct {
	AcctId       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 AcctId"`
	RltdAcctDtls GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 RltdAcctDtls,omitempty"`
}

type AccountManagementMessageReference5 struct {
	LkdRef      LinkedMessage5Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 LkdRef,omitempty"`
	StsReqTp    AccountManagementType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 StsReqTp"`
	AcctApplId  Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 AcctApplId,omitempty"`
	ExstgAcctId Account23                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 ExstgAcctId,omitempty"`
	InvstmtAcct InvestmentAccount77        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 InvstmtAcct,omitempty"`
}

// May be one of ACCM, ACCO, GACC, ACST
type AccountManagementType3Code string

type AdditionalReference13 struct {
	Ref     Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Ref"`
	RefIssr PartyIdentification125Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 RefIssr,omitempty"`
	MsgNm   Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	ReqForAcctMgmtStsRpt RequestForAccountManagementStatusReportV06 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 ReqForAcctMgmtStsRpt"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be one of MALE, FEMA
type GenderCode string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Issr,omitempty"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 SchmeNm,omitempty"`
}

type GenericIdentification81 struct {
	Id   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Id"`
	IdTp OtherIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 IdTp"`
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

type IndividualPerson30 struct {
	GvnNm   Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 GvnNm,omitempty"`
	MddlNm  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 MddlNm,omitempty"`
	Nm      Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Nm"`
	Gndr    GenderCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Gndr,omitempty"`
	BirthDt ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 BirthDt,omitempty"`
}

type IndividualPersonIdentification2Choice struct {
	IdNb   GenericIdentification81 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 IdNb"`
	PrsnNm IndividualPerson30      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 PrsnNm"`
}

type InvestmentAccount77 struct {
	AcctId    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 AcctId"`
	AcctNm    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 AcctNm,omitempty"`
	AcctDsgnt Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 AcctDsgnt,omitempty"`
	OwnrId    OwnerIdentification3Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 OwnrId,omitempty"`
	AcctSvcr  PartyIdentification125Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 AcctSvcr,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LinkedMessage5Choice struct {
	PrvsRef AdditionalReference13 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 PrvsRef"`
	OthrRef AdditionalReference13 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 OthrRef"`
}

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

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Adr,omitempty"`
}

type OtherIdentification3Choice struct {
	Cd    PartyIdentificationType7Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Cd"`
	Prtry GenericIdentification47      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Prtry"`
}

type OwnerIdentification3Choice struct {
	IndvOwnrId IndividualPersonIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 IndvOwnrId"`
	OrgOwnrId  PartyIdentification139                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 OrgOwnrId"`
}

type PartyIdentification125Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 AnyBIC"`
	PrtryId  GenericIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 NmAndAdr"`
}

type PartyIdentification139 struct {
	Pty PartyIdentification125Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Pty"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 LEI,omitempty"`
}

// May be one of ATIN, IDCD, NRIN, OTHR, PASS, POCD, SOCS, SRSA, GUNL, GTIN, ITIN, CPFA, AREG, DRLC, EMID, NINV, INCL, GIIN
type PartyIdentificationType7Code string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 Ctry"`
}

type RequestForAccountManagementStatusReportV06 struct {
	MsgId   MessageIdentification1             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 MsgId"`
	ReqDtls AccountManagementMessageReference5 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.06 ReqDtls"`
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

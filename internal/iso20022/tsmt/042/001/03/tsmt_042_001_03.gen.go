// Code generated by main. DO NOT EDIT.

package tsmt_042_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of SBTW, RSTW, RSBS, ARDM, ARCS, ARES, WAIT, UPDT, SBDS, ARBA
type Action1Code string

type BICIdentification1 struct {
	BIC BICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 BIC"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

// May be one of PROP, CLSD, PMTC, ESTD, ACTV, COMP, AMRQ, RARQ, CLRQ, SCRQ, SERQ, DARQ
type BaselineStatus3Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	TxRptReq TransactionReportRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 TxRptReq"`
}

type GenericIdentification4 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 Id"`
	IdTp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 IdTp"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 CreDtTm"`
}

type PartyIdentification28 struct {
	Nm      Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 Nm"`
	PrtryId GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 PrtryId,omitempty"`
}

type PendingActivity1 struct {
	Tp   Action1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 Tp"`
	Desc Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 Desc,omitempty"`
}

type ReportSpecification4 struct {
	TxId           []Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 TxId,omitempty"`
	TxSts          []TransactionStatus4    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 TxSts,omitempty"`
	SubmitrTxRef   []Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 SubmitrTxRef,omitempty"`
	NttiesToBeRptd []BICIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 NttiesToBeRptd,omitempty"`
	Crspdt         []BICIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 Crspdt,omitempty"`
	SubmitgBk      []BICIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 SubmitgBk,omitempty"`
	OblgrBk        []BICIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 OblgrBk,omitempty"`
	Buyr           []PartyIdentification28 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 Buyr,omitempty"`
	Sellr          []PartyIdentification28 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 Sellr,omitempty"`
	BuyrCtry       []CountryCode           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 BuyrCtry,omitempty"`
	SellrCtry      []CountryCode           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 SellrCtry,omitempty"`
	CrspdtCtry     []CountryCode           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 CrspdtCtry,omitempty"`
	PdgReqForActn  []PendingActivity1      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 PdgReqForActn,omitempty"`
}

type TransactionReportRequestV03 struct {
	ReqId      MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 ReqId"`
	RptSpcfctn ReportSpecification4   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 RptSpcfctn"`
}

type TransactionStatus4 struct {
	Sts BaselineStatus3Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 Sts"`
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

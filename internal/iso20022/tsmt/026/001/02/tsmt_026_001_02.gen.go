// Code generated by main. DO NOT EDIT.

package tsmt_026_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of COMP, CLSD, ACTV
type BaselineStatus2Code string

type Document struct {
	StsChngReq StatusChangeRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.026.001.02 StsChngReq"`
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

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.026.001.02 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.026.001.02 CreDtTm"`
}

type Reason2 struct {
	Desc Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.026.001.02 Desc"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.026.001.02 Id"`
}

type StatusChangeRequestV02 struct {
	ReqId        MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.026.001.02 ReqId"`
	TxId         SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.026.001.02 TxId"`
	SubmitrTxRef SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.026.001.02 SubmitrTxRef,omitempty"`
	ReqdSts      TransactionStatus3              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.026.001.02 ReqdSts"`
	ReqRsn       Reason2                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.026.001.02 ReqRsn,omitempty"`
}

type TransactionStatus3 struct {
	Sts BaselineStatus2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.026.001.02 Sts"`
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
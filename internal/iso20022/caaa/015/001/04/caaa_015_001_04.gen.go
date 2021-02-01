// Code generated by main. DO NOT EDIT.

package caaa_015_001_04

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorRejection2 struct {
	RjctRsn  RejectReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 RjctRsn"`
	AddtlInf Max500Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 AddtlInf,omitempty"`
	MsgInErr Max100KBinary     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 MsgInErr,omitempty"`
}

type AcceptorRejectionV04 struct {
	Hdr  Header13           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Hdr"`
	Rjct AcceptorRejection2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Rjct"`
}

type Document struct {
	AccptrRjctn AcceptorRejectionV04 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 AccptrRjctn"`
}

type GenericIdentification53 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Tp,omitempty"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 ShrtNm,omitempty"`
}

type GenericIdentification76 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 ShrtNm,omitempty"`
}

type Header13 struct {
	MsgFctn    MessageFunction4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 PrtcolVrsn"`
	XchgId     Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 XchgId,omitempty"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 CreDtTm"`
	InitgPty   GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 InitgPty,omitempty"`
	RcptPty    GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 RcptPty,omitempty"`
	Tracblt    []Traceability2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must be at least 1 items long
type Max500Text string

// Must be at least 1 items long
type Max6Text string

// May be one of AUTQ, AUTP, FAUQ, FAUP, CMPV, CMPK, FCMV, FCMK, RVRA, RVRR, FRVA, FRVR, CCAQ, CCAP, CCAV, CCAK, DGNP, DGNQ, RCLQ, RCLP, RJCT, DCCQ, DCCP
type MessageFunction4Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

// May be one of UNPR, IMSG, PARS, SECU, INTP, RCPP, DPMG, VERS, MSGT
type RejectReason1Code string

type Traceability2 struct {
	RlayId      GenericIdentification76 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 RlayId"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 TracDtTmOut"`
}

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
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

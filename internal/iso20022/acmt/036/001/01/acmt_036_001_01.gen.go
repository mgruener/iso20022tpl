// Code generated by main. DO NOT EDIT.

package acmt_036_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountSwitchDetails1 struct {
	UnqRefNb      Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 UnqRefNb"`
	RtgUnqRefNb   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 RtgUnqRefNb"`
	SwtchRcvdDtTm ISODateTime                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 SwtchRcvdDtTm,omitempty"`
	SwtchDt       ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 SwtchDt,omitempty"`
	SwtchTp       SwitchType1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 SwtchTp"`
	SwtchSts      SwitchStatus1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 SwtchSts,omitempty"`
	BalTrfWndw    BalanceTransferWindow1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 BalTrfWndw,omitempty"`
	Rspn          []ResponseDetails1         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 Rspn,omitempty"`
}

type AccountSwitchTerminationSwitchV01 struct {
	MsgId         MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 AcctSwtchDtls"`
	SplmtryData   []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 SplmtryData,omitempty"`
}

// May be one of DAYH, EARL
type BalanceTransferWindow1Code string

type Document struct {
	AcctSwtchTermntnSwtch AccountSwitchTerminationSwitchV01 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 AcctSwtchTermntnSwtch"`
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

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 CreDtTm"`
}

type ResponseDetails1 struct {
	RspnCd    Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 RspnCd"`
	AddtlDtls Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 AddtlDtls,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of ACPT, BTRQ, BTRS, COMP, REDT, REDE, REJT, REQU, TMTN
type SwitchStatus1Code string

// May be one of FULL, PART
type SwitchType1Code string

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

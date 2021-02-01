// Code generated by main. DO NOT EDIT.

package reda_016_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	PtyStsAdvc PartyStatusAdviceV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 PtyStsAdvc"`
}

// Must be at least 1 items long
type ExternalStatusReason1Code string

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 SchmeNm,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

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

type MessageHeader12 struct {
	MsgId         Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 MsgId"`
	CreDtTm       ISODateTime                  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 CreDtTm,omitempty"`
	OrgnlBizInstr OriginalBusinessInstruction1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 OrgnlBizInstr,omitempty"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 Adr,omitempty"`
}

type OriginalBusinessInstruction1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 MsgId"`
	MsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 MsgNmId,omitempty"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 CreDtTm,omitempty"`
}

type PartyIdentification120Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 NmAndAdr"`
}

type PartyIdentification136 struct {
	Id  PartyIdentification120Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 LEI,omitempty"`
}

type PartyStatus2 struct {
	Sts      Status6Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 Sts"`
	StsRsn   []StatusReasonInformation10 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 StsRsn,omitempty"`
	SysPtyId SystemPartyIdentification8  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 SysPtyId,omitempty"`
}

type PartyStatusAdviceV01 struct {
	MsgHdr      MessageHeader12      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 MsgHdr,omitempty"`
	PtySts      PartyStatus2         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 PtySts"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 SplmtryData,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 Ctry"`
}

// May be one of REJT, COMP, QUED
type Status6Code string

type StatusReason6Choice struct {
	Cd    ExternalStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 Prtry"`
}

type StatusReasonInformation10 struct {
	Rsn      StatusReason6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 Rsn"`
	AddtlInf Max140Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 AddtlInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemPartyIdentification8 struct {
	Id           PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 Id"`
	RspnsblPtyId PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.016.001.01 RspnsblPtyId,omitempty"`
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

// Code generated by main. DO NOT EDIT.

package seev_048_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 DtTm"`
}

type DateCode20Choice struct {
	Cd    DateType1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Prtry"`
}

type DateFormat46Choice struct {
	Dt   DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Dt"`
	DtCd DateCode20Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 DtCd"`
}

// May be one of UKWN
type DateType1Code string

type DisclosureRequestIdentification1 struct {
	IssrDsclsrReqId      Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 IssrDsclsrReqId"`
	FinInstrmId          SecurityIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 FinInstrmId"`
	ShrhldrsDsclsrRcrdDt DateFormat46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 ShrhldrsDsclsrRcrdDt"`
}

type Document struct {
	ShrhldrIdDsclsrRspnCxlAdvc ShareholderIdentificationDisclosureResponseCancellationAdviceV01 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 ShrhldrIdDsclsrRspnCxlAdvc"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Prtry"`
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

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Tp"`
}

type PartyIdentification195Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 PrtryId"`
	LEI     LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 LEI"`
}

type PartyIdentification215 struct {
	Nm Max350Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Nm"`
	Id PartyIdentification195Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Id"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Desc,omitempty"`
}

type ShareholderIdentificationDisclosureResponseCancellationAdviceV01 struct {
	DsclsrRspnId     Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 DsclsrRspnId"`
	IssrDsclsrReqRef DisclosureRequestIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 IssrDsclsrReqRef"`
	RspndgIntrmy     PartyIdentification215           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 RspndgIntrmy"`
	SplmtryData      []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 SplmtryData,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.048.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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

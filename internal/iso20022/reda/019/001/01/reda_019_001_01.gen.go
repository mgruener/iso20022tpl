// Code generated by main. DO NOT EDIT.

package reda_019_001_01

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

type DatePeriod2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 ToDt"`
}

type DatePeriodSearch1Choice struct {
	FrDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 FrDt"`
	ToDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 ToDt"`
	FrToDt DatePeriod2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 FrToDt"`
	EQDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 EQDt"`
	NEQDt  ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 NEQDt"`
}

type Document struct {
	SctiesAcctQry SecuritiesAccountQueryV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 SctiesAcctQry"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalSystemPartyType1Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 SchmeNm,omitempty"`
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

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type MessageHeader2 struct {
	MsgId   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 MsgId"`
	CreDtTm ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 CreDtTm,omitempty"`
	ReqTp   RequestType2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 ReqTp,omitempty"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Adr,omitempty"`
}

type PartyIdentification120Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 NmAndAdr"`
}

type PartyIdentification136 struct {
	Id  PartyIdentification120Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 LEI,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Ctry"`
}

// May be one of RT01, RT02, RT03, RT04, RT05
type RequestType1Code string

type RequestType2Choice struct {
	PmtCtrl RequestType1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 PmtCtrl"`
	Enqry   RequestType2Code       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Enqry"`
	Prtry   GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Prtry"`
}

// May be one of RT11, RT12, RT13, RT14, RT15
type RequestType2Code string

type SecuritiesAccountQueryV01 struct {
	MsgHdr      MessageHeader2                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 MsgHdr,omitempty"`
	SchCrit     SecuritiesAccountSearchCriteria2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 SchCrit"`
	RtrCrit     SecuritiesAccountReturnCriteria1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 RtrCrit,omitempty"`
	SplmtryData []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 SplmtryData,omitempty"`
}

type SecuritiesAccountReturnCriteria1 struct {
	AcctId       bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 AcctId,omitempty"`
	PtyId        bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 PtyId,omitempty"`
	PtyTp        bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 PtyTp,omitempty"`
	AcctSvcr     bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 AcctSvcr,omitempty"`
	AcctTp       bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 AcctTp,omitempty"`
	OpngDt       bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 OpngDt,omitempty"`
	ClsgDt       bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 ClsgDt,omitempty"`
	EndInvstrFlg bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 EndInvstrFlg,omitempty"`
	PricgSchme   bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 PricgSchme,omitempty"`
}

type SecuritiesAccountSearchCriteria2 struct {
	AcctId       Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 AcctId,omitempty"`
	AcctSvcr     PartyIdentification136             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 AcctSvcr,omitempty"`
	AcctOwnr     SystemPartyIdentification8         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 AcctOwnr,omitempty"`
	PtyTp        SystemPartyType1Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 PtyTp,omitempty"`
	OpngDt       DatePeriodSearch1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 OpngDt,omitempty"`
	ClsgDt       DatePeriodSearch1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 ClsgDt,omitempty"`
	AcctTp       SystemSecuritiesAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 AcctTp,omitempty"`
	EndInvstrFlg Exact4AlphaNumericText             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 EndInvstrFlg,omitempty"`
	PricgSchme   Exact4AlphaNumericText             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 PricgSchme,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemPartyIdentification8 struct {
	Id           PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Id"`
	RspnsblPtyId PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 RspnsblPtyId,omitempty"`
}

type SystemPartyType1Choice struct {
	Cd    ExternalSystemPartyType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Prtry"`
}

type SystemSecuritiesAccountType1Choice struct {
	Cd    SystemSecuritiesAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Cd"`
	Prtry GenericIdentification30          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.019.001.01 Prtry"`
}

// May be one of CSDP, CSDM, ICSA, TOFF, CSDO, ISSA
type SystemSecuritiesAccountType1Code string

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

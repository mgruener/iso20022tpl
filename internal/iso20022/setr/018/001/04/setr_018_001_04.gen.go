// Code generated by main. DO NOT EDIT.

package setr_018_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AdditionalReference8 struct {
	Ref     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Ref"`
	RefIssr PartyIdentification113 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 RefIssr,omitempty"`
	MsgNm   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateSecurityIdentification7 struct {
	Id    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Id"`
	IdSrc IdentificationSource1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 IdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern (BBG)[BCDFGHJKLMNPQRSTVWXYZ\d]{8}\d
type Bloomberg2Identifier string

// Must be at least 1 items long
type ConsolidatedTapeAssociationIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateFormat42Choice struct {
	YrMnth    ISOYearMonth `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 YrMnth"`
	YrMnthDay ISODate      `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 YrMnthDay"`
}

// May be one of DIST, ACCU
type DistributionPolicy1Code string

type Document struct {
	ReqForOrdrStsRpt RequestForOrderStatusReportV04 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 ReqForOrdrStsRpt"`
}

// Must be at least 1 items long
type EuroclearClearstreamIdentifier string

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 PlcAndNm"`
	Txt      Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Txt"`
}

type FinancialInstrument57 struct {
	Id          SecurityIdentification25Choice `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Id"`
	Nm          Max350Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Nm,omitempty"`
	ShrtNm      Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 ShrtNm,omitempty"`
	SplmtryId   Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 SplmtryId,omitempty"`
	ClssTp      Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 ClssTp,omitempty"`
	SctiesForm  FormOfSecurity1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 SctiesForm,omitempty"`
	DstrbtnPlcy DistributionPolicy1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 DstrbtnPlcy,omitempty"`
	PdctGrp     Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 PdctGrp,omitempty"`
	SrsId       Series1                        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 SrsId,omitempty"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Issr,omitempty"`
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

type ISOYearMonth time.Time

func (t *ISOYearMonth) UnmarshalText(text []byte) error {
	return (*xsdGYearMonth)(t).UnmarshalText(text)
}
func (t ISOYearMonth) MarshalText() ([]byte, error) {
	return xsdGYearMonth(t).MarshalText()
}

type IdentificationSource1Choice struct {
	Dmst  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Dmst"`
	Prtry Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Prtry"`
}

type InvestmentAccount58 struct {
	AcctId           Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 AcctId"`
	AcctNm           Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 AcctNm,omitempty"`
	AcctDsgnt        Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 AcctDsgnt,omitempty"`
	OwnrId           []PartyIdentification113        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 OwnrId,omitempty"`
	AcctSvcr         PartyIdentification113          `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 AcctSvcr,omitempty"`
	OrdrOrgtrElgblty OrderOriginatorEligibility1Code `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 OrdrOrgtrElgblty,omitempty"`
	SubAcctDtls      SubAccount6                     `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 SubAcctDtls,omitempty"`
}

type InvestmentFundOrder8 struct {
	MstrRef         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 MstrRef,omitempty"`
	OrdrRef         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 OrdrRef"`
	ClntRef         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 ClntRef,omitempty"`
	CxlRef          Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 CxlRef,omitempty"`
	DealRef         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 DealRef,omitempty"`
	InvstmtAcctDtls InvestmentAccount58   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 InvstmtAcctDtls,omitempty"`
	FinInstrmDtls   FinancialInstrument57 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 FinInstrmDtls,omitempty"`
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

type MessageAndBusinessReference10 struct {
	Ref     References62Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Ref,omitempty"`
	RltdRef AdditionalReference8   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 RltdRef,omitempty"`
	OrdrRef []InvestmentFundOrder8 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 OrdrRef,omitempty"`
}

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Adr,omitempty"`
}

// May be one of ELIG, RETL, PROF
type OrderOriginatorEligibility1Code string

type PartyIdentification113 struct {
	Pty PartyIdentification90Choice `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Pty"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 LEI,omitempty"`
}

type PartyIdentification90Choice struct {
	AnyBIC   AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 AnyBIC"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Ctry"`
}

// Must be at least 1 items long
type RICIdentifier string

type References62Choice struct {
	PrvsRef []AdditionalReference8 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 PrvsRef"`
	OthrRef []AdditionalReference8 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 OthrRef"`
}

type RequestForOrderStatusReportV04 struct {
	MsgId   MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 MsgId"`
	ReqDtls []MessageAndBusinessReference10 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 ReqDtls"`
	Xtnsn   []Extension1                    `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Xtnsn,omitempty"`
}

type SecurityIdentification25Choice struct {
	ISIN        ISINOct2015Identifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 ISIN"`
	SEDOL       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 SEDOL"`
	CUSIP       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 CUSIP"`
	RIC         RICIdentifier                         `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 RIC"`
	TckrSymb    TickerIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 TckrSymb"`
	Blmbrg      Bloomberg2Identifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Blmbrg"`
	CTA         ConsolidatedTapeAssociationIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 CTA"`
	QUICK       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 QUICK"`
	Wrtppr      string                                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Wrtppr"`
	Dtch        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Dtch"`
	Vlrn        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Vlrn"`
	SCVM        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 SCVM"`
	Belgn       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Belgn"`
	Cmon        EuroclearClearstreamIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Cmon"`
	OthrPrtryId AlternateSecurityIdentification7      `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 OthrPrtryId"`
}

type Series1 struct {
	SrsDt DateFormat42Choice `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 SrsDt,omitempty"`
	SrsNm Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 SrsNm,omitempty"`
}

type SubAccount6 struct {
	Id        Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Id"`
	Nm        Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Nm,omitempty"`
	Chrtc     Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Chrtc,omitempty"`
	AcctDsgnt Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 AcctDsgnt,omitempty"`
}

// Must be at least 1 items long
type TickerIdentifier string

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

type xsdGYearMonth time.Time

func (t *xsdGYearMonth) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYearMonth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
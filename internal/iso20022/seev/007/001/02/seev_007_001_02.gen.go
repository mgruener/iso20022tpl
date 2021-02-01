// Code generated by main. DO NOT EDIT.

package seev_007_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateFinancialInstrumentIdentification1 struct {
	DmstIdSrc  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 DmstIdSrc"`
	PrtryIdSrc Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 PrtryIdSrc"`
	Id         Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Id"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DetailedInstructionStatus2 struct {
	InstrId      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 InstrId"`
	AcctId       Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 AcctId,omitempty"`
	AcctOwnr     PartyIdentification9Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 AcctOwnr,omitempty"`
	SubAcctId    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 SubAcctId,omitempty"`
	RghtsHldr    []PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 RghtsHldr,omitempty"`
	StgInstr     bool                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 StgInstr"`
	VotePerRsltn []Vote4                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 VotePerRsltn"`
}

type Document struct {
	MtgVoteExctnConf MeetingVoteExecutionConfirmationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 MtgVoteExctnConf"`
}

// Must be at least 1 items long
type Extended350Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Issr,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type MeetingReference3 struct {
	MtgId          Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 MtgId,omitempty"`
	IssrMtgId      Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 IssrMtgId,omitempty"`
	MtgDtAndTm     ISODateTime                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 MtgDtAndTm"`
	Tp             MeetingType2Code               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Tp"`
	Clssfctn       MeetingTypeClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Clssfctn,omitempty"`
	XtndedClssfctn Extended350Code                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 XtndedClssfctn,omitempty"`
	Lctn           []PostalAddress1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Lctn,omitempty"`
}

// May be one of XMET, GMET, MIXD, SPCL
type MeetingType2Code string

// May be one of AMET, OMET, CLAS, ISSU, VRHI, CORT
type MeetingTypeClassification1Code string

type MeetingVoteExecutionConfirmationV02 struct {
	VoteExctnConfId MessageIdentification1       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 VoteExctnConfId"`
	RltdRef         MessageIdentification        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 RltdRef"`
	MtgRef          MeetingReference3            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 MtgRef"`
	RptgPty         PartyIdentification9Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 RptgPty"`
	SctyId          SecurityIdentification3      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 SctyId"`
	VoteInstr       []DetailedInstructionStatus2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 VoteInstr"`
}

type MessageIdentification struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Id"`
}

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Adr,omitempty"`
}

type PartyIdentification9Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Ctry"`
}

type SecurityIdentification3 struct {
	ISIN     ISINIdentifier                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 ISIN"`
	TckrSymb TickerIdentifier                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 TckrSymb,omitempty"`
	CUSIP    string                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 CUSIP,omitempty"`
	SEDOL    string                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 SEDOL,omitempty"`
	QUICK    string                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 QUICK,omitempty"`
	OthrId   AlternateFinancialInstrumentIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 OthrId,omitempty"`
}

// Must be at least 1 items long
type TickerIdentifier string

type Vote4 struct {
	IssrLabl  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 IssrLabl"`
	For       float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 For,omitempty"`
	Agnst     float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Agnst,omitempty"`
	Abstn     float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Abstn,omitempty"`
	Wthhld    float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Wthhld,omitempty"`
	WthMgmt   float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 WthMgmt,omitempty"`
	AgnstMgmt float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 AgnstMgmt,omitempty"`
	Dscrtnry  float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Dscrtnry,omitempty"`
	NoActn    float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 NoActn,omitempty"`
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
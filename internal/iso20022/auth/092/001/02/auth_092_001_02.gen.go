// Code generated by main. DO NOT EDIT.

package auth_092_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type CompetentAuthority1 struct {
	Id        Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 Id"`
	OnbrdgSts bool       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 OnbrdgSts"`
}

type CounterpartyData78 struct {
	RptSubmitgNtty OrganisationIdentification10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 RptSubmitgNtty"`
	RptgCtrPty     OrganisationIdentification10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 RptgCtrPty"`
}

type DerivativesStatistics3 struct {
	TtlSubmittd     StatisticsPerActionType1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 TtlSubmittd"`
	TtlRjctd        StatisticsPerActionType1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 TtlRjctd"`
	TtlCrrctdRjctns StatisticsPerActionType1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 TtlCrrctdRjctns"`
	TopRjctnRsns    TopReasonsForRejections2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 TopRjctnRsns"`
}

type DerivativesTradeRejectionStatisticalReportV02 struct {
	SttstcsPerCtrPty StatisticsPerCounterparty3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 SttstcsPerCtrPty"`
	SplmtryData      []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 SplmtryData,omitempty"`
}

type DetailedStatisticsPerCounterparty7 struct {
	RptgPrd       Period2               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 RptgPrd"`
	CtrPtyId      CounterpartyData78    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 CtrPtyId"`
	RjctnSttstcs  RejectionStatistics3  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 RjctnSttstcs"`
	CmptntAuthrty []CompetentAuthority1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 CmptntAuthrty,omitempty"`
}

type Document struct {
	DerivsTradRjctnSttstclRpt DerivativesTradeRejectionStatisticalReportV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 DerivsTradRjctnSttstclRpt"`
}

type GenericIdentification175 struct {
	Id      Max72Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 Issr,omitempty"`
}

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max500Text string

// Must be at least 1 items long
type Max72Text string

type OrganisationIdentification10Choice struct {
	LEI    LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 LEI"`
	Othr   OrganisationIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 Othr"`
	AnyBIC AnyBICDec2014Identifier      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 AnyBIC"`
}

type OrganisationIdentification36 struct {
	Id   GenericIdentification175 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 Id"`
	Nm   Max105Text               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 Nm,omitempty"`
	Dmcl Max500Text               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 Dmcl,omitempty"`
}

type Period2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 ToDt"`
}

type RejectionStatistics3 struct {
	TtlNbOfTechRjctns float64                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 TtlNbOfTechRjctns"`
	DerivsSttstcs     DerivativesStatistics3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 DerivsSttstcs"`
}

// May be one of NOTX
type ReportPeriodActivity1Code string

type StatisticsPerActionType1 struct {
	All float64 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 All"`
	New float64 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 New"`
	Mod float64 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 Mod"`
}

type StatisticsPerCounterparty3Choice struct {
	DataSetActn ReportPeriodActivity1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 DataSetActn"`
	Rpt         []DetailedStatisticsPerCounterparty7 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 Rpt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TopReasonsForRejections2 struct {
	All []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 All,omitempty"`
	New []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 New,omitempty"`
	Mod []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.02 Mod,omitempty"`
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
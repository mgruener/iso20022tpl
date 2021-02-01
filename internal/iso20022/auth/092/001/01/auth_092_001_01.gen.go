// Code generated by main. DO NOT EDIT.

package auth_092_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type CompetentAuthority1 struct {
	Id        Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 Id"`
	OnbrdgSts bool       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 OnbrdgSts"`
}

type CounterpartyData36 struct {
	RptSubmitgNtty LEIIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 RptSubmitgNtty"`
	RptgCtrPty     LEIIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 RptgCtrPty"`
}

type DerivativesStatistics3 struct {
	TtlSubmittd     StatisticsPerActionType1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 TtlSubmittd"`
	TtlRjctd        StatisticsPerActionType1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 TtlRjctd"`
	TtlCrrctdRjctns StatisticsPerActionType1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 TtlCrrctdRjctns"`
	TopRjctnRsns    TopReasonsForRejections2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 TopRjctnRsns"`
}

type DerivativesTradeRejectionStatisticalReportV01 struct {
	SttstcsPerCtrPty StatisticsPerCounterparty2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 SttstcsPerCtrPty"`
	SplmtryData      []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 SplmtryData,omitempty"`
}

type DetailedStatisticsPerCounterparty6 struct {
	RptgPrd       Period2               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 RptgPrd"`
	CtrPtyId      CounterpartyData36    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 CtrPtyId"`
	RjctnSttstcs  RejectionStatistics3  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 RjctnSttstcs"`
	CmptntAuthrty []CompetentAuthority1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 CmptntAuthrty,omitempty"`
}

type Document struct {
	DerivsTradRjctnSttstclRpt DerivativesTradeRejectionStatisticalReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 DerivsTradRjctnSttstclRpt"`
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
type Max350Text string

// Must be at least 1 items long
type Max35Text string

type Period2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 ToDt"`
}

type RejectionStatistics3 struct {
	TtlNbOfTechRjctns float64                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 TtlNbOfTechRjctns"`
	DerivsSttstcs     DerivativesStatistics3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 DerivsSttstcs"`
}

// May be one of NOTX
type ReportPeriodActivity1Code string

type StatisticsPerActionType1 struct {
	All float64 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 All"`
	New float64 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 New"`
	Mod float64 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 Mod"`
}

type StatisticsPerCounterparty2Choice struct {
	DataSetActn ReportPeriodActivity1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 DataSetActn"`
	Rpt         []DetailedStatisticsPerCounterparty6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 Rpt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TopReasonsForRejections2 struct {
	All []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 All,omitempty"`
	New []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 New,omitempty"`
	Mod []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.092.001.01 Mod,omitempty"`
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
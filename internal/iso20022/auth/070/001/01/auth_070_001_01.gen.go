// Code generated by main. DO NOT EDIT.

package auth_070_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type CollateralMarginCorrection3 struct {
	TechRcrdId     Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 TechRcrdId,omitempty"`
	RptgDtTm       ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 RptgDtTm"`
	EvtDt          ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 EvtDt"`
	CtrPty         Counterparty30              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 CtrPty"`
	CollPrtflId    Max52Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 CollPrtflId"`
	PstdMrgnOrColl PostedMarginOrCollateral3   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 PstdMrgnOrColl,omitempty"`
	RcvdMrgnOrColl ReceivedMarginOrCollateral3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 RcvdMrgnOrColl,omitempty"`
	SplmtryData    []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 SplmtryData,omitempty"`
}

type CollateralMarginError2 struct {
	TechRcrdId  Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 TechRcrdId,omitempty"`
	RptgDtTm    ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 RptgDtTm"`
	CtrPty      Counterparty30       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 CtrPty"`
	CollPrtflId Max52Text            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 CollPrtflId"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 SplmtryData,omitempty"`
}

type CollateralMarginMarginUpdate2 struct {
	TechRcrdId     Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 TechRcrdId,omitempty"`
	RptgDtTm       ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 RptgDtTm"`
	EvtDt          ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 EvtDt"`
	CtrPty         Counterparty30              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 CtrPty,omitempty"`
	CollPrtflId    Max52Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 CollPrtflId"`
	PstdMrgnOrColl PostedMarginOrCollateral3   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 PstdMrgnOrColl,omitempty"`
	RcvdMrgnOrColl ReceivedMarginOrCollateral3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 RcvdMrgnOrColl,omitempty"`
	SplmtryData    []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 SplmtryData,omitempty"`
}

type Counterparty30 struct {
	RptgCtrPty        OrganisationIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 RptgCtrPty"`
	OthrCtrPty        OrganisationIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 OthrCtrPty"`
	NttyRspnsblForRpt OrganisationIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 NttyRspnsblForRpt,omitempty"`
	RptSubmitgNtty    OrganisationIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 RptSubmitgNtty,omitempty"`
}

type Document struct {
	SctiesFincgRptgTxMrgnDataRpt SecuritiesFinancingReportingTransactionMarginDataReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 SctiesFincgRptgTxMrgnDataRpt"`
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
type Max140Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max50Text string

// Must be at least 1 items long
type Max52Text string

type OrganisationIdentification9Choice struct {
	LEI    LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 LEI"`
	ClntId Max50Text               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 ClntId"`
	AnyBIC AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 AnyBIC"`
}

type PostedMarginOrCollateral3 struct {
	InitlMrgnPstd ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 InitlMrgnPstd,omitempty"`
	VartnMrgnPstd ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 VartnMrgnPstd,omitempty"`
	XcssCollPstd  ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 XcssCollPstd,omitempty"`
}

type ReceivedMarginOrCollateral3 struct {
	InitlMrgnRcvd ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 InitlMrgnRcvd,omitempty"`
	VartnMrgnRcvd ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 VartnMrgnRcvd,omitempty"`
	XcssCollRcvd  ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 XcssCollRcvd,omitempty"`
}

// May be one of NOTX
type ReportPeriodActivity1Code string

type SecuritiesFinancingReportingTransactionMarginDataReportV01 struct {
	TradData    TradeData6Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 TradData"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 SplmtryData,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeData6Choice struct {
	DataSetActn ReportPeriodActivity1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 DataSetActn"`
	Rpt         []TradeReport6Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 Rpt"`
}

type TradeReport6Choice struct {
	New     CollateralMarginCorrection3   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 New"`
	Err     CollateralMarginError2        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 Err"`
	Crrctn  CollateralMarginCorrection3   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 Crrctn"`
	TradUpd CollateralMarginMarginUpdate2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.070.001.01 TradUpd"`
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

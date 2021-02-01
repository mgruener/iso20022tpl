// Code generated by main. DO NOT EDIT.

package auth_013_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of BILA, BROK
type BrokeredDeal1Code string

type CounterpartyIdentification3Choice struct {
	LEI         LEIIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 LEI"`
	SctrAndLctn SectorAndLocation1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 SctrAndLctn"`
	NmAndLctn   NameAndLocation1   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 NmAndLctn"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 DtTm"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 ToDtTm"`
}

type Document struct {
	MnyMktUscrdMktSttstclRpt MoneyMarketUnsecuredMarketStatisticalReportV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 MnyMktUscrdMktSttstclRpt"`
}

// May be one of CEOD, COPR, OTHR, ABCP, FRNT, CACM, DPST
type FinancialInstrumentProductType1Code string

type FloatingRateNote2 struct {
	RefRateIndx ISINOct2015Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 RefRateIndx"`
	BsisPtSprd  float64               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 BsisPtSprd"`
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

// May be one of FIXE, VARI
type InterestRateType1Code string

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max70Text string

type MoneyMarketReportHeader1 struct {
	RptgAgt LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 RptgAgt"`
	RefPrd  DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 RefPrd"`
}

// May be one of BORR, LEND
type MoneyMarketTransactionType1Code string

type MoneyMarketUnsecuredMarketStatisticalReportV02 struct {
	RptHdr      MoneyMarketReportHeader1     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 RptHdr"`
	UscrdMktRpt UnsecuredMarketReport4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 UscrdMktRpt"`
	SplmtryData []SupplementaryData1         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 SplmtryData,omitempty"`
}

type NameAndLocation1 struct {
	Nm   Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 Nm"`
	Lctn CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 Lctn"`
}

// May be one of NONO, NOVA
type NovationStatus1Code string

type Option12 struct {
	Tp      OptionType1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 Tp"`
	DtOrPrd OptionDateOrPeriod1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 DtOrPrd"`
}

type OptionDateOrPeriod1Choice struct {
	EarlstExrcDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 EarlstExrcDt"`
	NtcePrd      float64 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 NtcePrd"`
}

// May be one of CALL, PUTO
type OptionType1Code string

// May be one of NOTX, NORA
type ReportPeriodActivity3Code string

type SectorAndLocation1 struct {
	Sctr string      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 Sctr"`
	Lctn CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 Lctn"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of AMND, CANC, CORR, NEWT
type TransactionOperationType1Code string

type UnsecuredMarketReport4Choice struct {
	DataSetActn ReportPeriodActivity3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 DataSetActn"`
	Tx          []UnsecuredMarketTransaction4 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 Tx"`
}

type UnsecuredMarketTransaction4 struct {
	RptdTxSts       TransactionOperationType1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 RptdTxSts"`
	NvtnSts         NovationStatus1Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 NvtnSts,omitempty"`
	BrnchId         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 BrnchId,omitempty"`
	UnqTxIdr        Max105Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 UnqTxIdr,omitempty"`
	PrtryTxId       Max105Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 PrtryTxId"`
	RltdPrtryTxId   Max105Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 RltdPrtryTxId,omitempty"`
	CtrPtyPrtryTxId Max105Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 CtrPtyPrtryTxId,omitempty"`
	CtrPtyId        CounterpartyIdentification3Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 CtrPtyId"`
	TradDt          DateAndDateTimeChoice               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 TradDt"`
	SttlmDt         ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 SttlmDt"`
	MtrtyDt         ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 MtrtyDt"`
	TxTp            MoneyMarketTransactionType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 TxTp"`
	InstrmTp        FinancialInstrumentProductType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 InstrmTp"`
	TxNmnlAmt       ActiveCurrencyAndAmount             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 TxNmnlAmt"`
	DealPric        float64                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 DealPric"`
	RateTp          InterestRateType1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 RateTp"`
	DealRate        float64                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 DealRate,omitempty"`
	FltgRateNote    FloatingRateNote2                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 FltgRateNote,omitempty"`
	BrkrdDeal       BrokeredDeal1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 BrkrdDeal,omitempty"`
	CallPutOptn     []Option12                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 CallPutOptn,omitempty"`
	SplmtryData     []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.02 SplmtryData,omitempty"`
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

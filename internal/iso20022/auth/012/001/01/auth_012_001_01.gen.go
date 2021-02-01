// Code generated by main. DO NOT EDIT.

package auth_012_001_01

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

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

type Collateral14 struct {
	Valtn       SecuredCollateral2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Valtn"`
	Hrcut       float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Hrcut,omitempty"`
	SpclCollInd SpecialCollateral1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 SpclCollInd,omitempty"`
}

// May be one of NOPL, POOL
type CollateralPool1Code string

type CollateralValuation6 struct {
	NmnlAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 NmnlAmt,omitempty"`
	ISIN    ISINOct2015Identifier   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 ISIN"`
}

type CollateralValuation7 struct {
	PoolSts CollateralPool1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 PoolSts"`
	Tp      CFIOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Tp"`
	Sctr    string                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Sctr"`
	NmnlAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 NmnlAmt,omitempty"`
}

type CounterpartyIdentification2Choice struct {
	LEI  LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 LEI"`
	Othr ReportedPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Othr"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 DtTm"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 ToDtTm"`
}

type Document struct {
	MnyMktScrdMktSttstclRpt MoneyMarketSecuredMarketStatisticalReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 MnyMktScrdMktSttstclRpt"`
}

type FloatingRateNote2 struct {
	RefRateIndx ISINOct2015Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 RefRateIndx"`
	BsisPtSprd  float64               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 BsisPtSprd"`
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
	RptgAgt LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 RptgAgt"`
	RefPrd  DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 RefPrd"`
}

type MoneyMarketSecuredMarketStatisticalReportV01 struct {
	RptHdr      MoneyMarketReportHeader1   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 RptHdr"`
	ScrdMktRpt  SecuredMarketReport3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 ScrdMktRpt"`
	SplmtryData []SupplementaryData1       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 SplmtryData,omitempty"`
}

// May be one of BORR, LEND
type MoneyMarketTransactionType1Code string

type NameOrSector1Choice struct {
	Nm   Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Nm"`
	Sctr string    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Sctr"`
}

type Rate2 struct {
	Sgn  bool    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Sgn,omitempty"`
	Rate float64 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Rate"`
}

// May be one of NOTX
type ReportPeriodActivity1Code string

type ReportedPartyIdentification1 struct {
	NmOrSctr NameOrSector1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 NmOrSctr"`
	Lctn     CountryCode         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Lctn"`
}

type SecuredCollateral2Choice struct {
	SnglColl  CollateralValuation6   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 SnglColl"`
	MltplColl []CollateralValuation6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 MltplColl"`
	PoolColl  CollateralValuation6   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 PoolColl"`
	OthrColl  []CollateralValuation7 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 OthrColl"`
}

type SecuredMarketReport3Choice struct {
	DataSetActn ReportPeriodActivity1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 DataSetActn"`
	Tx          []SecuredMarketTransaction3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Tx"`
}

type SecuredMarketTransaction3 struct {
	RptdTxSts       TransactionOperationType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 RptdTxSts"`
	BrnchId         LEIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 BrnchId,omitempty"`
	UnqTxIdr        Max105Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 UnqTxIdr,omitempty"`
	PrtryTxId       Max105Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 PrtryTxId"`
	CtrPtyPrtryTxId Max105Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 CtrPtyPrtryTxId,omitempty"`
	CtrPtyId        CounterpartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 CtrPtyId"`
	TrptyAgtId      LEIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 TrptyAgtId,omitempty"`
	TradDt          DateAndDateTimeChoice             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 TradDt"`
	SttlmDt         ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 SttlmDt"`
	MtrtyDt         ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 MtrtyDt"`
	TxTp            MoneyMarketTransactionType1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 TxTp"`
	TxNmnlAmt       ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 TxNmnlAmt"`
	RateTp          InterestRateType1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 RateTp"`
	DealRate        Rate2                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 DealRate,omitempty"`
	FltgRateRpAgrmt FloatingRateNote2                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 FltgRateRpAgrmt,omitempty"`
	BrkrdDeal       BrokeredDeal1Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 BrkrdDeal,omitempty"`
	Coll            Collateral14                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Coll"`
	SplmtryData     []SupplementaryData1              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 SplmtryData,omitempty"`
}

// May be one of GENE, SPEC
type SpecialCollateral1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of AMND, CANC, CORR, NEWT
type TransactionOperationType1Code string

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

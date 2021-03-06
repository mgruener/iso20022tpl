// Code generated by main. DO NOT EDIT.

package colr_015_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification26 struct {
	Prtry SimpleIdentificationInformation4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type Agreement2 struct {
	AgrmtDtls  Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 AgrmtDtls"`
	AgrmtId    Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 AgrmtId,omitempty"`
	AgrmtDt    ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 AgrmtDt"`
	BaseCcy    CurrencyCode              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 BaseCcy"`
	AgrmtFrmwk AgreementFramework1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 AgrmtFrmwk,omitempty"`
}

type AgreementFramework1Choice struct {
	AgrmtFrmwk AgreementFramework1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 AgrmtFrmwk"`
	PrtryId    GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 PrtryId"`
}

// May be one of FBAA, BBAA, DERV, ISDA, NONR
type AgreementFramework1Code string

type AmountAndDirection20 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 CdtDbtInd,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CollateralAccount1 struct {
	Id Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Id"`
	Tp CollateralAccountIdentificationType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Tp,omitempty"`
	Nm Max70Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Nm,omitempty"`
}

type CollateralAccountIdentificationType1Choice struct {
	Tp    CollateralAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Tp,omitempty"`
	Prtry GenericIdentification29    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Prtry"`
}

// May be one of HOUS, CLIE, LIPR, MGIN, DFLT
type CollateralAccountType1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 DtTm"`
}

type DatePeriodDetails struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 ToDt"`
}

type Document struct {
	IntrstPmtStmt InterestPaymentStatementV03 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 IntrstPmtStmt"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be one of BFWD, PAYM, CCPC, COMM, CRDS, CRTL, CRSP, CCIR, CRPR, EQUI, EQPT, EQUS, EXTD, EXPT, FIXI, FORX, FORW, FUTR, OPTN, LIQU, OTCD, REPO, RVPO, SLOA, SBSC, SCRP, SLEB, SHSL, SCIR, SCIE, SWPT, TBAS, TRBD, TRCP
type ExposureType5Code string

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA
type Frequency1Code string

type GenericIdentification29 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 SchmeNm,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 SchmeNm,omitempty"`
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

type InterestCalculation3 struct {
	ClctnDt        ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 ClctnDt"`
	CollAcctId     AccountIdentification26 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 CollAcctId,omitempty"`
	FctvPrncplAmt  AmountAndDirection20    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 FctvPrncplAmt"`
	PrncplAmt      AmountAndDirection20    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 PrncplAmt,omitempty"`
	MvmntAmt       AmountAndDirection20    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 MvmntAmt,omitempty"`
	NbOfDays       float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 NbOfDays,omitempty"`
	FctvRate       float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 FctvRate"`
	IntrstRate     float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 IntrstRate,omitempty"`
	Sprd           float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Sprd,omitempty"`
	AcrdIntrstAmt  AmountAndDirection20    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 AcrdIntrstAmt"`
	AggtdIntrstAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 AggtdIntrstAmt,omitempty"`
}

type InterestPaymentStatementV03 struct {
	TxId        Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 TxId"`
	Agrmt       Agreement2           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Agrmt,omitempty"`
	Oblgtn      Obligation3          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Oblgtn"`
	StmtParams  Statement32          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 StmtParams"`
	Pgntn       Pagination           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Pgntn,omitempty"`
	IntrstStmt  InterestStatement3   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 IntrstStmt"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 SplmtryData,omitempty"`
}

type InterestStatement3 struct {
	IntrstPrd          DatePeriodDetails       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 IntrstPrd"`
	TtlIntrstAmtDueToA ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 TtlIntrstAmtDueToA,omitempty"`
	TtlIntrstAmtDueToB ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 TtlIntrstAmtDueToB,omitempty"`
	ValDt              ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 ValDt"`
	IntrstPmtReqId     Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 IntrstPmtReqId,omitempty"`
	IntrstClctn        []InterestCalculation3  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 IntrstClctn,omitempty"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Adr"`
}

type Obligation3 struct {
	PtyA       PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 PtyA"`
	SvcgPtyA   PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 SvcgPtyA,omitempty"`
	PtyB       PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 PtyB"`
	SvcgPtyB   PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 SvcgPtyB,omitempty"`
	CollAcctId CollateralAccount1          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 CollAcctId,omitempty"`
	XpsrTp     ExposureType5Code           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 XpsrTp,omitempty"`
	ValtnDt    DateAndDateTimeChoice       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 ValtnDt"`
}

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 LastPgInd"`
}

type PartyIdentification33Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 AnyBIC"`
	PrtryId  GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 PrtryId"`
	NmAndAdr NameAndAddress6         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 NmAndAdr"`
}

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Ctry"`
}

type SimpleIdentificationInformation4 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Id"`
}

type Statement32 struct {
	StmtId    Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 StmtId"`
	ActvtyInd bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 ActvtyInd"`
	Frqcy     Frequency1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Frqcy"`
	StmtDtTm  DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 StmtDtTm"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.03 Envlp"`
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

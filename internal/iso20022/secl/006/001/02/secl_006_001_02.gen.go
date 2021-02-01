// Code generated by main. DO NOT EDIT.

package secl_006_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Prtry"`
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

type AlternatePartyIdentification4 struct {
	IdTp    IdentificationType6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 IdTp"`
	Ctry    CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Ctry"`
	AltrnId Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 AltrnId"`
}

type AmountAndDirection21 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 CdtDbtInd,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of HOUS, CLIE, LIPR
type ClearingAccountType1Code string

type Collateral3 struct {
	PstHrcutVal ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 PstHrcutVal"`
	MktVal      ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 MktVal"`
	CollTp      CollateralType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 CollTp"`
}

// May be one of CASH, SECU
type CollateralType2Code string

type Contribution1 struct {
	Acct        AccountIdentification4Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Acct,omitempty"`
	ReqrdAmt    ActiveCurrencyAndAmount         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 ReqrdAmt"`
	IncrCvrgAmt ActiveCurrencyAndAmount         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 IncrCvrgAmt,omitempty"`
	NonClrMmb   PartyIdentificationAndAccount31 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 NonClrMmb,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 DtTm"`
}

type DefaultFund1 struct {
	DfltFndAcct   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 DfltFndAcct"`
	TtlDfltFndAmt ActiveCurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 TtlDfltFndAmt"`
	Cntrbtn       []Contribution1              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Cntrbtn,omitempty"`
	IncrCvrgAmt   ActiveCurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 IncrCvrgAmt,omitempty"`
}

type DefaultFundContributionReportV02 struct {
	RptParams   ReportParameters2           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 RptParams"`
	ClrMmb      PartyIdentification35Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 ClrMmb"`
	RptDtls     []DefaultFundReport1        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 RptDtls"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 SplmtryData,omitempty"`
}

type DefaultFundReport1 struct {
	DfltFndClctn   []DefaultFund1       `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 DfltFndClctn"`
	CollDesc       []Collateral3        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 CollDesc"`
	NetXcssOrDfcit AmountAndDirection21 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 NetXcssOrDfcit"`
}

type Document struct {
	DfltFndCntrbtnRpt DefaultFundContributionReportV02 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 DfltFndCntrbtnRpt"`
}

// May be one of DAIL, INDA, ONDE
type EventFrequency6Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Issr,omitempty"`
}

type GenericIdentification29 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 SchmeNm,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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

type IdentificationType6Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Cd"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Prtry"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Adr"`
}

type PartyIdentification33Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 AnyBIC"`
	PrtryId  GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 PrtryId"`
	NmAndAdr NameAndAddress6         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 NmAndAdr"`
}

type PartyIdentification35Choice struct {
	BIC     AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 BIC"`
	PrtryId GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 PrtryId"`
}

type PartyIdentificationAndAccount31 struct {
	Id       PartyIdentification33Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Id"`
	AltrnId  AlternatePartyIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 AltrnId,omitempty"`
	AddtlInf PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 AddtlInf,omitempty"`
	ClrAcct  SecuritiesAccount18           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 ClrAcct,omitempty"`
}

type PartyTextInformation1 struct {
	DclrtnDtls  Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 DclrtnDtls,omitempty"`
	PtyCtctDtls Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 PtyCtctDtls,omitempty"`
	RegnDtls    Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 RegnDtls,omitempty"`
}

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Ctry"`
}

type ReportParameters2 struct {
	RptId      Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 RptId"`
	RptDtAndTm DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 RptDtAndTm"`
	Frqcy      EventFrequency6Code   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Frqcy"`
	RptCcy     CurrencyCode          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 RptCcy"`
	ClctnDt    ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 ClctnDt,omitempty"`
}

type SecuritiesAccount18 struct {
	Id Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Id"`
	Tp ClearingAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Tp"`
	Nm Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Nm,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of ARNU, CCPT, CHTY, CORP, DRLC, FIIN, TXID
type TypeOfIdentification1Code string

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

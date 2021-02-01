// Code generated by main. DO NOT EDIT.

package tsmt_044_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification3Choice struct {
	IBAN      IBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 IBAN"`
	BBAN      BBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 BBAN"`
	UPIC      UPICIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 UPIC"`
	PrtryAcct SimpleIdentificationInformation2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 PrtryAcct"`
}

type Adjustment4 struct {
	Tp             AdjustmentType2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Tp"`
	OthrAdjstmntTp Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 OthrAdjstmntTp"`
	Drctn          AdjustmentDirection1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Drctn"`
	Amt            CurrencyAndAmount        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Amt"`
}

// May be one of ADDD, SUBS
type AdjustmentDirection1Code string

// May be one of REBA, DISC, CREN, SURC
type AdjustmentType2Code string

// Must match the pattern [a-zA-Z0-9]{1,30}
type BBANIdentifier string

type BICIdentification1 struct {
	BIC BICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 BIC"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

type CashAccount7 struct {
	Id  AccountIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Id"`
	Tp  CashAccountType2             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Tp,omitempty"`
	Ccy CurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Nm,omitempty"`
}

type CashAccountType2 struct {
	Cd    CashAccountType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyAndAmount struct {
	Value float64      `xml:",chardata"`
	Ccy   CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type Document struct {
	InttToPayNtfctn IntentToPayNotificationV01 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 InttToPayNtfctn"`
}

type DocumentIdentification7 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Id"`
	DtOfIsse ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 DtOfIsse"`
}

type FinancialInstitutionIdentification4Choice struct {
	BIC      BICIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 BIC"`
	NmAndAdr NameAndAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 NmAndAdr"`
}

// Must match the pattern [a-zA-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBANIdentifier string

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

type IntentToPay1 struct {
	ByPurchsOrdr ReportLine3      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 ByPurchsOrdr"`
	ByComrclInvc ReportLine4      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 ByComrclInvc"`
	XpctdPmtDt   ISODate          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 XpctdPmtDt"`
	SttlmTerms   SettlementTerms2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 SttlmTerms,omitempty"`
}

type IntentToPayNotificationV01 struct {
	NtfctnId     MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 NtfctnId"`
	TxId         SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 TxId"`
	SubmitrTxRef SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 SubmitrTxRef,omitempty"`
	BuyrBk       BICIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 BuyrBk"`
	SellrBk      BICIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 SellrBk"`
	InttToPay    IntentToPay1                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 InttToPay"`
}

type InvoiceIdentification1 struct {
	InvcNb Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 InvcNb"`
	IsseDt ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 IsseDt"`
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 CreDtTm"`
}

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Adr"`
}

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Ctry"`
}

type ReportLine2 struct {
	TxId          Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 TxId"`
	PurchsOrdrRef DocumentIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 PurchsOrdrRef"`
	Adjstmnt      []Adjustment4           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Adjstmnt,omitempty"`
	NetAmt        CurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 NetAmt"`
}

type ReportLine3 struct {
	PurchsOrdrRef DocumentIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 PurchsOrdrRef"`
	Adjstmnt      []Adjustment4           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Adjstmnt,omitempty"`
	NetAmt        CurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 NetAmt"`
}

type ReportLine4 struct {
	ComrclDocRef       InvoiceIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 ComrclDocRef"`
	Adjstmnt           []Adjustment4          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Adjstmnt,omitempty"`
	NetAmt             CurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 NetAmt"`
	BrkdwnByPurchsOrdr []ReportLine2          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 BrkdwnByPurchsOrdr"`
}

type SettlementTerms2 struct {
	CdtrAgt  FinancialInstitutionIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 CdtrAgt,omitempty"`
	CdtrAcct CashAccount7                              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 CdtrAcct"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Id"`
}

type SimpleIdentificationInformation2 struct {
	Id Max34Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.01 Id"`
}

// Must match the pattern [0-9]{8,17}
type UPICIdentifier string

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

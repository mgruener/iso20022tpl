// Code generated by main. DO NOT EDIT.

package tsmt_045_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Prtry"`
}

// May be one of SBTW, RSTW, RSBS, ARDM, ARCS, ARES, WAIT, UPDT, SBDS, ARBA, ARRO, CINR
type Action2Code string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type Adjustment6 struct {
	Tp    AdjustmentType1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Tp"`
	Drctn AdjustmentDirection1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Drctn"`
	Amt   CurrencyAndAmount        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Amt"`
}

// May be one of ADDD, SUBS
type AdjustmentDirection1Code string

type AdjustmentType1Choice struct {
	Tp             AdjustmentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Tp"`
	OthrAdjstmntTp Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 OthrAdjstmntTp"`
}

// May be one of REBA, DISC, CREN, SURC
type AdjustmentType2Code string

type BICIdentification1 struct {
	BIC BICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 BIC"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

// May be one of PROP, CLSD, PMTC, ESTD, ACTV, COMP, AMRQ, RARQ, CLRQ, SCRQ, SERQ, DARQ
type BaselineStatus3Code string

type BreakDown1Choice struct {
	ByPurchsOrdr ReportLine5 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 ByPurchsOrdr"`
	ByComrclInvc ReportLine6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 ByComrclInvc"`
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Id"`
	Tp  CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Nm,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyAndAmount struct {
	Value float64      `xml:",chardata"`
	Ccy   CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type Document struct {
	FwdInttToPayNtfctn ForwardIntentToPayNotificationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 FwdInttToPayNtfctn"`
}

type DocumentIdentification3 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Id"`
	Vrsn float64   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Vrsn"`
}

type DocumentIdentification5 struct {
	Id     Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Id"`
	IdIssr BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 IdIssr"`
}

type DocumentIdentification7 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Id"`
	DtOfIsse ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 DtOfIsse"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

type FinancialInstitutionIdentification4Choice struct {
	BIC      BICIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 BIC"`
	NmAndAdr NameAndAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 NmAndAdr"`
}

type ForwardIntentToPayNotificationV02 struct {
	NtfctnId          MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 NtfctnId"`
	TxId              SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 TxId"`
	EstblishdBaselnId DocumentIdentification3         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 EstblishdBaselnId"`
	TxSts             TransactionStatus4              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 TxSts"`
	UsrTxRef          []DocumentIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 UsrTxRef,omitempty"`
	BuyrBk            BICIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 BuyrBk"`
	SellrBk           BICIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 SellrBk"`
	InttToPay         IntentToPay2                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 InttToPay"`
	ReqForActn        PendingActivity2                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 ReqForActn,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Issr,omitempty"`
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

type IntentToPay2 struct {
	Brkdwn     BreakDown1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Brkdwn"`
	XpctdPmtDt ISODate          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 XpctdPmtDt"`
	SttlmTerms SettlementTerms3 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 SttlmTerms,omitempty"`
}

type InvoiceIdentification1 struct {
	InvcNb Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 InvcNb"`
	IsseDt ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 IsseDt"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 CreDtTm"`
}

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Adr"`
}

type PendingActivity2 struct {
	Tp   Action2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Tp"`
	Desc Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Desc,omitempty"`
}

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Ctry"`
}

type ReportLine5 struct {
	PurchsOrdrRef DocumentIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 PurchsOrdrRef"`
	Adjstmnt      []Adjustment6           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Adjstmnt,omitempty"`
	NetAmt        CurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 NetAmt"`
}

type ReportLine6 struct {
	ComrclDocRef       InvoiceIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 ComrclDocRef"`
	Adjstmnt           []Adjustment6          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Adjstmnt,omitempty"`
	NetAmt             CurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 NetAmt"`
	BrkdwnByPurchsOrdr []ReportLine7          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 BrkdwnByPurchsOrdr"`
}

type ReportLine7 struct {
	TxId          Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 TxId"`
	PurchsOrdrRef DocumentIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 PurchsOrdrRef"`
	Adjstmnt      []Adjustment6           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Adjstmnt,omitempty"`
	NetAmt        CurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 NetAmt"`
}

type SettlementTerms3 struct {
	CdtrAgt  FinancialInstitutionIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 CdtrAgt,omitempty"`
	CdtrAcct CashAccount24                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 CdtrAcct"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Id"`
}

type TransactionStatus4 struct {
	Sts BaselineStatus3Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Sts"`
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
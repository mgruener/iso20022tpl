// Code generated by main. DO NOT EDIT.

package sese_009_001_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Account25 struct {
	Id       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Id,omitempty"`
	AcctSvcr PartyIdentification113 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 AcctSvcr"`
}

type AdditionalReference8 struct {
	Ref     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Ref"`
	RefIssr PartyIdentification113 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 RefIssr,omitempty"`
	MsgNm   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of NCER, ELEC, PHYS
type BeneficiaryCertificationCompletion1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	ReqForTrfStsRpt RequestForTransferStatusReportV06 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 ReqForTrfStsRpt"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 PlcAndNm"`
	Txt      Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Txt"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 SchmeNm,omitempty"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Id,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type ISOYearMonth time.Time

func (t *ISOYearMonth) UnmarshalText(text []byte) error {
	return (*xsdGYearMonth)(t).UnmarshalText(text)
}
func (t ISOYearMonth) MarshalText() ([]byte, error) {
	return xsdGYearMonth(t).MarshalText()
}

// May be one of CASH, SECU
type IncomePreference2Code string

type Intermediary42 struct {
	Id   PartyIdentification113 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Id"`
	Acct Account25              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Acct,omitempty"`
	Role Role4Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Role,omitempty"`
}

type InvestmentAccount68 struct {
	OwnrId               []PartyIdentification113                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 OwnrId,omitempty"`
	AcctId               Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 AcctId"`
	AcctNm               Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 AcctNm,omitempty"`
	AcctDsgnt            Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 AcctDsgnt,omitempty"`
	IntrmyInf            []Intermediary42                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 IntrmyInf,omitempty"`
	SctiesForm           FormOfSecurity1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 SctiesForm,omitempty"`
	DmtrlsdInd           bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 DmtrlsdInd,omitempty"`
	IncmPref             IncomePreference2Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 IncmPref,omitempty"`
	BnfcryCertfctnCmpltn BeneficiaryCertificationCompletion1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 BnfcryCertfctnCmpltn,omitempty"`
	SfkpgPlc             SafekeepingPlaceFormat8Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 SfkpgPlc,omitempty"`
	AcctSvcr             PartyIdentification113                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 AcctSvcr,omitempty"`
	SubAcctDtls          SubAccount5                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 SubAcctDtls,omitempty"`
}

// May be one of FMCO, REGI, TRAG, INTR, DIST, CONC, UCL1, UCL2, TRAN
type InvestmentFundRole2Code string

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type MarketPracticeVersion1 struct {
	Nm Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Nm"`
	Dt ISOYearMonth `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Dt,omitempty"`
	Nb Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Nb,omitempty"`
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

type MessageAndBusinessReference11 struct {
	Ref             References63Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Ref,omitempty"`
	MstrRef         Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 MstrRef,omitempty"`
	TrfRef          Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 TrfRef"`
	ClntRef         AdditionalReference8 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 ClntRef,omitempty"`
	CxlRef          Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 CxlRef,omitempty"`
	InvstmtAcctDtls InvestmentAccount68  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 InvstmtAcctDtls,omitempty"`
}

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Adr,omitempty"`
}

type PartyIdentification113 struct {
	Pty PartyIdentification90Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Pty"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 LEI,omitempty"`
}

type PartyIdentification90Choice struct {
	AnyBIC   AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 AnyBIC"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Ctry"`
}

type References63Choice struct {
	PrvsRef AdditionalReference8 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 PrvsRef"`
	OthrRef AdditionalReference8 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 OthrRef"`
}

type RequestForTransferStatusReportV06 struct {
	MsgId        MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 MsgId"`
	ReqDtls      []MessageAndBusinessReference11 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 ReqDtls"`
	MktPrctcVrsn MarketPracticeVersion1          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 MktPrctcVrsn,omitempty"`
	Xtnsn        []Extension1                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Xtnsn,omitempty"`
}

type Role4Choice struct {
	Cd    InvestmentFundRole2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Prtry"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat8Choice struct {
	Id      SafekeepingPlaceTypeAndText6             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 TpAndId"`
	Prtry   GenericIdentification78                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Id"`
}

type SafekeepingPlaceTypeAndText6 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Id,omitempty"`
}

type SubAccount5 struct {
	Id    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Id"`
	Nm    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Nm,omitempty"`
	Chrtc Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.06 Chrtc,omitempty"`
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

type xsdGYearMonth time.Time

func (t *xsdGYearMonth) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYearMonth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
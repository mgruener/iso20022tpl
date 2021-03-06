// Code generated by main. DO NOT EDIT.

package seev_040_002_08

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification49 struct {
	SfkpgAcct RestrictedFINXMax35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 SfkpgAcct"`
	AcctOwnr  PartyIdentification136Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat34Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 SfkpgPlc,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, PRII, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH
type CorporateActionEventType20Code string

type CorporateActionEventType58Choice struct {
	Cd    CorporateActionEventType20Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Cd"`
	Prtry GenericIdentification47        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Prtry"`
}

type CorporateActionGeneralInformation121 struct {
	CorpActnEvtId      RestrictedFINXMax16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 CorpActnEvtId"`
	OffclCorpActnEvtId RestrictedFINXMax16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 OffclCorpActnEvtId,omitempty"`
	EvtTp              CorporateActionEventType58Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 EvtTp"`
	FinInstrmId        SecurityIdentification21         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 FinInstrmId,omitempty"`
}

type CorporateActionInstructionCancellationRequest002V08 struct {
	ChngInstrInd   bool                                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 ChngInstrInd,omitempty"`
	InstrId        DocumentIdentification37             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 InstrId"`
	CorpActnGnlInf CorporateActionGeneralInformation121 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 CorpActnGnlInf"`
	AcctDtls       AccountIdentification49              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 AcctDtls"`
	CorpActnInstr  CorporateActionOption128             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 CorpActnInstr"`
	PrtctInstr     ProtectInstruction7                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 PrtctInstr,omitempty"`
	SplmtryData    []SupplementaryData1                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 SplmtryData,omitempty"`
}

type CorporateActionOption128 struct {
	OptnNb   OptionNumber1Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 OptnNb"`
	OptnTp   CorporateActionOption29Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 OptnTp"`
	InstdQty Quantity40Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 InstdQty"`
}

type CorporateActionOption29Choice struct {
	Cd    CorporateActionOption9Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Cd"`
	Prtry GenericIdentification47    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Prtry"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CERT, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, TAXI, PRUN
type CorporateActionOption9Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	CorpActnInstrCxlReq CorporateActionInstructionCancellationRequest002V08 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 CorpActnInstrCxlReq"`
}

type DocumentIdentification37 struct {
	Id    RestrictedFINXMax16Text    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Id"`
	LkgTp ProcessingPosition10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 LkgTp,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity15Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 AmtsdVal"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 SchmeNm,omitempty"`
}

type GenericIdentification84 struct {
	Id      RestrictedFINXMax34Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 SchmeNm,omitempty"`
}

type GenericIdentification89 struct {
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Tp"`
	Id RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Id,omitempty"`
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

type IdentificationSource4Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Cd"`
	Prtry RestrictedFINExact2Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Prtry"`
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities4 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 AmtsdVal"`
}

type OtherIdentification2 struct {
	Id  RestrictedFINXMax31Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Sfx,omitempty"`
	Tp  IdentificationSource4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Tp"`
}

type PartyIdentification136Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 AnyBIC"`
	PrtryId GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 PrtryId"`
}

type ProcessingPosition10Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Prtry"`
}

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

type ProtectInstruction7 struct {
	TxTp    ProtectTransactionType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 TxTp"`
	TxId    RestrictedFINMax15Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 TxId,omitempty"`
	PrtctDt ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 PrtctDt,omitempty"`
}

// May be one of PROT
type ProtectTransactionType3Code string

// May be one of QALL
type Quantity1Code string

type Quantity40Choice struct {
	Cd                 Quantity1Code                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Cd"`
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities4       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 OrgnlAndCurFaceAmt"`
	Qty                FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Qty"`
}

// Must match the pattern XX|TS
type RestrictedFINExact2Text string

// Must be at least 1 items long
type RestrictedFINMax15Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,140}
type RestrictedFINXMax140Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax16Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax30Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,31}
type RestrictedFINXMax31Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax34Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,35}
type RestrictedFINXMax35Text string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat34Choice struct {
	Id      SafekeepingPlaceTypeAndText9           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Id"`
	Ctry    CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Ctry"`
	TpAndId SafekeepingPlaceTypeAndIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 TpAndId"`
	Prtry   GenericIdentification89                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Prtry"`
}

type SafekeepingPlaceTypeAndIdentification1 struct {
	SfkpgPlcTp SafekeepingPlace1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 SfkpgPlcTp"`
	Id         AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Id"`
}

type SafekeepingPlaceTypeAndText9 struct {
	SfkpgPlcTp SafekeepingPlace2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 SfkpgPlcTp"`
	Id         RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Id,omitempty"`
}

type SecurityIdentification21 struct {
	ISIN   ISINOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 ISIN,omitempty"`
	OthrId []OtherIdentification2   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 OthrId,omitempty"`
	Desc   RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Desc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.08 Envlp"`
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

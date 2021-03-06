// Code generated by main. DO NOT EDIT.

package caaa_006_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorCancellationResponse2 struct {
	Envt   CardPaymentEnvironment11 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Envt"`
	Tx     CardPaymentTransaction6  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Tx"`
	TxRspn CardPaymentTransaction10 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 TxRspn"`
}

type AcceptorCancellationResponseV02 struct {
	Hdr      Header1                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Hdr"`
	CxlRspn  AcceptorCancellationResponse2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 CxlRspn"`
	SctyTrlr ContentInformationType6       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 SctyTrlr"`
}

type Action1 struct {
	ActnTp    ActionType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 ActnTp"`
	MsgToPres ActionMessage1  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 MsgToPres,omitempty"`
}

type ActionMessage1 struct {
	MsgDstn      UserInterface1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 MsgDstn"`
	MsgCntt      Max256Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 MsgCntt"`
	MsgCnttSgntr Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 MsgCnttSgntr,omitempty"`
}

// May be one of DISP, PRNT, RFRL, CPTR, PINR, PINL, RQDT, BUSY, RQID
type ActionType1Code string

// May be one of EA2C, E3DC, DKPT, DKP9, UKPT, UKA1
type Algorithm2Code string

// May be one of MACC, MCCS, CMA1, CMD1
type Algorithm3Code string

// May be one of HS25, HS38, HS51
type Algorithm5Code string

// May be one of EA2C, E3DC
type Algorithm6Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification2 struct {
	Algo  Algorithm2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Algo"`
	Param Parameter1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Param,omitempty"`
}

type AlgorithmIdentification3 struct {
	Algo  Algorithm3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Algo"`
	Param Parameter1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Param,omitempty"`
}

type AlgorithmIdentification6 struct {
	Algo  Algorithm6Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Algo"`
	Param Parameter1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Param,omitempty"`
}

type AlgorithmIdentification7 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Algo"`
	Param Parameter2     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Param,omitempty"`
}

type AlgorithmIdentification8 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Algo"`
	Param Parameter3     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData2 struct {
	Vrsn        float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Vrsn,omitempty"`
	Rcpt        []Recipient2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Rcpt"`
	MACAlgo     AlgorithmIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 MACAlgo"`
	NcpsltdCntt EncapsulatedContent1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 NcpsltdCntt"`
	MAC         Max35Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 MAC"`
}

type AuthorisationResult1 struct {
	AuthstnNtty   GenericIdentification33 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 AuthstnNtty,omitempty"`
	RspnToAuthstn ResponseType1           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 RspnToAuthstn"`
	AuthstnCd     Min6Max8Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 AuthstnCd,omitempty"`
	CmpltnReqrd   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 CmpltnReqrd,omitempty"`
	TMSTrggr      TMSTrigger1             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 TMSTrggr,omitempty"`
}

type CardPaymentEnvironment11 struct {
	AcqrrId        GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 AcqrrId,omitempty"`
	MrchntId       GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 MrchntId,omitempty"`
	POIId          GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 POIId"`
	PrtctdCardData ContentInformationType5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 PrtctdCardData,omitempty"`
	PlainCardData  PlainCardData3          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 PlainCardData,omitempty"`
}

type CardPaymentTransaction10 struct {
	AuthstnRslt AuthorisationResult1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 AuthstnRslt"`
	Actn        []Action1            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Actn,omitempty"`
}

type CardPaymentTransaction6 struct {
	TxId         TransactionIdentifier1         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 TxId"`
	RcptTxId     Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 RcptTxId,omitempty"`
	RcncltnId    Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 RcncltnId,omitempty"`
	IntrchngData Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 IntrchngData,omitempty"`
	TxDtls       CardPaymentTransactionDetails6 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 TxDtls"`
}

type CardPaymentTransactionDetails6 struct {
	Ccy    CurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Ccy"`
	TtlAmt float64      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 TtlAmt"`
}

type CertificateIdentifier1 struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 IssrAndSrlNb"`
}

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 RltvDstngshdNm"`
}

type ContentInformationType5 struct {
	CnttTp     ContentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 CnttTp"`
	EnvlpdData EnvelopedData2   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 EnvlpdData"`
}

type ContentInformationType6 struct {
	CnttTp       ContentType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 CnttTp"`
	AuthntcdData []AuthenticatedData2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 AuthntcdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, ECRP, AUTH
type ContentType1Code string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type Document struct {
	AccptrCxlRspn AcceptorCancellationResponseV02 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 AccptrCxlRspn"`
}

type EncapsulatedContent1 struct {
	CnttTp ContentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 CnttTp"`
	Cntt   Max10000Binary   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Cntt,omitempty"`
}

type EncryptedContent2 struct {
	CnttTp         ContentType1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 CnttNcrptnAlgo"`
	NcrptdData     Max10000Binary           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 NcrptdData"`
}

type EnvelopedData2 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Vrsn,omitempty"`
	Rcpt       []Recipient2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Rcpt"`
	NcrptdCntt EncryptedContent2  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 NcrptdCntt"`
}

// May be no more than 10 items long
type Exact10Text string

type GenericIdentification31 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Tp"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 ShrtNm,omitempty"`
}

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 ShrtNm,omitempty"`
}

type GenericIdentification33 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Id,omitempty"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Tp"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 ShrtNm,omitempty"`
}

type Header1 struct {
	MsgFctn    MessageFunction1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 PrtcolVrsn"`
	XchgId     Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 XchgId"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 CreDtTm"`
	InitgPty   GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 InitgPty"`
	RcptPty    GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 RcptPty,omitempty"`
	Tracblt    []Traceability1         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Tracblt,omitempty"`
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

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 SrlNb"`
}

type KEK2 struct {
	Vrsn          float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Vrsn,omitempty"`
	KEKId         KEKIdentifier1           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 NcrptdKey"`
}

type KEKIdentifier1 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 KeyId"`
	KeyVrsn   Exact10Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 KeyVrsn"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 DerivtnId,omitempty"`
}

type KeyTransport2 struct {
	Vrsn          float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Vrsn"`
	RcptId        CertificateIdentifier1   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 NcrptdKey"`
}

type Max10000Binary []byte

func (t *Max10000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max256Text string

type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

type Max500Binary []byte

func (t *Max500Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max500Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// May be one of AUTQ, AUTP, FAUQ, FAUP, CMPV, CMPK, FCMV, FCMK, RVRA, RVRR, FRVA, FRVR, CCAQ, CCAP, CCAV, CCAK, DGNP, DGNQ, RCLQ, RCLP, RJCT
type MessageFunction1Code string

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 6 items long
type Min6Max8Text string

// Must match the pattern [0-9]{8,28}
type Min8Max28NumericText string

type Parameter1 struct {
	InitlstnVctr Max500Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 InitlstnVctr,omitempty"`
}

type Parameter2 struct {
	DgstAlgo     Algorithm5Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 MskGnrtrAlgo,omitempty"`
}

type Parameter3 struct {
	DgstAlgo Algorithm5Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 DgstAlgo,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

type PlainCardData3 struct {
	PAN       Min8Max28NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 PAN"`
	CardSeqNb Min2Max3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 CardSeqNb,omitempty"`
	FctvDt    ISOYearMonth         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 FctvDt,omitempty"`
	XpryDt    ISOYearMonth         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 XpryDt"`
}

type Recipient2Choice struct {
	KeyTrnsprt KeyTransport2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 KeyTrnsprt,omitempty"`
	KEK        KEK2          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 KEK,omitempty"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 AttrVal"`
}

// May be one of DECL, APPR, PART, TECH
type Response1Code string

type ResponseType1 struct {
	Rspn    Response1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Rspn"`
	RspnRsn Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 RspnRsn,omitempty"`
}

// May be one of CRIT, ASAP, DTIM
type TMSContactLevel1Code string

type TMSTrigger1 struct {
	TMSCtctLvl  TMSContactLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 TMSCtctLvl"`
	TMSId       Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 TMSId,omitempty"`
	TMSCtctDtTm ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 TMSCtctDtTm,omitempty"`
}

type Traceability1 struct {
	RlayId      GenericIdentification31 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 RlayId"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 TracDtTmOut"`
}

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 TxRef"`
}

// May be one of CDSP, CRCP, MDSP, MRCP
type UserInterface1Code string

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
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

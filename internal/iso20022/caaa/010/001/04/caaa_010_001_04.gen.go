// Code generated by main. DO NOT EDIT.

package caaa_010_001_04

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorReconciliationResponse3 struct {
	Envt   CardPaymentEnvironment38   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Envt"`
	TxRspn ResponseType1              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 TxRspn"`
	Tx     TransactionReconciliation3 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Tx"`
}

type AcceptorReconciliationResponseV04 struct {
	Hdr         Header10                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Hdr"`
	RcncltnRspn AcceptorReconciliationResponse3 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 RcncltnRspn"`
	SctyTrlr    ContentInformationType11        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 SctyTrlr"`
}

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5
type Algorithm12Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 MAC"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CardPaymentEnvironment38 struct {
	AcqrrId  GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 AcqrrId"`
	MrchntId GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 MrchntId,omitempty"`
	POIId    GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 POIId,omitempty"`
}

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 RltvDstngshdNm"`
}

type ContentInformationType11 struct {
	CnttTp       ContentType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 CnttTp"`
	AuthntcdData []AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 AuthntcdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type Document struct {
	AccptrRcncltnRspn AcceptorReconciliationResponseV04 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 AccptrRcncltnRspn"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Cntt,omitempty"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 ShrtNm,omitempty"`
}

type GenericIdentification53 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Tp,omitempty"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 ShrtNm,omitempty"`
}

type GenericIdentification76 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 ShrtNm,omitempty"`
}

type Header10 struct {
	MsgFctn    MessageFunction4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 PrtcolVrsn"`
	XchgId     Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 XchgId"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 CreDtTm"`
	InitgPty   GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 InitgPty"`
	RcptPty    GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 RcptPty,omitempty"`
	Tracblt    []Traceability2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 NcrptdKey"`
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
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

type Max5000Binary []byte

func (t *Max5000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max5000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

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

// May be one of AUTQ, AUTP, FAUQ, FAUP, CMPV, CMPK, FCMV, FCMK, RVRA, RVRR, FRVA, FRVR, CCAQ, CCAP, CCAV, CCAK, DGNP, DGNQ, RCLQ, RCLP, RJCT, DCCQ, DCCP
type MessageFunction4Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 BPddg,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 KeyIdr"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 AttrVal"`
}

// May be one of DECL, APPR, PART, TECH
type Response1Code string

type ResponseType1 struct {
	Rspn    Response1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Rspn"`
	RspnRsn Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 RspnRsn,omitempty"`
}

type Traceability2 struct {
	RlayId      GenericIdentification76 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 RlayId"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 TracDtTmOut"`
}

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 TxRef"`
}

type TransactionReconciliation3 struct {
	ClsPrd      bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 ClsPrd,omitempty"`
	RcncltnTxId TransactionIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 RcncltnTxId"`
	RcncltnId   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 RcncltnId"`
	TxTtls      []TransactionTotals3   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 TxTtls,omitempty"`
	AddtlTxData Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 AddtlTxData,omitempty"`
}

type TransactionTotals3 struct {
	POIGrpId     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 POIGrpId,omitempty"`
	CardPdctPrfl Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 CardPdctPrfl,omitempty"`
	Ccy          CurrencyCode               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Ccy,omitempty"`
	Tp           TypeTransactionTotals2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Tp"`
	TtlNb        float64                    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 TtlNb"`
	CmltvAmt     float64                    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 CmltvAmt"`
}

// May be one of CRDT, CRDR, DEBT, DBTR, DECL, FAIL
type TypeTransactionTotals2Code string

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

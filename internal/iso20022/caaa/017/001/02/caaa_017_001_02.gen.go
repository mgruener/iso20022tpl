// Code generated by main. DO NOT EDIT.

package caaa_017_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorCurrencyConversionResponse2 struct {
	Envt     CardPaymentEnvironment33 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Envt"`
	Tx       CardPaymentTransaction38 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Tx"`
	CcyConvs CurrencyConversion3      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CcyConvs"`
}

type AcceptorCurrencyConversionResponseV02 struct {
	Hdr          Header10                            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Hdr"`
	CcyConvsRspn AcceptorCurrencyConversionResponse2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CcyConvsRspn"`
	SctyTrlr     ContentInformationType11            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 SctyTrlr"`
}

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5
type Algorithm12Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of EA2C, E3DC, EA9C, EA5C
type Algorithm15Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 MAC"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// May be one of CTDP, CHCK, CRDT, CURR, CDBT, DFLT, EPRS, HEQL, ISTL, INVS, LCDT, MBNW, MNMK, MNMC, MTGL, RTRM, RVLV, SVNG, STBD, UVRL
type CardAccountType2Code string

type CardPaymentEnvironment33 struct {
	AcqrrId  GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 AcqrrId,omitempty"`
	MrchntId GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 MrchntId,omitempty"`
	POIId    GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 POIId"`
	Card     PaymentCard10           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Card,omitempty"`
	PmtTkn   CardPaymentToken2       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 PmtTkn,omitempty"`
}

type CardPaymentToken2 struct {
	TknChrtc     []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 TknChrtc,omitempty"`
	TknAssrncLvl float64     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 TknAssrncLvl,omitempty"`
}

type CardPaymentTransaction38 struct {
	SaleRefId    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 SaleRefId,omitempty"`
	TxId         TransactionIdentifier1          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 TxId"`
	RcptTxId     Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 RcptTxId,omitempty"`
	RcncltnId    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 RcncltnId,omitempty"`
	IntrchngData Max140Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 IntrchngData,omitempty"`
	TxDtls       CardPaymentTransactionDetails20 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 TxDtls"`
}

type CardPaymentTransactionDetails20 struct {
	Ccy         CurrencyCode         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Ccy"`
	TtlAmt      float64              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 TtlAmt"`
	DtldAmt     DetailedAmount7      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 DtldAmt,omitempty"`
	VldtyDt     ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 VldtyDt,omitempty"`
	AcctTp      CardAccountType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 AcctTp,omitempty"`
	ICCRltdData Max10000Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 ICCRltdData,omitempty"`
}

// May be one of COMM, CONS
type CardProductType1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 RltvDstngshdNm"`
}

type Commission18 struct {
	Rate     float64    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Rate"`
	AddtlInf Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 AddtlInf,omitempty"`
}

type Commission19 struct {
	Amt      float64    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Amt"`
	AddtlInf Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 AddtlInf,omitempty"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 EnvlpdData"`
}

type ContentInformationType11 struct {
	CnttTp       ContentType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CnttTp"`
	AuthntcdData []AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 AuthntcdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type CurrencyConversion2 struct {
	CcyConvsId    Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CcyConvsId,omitempty"`
	TrgtCcy       CurrencyDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 TrgtCcy"`
	RsltgAmt      float64          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 RsltgAmt"`
	XchgRate      float64          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 XchgRate"`
	XchgRateDcml  float64          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 XchgRateDcml,omitempty"`
	NvrtdXchgRate float64          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 NvrtdXchgRate,omitempty"`
	QtnDt         ISODateTime      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 QtnDt,omitempty"`
	VldUntil      ISODateTime      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 VldUntil,omitempty"`
	SrcCcy        CurrencyDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 SrcCcy"`
	OrgnlAmt      float64          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 OrgnlAmt"`
	ComssnDtls    []Commission19   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 ComssnDtls,omitempty"`
	MrkUpDtls     []Commission18   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 MrkUpDtls,omitempty"`
	DclrtnDtls    Max2048Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 DclrtnDtls,omitempty"`
}

type CurrencyConversion3 struct {
	Rslt    CurrencyConversionResponse1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Rslt"`
	RsltRsn Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 RsltRsn,omitempty"`
	Convs   CurrencyConversion2             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Convs,omitempty"`
}

// May be one of ODCC, DCCA, ICRD, IMER, IPRD, IRAT, NDCC
type CurrencyConversionResponse1Code string

type CurrencyDetails1 struct {
	AlphaCd CurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 AlphaCd"`
	NmrcCd  Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 NmrcCd"`
	Dcml    float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Dcml"`
	Nm      Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Nm,omitempty"`
}

type DetailedAmount4 struct {
	Amt  float64    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Amt"`
	Labl Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Labl,omitempty"`
}

type DetailedAmount7 struct {
	CshBck      float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CshBck,omitempty"`
	Grtty       float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Grtty,omitempty"`
	Fees        []DetailedAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Fees,omitempty"`
	Rbt         []DetailedAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Rbt,omitempty"`
	ValAddedTax []DetailedAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 ValAddedTax,omitempty"`
	Srchrg      []DetailedAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Srchrg,omitempty"`
}

type Document struct {
	AccptrCcyConvsRspn AcceptorCurrencyConversionResponseV02 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 AccptrCcyConvsRspn"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 NcrptdCntt,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 ShrtNm,omitempty"`
}

type GenericIdentification53 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Tp,omitempty"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 ShrtNm,omitempty"`
}

type GenericIdentification76 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 ShrtNm,omitempty"`
}

type Header10 struct {
	MsgFctn    MessageFunction4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 PrtcolVrsn"`
	XchgId     Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 XchgId"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CreDtTm"`
	InitgPty   GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 InitgPty"`
	RcptPty    GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 RcptPty,omitempty"`
	Tracblt    []Traceability2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Tracblt,omitempty"`
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

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 NcrptdKey"`
}

type Max10000Binary []byte

func (t *Max10000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max10Text string

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
type Max2048Text string

// Must be at least 1 items long
type Max350Text string

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

// May be one of AUTQ, AUTP, FAUQ, FAUP, CMPV, CMPK, FCMV, FCMK, RVRA, RVRR, FRVA, FRVR, CCAQ, CCAP, CCAV, CCAK, DGNP, DGNQ, RCLQ, RCLP, RJCT, DCCQ, DCCP
type MessageFunction4Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{8,28}
type Min8Max28NumericText string

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 BPddg,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

type PaymentCard10 struct {
	PrtctdCardData ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 PrtctdCardData,omitempty"`
	PlainCardData  PlainCardData8           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 PlainCardData,omitempty"`
	MskdPAN        string                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 MskdPAN,omitempty"`
	CardBrnd       Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CardBrnd,omitempty"`
	CardPdctTp     CardProductType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CardPdctTp,omitempty"`
}

type PlainCardData8 struct {
	PAN       Min8Max28NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 PAN"`
	CardSeqNb Min2Max3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 CardSeqNb,omitempty"`
	FctvDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 FctvDt,omitempty"`
	XpryDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 XpryDt"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 KeyIdr"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 AttrVal"`
}

type Traceability2 struct {
	RlayId      GenericIdentification76 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 RlayId"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 TracDtTmOut"`
}

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 TxRef"`
}

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

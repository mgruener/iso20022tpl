// Code generated by main. DO NOT EDIT.

package cain_011_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcquirerKeyExchangeInitiation1 struct {
	Envt CardTransactionEnvironment6 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Envt"`
	Tx   CardTransaction13           `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Tx"`
}

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5
type Algorithm12Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of ERS2, ERS1, RPSS
type Algorithm14Code string

// May be one of EA2C, E3DC, EA9C, EA5C
type Algorithm15Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Param,omitempty"`
}

type AlgorithmIdentification16 struct {
	Algo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Algo"`
}

type AlgorithmIdentification17 struct {
	Algo  Algorithm14Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Algo"`
	Param Parameter8      `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 MAC"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// May be one of KYDL, KYCG, KYVF
type CardServiceType3Code string

type CardTransaction13 struct {
	KeyXchgTp CardServiceType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KeyXchgTp"`
	InitrDtTm ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 InitrDtTm"`
	ReqdKey   []KEKIdentifier3     `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 ReqdKey,omitempty"`
	Key       []CryptographicKey6  `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Key,omitempty"`
	TxRspn    ResponseType2        `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 TxRspn,omitempty"`
}

type CardTransactionEnvironment6 struct {
	SndgInstn GenericIdentification73 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 SndgInstn"`
	RcvgInstn GenericIdentification73 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 RcvgInstn"`
}

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 RltvDstngshdNm"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 EnvlpdData"`
}

type ContentInformationType12 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 CnttTp"`
	EnvlpdData   EnvelopedData4     `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 EnvlpdData,omitempty"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 AuthntcdData,omitempty"`
	SgndData     SignedData4        `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 SgndData,omitempty"`
	DgstdData    DigestedData4      `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 DgstdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

type CryptographicKey6 struct {
	Nm           Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Nm,omitempty"`
	Id           Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Id"`
	Vrsn         Max256Text                `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Vrsn,omitempty"`
	Tp           CryptographicKeyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Tp"`
	Fctn         []KeyUsage1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Fctn"`
	ActvtnDt     ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 ActvtnDt,omitempty"`
	DeactvtnDt   ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 DeactvtnDt,omitempty"`
	NcrptdKeyVal ContentInformationType10  `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 NcrptdKeyVal,omitempty"`
	Cert         []Max5000Binary           `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Cert,omitempty"`
	ICCRltdData  Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 ICCRltdData,omitempty"`
}

// May be one of AES2, EDE3, DKP9, AES9, AES5, EDE4
type CryptographicKeyType3Code string

type DigestedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Vrsn,omitempty"`
	DgstAlgo    AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 NcpsltdCntt"`
	Dgst        Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Dgst"`
}

type Document struct {
	KeyXchgInitn KeyExchangeInitiation `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KeyXchgInitn"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 NcrptdCntt,omitempty"`
}

type GenericIdentification73 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Id"`
	Tp     PartyType9Code    `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Tp,omitempty"`
	Issr   PartyType9Code    `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 ShrtNm,omitempty"`
}

type GenericIdentification74 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Id"`
	Tp     PartyType10Code   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Tp"`
	Issr   PartyType10Code   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 ShrtNm,omitempty"`
}

type Header17 struct {
	MsgFctn        MessageFunction6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 MsgFctn"`
	PrtcolVrsn     Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 PrtcolVrsn"`
	XchgId         Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 XchgId"`
	ReTrnsmssnCntr Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 ReTrnsmssnCntr,omitempty"`
	CreDtTm        ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 CreDtTm"`
	InitgPty       GenericIdentification73 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 InitgPty"`
	RcptPty        GenericIdentification73 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 RcptPty,omitempty"`
	Tracblt        []Traceability3         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 DerivtnId,omitempty"`
}

type KEKIdentifier3 struct {
	Nm         Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Nm,omitempty"`
	Id         Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Id"`
	Vrsn       Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Vrsn,omitempty"`
	KeyChckVal Max35Binary `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KeyChckVal,omitempty"`
}

type KeyExchangeInitiation struct {
	Hdr          Header17                       `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Hdr"`
	KeyXchgInitn AcquirerKeyExchangeInitiation1 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KeyXchgInitn"`
	SctyTrlr     ContentInformationType12       `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 SctyTrlr"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 NcrptdKey"`
}

// May be one of ENCR, DCPT, DENC, DDEC, TRNI, TRNX, MACG, MACV, SIGG, SUGV, PINE, PIND, PINV, KEYG, KEYI, KEYX, KEYD
type KeyUsage1Code string

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

// Must be at least 1 items long
type Max256Text string

type Max3000Binary []byte

func (t *Max3000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max3000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

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

// May be one of RCAV, RCAN, RCAQ, REJA, REVV, REVN, REVQ, RCPV, RCPN, RCPQ, REJP, AUTV, AUTN, AUTQ, AUTP, FNCV, FNCN, FNCQ, RCIV, RCIN, RCIQ, REJI, KEYV, KEYQ, MGTV, MGTQ
type MessageFunction6Code string

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
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 BPddg,omitempty"`
}

type Parameter8 struct {
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 DgstAlgo"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 MskGnrtrAlgo"`
	SaltLngth    float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 SaltLngth"`
	TrlrFld      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 TrlrFld,omitempty"`
}

// May be one of ACCP, ACQR, ATMG, CISS, DLIS, HSTG, ITAG, MERC, OATM, OPOI
type PartyType10Code string

// May be one of ACQR, ACQP, CISS, CISP, CSCH, SCHP
type PartyType9Code string

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KeyIdr"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 AttrVal"`
}

// May be one of APPR, DECL, FRTH, PART, PRCS, UNPR
type Response3Code string

type ResponseType2 struct {
	Rslt         Response3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Rslt"`
	RsltDtls     ResultDetail1Code `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 RsltDtls,omitempty"`
	AddtlRsltInf Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 AddtlRsltInf,omitempty"`
}

// May be one of ACTF, ACQS, AMLV, AMTA, BANK, CRDR, CRDF, CSHI, CSHE, ACTC, CTVG, DBER, FEES, TXNL, CRDX, FMTR, ACEF, TXNG, FNDI, ACPI, AMTI, CHDI, CRDI, CTFV, NPRA, PINA, LBLU, CRDA, PINN, MACK, MACR, CRDL, LBLA, ISSU, ISST, ISSO, ISSF, ISSP, DATI, TXNV, TKID, TKKO, CSCV, PINV, AMTO, NPRC, OFFL, ONLP, TXNM, OTHR, BALO, SEQO, AMTL, NMBL, PINC, PIND, PINE, PINS, PINX, QMAX, RECD, CRDT, SECV, SFWE, SPCC, CRDW, NMBW, AMTW, VNDF, VNDR, SVSU, CRDU, CMKY, UNBC, UNBP, UNBO, ORGF, TXND, TXNU, TTLV, ACTT, SYSM, SYSP, FRDS, CNTC, SRCH, CRDS, ACKO
type ResultDetail1Code string

type SignedData4 struct {
	Vrsn        float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3        `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 NcpsltdCntt"`
	Cert        []Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Cert,omitempty"`
	Sgnr        []Signer3                   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Sgnr"`
}

type Signer3 struct {
	Vrsn      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Vrsn,omitempty"`
	SgnrId    Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 SgnrId,omitempty"`
	DgstAlgo  AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 DgstAlgo"`
	SgntrAlgo AlgorithmIdentification17 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 SgntrAlgo"`
	Sgntr     Max3000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Sgntr"`
}

type Traceability3 struct {
	RlayId      GenericIdentification74 `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 RlayId"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 TracDtTmOut"`
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

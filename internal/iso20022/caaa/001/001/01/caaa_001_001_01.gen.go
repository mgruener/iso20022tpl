// Code generated by main. DO NOT EDIT.

package caaa_001_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorAuthorisationRequest1 struct {
	Envt  CardPaymentEnvironment1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Envt"`
	Cntxt CardPaymentContext1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Cntxt"`
	Tx    CardPaymentTransaction1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Tx"`
}

type AcceptorAuthorisationRequestV01 struct {
	Hdr        Header1                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Hdr"`
	AuthstnReq AcceptorAuthorisationRequest1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AuthstnReq"`
	SctyTrlr   ContentInformationType3       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SctyTrlr"`
}

type Acquirer1 struct {
	Id         GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Id,omitempty"`
	ParamsVrsn ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 ParamsVrsn"`
}

type AddressVerification1 struct {
	AdrDgts    Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AdrDgts,omitempty"`
	PstlCdDgts Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PstlCdDgts,omitempty"`
}

// May be one of MACC, MCCS, UKPT, DKPT, E3DC, HS25, ERS2, ERSA
type Algorithm1Code string

type AlgorithmIdentification1 struct {
	Algo  Algorithm1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Algo"`
	Param Parameter1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Param,omitempty"`
}

// May be one of ATTD, SATT, UATT
type AttendanceContext1Code string

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData1 struct {
	Vrsn        float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Vrsn,omitempty"`
	Rcpt        []Recipient1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Rcpt"`
	MACAlgo     AlgorithmIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 MACAlgo"`
	NcpsltdCntt EncapsulatedContent1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 NcpsltdCntt"`
	MAC         Max35Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 MAC"`
}

// May be one of ICCD, AGNT, MERC
type AuthenticationEntity1Code string

// May be one of UKNW, BYPS, NPIN, FPIN, CPSG, PPSG, MANU, MERC, SCRT, SNCT, SCNL
type AuthenticationMethod1Code string

// May be one of PRST, BYPS, UNRD, NCSC
type CSCManagement1Code string

// May be one of DFLT, SVNG, CHCK, CRDT, UVRL, INVS, EPRS
type CardAccountType1Code string

// May be one of TAGC, PHYS, BRCD, MGST, CICC, DFLE, CTLS, ECTL
type CardDataReading1Code string

type CardPaymentContext1 struct {
	PmtCntxt  PaymentContext1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PmtCntxt"`
	SaleCntxt SaleContext1    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SaleCntxt,omitempty"`
}

type CardPaymentEnvironment1 struct {
	Acqrr   Acquirer1           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Acqrr,omitempty"`
	Mrchnt  Organisation5       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Mrchnt,omitempty"`
	POI     PointOfInteraction1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 POI"`
	Card    PaymentCard1        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Card"`
	Crdhldr Cardholder1         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Crdhldr,omitempty"`
}

// May be one of CAFT, ORCR, CRDP, CSHB, CSHW, CSHD, DEFR, RESA, LOAD, RFND, QUCH, BALC, CACT, CAVR, PINC, VALC
type CardPaymentServiceType1Code string

// May be one of AGGR, DCCV, GRTT, INSP, LOYT, NRES, PUCO, RECP, SOAF, UNAF, VCAU
type CardPaymentServiceType2Code string

// May be one of IRES, URES, PRES, ARES, FREC, RREC
type CardPaymentServiceType3Code string

type CardPaymentTransaction1 struct {
	TxCaptr      bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TxCaptr"`
	TxTp         CardPaymentServiceType1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TxTp"`
	AddtlSvc     []CardPaymentServiceType2Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AddtlSvc,omitempty"`
	SvcAttr      CardPaymentServiceType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SvcAttr,omitempty"`
	MrchntCtgyCd Min3Max4Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 MrchntCtgyCd"`
	TxId         TransactionIdentifier1         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TxId"`
	OrgnlTx      CardPaymentTransaction8        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 OrgnlTx,omitempty"`
	InitrTxId    Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 InitrTxId,omitempty"`
	RcncltnId    Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 RcncltnId,omitempty"`
	TxDtls       CardPaymentTransactionDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TxDtls"`
	AddtlTxData  Max70Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AddtlTxData,omitempty"`
}

type CardPaymentTransaction8 struct {
	TxId      TransactionIdentifier1        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TxId"`
	POIId     GenericIdentification32       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 POIId,omitempty"`
	InitrTxId Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 InitrTxId,omitempty"`
	RcptTxId  Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 RcptTxId,omitempty"`
	TxTp      CardPaymentServiceType1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TxTp"`
	AddtlSvc  []CardPaymentServiceType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AddtlSvc,omitempty"`
	SvcAttr   CardPaymentServiceType3Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SvcAttr,omitempty"`
	TxRslt    CardPaymentTransactionResult1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TxRslt,omitempty"`
}

type CardPaymentTransactionDetails1 struct {
	Ccy            CurrencyCode          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Ccy"`
	TtlAmt         float64               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TtlAmt"`
	AmtQlfr        TypeOfAmount1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AmtQlfr,omitempty"`
	DtldAmt        []DetailedAmount1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 DtldAmt,omitempty"`
	VldtyDt        ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 VldtyDt,omitempty"`
	OnLineRsn      OnLineReason1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 OnLineRsn,omitempty"`
	UattnddLvlCtgy Max35NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 UattnddLvlCtgy,omitempty"`
	AcctTp         CardAccountType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AcctTp,omitempty"`
	RcrngTx        RecurringTransaction1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 RcrngTx,omitempty"`
	Pdct           []Product1            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Pdct,omitempty"`
	ICCRltdData    Max10000Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 ICCRltdData,omitempty"`
}

type CardPaymentTransactionResult1 struct {
	AuthstnNtty   GenericIdentification33 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AuthstnNtty,omitempty"`
	RspnToAuthstn ResponseType1           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 RspnToAuthstn"`
	AuthstnCd     Min6Max8Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AuthstnCd,omitempty"`
}

type CardSecurityInformation1 struct {
	CSCMgmt CSCManagement1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CSCMgmt"`
	CSCVal  Min3Max4NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CSCVal,omitempty"`
}

type Cardholder1 struct {
	Id        []CardholderIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Id,omitempty"`
	Nm        Max45Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Nm,omitempty"`
	Lang      ISO2ALanguageCode           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Lang,omitempty"`
	Authntcn  []CardholderAuthentication1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Authntcn,omitempty"`
	AdrVrfctn AddressVerification1        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AdrVrfctn,omitempty"`
	PrsnlData Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PrsnlData,omitempty"`
}

type CardholderAuthentication1 struct {
	AuthntcnMtd       AuthenticationMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AuthntcnMtd"`
	AuthntcnNtty      AuthenticationEntity1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AuthntcnNtty"`
	AuthntcnVal       Max40Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AuthntcnVal,omitempty"`
	CrdhldrOnLinePIN  OnLinePIN1                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CrdhldrOnLinePIN,omitempty"`
	AuthntcnColltnInd Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AuthntcnColltnInd,omitempty"`
}

type CardholderIdentification1 struct {
	CrdhldrIdVal Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CrdhldrIdVal"`
	CrdhldrIdTp  PersonIdentificationType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CrdhldrIdTp"`
}

// May be one of MNSG, NPIN, FCPN, FEPN, FDSG, FBIO, MNVR, FBIG, APKI, PKIS, CHDT, SCEC
type CardholderVerificationCapability1Code string

type CertificateIdentifier1 struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 IssrAndSrlNb"`
}

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 RltvDstngshdNm"`
}

type ContentInformationType2 struct {
	CnttTp     ContentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CnttTp"`
	EnvlpdData EnvelopedData1   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 EnvlpdData"`
}

type ContentInformationType3 struct {
	CnttTp       ContentType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CnttTp"`
	AuthntcdData []AuthenticatedData1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AuthntcdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, ECRP, AUTH
type ContentType1Code string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DetailedAmount1 struct {
	Tp  TypeOfAmount2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Tp"`
	Val float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Val"`
}

type DisplayCapabilities1 struct {
	DispTp    UserInterface2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 DispTp"`
	NbOfLines Max3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 NbOfLines"`
	LineWidth Max3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 LineWidth"`
}

type Document struct {
	AccptrAuthstnReq AcceptorAuthorisationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AccptrAuthstnReq"`
}

type EncapsulatedContent1 struct {
	CnttTp ContentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CnttTp"`
	Cntt   Max10000Binary   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Cntt,omitempty"`
}

type EncryptedContent1 struct {
	CnttTp         ContentType1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CnttNcrptnAlgo"`
	NcrptdData     Max10000Binary           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 NcrptdData"`
}

type EnvelopedData1 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Vrsn,omitempty"`
	Rcpt       []Recipient1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Rcpt"`
	NcrptdCntt EncryptedContent1  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 NcrptdCntt"`
}

// May be no more than 10 items long
type Exact10Text string

// Must match the pattern [0-9]
type Exact1NumericText string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// May be one of DAIL, MNTH, YEAR
type Frequency4Code string

type GenericIdentification31 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Tp"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 ShrtNm,omitempty"`
}

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 ShrtNm,omitempty"`
}

type GenericIdentification33 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Id,omitempty"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Tp"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 ShrtNm,omitempty"`
}

type Header1 struct {
	MsgFctn    MessageFunction1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PrtcolVrsn"`
	XchgId     Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 XchgId"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CreDtTm"`
	InitgPty   GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 InitgPty"`
	RcptPty    GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 RcptPty,omitempty"`
	Tracblt    []Traceability1         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Tracblt,omitempty"`
}

// Must match the pattern [a-z]{2,2}
type ISO2ALanguageCode string

// Must match the pattern [A-Z]{3,3}
type ISO3ACountryCode string

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

type ISOYearMonth time.Time

func (t *ISOYearMonth) UnmarshalText(text []byte) error {
	return (*xsdGYearMonth)(t).UnmarshalText(text)
}
func (t ISOYearMonth) MarshalText() ([]byte, error) {
	return xsdGYearMonth(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SrlNb"`
}

type KEK1 struct {
	Vrsn          float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Vrsn,omitempty"`
	KEKId         KEKIdentifier1           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 NcrptdKey"`
}

type KEKIdentifier1 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 KeyId"`
	KeyVrsn   Exact10Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 KeyVrsn"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 DerivtnId,omitempty"`
}

type KeyTransport1 struct {
	Vrsn          float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Vrsn"`
	RcptId        CertificateIdentifier1   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 NcrptdKey"`
}

// May be one of FIXD, ABRD, NMDC, MOTO, HOME
type LocationCategory1Code string

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
type Max16Text string

// Must match the pattern [0-9]{1,2}
type Max2NumericText string

type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{1,35}
type Max35NumericText string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must be at least 1 items long
type Max40Text string

// Must be at least 1 items long
type Max45Text string

type Max500Binary []byte

func (t *Max500Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max500Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// May be one of AUTQ, AUTP, FAUQ, FAUP, CMPV, CMPK, FCMV, FCMK, RVRA, RVRR, FRVA, FRVR, CCAQ, CCAP, CCAV, CCAK, DGNP, DGNQ, RCLQ, RCLP, RJCT
type MessageFunction1Code string

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

// Must match the pattern [0-9]{3,4}
type Min3Max4NumericText string

// Must be at least 3 items long
type Min3Max4Text string

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

// May be one of OFLN, ONLN, SMON
type OnLineCapability1Code string

type OnLinePIN1 struct {
	NcrptdPINBlck ContentInformationType2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 NcrptdPINBlck"`
	PINFrmt       PINFormat1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PINFrmt"`
	AddtlInpt     Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AddtlInpt,omitempty"`
}

// May be one of RNDM, ICCF, MERF, TRMF, ISSF, FRLT, EXFL, TAMT, CBIN, UBIN, CPAN, FLOW, CRCY
type OnLineReason1Code string

type Organisation5 struct {
	Id        GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Id,omitempty"`
	CmonNm    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CmonNm,omitempty"`
	LctnCtgy  LocationCategory1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 LctnCtgy,omitempty"`
	Adr       Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Adr,omitempty"`
	CtryCd    ISO3ACountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CtryCd,omitempty"`
	SchmeData Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SchmeData,omitempty"`
}

// May be one of ISO0, ISO1, ISO2, ISO3
type PINFormat1Code string

// May be one of SOFT, EMVK, EMVO, MRIT, CHIT, SECM, PEDV
type POIComponentType1Code string

type Parameter1 struct {
	InitlstnVctr Max500Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 InitlstnVctr,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

type PaymentCard1 struct {
	PrtctdCardData ContentInformationType2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PrtctdCardData,omitempty"`
	PlainCardData  PlainCardData1          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PlainCardData,omitempty"`
	CardCtryCd     Exact3NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CardCtryCd,omitempty"`
	CardPdctPrfl   Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CardPdctPrfl,omitempty"`
	CardBrnd       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CardBrnd,omitempty"`
	AddtlCardData  Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AddtlCardData,omitempty"`
}

type PaymentContext1 struct {
	CardPres       bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CardPres,omitempty"`
	CrdhldrPres    bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CrdhldrPres,omitempty"`
	OnLineCntxt    bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 OnLineCntxt,omitempty"`
	AttndncCntxt   AttendanceContext1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AttndncCntxt,omitempty"`
	TxEnvt         TransactionEnvironment1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TxEnvt,omitempty"`
	TxChanl        TransactionChannel1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TxChanl,omitempty"`
	AttndntMsgCpbl bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AttndntMsgCpbl,omitempty"`
	AttndntLang    ISO2ALanguageCode           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AttndntLang,omitempty"`
	CardDataNtryMd CardDataReading1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CardDataNtryMd"`
	FllbckInd      bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 FllbckInd,omitempty"`
}

// May be one of PASS, DRLC, EEID, DRVR
type PersonIdentificationType4Code string

type PlainCardData1 struct {
	PAN        Min8Max28NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PAN"`
	CardSeqNb  Min2Max3NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CardSeqNb,omitempty"`
	FctvDt     ISOYearMonth             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 FctvDt,omitempty"`
	XpryDt     ISOYearMonth             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 XpryDt"`
	SvcCd      Exact3NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SvcCd,omitempty"`
	TrckData   []TrackData1             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TrckData,omitempty"`
	CardSctyCd CardSecurityInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CardSctyCd,omitempty"`
}

type PointOfInteraction1 struct {
	Id       GenericIdentification32         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Id"`
	SysNm    Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SysNm,omitempty"`
	GrpId    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 GrpId,omitempty"`
	Cpblties PointOfInteractionCapabilities1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Cpblties,omitempty"`
	Cmpnt    []PointOfInteractionComponent1  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Cmpnt,omitempty"`
}

type PointOfInteractionCapabilities1 struct {
	CardRdngCpblties      []CardDataReading1Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CardRdngCpblties,omitempty"`
	CrdhldrVrfctnCpblties []CardholderVerificationCapability1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CrdhldrVrfctnCpblties,omitempty"`
	OnLineCpblties        OnLineCapability1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 OnLineCpblties,omitempty"`
	DispCpblties          []DisplayCapabilities1                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 DispCpblties,omitempty"`
	PrtLineWidth          Max3NumericText                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PrtLineWidth,omitempty"`
}

type PointOfInteractionComponent1 struct {
	POICmpntTp POIComponentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 POICmpntTp"`
	ManfctrId  Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 ManfctrId,omitempty"`
	Mdl        Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Mdl,omitempty"`
	VrsnNb     Max16Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 VrsnNb,omitempty"`
	SrlNb      Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SrlNb,omitempty"`
	ApprvlNb   []Max70Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 ApprvlNb,omitempty"`
}

type Product1 struct {
	PdctCd       Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PdctCd"`
	UnitOfMeasr  UnitOfMeasure1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 UnitOfMeasr,omitempty"`
	PdctQty      float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PdctQty,omitempty"`
	UnitPric     float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 UnitPric,omitempty"`
	PdctAmt      float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PdctAmt"`
	TaxTp        Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TaxTp,omitempty"`
	AddtlPdctInf Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AddtlPdctInf,omitempty"`
}

type Recipient1Choice struct {
	KeyTrnsprt KeyTransport1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 KeyTrnsprt,omitempty"`
	KEK        KEK1          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 KEK,omitempty"`
}

type RecurringTransaction1 struct {
	SeqNb       Max2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SeqNb"`
	PrdUnit     Frequency4Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 PrdUnit"`
	InstlmtPrd  float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 InstlmtPrd"`
	TtlNbOfPmts float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TtlNbOfPmts"`
	IntrstChrgs float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 IntrstChrgs,omitempty"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AttrVal"`
}

// May be one of DECL, APPR, PART, TECH
type Response1Code string

type ResponseType1 struct {
	Rspn    Response1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Rspn"`
	RspnRsn Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 RspnRsn,omitempty"`
}

type SaleContext1 struct {
	SaleId        Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SaleId,omitempty"`
	SaleRefNb     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SaleRefNb,omitempty"`
	SaleRcncltnId Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 SaleRcncltnId,omitempty"`
	CshrId        Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 CshrId,omitempty"`
	ShftNb        Max2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 ShftNb,omitempty"`
	AddtlSaleData Max70Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 AddtlSaleData,omitempty"`
}

type Traceability1 struct {
	RlayId      GenericIdentification31 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 RlayId"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TracDtTmOut"`
}

type TrackData1 struct {
	TrckNb  Exact1NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TrckNb,omitempty"`
	TrckVal Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TrckVal"`
}

// May be one of MAIL, TLPH, ECOM, TVPY
type TransactionChannel1Code string

// May be one of MERC, PRIV, PUBL
type TransactionEnvironment1Code string

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 TxRef"`
}

// May be one of ACTL, ESTM, MAXI, DFLT, RPLT
type TypeOfAmount1Code string

// May be one of CSHB, GRTY, FEES, RBTS, VATX
type TypeOfAmount2Code string

// May be one of PIEC, TONS, FOOT, GBGA, USGA, GRAM, INCH, KILO, PUND, METR, CMET, MMET, LITR, CELI, MILI, GBOU, USOU, GBQA, USQA, GBPI, USPI, MILE, KMET, YARD, SQKI, HECT, ARES, SMET, SCMT, SMIL, SQMI, SQYA, SQFO, SQIN, ACRE
type UnitOfMeasure1Code string

// May be one of MDSP, CDSP
type UserInterface2Code string

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

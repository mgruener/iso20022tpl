// Code generated by main. DO NOT EDIT.

package caaa_007_001_04

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorCancellationAdvice4 struct {
	Envt  CardPaymentEnvironment36 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Envt"`
	Cntxt CardPaymentContext11     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Cntxt,omitempty"`
	Tx    CardPaymentTransaction44 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Tx"`
}

type AcceptorCancellationAdviceV04 struct {
	Hdr      Header11                    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Hdr"`
	CxlAdvc  AcceptorCancellationAdvice4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CxlAdvc"`
	SctyTrlr ContentInformationType11    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SctyTrlr"`
}

type Acquirer4 struct {
	Id         GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Id,omitempty"`
	ParamsVrsn Max256Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 ParamsVrsn"`
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
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Param,omitempty"`
}

// May be one of ATTD, SATT, UATT
type AttendanceContext1Code string

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 MAC"`
}

type AuthorisationResult5 struct {
	AuthstnNtty   GenericIdentification70 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AuthstnNtty,omitempty"`
	RspnToAuthstn ResponseType1           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 RspnToAuthstn"`
	AuthstnCd     Min6Max8Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AuthstnCd,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// May be one of TAGC, PHYS, BRCD, MGST, CICC, DFLE, CTLS, ECTL
type CardDataReading1Code string

// May be one of FFLB, SFLB, NFLB
type CardFallback1Code string

type CardPaymentContext11 struct {
	PmtCntxt  PaymentContext11 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 PmtCntxt"`
	SaleCntxt SaleContext1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SaleCntxt,omitempty"`
}

type CardPaymentEnvironment36 struct {
	Acqrr    Acquirer4           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Acqrr,omitempty"`
	Mrchnt   Organisation8       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Mrchnt,omitempty"`
	POI      PointOfInteraction4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 POI"`
	Card     PaymentCard11       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Card"`
	CstmrDvc CustomerDevice1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CstmrDvc,omitempty"`
	Wllt     CustomerDevice1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Wllt,omitempty"`
	PmtTkn   CardPaymentToken3   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 PmtTkn,omitempty"`
}

// May be one of IRES, URES, PRES, ARES, FREC, RREC
type CardPaymentServiceType3Code string

// May be one of BALC, CACT, CRDP, CAFH, CAVR, CSHW, CSHD, DEFR, LOAD, ORCR, PINC, QUCH, RFND, RESA, VALC, UNLD, CAFT, CAFL
type CardPaymentServiceType5Code string

// May be one of AGGR, DCCV, GRTT, LOYT, NRES, PUCO, RECP, SOAF, VCAU, INSI, INSA, CSHB
type CardPaymentServiceType6Code string

type CardPaymentToken3 struct {
	TknChrtc     []Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TknChrtc,omitempty"`
	TknRqstr     PaymentTokenIdentifiers1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TknRqstr,omitempty"`
	TknAssrncLvl float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TknAssrncLvl,omitempty"`
}

type CardPaymentTransaction37 struct {
	SaleRefId Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SaleRefId,omitempty"`
	TxId      TransactionIdentifier1        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TxId"`
	POIId     GenericIdentification32       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 POIId,omitempty"`
	InitrTxId Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 InitrTxId,omitempty"`
	RcptTxId  Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 RcptTxId,omitempty"`
	TxTp      CardPaymentServiceType5Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TxTp"`
	AddtlSvc  []CardPaymentServiceType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AddtlSvc,omitempty"`
	SvcAttr   CardPaymentServiceType3Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SvcAttr,omitempty"`
	TxRslt    CardPaymentTransactionResult2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TxRslt,omitempty"`
}

type CardPaymentTransaction44 struct {
	MrchntCtgyCd Min3Max4Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 MrchntCtgyCd"`
	SaleRefId    Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SaleRefId,omitempty"`
	TxId         TransactionIdentifier1         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TxId"`
	OrgnlTx      CardPaymentTransaction37       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 OrgnlTx,omitempty"`
	TxSucss      bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TxSucss"`
	Rvsl         bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Rvsl,omitempty"`
	FailrRsn     []FailureReason3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 FailrRsn,omitempty"`
	RcptTxId     Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 RcptTxId,omitempty"`
	RcncltnId    Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 RcncltnId,omitempty"`
	IntrchngData Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 IntrchngData,omitempty"`
	TxDtls       CardPaymentTransactionDetails7 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TxDtls"`
	AuthstnRslt  AuthorisationResult5           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AuthstnRslt,omitempty"`
	AddtlTxData  []Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AddtlTxData,omitempty"`
}

type CardPaymentTransactionDetails7 struct {
	Ccy         CurrencyCode   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Ccy"`
	TtlAmt      float64        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TtlAmt"`
	VldtyDt     ISODate        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 VldtyDt,omitempty"`
	ICCRltdData Max10000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 ICCRltdData,omitempty"`
}

type CardPaymentTransactionResult2 struct {
	AuthstnNtty   GenericIdentification70 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AuthstnNtty,omitempty"`
	RspnToAuthstn ResponseType1           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 RspnToAuthstn"`
	AuthstnCd     Min6Max8Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AuthstnCd,omitempty"`
}

// May be one of COMM, CONS
type CardProductType1Code string

// May be one of MNSG, NPIN, FCPN, FEPN, FDSG, FBIO, MNVR, FBIG, APKI, PKIS, CHDT, SCEC
type CardholderVerificationCapability1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 RltvDstngshdNm"`
}

type CommunicationCharacteristics2 struct {
	ComTp   POICommunicationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 ComTp"`
	RmotPty []PartyType7Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 RmotPty"`
	Actv    bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Actv"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 EnvlpdData"`
}

type ContentInformationType11 struct {
	CnttTp       ContentType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CnttTp"`
	AuthntcdData []AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AuthntcdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type CustomerDevice1 struct {
	Id    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Id,omitempty"`
	Tp    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Tp,omitempty"`
	Prvdr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Prvdr,omitempty"`
}

type DisplayCapabilities2 struct {
	DispTp    UserInterface2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 DispTp"`
	NbOfLines float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 NbOfLines"`
	LineWidth float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 LineWidth"`
}

type Document struct {
	AccptrCxlAdvc AcceptorCancellationAdviceV04 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AccptrCxlAdvc"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 NcrptdCntt,omitempty"`
}

// Must match the pattern [0-9]
type Exact1NumericText string

// Must match the pattern [a-zA-Z0-9]{3}
type Exact3AlphaNumericText string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// May be one of CDCL, CUCL, MALF, FDCL, NDCL, PART, SFRD, TIMO, LATE, UCMP, USND, SECU
type FailureReason3Code string

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 ShrtNm,omitempty"`
}

type GenericIdentification48 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Id"`
	Vrsn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Vrsn"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Issr"`
}

type GenericIdentification53 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Tp,omitempty"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 ShrtNm,omitempty"`
}

type GenericIdentification70 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Id,omitempty"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 ShrtNm,omitempty"`
}

type GenericIdentification76 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 ShrtNm,omitempty"`
}

type Header11 struct {
	MsgFctn        MessageFunction4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 MsgFctn"`
	PrtcolVrsn     Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 PrtcolVrsn"`
	XchgId         Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 XchgId"`
	ReTrnsmssnCntr Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 ReTrnsmssnCntr,omitempty"`
	CreDtTm        ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CreDtTm"`
	InitgPty       GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 InitgPty"`
	RcptPty        GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 RcptPty,omitempty"`
	Tracblt        []Traceability2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Tracblt,omitempty"`
}

// Must match the pattern [0-9]{3,3}
type ISO3NumericCountryCode string

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
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 NcrptdKey"`
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

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Must be at least 1 items long
type Max256Text string

// Must match the pattern [0-9]{1,2}
type Max2NumericText string

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

// Must be at least 1 items long
type Max3Text string

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

type MemoryCharacteristics1 struct {
	Id     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Id"`
	TtlSz  float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TtlSz"`
	FreeSz float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 FreeSz"`
	Unit   MemoryUnit1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Unit"`
}

// May be one of BYTE, EXAB, GIGA, KILO, MEGA, PETA, TERA
type MemoryUnit1Code string

// May be one of AUTQ, AUTP, FAUQ, FAUP, CMPV, CMPK, FCMV, FCMK, RVRA, RVRR, FRVA, FRVR, CCAQ, CCAP, CCAV, CCAK, DGNP, DGNQ, RCLQ, RCLP, RJCT, DCCQ, DCCP
type MessageFunction4Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

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

type Organisation8 struct {
	Id        GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Id,omitempty"`
	CmonNm    Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CmonNm,omitempty"`
	LctnCtgy  LocationCategory1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 LctnCtgy,omitempty"`
	Adr       Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Adr,omitempty"`
	CtryCd    ISO3NumericCountryCode  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CtryCd,omitempty"`
	SchmeData Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SchmeData,omitempty"`
}

// May be one of BLTH, ETHR, GPRS, GSMF, PSTN, RS23, USBD, USBH, WIFI
type POICommunicationType1Code string

// May be one of APPL, CERT, EVAL
type POIComponentAssessment1Code string

// May be one of WAIT, OUTD, OPER, DACT
type POIComponentStatus1Code string

// May be one of AQPP, APPR, TLPR, SCPR, SERV, TERM, DVCE, SECM, APLI, EMVK, EMVO, MDWR, DRVR, OPST, MRPR
type POIComponentType3Code string

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 BPddg,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

// May be one of ACQR, ITAG, PCPT, TMGT, SALE
type PartyType7Code string

type PaymentCard11 struct {
	PrtctdCardData ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 PrtctdCardData,omitempty"`
	PlainCardData  PlainCardData9           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 PlainCardData,omitempty"`
	MskdPAN        string                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 MskdPAN,omitempty"`
	IssrBIN        Max15NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 IssrBIN,omitempty"`
	CardCtryCd     Max3Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CardCtryCd,omitempty"`
	CardCcyCd      Exact3AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CardCcyCd,omitempty"`
	CardPdctPrfl   Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CardPdctPrfl,omitempty"`
	CardBrnd       Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CardBrnd,omitempty"`
	CardPdctTp     CardProductType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CardPdctTp,omitempty"`
	AddtlCardData  Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AddtlCardData,omitempty"`
}

type PaymentContext11 struct {
	CardPres       bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CardPres,omitempty"`
	CrdhldrPres    bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CrdhldrPres,omitempty"`
	OnLineCntxt    bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 OnLineCntxt,omitempty"`
	AttndncCntxt   AttendanceContext1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AttndncCntxt,omitempty"`
	TxEnvt         TransactionEnvironment1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TxEnvt,omitempty"`
	TxChanl        TransactionChannel3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TxChanl,omitempty"`
	CardDataNtryMd CardDataReading1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CardDataNtryMd"`
	FllbckInd      CardFallback1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 FllbckInd,omitempty"`
}

type PaymentTokenIdentifiers1 struct {
	PrvdrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 PrvdrId"`
	RqstrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 RqstrId"`
}

type PlainCardData9 struct {
	PAN       Min8Max28NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 PAN"`
	CardSeqNb Min2Max3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CardSeqNb,omitempty"`
	FctvDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 FctvDt,omitempty"`
	XpryDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 XpryDt"`
	SvcCd     Exact3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SvcCd,omitempty"`
	TrckData  []TrackData1         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TrckData,omitempty"`
}

type PointOfInteraction4 struct {
	Id       GenericIdentification32         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Id"`
	SysNm    Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SysNm,omitempty"`
	GrpId    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 GrpId,omitempty"`
	Cpblties PointOfInteractionCapabilities3 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Cpblties,omitempty"`
	Cmpnt    []PointOfInteractionComponent5  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Cmpnt,omitempty"`
}

type PointOfInteractionCapabilities3 struct {
	CardRdngCpblties      []CardDataReading1Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CardRdngCpblties,omitempty"`
	CrdhldrVrfctnCpblties []CardholderVerificationCapability1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CrdhldrVrfctnCpblties,omitempty"`
	PINLngthCpblties      float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 PINLngthCpblties,omitempty"`
	ApprvlCdLngth         float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 ApprvlCdLngth,omitempty"`
	CardCaptrCpbl         bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CardCaptrCpbl,omitempty"`
	OnLineCpblties        OnLineCapability1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 OnLineCpblties,omitempty"`
	DispCpblties          []DisplayCapabilities2                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 DispCpblties,omitempty"`
	PrtLineWidth          float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 PrtLineWidth,omitempty"`
	AvlblLang             []string                                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AvlblLang,omitempty"`
}

type PointOfInteractionComponent5 struct {
	Tp       POIComponentType3Code                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Tp"`
	Id       PointOfInteractionComponentIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Id"`
	Sts      PointOfInteractionComponentStatus2          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Sts,omitempty"`
	StdCmplc []GenericIdentification48                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 StdCmplc,omitempty"`
	Chrtcs   PointOfInteractionComponentCharacteristics2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Chrtcs,omitempty"`
	Assmnt   []PointOfInteractionComponentAssessment1    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Assmnt,omitempty"`
}

type PointOfInteractionComponentAssessment1 struct {
	Tp      POIComponentAssessment1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Tp"`
	Assgnr  []Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Assgnr"`
	DlvryDt ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 DlvryDt,omitempty"`
	XprtnDt ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 XprtnDt,omitempty"`
	Nb      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Nb"`
}

type PointOfInteractionComponentCharacteristics2 struct {
	Mmry           []MemoryCharacteristics1        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Mmry,omitempty"`
	Com            []CommunicationCharacteristics2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Com,omitempty"`
	SctyAccsMdls   float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SctyAccsMdls,omitempty"`
	SbcbrIdntyMdls float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SbcbrIdntyMdls,omitempty"`
	KeyChckVal     Max35Binary                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 KeyChckVal,omitempty"`
}

type PointOfInteractionComponentIdentification1 struct {
	ItmNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 ItmNb,omitempty"`
	PrvdrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 PrvdrId,omitempty"`
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Id,omitempty"`
	SrlNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SrlNb,omitempty"`
}

type PointOfInteractionComponentStatus2 struct {
	VrsnNb Max256Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 VrsnNb,omitempty"`
	Sts    POIComponentStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Sts,omitempty"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 KeyIdr"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AttrVal"`
}

// May be one of DECL, APPR, PART, TECH
type Response1Code string

type ResponseType1 struct {
	Rspn    Response1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Rspn"`
	RspnRsn Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 RspnRsn,omitempty"`
}

type SaleContext1 struct {
	SaleId        Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SaleId,omitempty"`
	SaleRefNb     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SaleRefNb,omitempty"`
	SaleRcncltnId Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 SaleRcncltnId,omitempty"`
	CshrId        Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 CshrId,omitempty"`
	ShftNb        Max2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 ShftNb,omitempty"`
	AddtlSaleData Max70Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 AddtlSaleData,omitempty"`
}

type Traceability2 struct {
	RlayId      GenericIdentification76 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 RlayId"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TracDtTmOut"`
}

type TrackData1 struct {
	TrckNb  Exact1NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TrckNb,omitempty"`
	TrckVal Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TrckVal"`
}

// May be one of MAIL, TLPH, ECOM, TVPY, SECM, MOBL
type TransactionChannel3Code string

// May be one of MERC, PRIV, PUBL
type TransactionEnvironment1Code string

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 TxRef"`
}

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

// Code generated by main. DO NOT EDIT.

package catm_001_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// May be one of EA2C, E3DC, DKPT, DKP9, UKPT, UKA1
type Algorithm2Code string

// May be one of MACC, MCCS, CMA1, CMD1
type Algorithm3Code string

// May be one of ERS2
type Algorithm4Code string

// May be one of HS25, HS38, HS51
type Algorithm5Code string

// May be one of EA2C, E3DC
type Algorithm6Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification2 struct {
	Algo  Algorithm2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Algo"`
	Param Parameter1     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Param,omitempty"`
}

type AlgorithmIdentification3 struct {
	Algo  Algorithm3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Algo"`
	Param Parameter1     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Param,omitempty"`
}

type AlgorithmIdentification4 struct {
	Algo Algorithm4Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Algo"`
}

type AlgorithmIdentification5 struct {
	Algo Algorithm5Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Algo"`
}

type AlgorithmIdentification6 struct {
	Algo  Algorithm6Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Algo"`
	Param Parameter1     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Param,omitempty"`
}

type AlgorithmIdentification7 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Algo"`
	Param Parameter2     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Param,omitempty"`
}

type AlgorithmIdentification8 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Algo"`
	Param Parameter3     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Param,omitempty"`
}

// May be one of ATTD, SATT, UATT
type AttendanceContext1Code string

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData2 struct {
	Vrsn        float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Vrsn,omitempty"`
	Rcpt        []Recipient2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Rcpt"`
	MACAlgo     AlgorithmIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 MACAlgo"`
	NcpsltdCntt EncapsulatedContent1     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 NcpsltdCntt"`
	MAC         Max35Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 MAC"`
}

// May be one of TAGC, PHYS, BRCD, MGST, CICC, DFLE, CTLS, ECTL
type CardDataReading1Code string

// May be one of MNSG, NPIN, FCPN, FEPN, FDSG, FBIO, MNVR, FBIG, APKI, PKIS, CHDT, SCEC
type CardholderVerificationCapability1Code string

type CertificateIdentifier1 struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 IssrAndSrlNb"`
}

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 RltvDstngshdNm"`
}

type CommunicationCharacteristics1 struct {
	ComTp   POICommunicationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 ComTp"`
	RmotPty PartyType7Code            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 RmotPty"`
	Actv    bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Actv"`
}

type ContentInformationType4 struct {
	CnttTp           ContentType1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 CnttTp"`
	EnvlpdData       EnvelopedData2         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 EnvlpdData,omitempty"`
	AuthntcdData     AuthenticatedData2     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 AuthntcdData,omitempty"`
	SgndData         SignedData2            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 SgndData,omitempty"`
	DgstdData        DigestedData2          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DgstdData,omitempty"`
	NmdKeyNcrptdData NamedKeyEncryptedData2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 NmdKeyNcrptdData,omitempty"`
}

type ContentInformationType5 struct {
	CnttTp     ContentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 CnttTp"`
	EnvlpdData EnvelopedData2   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 EnvlpdData"`
}

// May be one of DATA, SIGN, EVLP, DGST, ECRP, AUTH
type ContentType1Code string

// May be one of AQPR, APPR, TXCP, AKCP, DLGT, MGTP, MRPR, SCPR, SWPK, STRP, TRPR, VDPR
type DataSetCategory3Code string

type DataSetIdentification3 struct {
	Nm      Max256Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Nm,omitempty"`
	Tp      DataSetCategory3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Tp"`
	Vrsn    Max256Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Vrsn,omitempty"`
	CreDtTm ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 CreDtTm,omitempty"`
}

type DigestedData2 struct {
	Vrsn        float64                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent1       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 NcpsltdCntt"`
	Dgst        Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Dgst"`
}

type DisplayCapabilities1 struct {
	DispTp    UserInterface2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DispTp"`
	NbOfLines Max3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 NbOfLines"`
	LineWidth Max3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 LineWidth"`
}

type Document struct {
	StsRpt StatusReportV02 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 StsRpt"`
}

type EncapsulatedContent1 struct {
	CnttTp ContentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 CnttTp"`
	Cntt   Max10000Binary   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Cntt,omitempty"`
}

type EncryptedContent2 struct {
	CnttTp         ContentType1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 CnttNcrptnAlgo"`
	NcrptdData     Max10000Binary           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 NcrptdData"`
}

type EnvelopedData2 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Vrsn,omitempty"`
	Rcpt       []Recipient2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Rcpt"`
	NcrptdCntt EncryptedContent2  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 NcrptdCntt"`
}

// May be no more than 10 items long
type Exact10Text string

type GenericIdentification35 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Id"`
	Tp     PartyType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Tp,omitempty"`
	Issr   PartyType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 ShrtNm,omitempty"`
}

type GenericIdentification48 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Id"`
	Vrsn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Vrsn"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Issr"`
}

type Header4 struct {
	DwnldTrf bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DwnldTrf"`
	FrmtVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 FrmtVrsn"`
	XchgId   Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 XchgId"`
	CreDtTm  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 CreDtTm"`
	InitgPty GenericIdentification35 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 InitgPty"`
	RcptPty  GenericIdentification35 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 RcptPty,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 SrlNb"`
}

type KEK2 struct {
	Vrsn          float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Vrsn,omitempty"`
	KEKId         KEKIdentifier1           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 NcrptdKey"`
}

type KEKIdentifier1 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 KeyId"`
	KeyVrsn   Exact10Text     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 KeyVrsn"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DerivtnId,omitempty"`
}

type KeyTransport2 struct {
	Vrsn          float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Vrsn"`
	RcptId        CertificateIdentifier1   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 NcrptdKey"`
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

// Must match the pattern [0-9]{1,9}
type Max9NumericText string

type MemoryCharacteristics1 struct {
	Id     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Id"`
	TtlSz  float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 TtlSz"`
	FreeSz float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 FreeSz"`
	Unit   MemoryUnit1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Unit"`
}

// May be one of BYTE, EXAB, GIGA, KILO, MEGA, PETA, TERA
type MemoryUnit1Code string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type NamedKeyEncryptedData2 struct {
	Vrsn       float64           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Vrsn,omitempty"`
	KeyNm      Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 KeyNm,omitempty"`
	NcrptdCntt EncryptedContent2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 NcrptdCntt"`
}

// May be one of OFLN, ONLN, SMON
type OnLineCapability1Code string

// May be one of BLTH, ETHR, GPRS, GSMF, PSTN, RS23, USBD, USBH, WIFI
type POICommunicationType1Code string

// May be one of APPL, CERT, EVAL
type POIComponentAssessment1Code string

// May be one of WAIT, OUTD, OPER, DACT
type POIComponentStatus1Code string

// May be one of AQPP, APPR, TLPR, SCPR, SERV, TERM, DVCE, SECM, APLI, EMVK, EMVO, MDWR, DRVR, OPST, MRPR
type POIComponentType3Code string

type Parameter1 struct {
	InitlstnVctr Max500Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 InitlstnVctr,omitempty"`
}

type Parameter2 struct {
	DgstAlgo     Algorithm5Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 MskGnrtrAlgo,omitempty"`
}

type Parameter3 struct {
	DgstAlgo Algorithm5Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DgstAlgo,omitempty"`
}

// May be one of OPOI, ACCP, MERC, ACQR, ITAG, MTMG, TMGT
type PartyType5Code string

// May be one of ACCP, MERC, ACQR, ITAG, MTMG, TMGT
type PartyType6Code string

// May be one of ACQR, ITAG, PCPT, TMGT, SALE
type PartyType7Code string

type PointOfInteractionCapabilities1 struct {
	CardRdngCpblties      []CardDataReading1Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 CardRdngCpblties,omitempty"`
	CrdhldrVrfctnCpblties []CardholderVerificationCapability1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 CrdhldrVrfctnCpblties,omitempty"`
	OnLineCpblties        OnLineCapability1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 OnLineCpblties,omitempty"`
	DispCpblties          []DisplayCapabilities1                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DispCpblties,omitempty"`
	PrtLineWidth          Max3NumericText                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 PrtLineWidth,omitempty"`
}

type PointOfInteractionComponent3 struct {
	Tp       POIComponentType3Code                       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Tp"`
	Id       PointOfInteractionComponentIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Id"`
	Sts      PointOfInteractionComponentStatus1          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Sts,omitempty"`
	StdCmplc []GenericIdentification48                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 StdCmplc,omitempty"`
	Chrtcs   PointOfInteractionComponentCharacteristics1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Chrtcs,omitempty"`
	Assmnt   []PointOfInteractionComponentAssessment1    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Assmnt,omitempty"`
}

type PointOfInteractionComponentAssessment1 struct {
	Tp      POIComponentAssessment1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Tp"`
	Assgnr  []Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Assgnr"`
	DlvryDt ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DlvryDt,omitempty"`
	XprtnDt ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 XprtnDt,omitempty"`
	Nb      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Nb"`
}

type PointOfInteractionComponentCharacteristics1 struct {
	Mmry           []MemoryCharacteristics1        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Mmry,omitempty"`
	Com            []CommunicationCharacteristics1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Com,omitempty"`
	SctyAccsMdls   float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 SctyAccsMdls,omitempty"`
	SbcbrIdntyMdls float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 SbcbrIdntyMdls,omitempty"`
	KeyChckVal     Max35Binary                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 KeyChckVal,omitempty"`
}

type PointOfInteractionComponentIdentification1 struct {
	ItmNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 ItmNb,omitempty"`
	PrvdrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 PrvdrId,omitempty"`
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Id,omitempty"`
	SrlNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 SrlNb,omitempty"`
}

type PointOfInteractionComponentStatus1 struct {
	VrsnNb Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 VrsnNb,omitempty"`
	Sts    POIComponentStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Sts,omitempty"`
}

type Recipient2Choice struct {
	KeyTrnsprt KeyTransport2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 KeyTrnsprt,omitempty"`
	KEK        KEK2          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 KEK,omitempty"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 AttrVal"`
}

type SignedData2 struct {
	Vrsn        float64                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent1       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 NcpsltdCntt"`
	Cert        []Max3000Binary            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Cert,omitempty"`
	Sgnr        []Signer2                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Sgnr"`
}

type Signer2 struct {
	Vrsn      float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Vrsn,omitempty"`
	SgnrId    CertificateIdentifier1   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 SgnrId"`
	DgstAlgo  AlgorithmIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DgstAlgo"`
	SgntrAlgo AlgorithmIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 SgntrAlgo"`
	Sgntr     Max500Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Sgntr"`
}

type StatusReport2 struct {
	POIId       GenericIdentification35      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 POIId"`
	TermnlMgrId GenericIdentification35      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 TermnlMgrId,omitempty"`
	DataSet     []TerminalManagementDataSet4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DataSet"`
}

type StatusReportContent2 struct {
	POICpblties  PointOfInteractionCapabilities1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 POICpblties,omitempty"`
	POICmpnt     []PointOfInteractionComponent3  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 POICmpnt,omitempty"`
	AttndncCntxt AttendanceContext1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 AttndncCntxt,omitempty"`
	POIDtTm      ISODateTime                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 POIDtTm"`
	DataSetReqrd TerminalManagementDataSet7      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DataSetReqrd,omitempty"`
	Evt          []TMSEvent2                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Evt,omitempty"`
	Errs         Max140Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Errs,omitempty"`
}

type StatusReportV02 struct {
	Hdr      Header4                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Hdr"`
	StsRpt   StatusReport2           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 StsRpt"`
	SctyTrlr ContentInformationType4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 SctyTrlr"`
}

type TMSActionIdentification2 struct {
	ActnTp    TerminalManagementAction1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 ActnTp"`
	DataSetId DataSetIdentification3        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 DataSetId,omitempty"`
}

type TMSEvent2 struct {
	TmStmp      ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 TmStmp"`
	Rslt        TerminalManagementActionResult1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Rslt"`
	ActnId      TMSActionIdentification2            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 ActnId"`
	AddtlErrInf Max70Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 AddtlErrInf,omitempty"`
}

// May be one of ACTV, DCTV, DELT, DWNL, INST, RSTR, UPLD
type TerminalManagementAction1Code string

// May be one of ACCD, CNTE, FMTE, INVC, LENE, OVER, MISS, NSUP, SIGE, SUCC, SYNE, TIMO, UKDT, UKRF
type TerminalManagementActionResult1Code string

type TerminalManagementDataSet4 struct {
	Id      DataSetIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Id"`
	SeqCntr Max9NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 SeqCntr,omitempty"`
	Cntt    StatusReportContent2   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Cntt"`
}

type TerminalManagementDataSet7 struct {
	Id        DataSetIdentification3  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Id"`
	POIChllng Max140Binary            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 POIChllng,omitempty"`
	TMChllng  Max140Binary            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 TMChllng,omitempty"`
	NcrptdKey ContentInformationType5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 NcrptdKey,omitempty"`
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

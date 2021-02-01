// Code generated by main. DO NOT EDIT.

package catm_001_001_09

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of HS25, HS38, HS51, HS01, SH31, SH32, SH33, SH35, SHK1, SHK2
type Algorithm16Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5, CMA2, CM31, CM32, CM33, MCS3, CCA1, CCA2, CCA3
type Algorithm17Code string

// May be one of ERS2, ERS1, RPSS, ECC5, ECC1, ECC4, ECC2, ECC3, ERS3, ECP2, ECP3, ECP5
type Algorithm19Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA2, EA9C, EA5C, DA12, DA19, DA25, N108, EA5R, EA9R, EA2R, E3DR, E36C, E36R, SD5C, UKA1, UKA3
type Algorithm24Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Param,omitempty"`
}

type AlgorithmIdentification18 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Algo"`
	Param Parameter9     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Param,omitempty"`
}

type AlgorithmIdentification19 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Algo"`
	Param Parameter10    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Param,omitempty"`
}

type AlgorithmIdentification20 struct {
	Algo  Algorithm19Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Algo"`
	Param Parameter11     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Param,omitempty"`
}

type AlgorithmIdentification21 struct {
	Algo Algorithm16Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Algo"`
}

type AlgorithmIdentification22 struct {
	Algo  Algorithm17Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Param,omitempty"`
}

type AlgorithmIdentification29 struct {
	Algo  Algorithm24Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Algo"`
	Param Parameter12     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Param,omitempty"`
}

// May be one of ATTD, SATT, UATT
type AttendanceContext1Code string

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData6 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Vrsn,omitempty"`
	Rcpt        []Recipient8Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Rcpt"`
	MACAlgo     AlgorithmIdentification22 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 MAC"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// May be one of TAGC, PHYS, BRCD, MGST, CICC, DFLE, CTLS, ECTL, CDFL, SICC, UNKW, QRCD, OPTC
type CardDataReading8Code string

// May be one of APKI, CHDT, MNSG, MNVR, FBIG, FBIO, FDSG, FCPN, FEPN, NPIN, PKIS, SCEC, NBIO, NOVF, OTHR
type CardholderVerificationCapability4Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 RltvDstngshdNm"`
}

type CommunicationCharacteristics5 struct {
	ComTp      POICommunicationType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 ComTp"`
	RmotPty    []PartyType7Code            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 RmotPty"`
	Actv       bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Actv"`
	Params     NetworkParameters7          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Params,omitempty"`
	PhysIntrfc PhysicalInterfaceParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 PhysIntrfc,omitempty"`
}

type ContentInformationType21 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 CnttTp"`
	AuthntcdData AuthenticatedData6 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AuthntcdData,omitempty"`
	SgndData     SignedData5        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SgndData,omitempty"`
}

type ContentInformationType23 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 CnttTp"`
	EnvlpdData   EnvelopedData7     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 EnvlpdData,omitempty"`
	AuthntcdData AuthenticatedData6 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AuthntcdData,omitempty"`
	SgndData     SignedData5        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SgndData,omitempty"`
	DgstdData    DigestedData5      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DgstdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

type CryptographicKey14 struct {
	Id           Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Id"`
	AddtlId      Max35Binary               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AddtlId,omitempty"`
	Nm           Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Nm,omitempty"`
	SctyPrfl     Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SctyPrfl,omitempty"`
	ItmNb        Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 ItmNb,omitempty"`
	Vrsn         Max256Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Vrsn"`
	Tp           CryptographicKeyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Tp,omitempty"`
	Fctn         []KeyUsage1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Fctn,omitempty"`
	ActvtnDt     ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 ActvtnDt,omitempty"`
	DeactvtnDt   ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DeactvtnDt,omitempty"`
	KeyVal       ContentInformationType23  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 KeyVal,omitempty"`
	KeyChckVal   Max35Binary               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 KeyChckVal,omitempty"`
	AddtlMgmtInf []GenericInformation1     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AddtlMgmtInf,omitempty"`
}

// May be one of AES2, EDE3, DKP9, AES9, AES5, EDE4
type CryptographicKeyType3Code string

// May be one of AQPR, APPR, TXCP, AKCP, DLGT, MGTP, MRPR, SCPR, SWPK, STRP, TRPR, VDPR, PARA, TMSP, CRTF, LOGF, CMRQ, MDFL, SOFT, CONF, RPFL
type DataSetCategory14Code string

type DataSetIdentification8 struct {
	Nm      Max256Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Nm,omitempty"`
	Tp      DataSetCategory14Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Tp"`
	Vrsn    Max256Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Vrsn,omitempty"`
	CreDtTm ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 CreDtTm,omitempty"`
}

type DataSetRequest1 struct {
	Id               DataSetIdentification8   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Id"`
	POIChllng        Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 POIChllng,omitempty"`
	TMChllng         Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 TMChllng,omitempty"`
	SsnKey           CryptographicKey14       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SsnKey,omitempty"`
	DlgtnProof       Max5000Binary            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DlgtnProof,omitempty"`
	PrtctdDlgtnProof ContentInformationType23 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 PrtctdDlgtnProof,omitempty"`
}

type DigestedData5 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Vrsn,omitempty"`
	DgstAlgo    AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 NcpsltdCntt"`
	Dgst        Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Dgst"`
}

type DisplayCapabilities4 struct {
	Dstn      []UserInterface4Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Dstn"`
	AvlblFrmt []OutputFormat1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AvlblFrmt,omitempty"`
	NbOfLines float64              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 NbOfLines,omitempty"`
	LineWidth float64              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 LineWidth,omitempty"`
	AvlblLang []string             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AvlblLang,omitempty"`
}

type Document struct {
	StsRpt StatusReportV09 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 StsRpt"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Cntt,omitempty"`
}

type EncryptedContent6 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 CnttNcrptnAlgo,omitempty"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 NcrptdData"`
}

// May be one of TR31, TR34, I238
type EncryptionFormat2Code string

type EnvelopedData7 struct {
	Vrsn       float64                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Vrsn,omitempty"`
	OrgtrInf   OriginatorInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 OrgtrInf,omitempty"`
	Rcpt       []Recipient8Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Rcpt"`
	NcrptdCntt EncryptedContent6      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 NcrptdCntt,omitempty"`
}

// May be one of ONDM, IMMD, ASAP, AGRP, NBLT, TTLT, CYCL, NONE, BLCK
type ExchangePolicy2Code string

type ExternallyDefinedData1 struct {
	Id        Max1025Text              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Id"`
	Val       Max100KBinary            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Val,omitempty"`
	PrtctdVal ContentInformationType23 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 PrtctdVal,omitempty"`
	Tp        Max1025Text              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Tp,omitempty"`
}

type GenericIdentification176 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Id"`
	Tp     PartyType33Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Tp,omitempty"`
	Issr   PartyType33Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 ShrtNm,omitempty"`
}

type GenericIdentification177 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Id"`
	Tp       PartyType33Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Tp,omitempty"`
	Issr     PartyType33Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 ShrtNm,omitempty"`
	RmotAccs NetworkParameters7 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 RmotAccs,omitempty"`
	Glctn    Geolocation1       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Glctn,omitempty"`
}

type GenericIdentification48 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Id"`
	Vrsn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Vrsn"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Issr"`
}

type GenericInformation1 struct {
	Nm  Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Nm"`
	Val Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Val,omitempty"`
}

type Geolocation1 struct {
	GeogcCordints GeolocationGeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 GeogcCordints,omitempty"`
	UTMCordints   GeolocationUTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 UTMCordints,omitempty"`
}

type GeolocationGeographicCoordinates1 struct {
	Lat  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Lat"`
	Long Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Long"`
}

type GeolocationUTMCoordinates1 struct {
	UTMZone    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 UTMZone"`
	UTMEstwrd  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 UTMEstwrd"`
	UTMNrthwrd Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 UTMNrthwrd"`
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
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SrlNb"`
}

type KEK7 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DerivtnId,omitempty"`
}

type KeyTransport5 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 NcrptdKey"`
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

// Must be at least 1 items long
type Max1025Text string

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
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

type Max2KBinary []byte

func (t *Max2KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max2KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

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
type Max500Text string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// Must match the pattern [0-9]{1,9}
type Max9NumericText string

type MemoryCharacteristics1 struct {
	Id     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Id"`
	TtlSz  float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 TtlSz"`
	FreeSz float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 FreeSz"`
	Unit   MemoryUnit1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Unit"`
}

// May be one of BYTE, EXAB, GIGA, KILO, MEGA, PETA, TERA
type MemoryUnit1Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type NetworkParameters7 struct {
	Adr        []NetworkParameters9 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SctyPrfl,omitempty"`
}

type NetworkParameters9 struct {
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 NtwkTp"`
	AdrVal Max500Text       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AdrVal"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

// May be one of OFLN, ONLN, SMON
type OnLineCapability1Code string

type OriginatorInformation1 struct {
	Cert []Max5000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Cert,omitempty"`
}

// May be one of MREF, TEXT, HTML
type OutputFormat1Code string

// May be one of BLTH, ETHR, GPRS, GSMF, PSTN, RS23, USBD, USBH, WIFI, WT2G, WT3G, WT4G, WT5G
type POICommunicationType2Code string

// May be one of APPL, CERT, EVAL
type POIComponentAssessment1Code string

// May be one of WAIT, OUTD, OPER, DACT
type POIComponentStatus1Code string

// May be one of AQPP, APPR, TLPR, SCPR, SERV, TERM, DVCE, SECM, APLI, EMVK, EMVO, MDWR, DRVR, OPST, MRPR, CRTF, TMSP, SACP, SAPR, LOGF, MDFL, SOFT, CONF, RPFL
type POIComponentType6Code string

type PackageType1 struct {
	PackgId     GenericIdentification176 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 PackgId,omitempty"`
	PackgLngth  float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 PackgLngth,omitempty"`
	OffsetStart float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 OffsetStart,omitempty"`
	OffsetEnd   float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 OffsetEnd,omitempty"`
	PackgBlck   []ExternallyDefinedData1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 PackgBlck,omitempty"`
}

type Parameter10 struct {
	NcrptnFrmt   EncryptionFormat2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 MskGnrtrAlgo,omitempty"`
}

type Parameter11 struct {
	DgstAlgo     Algorithm16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DgstAlgo"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 MskGnrtrAlgo"`
	SaltLngth    float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SaltLngth"`
	TrlrFld      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 TrlrFld,omitempty"`
}

type Parameter12 struct {
	NcrptnFrmt   EncryptionFormat2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 BPddg,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DgstAlgo,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 BPddg,omitempty"`
}

type Parameter9 struct {
	DgstAlgo Algorithm16Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DgstAlgo,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS, MTMG, TAXH, TMGT
type PartyType33Code string

// May be one of OPOI, ACCP, MERC, ACQR, ITAG, MTMG, TMGT
type PartyType5Code string

// May be one of ACQR, ITAG, PCPT, TMGT, SALE
type PartyType7Code string

type PhysicalInterfaceParameter1 struct {
	IntrfcNm    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 IntrfcNm"`
	IntrfcTp    POICommunicationType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 IntrfcTp,omitempty"`
	UsrNm       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 UsrNm,omitempty"`
	AccsCd      Max35Binary               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AccsCd,omitempty"`
	SctyPrfl    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SctyPrfl,omitempty"`
	AddtlParams Max2KBinary               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AddtlParams,omitempty"`
}

type PointOfInteractionCapabilities9 struct {
	CardRdngCpblties      []CardDataReading8Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 CardRdngCpblties,omitempty"`
	CrdhldrVrfctnCpblties []CardholderVerificationCapability4Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 CrdhldrVrfctnCpblties,omitempty"`
	PINLngthCpblties      float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 PINLngthCpblties,omitempty"`
	ApprvlCdLngth         float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 ApprvlCdLngth,omitempty"`
	MxScrptLngth          float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 MxScrptLngth,omitempty"`
	CardCaptrCpbl         bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 CardCaptrCpbl,omitempty"`
	OnLineCpblties        OnLineCapability1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 OnLineCpblties,omitempty"`
	MsgCpblties           []DisplayCapabilities4                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 MsgCpblties,omitempty"`
}

type PointOfInteractionComponent10 struct {
	Tp       POIComponentType6Code                       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Tp"`
	SubTpInf Max70Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SubTpInf,omitempty"`
	Id       PointOfInteractionComponentIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Id"`
	Sts      PointOfInteractionComponentStatus3          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Sts,omitempty"`
	StdCmplc []GenericIdentification48                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 StdCmplc,omitempty"`
	Chrtcs   PointOfInteractionComponentCharacteristics6 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Chrtcs,omitempty"`
	Assmnt   []PointOfInteractionComponentAssessment1    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Assmnt,omitempty"`
	Packg    []PackageType1                              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Packg,omitempty"`
}

type PointOfInteractionComponentAssessment1 struct {
	Tp      POIComponentAssessment1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Tp"`
	Assgnr  []Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Assgnr"`
	DlvryDt ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DlvryDt,omitempty"`
	XprtnDt ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 XprtnDt,omitempty"`
	Nb      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Nb"`
}

type PointOfInteractionComponentCharacteristics6 struct {
	Mmry           []MemoryCharacteristics1        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Mmry,omitempty"`
	Com            []CommunicationCharacteristics5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Com,omitempty"`
	SctyAccsMdls   float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SctyAccsMdls,omitempty"`
	SbcbrIdntyMdls float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SbcbrIdntyMdls,omitempty"`
	SctyElmt       []CryptographicKey14            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SctyElmt,omitempty"`
}

type PointOfInteractionComponentIdentification1 struct {
	ItmNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 ItmNb,omitempty"`
	PrvdrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 PrvdrId,omitempty"`
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Id,omitempty"`
	SrlNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SrlNb,omitempty"`
}

type PointOfInteractionComponentStatus3 struct {
	VrsnNb Max256Text              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 VrsnNb,omitempty"`
	Sts    POIComponentStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Sts,omitempty"`
	XpryDt ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 XpryDt,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 KeyIdr"`
}

type Recipient8Choice struct {
	KeyTrnsprt KeyTransport5  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 KeyTrnsprt"`
	KEK        KEK7           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AttrVal"`
}

type SignedData5 struct {
	Vrsn        float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DgstAlgo,omitempty"`
	NcpsltdCntt EncapsulatedContent3        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 NcpsltdCntt,omitempty"`
	Cert        []Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Cert,omitempty"`
	Sgnr        []Signer4                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Sgnr,omitempty"`
}

type Signer4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Vrsn,omitempty"`
	SgnrId      Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SgnrId,omitempty"`
	DgstAlgo    AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DgstAlgo"`
	SgndAttrbts []GenericInformation1     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SgndAttrbts,omitempty"`
	SgntrAlgo   AlgorithmIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SgntrAlgo"`
	Sgntr       Max3000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Sgntr"`
}

type StatusReport9 struct {
	POIId       GenericIdentification176    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 POIId"`
	InitgTrggr  TriggerInformation2         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 InitgTrggr,omitempty"`
	TermnlMgrId GenericIdentification176    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 TermnlMgrId"`
	DataSet     StatusReportDataSetRequest1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DataSet"`
}

type StatusReportContent9 struct {
	POICpblties  PointOfInteractionCapabilities9 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 POICpblties,omitempty"`
	POICmpnt     []PointOfInteractionComponent10 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 POICmpnt,omitempty"`
	AttndncCntxt AttendanceContext1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AttndncCntxt,omitempty"`
	POIDtTm      ISODateTime                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 POIDtTm"`
	DataSetReqrd []DataSetRequest1               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DataSetReqrd,omitempty"`
	Evt          []TMSEvent7                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Evt,omitempty"`
	Errs         []Max140Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Errs,omitempty"`
}

type StatusReportDataSetRequest1 struct {
	Id      DataSetIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Id"`
	SeqCntr Max9NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SeqCntr,omitempty"`
	Cntt    StatusReportContent9   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Cntt"`
}

type StatusReportV09 struct {
	Hdr      TMSHeader1               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Hdr"`
	StsRpt   StatusReport9            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 StsRpt"`
	SctyTrlr ContentInformationType21 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SctyTrlr,omitempty"`
}

type TMSActionIdentification6 struct {
	ActnTp    TerminalManagementAction4Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 ActnTp"`
	DataSetId DataSetIdentification8        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DataSetId,omitempty"`
}

type TMSEvent7 struct {
	TmStmp      ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 TmStmp"`
	Rslt        TerminalManagementActionResult4Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Rslt"`
	ActnId      TMSActionIdentification6            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 ActnId"`
	AddtlErrInf Max70Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AddtlErrInf,omitempty"`
	TermnlMgrId Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 TermnlMgrId,omitempty"`
}

type TMSHeader1 struct {
	DwnldTrf bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 DwnldTrf"`
	FrmtVrsn Max6Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 FrmtVrsn"`
	XchgId   float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 XchgId"`
	CreDtTm  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 CreDtTm"`
	InitgPty GenericIdentification176 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 InitgPty"`
	RcptPty  GenericIdentification177 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 RcptPty,omitempty"`
	Tracblt  []Traceability8          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 Tracblt,omitempty"`
}

// May be one of DCTV, DELT, DWNL, INST, RSTR, UPLD, UPDT, BIND, RBND, UBND, ACTV
type TerminalManagementAction4Code string

// May be one of ACCD, CNTE, FMTE, INVC, LENE, OVER, MISS, NSUP, SIGE, SUCC, SYNE, TIMO, UKDT, UKRF, INDP, IDMP, DPRU, AERR, CMER, ULER
type TerminalManagementActionResult4Code string

type Traceability8 struct {
	RlayId      GenericIdentification177 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 RlayId"`
	PrtcolNm    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 PrtcolNm,omitempty"`
	PrtcolVrsn  Max6Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 PrtcolVrsn,omitempty"`
	TracDtTmIn  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 TracDtTmIn"`
	TracDtTmOut ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 TracDtTmOut"`
}

type TriggerInformation2 struct {
	TrggrSrc PartyType5Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 TrggrSrc"`
	SrcId    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 SrcId"`
	TrggrTp  ExchangePolicy2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 TrggrTp"`
	AddtlInf Max70Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.09 AddtlInf,omitempty"`
}

// May be one of CDSP, CRCP, MDSP, MRCP, CRDO
type UserInterface4Code string

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
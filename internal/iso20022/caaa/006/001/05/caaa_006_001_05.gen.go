// Code generated by main. DO NOT EDIT.

package caaa_006_001_05

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorCancellationResponse5 struct {
	Envt   CardPaymentEnvironment46 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Envt"`
	Tx     CardPaymentTransaction57 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Tx"`
	TxRspn CardPaymentTransaction58 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TxRspn"`
}

type AcceptorCancellationResponseV05 struct {
	Hdr      Header30                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Hdr"`
	CxlRspn  AcceptorCancellationResponse5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 CxlRspn"`
	SctyTrlr ContentInformationType15      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 SctyTrlr,omitempty"`
}

type Action6 struct {
	ActnTp    ActionType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 ActnTp"`
	MsgToPres ActionMessage2  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 MsgToPres,omitempty"`
}

type ActionMessage2 struct {
	MsgDstn      UserInterface4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 MsgDstn"`
	Frmt         OutputFormat1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Frmt,omitempty"`
	MsgCntt      Max20000Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 MsgCntt"`
	MsgCnttSgntr Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 MsgCnttSgntr,omitempty"`
}

// May be one of BUSY, CPTR, DISP, NOVR, RQID, PINL, PINR, PRNT, RFRL, RQDT, DCCQ
type ActionType3Code string

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

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
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 MAC"`
}

type AuthorisationResult12 struct {
	AuthstnNtty   GenericIdentification90 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 AuthstnNtty,omitempty"`
	RspnToAuthstn ResponseType5           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 RspnToAuthstn"`
	AuthstnCd     Min6Max8Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 AuthstnCd,omitempty"`
	TMSTrggr      TMSTrigger1             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TMSTrggr,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CardPaymentEnvironment46 struct {
	AcqrrId  GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 AcqrrId,omitempty"`
	MrchntId GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 MrchntId,omitempty"`
	POIId    GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 POIId"`
	Card     PaymentCard19           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Card,omitempty"`
	PmtTkn   CardPaymentToken2       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 PmtTkn,omitempty"`
}

type CardPaymentToken2 struct {
	TknChrtc     []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TknChrtc,omitempty"`
	TknAssrncLvl float64     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TknAssrncLvl,omitempty"`
}

type CardPaymentTransaction57 struct {
	SaleRefId    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 SaleRefId,omitempty"`
	TxId         TransactionIdentifier1          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TxId"`
	InitrTxId    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 InitrTxId,omitempty"`
	RcptTxId     Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 RcptTxId,omitempty"`
	RcncltnId    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 RcncltnId,omitempty"`
	IntrchngData Max140Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 IntrchngData,omitempty"`
	TxDtls       CardPaymentTransactionDetails35 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TxDtls"`
}

type CardPaymentTransaction58 struct {
	AuthstnRslt AuthorisationResult12 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 AuthstnRslt"`
	Actn        []Action6             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Actn,omitempty"`
}

type CardPaymentTransactionDetails35 struct {
	Ccy         ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Ccy"`
	TtlAmt      float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TtlAmt"`
	ICCRltdData Max10000Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 ICCRltdData,omitempty"`
}

// May be one of COMM, CONS
type CardProductType1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 RltvDstngshdNm"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 EnvlpdData"`
}

type ContentInformationType15 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 CnttTp"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 AuthntcdData"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

type Document struct {
	AccptrCxlRspn AcceptorCancellationResponseV05 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 AccptrCxlRspn"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 NcrptdCntt,omitempty"`
}

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 ShrtNm,omitempty"`
}

type GenericIdentification53 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Tp,omitempty"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 ShrtNm,omitempty"`
}

type GenericIdentification76 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 ShrtNm,omitempty"`
}

type GenericIdentification90 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Id,omitempty"`
	Tp     PartyType14Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 ShrtNm,omitempty"`
}

type GenericIdentification94 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Id"`
	Tp       PartyType3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Tp,omitempty"`
	Issr     PartyType4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 ShrtNm,omitempty"`
	RmotAccs NetworkParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 RmotAccs,omitempty"`
}

type Header30 struct {
	MsgFctn    MessageFunction10Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 PrtcolVrsn"`
	XchgId     float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 XchgId"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 CreDtTm"`
	InitgPty   GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 InitgPty"`
	RcptPty    GenericIdentification94 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 RcptPty,omitempty"`
	Tracblt    []Traceability5         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 NcrptdKey"`
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

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
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
type Max20000Text string

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
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// May be one of AUTQ, AUTP, FAUQ, FAUP, CMPV, CMPK, FCMV, FCMK, RVRA, RVRR, FRVA, FRVR, CCAQ, CCAP, CCAV, CCAK, DGNP, DGNQ, RCLQ, RCLP, DCCQ, DCCP
type MessageFunction10Code string

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

// Must be at least 6 items long
type Min6Max8Text string

// Must match the pattern [0-9]{8,28}
type Min8Max28NumericText string

type NetworkParameters4 struct {
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 NtwkTp"`
	AdrVal Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 AdrVal"`
}

type NetworkParameters5 struct {
	Adr        []NetworkParameters4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 SctyPrfl,omitempty"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

// May be one of MREF, TEXT, HTML
type OutputFormat1Code string

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 BPddg,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS, ICCA
type PartyType14Code string

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

type PaymentCard19 struct {
	PrtctdCardData ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 PrtctdCardData,omitempty"`
	PlainCardData  PlainCardData8           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 PlainCardData,omitempty"`
	PmtAcctRef     Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 PmtAcctRef,omitempty"`
	MskdPAN        string                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 MskdPAN,omitempty"`
	CardBrnd       Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 CardBrnd,omitempty"`
	CardPdctTp     CardProductType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 CardPdctTp,omitempty"`
}

type PlainCardData8 struct {
	PAN       Min8Max28NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 PAN"`
	CardSeqNb Min2Max3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 CardSeqNb,omitempty"`
	FctvDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 FctvDt,omitempty"`
	XpryDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 XpryDt"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 KeyIdr"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 AttrVal"`
}

// May be one of APPR, DECL, PART
type Response4Code string

type ResponseType5 struct {
	Rspn         Response4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Rspn"`
	RspnRsn      Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 RspnRsn,omitempty"`
	AddtlRspnInf Max140Text    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 AddtlRspnInf,omitempty"`
}

// May be one of CRIT, ASAP, DTIM
type TMSContactLevel1Code string

type TMSTrigger1 struct {
	TMSCtctLvl  TMSContactLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TMSCtctLvl"`
	TMSId       Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TMSId,omitempty"`
	TMSCtctDtTm ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TMSCtctDtTm,omitempty"`
}

type Traceability5 struct {
	RlayId      GenericIdentification76 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 RlayId"`
	PrtcolNm    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 PrtcolNm,omitempty"`
	PrtcolVrsn  Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 PrtcolVrsn,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TracDtTmOut"`
}

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 TxRef"`
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

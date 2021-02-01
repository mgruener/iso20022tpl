// Code generated by main. DO NOT EDIT.

package catm_007_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5
type Algorithm12Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of ERS2, ERS1, RPSS
type Algorithm14Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Param,omitempty"`
}

type AlgorithmIdentification16 struct {
	Algo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Algo"`
}

type AlgorithmIdentification17 struct {
	Algo  Algorithm14Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Algo"`
	Param Parameter8      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

// May be one of EMAL, CHLG
type AttributeType2Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 MAC"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// May be one of CRTC, CRTR, CRTK, WLSR, WLSA
type CardPaymentServiceType10Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 RltvDstngshdNm"`
}

type CertificateManagementRequest1 struct {
	POIId            GenericIdentification72      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 POIId"`
	TMId             GenericIdentification72      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 TMId,omitempty"`
	CertSvc          CardPaymentServiceType10Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 CertSvc"`
	SctyDomn         Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SctyDomn,omitempty"`
	BinryCertfctnReq Max20000Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 BinryCertfctnReq,omitempty"`
	CertfctnReq      CertificationRequest1        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 CertfctnReq,omitempty"`
	ClntCert         Max10KBinary                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 ClntCert,omitempty"`
	WhtListId        PointOfInteraction6          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 WhtListId,omitempty"`
}

type CertificateManagementRequestV01 struct {
	Hdr         Header29                      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Hdr"`
	CertMgmtReq CertificateManagementRequest1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 CertMgmtReq"`
	SctyTrlr    ContentInformationType13      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SctyTrlr,omitempty"`
}

type CertificationRequest1 struct {
	CertReqInf CertificationRequest2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 CertReqInf"`
	KeyId      Max140Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 KeyId,omitempty"`
	KeyVrsn    Max140Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 KeyVrsn,omitempty"`
}

type CertificationRequest2 struct {
	Vrsn           float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Vrsn,omitempty"`
	SbjtNm         CertificateIssuer1           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SbjtNm,omitempty"`
	SbjtPblcKeyInf PublicRSAKey2                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SbjtPblcKeyInf"`
	Attr           []RelativeDistinguishedName2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Attr"`
}

type ContentInformationType13 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 CnttTp"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 AuthntcdData,omitempty"`
	SgndData     SignedData4        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SgndData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

type Document struct {
	CertMgmtReq CertificateManagementRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 CertMgmtReq"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Cntt,omitempty"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type GenericIdentification72 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Id"`
	Issr   PartyType6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 ShrtNm,omitempty"`
}

type GenericIdentification93 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Id"`
	Issr     PartyType6Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 ShrtNm,omitempty"`
	RmotAccs NetworkParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 RmotAccs,omitempty"`
}

type Header29 struct {
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 PrtcolVrsn"`
	XchgId     float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 XchgId,omitempty"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 CreDtTm"`
	InitgPty   GenericIdentification72 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 InitgPty"`
	RcptPty    GenericIdentification93 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 RcptPty,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 NcrptdKey"`
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
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type NetworkParameters4 struct {
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 NtwkTp"`
	AdrVal Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 AdrVal"`
}

type NetworkParameters5 struct {
	Adr        []NetworkParameters4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SctyPrfl,omitempty"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 BPddg,omitempty"`
}

type Parameter8 struct {
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 DgstAlgo"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 MskGnrtrAlgo"`
	SaltLngth    float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SaltLngth"`
	TrlrFld      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 TrlrFld,omitempty"`
}

// May be one of ACCP, MERC, ACQR, ITAG, MTMG, TMGT
type PartyType6Code string

type PointOfInteraction6 struct {
	ManfctrIdr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 ManfctrIdr"`
	Mdl        Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Mdl"`
	SrlNb      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SrlNb"`
}

type PublicRSAKey1 struct {
	Mdlus Max5000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Mdlus"`
	Expnt Max5000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Expnt"`
}

type PublicRSAKey2 struct {
	Algo       Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Algo,omitempty"`
	PblcKeyVal PublicRSAKey1  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 PblcKeyVal"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 KeyIdr"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 AttrVal"`
}

type RelativeDistinguishedName2 struct {
	AttrTp  AttributeType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 AttrVal"`
}

type SignedData4 struct {
	Vrsn        float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 NcpsltdCntt"`
	Cert        []Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Cert,omitempty"`
	Sgnr        []Signer3                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Sgnr"`
}

type Signer3 struct {
	Vrsn      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Vrsn,omitempty"`
	SgnrId    Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SgnrId,omitempty"`
	DgstAlgo  AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 DgstAlgo"`
	SgntrAlgo AlgorithmIdentification17 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 SgntrAlgo"`
	Sgntr     Max3000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Sgntr"`
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

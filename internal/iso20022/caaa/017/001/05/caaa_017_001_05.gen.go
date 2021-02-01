// Code generated by main. DO NOT EDIT.

package caaa_017_001_05

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorCurrencyConversionResponse5 struct {
	Envt         CardPaymentEnvironment69 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Envt"`
	Tx           CardPaymentTransaction77 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Tx"`
	CcyConvsRslt CurrencyConversion16     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CcyConvsRslt"`
}

type AcceptorCurrencyConversionResponseV05 struct {
	Hdr          Header35                            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Hdr"`
	CcyConvsRspn AcceptorCurrencyConversionResponse5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CcyConvsRspn"`
	SctyTrlr     ContentInformationType16            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 SctyTrlr,omitempty"`
}

type ActionMessage5 struct {
	Frmt    OutputFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Frmt,omitempty"`
	MsgCntt Max20000Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 MsgCntt"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of HS25, HS38, HS51, HS01, SH31, SH32, SH33, SH35, SHK1, SHK2
type Algorithm16Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5, CMA2, CM31, CM32, CM33, MCS3, CCA1, CCA2, CCA3
type Algorithm17Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C, DA12, DA19, DA25, N108, EA5R, EA9R, EA2R, E3DR, E36C, E36R, SD5C
type Algorithm18Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification18 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Algo"`
	Param Parameter9     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Param,omitempty"`
}

type AlgorithmIdentification19 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Algo"`
	Param Parameter10    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Param,omitempty"`
}

type AlgorithmIdentification22 struct {
	Algo  Algorithm17Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Param,omitempty"`
}

type AlgorithmIdentification23 struct {
	Algo  Algorithm18Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Algo"`
	Param Parameter12     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Param,omitempty"`
}

type AlgorithmIdentification24 struct {
	Algo  Algorithm18Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Algo"`
	Param Parameter12     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData5 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Vrsn,omitempty"`
	Rcpt        []Recipient6Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Rcpt"`
	MACAlgo     AlgorithmIdentification22 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 MAC"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// May be one of CTDP, CHCK, CRDT, CURR, CDBT, DFLT, EPRS, HEQL, ISTL, INVS, LCDT, MBNW, MNMK, MNMC, MTGL, RTRM, RVLV, SVNG, STBD, UVRL, PRPD, FLTC
type CardAccountType3Code string

type CardPaymentEnvironment69 struct {
	AcqrrId  GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AcqrrId,omitempty"`
	MrchntId GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 MrchntId,omitempty"`
	POIId    GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 POIId,omitempty"`
	Card     PaymentCard28           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Card,omitempty"`
	PmtTkn   CardPaymentToken4       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 PmtTkn,omitempty"`
}

type CardPaymentToken4 struct {
	Tkn           Min8Max28NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Tkn,omitempty"`
	CardSeqNb     Min2Max3NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CardSeqNb,omitempty"`
	TknXpryDt     Max10Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TknXpryDt,omitempty"`
	TknChrtc      []Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TknChrtc,omitempty"`
	TknRqstr      PaymentTokenIdentifiers1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TknRqstr,omitempty"`
	TknAssrncLvl  float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TknAssrncLvl,omitempty"`
	TknAssrncData Max500Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TknAssrncData,omitempty"`
}

type CardPaymentTransaction77 struct {
	SaleRefId     Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 SaleRefId,omitempty"`
	TxId          TransactionIdentifier1          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TxId"`
	InitrTxId     Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 InitrTxId,omitempty"`
	RcptTxId      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 RcptTxId,omitempty"`
	RcncltnId     Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 RcncltnId,omitempty"`
	IntrchngData  Max140Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 IntrchngData,omitempty"`
	TxDtls        CardPaymentTransactionDetails28 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TxDtls"`
	MrchntRefData Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 MrchntRefData,omitempty"`
}

type CardPaymentTransactionDetails28 struct {
	Ccy         ActiveCurrencyCode   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Ccy"`
	TtlAmt      float64              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TtlAmt"`
	DtldAmt     DetailedAmount15     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 DtldAmt,omitempty"`
	InvcAmt     DetailedAmount4      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 InvcAmt,omitempty"`
	VldtyDt     ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 VldtyDt,omitempty"`
	AcctTp      CardAccountType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AcctTp,omitempty"`
	ICCRltdData Max10000Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 ICCRltdData,omitempty"`
}

// May be one of COMM, CONS
type CardProductType1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 RltvDstngshdNm"`
}

type Commission18 struct {
	Rate     float64    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Rate"`
	AddtlInf Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AddtlInf,omitempty"`
}

type Commission19 struct {
	Amt      float64    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Amt"`
	AddtlInf Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AddtlInf,omitempty"`
}

type ContentInformationType16 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CnttTp"`
	AuthntcdData AuthenticatedData5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AuthntcdData"`
}

type ContentInformationType17 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CnttTp"`
	EnvlpdData EnvelopedData5   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 EnvlpdData"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

type CurrencyConversion14 struct {
	CcyConvsId    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CcyConvsId,omitempty"`
	TrgtCcy       CurrencyDetails3       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TrgtCcy"`
	RsltgAmt      float64                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 RsltgAmt"`
	XchgRate      float64                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 XchgRate"`
	NvrtdXchgRate float64                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 NvrtdXchgRate,omitempty"`
	QtnDt         ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 QtnDt,omitempty"`
	VldUntil      ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 VldUntil,omitempty"`
	SrcCcy        CurrencyDetails2       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 SrcCcy"`
	OrgnlAmt      OriginalAmountDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 OrgnlAmt"`
	ComssnDtls    []Commission19         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 ComssnDtls,omitempty"`
	MrkUpDtls     []Commission18         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 MrkUpDtls,omitempty"`
	DclrtnDtls    ActionMessage5         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 DclrtnDtls,omitempty"`
}

type CurrencyConversion16 struct {
	Rslt      CurrencyConversionResponse3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Rslt"`
	RsltRsn   Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 RsltRsn,omitempty"`
	ConvsDtls []CurrencyConversion14          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 ConvsDtls,omitempty"`
}

// May be one of ODCC, DCCA, ICRD, IMER, IPRD, IRAT, NDCC, REST, CATG
type CurrencyConversionResponse3Code string

type CurrencyDetails2 struct {
	AlphaCd ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AlphaCd,omitempty"`
	NmrcCd  Exact3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 NmrcCd,omitempty"`
	Dcml    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Dcml,omitempty"`
	Nm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Nm,omitempty"`
}

type CurrencyDetails3 struct {
	AlphaCd ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AlphaCd"`
	NmrcCd  Exact3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 NmrcCd"`
	Dcml    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Dcml"`
	Nm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Nm,omitempty"`
}

type DetailedAmount15 struct {
	AmtGoodsAndSvcs float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AmtGoodsAndSvcs,omitempty"`
	CshBck          float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CshBck,omitempty"`
	Grtty           float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Grtty,omitempty"`
	Fees            []DetailedAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Fees,omitempty"`
	Rbt             []DetailedAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Rbt,omitempty"`
	ValAddedTax     []DetailedAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 ValAddedTax,omitempty"`
	Srchrg          []DetailedAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Srchrg,omitempty"`
}

type DetailedAmount4 struct {
	Amt  float64    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Amt"`
	Labl Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Labl,omitempty"`
}

type Document struct {
	AccptrCcyConvsRspn AcceptorCurrencyConversionResponseV05 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AccptrCcyConvsRspn"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Cntt,omitempty"`
}

type EncryptedContent4 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification24 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CnttNcrptnAlgo,omitempty"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 NcrptdData"`
}

// May be one of TR31, TR34, I238
type EncryptionFormat2Code string

type EnvelopedData5 struct {
	Vrsn       float64                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Vrsn,omitempty"`
	OrgtrInf   OriginatorInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 OrgtrInf,omitempty"`
	Rcpt       []Recipient6Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Rcpt"`
	NcrptdCntt EncryptedContent4      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 NcrptdCntt,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{3}
type Exact3AlphaNumericText string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 ShrtNm,omitempty"`
}

type GenericIdentification53 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Tp,omitempty"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 ShrtNm,omitempty"`
}

type GenericIdentification76 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 ShrtNm,omitempty"`
}

type GenericIdentification94 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Id"`
	Tp       PartyType3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Tp,omitempty"`
	Issr     PartyType4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 ShrtNm,omitempty"`
	RmotAccs NetworkParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 RmotAccs,omitempty"`
}

type Header35 struct {
	MsgFctn    MessageFunction14Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 PrtcolVrsn"`
	XchgId     float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 XchgId"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CreDtTm"`
	InitgPty   GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 InitgPty"`
	RcptPty    GenericIdentification94 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 RcptPty,omitempty"`
	Tracblt    []Traceability5         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Tracblt,omitempty"`
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
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 SrlNb"`
}

type KEK5 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification23 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 DerivtnId,omitempty"`
}

type KeyTransport5 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 NcrptdKey"`
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
type Max104Text string

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

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Must be at least 1 items long
type Max20000Text string

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

// Must be at least 1 items long
type Max37Text string

// Must be at least 1 items long
type Max3Text string

// Must be at least 1 items long
type Max45Text string

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

// Must be at least 1 items long
type Max76Text string

// May be one of AUTQ, AUTP, CCAV, CCAK, CCAQ, CCAP, CMPV, CMPK, DCAV, DCRR, DCCQ, DCCP, DGNP, DGNQ, FAUQ, FAUP, FCMV, FCMK, FRVA, FRVR, RCLQ, RCLP, RVRA, RVRR, CDDQ, CDDK, CDDR, CDDP
type MessageFunction14Code string

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

type NetworkParameters4 struct {
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 NtwkTp"`
	AdrVal Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AdrVal"`
}

type NetworkParameters5 struct {
	Adr        []NetworkParameters4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 SctyPrfl,omitempty"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

type OriginalAmountDetails1 struct {
	ActlAmt float64 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 ActlAmt,omitempty"`
	MinAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 MinAmt,omitempty"`
	MaxAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 MaxAmt,omitempty"`
}

type OriginatorInformation1 struct {
	Cert []Max5000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Cert,omitempty"`
}

// May be one of MREF, TEXT, HTML
type OutputFormat1Code string

type Parameter10 struct {
	NcrptnFrmt   EncryptionFormat2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 MskGnrtrAlgo,omitempty"`
}

type Parameter12 struct {
	NcrptnFrmt   EncryptionFormat2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 BPddg,omitempty"`
}

type Parameter9 struct {
	DgstAlgo Algorithm16Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 DgstAlgo,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

type PaymentCard28 struct {
	PrtctdCardData ContentInformationType17 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 PrtctdCardData,omitempty"`
	PrvtCardData   Max100KBinary            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 PrvtCardData,omitempty"`
	PlainCardData  PlainCardData15          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 PlainCardData,omitempty"`
	PmtAcctRef     Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 PmtAcctRef,omitempty"`
	MskdPAN        string                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 MskdPAN,omitempty"`
	IssrBIN        Max15NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 IssrBIN,omitempty"`
	CardCtryCd     Max3Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CardCtryCd,omitempty"`
	CardCcyCd      Exact3AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CardCcyCd,omitempty"`
	CardPdctPrfl   Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CardPdctPrfl,omitempty"`
	CardBrnd       Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CardBrnd,omitempty"`
	CardPdctTp     CardProductType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CardPdctTp,omitempty"`
	CardPdctSubTp  Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CardPdctSubTp,omitempty"`
	IntrnlCard     bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 IntrnlCard,omitempty"`
	AllwdPdct      []Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AllwdPdct,omitempty"`
	SvcOptn        Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 SvcOptn,omitempty"`
	AddtlCardData  Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AddtlCardData,omitempty"`
}

type PaymentTokenIdentifiers1 struct {
	PrvdrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 PrvdrId"`
	RqstrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 RqstrId"`
}

type PlainCardData15 struct {
	PAN       Min8Max28NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 PAN"`
	CardSeqNb Min2Max3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CardSeqNb,omitempty"`
	FctvDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 FctvDt,omitempty"`
	XpryDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 XpryDt"`
	SvcCd     Exact3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 SvcCd,omitempty"`
	Trck1     Max76Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Trck1,omitempty"`
	Trck2     Max37Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Trck2,omitempty"`
	Trck3     Max104Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 Trck3,omitempty"`
	CrdhldrNm Max45Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 CrdhldrNm,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 KeyIdr"`
}

type Recipient6Choice struct {
	KeyTrnsprt KeyTransport5  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 KeyTrnsprt"`
	KEK        KEK5           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 AttrVal"`
}

type Traceability5 struct {
	RlayId      GenericIdentification76 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 RlayId"`
	PrtcolNm    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 PrtcolNm,omitempty"`
	PrtcolVrsn  Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 PrtcolVrsn,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TracDtTmOut"`
}

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.05 TxRef"`
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
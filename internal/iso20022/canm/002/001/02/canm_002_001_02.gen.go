// Code generated by main. DO NOT EDIT.

package canm_002_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type Action9 struct {
	Dstn       PartyType20Code  `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Dstn,omitempty"`
	ActnTp     ActionType11Code `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 ActnTp,omitempty"`
	OthrActnTp Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrActnTp,omitempty"`
	ActnInf    Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 ActnInf,omitempty"`
}

// May be one of FILE, MOBL, OTHN, OTHP, PECR, POFS
type ActionDestination1Code string

// May be one of ACTV, DEAC, DISP, FUPD, PRNT, SNDM
type ActionType10Code string

// May be one of CNTI, CNIS, CNTA, CNAS, CPTR, CHDV, VIPM, TRCK, TRXR, OTHN, OTHP, SIGN
type ActionType11Code string

type AdditionalAction1 struct {
	Tp         ActionType10Code       `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Tp,omitempty"`
	Dstn       PartyType21Code        `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Dstn,omitempty"`
	OthrDstn   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrDstn,omitempty"`
	DstnTp     ActionDestination1Code `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 DstnTp,omitempty"`
	OthrDstnTp Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrDstnTp,omitempty"`
	DstnAdr    Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 DstnAdr,omitempty"`
	Frmt       OutputFormat4Code      `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Frmt,omitempty"`
	OthrFrmt   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrFrmt,omitempty"`
	Cntt       Content1               `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Cntt,omitempty"`
}

type AdditionalData1 struct {
	Tp  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Tp,omitempty"`
	Val Max2048Text `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Val,omitempty"`
}

type AdditionalFee1 struct {
	Tp         TypeOfAmount10Code `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Tp"`
	OthrTp     Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrTp,omitempty"`
	FeePrgm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 FeePrgm,omitempty"`
	FeeDscrptr Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 FeeDscrptr,omitempty"`
	Amt        FeeAmount2         `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Amt"`
	Labl       Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Labl,omitempty"`
}

type AdditionalInformation20 struct {
	Rcpt     PartyType22Code      `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Rcpt,omitempty"`
	OthrRcpt Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrRcpt,omitempty"`
	Trgt     []UserInterface6Code `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Trgt,omitempty"`
	OthrTrgt Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrTrgt,omitempty"`
	Frmt     OutputFormat4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Frmt,omitempty"`
	OthrFrmt Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrFrmt,omitempty"`
	Tp       Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Tp,omitempty"`
	Val      Max20KText           `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Val"`
}

type ApprovalData1 struct {
	ApprvlNtty ApprovalEntity1        `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 ApprvlNtty,omitempty"`
	ApprvlCd   Exact6AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 ApprvlCd,omitempty"`
}

type ApprovalEntity1 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Id,omitempty"`
	Tp     PartyType26Code   `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Tp"`
	OthrTp Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrTp,omitempty"`
	Assgnr PartyType9Code    `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Assgnr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 ShrtNm,omitempty"`
}

type BatchManagementInformation1 struct {
	ColltnId         Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 ColltnId,omitempty"`
	BtchId           Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 BtchId"`
	MsgSeqNb         Max15NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 MsgSeqNb,omitempty"`
	MsgChcksmInptVal Max140Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 MsgChcksmInptVal,omitempty"`
}

type CardProgrammeMode1 struct {
	Tp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Tp,omitempty"`
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Id"`
}

type Content1 struct {
	Val    Max20KText   `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Val"`
	Sgntr  Max140Binary `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Sgntr,omitempty"`
	CertId Max70Text    `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 CertId,omitempty"`
}

type ContentInformationType20 struct {
	MACData MACData1          `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 MACData"`
	MAC     Max8HexBinaryText `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 MAC"`
}

type Context8 struct {
	TxCntxt TransactionContext5 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 TxCntxt,omitempty"`
}

type Document struct {
	NtwkMgmtRspn NetworkManagementResponseV02 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 NtwkMgmtRspn"`
}

// May be no more than 12 items long
type Exact12Text string

// May be no more than 15 items long
type Exact15Text string

// Must match the pattern ([0-9A-F][0-9A-F]){1}
type Exact1HexBinaryText string

// Must match the pattern [a-zA-Z0-9]{2}
type Exact2AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// Must match the pattern [a-zA-Z0-9\s]{6}
type Exact6AlphaNumericText string

type FeeAmount2 struct {
	Amt      float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Amt"`
	Ccy      ISO3NumericCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Ccy,omitempty"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 XchgRate,omitempty"`
	QtnDt    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 QtnDt,omitempty"`
	Sgn      bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Sgn,omitempty"`
}

type GenericIdentification172 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Id"`
	Tp     PartyType17Code   `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Tp,omitempty"`
	OthrTp Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrTp,omitempty"`
	Assgnr PartyType18Code   `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Assgnr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 ShrtNm,omitempty"`
}

type Header43 struct {
	MsgFctn        MessageFunction23Code       `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 MsgFctn"`
	PrtcolVrsn     Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 PrtcolVrsn"`
	XchgId         Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 XchgId,omitempty"`
	ReTrnsmssnCntr Max3NumericText             `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 ReTrnsmssnCntr,omitempty"`
	CreDtTm        ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 CreDtTm"`
	BtchMgmtInf    BatchManagementInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 BtchMgmtInf,omitempty"`
	InitgPty       GenericIdentification172    `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 InitgPty"`
	RcptPty        GenericIdentification172    `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 RcptPty,omitempty"`
	TracData       []AdditionalData1           `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 TracData,omitempty"`
	Tracblt        []Traceability7             `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Tracblt,omitempty"`
}

// Must match the pattern [0-9]{3,3}
type ISO3NumericCurrencyCode string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type MACData1 struct {
	Ctrl         Exact1HexBinaryText `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Ctrl"`
	KeySetIdr    Max8NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 KeySetIdr"`
	DrvdInf      Max32HexBinaryText  `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 DrvdInf,omitempty"`
	Algo         Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Algo"`
	KeyLngth     Max4NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 KeyLngth,omitempty"`
	KeyPrtcn     Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 KeyPrtcn,omitempty"`
	KeyIndx      Max5NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 KeyIndx,omitempty"`
	PddgMtd      Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 PddgMtd,omitempty"`
	InitlstnVctr Max32HexBinaryText  `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 InitlstnVctr,omitempty"`
}

// Must be at least 1 items long
type Max1000Text string

// Must match the pattern [0-9]{1,12}
type Max12NumericText string

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
type Max2048Text string

// Must be at least 1 items long
type Max20KText string

// Must match the pattern [0-9]{1,2}
type Max2NumericText string

// Must match the pattern ([0-9A-F][0-9A-F]){1,32}
type Max32HexBinaryText string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must match the pattern [0-9]{1,4}
type Max4NumericText string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max70Text string

// Must match the pattern ([0-9A-F][0-9A-F]){1,8}
type Max8HexBinaryText string

// Must match the pattern [0-9]{1,8}
type Max8NumericText string

// May be one of ADVC, REQU
type MessageFunction23Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type NetworkManagementResponse1 struct {
	Cntxt       Context8             `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Cntxt,omitempty"`
	Tx          Transaction99        `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Tx"`
	PrcgRslt    ProcessingResult1    `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 PrcgRslt"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 SplmtryData,omitempty"`
}

type NetworkManagementResponseV02 struct {
	Hdr      Header43                   `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Hdr"`
	Body     NetworkManagementResponse1 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Body"`
	SctyTrlr ContentInformationType20   `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 SctyTrlr,omitempty"`
}

// May be one of ECTS, ESFW, OTHN, OTHP, SGNF, SGNN, DSFW, TSUN, MOSB, SPIN, IART, SYCL, DRBI, ERBI
type NetworkManagementType1Code string

// May be one of FLNM, MREF, OTHN, OTHP, SMSI, TEXT, URLI, HTML
type OutputFormat4Code string

// May be one of OTHN, OTHP, ACQR, ACQP, CISS, CISP, AGNT
type PartyType17Code string

// May be one of ACQR, CISS, CSCH, AGNT
type PartyType18Code string

// May be one of ACCP, ACQR, CRDH, CISS, AGNT
type PartyType20Code string

// May be one of ACCP, CRDH, OTHN, OTHP
type PartyType21Code string

// May be one of CRDH, MERC, OTHN, OTHP, AGNT
type PartyType22Code string

// May be one of ACCP, ACQR, ICCA, CISS, DLIS, AGNT, OTHN, OTHP
type PartyType26Code string

// May be one of ACQR, ACQP, CISS, CISP, CSCH, SCHP
type PartyType9Code string

type ProcessingResult1 struct {
	ApprvlData    ApprovalData1             `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 ApprvlData,omitempty"`
	RsltData      ResultData1               `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 RsltData,omitempty"`
	OrgnlRsltData ResultData1               `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OrgnlRsltData,omitempty"`
	ActnReqrd     bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 ActnReqrd,omitempty"`
	Actn          []Action9                 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Actn,omitempty"`
	AddtlActn     []AdditionalAction1       `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 AddtlActn,omitempty"`
	AddtlInf      []AdditionalInformation20 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 AddtlInf,omitempty"`
}

// May be one of PRCS, UNPR, UNRV, REJT, TECH, OTHN, OTHP
type Response8Code string

type ResultData1 struct {
	Rslt         Response8Code          `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Rslt,omitempty"`
	OthrRslt     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrRslt,omitempty"`
	RsltDtls     Exact2AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 RsltDtls"`
	OthrRsltDtls Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrRsltDtls,omitempty"`
	AddtlRsltInf []AdditionalData1      `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 AddtlRsltInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type Traceability7 struct {
	RlayId      GenericIdentification172 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 RlayId"`
	TracDtTmIn  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 TracDtTmIn,omitempty"`
	TracDtTmOut ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 TracDtTmOut,omitempty"`
}

type Transaction99 struct {
	NtwkMgmtTp     NetworkManagementType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 NtwkMgmtTp,omitempty"`
	OthrNtwkMgmtTp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 OthrNtwkMgmtTp,omitempty"`
	MsgRsn         []Exact4NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 MsgRsn,omitempty"`
	AltrnMsgRsn    []Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 AltrnMsgRsn,omitempty"`
	TxId           TransactionIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 TxId"`
	NbOfMsgs       float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 NbOfMsgs,omitempty"`
	MaxNbOfMsgs    float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 MaxNbOfMsgs,omitempty"`
	AddtlFees      []AdditionalFee1            `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 AddtlFees,omitempty"`
	TxDesc         Max1000Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 TxDesc,omitempty"`
	AddtlData      []AdditionalData1           `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 AddtlData,omitempty"`
}

type TransactionContext5 struct {
	CardPrgrmmApld CardProgrammeMode1 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 CardPrgrmmApld,omitempty"`
}

type TransactionIdentification12 struct {
	SysTracAudtNb      Max12NumericText                    `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 SysTracAudtNb"`
	TrnsmssnDtTm       ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 TrnsmssnDtTm"`
	RtrvlRefNb         Exact12Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 RtrvlRefNb,omitempty"`
	LifeCyclTracIdData TransactionLifeCycleIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 LifeCyclTracIdData,omitempty"`
}

type TransactionLifeCycleIdentification2 struct {
	Id Exact15Text `xml:"urn:iso:std:iso:20022:tech:xsd:canm.002.001.02 Id"`
}

// May be one of INTC, FEEP, OTHN, OTHP, FEEA
type TypeOfAmount10Code string

// May be one of CDSP, CRCP, MDSP, MRCP, CRDO, FILE, CHAP, MRAP, MRIN
type UserInterface6Code string

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

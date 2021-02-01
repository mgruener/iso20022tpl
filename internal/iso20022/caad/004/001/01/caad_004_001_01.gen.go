// Code generated by main. DO NOT EDIT.

package caad_004_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AdditionalData1 struct {
	Tp  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Tp,omitempty"`
	Val Max2048Text `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Val,omitempty"`
}

type AdditionalInformation21 struct {
	Rcpt PartyType23Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Rcpt,omitempty"`
	Trgt []UserInterface7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Trgt,omitempty"`
	Frmt OutputFormat4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Frmt,omitempty"`
	Val  Max20KText           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Val"`
}

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of HS25, HS38, HS51
type Algorithm20Code string

// May be one of EA2C, E3DC, EA9C, EA5C, EA2R, EA9R, EA5R, E3DR, E36C, E36R, SD5C
type Algorithm23Code string

// May be one of HS25, HS38, HS51
type Algorithm5Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification25 struct {
	Algo  Algorithm23Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Param,omitempty"`
}

type AlgorithmIdentification26 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Algo"`
	Param Algorithm5Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Param,omitempty"`
}

type AlgorithmIdentification27 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Algo"`
	Param Parameter13    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Param,omitempty"`
}

type AlgorithmIdentification28 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Algo"`
	Param Parameter14     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Param,omitempty"`
}

type Amount14 struct {
	Amt float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Amt"`
	Ccy ISO3NumericCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Ccy,omitempty"`
	Sgn bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Sgn,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type BatchManagementInformation1 struct {
	ColltnId         Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ColltnId,omitempty"`
	BtchId           Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 BtchId"`
	MsgSeqNb         Max15NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 MsgSeqNb,omitempty"`
	MsgChcksmInptVal Max140Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 MsgChcksmInptVal,omitempty"`
}

type BatchTransferResponse1 struct {
	Envt        Environment4         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Envt"`
	Cntxt       Context8             `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Cntxt,omitempty"`
	Tx          Transaction96        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Tx"`
	PrcgRslt    ProcessingResult3    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 PrcgRslt"`
	PrtctdData  []ProtectedData1     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 PrtctdData,omitempty"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 SplmtryData,omitempty"`
}

type BatchTransferResponseV01 struct {
	Hdr      Header39                 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Hdr"`
	Body     BatchTransferResponse1   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Body"`
	SctyTrlr ContentInformationType20 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 SctyTrlr,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CardProgrammeMode1 struct {
	Tp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Tp,omitempty"`
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Id"`
}

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RltvDstngshdNm"`
}

type ClearingBatchData1 struct {
	ClrMtd      ClearingMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrMtd"`
	OthrClrMtd  Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 OthrClrMtd,omitempty"`
	ClrPrty     Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrPrty,omitempty"`
	ClrDt       ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrDt,omitempty"`
	ClrTtls     ClearingTotals1     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrTtls,omitempty"`
	IntrchngFee Amount14            `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 IntrchngFee,omitempty"`
	AgtFee      Amount14            `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AgtFee,omitempty"`
}

type ClearingControlTotals1 struct {
	ClrCtrlCnt float64  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrCtrlCnt"`
	ClrCtrlAmt Amount14 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrCtrlAmt"`
}

// May be one of DAYC, DFRD, INST, OTHN, OTHP, RLTM
type ClearingMethod2Code string

type ClearingRecordData1 struct {
	Orgtr       Originator1         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Orgtr,omitempty"`
	Dstn        Destination1        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Dstn,omitempty"`
	ClrMtd      ClearingMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrMtd"`
	OthrClrMtd  Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 OthrClrMtd,omitempty"`
	ClrPrty     Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrPrty,omitempty"`
	ClrDt       ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrDt,omitempty"`
	ClrAmt      Amount14            `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrAmt,omitempty"`
	IntrchngFee Amount14            `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 IntrchngFee,omitempty"`
	AgtFee      Amount14            `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AgtFee,omitempty"`
	OthrAmt     OtherAmount2        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 OthrAmt,omitempty"`
}

type ClearingTotals1 struct {
	Cnt       float64  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Cnt"`
	AcmltdAmt Amount14 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AcmltdAmt"`
}

type ContentInformationType20 struct {
	MACData MACData1          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 MACData"`
	MAC     Max8HexBinaryText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 MAC"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// May be one of EVLP, IFSE
type ContentType3Code string

type Context8 struct {
	TxCntxt TransactionContext5 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 TxCntxt,omitempty"`
}

type Destination1 struct {
	Id     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Id"`
	Assgnr Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Assgnr,omitempty"`
	Ctry   ISO3NumericCountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Ctry,omitempty"`
	ShrtNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ShrtNm,omitempty"`
}

type Document struct {
	BtchTrfRspn BatchTransferResponseV01 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 BtchTrfRspn"`
}

type EncryptedContent5 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification25 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 CnttNcrptnAlgo"`
	NcrptdDataElmt []EncryptedDataElement1   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NcrptdDataElmt"`
}

type EncryptedData1 struct {
	Ctrl           Exact1HexBinaryText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Ctrl,omitempty"`
	KeySetIdr      Max8NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeySetIdr,omitempty"`
	DrvdInf        Max32HexBinaryText      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 DrvdInf,omitempty"`
	Algo           Max2NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Algo,omitempty"`
	KeyLngth       Max4NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyLngth,omitempty"`
	KeyPrtcn       Max2NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyPrtcn,omitempty"`
	KeyIndx        Max5NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyIndx,omitempty"`
	PddgMtd        Max2NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 PddgMtd,omitempty"`
	NcrptdDataFrmt Max2NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NcrptdDataFrmt,omitempty"`
	NcrptdDataElmt []EncryptedDataElement1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NcrptdDataElmt"`
}

type EncryptedData1Choice struct {
	BinryData   Max100KBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 BinryData"`
	HexBinryVal string        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 HexBinryVal"`
}

type EncryptedDataElement1 struct {
	Id                   ExternalEncryptedElementIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Id,omitempty"`
	OthrId               Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 OthrId,omitempty"`
	NcrptdData           EncryptedData1Choice                        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NcrptdData"`
	ClearTxtDataFrmt     EncryptedDataFormat1Code                    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClearTxtDataFrmt,omitempty"`
	OthrClearTxtDataFrmt Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 OthrClearTxtDataFrmt,omitempty"`
}

// May be one of ASCI, BINF, EBCD, HEXF, OTHN, OTHP
type EncryptedDataFormat1Code string

// May be one of TR34, TR31, CTCE, CBCE
type EncryptionFormat3Code string

type EnvelopedData6 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Vrsn,omitempty"`
	Rcpt       []Recipient7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Rcpt"`
	NcrptdCntt EncryptedContent5  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NcrptdCntt,omitempty"`
}

type Environment4 struct {
	Orgtr Originator1  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Orgtr,omitempty"`
	Dstn  Destination1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Dstn"`
}

// May be no more than 12 items long
type Exact12Text string

// May be no more than 15 items long
type Exact15Text string

// Must match the pattern ([0-9A-F][0-9A-F]){1}
type Exact1HexBinaryText string

// Must match the pattern [a-zA-Z0-9]{2}
type Exact2AlphaNumericText string

// Must match the pattern [0-9]{2}
type Exact2NumericText string

// Must match the pattern ([0-9A-F][0-9A-F]){1,3}
type ExternalEncryptedElementIdentification1Code string

type GenericIdentification172 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Id"`
	Tp     PartyType17Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Tp,omitempty"`
	OthrTp Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 OthrTp,omitempty"`
	Assgnr PartyType18Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Assgnr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ShrtNm,omitempty"`
}

type Header39 struct {
	MsgFctn        MessageFunction17Code       `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 MsgFctn"`
	PrtcolVrsn     Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 PrtcolVrsn"`
	XchgId         Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 XchgId,omitempty"`
	ReTrnsmssnCntr Max3NumericText             `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ReTrnsmssnCntr,omitempty"`
	CreDtTm        ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 CreDtTm"`
	BtchMgmtInf    BatchManagementInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 BtchMgmtInf,omitempty"`
	InitgPty       GenericIdentification172    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 InitgPty"`
	RcptPty        GenericIdentification172    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RcptPty,omitempty"`
	TracData       []AdditionalData1           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 TracData,omitempty"`
	Tracblt        []Traceability7             `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Tracblt,omitempty"`
}

// Must match the pattern [0-9]{3,3}
type ISO3NumericCountryCode string

// Must match the pattern [0-9]{3,3}
type ISO3NumericCurrencyCode string

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
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 SrlNb"`
}

type KEK6 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Vrsn,omitempty"`
	KEKId         KEKIdentifier6            `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification28 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NcrptdKey,omitempty"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 DerivtnId,omitempty"`
}

type KEKIdentifier6 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyVrsn,omitempty"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 DerivtnId,omitempty"`
}

type KeyTransport6 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification27 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NcrptdKey"`
}

type MACData1 struct {
	Ctrl         Exact1HexBinaryText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Ctrl"`
	KeySetIdr    Max8NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeySetIdr"`
	DrvdInf      Max32HexBinaryText  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 DrvdInf,omitempty"`
	Algo         Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Algo"`
	KeyLngth     Max4NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyLngth,omitempty"`
	KeyPrtcn     Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyPrtcn,omitempty"`
	KeyIndx      Max5NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyIndx,omitempty"`
	PddgMtd      Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 PddgMtd,omitempty"`
	InitlstnVctr Max32HexBinaryText  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 InitlstnVctr,omitempty"`
}

// Must be at least 1 items long
type Max1000Text string

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

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

// Must match the pattern [0-9]{1,23}
type Max23NumericText string

type Max2MBBinary []byte

func (t *Max2MBBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max2MBBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{1,2}
type Max2NumericText string

// Must match the pattern ([0-9A-F][0-9A-F]){1,32}
type Max32HexBinaryText string

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

// Must match the pattern [0-9]{1,4}
type Max4NumericText string

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

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max70Text string

// Must match the pattern ([0-9A-F][0-9A-F]){1,8}
type Max8HexBinaryText string

// Must match the pattern [0-9]{1,8}
type Max8NumericText string

// Must be at least 1 items long
type Max99Text string

// May be one of NOTI, REQU, ADVC
type MessageFunction17Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Originator1 struct {
	Id     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Id"`
	Assgnr Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Assgnr,omitempty"`
	Ctry   ISO3NumericCountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Ctry,omitempty"`
	ShrtNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ShrtNm,omitempty"`
}

type OtherAmount2 struct {
	ClrCnt      float64  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrCnt"`
	ClrAmt      Amount14 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrAmt"`
	IntrchngFee Amount14 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 IntrchngFee,omitempty"`
	AgtFee      Amount14 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AgtFee,omitempty"`
}

// May be one of FLNM, MREF, OTHN, OTHP, SMSI, TEXT, URLI, HTML
type OutputFormat4Code string

type Parameter13 struct {
	DgstAlgo     Algorithm20Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification26 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 MskGnrtrAlgo,omitempty"`
}

type Parameter14 struct {
	NcrptnFrmt   EncryptionFormat3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 BPddg,omitempty"`
}

// May be one of OTHN, OTHP, ACQR, ACQP, CISS, CISP, AGNT
type PartyType17Code string

// May be one of ACQR, CISS, CSCH, AGNT
type PartyType18Code string

// May be one of OTHN, OTHP, CLRA
type PartyType23Code string

type ProcessingResult3 struct {
	RsltData      ResultData2 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RsltData"`
	OrgnlRsltData ResultData1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 OrgnlRsltData,omitempty"`
}

type ProtectedData1 struct {
	CnttTp     ContentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 CnttTp"`
	EnvlpdData EnvelopedData6   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 EnvlpdData,omitempty"`
	NcrptdData EncryptedData1   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NcrptdData,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyIdr"`
}

type Recipient7Choice struct {
	KeyTrnsprt KeyTransport6  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyTrnsprt"`
	KEK        KEK6           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KEK"`
	KeyIdr     KEKIdentifier6 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyIdr"`
}

type Record1 struct {
	SeqCntr           float64              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 SeqCntr"`
	RcrdChcksmInptVal Max140Binary         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RcrdChcksmInptVal,omitempty"`
	ClrRcrdData       ClearingRecordData1  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrRcrdData,omitempty"`
	RcrdMsg           RecordMessage1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RcrdMsg"`
}

type RecordMessage1Choice struct {
	AdddmInitn      Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AdddmInitn"`
	AdddmRspn       Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AdddmRspn"`
	Amdmnt          Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Amdmnt"`
	AuthstnInitn    Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AuthstnInitn"`
	AuthstnRspn     Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AuthstnRspn"`
	CardMgmtInitn   Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 CardMgmtInitn"`
	CardMgmtRspn    Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 CardMgmtRspn"`
	ChrgBckInitn    Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ChrgBckInitn"`
	ChrgBckRspn     Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ChrgBckRspn"`
	Err             Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Err"`
	FeeColltnInitn  Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 FeeColltnInitn"`
	FeeColltnRspn   Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 FeeColltnRspn"`
	FileActnInitn   Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 FileActnInitn"`
	FileActnRspn    Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 FileActnRspn"`
	FinInitn        Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 FinInitn"`
	FinRspn         Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 FinRspn"`
	FrdDspstnInitn  Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 FrdDspstnInitn"`
	FrdDspstnRspn   Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 FrdDspstnRspn"`
	FrdRptgInitn    Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 FrdRptgInitn"`
	FrdRptgRspn     Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 FrdRptgRspn"`
	NqryInitn       Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NqryInitn"`
	NqryRspn        Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NqryRspn"`
	KeyXchgInitn    Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyXchgInitn"`
	KeyXchgRspn     Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 KeyXchgRspn"`
	NtwkMgmtInitn   Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NtwkMgmtInitn"`
	NtwkMgmtRspn    Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NtwkMgmtRspn"`
	RcncltnInitn    Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RcncltnInitn"`
	RcncltnRspn     Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RcncltnRspn"`
	RtrvlInitn      Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RtrvlInitn"`
	RtrvlRspn       Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RtrvlRspn"`
	RtrvlFlfmtInitn Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RtrvlFlfmtInitn"`
	RtrvlFlfmtRspn  Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RtrvlFlfmtRspn"`
	RvslInitn       Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RvslInitn"`
	RvslRspn        Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RvslRspn"`
	SttlmRptgInitn  Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 SttlmRptgInitn"`
	SttlmRptgRspn   Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 SttlmRptgRspn"`
	VrfctnInitn     Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 VrfctnInitn"`
	VrfctnRspn      Max2MBBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 VrfctnRspn"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AttrVal"`
}

// May be one of UNPR, PPRC, PRCS, REJT, OTHP, OTHN, TECH, UNRV
type Response6Code string

// May be one of PRCS, UNPR, UNRV, REJT, TECH, OTHN, OTHP
type Response8Code string

type ResultData1 struct {
	Rslt         Response8Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Rslt,omitempty"`
	OthrRslt     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 OthrRslt,omitempty"`
	RsltDtls     Exact2AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RsltDtls"`
	OthrRsltDtls Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 OthrRsltDtls,omitempty"`
	AddtlRsltInf []AdditionalData1      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AddtlRsltInf,omitempty"`
}

type ResultData2 struct {
	Rslt         Response6Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Rslt,omitempty"`
	OthrRslt     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 OthrRslt,omitempty"`
	RsltDtls     Exact2AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RsltDtls"`
	OthrRsltDtls Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 OthrRsltDtls,omitempty"`
	AddtlRsltInf []AdditionalData1      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AddtlRsltInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type Traceability7 struct {
	RlayId      GenericIdentification172 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RlayId"`
	TracDtTmIn  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 TracDtTmIn,omitempty"`
	TracDtTmOut ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 TracDtTmOut,omitempty"`
}

type Transaction96 struct {
	TxId        TransactionIdentification10 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 TxId,omitempty"`
	BtchId      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 BtchId,omitempty"`
	OrgnlBtchId Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 OrgnlBtchId,omitempty"`
	NbOfMsgs    float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 NbOfMsgs,omitempty"`
	BtchChcksm  Max35Binary                 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 BtchChcksm,omitempty"`
	ReqAck      bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ReqAck,omitempty"`
	ClrBtchData []ClearingBatchData1        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrBtchData,omitempty"`
	ClrCtrlTtls ClearingControlTotals1      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 ClrCtrlTtls,omitempty"`
	AgtData     []AdditionalInformation21   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AgtData,omitempty"`
	Rcrd        []Record1                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Rcrd,omitempty"`
	AddtlData   []AdditionalData1           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AddtlData,omitempty"`
}

type TransactionContext5 struct {
	CardPrgrmmApld CardProgrammeMode1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 CardPrgrmmApld,omitempty"`
}

type TransactionIdentification10 struct {
	LclDtTm             ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 LclDtTm"`
	TmZone              Max70Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 TmZone,omitempty"`
	TxRef               Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 TxRef,omitempty"`
	TrnsmssnDtTm        ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 TrnsmssnDtTm,omitempty"`
	SysTracAudtNb       Max12NumericText                    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 SysTracAudtNb"`
	RtrvlRefNb          Exact12Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 RtrvlRefNb"`
	LifeCyclSpprtInd    Exact2NumericText                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 LifeCyclSpprtInd,omitempty"`
	LifeCyclTracIdData  TransactionLifeCycleIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 LifeCyclTracIdData,omitempty"`
	LifeCyclTracIdMssng Max70Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 LifeCyclTracIdMssng,omitempty"`
	AcqrrRefData        Max99Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AcqrrRefData,omitempty"`
	AcqrrRefNb          Max23NumericText                    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AcqrrRefNb,omitempty"`
	CardIssrRefData     Max1000Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 CardIssrRefData,omitempty"`
}

type TransactionLifeCycleIdentification1 struct {
	Id              Exact15Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 Id"`
	AuthstnSeqNb    Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AuthstnSeqNb,omitempty"`
	PresntmntSeqNb  Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 PresntmntSeqNb,omitempty"`
	PresntmntSeqCnt Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 PresntmntSeqCnt,omitempty"`
	AuthntcnTkn     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.004.001.01 AuthntcnTkn,omitempty"`
}

// May be one of OCAI, CLRL, CLRA
type UserInterface7Code string

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

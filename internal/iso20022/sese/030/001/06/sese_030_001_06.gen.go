// Code generated by main. DO NOT EDIT.

package sese_030_001_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AcceptedReason9Choice struct {
	Cd    AcknowledgementReason8Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prtry"`
}

type AcceptedStatus7Choice struct {
	NoSpcfdRsn NoReasonCode            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 NoSpcfdRsn"`
	Rsn        []AcceptedStatusReason8 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Rsn"`
}

type AcceptedStatusReason8 struct {
	RsnCd       AcceptedReason9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 RsnCd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AddtlRsnInf,omitempty"`
}

// May be one of NARR
type AcknowledgementReason8Code string

type AdditionalInformation11 struct {
	AcctOwnrTxId Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AcctOwnrTxId,omitempty"`
	ClssfctnTp   ClassificationType32Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 ClssfctnTp,omitempty"`
	SfkpgAcct    SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 SfkpgAcct,omitempty"`
	FinInstrmId  SecurityIdentification19           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 FinInstrmId,omitempty"`
	Qty          FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Qty,omitempty"`
	FctvDt       DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 FctvDt,omitempty"`
	XpryDt       DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 XpryDt,omitempty"`
	CutOffDt     DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 CutOffDt,omitempty"`
	Invstr       PartyIdentification100             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Invstr,omitempty"`
	DlvrgPty1    PartyIdentificationAndAccount117   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 DlvrgPty1,omitempty"`
	RcvgPty1     PartyIdentificationAndAccount117   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 RcvgPty1,omitempty"`
	PrcgSts      ProcessingStatus56Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 PrcgSts,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of LAMI, NBOR, YBOR, RTRN
type AutoBorrowing2Code string

type AutomaticBorrowing7Choice struct {
	Cd    AutoBorrowing2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prtry"`
}

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

type CancellationReason11 struct {
	Cd          CancellationReason22Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	AddtlRsnInf Max210Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AddtlRsnInf,omitempty"`
}

type CancellationReason22Choice struct {
	Cd    CancelledStatusReason14Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prtry"`
}

type CancelledStatus10Choice struct {
	NoSpcfdRsn NoReasonCode         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 NoSpcfdRsn"`
	Rsn        CancellationReason11 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Rsn"`
}

// May be one of NARR
type CancelledStatusReason14Code string

type ClassificationType32Choice struct {
	ClssfctnFinInstrm CFIOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AltrnClssfctn"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 DtTm"`
}

type Document struct {
	SctiesSttlmCondsModReq SecuritiesSettlementConditionsModificationRequestV06 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 SctiesSttlmCondsModReq"`
}

type DocumentNumber5Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 LngNb"`
	PrtryNb GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AmtsdVal"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 SchmeNm,omitempty"`
}

type HoldIndicator6 struct {
	Ind bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Ind"`
	Rsn []RegistrationReason5 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Rsn,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prtry"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be one of LINK, UNLK, SOFT
type LinkageType1Code string

type LinkageType3Choice struct {
	Cd    LinkageType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prtry"`
}

type Linkages39 struct {
	PrcgPos ProcessingPosition8Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 PrcgPos,omitempty"`
	MsgNb   DocumentNumber5Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 MsgNb,omitempty"`
	Ref     References46Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Ref"`
	RefOwnr PartyIdentification92Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 RefOwnr,omitempty"`
}

type MatchingDenied3Choice struct {
	Cd    MatchingProcess1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prtry"`
}

// May be one of UNMT, MTRE
type MatchingProcess1Code string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Adr,omitempty"`
}

// May be one of NORE
type NoReasonCode string

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Tp"`
}

type PartyIdentification100 struct {
	Id  PartyIdentification71Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 LEI,omitempty"`
}

type PartyIdentification71Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 NmAndAdr"`
}

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 PrtryId"`
}

type PartyIdentification98 struct {
	Id  PartyIdentification92Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 LEI,omitempty"`
}

type PartyIdentificationAndAccount117 struct {
	Id        PartyIdentification71Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Id"`
	LEI       LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 LEI,omitempty"`
	SfkpgAcct SecuritiesAccount19         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 SfkpgAcct,omitempty"`
	PrcgId    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 PrcgId,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Ctry"`
}

type PriorityNumeric4Choice struct {
	Nmrc  Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Nmrc"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prtry"`
}

// May be one of AFTE, BEFO, WITH
type ProcessingPosition4Code string

type ProcessingPosition8Choice struct {
	Cd    ProcessingPosition4Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prtry"`
}

type ProcessingStatus56Choice struct {
	Rjctd  RejectedStatus17Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Rjctd"`
	Canc   CancelledStatus10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Canc"`
	Accptd AcceptedStatus7Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Accptd"`
}

type References18 struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AcctOwnrTxId,omitempty"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 PrcrTxId,omitempty"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 PoolId,omitempty"`
	CmonId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 CmonId,omitempty"`
	TradId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 TradId,omitempty"`
}

type References46Choice struct {
	SctiesSttlmTxId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 SctiesSttlmTxId"`
	IntraPosMvmntId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 IntraPosMvmntId"`
	IntraBalMvmntId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 IntraBalMvmntId"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AcctSvcrTxId"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 MktInfrstrctrTxId"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 PoolId"`
	CmonId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 CmonId"`
	TradId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 TradId"`
	OthrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 OthrTxId"`
}

type Registration10Choice struct {
	Cd    Registration2Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prtry"`
}

// May be one of PTYH, CSDH, CDEL, CVAL
type Registration2Code string

type RegistrationReason5 struct {
	Cd       Registration10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	AddtlInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AddtlInf,omitempty"`
}

type RejectedStatus17Choice struct {
	NoSpcfdRsn NoReasonCode        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 NoSpcfdRsn"`
	Rsn        []RejectionReason27 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Rsn"`
}

type RejectionReason25Choice struct {
	Cd    RejectionReason40Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prtry"`
}

type RejectionReason27 struct {
	Cd          RejectionReason25Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AddtlRsnInf,omitempty"`
}

// May be one of SETS, DDAT, CASY, DDEA, DEPT, DMON, DQUA, DSEC, DTRD, ICAG, ICUS, IEXE, NARR, NCRR, PLCE, RTGS, SETR
type RejectionReason40Code string

type RequestDetails15 struct {
	Ref          References18                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Ref"`
	AutomtcBrrwg AutomaticBorrowing7Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AutomtcBrrwg,omitempty"`
	RtnInd       bool                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 RtnInd,omitempty"`
	Lkg          LinkageType3Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Lkg,omitempty"`
	Prty         PriorityNumeric4Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prty,omitempty"`
	OthrPrcg     []GenericIdentification30           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 OthrPrcg,omitempty"`
	PrtlSttlmInd SettlementTransactionCondition5Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 PrtlSttlmInd,omitempty"`
	SctiesRTGS   SecuritiesRTGS4Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 SctiesRTGS,omitempty"`
	HldInd       HoldIndicator6                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 HldInd,omitempty"`
	MtchgDnl     MatchingDenied3Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 MtchgDnl,omitempty"`
	UnltrlSplt   UnilateralSplit3Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 UnltrlSplt,omitempty"`
	Lnkgs        []Linkages39                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Lnkgs,omitempty"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Nm,omitempty"`
}

type SecuritiesRTGS4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prtry"`
}

type SecuritiesSettlementConditionsModificationRequestV06 struct {
	AcctOwnr    PartyIdentification98     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount19       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 SfkpgAcct"`
	ReqDtls     []RequestDetails15        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 ReqDtls"`
	AddtlInf    []AdditionalInformation11 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 AddtlInf,omitempty"`
	SplmtryData []SupplementaryData1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 SplmtryData,omitempty"`
}

// May be one of TRAD
type SecuritiesTransactionType5Code string

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Desc,omitempty"`
}

// May be one of PART, NPAR, PARC, PARQ
type SettlementTransactionCondition5Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type UnilateralSplit3Choice struct {
	Cd    SecuritiesTransactionType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.06 Prtry"`
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

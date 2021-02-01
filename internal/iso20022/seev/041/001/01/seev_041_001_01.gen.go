// Code generated by main. DO NOT EDIT.

package seev_041_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AcceptedReason1Choice struct {
	Cd    AcknowledgementReason4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Prtry"`
}

type AcceptedStatus1Choice struct {
	NoSpcfdRsn NoReasonCode            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 NoSpcfdRsn"`
	Rsn        []AcceptedStatusReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Rsn"`
}

type AcceptedStatusReason1 struct {
	RsnCd       AcceptedReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 RsnCd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, LATE, NSTP, OTHR
type AcknowledgementReason4Code string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateIdentification1 struct {
	Id    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Id"`
	IdSrc IdentificationSource1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 IdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CancelledReason1Choice struct {
	Cd    CancelledStatusReason6Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Prtry"`
}

type CancelledStatus1Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 NoSpcfdRsn"`
	Rsn        []CancelledStatusReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Rsn"`
}

type CancelledStatusReason4 struct {
	RsnCd       CancelledReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 RsnCd"`
	AddtlRsnInf Max210Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AddtlRsnInf,omitempty"`
}

// May be one of CANI, CANO, CANS, CSUB, OTHR
type CancelledStatusReason6Code string

type CashAccountIdentification5Choice struct {
	IBAN  IBAN2007Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 IBAN"`
	Prtry Max34Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Prtry"`
}

// May be one of CODU, COPY, DUPL
type CopyDuplicate1Code string

type CorporateActionEventType3Choice struct {
	Cd    CorporateActionEventType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Cd"`
	Prtry GenericIdentification20       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Prtry"`
}

// May be one of ACTV, ATTI, BIDS, BONU, BPUT, BRUP, CAPG, CAPI, CERT, CHAN, CLSA, CONS, CONV, COOP, DECR, DETI, DFLT, DLST, DRAW, DRIP, DSCL, DTCH, DVCA, DVOP, DVSC, DVSE, EXOF, EXRI, EXTM, EXWA, CAPD, INCR, INTR, LIQU, MCAL, MRGR, ODLT, OTHR, PARI, PCAL, PDEF, PINK, PLAC, PPMT, PRED, PRII, PRIO, REDM, REDO, REMK, RHDI, RHTS, SHPR, SMAL, SOFF, SPLF, SPLR, SUSP, TEND, TREC, WRTH, WTRC, CREV
type CorporateActionEventType6Code string

type CorporateActionGeneralInformation9 struct {
	CorpActnEvtId      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 OffclCorpActnEvtId,omitempty"`
	ClssActnNb         Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 ClssActnNb,omitempty"`
	EvtTp              CorporateActionEventType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 EvtTp"`
}

type CorporateActionInstructionCancellationRequestStatusAdviceV01 struct {
	Id             DocumentIdentification11                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Id"`
	InstrCxlReqId  DocumentIdentification9                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 InstrCxlReqId,omitempty"`
	OthrDocId      []DocumentIdentification14                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 OthrDocId,omitempty"`
	CorpActnGnlInf CorporateActionGeneralInformation9            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 CorpActnGnlInf"`
	InstrCxlReqSts []InstructionCancellationRequestStatus1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 InstrCxlReqSts"`
	CorpActnInstr  CorporateActionOption9                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 CorpActnInstr,omitempty"`
	AddtlInf       CorporateActionNarrative10                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AddtlInf,omitempty"`
	MsgOrgtr       PartyIdentification10Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 MsgOrgtr,omitempty"`
	MsgRcpt        PartyIdentification10Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 MsgRcpt,omitempty"`
	Xtnsn          []Extension2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Xtnsn,omitempty"`
}

type CorporateActionNarrative10 struct {
	AddtlTxt     []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AddtlTxt,omitempty"`
	PtyCtctNrrtv []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 PtyCtctNrrtv,omitempty"`
}

type CorporateActionOption5Choice struct {
	Cd    CorporateActionOption6Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Prtry"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, TAXI
type CorporateActionOption6Code string

type CorporateActionOption9 struct {
	OptnNb           OptionNumber1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 OptnNb"`
	OptnTp           CorporateActionOption5Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 OptnTp"`
	AcctOwnr         PartyIdentification13Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AcctOwnr,omitempty"`
	SfkpgAcct        Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 SfkpgAcct,omitempty"`
	CshAcct          CashAccountIdentification5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 CshAcct,omitempty"`
	SfkpgPlc         SafekeepingPlaceFormat2Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 SfkpgPlc,omitempty"`
	SctyId           SecurityIdentification11         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 SctyId,omitempty"`
	TtlElgblBal      SignedQuantityFormat1            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 TtlElgblBal,omitempty"`
	InstdBal         SignedQuantityFormat1            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 InstdBal,omitempty"`
	UinstdBal        SignedQuantityFormat1            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 UinstdBal,omitempty"`
	StsQtyOrQtyToRcv StatusOrQuantityToReceive1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 StsQtyOrQtyToRcv,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 DtTm"`
}

type Document struct {
	CorpActnInstrCxlReqStsAdvc CorporateActionInstructionCancellationRequestStatusAdviceV01 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 CorpActnInstrCxlReqStsAdvc"`
}

type DocumentIdentification11 struct {
	Id       Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Id"`
	CreDtTm  DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 CreDtTm,omitempty"`
	CpyDplct CopyDuplicate1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 CpyDplct,omitempty"`
}

type DocumentIdentification14 struct {
	Id    DocumentIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Id"`
	DocNb DocumentNumber1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 DocNb,omitempty"`
}

type DocumentIdentification1Choice struct {
	AcctSvcrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AcctSvcrDocId"`
	AcctOwnrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AcctOwnrDocId"`
}

type DocumentIdentification9 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Id"`
}

type DocumentNumber1Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 LngNb"`
	PrtryNb GenericIdentification19           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type Extension2 struct {
	PlcAndNm   Max350Text         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 PlcAndNm,omitempty"`
	XtnsnEnvlp ExtensionEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 XtnsnEnvlp"`
}

type ExtensionEnvelope1 struct {
	Item string `xml:",any"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AmtsdVal"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Id,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

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

type IdentificationSource1Choice struct {
	Dmst  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Dmst"`
	Prtry Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Prtry"`
}

type InstructionCancellationRequestStatus1Choice struct {
	CxlCmpltd CancelledStatus1Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 CxlCmpltd"`
	Accptd    AcceptedStatus1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Accptd"`
	Rjctd     RejectedStatus1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Rjctd"`
	PdgCxl    PendingCancellationStatus1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 PdgCxl"`
	PrtrySts  ProprietaryStatusAndReason1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 PrtrySts"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Adr,omitempty"`
}

// May be one of NORE
type NoReasonCode string

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AmtsdVal"`
}

type PartyIdentification10Choice struct {
	BICOrBEI AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 BICOrBEI"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 NmAndAdr"`
}

type PartyIdentification13Choice struct {
	BICOrBEI AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 BICOrBEI"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 PrtryId"`
}

type PendingCancellationReason1Choice struct {
	Cd    PendingCancellationReason4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Prtry"`
}

// May be one of ADEA, DQUA, LATE, OTHR
type PendingCancellationReason4Code string

type PendingCancellationStatus1Choice struct {
	NotSpcfdRsn NoReasonCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 NotSpcfdRsn"`
	Rsn         []PendingCancellationStatusReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Rsn"`
}

type PendingCancellationStatusReason1 struct {
	RsnCd       PendingCancellationReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 RsnCd"`
	AddtlRsnInf Max210Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AddtlRsnInf,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Ctry"`
}

type ProprietaryQuantity2 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 QtyTp"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 SchmeNm,omitempty"`
}

type ProprietaryReason1 struct {
	Rsn         GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason1 struct {
	PrtrySts GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 PrtrySts"`
	PrtryRsn []ProprietaryReason1    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 PrtryRsn,omitempty"`
}

type Quantity2Choice struct {
	Qty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Qty"`
	PrtryQty ProprietaryQuantity2               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 PrtryQty"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 OrgnlAndCurFace"`
}

type RejectedReason1Choice struct {
	Cd    RejectionReason17Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Prtry"`
}

type RejectedStatus1Choice struct {
	NoSpcfdRsn NoReasonCode            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 NoSpcfdRsn"`
	Rsn        []RejectedStatusReason8 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Rsn"`
}

type RejectedStatusReason8 struct {
	RsnCd       RejectedReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 RsnCd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, CANC, DCAN, DPRG, DQUA, DSEC, EVNM, INIR, INTV, INVA, LACK, LATE, OTHR, NMTY, OPNM, OPTY, REFT, SAFE, ULNK, CERT
type RejectionReason17Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat2Choice struct {
	Id      SafekeepingPlaceTypeAndText2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Id"`
}

type SafekeepingPlaceTypeAndText2 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Id,omitempty"`
}

type SecurityIdentification11 struct {
	Id   SecurityIdentification11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Id"`
	Desc Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Desc,omitempty"`
}

type SecurityIdentification11Choice struct {
	ISIN   ISINIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 ISIN"`
	OthrId AlternateIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 OthrId"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat1 struct {
	ShrtLngPos ShortLong1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 ShrtLngPos"`
	QtyChc     Quantity2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 QtyChc"`
}

type StatusOrQuantityToReceive1Choice struct {
	StsQty   Quantity6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 StsQty"`
	QtyToRcv Quantity6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 QtyToRcv"`
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
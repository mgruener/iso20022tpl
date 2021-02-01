// Code generated by main. DO NOT EDIT.

package seev_034_001_11

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AcceptedReason10Choice struct {
	Cd    AcknowledgementReason7Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Prtry"`
}

type AcceptedStatus8Choice struct {
	NoSpcfdRsn NoReasonCode            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 NoSpcfdRsn"`
	Rsn        []AcceptedStatusReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Rsn"`
}

type AcceptedStatusReason9 struct {
	RsnCd       AcceptedReason10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 RsnCd"`
	AddtlRsnInf Max210Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AddtlRsnInf,omitempty"`
}

// May be one of NSTP, OTHR
type AcknowledgementReason7Code string

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type CancelledReason8Choice struct {
	Cd    CancelledStatusReason6Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Prtry"`
}

type CancelledStatus12Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 NoSpcfdRsn"`
	Rsn        []CancelledStatusReason11 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Rsn"`
}

type CancelledStatusReason11 struct {
	RsnCd       CancelledReason8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 RsnCd"`
	AddtlRsnInf Max210Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AddtlRsnInf,omitempty"`
}

// May be one of CANI, CANO, CANS, CSUB, OTHR
type CancelledStatusReason6Code string

type CashAccountIdentification5Choice struct {
	IBAN  IBAN2007Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 IBAN"`
	Prtry Max34Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Prtry"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH
type CorporateActionEventType29Code string

type CorporateActionEventType85Choice struct {
	Cd    CorporateActionEventType29Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Prtry"`
}

type CorporateActionGeneralInformation137 struct {
	CorpActnEvtId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 OffclCorpActnEvtId,omitempty"`
	ClssActnNb         Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 ClssActnNb,omitempty"`
	EvtTp              CorporateActionEventType85Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 EvtTp"`
}

type CorporateActionInstructionStatusAdviceV11 struct {
	InstrId        DocumentIdentification9               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 InstrId,omitempty"`
	OthrDocId      []DocumentIdentification33            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 OthrDocId,omitempty"`
	CorpActnGnlInf CorporateActionGeneralInformation137  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 CorpActnGnlInf"`
	InstrPrcgSts   []InstructionProcessingStatus38Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 InstrPrcgSts"`
	CorpActnInstr  CorporateActionOption167              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 CorpActnInstr,omitempty"`
	PrtctInstr     ProtectInstruction2                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 PrtctInstr,omitempty"`
	AddtlInf       CorporateActionNarrative10            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AddtlInf,omitempty"`
	SplmtryData    []SupplementaryData1                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 SplmtryData,omitempty"`
}

type CorporateActionNarrative10 struct {
	AddtlTxt     []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AddtlTxt,omitempty"`
	PtyCtctNrrtv []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 PtyCtctNrrtv,omitempty"`
}

// May be one of ABST, BSPL, BUYA, CASE, CASH, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, QINV, SECU, SLLE, TAXI, PRUN
type CorporateActionOption14Code string

type CorporateActionOption167 struct {
	OptnNb      OptionNumber1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 OptnNb"`
	OptnTp      CorporateActionOption32Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 OptnTp"`
	OptnFeatrs  OptionFeaturesFormat25Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 OptnFeatrs,omitempty"`
	AcctOwnr    PartyIdentification127Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AcctOwnr,omitempty"`
	SfkpgAcct   Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 SfkpgAcct,omitempty"`
	CshAcct     CashAccountIdentification5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 CshAcct,omitempty"`
	SfkpgPlc    SafekeepingPlaceFormat28Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 SfkpgPlc,omitempty"`
	FinInstrmId SecurityIdentification19         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 FinInstrmId,omitempty"`
	TtlElgblBal SignedQuantityFormat7            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 TtlElgblBal,omitempty"`
	InstdBal    SignedQuantityFormat7            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 InstdBal,omitempty"`
	UinstdBal   SignedQuantityFormat7            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 UinstdBal,omitempty"`
	StsQty      Quantity6Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 StsQty,omitempty"`
	StsCshAmt   ActiveCurrencyAndAmount          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 StsCshAmt,omitempty"`
}

type CorporateActionOption32Choice struct {
	Cd    CorporateActionOption14Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	CorpActnInstrStsAdvc CorporateActionInstructionStatusAdviceV11 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 CorpActnInstrStsAdvc"`
}

type DocumentIdentification33 struct {
	Id    DocumentIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Id"`
	DocNb DocumentNumber5Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 DocNb,omitempty"`
}

type DocumentIdentification3Choice struct {
	AcctSvcrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AcctSvcrDocId"`
	AcctOwnrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AcctOwnrDocId"`
}

type DocumentIdentification9 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Id"`
}

type DocumentNumber5Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 LngNb"`
	PrtryNb GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity18Choice struct {
	Unit    float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Unit"`
	FaceAmt float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 FaceAmt"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AmtsdVal"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Id,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Prtry"`
}

type InstructionProcessingStatus38Choice struct {
	Canc               CancelledStatus12Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Canc"`
	AccptdForFrthrPrcg AcceptedStatus8Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AccptdForFrthrPrcg"`
	Rjctd              RejectedStatus32Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Rjctd"`
	Pdg                PendingStatus57Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Pdg"`
	DfltActn           NoSpecifiedReason1          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 DfltActn"`
	StgInstr           NoSpecifiedReason1          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 StgInstr"`
	PrtrySts           ProprietaryStatusAndReason6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 PrtrySts"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max15Text string

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

// May be one of NORE
type NoReasonCode string

type NoSpecifiedReason1 struct {
	NoSpcfdRsn NoReasonCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 NoSpcfdRsn"`
}

// May be one of OPLF
type OptionFeatures12Code string

type OptionFeaturesFormat25Choice struct {
	Cd    OptionFeatures12Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Prtry"`
}

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Tp"`
}

type PartyIdentification127Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 PrtryId"`
}

// May be one of ADEA, OTHR, FULL, MCER, MONY, LACK, LATE, DQUA, PENR, CERT, DQCS, ITAX, NTAX, MTAX
type PendingReason18Code string

type PendingReason53Choice struct {
	Cd    PendingReason18Code     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Prtry"`
}

type PendingStatus57Choice struct {
	NoSpcfdRsn NoReasonCode            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 NoSpcfdRsn"`
	Rsn        []PendingStatusReason17 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Rsn"`
}

type PendingStatusReason17 struct {
	RsnCd       PendingReason53Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 RsnCd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AddtlRsnInf,omitempty"`
}

type ProprietaryQuantity8 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 QtyTp"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 SchmeNm,omitempty"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 PrtryRsn,omitempty"`
}

type ProtectInstruction2 struct {
	TxTp           ProtectTransactionType2Code         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 TxTp"`
	PrtctTxSts     ProtectInstructionStatus3Code       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 PrtctTxSts,omitempty"`
	TxId           Max15Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 TxId,omitempty"`
	PrtctSfkpgAcct Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 PrtctSfkpgAcct,omitempty"`
	PrtctDt        ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 PrtctDt,omitempty"`
	UcvrdPrtctQty  FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 UcvrdPrtctQty,omitempty"`
}

// May be one of OPEN, COVR, EXPI
type ProtectInstructionStatus3Code string

// May be one of PROT, COVP, COVR
type ProtectTransactionType2Code string

type Quantity19Choice struct {
	Qty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Qty"`
	PrtryQty ProprietaryQuantity8               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 PrtryQty"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 OrgnlAndCurFace"`
}

type RejectedReason31Choice struct {
	Cd    RejectionReason54Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Prtry"`
}

type RejectedStatus32Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 NoSpcfdRsn"`
	Rsn        []RejectedStatusReason30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Rsn"`
}

type RejectedStatusReason30 struct {
	RsnCd       RejectedReason31Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 RsnCd"`
	AddtlRsnInf Max210Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, CERT, INVA, OPTY, ULNK, DSEC, LACK, LATE, NMTY, FULL, CANC, INTV, OPNM, OTHR, DQUA, REFT, SAFE, EVNM, DQCS, DQCC, DQAM, IRDQ, DQBV, DQBI, SHAR, ITAX, NTAX, MTAX
type RejectionReason54Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat28Choice struct {
	Id      SafekeepingPlaceTypeAndText6           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Id"`
	Ctry    CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Ctry"`
	TpAndId SafekeepingPlaceTypeAndIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 TpAndId"`
	Prtry   GenericIdentification78                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Prtry"`
}

type SafekeepingPlaceTypeAndIdentification1 struct {
	SfkpgPlcTp SafekeepingPlace1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 SfkpgPlcTp"`
	Id         AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Id"`
}

type SafekeepingPlaceTypeAndText6 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Id,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Desc,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat7 struct {
	ShrtLngPos ShortLong1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 ShrtLngPos"`
	QtyChc     Quantity19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 QtyChc"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.11 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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

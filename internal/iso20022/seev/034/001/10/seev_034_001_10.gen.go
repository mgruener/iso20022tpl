// Code generated by main. DO NOT EDIT.

package seev_034_001_10

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AcceptedReason10Choice struct {
	Cd    AcknowledgementReason7Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Prtry"`
}

type AcceptedStatus8Choice struct {
	NoSpcfdRsn NoReasonCode            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 NoSpcfdRsn"`
	Rsn        []AcceptedStatusReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Rsn"`
}

type AcceptedStatusReason9 struct {
	RsnCd       AcceptedReason10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 RsnCd"`
	AddtlRsnInf Max210Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AddtlRsnInf,omitempty"`
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
	Cd    CancelledStatusReason6Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Prtry"`
}

type CancelledStatus12Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 NoSpcfdRsn"`
	Rsn        []CancelledStatusReason11 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Rsn"`
}

type CancelledStatusReason11 struct {
	RsnCd       CancelledReason8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 RsnCd"`
	AddtlRsnInf Max210Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AddtlRsnInf,omitempty"`
}

// May be one of CANI, CANO, CANS, CSUB, OTHR
type CancelledStatusReason6Code string

type CashAccountIdentification5Choice struct {
	IBAN  IBAN2007Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 IBAN"`
	Prtry Max34Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Prtry"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, PRII, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH
type CorporateActionEventType20Code string

type CorporateActionEventType52Choice struct {
	Cd    CorporateActionEventType20Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Prtry"`
}

type CorporateActionGeneralInformation109 struct {
	CorpActnEvtId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 OffclCorpActnEvtId,omitempty"`
	ClssActnNb         Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 ClssActnNb,omitempty"`
	EvtTp              CorporateActionEventType52Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 EvtTp"`
}

type CorporateActionInstructionStatusAdviceV10 struct {
	InstrId        DocumentIdentification9               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 InstrId,omitempty"`
	OthrDocId      []DocumentIdentification33            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 OthrDocId,omitempty"`
	CorpActnGnlInf CorporateActionGeneralInformation109  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 CorpActnGnlInf"`
	InstrPrcgSts   []InstructionProcessingStatus31Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 InstrPrcgSts"`
	CorpActnInstr  CorporateActionOption151              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 CorpActnInstr,omitempty"`
	PrtctInstr     ProtectInstruction2                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 PrtctInstr,omitempty"`
	AddtlInf       CorporateActionNarrative10            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AddtlInf,omitempty"`
	SplmtryData    []SupplementaryData1                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 SplmtryData,omitempty"`
}

type CorporateActionNarrative10 struct {
	AddtlTxt     []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AddtlTxt,omitempty"`
	PtyCtctNrrtv []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 PtyCtctNrrtv,omitempty"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, TAXI, PRUN
type CorporateActionOption10Code string

type CorporateActionOption151 struct {
	OptnNb      OptionNumber1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 OptnNb"`
	OptnTp      CorporateActionOption21Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 OptnTp"`
	OptnFeatrs  OptionFeaturesFormat25Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 OptnFeatrs,omitempty"`
	AcctOwnr    PartyIdentification127Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AcctOwnr,omitempty"`
	SfkpgAcct   Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 SfkpgAcct,omitempty"`
	CshAcct     CashAccountIdentification5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 CshAcct,omitempty"`
	SfkpgPlc    SafekeepingPlaceFormat28Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 SfkpgPlc,omitempty"`
	FinInstrmId SecurityIdentification19         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 FinInstrmId,omitempty"`
	TtlElgblBal SignedQuantityFormat7            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 TtlElgblBal,omitempty"`
	InstdBal    SignedQuantityFormat7            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 InstdBal,omitempty"`
	UinstdBal   SignedQuantityFormat7            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 UinstdBal,omitempty"`
	StsQty      Quantity6Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 StsQty,omitempty"`
	StsCshAmt   ActiveCurrencyAndAmount          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 StsCshAmt,omitempty"`
}

type CorporateActionOption21Choice struct {
	Cd    CorporateActionOption10Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	CorpActnInstrStsAdvc CorporateActionInstructionStatusAdviceV10 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 CorpActnInstrStsAdvc"`
}

type DocumentIdentification33 struct {
	Id    DocumentIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Id"`
	DocNb DocumentNumber5Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 DocNb,omitempty"`
}

type DocumentIdentification3Choice struct {
	AcctSvcrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AcctSvcrDocId"`
	AcctOwnrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AcctOwnrDocId"`
}

type DocumentIdentification9 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Id"`
}

type DocumentNumber5Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 LngNb"`
	PrtryNb GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity18Choice struct {
	Unit    float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Unit"`
	FaceAmt float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 FaceAmt"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AmtsdVal"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Id,omitempty"`
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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Prtry"`
}

type InstructionProcessingStatus31Choice struct {
	Canc               CancelledStatus12Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Canc"`
	AccptdForFrthrPrcg AcceptedStatus8Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AccptdForFrthrPrcg"`
	Rjctd              RejectedStatus25Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Rjctd"`
	Pdg                PendingStatus52Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Pdg"`
	DfltActn           NoSpecifiedReason1          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 DfltActn"`
	StgInstr           NoSpecifiedReason1          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 StgInstr"`
	PrtrySts           ProprietaryStatusAndReason6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 PrtrySts"`
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
	NoSpcfdRsn NoReasonCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 NoSpcfdRsn"`
}

// May be one of OPLF
type OptionFeatures12Code string

type OptionFeaturesFormat25Choice struct {
	Cd    OptionFeatures12Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Prtry"`
}

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Tp"`
}

type PartyIdentification127Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 PrtryId"`
}

// May be one of ADEA, OTHR, FULL, MCER, MONY, LACK, LATE, DQUA, PENR, CERT, DQCS
type PendingReason14Code string

type PendingReason48Choice struct {
	Cd    PendingReason14Code     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Prtry"`
}

type PendingStatus52Choice struct {
	NoSpcfdRsn NoReasonCode            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 NoSpcfdRsn"`
	Rsn        []PendingStatusReason13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Rsn"`
}

type PendingStatusReason13 struct {
	RsnCd       PendingReason48Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 RsnCd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AddtlRsnInf,omitempty"`
}

type ProprietaryQuantity8 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 QtyTp"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 SchmeNm,omitempty"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 PrtryRsn,omitempty"`
}

type ProtectInstruction2 struct {
	TxTp           ProtectTransactionType2Code         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 TxTp"`
	PrtctTxSts     ProtectInstructionStatus3Code       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 PrtctTxSts,omitempty"`
	TxId           Max15Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 TxId,omitempty"`
	PrtctSfkpgAcct Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 PrtctSfkpgAcct,omitempty"`
	PrtctDt        ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 PrtctDt,omitempty"`
	UcvrdPrtctQty  FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 UcvrdPrtctQty,omitempty"`
}

// May be one of OPEN, COVR, EXPI
type ProtectInstructionStatus3Code string

// May be one of PROT, COVP, COVR
type ProtectTransactionType2Code string

type Quantity19Choice struct {
	Qty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Qty"`
	PrtryQty ProprietaryQuantity8               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 PrtryQty"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 OrgnlAndCurFace"`
}

type RejectedReason24Choice struct {
	Cd    RejectionReason48Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Prtry"`
}

type RejectedStatus25Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 NoSpcfdRsn"`
	Rsn        []RejectedStatusReason23 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Rsn"`
}

type RejectedStatusReason23 struct {
	RsnCd       RejectedReason24Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 RsnCd"`
	AddtlRsnInf Max210Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, CERT, INVA, OPTY, ULNK, DSEC, LACK, LATE, NMTY, FULL, CANC, INTV, OPNM, OTHR, DQUA, REFT, SAFE, EVNM, DQCS, DQCC, DQAM, IRDQ, DQBV, DQBI, SHAR
type RejectionReason48Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat28Choice struct {
	Id      SafekeepingPlaceTypeAndText6           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Id"`
	Ctry    CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Ctry"`
	TpAndId SafekeepingPlaceTypeAndIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 TpAndId"`
	Prtry   GenericIdentification78                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Prtry"`
}

type SafekeepingPlaceTypeAndIdentification1 struct {
	SfkpgPlcTp SafekeepingPlace1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 SfkpgPlcTp"`
	Id         AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Id"`
}

type SafekeepingPlaceTypeAndText6 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Id,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Desc,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat7 struct {
	ShrtLngPos ShortLong1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 ShrtLngPos"`
	QtyChc     Quantity19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 QtyChc"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.10 Envlp"`
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

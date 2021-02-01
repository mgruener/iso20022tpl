// Code generated by main. DO NOT EDIT.

package seev_042_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification25 struct {
	SfkpgAcct         Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SfkpgAcct"`
	AcctOwnr          PartyIdentification36Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 AcctOwnr,omitempty"`
	SfkpgPlc          SafekeepingPlaceFormat2Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SfkpgPlc,omitempty"`
	CorpActnEvtAndBal []CorporateActionEventAndBalance7 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 CorpActnEvtAndBal,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type BalanceFormat1Choice struct {
	Bal         SignedQuantityFormat1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Bal"`
	ElgblBal    SignedQuantityFormat2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 ElgblBal"`
	NotElgblBal SignedQuantityFormat2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 NotElgblBal"`
}

type CorporateActionBalanceDetails9 struct {
	TtlElgblBal      Quantity3Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 TtlElgblBal"`
	UinstdBal        BalanceFormat1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 UinstdBal"`
	TtlInstdBalDtls  InstructedBalanceDetails3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 TtlInstdBalDtls"`
	BlckdBal         SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 BlckdBal,omitempty"`
	BrrwdBal         SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 BrrwdBal,omitempty"`
	CollInBal        SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 CollInBal,omitempty"`
	CollOutBal       SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 CollOutBal,omitempty"`
	OnLnBal          SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 OnLnBal,omitempty"`
	OutForRegnBal    SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 OutForRegnBal,omitempty"`
	SttlmPosBal      SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SttlmPosBal,omitempty"`
	StrtPosBal       SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 StrtPosBal,omitempty"`
	TradDtPosBal     SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 TradDtPosBal,omitempty"`
	InTrnsShipmntBal SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 InTrnsShipmntBal,omitempty"`
	RegdBal          SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 RegdBal,omitempty"`
	OblgtdBal        SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 OblgtdBal,omitempty"`
	PdgDlvryBal      []PendingBalance1         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 PdgDlvryBal,omitempty"`
	PdgRctBal        []PendingBalance1         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 PdgRctBal,omitempty"`
}

// May be one of MKDT, RDDT, EARD
type CorporateActionDeadline1Code string

type CorporateActionEventAndBalance7 struct {
	GnlInf      EventInformation5              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 GnlInf"`
	UndrlygScty SecurityIdentification14       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 UndrlygScty"`
	Bal         CorporateActionBalanceDetails9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Bal,omitempty"`
	SplmtryData []SupplementaryData1           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SplmtryData,omitempty"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, CLSA, COOP, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, RHDI, PRII, LIQU, EXTM, MRGR, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH, NOOF
type CorporateActionEventType10Code string

type CorporateActionEventType11Choice struct {
	Cd    CorporateActionEventType10Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Prtry"`
}

type CorporateActionInstructionStatementReportV04 struct {
	Pgntn           Pagination                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Pgntn"`
	StmtGnlDtls     Statement12               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 StmtGnlDtls"`
	AcctAndStmtDtls []AccountIdentification25 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 AcctAndStmtDtls"`
	SplmtryData     []SupplementaryData1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SplmtryData,omitempty"`
}

type CorporateActionMandatoryVoluntary1Choice struct {
	Cd    CorporateActionMandatoryVoluntary1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Cd"`
	Prtry GenericIdentification20                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Prtry"`
}

// May be one of MAND, CHOS, VOLU
type CorporateActionMandatoryVoluntary1Code string

type CorporateActionOption10Choice struct {
	Cd    CorporateActionOption7Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Prtry"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CEXC, CONN, CONY, CTEN, EXER, LAPS, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, PRUN
type CorporateActionOption7Code string

// May be one of MASE, SAME
type CorporateActionStatementReportingType1Code string

// May be one of MISS, ALLL
type CorporateActionStatementType1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 DtTm"`
}

type DateOrDateTimePeriodChoice struct {
	Dt   DatePeriodDetails     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Dt"`
	DtTm DateTimePeriodDetails `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 DtTm"`
}

type DatePeriodDetails struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 ToDt"`
}

type DateTimePeriodDetails struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 ToDtTm"`
}

type DeadlineCode1Choice struct {
	Cd    CorporateActionDeadline1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Cd"`
	Prtry GenericIdentification20      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Prtry"`
}

type DefaultProcessingOrStandingInstruction1Choice struct {
	DfltOptnInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 DfltOptnInd"`
	StgInstrInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 StgInstrInd"`
}

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	CorpActnInstrStmtRpt CorporateActionInstructionStatementReportV04 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 CorpActnInstrStmtRpt"`
}

// May be one of YEAR, ADHO, MNTH, DAIL, INDA, WEEK
type EventFrequency4Code string

type EventInformation5 struct {
	CorpActnEvtId      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 OffclCorpActnEvtId,omitempty"`
	EvtTp              CorporateActionEventType11Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 EvtTp"`
	MndtryVlntryEvtTp  CorporateActionMandatoryVoluntary1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 MndtryVlntryEvtTp"`
	LastNtfctnId       NotificationIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 LastNtfctnId,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 AmtsdVal"`
}

type Frequency4Choice struct {
	Cd    EventFrequency4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Prtry"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Id,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Prtry"`
}

type InstructedBalanceDetails3 struct {
	TtlInstdBal BalanceFormat1Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 TtlInstdBal"`
	OptnDtls    []InstructedCorporateActionOption4 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 OptnDtls,omitempty"`
}

type InstructedCorporateActionOption4 struct {
	OptnNb   Exact3NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 OptnNb,omitempty"`
	OptnTp   CorporateActionOption10Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 OptnTp"`
	InstdBal BalanceFormat1Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 InstdBal"`
	DfltActn DefaultProcessingOrStandingInstruction1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 DfltActn,omitempty"`
	DdlnDtTm ISODateTime                                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 DdlnDtTm"`
	DdlnTp   DeadlineCode1Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 DdlnTp"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

type NotificationIdentification1 struct {
	Id      Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Id"`
	CreDtTm DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 CreDtTm,omitempty"`
}

type OriginalAndCurrentQuantities2 struct {
	ShrtLngPos ShortLong1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 ShrtLngPos"`
	FaceAmt    float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 FaceAmt"`
	AmtsdVal   float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Tp"`
}

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 LastPgInd"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 PrtryId"`
}

type PendingBalance1 struct {
	Bal    SignedQuantityFormat2              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Bal"`
	PdgTxs []SettlementTypeAndIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 PdgTxs,omitempty"`
}

type ProprietaryQuantity2 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 QtyTp"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SchmeNm,omitempty"`
}

type ProprietaryQuantity3 struct {
	ShrtLngPos ShortLong1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 ShrtLngPos,omitempty"`
	Qty        float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Qty"`
	QtyTp      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 QtyTp"`
	Issr       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Issr"`
	SchmeNm    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SchmeNm,omitempty"`
}

type Quantity2Choice struct {
	Qty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Qty"`
	PrtryQty ProprietaryQuantity2               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 PrtryQty"`
}

type Quantity3Choice struct {
	QtyChc   Quantity4Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 QtyChc"`
	PrtryQty ProprietaryQuantity3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 PrtryQty"`
}

type Quantity4Choice struct {
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 OrgnlAndCurFaceAmt"`
	SgndQty            SignedQuantityFormat2         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SgndQty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat2Choice struct {
	Id      SafekeepingPlaceTypeAndText2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Id"`
}

type SafekeepingPlaceTypeAndText2 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Id,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Desc,omitempty"`
}

type SettlementTypeAndIdentification2 struct {
	Pmt     DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Pmt"`
	TxId    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 TxId"`
	SttlmDt DateAndDateTimeChoice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 SttlmDt,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat1 struct {
	ShrtLngPos ShortLong1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 ShrtLngPos"`
	QtyChc     Quantity2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 QtyChc"`
}

type SignedQuantityFormat2 struct {
	ShrtLngPos ShortLong1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 ShrtLngPos"`
	Qty        FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Qty"`
}

type Statement12 struct {
	StmtTp        CorporateActionStatementType1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 StmtTp"`
	RptgTp        CorporateActionStatementReportingType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 RptgTp"`
	StmtId        Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 StmtId"`
	RptNb         Max5NumericText                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 RptNb,omitempty"`
	StmtDtTm      DateAndDateTimeChoice                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 StmtDtTm"`
	Frqcy         Frequency4Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Frqcy"`
	UpdTp         UpdateType2Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 UpdTp"`
	ActvtyInd     bool                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 ActvtyInd"`
	NtfctnDdlnPrd DateOrDateTimePeriodChoice                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 NtfctnDdlnPrd,omitempty"`
}

// May be one of COMP, DELT
type StatementUpdateType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type UpdateType2Choice struct {
	Cd    StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Cd"`
	Prtry GenericIdentification20  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Prtry"`
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
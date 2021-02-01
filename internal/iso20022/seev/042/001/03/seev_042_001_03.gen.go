// Code generated by main. DO NOT EDIT.

package seev_042_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification23 struct {
	SfkpgAcct         Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SfkpgAcct"`
	AcctOwnr          PartyIdentification36Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 AcctOwnr,omitempty"`
	SfkpgPlc          SafekeepingPlaceFormat2Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SfkpgPlc,omitempty"`
	CorpActnEvtAndBal []CorporateActionEventAndBalance5 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 CorpActnEvtAndBal,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type BalanceFormat1Choice struct {
	Bal         SignedQuantityFormat1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Bal"`
	ElgblBal    SignedQuantityFormat2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 ElgblBal"`
	NotElgblBal SignedQuantityFormat2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 NotElgblBal"`
}

type CorporateActionBalanceDetails9 struct {
	TtlElgblBal      Quantity3Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 TtlElgblBal"`
	UinstdBal        BalanceFormat1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 UinstdBal"`
	TtlInstdBalDtls  InstructedBalanceDetails3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 TtlInstdBalDtls"`
	BlckdBal         SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 BlckdBal,omitempty"`
	BrrwdBal         SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 BrrwdBal,omitempty"`
	CollInBal        SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 CollInBal,omitempty"`
	CollOutBal       SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 CollOutBal,omitempty"`
	OnLnBal          SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 OnLnBal,omitempty"`
	OutForRegnBal    SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 OutForRegnBal,omitempty"`
	SttlmPosBal      SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SttlmPosBal,omitempty"`
	StrtPosBal       SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 StrtPosBal,omitempty"`
	TradDtPosBal     SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 TradDtPosBal,omitempty"`
	InTrnsShipmntBal SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 InTrnsShipmntBal,omitempty"`
	RegdBal          SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 RegdBal,omitempty"`
	OblgtdBal        SignedQuantityFormat2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 OblgtdBal,omitempty"`
	PdgDlvryBal      []PendingBalance1         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 PdgDlvryBal,omitempty"`
	PdgRctBal        []PendingBalance1         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 PdgRctBal,omitempty"`
}

// May be one of MKDT, RDDT, EARD
type CorporateActionDeadline1Code string

type CorporateActionEventAndBalance5 struct {
	GnlInf      EventInformation3              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 GnlInf"`
	UndrlygScty SecurityIdentification14       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 UndrlygScty"`
	Bal         CorporateActionBalanceDetails9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Bal,omitempty"`
	SplmtryData []SupplementaryData1           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SplmtryData,omitempty"`
}

type CorporateActionEventType7Choice struct {
	Cd    CorporateActionEventType8Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Cd"`
	Prtry GenericIdentification20       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Prtry"`
}

// May be one of ACTV, ATTI, BIDS, BONU, BPUT, BRUP, CAPG, CAPI, CERT, CHAN, CLSA, CONS, CONV, COOP, DECR, DETI, DFLT, DLST, DRAW, DRIP, DSCL, DTCH, DVCA, DVOP, DVSC, DVSE, EXOF, EXRI, EXTM, EXWA, CAPD, INCR, INTR, LIQU, MCAL, MRGR, ODLT, OTHR, PARI, PCAL, PDEF, PINK, PLAC, PPMT, PRED, PRII, PRIO, REDM, REDO, REMK, RHDI, RHTS, SHPR, SMAL, SOFF, SPLF, SPLR, SUSP, TEND, TREC, WRTH, WTRC, CREV, DRCA
type CorporateActionEventType8Code string

type CorporateActionInstructionStatementReportV03 struct {
	Pgntn           Pagination                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Pgntn"`
	StmtGnlDtls     Statement12               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 StmtGnlDtls"`
	AcctAndStmtDtls []AccountIdentification23 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 AcctAndStmtDtls"`
	SplmtryData     []SupplementaryData1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SplmtryData,omitempty"`
}

type CorporateActionMandatoryVoluntary1Choice struct {
	Cd    CorporateActionMandatoryVoluntary1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Cd"`
	Prtry GenericIdentification20                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Prtry"`
}

// May be one of MAND, CHOS, VOLU
type CorporateActionMandatoryVoluntary1Code string

type CorporateActionOption10Choice struct {
	Cd    CorporateActionOption7Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Prtry"`
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
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 DtTm"`
}

type DateOrDateTimePeriodChoice struct {
	Dt   DatePeriodDetails     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Dt"`
	DtTm DateTimePeriodDetails `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 DtTm"`
}

type DatePeriodDetails struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 ToDt"`
}

type DateTimePeriodDetails struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 ToDtTm"`
}

type DeadlineCode1Choice struct {
	Cd    CorporateActionDeadline1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Cd"`
	Prtry GenericIdentification20      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Prtry"`
}

type DefaultProcessingOrStandingInstruction1Choice struct {
	DfltOptnInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 DfltOptnInd"`
	StgInstrInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 StgInstrInd"`
}

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	CorpActnInstrStmtRpt CorporateActionInstructionStatementReportV03 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 CorpActnInstrStmtRpt"`
}

// May be one of YEAR, ADHO, MNTH, DAIL, INDA, WEEK
type EventFrequency4Code string

type EventInformation3 struct {
	CorpActnEvtId      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 OffclCorpActnEvtId,omitempty"`
	EvtTp              CorporateActionEventType7Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 EvtTp"`
	MndtryVlntryEvtTp  CorporateActionMandatoryVoluntary1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 MndtryVlntryEvtTp"`
	LastNtfctnId       NotificationIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 LastNtfctnId,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 AmtsdVal"`
}

type Frequency4Choice struct {
	Cd    EventFrequency4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Prtry"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Id,omitempty"`
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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Prtry"`
}

type InstructedBalanceDetails3 struct {
	TtlInstdBal BalanceFormat1Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 TtlInstdBal"`
	OptnDtls    []InstructedCorporateActionOption4 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 OptnDtls,omitempty"`
}

type InstructedCorporateActionOption4 struct {
	OptnNb   Exact3NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 OptnNb,omitempty"`
	OptnTp   CorporateActionOption10Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 OptnTp"`
	InstdBal BalanceFormat1Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 InstdBal"`
	DfltActn DefaultProcessingOrStandingInstruction1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 DfltActn,omitempty"`
	DdlnDtTm ISODateTime                                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 DdlnDtTm"`
	DdlnTp   DeadlineCode1Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 DdlnTp"`
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
	Id      Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Id"`
	CreDtTm DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 CreDtTm,omitempty"`
}

type OriginalAndCurrentQuantities2 struct {
	ShrtLngPos ShortLong1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 ShrtLngPos"`
	FaceAmt    float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 FaceAmt"`
	AmtsdVal   float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Tp"`
}

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 LastPgInd"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 PrtryId"`
}

type PendingBalance1 struct {
	Bal    SignedQuantityFormat2              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Bal"`
	PdgTxs []SettlementTypeAndIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 PdgTxs,omitempty"`
}

type ProprietaryQuantity2 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 QtyTp"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SchmeNm,omitempty"`
}

type ProprietaryQuantity3 struct {
	ShrtLngPos ShortLong1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 ShrtLngPos,omitempty"`
	Qty        float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Qty"`
	QtyTp      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 QtyTp"`
	Issr       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Issr"`
	SchmeNm    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SchmeNm,omitempty"`
}

type Quantity2Choice struct {
	Qty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Qty"`
	PrtryQty ProprietaryQuantity2               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 PrtryQty"`
}

type Quantity3Choice struct {
	QtyChc   Quantity4Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 QtyChc"`
	PrtryQty ProprietaryQuantity3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 PrtryQty"`
}

type Quantity4Choice struct {
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 OrgnlAndCurFaceAmt"`
	SgndQty            SignedQuantityFormat2         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SgndQty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat2Choice struct {
	Id      SafekeepingPlaceTypeAndText2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Id"`
}

type SafekeepingPlaceTypeAndText2 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Id,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Desc,omitempty"`
}

type SettlementTypeAndIdentification2 struct {
	Pmt     DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Pmt"`
	TxId    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 TxId"`
	SttlmDt DateAndDateTimeChoice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 SttlmDt,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat1 struct {
	ShrtLngPos ShortLong1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 ShrtLngPos"`
	QtyChc     Quantity2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 QtyChc"`
}

type SignedQuantityFormat2 struct {
	ShrtLngPos ShortLong1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 ShrtLngPos"`
	Qty        FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Qty"`
}

type Statement12 struct {
	StmtTp        CorporateActionStatementType1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 StmtTp"`
	RptgTp        CorporateActionStatementReportingType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 RptgTp"`
	StmtId        Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 StmtId"`
	RptNb         Max5NumericText                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 RptNb,omitempty"`
	StmtDtTm      DateAndDateTimeChoice                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 StmtDtTm"`
	Frqcy         Frequency4Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Frqcy"`
	UpdTp         UpdateType2Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 UpdTp"`
	ActvtyInd     bool                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 ActvtyInd"`
	NtfctnDdlnPrd DateOrDateTimePeriodChoice                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 NtfctnDdlnPrd,omitempty"`
}

// May be one of COMP, DELT
type StatementUpdateType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type UpdateType2Choice struct {
	Cd    StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Cd"`
	Prtry GenericIdentification20  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.03 Prtry"`
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
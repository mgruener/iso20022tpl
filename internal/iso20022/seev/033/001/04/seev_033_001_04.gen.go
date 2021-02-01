// Code generated by main. DO NOT EDIT.

package seev_033_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountAndBalance17 struct {
	SfkpgAcct Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SfkpgAcct"`
	AcctOwnr  PartyIdentification36Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat2Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SfkpgPlc,omitempty"`
	Bal       CorporateActionBalanceDetails12 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Bal,omitempty"`
}

type ActiveCurrencyAnd13DecimalAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternatePartyIdentification2 struct {
	IdTp    IdentificationType4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 IdTp"`
	Ctry    CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Ctry"`
	AltrnId Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AltrnId"`
}

type AmountPrice3 struct {
	AmtPricTp AmountPriceType1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AmtPricTp"`
	PricVal   ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PricVal"`
}

type AmountPricePerAmount2 struct {
	AmtPricTp AmountPriceType1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AmtPricTp"`
	PricVal   ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PricVal"`
	Amt       ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Amt"`
}

type AmountPricePerFinancialInstrumentQuantity3 struct {
	AmtPricTp    AmountPriceType1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AmtPricTp"`
	PricVal      ActiveCurrencyAnd13DecimalAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PricVal"`
	FinInstrmQty FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 FinInstrmQty"`
}

// May be one of ACTU, DISC, PLOT, PREM
type AmountPriceType1Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type BalanceFormat1Choice struct {
	Bal         SignedQuantityFormat1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Bal"`
	ElgblBal    SignedQuantityFormat2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ElgblBal"`
	NotElgblBal SignedQuantityFormat2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 NotElgblBal"`
}

// May be one of ACCI, NCOM, QIBB
type BeneficiaryCertificationType5Code string

type BeneficiaryCertificationType6Choice struct {
	Cd    BeneficiaryCertificationType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Cd"`
	Prtry GenericIdentification20           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Prtry"`
}

// Must match the pattern [A-Z]{1,6}
type CFIIdentifier string

type ClassificationType2Choice struct {
	ClssfctnFinInstrm CFIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AltrnClssfctn"`
}

type CorporateActionBalanceDetails12 struct {
	TtlElgblBal      Quantity3Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 TtlElgblBal,omitempty"`
	BlckdBal         BalanceFormat1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 BlckdBal,omitempty"`
	BrrwdBal         BalanceFormat1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 BrrwdBal,omitempty"`
	CollInBal        BalanceFormat1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CollInBal,omitempty"`
	CollOutBal       BalanceFormat1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CollOutBal,omitempty"`
	OnLnBal          BalanceFormat1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OnLnBal,omitempty"`
	PdgDlvryBal      []BalanceFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PdgDlvryBal,omitempty"`
	PdgRctBal        []BalanceFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PdgRctBal,omitempty"`
	OutForRegnBal    BalanceFormat1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OutForRegnBal,omitempty"`
	SttlmPosBal      BalanceFormat1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SttlmPosBal,omitempty"`
	StrtPosBal       BalanceFormat1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 StrtPosBal,omitempty"`
	TradDtPosBal     BalanceFormat1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 TradDtPosBal,omitempty"`
	InTrnsShipmntBal BalanceFormat1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 InTrnsShipmntBal,omitempty"`
	RegdBal          BalanceFormat1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 RegdBal,omitempty"`
}

// May be one of BERE, CERT, DEPH, GPPH, GTGP, GTPH, NAME, PHDE, REBE, TERM
type CorporateActionChangeType2Code string

type CorporateActionChangeTypeFormat2Choice struct {
	Cd    CorporateActionChangeType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Prtry"`
}

type CorporateActionEventReference1 struct {
	EvtId CorporateActionEventReference1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 EvtId"`
	LkgTp ProcessingPosition1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 LkgTp,omitempty"`
}

type CorporateActionEventReference1Choice struct {
	LkdOffclCorpActnEvtId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 LkdOffclCorpActnEvtId"`
	LkdCorpActnId         Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 LkdCorpActnId"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, CLSA, COOP, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, RHDI, PRII, LIQU, EXTM, MRGR, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH, NOOF
type CorporateActionEventType10Code string

type CorporateActionEventType11Choice struct {
	Cd    CorporateActionEventType10Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Prtry"`
}

type CorporateActionGeneralInformation55 struct {
	CorpActnEvtId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OffclCorpActnEvtId,omitempty"`
	EvtTp              CorporateActionEventType11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 EvtTp"`
	UndrlygScty        FinancialInstrumentAttributes32  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 UndrlygScty,omitempty"`
}

type CorporateActionInstructionV04 struct {
	ChngInstrInd   bool                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ChngInstrInd,omitempty"`
	CancInstrId    DocumentIdentification15            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CancInstrId,omitempty"`
	InstrCxlReqId  DocumentIdentification15            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 InstrCxlReqId,omitempty"`
	OthrDocId      []DocumentIdentification13          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OthrDocId,omitempty"`
	EvtsLkg        []CorporateActionEventReference1    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 EvtsLkg,omitempty"`
	CorpActnGnlInf CorporateActionGeneralInformation55 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CorpActnGnlInf"`
	AcctDtls       AccountAndBalance17                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AcctDtls"`
	BnfclOwnrDtls  []PartyIdentification56             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 BnfclOwnrDtls,omitempty"`
	CorpActnInstr  CorporateActionOption57             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CorpActnInstr"`
	AddtlInf       CorporateActionNarrative21          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AddtlInf,omitempty"`
	SplmtryData    []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SplmtryData,omitempty"`
}

type CorporateActionNarrative21 struct {
	AddtlTxt       []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AddtlTxt,omitempty"`
	NrrtvVrsn      []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 NrrtvVrsn,omitempty"`
	RegnDtls       []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 RegnDtls,omitempty"`
	PtyCtctNrrtv   []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PtyCtctNrrtv,omitempty"`
	Dsclmr         []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Dsclmr,omitempty"`
	BsktOrIndxInf  []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 BsktOrIndxInf,omitempty"`
	CertfctnBrkdwn []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CertfctnBrkdwn,omitempty"`
}

type CorporateActionNarrative8 struct {
	InfToCmplyWth    []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 InfToCmplyWth,omitempty"`
	DlvryDtls        []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 DlvryDtls,omitempty"`
	FXInstrsAddtlInf []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 FXInstrsAddtlInf,omitempty"`
	Dsclmr           []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Dsclmr,omitempty"`
	InstrAddtlInf    []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 InstrAddtlInf,omitempty"`
}

type CorporateActionOption12Choice struct {
	Cd    CorporateActionOption9Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Prtry"`
}

type CorporateActionOption57 struct {
	OptnNb          OptionNumber1Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OptnNb"`
	OptnTp          CorporateActionOption12Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OptnTp"`
	FrctnDspstn     FractionDispositionType17Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 FrctnDspstn,omitempty"`
	ChngTp          []CorporateActionChangeTypeFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ChngTp,omitempty"`
	ElgblForCollInd bool                                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ElgblForCollInd,omitempty"`
	CcyToBuy        ActiveCurrencyCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CcyToBuy,omitempty"`
	CcyToSell       ActiveCurrencyCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CcyToSell,omitempty"`
	CcyOptn         ActiveCurrencyCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CcyOptn,omitempty"`
	SctyId          SecurityIdentification14                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SctyId,omitempty"`
	SctiesQty       SecuritiesOption2                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SctiesQty"`
	ExctnReqdDtTm   DateAndDateTimeChoice                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ExctnReqdDtTm,omitempty"`
	RateAndAmtDtls  CorporateActionRate8                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 RateAndAmtDtls,omitempty"`
	PricDtls        CorporateActionPrice29                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PricDtls,omitempty"`
	AddtlInf        CorporateActionNarrative8                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AddtlInf,omitempty"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CERT, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, TAXI, PRUN
type CorporateActionOption9Code string

type CorporateActionPrice29 struct {
	IndctvOrMktPric       IndicativeOrMarketPrice2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 IndctvOrMktPric,omitempty"`
	IssePric              PriceFormat5Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 IssePric,omitempty"`
	GncCshPricRcvdPerPdct PriceFormat21Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 GncCshPricRcvdPerPdct,omitempty"`
	GncCshPricPdPerPdct   PriceFormat5Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 GncCshPricPdPerPdct,omitempty"`
}

type CorporateActionRate8 struct {
	PropsdRate    float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PropsdRate,omitempty"`
	OvrsbcptRate  RateAndAmountFormat12Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OvrsbcptRate,omitempty"`
	ReqdTaxtnRate float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ReqdTaxtnRate,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 DtTm"`
}

type Document struct {
	CorpActnInstr CorporateActionInstructionV04 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CorpActnInstr"`
}

type DocumentIdentification13 struct {
	Id    DocumentIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Id"`
	DocNb DocumentNumber1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 DocNb,omitempty"`
	LkgTp ProcessingPosition1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 LkgTp,omitempty"`
}

type DocumentIdentification15 struct {
	Id    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Id"`
	LkgTp ProcessingPosition1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 LkgTp,omitempty"`
}

type DocumentIdentification1Choice struct {
	AcctSvcrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AcctSvcrDocId"`
	AcctOwnrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AcctOwnrDocId"`
}

type DocumentNumber1Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 LngNb"`
	PrtryNb GenericIdentification19           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentAttributes32 struct {
	FinInstrmId       SecurityIdentification14               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 FinInstrmId,omitempty"`
	PlcOfListg        MarketIdentification3Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PlcOfListg,omitempty"`
	DayCntBsis        InterestComputationMethodFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 DayCntBsis,omitempty"`
	ClssfctnTp        ClassificationType2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ClssfctnTp,omitempty"`
	DnmtnCcy          ActiveOrHistoricCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 DnmtnCcy,omitempty"`
	NxtCpnDt          ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 NxtCpnDt,omitempty"`
	XpryDt            ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 XpryDt,omitempty"`
	FltgRateFxgDt     ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 FltgRateFxgDt,omitempty"`
	MtrtyDt           ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 MtrtyDt,omitempty"`
	IsseDt            ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 IsseDt,omitempty"`
	NxtCllblDt        ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 NxtCllblDt,omitempty"`
	PutblDt           ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PutblDt,omitempty"`
	DtdDt             ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 DtdDt,omitempty"`
	ConvsDt           ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ConvsDt,omitempty"`
	PrvsFctr          float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PrvsFctr,omitempty"`
	NxtFctr           float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 NxtFctr,omitempty"`
	IntrstRate        float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 IntrstRate,omitempty"`
	NxtIntrstRate     float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 NxtIntrstRate,omitempty"`
	MinNmnlQty        FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 MinNmnlQty,omitempty"`
	MinExrcblQty      FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 MinExrcblQty,omitempty"`
	MinExrcblMltplQty FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 MinExrcblMltplQty,omitempty"`
	CtrctSz           FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CtrctSz,omitempty"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AmtsdVal"`
}

// May be one of BUYU, CINL, EXPI, DIST
type FractionDispositionType10Code string

type FractionDispositionType17Choice struct {
	Cd    FractionDispositionType10Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Cd"`
	Prtry GenericIdentification20       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Prtry"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Id,omitempty"`
}

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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Prtry"`
}

type IdentificationType4Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Cd"`
	Prtry GenericIdentification20   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Prtry"`
}

type IndicativeOrMarketPrice2Choice struct {
	IndctvPric PriceFormat5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 IndctvPric"`
	MktPric    PriceFormat5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 MktPric"`
}

type InstructedOrQuantityToReceive1Choice struct {
	InstdQty Quantity5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 InstdQty"`
	QtyToRcv Quantity5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 QtyToRcv"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat1Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MarketIdentification3Choice struct {
	MktIdrCd MICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 MktIdrCd"`
	Desc     Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Desc"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Adr,omitempty"`
}

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AmtsdVal"`
}

type OriginalAndCurrentQuantities2 struct {
	ShrtLngPos ShortLong1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ShrtLngPos"`
	FaceAmt    float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 FaceAmt"`
	AmtsdVal   float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Tp"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PrtryId"`
}

type PartyIdentification48Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 NmAndAdr"`
}

type PartyIdentification56 struct {
	OwnrId        PartyIdentification48Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OwnrId"`
	AltrnId       []AlternatePartyIdentification2       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AltrnId,omitempty"`
	DmclCtry      CountryCode                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 DmclCtry,omitempty"`
	NonDmclCtry   []CountryCode                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 NonDmclCtry,omitempty"`
	OwndSctiesQty FinancialInstrumentQuantity1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OwndSctiesQty"`
	CertfctnTp    []BeneficiaryCertificationType6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CertfctnTp,omitempty"`
	DclrtnDtls    Max350Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 DclrtnDtls,omitempty"`
}

type PercentagePrice1 struct {
	PctgPricTp PriceRateType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PctgPricTp"`
	PricVal    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PricVal"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Ctry"`
}

type PriceFormat21Choice struct {
	PctgPric               PercentagePrice1                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PctgPric"`
	AmtPric                AmountPrice3                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AmtPric"`
	NotSpcfdPric           PriceValueType9Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 NotSpcfdPric"`
	AmtPricPerFinInstrmQty AmountPricePerFinancialInstrumentQuantity3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AmtPricPerFinInstrmQty"`
	AmtPricPerAmt          AmountPricePerAmount2                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AmtPricPerAmt"`
}

type PriceFormat5Choice struct {
	PctgPric PercentagePrice1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PctgPric"`
	AmtPric  AmountPrice3     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 AmtPric"`
}

// May be one of DISC, PREM, PRCT, YIEL
type PriceRateType3Code string

// May be one of TBSP, UNSP, UKWN
type PriceValueType9Code string

type ProcessingPosition1Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Prtry"`
}

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

type ProprietaryQuantity2 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 QtyTp"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SchmeNm,omitempty"`
}

type ProprietaryQuantity3 struct {
	ShrtLngPos ShortLong1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ShrtLngPos,omitempty"`
	Qty        float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Qty"`
	QtyTp      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 QtyTp"`
	Issr       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Issr"`
	SchmeNm    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SchmeNm,omitempty"`
}

// May be one of QALL
type Quantity1Code string

type Quantity2Choice struct {
	Qty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Qty"`
	PrtryQty ProprietaryQuantity2               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PrtryQty"`
}

type Quantity3Choice struct {
	QtyChc   Quantity4Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 QtyChc"`
	PrtryQty ProprietaryQuantity3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PrtryQty"`
}

type Quantity4Choice struct {
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OrgnlAndCurFaceAmt"`
	SgndQty            SignedQuantityFormat2         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SgndQty"`
}

type Quantity5Choice struct {
	Cd                 Quantity1Code                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Cd"`
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OrgnlAndCurFaceAmt"`
	Qty                FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Qty"`
}

type RateAndAmountFormat12Choice struct {
	Rate float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Rate"`
	Amt  ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Amt"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat2Choice struct {
	Id      SafekeepingPlaceTypeAndText2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Id"`
}

type SafekeepingPlaceTypeAndText2 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Id,omitempty"`
}

type SecuritiesOption2 struct {
	CondlQty                        FinancialInstrumentQuantity1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 CondlQty,omitempty"`
	OverAndAbovNrmlNsrdEntitlmntQty FinancialInstrumentQuantity1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OverAndAbovNrmlNsrdEntitlmntQty,omitempty"`
	InstdOrQtyToRcv                 InstructedOrQuantityToReceive1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 InstdOrQtyToRcv"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Desc,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat1 struct {
	ShrtLngPos ShortLong1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ShrtLngPos"`
	QtyChc     Quantity2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 QtyChc"`
}

type SignedQuantityFormat2 struct {
	ShrtLngPos ShortLong1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 ShrtLngPos"`
	Qty        FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Qty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of ARNU, CCPT, CHTY, CORP, DRLC, FIIN, TXID
type TypeOfIdentification1Code string

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

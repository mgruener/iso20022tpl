// Code generated by main. DO NOT EDIT.

package seev_033_001_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountAndBalance35 struct {
	SfkpgAcct Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SfkpgAcct"`
	AcctOwnr  PartyIdentification92Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat8Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SfkpgPlc,omitempty"`
	Bal       CorporateActionBalanceDetails32 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Bal,omitempty"`
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

type AlternatePartyIdentification7 struct {
	IdTp    IdentificationType42Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 IdTp"`
	Ctry    CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Ctry"`
	AltrnId Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AltrnId"`
}

type AmountPrice3 struct {
	AmtPricTp AmountPriceType1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AmtPricTp"`
	PricVal   ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PricVal"`
}

type AmountPricePerAmount2 struct {
	AmtPricTp AmountPriceType1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AmtPricTp"`
	PricVal   ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PricVal"`
	Amt       ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Amt"`
}

type AmountPricePerFinancialInstrumentQuantity6 struct {
	AmtPricTp    AmountPriceType1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AmtPricTp"`
	PricVal      ActiveCurrencyAnd13DecimalAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PricVal"`
	FinInstrmQty FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 FinInstrmQty"`
}

// May be one of ACTU, DISC, PLOT, PREM
type AmountPriceType1Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type BalanceFormat5Choice struct {
	Bal         SignedQuantityFormat7 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Bal"`
	ElgblBal    SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ElgblBal"`
	NotElgblBal SignedQuantityFormat6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 NotElgblBal"`
}

type BeneficiaryCertificationType10Choice struct {
	Cd    BeneficiaryCertificationType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Cd"`
	Prtry GenericIdentification30           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Prtry"`
}

// May be one of ACCI, NCOM, QIBB
type BeneficiaryCertificationType5Code string

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

type ClassificationType32Choice struct {
	ClssfctnFinInstrm CFIOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AltrnClssfctn"`
}

type CorporateActionBalanceDetails32 struct {
	TtlElgblBal      Quantity17Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 TtlElgblBal,omitempty"`
	BlckdBal         BalanceFormat5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 BlckdBal,omitempty"`
	BrrwdBal         BalanceFormat5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 BrrwdBal,omitempty"`
	CollInBal        BalanceFormat5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CollInBal,omitempty"`
	CollOutBal       BalanceFormat5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CollOutBal,omitempty"`
	OnLnBal          BalanceFormat5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 OnLnBal,omitempty"`
	PdgDlvryBal      []BalanceFormat5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PdgDlvryBal,omitempty"`
	PdgRctBal        []BalanceFormat5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PdgRctBal,omitempty"`
	OutForRegnBal    BalanceFormat5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 OutForRegnBal,omitempty"`
	SttlmPosBal      BalanceFormat5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SttlmPosBal,omitempty"`
	StrtPosBal       BalanceFormat5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 StrtPosBal,omitempty"`
	TradDtPosBal     BalanceFormat5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 TradDtPosBal,omitempty"`
	InTrnsShipmntBal BalanceFormat5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 InTrnsShipmntBal,omitempty"`
	RegdBal          BalanceFormat5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 RegdBal,omitempty"`
}

// May be one of BERE, CERT, DEPH, GPPH, GTGP, GTPH, NAME, PHDE, REBE, TERM
type CorporateActionChangeType2Code string

type CorporateActionChangeTypeFormat6Choice struct {
	Cd    CorporateActionChangeType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Prtry"`
}

type CorporateActionEventReference3 struct {
	EvtId CorporateActionEventReference3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 EvtId"`
	LkgTp ProcessingPosition7Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 LkgTp,omitempty"`
}

type CorporateActionEventReference3Choice struct {
	LkdOffclCorpActnEvtId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 LkdOffclCorpActnEvtId"`
	LkdCorpActnId         Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 LkdCorpActnId"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, CLSA, COOP, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, PRII, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH
type CorporateActionEventType15Code string

type CorporateActionEventType32Choice struct {
	Cd    CorporateActionEventType15Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Prtry"`
}

type CorporateActionGeneralInformation88 struct {
	CorpActnEvtId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 OffclCorpActnEvtId,omitempty"`
	EvtTp              CorporateActionEventType32Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 EvtTp"`
	UndrlygScty        FinancialInstrumentAttributes65  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 UndrlygScty,omitempty"`
}

type CorporateActionInstructionV06 struct {
	ChngInstrInd   bool                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ChngInstrInd,omitempty"`
	CancInstrId    DocumentIdentification31            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CancInstrId,omitempty"`
	InstrCxlReqId  DocumentIdentification31            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 InstrCxlReqId,omitempty"`
	OthrDocId      []DocumentIdentification32          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 OthrDocId,omitempty"`
	EvtsLkg        []CorporateActionEventReference3    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 EvtsLkg,omitempty"`
	CorpActnGnlInf CorporateActionGeneralInformation88 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CorpActnGnlInf"`
	AcctDtls       AccountAndBalance35                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AcctDtls"`
	BnfclOwnrDtls  []PartyIdentification93             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 BnfclOwnrDtls,omitempty"`
	CorpActnInstr  CorporateActionOption118            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CorpActnInstr"`
	AddtlInf       CorporateActionNarrative30          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AddtlInf,omitempty"`
	SplmtryData    []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SplmtryData,omitempty"`
}

type CorporateActionNarrative30 struct {
	RegnDtls       []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 RegnDtls,omitempty"`
	PtyCtctNrrtv   []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PtyCtctNrrtv,omitempty"`
	CertfctnBrkdwn []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CertfctnBrkdwn,omitempty"`
}

type CorporateActionNarrative32 struct {
	InfToCmplyWth    []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 InfToCmplyWth,omitempty"`
	DlvryDtls        []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 DlvryDtls,omitempty"`
	FXInstrsAddtlInf []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 FXInstrsAddtlInf,omitempty"`
	InstrAddtlInf    []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 InstrAddtlInf,omitempty"`
}

type CorporateActionOption118 struct {
	OptnNb          OptionNumber1Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 OptnNb"`
	OptnTp          CorporateActionOption20Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 OptnTp"`
	FrctnDspstn     FractionDispositionType28Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 FrctnDspstn,omitempty"`
	ChngTp          []CorporateActionChangeTypeFormat6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ChngTp,omitempty"`
	ElgblForCollInd bool                                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ElgblForCollInd,omitempty"`
	CcyToBuy        ActiveCurrencyCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CcyToBuy,omitempty"`
	CcyToSell       ActiveCurrencyCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CcyToSell,omitempty"`
	CcyOptn         ActiveCurrencyCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CcyOptn,omitempty"`
	SctyId          SecurityIdentification19                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SctyId,omitempty"`
	SctiesQty       SecuritiesOption52                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SctiesQty"`
	ExctnReqdDtTm   DateAndDateTimeChoice                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ExctnReqdDtTm,omitempty"`
	RateAndAmtDtls  CorporateActionRate71                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 RateAndAmtDtls,omitempty"`
	PricDtls        CorporateActionPrice60                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PricDtls,omitempty"`
	AddtlInf        CorporateActionNarrative32               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AddtlInf,omitempty"`
}

type CorporateActionOption20Choice struct {
	Cd    CorporateActionOption9Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Prtry"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CERT, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, TAXI, PRUN
type CorporateActionOption9Code string

type CorporateActionPrice60 struct {
	IndctvOrMktPric       IndicativeOrMarketPrice8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 IndctvOrMktPric,omitempty"`
	IssePric              PriceFormat50Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 IssePric,omitempty"`
	GncCshPricRcvdPerPdct PriceFormat49Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 GncCshPricRcvdPerPdct,omitempty"`
	GncCshPricPdPerPdct   PriceFormat50Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 GncCshPricPdPerPdct,omitempty"`
}

type CorporateActionRate71 struct {
	PropsdRate         float64                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PropsdRate,omitempty"`
	OvrsbcptRate       RateAndAmountFormat39Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 OvrsbcptRate,omitempty"`
	ReqdWhldgTaxRate   []RateAndAmountFormat40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ReqdWhldgTaxRate,omitempty"`
	ReqdScndLvlTaxRate []RateAndAmountFormat40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ReqdScndLvlTaxRate,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 DtTm"`
}

type Document struct {
	CorpActnInstr CorporateActionInstructionV06 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CorpActnInstr"`
}

type DocumentIdentification31 struct {
	Id    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Id"`
	LkgTp ProcessingPosition7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 LkgTp,omitempty"`
}

type DocumentIdentification32 struct {
	Id    DocumentIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Id"`
	DocNb DocumentNumber5Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 DocNb,omitempty"`
	LkgTp ProcessingPosition7Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 LkgTp,omitempty"`
}

type DocumentIdentification3Choice struct {
	AcctSvcrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AcctSvcrDocId"`
	AcctOwnrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AcctOwnrDocId"`
}

type DocumentNumber5Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 LngNb"`
	PrtryNb GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentAttributes65 struct {
	FinInstrmId   SecurityIdentification19               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 FinInstrmId,omitempty"`
	PlcOfListg    MarketIdentification3Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PlcOfListg,omitempty"`
	DayCntBsis    InterestComputationMethodFormat4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 DayCntBsis,omitempty"`
	ClssfctnTp    ClassificationType32Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ClssfctnTp,omitempty"`
	DnmtnCcy      ActiveOrHistoricCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 DnmtnCcy,omitempty"`
	NxtCpnDt      ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 NxtCpnDt,omitempty"`
	XpryDt        ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 XpryDt,omitempty"`
	FltgRateFxgDt ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 FltgRateFxgDt,omitempty"`
	MtrtyDt       ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 MtrtyDt,omitempty"`
	IsseDt        ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 IsseDt,omitempty"`
	NxtCllblDt    ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 NxtCllblDt,omitempty"`
	PutblDt       ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PutblDt,omitempty"`
	DtdDt         ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 DtdDt,omitempty"`
	ConvsDt       ISODate                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ConvsDt,omitempty"`
	PrvsFctr      float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PrvsFctr,omitempty"`
	NxtFctr       float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 NxtFctr,omitempty"`
	IntrstRate    float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 IntrstRate,omitempty"`
	NxtIntrstRate float64                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 NxtIntrstRate,omitempty"`
	MinNmnlQty    FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 MinNmnlQty,omitempty"`
	CtrctSz       FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CtrctSz,omitempty"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AmtsdVal"`
}

// May be one of BUYU, CINL, EXPI, DIST
type FractionDispositionType10Code string

type FractionDispositionType28Choice struct {
	Cd    FractionDispositionType10Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Cd"`
	Prtry GenericIdentification30       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Prtry"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Id,omitempty"`
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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Prtry"`
}

type IdentificationType42Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Cd"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Prtry"`
}

type IndicativeOrMarketPrice8Choice struct {
	IndctvPric PriceFormat50Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 IndctvPric"`
	MktPric    PriceFormat50Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 MktPric"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

type InterestComputationMethodFormat4Choice struct {
	Cd    InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MarketIdentification3Choice struct {
	MktIdrCd MICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 MktIdrCd"`
	Desc     Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Desc"`
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
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Adr,omitempty"`
}

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AmtsdVal"`
}

type OriginalAndCurrentQuantities6 struct {
	ShrtLngPos ShortLong1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ShrtLngPos"`
	FaceAmt    float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 FaceAmt"`
	AmtsdVal   float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Tp"`
}

type PartyIdentification71Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 NmAndAdr"`
}

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PrtryId"`
}

type PartyIdentification93 struct {
	OwnrId         PartyIdentification71Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 OwnrId"`
	AltrnId        []AlternatePartyIdentification7        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AltrnId,omitempty"`
	DmclCtry       CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 DmclCtry,omitempty"`
	NonDmclCtry    []CountryCode                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 NonDmclCtry,omitempty"`
	OwndSctiesQty  FinancialInstrumentQuantity1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 OwndSctiesQty"`
	CertfctnTp     []BeneficiaryCertificationType10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CertfctnTp,omitempty"`
	CertfctnBrkdwn Max350Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CertfctnBrkdwn,omitempty"`
}

type PercentagePrice1 struct {
	PctgPricTp PriceRateType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PctgPricTp"`
	PricVal    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PricVal"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Ctry"`
}

type PriceFormat49Choice struct {
	PctgPric               PercentagePrice1                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PctgPric"`
	AmtPric                AmountPrice3                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AmtPric"`
	NotSpcfdPric           PriceValueType9Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 NotSpcfdPric"`
	AmtPricPerFinInstrmQty AmountPricePerFinancialInstrumentQuantity6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AmtPricPerFinInstrmQty"`
	AmtPricPerAmt          AmountPricePerAmount2                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AmtPricPerAmt"`
	IndxPts                float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 IndxPts"`
}

type PriceFormat50Choice struct {
	PctgPric PercentagePrice1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PctgPric"`
	AmtPric  AmountPrice3     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 AmtPric"`
}

// May be one of DISC, PREM, PRCT, YIEL
type PriceRateType3Code string

// May be one of TBSP, UNSP, UKWN
type PriceValueType9Code string

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

type ProcessingPosition7Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Prtry"`
}

type ProprietaryQuantity7 struct {
	ShrtLngPos ShortLong1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ShrtLngPos,omitempty"`
	Qty        float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Qty"`
	QtyTp      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 QtyTp"`
	Issr       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Issr"`
	SchmeNm    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SchmeNm,omitempty"`
}

type ProprietaryQuantity8 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 QtyTp"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SchmeNm,omitempty"`
}

type Quantity17Choice struct {
	QtyChc   Quantity18Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 QtyChc"`
	PrtryQty ProprietaryQuantity7 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PrtryQty"`
}

type Quantity18Choice struct {
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities6 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 OrgnlAndCurFaceAmt"`
	SgndQty            SignedQuantityFormat6         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SgndQty"`
}

type Quantity19Choice struct {
	Qty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Qty"`
	PrtryQty ProprietaryQuantity8               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PrtryQty"`
}

// May be one of QALL
type Quantity1Code string

type Quantity20Choice struct {
	Cd                 Quantity1Code                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Cd"`
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 OrgnlAndCurFaceAmt"`
	Qty                FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Qty"`
}

type RateAndAmountFormat39Choice struct {
	Rate float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Rate"`
	Amt  ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Amt"`
}

type RateAndAmountFormat40Choice struct {
	Rate          float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Rate"`
	Amt           ActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Amt"`
	RateTpAndRate RateTypeAndPercentageRate8       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 RateTpAndRate"`
}

type RateType42Choice struct {
	Cd    WithholdingTaxRateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Prtry"`
}

type RateTypeAndPercentageRate8 struct {
	RateTp RateType42Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 RateTp"`
	Rate   float64          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Rate"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat8Choice struct {
	Id      SafekeepingPlaceTypeAndText6             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 TpAndId"`
	Prtry   GenericIdentification78                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Id"`
}

type SafekeepingPlaceTypeAndText6 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Id,omitempty"`
}

type SecuritiesOption52 struct {
	CondlQty FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 CondlQty,omitempty"`
	InstdQty Quantity20Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 InstdQty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Desc,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat6 struct {
	ShrtLngPos ShortLong1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ShrtLngPos"`
	Qty        FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Qty"`
}

type SignedQuantityFormat7 struct {
	ShrtLngPos ShortLong1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 ShrtLngPos"`
	QtyChc     Quantity19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 QtyChc"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.06 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of ARNU, CCPT, CHTY, CORP, DRLC, FIIN, TXID
type TypeOfIdentification1Code string

// May be one of BWIT, FTCA, NRAT
type WithholdingTaxRateType1Code string

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

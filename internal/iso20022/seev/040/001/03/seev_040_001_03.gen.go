// Code generated by main. DO NOT EDIT.

package seev_040_001_03

type AccountIdentification17 struct {
	SfkpgAcct Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 SfkpgAcct"`
	AcctOwnr  PartyIdentification36Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 SfkpgPlc,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CorporateActionEventType7Choice struct {
	Cd    CorporateActionEventType8Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Cd"`
	Prtry GenericIdentification20       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Prtry"`
}

// May be one of ACTV, ATTI, BIDS, BONU, BPUT, BRUP, CAPG, CAPI, CERT, CHAN, CLSA, CONS, CONV, COOP, DECR, DETI, DFLT, DLST, DRAW, DRIP, DSCL, DTCH, DVCA, DVOP, DVSC, DVSE, EXOF, EXRI, EXTM, EXWA, CAPD, INCR, INTR, LIQU, MCAL, MRGR, ODLT, OTHR, PARI, PCAL, PDEF, PINK, PLAC, PPMT, PRED, PRII, PRIO, REDM, REDO, REMK, RHDI, RHTS, SHPR, SMAL, SOFF, SPLF, SPLR, SUSP, TEND, TREC, WRTH, WTRC, CREV, DRCA
type CorporateActionEventType8Code string

type CorporateActionGeneralInformation33 struct {
	CorpActnEvtId      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 OffclCorpActnEvtId,omitempty"`
	EvtTp              CorporateActionEventType7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 EvtTp"`
	FinInstrmId        SecurityIdentification14        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 FinInstrmId,omitempty"`
}

type CorporateActionInstructionCancellationRequestV03 struct {
	ChngInstrInd   bool                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 ChngInstrInd,omitempty"`
	InstrId        DocumentIdentification15            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 InstrId"`
	CorpActnGnlInf CorporateActionGeneralInformation33 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 CorpActnGnlInf"`
	AcctDtls       AccountIdentification17             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 AcctDtls"`
	CorpActnInstr  CorporateActionOption42             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 CorpActnInstr"`
	SplmtryData    []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 SplmtryData,omitempty"`
}

type CorporateActionOption12Choice struct {
	Cd    CorporateActionOption9Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Cd"`
	Prtry GenericIdentification20    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Prtry"`
}

type CorporateActionOption42 struct {
	OptnNb          OptionNumber1Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 OptnNb"`
	OptnTp          CorporateActionOption12Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 OptnTp"`
	InstdOrQtyToRcv InstructedOrQuantityToReceive1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 InstdOrQtyToRcv"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CERT, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, TAXI, PRUN
type CorporateActionOption9Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	CorpActnInstrCxlReq CorporateActionInstructionCancellationRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 CorpActnInstrCxlReq"`
}

type DocumentIdentification15 struct {
	Id    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Id"`
	LkgTp ProcessingPosition1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 LkgTp,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 AmtsdVal"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Id,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Prtry"`
}

type InstructedOrQuantityToReceive1Choice struct {
	InstdQty Quantity5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 InstdQty"`
	QtyToRcv Quantity5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 QtyToRcv"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

type OptionNumber1Choice struct {
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Tp"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 PrtryId"`
}

type ProcessingPosition1Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Prtry"`
}

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

// May be one of QALL
type Quantity1Code string

type Quantity5Choice struct {
	Cd                 Quantity1Code                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Cd"`
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 OrgnlAndCurFaceAmt"`
	Qty                FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Qty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat2Choice struct {
	Id      SafekeepingPlaceTypeAndText2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Id"`
}

type SafekeepingPlaceTypeAndText2 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Id,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Desc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

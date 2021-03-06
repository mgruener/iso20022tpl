// Code generated by main. DO NOT EDIT.

package seev_040_001_07

type AccountIdentification31 struct {
	SfkpgAcct Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 SfkpgAcct"`
	AcctOwnr  PartyIdentification92Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 SfkpgPlc,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, PRII, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH
type CorporateActionEventType20Code string

type CorporateActionEventType52Choice struct {
	Cd    CorporateActionEventType20Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Prtry"`
}

type CorporateActionGeneralInformation110 struct {
	CorpActnEvtId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 OffclCorpActnEvtId,omitempty"`
	EvtTp              CorporateActionEventType52Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 EvtTp"`
	FinInstrmId        SecurityIdentification19         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 FinInstrmId,omitempty"`
}

type CorporateActionInstructionCancellationRequestV07 struct {
	ChngInstrInd   bool                                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 ChngInstrInd,omitempty"`
	InstrId        DocumentIdentification31             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 InstrId"`
	CorpActnGnlInf CorporateActionGeneralInformation110 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 CorpActnGnlInf"`
	AcctDtls       AccountIdentification31              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 AcctDtls"`
	CorpActnInstr  CorporateActionOption120             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 CorpActnInstr"`
	SplmtryData    []SupplementaryData1                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 SplmtryData,omitempty"`
}

type CorporateActionOption120 struct {
	OptnNb   OptionNumber1Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 OptnNb"`
	OptnTp   CorporateActionOption20Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 OptnTp"`
	InstdQty Quantity20Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 InstdQty"`
}

type CorporateActionOption20Choice struct {
	Cd    CorporateActionOption9Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Prtry"`
}

// May be one of ABST, AMGT, BSPL, BUYA, CASE, CASH, CERT, CEXC, CONN, CONY, CTEN, EXER, LAPS, MKDW, MKUP, MNGT, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, PROX, QINV, SECU, SLLE, SPLI, TAXI, PRUN
type CorporateActionOption9Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	CorpActnInstrCxlReq CorporateActionInstructionCancellationRequestV07 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 CorpActnInstrCxlReq"`
}

type DocumentIdentification31 struct {
	Id    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Id"`
	LkgTp ProcessingPosition7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 LkgTp,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 AmtsdVal"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Id,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Prtry"`
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
	Nb Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Nb"`
	Cd OptionNumber1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Cd"`
}

// May be one of UNSO
type OptionNumber1Code string

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Tp"`
}

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 PrtryId"`
}

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

type ProcessingPosition7Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Prtry"`
}

// May be one of QALL
type Quantity1Code string

type Quantity20Choice struct {
	Cd                 Quantity1Code                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Cd"`
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 OrgnlAndCurFaceAmt"`
	Qty                FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Qty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat8Choice struct {
	Id      SafekeepingPlaceTypeAndText6             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 TpAndId"`
	Prtry   GenericIdentification78                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Id"`
}

type SafekeepingPlaceTypeAndText6 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Id,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Desc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.07 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

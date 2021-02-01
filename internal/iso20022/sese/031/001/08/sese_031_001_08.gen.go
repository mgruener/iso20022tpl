// Code generated by main. DO NOT EDIT.

package sese_031_001_08

type AcknowledgedAcceptedStatus21Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Rsn"`
}

type AcknowledgementReason12Choice struct {
	Cd    AcknowledgementReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}

// May be one of ADEA, SMPG, OTHR, CDCY, CDRG, CDRE, NSTP, RQWV, LATE
type AcknowledgementReason5Code string

type AcknowledgementReason9 struct {
	Cd          AcknowledgementReason12Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	AddtlRsnInf Max210Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AddtlRsnInf,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// May be one of LAMI, NBOR, YBOR, RTRN
type AutoBorrowing2Code string

type AutomaticBorrowing7Choice struct {
	Cd    AutoBorrowing2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}

type DeniedReason10 struct {
	Cd          DeniedReason15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AddtlRsnInf,omitempty"`
}

type DeniedReason15Choice struct {
	Cd    DeniedReason6Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}

// May be one of ADEA, CDCY, CDRE, DCAN, DSET, DPRG, DREP, LATE, OTHR, CDRG
type DeniedReason6Code string

type DeniedStatus15Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 NoSpcfdRsn"`
	Rsn        []DeniedReason10 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Rsn"`
}

type Document struct {
	SctiesSttlmCondModStsAdvc SecuritiesSettlementConditionModificationStatusAdviceV08 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 SctiesSttlmCondModStsAdvc"`
}

type DocumentNumber5Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 LngNb"`
	PrtryNb GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 SchmeNm,omitempty"`
}

type HoldIndicator6 struct {
	Ind bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Ind"`
	Rsn []RegistrationReason5 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Rsn,omitempty"`
}

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

type Identification14 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Id"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be one of LINK, UNLK, SOFT
type LinkageType1Code string

type LinkageType3Choice struct {
	Cd    LinkageType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}

type Linkages53 struct {
	PrcgPos ProcessingPosition8Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 PrcgPos,omitempty"`
	MsgNb   DocumentNumber5Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 MsgNb,omitempty"`
	Ref     References65Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Ref"`
	RefOwnr PartyIdentification127Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 RefOwnr,omitempty"`
}

type MatchingDenied3Choice struct {
	Cd    MatchingProcess1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}

// May be one of UNMT, MTRE
type MatchingProcess1Code string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max52Text string

// Must be at least 1 items long
type Max70Text string

// May be one of NORE
type NoReasonCode string

type PartyIdentification127Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 PrtryId"`
}

type PartyIdentification144 struct {
	Id  PartyIdentification127Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 LEI,omitempty"`
}

type PendingReason16 struct {
	Cd          PendingReason28Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AddtlRsnInf,omitempty"`
}

type PendingReason28Choice struct {
	Cd    PendingReason6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}

// May be one of ADEA, CONF, OTHR, CDRG, CDCY, CDRE
type PendingReason6Code string

type PendingStatus38Choice struct {
	NoSpcfdRsn NoReasonCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 NoSpcfdRsn"`
	Rsn        []PendingReason16 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Rsn"`
}

type PriorityNumeric4Choice struct {
	Nmrc  Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Nmrc"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}

// May be one of AFTE, BEFO, WITH
type ProcessingPosition4Code string

type ProcessingPosition8Choice struct {
	Cd    ProcessingPosition4Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}

type ProcessingStatus50Choice struct {
	AckdAccptd AcknowledgedAcceptedStatus21Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AckdAccptd"`
	Rjctd      RejectionOrRepairStatus31Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Rjctd"`
	Cmpltd     ProprietaryReason4                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cmpltd"`
	Dnd        DeniedStatus15Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Dnd"`
	Pdg        PendingStatus38Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Pdg"`
	Prtry      ProprietaryStatusAndReason6        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 PrtryRsn,omitempty"`
}

type References23 struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AcctOwnrTxId,omitempty"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 PrcrTxId,omitempty"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 PoolId,omitempty"`
	CmonId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 CmonId,omitempty"`
	TradId            Max52Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 TradId,omitempty"`
}

type References65Choice struct {
	SctiesSttlmTxId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 SctiesSttlmTxId"`
	IntraPosMvmntId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 IntraPosMvmntId"`
	IntraBalMvmntId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 IntraBalMvmntId"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AcctSvcrTxId"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 MktInfrstrctrTxId"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 PoolId"`
	CmonId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 CmonId"`
	TradId            Max52Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 TradId"`
	OthrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 OthrTxId"`
}

type Registration10Choice struct {
	Cd    Registration2Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}

// May be one of PTYH, CSDH, CDEL, CVAL
type Registration2Code string

type RegistrationReason5 struct {
	Cd       Registration10Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	AddtlInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AddtlInf,omitempty"`
}

type RejectionAndRepairReason25Choice struct {
	Cd    RejectionReason27Code   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}

type RejectionOrRepairReason25 struct {
	Cd          RejectionAndRepairReason25Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	AddtlRsnInf Max210Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairStatus31Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason25 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Rsn"`
}

// May be one of ADEA, LATE, SAFE, NRGM, NRGN, OTHR, REFE, INVM, INVL
type RejectionReason27Code string

type RequestDetails20 struct {
	Ref          References23                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Ref"`
	RstrctnRef   []RestrictionIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 RstrctnRef,omitempty"`
	AutomtcBrrwg AutomaticBorrowing7Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AutomtcBrrwg,omitempty"`
	RtnInd       bool                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 RtnInd,omitempty"`
	Lkg          LinkageType3Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Lkg,omitempty"`
	Prty         PriorityNumeric4Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prty,omitempty"`
	OthrPrcg     []GenericIdentification30           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 OthrPrcg,omitempty"`
	PrtlSttlmInd SettlementTransactionCondition5Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 PrtlSttlmInd,omitempty"`
	SctiesRTGS   SecuritiesRTGS4Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 SctiesRTGS,omitempty"`
	HldInd       HoldIndicator6                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 HldInd,omitempty"`
	MtchgDnl     MatchingDenied3Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 MtchgDnl,omitempty"`
	UnltrlSplt   UnilateralSplit3Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 UnltrlSplt,omitempty"`
	Lnkgs        []Linkages53                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Lnkgs,omitempty"`
}

type RestrictionIdentification1 struct {
	Cd RestrictionReference1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	Id Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Id"`
}

// May be one of ADDC, ADDS, REMC, REMS
type RestrictionReference1Code string

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Nm,omitempty"`
}

type SecuritiesRTGS4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Ind"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}

type SecuritiesSettlementConditionModificationStatusAdviceV08 struct {
	ReqRef      Identification14         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 ReqRef"`
	AcctOwnr    PartyIdentification144   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount19      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 SfkpgAcct,omitempty"`
	ReqDtls     RequestDetails20         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 ReqDtls,omitempty"`
	PrcgSts     ProcessingStatus50Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 PrcgSts"`
	SplmtryData []SupplementaryData1     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 SplmtryData,omitempty"`
}

// May be one of TRAD
type SecuritiesTransactionType5Code string

// May be one of PART, NPAR, PARC, PARQ
type SettlementTransactionCondition5Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type UnilateralSplit3Choice struct {
	Cd    SecuritiesTransactionType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.08 Prtry"`
}
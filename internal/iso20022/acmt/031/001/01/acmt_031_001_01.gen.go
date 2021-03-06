// Code generated by main. DO NOT EDIT.

package acmt_031_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

type AccountSwitchDetails1 struct {
	UnqRefNb      Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 UnqRefNb"`
	RtgUnqRefNb   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RtgUnqRefNb"`
	SwtchRcvdDtTm ISODateTime                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 SwtchRcvdDtTm,omitempty"`
	SwtchDt       ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 SwtchDt,omitempty"`
	SwtchTp       SwitchType1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 SwtchTp"`
	SwtchSts      SwitchStatus1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 SwtchSts,omitempty"`
	BalTrfWndw    BalanceTransferWindow1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 BalTrfWndw,omitempty"`
	Rspn          []ResponseDetails1         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Rspn,omitempty"`
}

type AccountSwitchRequestBalanceTransferV01 struct {
	MsgId         MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 AcctSwtchDtls"`
	NewAcct       CashAccount36          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 NewAcct"`
	NmntdAcct     CashAccount36          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 NmntdAcct,omitempty"`
	BalTrf        []BalanceTransfer1     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 BalTrf,omitempty"`
	SplmtryData   []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 SplmtryData,omitempty"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

type BalanceTransfer1 struct {
	BalTrfRef     BalanceTransferReference1    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 BalTrfRef,omitempty"`
	BalTrfMtd     SettlementMethod1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 BalTrfMtd,omitempty"`
	BalTrfFndgLmt BalanceTransferFundingLimit1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 BalTrfFndgLmt,omitempty"`
}

type BalanceTransferFundingLimit1 struct {
	CcyAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CcyAmt"`
}

type BalanceTransferReference1 struct {
	BalTrfRef Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 BalTrfRef"`
}

// May be one of DAYH, EARL
type BalanceTransferWindow1Code string

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 PstlAdr,omitempty"`
}

// May be one of FWNG, PREC
type BusinessDayConvention1Code string

type CashAccount24 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Id"`
	Tp  CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Nm,omitempty"`
}

type CashAccount36 struct {
	Id   AccountIdentification4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Id"`
	Tp   CashAccountType2Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Ccy,omitempty"`
	Nm   Max70Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Nm,omitempty"`
	Ownr PartyIdentification125                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Ownr,omitempty"`
	Svcr BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Svcr,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

// May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code string

type Cheque7 struct {
	ChqTp       ChequeType2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 ChqTp,omitempty"`
	ChqNb       Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 ChqNb,omitempty"`
	ChqFr       NameAndAddress10            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 ChqFr,omitempty"`
	DlvryMtd    ChequeDeliveryMethod1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 DlvryMtd,omitempty"`
	DlvrTo      NameAndAddress10            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 DlvrTo,omitempty"`
	InstrPrty   Priority2Code               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 InstrPrty,omitempty"`
	ChqMtrtyDt  ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 ChqMtrtyDt,omitempty"`
	FrmsCd      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 FrmsCd,omitempty"`
	MemoFld     []Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 MemoFld,omitempty"`
	RgnlClrZone Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RgnlClrZone,omitempty"`
	PrtLctn     Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 PrtLctn,omitempty"`
	Sgntr       []Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Sgntr,omitempty"`
}

// May be one of MLDB, MLCD, MLFA, CRDB, CRCD, CRFA, PUDB, PUCD, PUFA, RGDB, RGCD, RGFA
type ChequeDelivery1Code string

type ChequeDeliveryMethod1Choice struct {
	Cd    ChequeDelivery1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

// May be one of CCHQ, CCCH, BCHQ, DRFT, ELDR
type ChequeType2Code string

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Othr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type CreditTransferTransaction27 struct {
	PmtId           PaymentIdentification1                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 PmtId"`
	PmtTpInf        PaymentTypeInformation19                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 PmtTpInf,omitempty"`
	TaxRateMrkr     TaxRateMarker1Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TaxRateMrkr,omitempty"`
	Amt             ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Amt"`
	ChrgBr          ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 ChrgBr,omitempty"`
	ChqInstr        Cheque7                                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 ChqInstr,omitempty"`
	Frqcy           Frequency1                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Frqcy,omitempty"`
	TrfInstr        TransferInstruction1                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TrfInstr,omitempty"`
	UltmtDbtr       PartyIdentification125                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 UltmtDbtr,omitempty"`
	IntrmyAgt1      BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 IntrmyAgt1,omitempty"`
	IntrmyAgt2      BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 IntrmyAgt2,omitempty"`
	IntrmyAgt3      BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 IntrmyAgt3,omitempty"`
	CdtrAgt         BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CdtrAgt"`
	Cdtr            PartyIdentification125                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cdtr,omitempty"`
	CdtrAcct        CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CdtrAcct,omitempty"`
	UltmtCdtr       PartyIdentification125                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 UltmtCdtr,omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent1               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 InstrForCdtrAgt,omitempty"`
	Purp            Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Purp,omitempty"`
	RgltryRptg      []RegulatoryReporting3                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RgltryRptg,omitempty"`
	Tax             TaxInformation6                              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Tax,omitempty"`
	RltdRmtInf      []RemittanceLocation2                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RltdRmtInf,omitempty"`
	RmtInf          RemittanceInformation14                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RmtInf,omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Tp,omitempty"`
	Ref Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CdOrPrtry"`
	Issr      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Issr,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 ToDt"`
}

type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

type Document struct {
	AcctSwtchReqBalTrf AccountSwitchRequestBalanceTransferV01 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 AcctSwtchReqBalTrf"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CdtDbtInd,omitempty"`
	Rsn       Max4Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Rsn,omitempty"`
	AddtlInf  Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 AddtlInf,omitempty"`
}

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT
type DocumentType5Code string

type EndPoint1Choice struct {
	NbOfPmts  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 NbOfPmts,omitempty"`
	LastPmtDt ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 LastPmtDt,omitempty"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalCategoryPurpose1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalDiscountAmountType1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalLocalInstrument1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalPurpose1Code string

// Must be at least 1 items long
type ExternalServiceLevel1Code string

// Must be at least 1 items long
type ExternalTaxAmountType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       BICFIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Othr,omitempty"`
}

type Frequency1 struct {
	Seq                 Max3NumericText            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Seq,omitempty"`
	StartDt             ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 StartDt"`
	EndPtChc            EndPoint1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 EndPtChc"`
	ReqdFrqcyPttrn      Frequency37Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 ReqdFrqcyPttrn,omitempty"`
	NonWorkgDayAdjstmnt BusinessDayConvention1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 NonWorkgDayAdjstmnt,omitempty"`
}

// May be one of NEVR, YEAR, RATE, MIAN, QURT
type Frequency10Code string

type Frequency37Choice struct {
	Cd    Frequency10Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Issr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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

// May be one of CHQB, HOLD, PHOB, TELB
type Instruction3Code string

type InstructionForCreditorAgent1 struct {
	Cd       Instruction3Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd,omitempty"`
	InstrInf Max140Text       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 InstrInf,omitempty"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

// Must be at least 1 items long
type Max10Text string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must be at least 1 items long
type Max4Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CreDtTm"`
}

type NameAndAddress10 struct {
	Nm  Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Nm"`
	Adr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Adr"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OrganisationIdentification8 struct {
	AnyBIC AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

type Party34Choice struct {
	OrgId  OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 OrgId"`
	PrvtId PersonIdentification13      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 PrvtId"`
}

type PartyIdentification125 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 PstlAdr,omitempty"`
	Id        Party34Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CtctDtls,omitempty"`
}

type PaymentIdentification1 struct {
	InstrId    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 InstrId,omitempty"`
	EndToEndId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 EndToEndId"`
}

type PaymentTypeInformation19 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 InstrPrty,omitempty"`
	SvcLvl    ServiceLevel8Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CtgyPurp,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 AdrLine,omitempty"`
}

// May be one of HIGH, NORM
type Priority2Code string

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

type ReferredDocumentInformation3 struct {
	Tp     ReferredDocumentType2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Tp,omitempty"`
	Nb     Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Nb,omitempty"`
	RltdDt ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RltdDt,omitempty"`
}

type ReferredDocumentType1Choice struct {
	Cd    DocumentType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

type ReferredDocumentType2 struct {
	CdOrPrtry ReferredDocumentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CdOrPrtry"`
	Issr      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Issr,omitempty"`
}

type RegulatoryAuthority2 struct {
	Nm   Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Nm,omitempty"`
	Ctry CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Ctry,omitempty"`
}

type RegulatoryReporting3 struct {
	DbtCdtRptgInd RegulatoryReportingType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 DbtCdtRptgInd,omitempty"`
	Authrty       RegulatoryAuthority2             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Authrty,omitempty"`
	Dtls          []StructuredRegulatoryReporting3 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Dtls,omitempty"`
}

// May be one of CRED, DEBT, BOTH
type RegulatoryReportingType1Code string

type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RmtdAmt,omitempty"`
}

type RemittanceInformation14 struct {
	Ustrd []Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation14 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Strd,omitempty"`
}

type RemittanceLocation2 struct {
	RmtId             Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RmtId,omitempty"`
	RmtLctnMtd        RemittanceLocationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RmtLctnMtd,omitempty"`
	RmtLctnElctrncAdr Max2048Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RmtLctnElctrncAdr,omitempty"`
	RmtLctnPstlAdr    NameAndAddress10              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RmtLctnPstlAdr,omitempty"`
}

// May be one of FAXI, EDIC, URID, EMAL, POST, SMSM
type RemittanceLocationMethod2Code string

type ResponseDetails1 struct {
	RspnCd    Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RspnCd"`
	AddtlDtls Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 AddtlDtls,omitempty"`
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

type SettlementMethod1Choice struct {
	Cdt CreditTransferTransaction27 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cdt"`
	Dbt CreditTransferTransaction27 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Dbt"`
}

type StructuredRegulatoryReporting3 struct {
	Tp   Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Tp,omitempty"`
	Dt   ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Dt,omitempty"`
	Ctry CountryCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Ctry,omitempty"`
	Cd   Max10Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd,omitempty"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Amt,omitempty"`
	Inf  []Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Inf,omitempty"`
}

type StructuredRemittanceInformation14 struct {
	RfrdDocInf  []ReferredDocumentInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification125         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Invcr,omitempty"`
	Invcee      PartyIdentification125         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Invcee,omitempty"`
	AddtlRmtInf []Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 AddtlRmtInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of ACPT, BTRQ, BTRS, COMP, REDT, REDE, REJT, REQU, TMTN
type SwitchStatus1Code string

// May be one of FULL, PART
type SwitchType1Code string

type TaxAmount2 struct {
	Rate         float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Rate,omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TaxblBaseAmt,omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails2               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Dtls,omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry"`
}

type TaxAuthorisation1 struct {
	Titl Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Titl,omitempty"`
	Nm   Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Nm,omitempty"`
}

type TaxInformation6 struct {
	Cdtr            TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Dbtr,omitempty"`
	AdmstnZn        Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 AdmstnZn,omitempty"`
	RefNb           Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RefNb,omitempty"`
	Mtd             Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TtlTaxAmt,omitempty"`
	Dt              ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Dt,omitempty"`
	SeqNb           float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Rcrd,omitempty"`
}

type TaxParty1 struct {
	TaxId  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TaxId,omitempty"`
	RegnId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RegnId,omitempty"`
	TaxTp  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxId   Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TaxId,omitempty"`
	RegnId  Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 RegnId,omitempty"`
	TaxTp   Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TaxTp,omitempty"`
	Authstn TaxAuthorisation1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Authstn,omitempty"`
}

type TaxPeriod2 struct {
	Yr     ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Tp,omitempty"`
	FrToDt DatePeriod2          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 FrToDt,omitempty"`
}

// May be one of ALPR, ALIT, GRSS
type TaxRateMarker1Code string

type TaxRecord2 struct {
	Tp       Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Tp,omitempty"`
	Ctgy     Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Ctgy,omitempty"`
	CtgyDtls Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CtgyDtls,omitempty"`
	DbtrSts  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 DbtrSts,omitempty"`
	CertId   Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 CertId,omitempty"`
	FrmsCd   Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 FrmsCd,omitempty"`
	Prd      TaxPeriod2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prd,omitempty"`
	TaxAmt   TaxAmount2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TaxAmt,omitempty"`
	AddtlInf Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 AddtlInf,omitempty"`
}

type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Amt"`
}

// May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code string

type TransferInstruction1 struct {
	TrfInd    bool        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 TrfInd,omitempty"`
	Cd        Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Cd"`
	Prtry     Max256Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Prtry,omitempty"`
	StartDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 StartDtTm,omitempty"`
	StartDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 StartDt,omitempty"`
	Desc      Max350Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.01 Desc,omitempty"`
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

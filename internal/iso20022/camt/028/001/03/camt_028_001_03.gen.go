// Code generated by main. DO NOT EDIT.

package camt_028_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AdditionalPaymentInformationV03 struct {
	Assgnmt CaseAssignment2                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Assgnmt"`
	Case    Case2                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Case"`
	Undrlyg UnderlyingTransaction1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Undrlyg"`
	Inf     PaymentComplementaryInformation2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Inf"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AmountType3Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 EqvtAmt"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

type BranchAndFinancialInstitutionIdentification4 struct {
	FinInstnId FinancialInstitutionIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 PstlAdr,omitempty"`
}

type Case2 struct {
	Id             Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Id"`
	Cretr          Party7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cretr"`
	ReopCaseIndctn bool         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 ReopCaseIndctn,omitempty"`
}

type CaseAssignment2 struct {
	Id      Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Id"`
	Assgnr  Party7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Assgnr"`
	Assgne  Party7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Assgne"`
	CreDtTm ISODateTime  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CreDtTm"`
}

type CashAccount16 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Id"`
	Tp  CashAccountType2             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Nm,omitempty"`
}

type CashAccountType2 struct {
	Cd    CashAccountType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

// May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code string

// May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Othr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Tp,omitempty"`
	Ref Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CdOrPrtry"`
	Issr      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Issr,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CtryOfBirth"`
}

type Document struct {
	AddtlPmtInf AdditionalPaymentInformationV03 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 AddtlPmtInf"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CdtDbtInd,omitempty"`
	Rsn       Max4Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Rsn,omitempty"`
	AddtlInf  Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 AddtlInf,omitempty"`
}

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT
type DocumentType5Code string

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Amt"`
	CcyOfTrf ActiveOrHistoricCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CcyOfTrf"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashClearingSystem1Code string

// Must be at least 1 items long
type ExternalCategoryPurpose1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

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

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

type FinancialInstitutionIdentification7 struct {
	BIC         BICIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 BIC,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Issr,omitempty"`
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

// May be one of PHOA, TELA
type Instruction4Code string

type InstructionForCreditorAgent1 struct {
	Cd       Instruction3Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd,omitempty"`
	InstrInf Max140Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 InstrInf,omitempty"`
}

type InstructionForNextAgent1 struct {
	Cd       Instruction4Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd,omitempty"`
	InstrInf Max140Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 InstrInf,omitempty"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max4Text string

// Must be at least 1 items long
type Max70Text string

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OrganisationIdentification4 struct {
	BICOrBEI AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 BICOrBEI,omitempty"`
	Othr     []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

type OriginalGroupInformation3 struct {
	OrgnlMsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlMsgId"`
	OrgnlMsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlCreDtTm,omitempty"`
}

type Party6Choice struct {
	OrgId  OrganisationIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 PrvtId"`
}

type Party7Choice struct {
	Pty PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Pty"`
	Agt BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Agt"`
}

type PartyIdentification32 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 PstlAdr,omitempty"`
	Id        Party6Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CtctDtls,omitempty"`
}

type PaymentComplementaryInformation2 struct {
	InstrId          Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 InstrId,omitempty"`
	EndToEndId       Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 EndToEndId,omitempty"`
	TxId             Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 TxId,omitempty"`
	PmtTpInf         PaymentTypeInformation22                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 PmtTpInf,omitempty"`
	ReqdExctnDt      ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 ReqdExctnDt,omitempty"`
	ReqdColltnDt     ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 ReqdColltnDt,omitempty"`
	IntrBkSttlmDt    ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 IntrBkSttlmDt,omitempty"`
	Amt              AmountType3Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Amt,omitempty"`
	IntrBkSttlmAmt   ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 IntrBkSttlmAmt,omitempty"`
	ChrgBr           ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 ChrgBr,omitempty"`
	UltmtDbtr        PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 UltmtDbtr,omitempty"`
	Dbtr             PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Dbtr,omitempty"`
	DbtrAcct         CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 DbtrAcct,omitempty"`
	DbtrAgt          BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 DbtrAgt,omitempty"`
	DbtrAgtAcct      CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 DbtrAgtAcct,omitempty"`
	SttlmInf         SettlementInformation13                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 SttlmInf,omitempty"`
	IntrmyAgt1       BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct   CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2       BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct   CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3       BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct   CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 IntrmyAgt3Acct,omitempty"`
	CdtrAgt          BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CdtrAgt,omitempty"`
	CdtrAgtAcct      CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CdtrAgtAcct,omitempty"`
	Cdtr             PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cdtr,omitempty"`
	CdtrAcct         CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CdtrAcct,omitempty"`
	UltmtCdtr        PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 UltmtCdtr,omitempty"`
	Purp             Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Purp,omitempty"`
	InstrForDbtrAgt  Max140Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 InstrForDbtrAgt,omitempty"`
	PrvsInstgAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 PrvsInstgAgt,omitempty"`
	PrvsInstgAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 PrvsInstgAgtAcct,omitempty"`
	InstrForNxtAgt   []InstructionForNextAgent1                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 InstrForNxtAgt,omitempty"`
	InstrForCdtrAgt  []InstructionForCreditorAgent1               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 InstrForCdtrAgt,omitempty"`
	RmtInf           RemittanceInformation5                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 RmtInf,omitempty"`
}

type PaymentTypeInformation22 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 ClrChanl,omitempty"`
	SvcLvl    ServiceLevel8Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 LclInstrm,omitempty"`
	SeqTp     SequenceType1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 SeqTp,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CtgyPurp,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 AdrLine,omitempty"`
}

// May be one of HIGH, NORM
type Priority2Code string

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

type ReferredDocumentInformation3 struct {
	Tp     ReferredDocumentType2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Tp,omitempty"`
	Nb     Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Nb,omitempty"`
	RltdDt ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 RltdDt,omitempty"`
}

type ReferredDocumentType1Choice struct {
	Cd    DocumentType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

type ReferredDocumentType2 struct {
	CdOrPrtry ReferredDocumentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CdOrPrtry"`
	Issr      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Issr,omitempty"`
}

type RemittanceAmount1 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 DuePyblAmt,omitempty"`
	DscntApldAmt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CdtNoteAmt,omitempty"`
	TaxAmt            ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 RmtdAmt,omitempty"`
}

type RemittanceInformation5 struct {
	Ustrd []Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Strd,omitempty"`
}

// May be one of FRST, RCUR, FNAL, OOFF
type SequenceType1Code string

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Prtry"`
}

type SettlementInformation13 struct {
	SttlmMtd             SettlementMethod1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 SttlmMtd"`
	SttlmAcct            CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 SttlmAcct,omitempty"`
	ClrSys               ClearingSystemIdentification3Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 ClrSys,omitempty"`
	InstgRmbrsmntAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 ThrdRmbrsmntAgtAcct,omitempty"`
}

// May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

type StructuredRemittanceInformation7 struct {
	RfrdDocInf  []ReferredDocumentInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount1              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification32          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Invcr,omitempty"`
	Invcee      PartyIdentification32          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Invcee,omitempty"`
	AddtlRmtInf []Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 AddtlRmtInf,omitempty"`
}

type UnderlyingGroupInformation1 struct {
	OrgnlMsgId         Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlMsgId"`
	OrgnlMsgNmId       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlMsgNmId"`
	OrgnlCreDtTm       ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlCreDtTm,omitempty"`
	OrgnlMsgDlvryChanl Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlMsgDlvryChanl,omitempty"`
}

type UnderlyingPaymentInstruction1 struct {
	OrgnlGrpInf     UnderlyingGroupInformation1       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlGrpInf,omitempty"`
	OrgnlPmtInfId   Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlPmtInfId,omitempty"`
	OrgnlInstrId    Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlEndToEndId,omitempty"`
	OrgnlInstdAmt   ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlInstdAmt"`
	ReqdExctnDt     ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 ReqdExctnDt,omitempty"`
	ReqdColltnDt    ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 ReqdColltnDt,omitempty"`
}

type UnderlyingPaymentTransaction1 struct {
	OrgnlGrpInf         UnderlyingGroupInformation1       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlGrpInf,omitempty"`
	OrgnlInstrId        Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId     Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlEndToEndId,omitempty"`
	OrgnlTxId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlTxId,omitempty"`
	OrgnlIntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlIntrBkSttlmDt"`
}

type UnderlyingStatementEntry1 struct {
	OrgnlGrpInf OriginalGroupInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlGrpInf,omitempty"`
	OrgnlStmtId Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlStmtId,omitempty"`
	OrgnlNtryId Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 OrgnlNtryId,omitempty"`
}

type UnderlyingTransaction1Choice struct {
	Initn    UnderlyingPaymentInstruction1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 Initn"`
	IntrBk   UnderlyingPaymentTransaction1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 IntrBk"`
	StmtNtry UnderlyingStatementEntry1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.03 StmtNtry"`
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

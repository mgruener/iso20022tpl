// Code generated by main. DO NOT EDIT.

package acmt_008_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AccountContract2 struct {
	TrgtGoLiveDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 TrgtGoLiveDt,omitempty"`
	TrgtClsgDt   ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 TrgtClsgDt,omitempty"`
	UrgcyFlg     bool    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 UrgcyFlg,omitempty"`
}

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Othr"`
}

type AccountOpeningAmendmentRequestV02 struct {
	Refs             References4                                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Refs"`
	Fr               OrganisationIdentification8                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Fr,omitempty"`
	CtrctDts         AccountContract2                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 CtrctDts,omitempty"`
	UndrlygMstrAgrmt ContractDocument1                            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 UndrlygMstrAgrmt,omitempty"`
	Acct             CustomerAccount4                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Acct"`
	AcctSvcrId       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 AcctSvcrId"`
	Org              Organisation12                               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Org"`
	Mndt             []OperationMandate2                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Mndt,omitempty"`
	Grp              []Group1                                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Grp,omitempty"`
	RefAcct          CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 RefAcct,omitempty"`
	DgtlSgntr        []PartyAndSignature2                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 DgtlSgntr,omitempty"`
	SplmtryData      []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 SplmtryData,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Prtry"`
}

// May be one of ENAB, DISA, DELE, FORM
type AccountStatus3Code string

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type Authorisation2 struct {
	MaxAmtByTx          FixedAmountOrUnlimited1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 MaxAmtByTx,omitempty"`
	MaxAmtByPrd         []MaximumAmountByPeriod1      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 MaxAmtByPrd,omitempty"`
	MaxAmtByBlkSubmissn FixedAmountOrUnlimited1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 MaxAmtByBlkSubmissn,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

type BankTransactionCodeStructure4 struct {
	Domn  BankTransactionCodeStructure5            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Domn,omitempty"`
	Prtry ProprietaryBankTransactionCodeStructure1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Prtry,omitempty"`
}

type BankTransactionCodeStructure5 struct {
	Cd   ExternalBankTransactionDomain1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	Fmly BankTransactionCodeStructure6      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Fmly"`
}

type BankTransactionCodeStructure6 struct {
	Cd        ExternalBankTransactionFamily1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	SubFmlyCd ExternalBankTransactionSubFamily1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 SubFmlyCd"`
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 PstlAdr,omitempty"`
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Id"`
	Tp  CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Nm,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Prtry"`
}

type Channel2Choice struct {
	Cd    CommunicationMethod3Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	Prtry Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 MmbId"`
}

type CodeOrProprietary1Choice struct {
	Cd    Max4Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Prtry"`
}

type CommunicationFormat1Choice struct {
	Cd    ExternalCommunicationFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	Prtry Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Prtry"`
}

type CommunicationMethod2Choice struct {
	Cd    CommunicationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	Prtry Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Prtry"`
}

// May be one of EMAL, FAXI, FILE, ONLI, POST
type CommunicationMethod2Code string

// May be one of EMAL, FAXI, POST, PHON, FILE, ONLI
type CommunicationMethod3Code string

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Othr,omitempty"`
}

type ContractDocument1 struct {
	Ref      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Ref"`
	SgnOffDt ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 SgnOffDt,omitempty"`
	Vrsn     Max6Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Vrsn,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CustomerAccount4 struct {
	Id               AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Id,omitempty"`
	Nm               Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Nm,omitempty"`
	Sts              AccountStatus3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Sts,omitempty"`
	Tp               CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Tp,omitempty"`
	Ccy              ActiveCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Ccy"`
	MnthlyPmtVal     float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 MnthlyPmtVal,omitempty"`
	MnthlyRcvdVal    float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 MnthlyRcvdVal,omitempty"`
	MnthlyTxNb       Max5NumericText              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 MnthlyTxNb,omitempty"`
	AvrgBal          float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 AvrgBal,omitempty"`
	AcctPurp         Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 AcctPurp,omitempty"`
	FlrNtfctnAmt     float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 FlrNtfctnAmt,omitempty"`
	ClngNtfctnAmt    float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 ClngNtfctnAmt,omitempty"`
	StmtFrqcyAndFrmt []StatementFrequencyAndForm1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 StmtFrqcyAndFrmt,omitempty"`
	ClsgDt           ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 ClsgDt,omitempty"`
	Rstrctn          []Restriction1               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Rstrctn,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 CtryOfBirth"`
}

type Document struct {
	AcctOpngAmdmntReq AccountOpeningAmendmentRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 AcctOpngAmdmntReq"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalBankTransactionDomain1Code string

// Must be at least 1 items long
type ExternalBankTransactionFamily1Code string

// Must be at least 1 items long
type ExternalBankTransactionSubFamily1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalCommunicationFormat1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Prtry"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       BICFIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Othr,omitempty"`
}

type FixedAmountOrUnlimited1Choice struct {
	Amt    ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Amt"`
	NotLtd Unlimited9Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 NotLtd"`
}

// May be one of YEAR, DAIL, MNTH, QURT, MIAN, TEND, MOVE, WEEK, INDA
type Frequency7Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Issr"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Issr,omitempty"`
}

type Group1 struct {
	GrpId Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 GrpId"`
	Pty   []PartyAndCertificate2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Pty"`
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

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max140Text string

// Must match the pattern [\+]{0,1}[0-9]{1,15}
type Max15PlusSignedNumericText string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max4Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

type MaximumAmountByPeriod1 struct {
	MaxAmt   ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 MaxAmt"`
	NbOfDays Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 NbOfDays"`
}

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 CreDtTm"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OperationMandate2 struct {
	Id           Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Id"`
	AplblChanl   []Channel2Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 AplblChanl"`
	ReqrdSgntrNb Max15PlusSignedNumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 ReqrdSgntrNb"`
	SgntrOrdrInd bool                            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 SgntrOrdrInd"`
	MndtHldr     []PartyAndAuthorisation1        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 MndtHldr,omitempty"`
	BkOpr        []BankTransactionCodeStructure4 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 BkOpr"`
	StartDt      ISODate                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 StartDt,omitempty"`
	EndDt        ISODate                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 EndDt,omitempty"`
}

type Organisation12 struct {
	FullLglNm    Max350Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 FullLglNm"`
	TradgNm      Max350Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 TradgNm,omitempty"`
	CtryOfOpr    CountryCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 CtryOfOpr"`
	RegnDt       ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 RegnDt,omitempty"`
	OprlAdr      PostalAddress6              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 OprlAdr,omitempty"`
	BizAdr       PostalAddress6              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 BizAdr,omitempty"`
	LglAdr       PostalAddress6              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 LglAdr"`
	BllgAdr      PostalAddress6              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 BllgAdr,omitempty"`
	OrgId        OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 OrgId"`
	RprtvOffcr   []PartyIdentification40     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 RprtvOffcr,omitempty"`
	TrsrMgr      PartyIdentification40       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 TrsrMgr,omitempty"`
	MainMndtHldr []PartyIdentification40     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 MainMndtHldr,omitempty"`
	Sndr         []PartyIdentification40     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Sndr,omitempty"`
	LglRprtv     []PartyIdentification40     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 LglRprtv,omitempty"`
}

type OrganisationIdentification8 struct {
	AnyBIC AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Prtry"`
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 PrvtId"`
}

type PartyAndAuthorisation1 struct {
	PtyOrGrp  PartyOrGroup1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 PtyOrGrp"`
	SgntrOrdr Max15PlusSignedNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 SgntrOrdr,omitempty"`
	Authstn   Authorisation2             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Authstn"`
}

type PartyAndCertificate2 struct {
	Pty  PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Pty"`
	Cert Max10KBinary          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cert,omitempty"`
}

type PartyAndSignature2 struct {
	Pty   PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Pty"`
	Sgntr ProprietaryData3      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Sgntr"`
}

type PartyIdentification40 struct {
	Nm        Max140Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Nm,omitempty"`
	PstlAdr   PostalAddress6        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 PstlAdr,omitempty"`
	Id        PersonIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Id,omitempty"`
	CtryOfRes CountryCode           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 CtctDtls,omitempty"`
}

type PartyIdentification43 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 PstlAdr,omitempty"`
	Id        Party11Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 CtctDtls,omitempty"`
}

type PartyOrGroup1Choice struct {
	GrpId Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 GrpId"`
	Pty   PartyAndCertificate2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Pty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 AdrLine,omitempty"`
}

type ProprietaryBankTransactionCodeStructure1 struct {
	Cd   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Cd"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Issr,omitempty"`
}

type ProprietaryData3 struct {
	Item string `xml:",any"`
}

type References4 struct {
	MsgId       MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 MsgId"`
	PrcId       MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 PrcId"`
	AttchdDocNm []Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 AttchdDocNm,omitempty"`
}

type Restriction1 struct {
	RstrctnTp CodeOrProprietary1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 RstrctnTp"`
	VldFr     ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 VldFr"`
	VldUntil  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 VldUntil,omitempty"`
}

type StatementFrequencyAndForm1 struct {
	Frqcy    Frequency7Code             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Frqcy"`
	ComMtd   CommunicationMethod2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 ComMtd"`
	DlvryAdr Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 DlvryAdr"`
	Frmt     CommunicationFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Frmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// Must match the pattern UNLIMITED
type Unlimited9Text string

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
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

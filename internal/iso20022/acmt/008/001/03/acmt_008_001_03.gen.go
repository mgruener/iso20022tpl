// Code generated by main. DO NOT EDIT.

package acmt_008_001_03

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AccountContract2 struct {
	TrgtGoLiveDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 TrgtGoLiveDt,omitempty"`
	TrgtClsgDt   ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 TrgtClsgDt,omitempty"`
	UrgcyFlg     bool    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 UrgcyFlg,omitempty"`
}

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Othr"`
}

type AccountOpeningAmendmentRequestV03 struct {
	Refs             References4                                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Refs"`
	Fr               OrganisationIdentification29                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Fr,omitempty"`
	CtrctDts         AccountContract2                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 CtrctDts,omitempty"`
	UndrlygMstrAgrmt ContractDocument1                            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 UndrlygMstrAgrmt,omitempty"`
	Acct             CustomerAccount4                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Acct"`
	AcctSvcrId       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 AcctSvcrId"`
	Org              Organisation33                               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Org"`
	Mndt             []OperationMandate4                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Mndt,omitempty"`
	Grp              []Group4                                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Grp,omitempty"`
	RefAcct          CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 RefAcct,omitempty"`
	DgtlSgntr        []PartyAndSignature3                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 DgtlSgntr,omitempty"`
	SplmtryData      []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 SplmtryData,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry"`
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

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type Authorisation2 struct {
	MaxAmtByTx          FixedAmountOrUnlimited1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 MaxAmtByTx,omitempty"`
	MaxAmtByPrd         []MaximumAmountByPeriod1      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 MaxAmtByPrd,omitempty"`
	MaxAmtByBlkSubmissn FixedAmountOrUnlimited1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 MaxAmtByBlkSubmissn,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BankTransactionCodeStructure4 struct {
	Domn  BankTransactionCodeStructure5            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Domn,omitempty"`
	Prtry ProprietaryBankTransactionCodeStructure1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry,omitempty"`
}

type BankTransactionCodeStructure5 struct {
	Cd   ExternalBankTransactionDomain1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Fmly BankTransactionCodeStructure6      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Fmly"`
}

type BankTransactionCodeStructure6 struct {
	Cd        ExternalBankTransactionFamily1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	SubFmlyCd ExternalBankTransactionSubFamily1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 SubFmlyCd"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry"`
}

type Channel2Choice struct {
	Cd    CommunicationMethod3Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Prtry Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 MmbId"`
}

type CodeOrProprietary1Choice struct {
	Cd    Max4Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry"`
}

type CommunicationFormat1Choice struct {
	Cd    ExternalCommunicationFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Prtry Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry"`
}

type CommunicationMethod2Choice struct {
	Cd    CommunicationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Prtry Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry"`
}

// May be one of EMAL, FAXI, FILE, ONLI, POST
type CommunicationMethod2Code string

// May be one of EMAL, FAXI, POST, PHON, FILE, ONLI
type CommunicationMethod3Code string

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PrefrdMtd,omitempty"`
}

type ContractDocument1 struct {
	Ref      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Ref"`
	SgnOffDt ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 SgnOffDt,omitempty"`
	Vrsn     Max6Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Vrsn,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CustomerAccount4 struct {
	Id               AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id,omitempty"`
	Nm               Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Nm,omitempty"`
	Sts              AccountStatus3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Sts,omitempty"`
	Tp               CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Tp,omitempty"`
	Ccy              ActiveCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Ccy"`
	MnthlyPmtVal     float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 MnthlyPmtVal,omitempty"`
	MnthlyRcvdVal    float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 MnthlyRcvdVal,omitempty"`
	MnthlyTxNb       Max5NumericText              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 MnthlyTxNb,omitempty"`
	AvrgBal          float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 AvrgBal,omitempty"`
	AcctPurp         Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 AcctPurp,omitempty"`
	FlrNtfctnAmt     float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 FlrNtfctnAmt,omitempty"`
	ClngNtfctnAmt    float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 ClngNtfctnAmt,omitempty"`
	StmtFrqcyAndFrmt []StatementFrequencyAndForm1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 StmtFrqcyAndFrmt,omitempty"`
	ClsgDt           ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 ClsgDt,omitempty"`
	Rstrctn          []Restriction1               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Rstrctn,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 CtryOfBirth"`
}

type Document struct {
	AcctOpngAmdmntReq AccountOpeningAmendmentRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 AcctOpngAmdmntReq"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

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

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Othr,omitempty"`
}

type FixedAmountOrUnlimited1Choice struct {
	Amt    ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Amt"`
	NotLtd Unlimited9Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 NotLtd"`
}

// May be one of YEAR, DAIL, MNTH, QURT, MIAN, TEND, MOVE, WEEK, INDA
type Frequency7Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Issr"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Issr,omitempty"`
}

type Group4 struct {
	GrpId Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 GrpId"`
	Pty   []PartyAndCertificate4 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Pty"`
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

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max128Text string

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
	MaxAmt   ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 MaxAmt"`
	NbOfDays Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 NbOfDays"`
}

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 CreDtTm"`
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OperationMandate4 struct {
	Id           Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id"`
	AplblChanl   []Channel2Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 AplblChanl"`
	ReqrdSgntrNb Max15PlusSignedNumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 ReqrdSgntrNb"`
	SgntrOrdrInd bool                            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 SgntrOrdrInd"`
	MndtHldr     []PartyAndAuthorisation4        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 MndtHldr,omitempty"`
	BkOpr        []BankTransactionCodeStructure4 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 BkOpr"`
	StartDt      ISODate                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 StartDt,omitempty"`
	EndDt        ISODate                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 EndDt,omitempty"`
}

type Organisation33 struct {
	FullLglNm    Max350Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 FullLglNm"`
	TradgNm      Max350Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 TradgNm,omitempty"`
	CtryOfOpr    CountryCode                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 CtryOfOpr"`
	RegnDt       ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 RegnDt,omitempty"`
	OprlAdr      PostalAddress24              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 OprlAdr,omitempty"`
	BizAdr       PostalAddress24              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 BizAdr,omitempty"`
	LglAdr       PostalAddress24              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 LglAdr"`
	BllgAdr      PostalAddress24              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 BllgAdr,omitempty"`
	OrgId        OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 OrgId"`
	RprtvOffcr   []PartyIdentification137     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 RprtvOffcr,omitempty"`
	TrsrMgr      PartyIdentification137       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 TrsrMgr,omitempty"`
	MainMndtHldr []PartyIdentification137     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 MainMndtHldr,omitempty"`
	Sndr         []PartyIdentification137     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Sndr,omitempty"`
	LglRprtv     []PartyIdentification137     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 LglRprtv,omitempty"`
}

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PrvtId"`
}

type PartyAndAuthorisation4 struct {
	PtyOrGrp  PartyOrGroup2Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PtyOrGrp"`
	SgntrOrdr Max15PlusSignedNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 SgntrOrdr,omitempty"`
	Authstn   Authorisation2             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Authstn"`
}

type PartyAndCertificate4 struct {
	Pty  PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Pty"`
	Cert Max10KBinary           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cert,omitempty"`
}

type PartyAndSignature3 struct {
	Pty   PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Pty"`
	Sgntr SkipPayload            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Sgntr"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 CtctDtls,omitempty"`
}

type PartyIdentification137 struct {
	Nm        Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Nm,omitempty"`
	PstlAdr   PostalAddress24        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PstlAdr,omitempty"`
	Id        PersonIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id,omitempty"`
	CtryOfRes CountryCode            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 CtryOfRes,omitempty"`
	CtctDtls  Contact4               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 CtctDtls,omitempty"`
}

type PartyOrGroup2Choice struct {
	GrpId Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 GrpId"`
	Pty   PartyAndCertificate4 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Pty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type ProprietaryBankTransactionCodeStructure1 struct {
	Cd   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Issr,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Prtry"`
}

type References4 struct {
	MsgId       MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 MsgId"`
	PrcId       MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PrcId"`
	AttchdDocNm []Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 AttchdDocNm,omitempty"`
}

type Restriction1 struct {
	RstrctnTp CodeOrProprietary1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 RstrctnTp"`
	VldFr     ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 VldFr"`
	VldUntil  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 VldUntil,omitempty"`
}

type SkipPayload struct {
	Item string `xml:",any"`
}

type StatementFrequencyAndForm1 struct {
	Frqcy    Frequency7Code             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Frqcy"`
	ComMtd   CommunicationMethod2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 ComMtd"`
	DlvryAdr Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 DlvryAdr"`
	Frmt     CommunicationFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Frmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 Envlp"`
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
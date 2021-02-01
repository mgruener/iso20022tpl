// Code generated by main. DO NOT EDIT.

package camt_005_001_08

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountCashEntryReturnCriteria3 struct {
	NtryRefInd  bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NtryRefInd,omitempty"`
	AcctTpInd   bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AcctTpInd,omitempty"`
	NtryAmtInd  bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NtryAmtInd,omitempty"`
	AcctCcyInd  bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AcctCcyInd,omitempty"`
	NtryStsInd  bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NtryStsInd,omitempty"`
	NtryDtInd   bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NtryDtInd,omitempty"`
	AcctSvcrInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AcctSvcrInd,omitempty"`
	AcctOwnrInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AcctOwnrInd,omitempty"`
}

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Othr"`
}

type AccountIdentificationSearchCriteria2Choice struct {
	EQ     AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 EQ"`
	CTTxt  Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CTTxt"`
	NCTTxt Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NCTTxt"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Prtry"`
}

type ActiveAmountRange3Choice struct {
	ImpldCcyAndAmtRg ImpliedCurrencyAndAmountRange1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ImpldCcyAndAmtRg"`
	CcyAndAmtRg      ActiveCurrencyAndAmountRange3  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CcyAndAmtRg"`
}

type ActiveCurrencyAndAmountRange3 struct {
	Amt       ImpliedCurrencyAmountRange1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CdtDbtInd,omitempty"`
	Ccy       ActiveCurrencyCode                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Ccy"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricAmountRange2Choice struct {
	ImpldCcyAndAmtRg ImpliedCurrencyAndAmountRange1          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ImpldCcyAndAmtRg"`
	CcyAndAmtRg      ActiveOrHistoricCurrencyAndAmountRange2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CcyAndAmtRg"`
}

type ActiveOrHistoricCurrencyAndAmountRange2 struct {
	Amt       ImpliedCurrencyAmountRange1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CdtDbtInd,omitempty"`
	Ccy       ActiveOrHistoricCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Ccy"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Prtry"`
}

type AmountRangeBoundary1 struct {
	BdryAmt float64 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 BdryAmt"`
	Incl    bool    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Incl"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PstlAdr,omitempty"`
}

type CashAccountEntrySearch6 struct {
	AcctId     []AccountIdentificationSearchCriteria2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AcctId,omitempty"`
	NtryAmt    []ActiveOrHistoricAmountRange2Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NtryAmt,omitempty"`
	NtryAmtCcy []ActiveOrHistoricCurrencyCode               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NtryAmtCcy,omitempty"`
	CdtDbtInd  CreditDebitCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CdtDbtInd,omitempty"`
	NtrySts    []EntryStatus1Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NtrySts,omitempty"`
	NtryDt     []DateAndDateTimeSearch3Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NtryDt,omitempty"`
	AcctOwnr   PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AcctOwnr,omitempty"`
	AcctSvcr   BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AcctSvcr,omitempty"`
}

// May be one of PDNG, FINL
type CashPaymentStatus2Code string

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Prtry"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Cd"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 MmbId"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PrefrdMtd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeSearch3Choice struct {
	DtTmSch DateTimePeriod1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 DtTmSch"`
	DtSch   DatePeriodSearch1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 DtSch"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ToDt"`
}

type DatePeriodSearch1Choice struct {
	FrDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 FrDt"`
	ToDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ToDt"`
	FrToDt DatePeriod2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 FrToDt"`
	EQDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 EQDt"`
	NEQDt  ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NEQDt"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ToDtTm"`
}

type DateTimePeriod1Choice struct {
	FrDtTm ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 FrDtTm"`
	ToDtTm ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ToDtTm"`
	DtTmRg DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 DtTmRg"`
}

type Document struct {
	GetTx GetTransactionV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 GetTx"`
}

// May be one of BOOK, PDNG, FUTR
type EntryStatus1Code string

// Must match the pattern [BEOVW]{1,1}[0-9]{2,2}|DUM
type EntryTypeIdentifier string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashClearingSystem1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPaymentControlRequestType1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// May be one of STLD, RJTD, CAND, FNLD
type FinalStatusCode string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Othr,omitempty"`
}

type FromToAmountRange1 struct {
	FrAmt AmountRangeBoundary1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 FrAmt"`
	ToAmt AmountRangeBoundary1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ToAmt"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Issr,omitempty"`
}

type GetTransactionV08 struct {
	MsgHdr      MessageHeader9       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 MsgHdr"`
	TxQryDef    TransactionQuery5    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 TxQryDef,omitempty"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 SplmtryData,omitempty"`
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

type ImpliedCurrencyAmountRange1Choice struct {
	FrAmt   AmountRangeBoundary1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 FrAmt"`
	ToAmt   AmountRangeBoundary1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ToAmt"`
	FrToAmt FromToAmountRange1   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 FrToAmt"`
	EQAmt   float64              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 EQAmt"`
	NEQAmt  float64              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NEQAmt"`
}

type ImpliedCurrencyAndAmountRange1 struct {
	Amt       ImpliedCurrencyAmountRange1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CdtDbtInd,omitempty"`
}

// May be one of PBEN, TTIL, TFRO
type Instruction1Code string

type InstructionStatusReturnCriteria1 struct {
	PmtInstrStsInd     bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtInstrStsInd"`
	PmtInstrStsDtTmInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtInstrStsDtTmInd,omitempty"`
	PmtInstrStsRsnInd  bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtInstrStsRsnInd,omitempty"`
}

type InstructionStatusSearch5 struct {
	PmtInstrSts     PaymentStatusCodeSearch2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtInstrSts,omitempty"`
	PmtInstrStsDtTm DateTimePeriod1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtInstrStsDtTm,omitempty"`
	PrtryStsRsn     Max4AlphaNumericText           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PrtryStsRsn,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LongPaymentIdentification2 struct {
	TxId           Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 TxId,omitempty"`
	UETR           UUIDv4Identifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 UETR,omitempty"`
	IntrBkSttlmAmt float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IntrBkSttlmAmt"`
	IntrBkSttlmDt  ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IntrBkSttlmDt"`
	PmtMtd         PaymentOrigin1Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtMtd,omitempty"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstdAgt"`
	NtryTp         EntryTypeIdentifier                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NtryTp,omitempty"`
	EndToEndId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 EndToEndId,omitempty"`
}

// Must be at least 1 items long
type Max128Text string

// Must be at least 1 items long
type Max140Text string

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

// Must be at least 1 items long
type Max70Text string

type MessageHeader9 struct {
	MsgId   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 MsgId"`
	CreDtTm ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CreDtTm,omitempty"`
	ReqTp   RequestType4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ReqTp,omitempty"`
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Prtry"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PrvtId"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Agt"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CtctDtls,omitempty"`
}

type PaymentIdentification6Choice struct {
	TxId      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 TxId"`
	QId       QueueTransactionIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 QId"`
	LngBizId  LongPaymentIdentification2      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 LngBizId"`
	ShrtBizId ShortPaymentIdentification2     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ShrtBizId"`
	PrtryId   Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PrtryId"`
}

// May be one of BDT, BCT, CDT, CCT, CHK, BKT, DCP, CCP, RTI, CAN
type PaymentInstrument1Code string

type PaymentOrigin1Choice struct {
	FINMT    Max3NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 FINMT"`
	XMLMsgNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 XMLMsgNm"`
	Prtry    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Prtry"`
	Instrm   PaymentInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Instrm"`
}

type PaymentReturnCriteria4 struct {
	MsgIdInd            bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 MsgIdInd,omitempty"`
	ReqdExctnDtInd      bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ReqdExctnDtInd,omitempty"`
	InstrInd            bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstrInd,omitempty"`
	InstrStsRtrCrit     InstructionStatusReturnCriteria1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstrStsRtrCrit,omitempty"`
	InstdAmtInd         bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstdAmtInd,omitempty"`
	CdtDbtInd           bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CdtDbtInd,omitempty"`
	IntrBkSttlmAmtInd   bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IntrBkSttlmAmtInd,omitempty"`
	PrtyInd             bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PrtyInd,omitempty"`
	PrcgVldtyTmInd      bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PrcgVldtyTmInd,omitempty"`
	PurpInd             bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PurpInd,omitempty"`
	InstrCpyInd         bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstrCpyInd,omitempty"`
	PmtMTInd            bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtMTInd,omitempty"`
	PmtTpInd            bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtTpInd,omitempty"`
	TxIdInd             bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 TxIdInd,omitempty"`
	IntrBkSttlmDtInd    bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IntrBkSttlmDtInd,omitempty"`
	EndToEndIdInd       bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 EndToEndIdInd,omitempty"`
	PmtMtdInd           bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtMtdInd,omitempty"`
	DbtrInd             bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 DbtrInd,omitempty"`
	DbtrAgtInd          bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 DbtrAgtInd,omitempty"`
	InstgRmbrsmntAgtInd bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstgRmbrsmntAgtInd,omitempty"`
	InstdRmbrsmntAgtInd bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstdRmbrsmntAgtInd,omitempty"`
	IntrmyInd           bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IntrmyInd,omitempty"`
	CdtrAgtInd          bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CdtrAgtInd,omitempty"`
	CdtrInd             bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CdtrInd,omitempty"`
}

type PaymentSearch8 struct {
	MsgId             []Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 MsgId,omitempty"`
	ReqdExctnDt       []DateAndDateTimeSearch3Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 ReqdExctnDt,omitempty"`
	PmtId             []PaymentIdentification6Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtId,omitempty"`
	Sts               []InstructionStatusSearch5           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Sts,omitempty"`
	InstdAmt          []ActiveOrHistoricAmountRange2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstdAmt,omitempty"`
	InstdAmtCcy       []ActiveOrHistoricCurrencyCode       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstdAmtCcy,omitempty"`
	CdtDbtInd         CreditDebitCode                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CdtDbtInd,omitempty"`
	IntrBkSttlmAmt    []ActiveAmountRange3Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IntrBkSttlmAmt,omitempty"`
	IntrBkSttlmAmtCcy []ActiveCurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IntrBkSttlmAmtCcy,omitempty"`
	PmtMtd            []PaymentOrigin1Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtMtd,omitempty"`
	PmtTp             []PaymentType4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtTp,omitempty"`
	Prty              []Priority1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Prty,omitempty"`
	PrcgVldtyTm       []DateTimePeriod1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PrcgVldtyTm,omitempty"`
	Instr             []Instruction1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Instr,omitempty"`
	TxId              []Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 TxId,omitempty"`
	IntrBkSttlmDt     []ISODate                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IntrBkSttlmDt,omitempty"`
	EndToEndId        []Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 EndToEndId,omitempty"`
	Pties             PaymentTransactionParty3             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Pties,omitempty"`
}

type PaymentStatusCodeSearch2Choice struct {
	PdgSts       PendingStatus4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PdgSts"`
	FnlSts       FinalStatusCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 FnlSts"`
	PdgAndFnlSts CashPaymentStatus2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PdgAndFnlSts"`
}

type PaymentTransactionParty3 struct {
	InstgAgt         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstgAgt,omitempty"`
	InstdAgt         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstdAgt,omitempty"`
	UltmtDbtr        Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 UltmtDbtr,omitempty"`
	Dbtr             Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Dbtr,omitempty"`
	DbtrAgt          BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 DbtrAgt,omitempty"`
	InstgRmbrsmntAgt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstgRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstdRmbrsmntAgt,omitempty"`
	IntrmyAgt1       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IntrmyAgt1,omitempty"`
	IntrmyAgt2       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IntrmyAgt2,omitempty"`
	IntrmyAgt3       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IntrmyAgt3,omitempty"`
	CdtrAgt          BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CdtrAgt,omitempty"`
	Cdtr             Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Cdtr,omitempty"`
	UltmtCdtr        Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 UltmtCdtr,omitempty"`
}

// May be one of CBS, BCK, BAL, CLS, CTR, CBH, CBP, DPG, DPN, EXP, TCH, LMT, LIQ, DPP, DPH, DPS, STF, TRP, TCS, LOA, LOR, TCP, OND, MGL
type PaymentType3Code string

type PaymentType4Choice struct {
	Cd    PaymentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Cd"`
	Prtry Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Prtry"`
}

// May be one of ACPD, VALD, MATD, AUTD, INVD, UMAC, STLE, STLM, SSPD, PCAN, PSTL, PFST, SMLR, RMLR, SRBL, AVLB, SRML
type PendingStatus4Code string

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type Priority1Choice struct {
	Cd    Priority5Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Cd"`
	Prtry Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Prtry"`
}

// May be one of HIGH, LOWW, NORM, URGT
type Priority5Code string

// May be one of ALLL, CHNG, MODF, DELD
type QueryType2Code string

type QueueTransactionIdentification1 struct {
	QId    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 QId"`
	PosInQ Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PosInQ"`
}

// May be one of STND, PRPR
type ReportIndicator1Code string

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Enqry"`
	Prtry   GenericIdentification1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Prtry"`
}

type ShortPaymentIdentification2 struct {
	TxId          Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 TxId"`
	IntrBkSttlmDt ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 IntrBkSttlmDt"`
	InstgAgt      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 InstgAgt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemReturnCriteria2 struct {
	SysIdInd  bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 SysIdInd,omitempty"`
	MmbIdInd  bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 MmbIdInd,omitempty"`
	CtryIdInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 CtryIdInd,omitempty"`
	AcctIdInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AcctIdInd,omitempty"`
}

type SystemSearch4 struct {
	SysId  []ClearingSystemIdentification3Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 SysId,omitempty"`
	MmbId  []BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 MmbId,omitempty"`
	Ctry   CountryCode                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 Ctry,omitempty"`
	AcctId AccountIdentification4Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AcctId,omitempty"`
}

type TransactionCriteria5Choice struct {
	QryNm   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 QryNm"`
	NewCrit TransactionCriteria8 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NewCrit"`
}

type TransactionCriteria8 struct {
	NewQryNm Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 NewQryNm,omitempty"`
	SchCrit  []TransactionSearchCriteria8 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 SchCrit,omitempty"`
	StmtRpt  ReportIndicator1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 StmtRpt,omitempty"`
	RtrCrit  TransactionReturnCriteria5   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 RtrCrit,omitempty"`
}

type TransactionQuery5 struct {
	QryTp  QueryType2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 QryTp,omitempty"`
	TxCrit TransactionCriteria5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 TxCrit,omitempty"`
}

type TransactionReturnCriteria5 struct {
	PmtToRtrCrit       SystemReturnCriteria2           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtToRtrCrit,omitempty"`
	PmtFrRtrCrit       SystemReturnCriteria2           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtFrRtrCrit,omitempty"`
	AcctCshNtryRtrCrit AccountCashEntryReturnCriteria3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AcctCshNtryRtrCrit,omitempty"`
	PmtRtrCrit         PaymentReturnCriteria4          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtRtrCrit,omitempty"`
}

type TransactionSearchCriteria8 struct {
	PmtTo       []SystemSearch4         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtTo,omitempty"`
	PmtFr       []SystemSearch4         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtFr,omitempty"`
	PmtSch      PaymentSearch8          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 PmtSch,omitempty"`
	AcctNtrySch CashAccountEntrySearch6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 AcctNtrySch,omitempty"`
}

// Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier string

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

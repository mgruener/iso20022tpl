// Code generated by main. DO NOT EDIT.

package tsin_012_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// May be one of DEAC, HOLD, MDFY, REAC, OPEN, SYNC, VRFY
type AgreementItemAction1Code string

// May be one of HS25, HS38, HS51
type Algorithm5Code string

type AlgorithmAndDigest1 struct {
	DgstAlgo Algorithm5Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 DgstAlgo"`
	Dgst     Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Dgst"`
}

type AmountAndPeriod1 struct {
	Amt     ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Amt"`
	StartDt ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 StartDt,omitempty"`
	EndDt   ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 EndDt,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

type BinaryFile1 struct {
	MIMETp         Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 MIMETp,omitempty"`
	NcodgTp        Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 NcodgTp,omitempty"`
	CharSet        Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CharSet,omitempty"`
	InclBinryObjct Max100KBinary `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 InclBinryObjct,omitempty"`
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PstlAdr,omitempty"`
}

type BusinessApplicationHeader1 struct {
	CharSet    string                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CharSet,omitempty"`
	Fr         Party9Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Fr"`
	To         Party9Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 To"`
	BizMsgIdr  Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 BizMsgIdr"`
	MsgDefIdr  Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 MsgDefIdr"`
	BizSvc     Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 BizSvc,omitempty"`
	CreDt      ISONormalisedDateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CreDt"`
	CpyDplct   CopyDuplicate1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CpyDplct,omitempty"`
	PssblDplct bool                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PssblDplct,omitempty"`
	Prty       string                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Prty,omitempty"`
	Sgntr      SignatureEnvelope     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Sgntr,omitempty"`
}

type BusinessLetter1 struct {
	ApplCntxt   Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 ApplCntxt,omitempty"`
	LttrIdr     QualifiedDocumentInformation1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 LttrIdr"`
	Dt          ISODate                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Dt"`
	RltdLttr    []QualifiedDocumentInformation1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 RltdLttr,omitempty"`
	RltdMsg     []QualifiedDocumentInformation1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 RltdMsg,omitempty"`
	CnttIdr     []Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CnttIdr,omitempty"`
	InstrPrty   Priority3Code                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 InstrPrty,omitempty"`
	Orgtr       QualifiedPartyIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Orgtr"`
	PmryRcpt    []QualifiedPartyIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PmryRcpt"`
	Sndr        []QualifiedPartyIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Sndr,omitempty"`
	AuthstnUsr  []QualifiedPartyIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AuthstnUsr"`
	RspnRcpt    []QualifiedPartyIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 RspnRcpt,omitempty"`
	CpyRcpt     []QualifiedPartyIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CpyRcpt,omitempty"`
	OthrPty     []QualifiedPartyIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 OthrPty,omitempty"`
	AssoctdDoc  []QualifiedDocumentInformation1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AssoctdDoc,omitempty"`
	GovngCtrct  []QualifiedDocumentInformation1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 GovngCtrct,omitempty"`
	LglCntxt    []GovernanceRules2               `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 LglCntxt,omitempty"`
	AddtlInf    Max2000Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AddtlInf,omitempty"`
	Ntce        Max350Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Ntce,omitempty"`
	VldtnStsInf ValidationStatusInformation1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 VldtnStsInf,omitempty"`
	DgtlSgntr   []QualifiedPartyAndXMLSignature1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 DgtlSgntr,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Othr,omitempty"`
}

type Contacts3 struct {
	NmPrfx    NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 NmPrfx,omitempty"`
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Nm,omitempty"`
	PhneNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PhneNb,omitempty"`
	MobNb     PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 MobNb,omitempty"`
	FaxNb     PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 FaxNb,omitempty"`
	EmailAdr  Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 EmailAdr,omitempty"`
	Othr      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Othr,omitempty"`
	JobTitl   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 JobTitl,omitempty"`
	Rspnsblty Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Rspnsblty,omitempty"`
	Dept      Max70Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Dept,omitempty"`
}

// May be one of CODU, COPY, DUPL
type CopyDuplicate1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CountrySubdivision1Choice struct {
	Cd    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Cd"`
	Prtry GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Prtry"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CtryOfBirth"`
}

type Document struct {
	PtyRegnAndGrntAck PartyRegistrationAndGuaranteeAcknowledgementV01 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PtyRegnAndGrntAck"`
}

type EncapsulatedBusinessMessage1 struct {
	Hdr  BusinessApplicationHeader1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Hdr,omitempty"`
	Prfx string                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Prfx,omitempty"`
	Prtl bool                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Prtl"`
	Msg  StrictPayload              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Msg"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalDocumentPurpose1Code string

// Must be at least 1 items long
type ExternalDocumentType1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalTradeMarket1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Prtry"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       BICFIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Othr,omitempty"`
}

type FinancialItemParameters1 struct {
	Idr         Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Idr"`
	IsseDt      ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 IsseDt"`
	RltdItm     []string                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 RltdItm,omitempty"`
	DocPurp     ExternalDocumentPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 DocPurp,omitempty"`
	LangCd      string                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 LangCd,omitempty"`
	Issr        string                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Issr,omitempty"`
	Rcpt        string                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Rcpt,omitempty"`
	Buyr        string                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Buyr,omitempty"`
	Sellr       string                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Sellr,omitempty"`
	SellrFinAgt string                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 SellrFinAgt,omitempty"`
	BuyrFinAgt  string                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 BuyrFinAgt,omitempty"`
	GovngCtrct  []string                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 GovngCtrct,omitempty"`
	LglCntxt    string                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 LglCntxt,omitempty"`
	Ccy         CurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Ccy,omitempty"`
	DbtAcct     AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 DbtAcct,omitempty"`
	CdtAcct     AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CdtAcct,omitempty"`
	TradMkt     TradeMarket1Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 TradMkt,omitempty"`
}

type FinancingAgreementItem1 struct {
	ItmCntxt     FinancialItemParameters1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 ItmCntxt"`
	ItmActn      AgreementItemAction1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 ItmActn,omitempty"`
	PmtInstrm    PaymentInstrumentCode        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PmtInstrm,omitempty"`
	VldtnStsInf  ValidationStatusInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 VldtnStsInf,omitempty"`
	Ratg         bool                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Ratg"`
	ReopIndctn   bool                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 ReopIndctn"`
	Grnt         []GuaranteeDetails1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Grnt,omitempty"`
	GrntSts      ValidationStatusInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 GrntSts,omitempty"`
	RltdGrntLttr string                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 RltdGrntLttr,omitempty"`
	AssoctdDoc   []string                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AssoctdDoc,omitempty"`
	AddtlInf     []Max2000Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AddtlInf,omitempty"`
}

type FinancingAgreementList1 struct {
	Idr         Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Idr"`
	Dt          ISODate                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Dt"`
	RltdDoc     []string                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 RltdDoc,omitempty"`
	AgrmtRqstr  string                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AgrmtRqstr"`
	AgrmtRspndr string                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AgrmtRspndr"`
	GrntApplcnt string                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 GrntApplcnt"`
	GrntNbfcry  string                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 GrntNbfcry"`
	GrntIssr    string                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 GrntIssr"`
	NtfctnInf   []FinancingNotificationParties1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 NtfctnInf,omitempty"`
	Itm         []FinancingAgreementItem1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Itm"`
	ItmCnt      Max15NumericText                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 ItmCnt"`
	CtrlSum     float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CtrlSum,omitempty"`
	AddtlInf    Max2000Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AddtlInf,omitempty"`
	VldtnStsInf ValidationStatusInformation1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 VldtnStsInf,omitempty"`
}

type FinancingNotificationParties1 struct {
	NtifngPty  string   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 NtifngPty"`
	NtfctnRcvr string   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 NtfctnRcvr"`
	AckRcvr    []string `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AckRcvr,omitempty"`
}

// May be one of CA01, CA02, AC01, AC04, AC06, BE08, BE09, BE10, BE11, DT02, ID01, ID02, ID03, MI01, NA01, CA03
type FinancingStatusReason1Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Issr,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Issr,omitempty"`
}

type GovernanceIdentification1Choice struct {
	Cd    GovernanceIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Cd"`
	Prtry GenericIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Prtry"`
}

// May be one of ISPR, NONE, UCPR, URDG
type GovernanceIdentification1Code string

type GovernanceRules2 struct {
	Id       string                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id"`
	RuleId   GovernanceIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 RuleId"`
	AplblLaw Location1                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AplblLaw,omitempty"`
	Jursdctn []Location1                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Jursdctn,omitempty"`
}

type GuaranteeDetails1 struct {
	Issr       string                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Issr,omitempty"`
	Pos        int                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Pos,omitempty"`
	Desc       Max2000Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Desc,omitempty"`
	GrntedAmt  []AmountAndPeriod1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 GrntedAmt,omitempty"`
	Xcss       []AmountAndPeriod1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Xcss,omitempty"`
	CvrdPctg   []PercentageAndPeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CvrdPctg,omitempty"`
	AssoctdDoc []string               `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AssoctdDoc,omitempty"`
	AddtlInf   []Max2000Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AddtlInf,omitempty"`
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

type ISONormalisedDateTime time.Time

func (t *ISONormalisedDateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISONormalisedDateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type LegalOrganisation1 struct {
	Id Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id,omitempty"`
	Nm Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Nm,omitempty"`
}

type Location1 struct {
	Ctry        CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Ctry,omitempty"`
	CtrySubDvsn CountrySubdivision1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CtrySubDvsn,omitempty"`
	Txt         []Max2000Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Txt,omitempty"`
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2000Text string

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

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Adr,omitempty"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OrganisationIdentification6 struct {
	BIC  AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 BIC,omitempty"`
	Othr []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Othr,omitempty"`
}

type OrganisationIdentification7 struct {
	AnyBIC AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Prtry"`
}

type Party10Choice struct {
	OrgId  OrganisationIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PrvtId"`
}

type Party8Choice struct {
	OrgId  OrganisationIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PrvtId"`
}

type Party9Choice struct {
	OrgId PartyIdentification42                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 OrgId"`
	FIId  BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 FIId"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 NmAndAdr"`
}

type PartyIdentification42 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PstlAdr,omitempty"`
	Id        Party10Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CtctDtls,omitempty"`
}

type PartyIdentification45 struct {
	Id        Party8Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id,omitempty"`
	Nm        Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PstlAdr,omitempty"`
	CtryOfRes CountryCode    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CtryOfRes,omitempty"`
	CtctDtls  []Contacts3    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CtctDtls,omitempty"`
}

type PartyRegistrationAndGuaranteeAcknowledgementV01 struct {
	Hdr       BusinessLetter1                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Hdr"`
	AckList   []FinancingAgreementList1      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AckList"`
	AckCnt    Max15NumericText               `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AckCnt"`
	ItmCnt    Max15NumericText               `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 ItmCnt,omitempty"`
	CtrlSum   float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CtrlSum,omitempty"`
	AttchdMsg []EncapsulatedBusinessMessage1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AttchdMsg,omitempty"`
}

// May be one of BDT, BCT, CDT, CCT, CHK, BKT, DCP, CCP, RTI, CAN
type PaymentInstrumentCode string

type PercentageAndPeriod1 struct {
	Pctg    float64 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Pctg"`
	StartDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 StartDt,omitempty"`
	EndDt   ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 EndDt,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Ctry"`
}

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AdrLine,omitempty"`
}

// May be one of URGT, HIGH, NORM
type Priority3Code string

type QualifiedDocumentInformation1 struct {
	Id           string                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id"`
	Issr         string                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Issr,omitempty"`
	ItmListIdr   Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 ItmListIdr,omitempty"`
	ItmIdr       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 ItmIdr,omitempty"`
	Dt           ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Dt,omitempty"`
	Vrsn         Max6Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Vrsn,omitempty"`
	ElctrncOrgnl bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 ElctrncOrgnl"`
	Dgst         []AlgorithmAndDigest1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Dgst,omitempty"`
	DocTp        ExternalDocumentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 DocTp,omitempty"`
	URL          Max2048Text               `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 URL,omitempty"`
	AttchdFile   []BinaryFile1             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AttchdFile,omitempty"`
}

type QualifiedPartyAndXMLSignature1 struct {
	Pty   string            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Pty,omitempty"`
	Sgntr SignatureEnvelope `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Sgntr"`
}

type QualifiedPartyIdentification1 struct {
	Id       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Id"`
	Pty      []SingleQualifiedPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Pty"`
	ShrtId   PartyIdentification2Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 ShrtId,omitempty"`
	Role     GenericIdentification1                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Role,omitempty"`
	RoleDesc Max256Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 RoleDesc,omitempty"`
}

type SignatureEnvelope struct {
	Item string `xml:",any"`
}

type SingleQualifiedPartyIdentification1 struct {
	BasePty TradeParty1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 BasePty"`
	RltvIdr []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 RltvIdr,omitempty"`
}

type StatusReason4Choice struct {
	Cd    FinancingStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Cd"`
	Prtry Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Prtry"`
}

type StrictPayload struct {
	Item string `xml:",any"`
}

// May be one of NONE, MASA, MISA, SISA, IISA, CUYP, PRYP, ASTR, EMPY, EMCY, EPRY, ECYE, NFPI, NFQP, DECP, IRAC, IRAR, KEOG, PFSP, 401K, SIRA, 403B, 457X, RIRA, RIAN, RCRF, RCIP, EIFP, EIOP
type TaxExemptReason1Code string

type TaxExemptionReasonFormatChoice struct {
	Ustrd Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Ustrd"`
	Strd  TaxExemptReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Strd"`
}

type TaxParty3 struct {
	TaxId       Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 TaxId,omitempty"`
	TaxTp       Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 TaxTp,omitempty"`
	RegnId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 RegnId,omitempty"`
	TaxXmptnRsn []TaxExemptionReasonFormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 TaxXmptnRsn,omitempty"`
}

// May be one of RCCF, RCER
type TechnicalValidationStatus1Code string

type TradeMarket1Choice struct {
	Cd    ExternalTradeMarket1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Cd"`
	Prtry GenericIdentification20  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Prtry"`
}

type TradeParty1 struct {
	PtyId  PartyIdentification45 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 PtyId"`
	LglOrg LegalOrganisation1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 LglOrg,omitempty"`
	TaxPty []TaxParty3           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 TaxPty,omitempty"`
}

type ValidationStatusInformation1 struct {
	Sts            TechnicalValidationStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 Sts"`
	StsRsn         StatusReason4Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 StsRsn,omitempty"`
	AddtlStsRsnInf []Max105Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.012.001.01 AddtlStsRsnInf,omitempty"`
}

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
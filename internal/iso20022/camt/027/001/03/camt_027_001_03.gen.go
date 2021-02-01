// Code generated by main. DO NOT EDIT.

package camt_027_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Prtry"`
}

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
type BICIdentifier string

type BranchAndFinancialInstitutionIdentification4 struct {
	FinInstnId FinancialInstitutionIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 PstlAdr,omitempty"`
}

type Case2 struct {
	Id             Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Id"`
	Cretr          Party7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Cretr"`
	ReopCaseIndctn bool         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 ReopCaseIndctn,omitempty"`
}

type CaseAssignment2 struct {
	Id      Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Id"`
	Assgnr  Party7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Assgnr"`
	Assgne  Party7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Assgne"`
	CreDtTm ISODateTime  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 CreDtTm"`
}

type CashAccount16 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Id"`
	Tp  CashAccountType2             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Nm,omitempty"`
}

type CashAccountType2 struct {
	Cd    CashAccountType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

type ClaimNonReceiptV03 struct {
	Assgnmt   CaseAssignment2              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Assgnmt"`
	Case      Case2                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Case"`
	Undrlyg   UnderlyingTransaction1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Undrlyg"`
	CoverDtls MissingCover2                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 CoverDtls,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Othr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 CtryOfBirth"`
}

type Document struct {
	ClmNonRct ClaimNonReceiptV03 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 ClmNonRct"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Prtry"`
}

type FinancialInstitutionIdentification7 struct {
	BIC         BICIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 BIC,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Issr,omitempty"`
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
type Max70Text string

type MissingCover2 struct {
	MssngCoverInd bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 MssngCoverInd"`
	CoverCrrctn   SettlementInformation15 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 CoverCrrctn,omitempty"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OrganisationIdentification4 struct {
	BICOrBEI AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 BICOrBEI,omitempty"`
	Othr     []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Prtry"`
}

type OriginalGroupInformation3 struct {
	OrgnlMsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlMsgId"`
	OrgnlMsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlCreDtTm,omitempty"`
}

type Party6Choice struct {
	OrgId  OrganisationIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 PrvtId"`
}

type Party7Choice struct {
	Pty PartyIdentification32                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Pty"`
	Agt BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Agt"`
}

type PartyIdentification32 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 PstlAdr,omitempty"`
	Id        Party6Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 CtctDtls,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 AdrLine,omitempty"`
}

type SettlementInformation15 struct {
	InstgRmbrsmntAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct CashAccount16                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 InstdRmbrsmntAgtAcct,omitempty"`
}

type UnderlyingGroupInformation1 struct {
	OrgnlMsgId         Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlMsgId"`
	OrgnlMsgNmId       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlMsgNmId"`
	OrgnlCreDtTm       ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlCreDtTm,omitempty"`
	OrgnlMsgDlvryChanl Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlMsgDlvryChanl,omitempty"`
}

type UnderlyingPaymentInstruction1 struct {
	OrgnlGrpInf     UnderlyingGroupInformation1       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlGrpInf,omitempty"`
	OrgnlPmtInfId   Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlPmtInfId,omitempty"`
	OrgnlInstrId    Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlEndToEndId,omitempty"`
	OrgnlInstdAmt   ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlInstdAmt"`
	ReqdExctnDt     ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 ReqdExctnDt,omitempty"`
	ReqdColltnDt    ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 ReqdColltnDt,omitempty"`
}

type UnderlyingPaymentTransaction1 struct {
	OrgnlGrpInf         UnderlyingGroupInformation1       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlGrpInf,omitempty"`
	OrgnlInstrId        Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId     Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlEndToEndId,omitempty"`
	OrgnlTxId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlTxId,omitempty"`
	OrgnlIntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlIntrBkSttlmDt"`
}

type UnderlyingStatementEntry1 struct {
	OrgnlGrpInf OriginalGroupInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlGrpInf,omitempty"`
	OrgnlStmtId Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlStmtId,omitempty"`
	OrgnlNtryId Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 OrgnlNtryId,omitempty"`
}

type UnderlyingTransaction1Choice struct {
	Initn    UnderlyingPaymentInstruction1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 Initn"`
	IntrBk   UnderlyingPaymentTransaction1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 IntrBk"`
	StmtNtry UnderlyingStatementEntry1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.03 StmtNtry"`
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
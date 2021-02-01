// Code generated by main. DO NOT EDIT.

package camt_080_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Othr"`
}

type AccountIdentificationSearchCriteria2Choice struct {
	EQ     AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 EQ"`
	CTTxt  Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 CTTxt"`
	NCTTxt Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 NCTTxt"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Prtry"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 PstlAdr,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 MmbId"`
}

// May be one of CODU, COPY, DUPL
type CopyDuplicate1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 DtTm"`
}

type DateAndDateTimeSearch5Choice struct {
	Dt   DatePeriodSearch1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Dt"`
	DtTm DateTimeSearch2Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 DtTm"`
}

type DatePeriod2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 ToDt"`
}

type DatePeriodSearch1Choice struct {
	FrDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 FrDt"`
	ToDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 ToDt"`
	FrToDt DatePeriod2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 FrToDt"`
	EQDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 EQDt"`
	NEQDt  ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 NEQDt"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 ToDtTm"`
}

type DateTimeSearch2Choice struct {
	FrDtTm   ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 FrDtTm"`
	ToDtTm   ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 ToDtTm"`
	FrToDtTm DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 FrToDtTm"`
	EQDtTm   ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 EQDtTm"`
	NEQDtTm  ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 NEQDtTm"`
}

type Document struct {
	IntraBalMvmntModQry IntraBalanceMovementModificationQueryV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 IntraBalMvmntModQry"`
}

type DocumentIdentification51 struct {
	Id       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Id"`
	CreDtTm  DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 CreDtTm,omitempty"`
	CpyDplct CopyDuplicate1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 CpyDplct,omitempty"`
	MsgOrgtr PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 MsgOrgtr,omitempty"`
	MsgRcpt  PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 MsgRcpt,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 SchmeNm,omitempty"`
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

type IntraBalanceMovementModificationQueryV01 struct {
	Id          DocumentIdentification51     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Id,omitempty"`
	QryDef      IntraBalanceQueryDefinition8 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 QryDef"`
	SplmtryData []SupplementaryData1         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 SplmtryData,omitempty"`
}

type IntraBalanceQueryCriteria8 struct {
	ModReqId    []Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 ModReqId,omitempty"`
	PrcgSts     []ModificationProcessingStatus9Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 PrcgSts,omitempty"`
	CshAcct     []AccountIdentificationSearchCriteria2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 CshAcct,omitempty"`
	CshAcctOwnr []SystemPartyIdentification8                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 CshAcctOwnr,omitempty"`
	CshAcctSvcr BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 CshAcctSvcr,omitempty"`
	MsgOrgtr    []SystemPartyIdentification8                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 MsgOrgtr,omitempty"`
	CreDtTm     DateAndDateTimeSearch5Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 CreDtTm,omitempty"`
}

type IntraBalanceQueryDefinition8 struct {
	QryTp   MovementResponseType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 QryTp"`
	SchCrit IntraBalanceQueryCriteria8 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 SchCrit"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

// May be one of PACK, REJT, MODC, DEND, MODP, REPR
type ModificationProcessingStatus1Code string

type ModificationProcessingStatus9Choice struct {
	Cd    ModificationProcessingStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Cd"`
	Prtry GenericIdentification30           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Prtry"`
}

// May be one of FULL, STTS
type MovementResponseType1Code string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Adr,omitempty"`
}

type PartyIdentification120Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 NmAndAdr"`
}

type PartyIdentification136 struct {
	Id  PartyIdentification120Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 LEI,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Ctry"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 AdrLine,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemPartyIdentification8 struct {
	Id           PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 Id"`
	RspnsblPtyId PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.080.001.01 RspnsblPtyId,omitempty"`
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

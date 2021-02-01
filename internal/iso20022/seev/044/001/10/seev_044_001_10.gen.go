// Code generated by main. DO NOT EDIT.

package seev_044_001_10

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification10 struct {
	IdCd SafekeepingAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 IdCd"`
}

type AccountIdentification40Choice struct {
	ForAllAccts AccountIdentification10   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 ForAllAccts"`
	AcctsList   []AccountIdentification46 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 AcctsList"`
}

type AccountIdentification46 struct {
	SfkpgAcct Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 SfkpgAcct"`
	AcctOwnr  PartyIdentification127Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat28Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 SfkpgPlc,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type CorporateAction45 struct {
	DtDtls  CorporateActionDate59                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 DtDtls,omitempty"`
	EvtStag CorporateActionEventStageFormat14Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 EvtStag,omitempty"`
	LtryTp  LotteryTypeFormat4Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 LtryTp,omitempty"`
}

type CorporateActionDate59 struct {
	RcrdDt   DateFormat43Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 RcrdDt,omitempty"`
	ExDvddDt DateFormat43Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 ExDvddDt,omitempty"`
}

// May be one of FULL, PART, RESC
type CorporateActionEventStage4Code string

type CorporateActionEventStageFormat14Choice struct {
	Cd    CorporateActionEventStage4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Prtry"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH, ACCU, TNDP
type CorporateActionEventType32Code string

type CorporateActionEventType86Choice struct {
	Cd    CorporateActionEventType32Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Prtry"`
}

type CorporateActionGeneralInformation140 struct {
	CorpActnEvtId      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 OffclCorpActnEvtId,omitempty"`
	ClssActnNb         Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 ClssActnNb,omitempty"`
	EvtTp              CorporateActionEventType86Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 EvtTp"`
	MndtryVlntryEvtTp  CorporateActionMandatoryVoluntary3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 MndtryVlntryEvtTp"`
	FinInstrmId        SecurityIdentification19                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 FinInstrmId"`
}

// May be one of MAND, CHOS, VOLU
type CorporateActionMandatoryVoluntary1Code string

type CorporateActionMandatoryVoluntary3Choice struct {
	Cd    CorporateActionMandatoryVoluntary1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Cd"`
	Prtry GenericIdentification30                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Prtry"`
}

type CorporateActionMovementPreliminaryAdviceCancellationAdviceV10 struct {
	MvmntPrlimryAdvcId DocumentIdentification31             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 MvmntPrlimryAdvcId"`
	CorpActnGnlInf     CorporateActionGeneralInformation140 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 CorpActnGnlInf"`
	AcctDtls           AccountIdentification40Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 AcctDtls"`
	CorpActnDtls       CorporateAction45                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 CorpActnDtls,omitempty"`
	IssrAgt            []PartyIdentification120Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 IssrAgt,omitempty"`
	PngAgt             []PartyIdentification120Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 PngAgt,omitempty"`
	SubPngAgt          []PartyIdentification120Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 SubPngAgt,omitempty"`
	Regar              PartyIdentification120Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Regar,omitempty"`
	RsellngAgt         []PartyIdentification120Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 RsellngAgt,omitempty"`
	PhysSctiesAgt      PartyIdentification120Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 PhysSctiesAgt,omitempty"`
	DrpAgt             PartyIdentification120Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 DrpAgt,omitempty"`
	SlctnAgt           []PartyIdentification120Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 SlctnAgt,omitempty"`
	InfAgt             PartyIdentification120Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 InfAgt,omitempty"`
	SplmtryData        []SupplementaryData1                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 SplmtryData,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 DtTm"`
}

type DateCode19Choice struct {
	Cd    DateType8Code           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Prtry"`
}

type DateFormat43Choice struct {
	Dt   DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Dt"`
	DtCd DateCode19Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 DtCd"`
}

// May be one of UKWN, ONGO
type DateType8Code string

type Document struct {
	CorpActnMvmntPrlimryAdvcCxlAdvc CorporateActionMovementPreliminaryAdviceCancellationAdviceV10 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 CorpActnMvmntPrlimryAdvcCxlAdvc"`
}

type DocumentIdentification31 struct {
	Id    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Id"`
	LkgTp ProcessingPosition7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 LkgTp,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Id,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Prtry"`
}

// May be one of ORIG, SUPP
type LotteryType1Code string

type LotteryTypeFormat4Choice struct {
	Cd    LotteryType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Prtry"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Adr,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Tp"`
}

type PartyIdentification120Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 NmAndAdr"`
}

type PartyIdentification127Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 PrtryId"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Ctry"`
}

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

type ProcessingPosition7Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Prtry"`
}

// May be one of GENR
type SafekeepingAccountIdentification1Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat28Choice struct {
	Id      SafekeepingPlaceTypeAndText6           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Id"`
	Ctry    CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Ctry"`
	TpAndId SafekeepingPlaceTypeAndIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 TpAndId"`
	Prtry   GenericIdentification78                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Prtry"`
}

type SafekeepingPlaceTypeAndIdentification1 struct {
	SfkpgPlcTp SafekeepingPlace1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 SfkpgPlcTp"`
	Id         AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Id"`
}

type SafekeepingPlaceTypeAndText6 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Id,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Desc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.10 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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

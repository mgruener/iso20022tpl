// Code generated by main. DO NOT EDIT.

package seev_004_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateIdentification1 struct {
	Id    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Id"`
	IdSrc IdentificationSource1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 IdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type AttendanceCard2 struct {
	AttndncCardLbllg Max105Text         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 AttndncCardLbllg,omitempty"`
	DlvryMtd         DeliveryPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 DlvryMtd"`
	OthrAdr          NameAndAddress9    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 OthrAdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of EMPL, INDI, ENTR, OADR
type DeliveryPlace1Code string

type Document struct {
	MtgInstr MeetingInstructionV04 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 MtgInstr"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type Extension2 struct {
	PlcAndNm   Max350Text         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 PlcAndNm,omitempty"`
	XtnsnEnvlp ExtensionEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 XtnsnEnvlp"`
}

type ExtensionEnvelope1 struct {
	Item string `xml:",any"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Issr"`
}

type GenericIdentification4 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Id"`
	IdTp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 IdTp"`
}

type GenericIdentification5 struct {
	Issr  Max8Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Issr"`
	Inf   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Inf"`
	Nrrtv Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Nrrtv,omitempty"`
}

type HoldingBalance5 struct {
	Bal      UnitOrFaceAmountOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Bal"`
	BalTp    SecuritiesEntryType2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 BalTp,omitempty"`
	SfkpgPlc SafekeepingPlaceFormatChoice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 SfkpgPlc,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IdentificationSource1Choice struct {
	Dmst  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Dmst"`
	Prtry Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Prtry"`
}

type IndividualPerson17 struct {
	BirthNm         Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 BirthNm"`
	GvnNm           Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 GvnNm,omitempty"`
	Id              PersonIdentification6      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Id,omitempty"`
	Adr             LongPostalAddress2Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Adr,omitempty"`
	EmplngPty       PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 EmplngPty,omitempty"`
	AttndncCardDtls AttendanceCard2            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 AttndncCardDtls"`
}

type Instruction2 struct {
	InstrId       Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 InstrId"`
	ReqdExctnDt   ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 ReqdExctnDt,omitempty"`
	VoteExctnConf bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 VoteExctnConf"`
	AcctDtls      SafekeepingAccount4         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 AcctDtls"`
	Prxy          Proxy4                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Prxy,omitempty"`
	VoteDtls      VoteDetails2                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 VoteDtls,omitempty"`
	MtgAttndee    []IndividualPerson17        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 MtgAttndee,omitempty"`
	SpcfcInstrReq SpecificInstructionRequest1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 SpcfcInstrReq,omitempty"`
}

type LongPostalAddress2Choice struct {
	Ustrd Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Ustrd"`
	Strd  PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Strd"`
}

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

// Must be at least 1 items long
type Max8Text string

type MeetingInstructionV04 struct {
	Id       MessageIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Id"`
	MtgRef   MeetingReference4          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 MtgRef"`
	InstgPty PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 InstgPty"`
	SctyId   SecurityIdentification11   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 SctyId"`
	Instr    []Instruction2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Instr"`
	Xtnsn    []Extension2               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Xtnsn,omitempty"`
}

type MeetingReference4 struct {
	MtgId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 MtgId,omitempty"`
	IssrMtgId  Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 IssrMtgId,omitempty"`
	MtgDtAndTm ISODateTime                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 MtgDtAndTm"`
	Tp         MeetingType2Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Tp"`
	Clssfctn   MeetingTypeClassification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Clssfctn,omitempty"`
	Lctn       []PostalAddress1                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Lctn,omitempty"`
}

// May be one of XMET, GMET, MIXD, SPCL
type MeetingType2Code string

type MeetingTypeClassification1Choice struct {
	Cd    MeetingTypeClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Cd"`
	Prtry GenericIdentification13        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Prtry"`
}

// May be one of AMET, OMET, CLAS, ISSU, VRHI, CORT
type MeetingTypeClassification1Code string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Adr,omitempty"`
}

type NameAndAddress9 struct {
	Nm  Max350Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Nm"`
	Adr LongPostalAddress2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Adr,omitempty"`
}

type PartyIdentification3 struct {
	BICOrBEI AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 BICOrBEI"`
}

type PartyIdentification9Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 NmAndAdr"`
}

type PersonIdentification6 struct {
	Issr     Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Issr,omitempty"`
	PrsnIdTp PersonIdentificationType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 PrsnIdTp"`
}

type PersonIdentificationType1Choice struct {
	PsptNb      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 PsptNb"`
	TaxIdNb     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 TaxIdNb"`
	SclSctyNb   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 SclSctyNb"`
	MplyrIdNb   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 MplyrIdNb"`
	DrvrsLicNb  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 DrvrsLicNb"`
	AlnRegnNb   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 AlnRegnNb"`
	IdntyCardNb Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 IdntyCardNb"`
	OthrId      GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 OthrId"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Ctry"`
}

type Proxy4 struct {
	PrxyTp                ProxyType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 PrxyTp"`
	PrsnDtls              IndividualPerson17 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 PrsnDtls,omitempty"`
	VoteInstrForAgndRsltn Vote2Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 VoteInstrForAgndRsltn,omitempty"`
}

// May be one of CHRM, DISC, HLDR
type ProxyType2Code string

// May be one of QALL
type Quantity1Code string

type SafekeepingAccount4 struct {
	AcctId      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 AcctId"`
	AcctOwnr    PartyIdentification9Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 AcctOwnr,omitempty"`
	SubAcctDtls SubAccount2                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 SubAcctDtls,omitempty"`
	InstdBal    []HoldingBalance5            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 InstdBal"`
	RghtsHldr   []PartyIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 RghtsHldr,omitempty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

type SafekeepingPlaceAsCodeAndPartyIdentification struct {
	PlcSfkpg SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 PlcSfkpg"`
	Nrrtv    Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Nrrtv,omitempty"`
	Pty      PartyIdentification3  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Pty,omitempty"`
}

type SafekeepingPlaceFormatChoice struct {
	Id       SafekeepingPlaceAsCodeAndPartyIdentification `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Id"`
	IdAsDSS  GenericIdentification5                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 IdAsDSS"`
	IdAsCtry CountryCode                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 IdAsCtry"`
}

// May be one of BLOK, ELIG, PEND, PENR, NOMI, SETD, BORR, LOAN, SPOS, TRAD, COLI, COLO, UNBA, INBA, REGO
type SecuritiesEntryType2Code string

type SecurityIdentification11 struct {
	Id   SecurityIdentification11Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Id"`
	Desc Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Desc,omitempty"`
}

type SecurityIdentification11Choice struct {
	ISIN   ISINIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 ISIN"`
	OthrId AlternateIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 OthrId"`
}

type SpecificInstructionRequest1 struct {
	PrtcptnRegn bool `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 PrtcptnRegn,omitempty"`
	BlckgScties bool `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 BlckgScties,omitempty"`
	SctiesRegn  bool `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 SctiesRegn,omitempty"`
}

type SubAccount2 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Id"`
}

type UnitOrFaceAmountOrCode1Choice struct {
	Unit    float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Unit"`
	FaceAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 FaceAmt"`
	Cd      Quantity1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Cd"`
}

type Vote2Choice struct {
	VoteInstr    []Vote4 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 VoteInstr"`
	GblVoteInstr []Vote3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 GblVoteInstr"`
}

type Vote3 struct {
	IssrLabl Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 IssrLabl"`
	VoteOptn VoteInstruction2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 VoteOptn"`
}

type Vote4 struct {
	IssrLabl  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 IssrLabl"`
	For       float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 For,omitempty"`
	Agnst     float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Agnst,omitempty"`
	Abstn     float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Abstn,omitempty"`
	Wthhld    float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Wthhld,omitempty"`
	WthMgmt   float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 WthMgmt,omitempty"`
	AgnstMgmt float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 AgnstMgmt,omitempty"`
	Dscrtnry  float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Dscrtnry,omitempty"`
	NoActn    float64   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 NoActn,omitempty"`
}

type VoteDetails2 struct {
	VoteInstrForAgndRsltn Vote2Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 VoteInstrForAgndRsltn"`
	VoteInstrForMtgRsltn  VoteInstructionForMeetingResolution1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 VoteInstrForMtgRsltn,omitempty"`
}

// May be one of CFOR, CAGS, ABST, WTHH, WMGT, AMGT, NOAC, DISC
type VoteInstruction2Code string

// May be one of CHRM, CAGS, CFOR, ABST, WTHH, WMGT, AMGT, NOAC
type VoteInstructionAtMeeting1Code string

type VoteInstructionForMeetingResolution1Choice struct {
	VoteIndctn VoteInstructionAtMeeting1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 VoteIndctn"`
	Shrhldr    NameAndAddress9               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.04 Shrhldr"`
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

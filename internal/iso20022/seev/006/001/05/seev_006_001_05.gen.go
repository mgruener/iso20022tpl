// Code generated by main. DO NOT EDIT.

package seev_006_001_05

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AdditionalStatus1 struct {
	Rsn      InstructionRejectionStatus1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Rsn"`
	AddtlInf Max350Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 AddtlInf,omitempty"`
}

type AdditionalStatus2 struct {
	Rsn      CancellationRejectionStatus1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Rsn"`
	AddtlInf Max350Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 AddtlInf,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type CancellationProcessingStatus1 struct {
	Sts      CancellationStatus3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Sts"`
	AddtlInf Max350Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 AddtlInf,omitempty"`
}

type CancellationRejectionStatus1Choice struct {
	Cd    RejectionReason2Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Prtry"`
}

type CancellationStatus2Choice struct {
	PrcgSts  CancellationProcessingStatus1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 PrcgSts"`
	RjctnSts AdditionalStatus2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 RjctnSts"`
}

// May be one of PACK, CANP, CAND, RCIS
type CancellationStatus3Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DetailedInstructionStatus11 struct {
	InstrId   Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 InstrId"`
	AcctId    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 AcctId,omitempty"`
	SubAcctId Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 SubAcctId,omitempty"`
	InstrSts  InstructionStatus6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 InstrSts"`
}

type Document struct {
	MtgInstrSts MeetingInstructionStatusV05 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 MtgInstrSts"`
}

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Issr"`
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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Prtry"`
}

type InstructionProcessingStatus3 struct {
	Sts      Status7Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Sts"`
	AddtlInf Max350Text  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 AddtlInf,omitempty"`
}

type InstructionRejectionStatus1Choice struct {
	Cd    RejectionReason1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Prtry"`
}

type InstructionStatus5Choice struct {
	GblInstrSts  InstructionStatus6Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 GblInstrSts"`
	DtldInstrSts []DetailedInstructionStatus11 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 DtldInstrSts"`
}

type InstructionStatus6Choice struct {
	PrcgSts  InstructionProcessingStatus3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 PrcgSts"`
	RjctnSts AdditionalStatus1            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 RjctnSts"`
}

type InstructionType1Choice struct {
	InstrId    MessageIdentification `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 InstrId"`
	InstrCxlId MessageIdentification `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 InstrCxlId"`
}

type InstructionTypeStatus2Choice struct {
	InstrSts InstructionStatus5Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 InstrSts"`
	CxlSts   CancellationStatus2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 CxlSts"`
}

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

type MeetingInstructionStatusV05 struct {
	InstrTp     InstructionType1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 InstrTp"`
	MtgRef      MeetingReference7            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 MtgRef"`
	FinInstrmId SecurityIdentification14     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 FinInstrmId"`
	InstrTpSts  InstructionTypeStatus2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 InstrTpSts"`
	SplmtryData []SupplementaryData1         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 SplmtryData,omitempty"`
}

type MeetingReference7 struct {
	MtgId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 MtgId,omitempty"`
	IssrMtgId  Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 IssrMtgId,omitempty"`
	MtgDtAndTm ISODateTime                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 MtgDtAndTm"`
	Tp         MeetingType3Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Tp"`
	Clssfctn   MeetingTypeClassification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Clssfctn,omitempty"`
	Lctn       []PostalAddress1                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Lctn,omitempty"`
}

// May be one of XMET, GMET, MIXD, SPCL, BMET
type MeetingType3Code string

type MeetingTypeClassification1Choice struct {
	Cd    MeetingTypeClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Cd"`
	Prtry GenericIdentification13        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Prtry"`
}

// May be one of AMET, OMET, CLAS, ISSU, VRHI, CORT
type MeetingTypeClassification1Code string

type MessageIdentification struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Id"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Tp"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Ctry"`
}

// May be one of ADEA, LATE, DQUA, IPOS, LACK, SAFE, RBIS, EVNM, ULNK, PRXY, PART, SPLT, IPOA, IREG, DSEC
type RejectionReason1Code string

// May be one of ULNK, RBIS, INIR, ADEA, LATE, DCAN, DPRG
type RejectionReason2Code string

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Desc,omitempty"`
}

// May be one of CSUB, PACK, CAND, COMP, NOIN, RCIS, STIN
type Status7Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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

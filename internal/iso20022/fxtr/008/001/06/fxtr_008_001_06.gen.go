// Code generated by main. DO NOT EDIT.

package fxtr_008_001_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// May be one of POST, PREA, UNAL
type AllocationIndicator1Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type ClearingBrokerIdentification1 struct {
	SdInd     SideIndicator1Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 SdInd"`
	ClrBrkrId Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ClrBrkrId"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Prtry"`
}

// May be one of FULL, ONEW, PART, UNCO
type CollateralisationIndicator1Code string

// May be one of L, A, C, I, F, O, R, U
type CorporateSectorIdentifier1Code string

type CounterpartySideTransactionReporting1 struct {
	RptgJursdctn     Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 RptgJursdctn,omitempty"`
	RptgPty          PartyIdentification73Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 RptgPty,omitempty"`
	CtrPtySdUnqTxIdr []UniqueTransactionIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 CtrPtySdUnqTxIdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 DtTm"`
}

type Document struct {
	FXTradStsNtfctn ForeignExchangeTradeStatusNotificationV06 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 FXTradStsNtfctn"`
}

// May be no more than 42 items long
type Exact42Text string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type ForeignExchangeTradeStatusNotificationV06 struct {
	TradData    TradeData15          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 TradData"`
	RgltryRptg  RegulatoryReporting6 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 RgltryRptg,omitempty"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 SplmtryData,omitempty"`
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

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Prtry"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max10Text string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max20Text string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max52Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress8 struct {
	Nm         Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Nm"`
	Adr        PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Adr,omitempty"`
	AltrntvIdr []Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 AltrntvIdr,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Tp"`
}

type PartyIdentification44 struct {
	AnyBIC     AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 AnyBIC"`
	AltrntvIdr []Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 AltrntvIdr,omitempty"`
}

type PartyIdentification59 struct {
	PtyNm      Max34Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 PtyNm,omitempty"`
	AnyBIC     PartyIdentification44               `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 AnyBIC,omitempty"`
	AcctNb     Max34Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 AcctNb,omitempty"`
	Adr        Max105Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Adr,omitempty"`
	ClrSysId   ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ClrSysId,omitempty"`
	LglNttyIdr LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 LglNttyIdr,omitempty"`
}

type PartyIdentification73Choice struct {
	NmAndAdr NameAndAddress8       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 NmAndAdr"`
	AnyBIC   PartyIdentification44 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 AnyBIC"`
	PtyId    PartyIdentification59 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 PtyId"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Ctry"`
}

type RegulatoryReporting6 struct {
	TradgSdTxRptg          []TradingSideTransactionReporting1      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 TradgSdTxRptg,omitempty"`
	CtrPtySdTxRptg         []CounterpartySideTransactionReporting1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 CtrPtySdTxRptg,omitempty"`
	CntrlCtrPtyClrHs       PartyIdentification73Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 CntrlCtrPtyClrHs,omitempty"`
	ClrBrkr                PartyIdentification73Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ClrBrkr,omitempty"`
	ClrXcptnPty            PartyIdentification73Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ClrXcptnPty,omitempty"`
	ClrBrkrId              ClearingBrokerIdentification1           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ClrBrkrId,omitempty"`
	ClrThrshldInd          bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ClrThrshldInd,omitempty"`
	ClrdPdctId             Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ClrdPdctId,omitempty"`
	UndrlygPdctIdr         UnderlyingProductIdentifier1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 UndrlygPdctIdr,omitempty"`
	AllcnInd               AllocationIndicator1Code                `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 AllcnInd,omitempty"`
	CollstnInd             CollateralisationIndicator1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 CollstnInd,omitempty"`
	ExctnVn                Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ExctnVn,omitempty"`
	ExctnTmstmp            DateAndDateTimeChoice                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ExctnTmstmp,omitempty"`
	NonStdFlg              bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 NonStdFlg,omitempty"`
	LkSwpId                Exact42Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 LkSwpId,omitempty"`
	FinNtrOfTheCtrPtyInd   bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 FinNtrOfTheCtrPtyInd,omitempty"`
	CollPrtflInd           bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 CollPrtflInd,omitempty"`
	CollPrtflCd            Max10Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 CollPrtflCd,omitempty"`
	PrtflCmprssnInd        bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 PrtflCmprssnInd,omitempty"`
	CorpSctrInd            CorporateSectorIdentifier1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 CorpSctrInd,omitempty"`
	TradWthNonEEACtrPtyInd bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 TradWthNonEEACtrPtyInd,omitempty"`
	NtrgrpTradInd          bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 NtrgrpTradInd,omitempty"`
	ComrclOrTrsrFincgInd   bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ComrclOrTrsrFincgInd,omitempty"`
	FinInstrmId            SecurityIdentification19                `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 FinInstrmId,omitempty"`
	ConfDtAndTmstmp        ISODateTime                             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ConfDtAndTmstmp,omitempty"`
	ClrTmstmp              ISOTime                                 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ClrTmstmp,omitempty"`
	AddtlRptgInf           Max210Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 AddtlRptgInf,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Desc,omitempty"`
}

// May be one of CCPL, CLNT
type SideIndicator1Code string

type Status27Choice struct {
	Cd    TradeStatus6Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Cd"`
	Prtry Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Prtry"`
}

type Status28Choice struct {
	Cd    TradeStatus7Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Cd"`
	Prtry Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Prtry"`
}

type StatusAndSubStatus2 struct {
	StsCd    Status27Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 StsCd"`
	SubStsCd Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 SubStsCd,omitempty"`
}

// May be one of SMDY
type StatusSubType2Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeData15 struct {
	MsgId              Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 MsgId"`
	OrgtrRef           Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 OrgtrRef,omitempty"`
	MtchgSysUnqRef     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 MtchgSysUnqRef"`
	MtchgSysMtchgRef   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 MtchgSysMtchgRef,omitempty"`
	MtchgSysMtchdSdRef Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 MtchgSysMtchdSdRef,omitempty"`
	StsOrgtr           Max20Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 StsOrgtr,omitempty"`
	CurSts             StatusAndSubStatus2    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 CurSts"`
	CurStsSubTp        StatusSubType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 CurStsSubTp,omitempty"`
	CurStsDtTm         ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 CurStsDtTm,omitempty"`
	PrvsSts            Status28Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 PrvsSts,omitempty"`
	PrvsStsSubTp       StatusSubType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 PrvsStsSubTp,omitempty"`
	PrvsStsDtTm        ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 PrvsStsDtTm,omitempty"`
	PdctTp             Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 PdctTp,omitempty"`
	SttlmSsnIdr        Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 SttlmSsnIdr,omitempty"`
	SpltTradInd        bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 SpltTradInd,omitempty"`
}

// May be one of INVA, FMTC, SMAP, RJCT, RSCD, STLD, SPLI, UMTC, SMAT, FUMT, NETT, PFIX, OMTC
type TradeStatus6Code string

// May be one of INVA, UMTC, FMTC, SMAT, SUSP, SMAP, PFIX, FUMT
type TradeStatus7Code string

type TradingSideTransactionReporting1 struct {
	RptgJursdctn    Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 RptgJursdctn,omitempty"`
	RptgPty         PartyIdentification73Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 RptgPty,omitempty"`
	TradgSdUnqTxIdr []UniqueTransactionIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 TradgSdUnqTxIdr,omitempty"`
}

// May be one of FORW, NDFO, SPOT, SWAP
type UnderlyingProductIdentifier1Code string

type UniqueTransactionIdentifier2 struct {
	UnqTxIdr    Max52Text   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 UnqTxIdr"`
	PrrUnqTxIdr []Max52Text `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 PrrUnqTxIdr,omitempty"`
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

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

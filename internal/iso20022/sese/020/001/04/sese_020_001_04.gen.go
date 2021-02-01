// Code generated by main. DO NOT EDIT.

package sese_020_001_04

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

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AmountAndDirection8 struct {
	Amt                 ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 CdtDbtInd"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 OrgnlCcyAndOrdrdAmt,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 DtTm"`
}

// May be one of VARI
type DateType3Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesTxCxlReq SecuritiesTransactionCancellationRequestV04 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 SctiesTxCxlReq"`
}

type DocumentNumber1Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 LngNb"`
	PrtryNb GenericIdentification19           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FXCancellation1Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Ind"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Prtry"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 AmtsdVal"`
}

type GenericDocumentIdentification1 struct {
	MsgNb DocumentNumber1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 MsgNb,omitempty"`
	Id    Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Id"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Prtry"`
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
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Adr,omitempty"`
}

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Tp"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 PrtryId"`
}

type PartyIdentification37Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 NmAndAdr"`
	Ctry     CountryCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Ctry"`
}

type PartyIdentification43Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 NmAndAdr"`
}

type PartyIdentification44Choice struct {
	AnyBIC   AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 AnyBIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Ctry"`
}

type PartyIdentification46 struct {
	Id     PartyIdentification44Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Id"`
	PrcgId Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 PrcgId,omitempty"`
}

type PartyIdentificationAndAccount43 struct {
	Id        PartyIdentification43Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Id"`
	SfkpgAcct SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 SfkpgAcct,omitempty"`
	PrcgId    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 PrcgId,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Ctry"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 OrgnlAndCurFace"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type References2Choice struct {
	SctiesSttlmTxId SettlementTypeAndIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 SctiesSttlmTxId"`
	SctiesFincgTxId SettlementTypeAndIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 SctiesFincgTxId"`
	IntraPosMvmntId Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 IntraPosMvmntId"`
	OthrTxId        GenericDocumentIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 OthrTxId"`
}

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Nm,omitempty"`
}

type SecuritiesTransactionCancellationRequestV04 struct {
	AcctOwnrTxId      References2Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 AcctOwnrTxId"`
	AcctSvcrTxId      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 PrcrTxId,omitempty"`
	AcctOwnr          PartyIdentification36Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 AcctOwnr,omitempty"`
	SfkpgAcct         SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 SfkpgAcct"`
	TxDtls            TransactionDetails28        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 TxDtls,omitempty"`
	FxCxl             FXCancellation1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 FxCxl,omitempty"`
	SplmtryData       []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 SplmtryData,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Desc,omitempty"`
}

type SettlementDate1Choice struct {
	Dt   DateAndDateTimeChoice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Dt"`
	DtCd SettlementDateCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 DtCd"`
}

// May be one of WISS
type SettlementDate4Code string

type SettlementDateCode1Choice struct {
	Cd    SettlementDate4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Prtry"`
}

type SettlementParties12 struct {
	Dpstry PartyIdentification46           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount43 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount43 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount43 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount43 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount43 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Pty5,omitempty"`
}

type SettlementTypeAndIdentification3 struct {
	TxId          Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 TxId"`
	SctiesMvmntTp ReceiveDelivery1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 SctiesMvmntTp"`
	Pmt           DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Pmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeDate1Choice struct {
	Dt   DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Dt"`
	DtCd TradeDateCode1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 DtCd"`
}

type TradeDateCode1Choice struct {
	Cd    DateType3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Prtry"`
}

type TransactionDetails28 struct {
	FinInstrmId     SecurityIdentification14    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 FinInstrmId"`
	TradDt          TradeDate1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 TradDt,omitempty"`
	SttlmDt         SettlementDate1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 SttlmDt"`
	SttlmQty        Quantity6Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 SttlmQty"`
	SttlmAmt        AmountAndDirection8         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 SttlmAmt,omitempty"`
	DlvrgSttlmPties SettlementParties12         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties  SettlementParties12         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 RcvgSttlmPties,omitempty"`
	Invstr          PartyIdentification37Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.04 Invstr,omitempty"`
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
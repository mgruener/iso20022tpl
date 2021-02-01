// Code generated by main. DO NOT EDIT.

package sese_020_002_06_0

import (
	"bytes"
	"encoding/xml"
	"time"
)

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AmountAndDirection67 struct {
	Amt                 RestrictedFINActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Amt"`
	CdtDbtInd           CreditDebitCode                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 CdtDbtInd"`
	OrgnlCcyAndOrdrdAmt RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 OrgnlCcyAndOrdrdAmt,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type CancellationReason27 struct {
	Cd            CancellationReason37Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Cd"`
	CorpActnEvtId RestrictedFINMax16Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 CorpActnEvtId,omitempty"`
}

type CancellationReason37Choice struct {
	Cd    CancelledStatusReason16Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Cd"`
	Prtry GenericIdentification47     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Prtry"`
}

// May be one of SCEX, OTHR, CXLR, BYIY, CTHP, CANZ, CANT, CSUB, CANS, CANI, CORP
type CancelledStatusReason16Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 DtTm"`
}

// May be one of VARI
type DateType3Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesTxCxlReq SecuritiesTransactionCancellationRequest002V06 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SctiesTxCxlReq"`
}

type DocumentNumber16Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 ShrtNb"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 LngNb"`
	PrtryNb GenericIdentification163          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 PrtryNb"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FXCancellation4Choice struct {
	Ind   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Ind"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Prtry"`
}

type FinancialInstrumentQuantity15Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 AmtsdVal"`
}

type GenericDocumentIdentification6 struct {
	MsgNb DocumentNumber16Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 MsgNb,omitempty"`
	Id    RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Id"`
}

type GenericIdentification163 struct {
	Id      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Id"`
	Issr    Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Issr"`
	SchmeNm Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SchmeNm,omitempty"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SchmeNm,omitempty"`
}

type GenericIdentification84 struct {
	Id      RestrictedFINXMax34Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type IdentificationSource4Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Cd"`
	Prtry RestrictedFINExact2Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Prtry"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

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

type NameAndAddress12 struct {
	Nm RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Nm"`
}

type OriginalAndCurrentQuantities4 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 AmtsdVal"`
}

type OtherIdentification2 struct {
	Id  RestrictedFINXMax31Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Sfx,omitempty"`
	Tp  IdentificationSource4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Tp"`
}

type PartyIdentification136Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 AnyBIC"`
	PrtryId GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 PrtryId"`
}

type PartyIdentification137Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 AnyBIC"`
	PrtryId  GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 PrtryId"`
	NmAndAdr NameAndAddress12        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 NmAndAdr"`
}

type PartyIdentification145Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 AnyBIC"`
	NmAndAdr NameAndAddress12        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 NmAndAdr"`
	Ctry     CountryCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Ctry"`
}

type PartyIdentification156 struct {
	Id  PartyIdentification136Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 LEI,omitempty"`
}

type PartyIdentification170 struct {
	Id  PartyIdentification176Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 LEI,omitempty"`
}

type PartyIdentification176Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 AnyBIC"`
	PrtryId  GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 PrtryId"`
	NmAndAdr NameAndAddress12        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 NmAndAdr"`
	Ctry     CountryCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Ctry"`
}

type PartyIdentification191 struct {
	Id     PartyIdentification145Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Id"`
	LEI    LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 LEI,omitempty"`
	PrcgId RestrictedFINXMax16Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 PrcgId,omitempty"`
}

type PartyIdentificationAndAccount190 struct {
	Id        PartyIdentification137Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Id"`
	LEI       LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 LEI,omitempty"`
	SfkpgAcct SecuritiesAccount30          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SfkpgAcct,omitempty"`
	PrcgId    RestrictedFINXMax16Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 PrcgId,omitempty"`
}

type Quantity10Choice struct {
	Qty             FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities4       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 OrgnlAndCurFace"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type References60Choice struct {
	SctiesSttlmTxId SettlementTypeAndIdentification22 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SctiesSttlmTxId"`
	SctiesFincgTxId SettlementTypeAndIdentification22 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SctiesFincgTxId"`
	IntraPosMvmntId RestrictedFINXMax16Text           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 IntraPosMvmntId"`
	OthrTxId        GenericDocumentIdentification6    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 OthrTxId"`
}

type RestrictedFINActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

type RestrictedFINActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern XX|TS
type RestrictedFINExact2Text string

// Must match the pattern ([^/]+/)+([^/]+)|([^/]*)
type RestrictedFINMax16Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,140}
type RestrictedFINXMax140Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax16Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,31}
type RestrictedFINXMax31Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax34Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,35}
type RestrictedFINXMax35Text string

type SecuritiesAccount30 struct {
	Id RestrictedFINXMax35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Id"`
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Nm,omitempty"`
}

type SecuritiesTransactionCancellationRequest002V06 struct {
	AcctOwnrTxId      References60Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 AcctOwnrTxId"`
	AcctSvcrTxId      RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 PrcrTxId,omitempty"`
	AcctOwnr          PartyIdentification156  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 AcctOwnr,omitempty"`
	SfkpgAcct         SecuritiesAccount30     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SfkpgAcct"`
	TxDtls            TransactionDetails120   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 TxDtls,omitempty"`
	CxlRsn            CancellationReason27    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 CxlRsn,omitempty"`
	FxCxl             FXCancellation4Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 FxCxl,omitempty"`
	SplmtryData       []SupplementaryData1    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SplmtryData,omitempty"`
}

type SecurityIdentification20 struct {
	ISIN   ISINOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 ISIN,omitempty"`
	OthrId []OtherIdentification2   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 OthrId,omitempty"`
	Desc   RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Desc,omitempty"`
}

type SettlementDate20Choice struct {
	Dt   DateAndDateTime2Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Dt"`
	DtCd SettlementDateCode9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 DtCd"`
}

// May be one of WISS
type SettlementDate4Code string

type SettlementDateCode9Choice struct {
	Cd    SettlementDate4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Prtry"`
}

type SettlementParties90 struct {
	Dpstry PartyIdentification191           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount190 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount190 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount190 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount190 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount190 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Pty5,omitempty"`
}

type SettlementTypeAndIdentification22 struct {
	TxId          RestrictedFINXMax16Text  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 TxId"`
	SctiesMvmntTp ReceiveDelivery1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SctiesMvmntTp"`
	Pmt           DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Pmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeDate9Choice struct {
	Dt   DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Dt"`
	DtCd TradeDateCode4Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 DtCd"`
}

type TradeDateCode4Choice struct {
	Cd    DateType3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Cd"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Prtry"`
}

type TransactionDetails120 struct {
	FinInstrmId     SecurityIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 FinInstrmId"`
	TradDt          TradeDate9Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 TradDt,omitempty"`
	SttlmDt         SettlementDate20Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SttlmDt"`
	SttlmQty        Quantity10Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SttlmQty"`
	SttlmAmt        AmountAndDirection67     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 SttlmAmt,omitempty"`
	DlvrgSttlmPties SettlementParties90      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties  SettlementParties90      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 RcvgSttlmPties,omitempty"`
	Invstr          PartyIdentification170   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.06 Invstr,omitempty"`
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

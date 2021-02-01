// Code generated by main. DO NOT EDIT.

package sese_040_001_02

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

type AmountAndDirection51 struct {
	Amt                 ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 CdtDbtInd"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 OrgnlCcyAndOrdrdAmt,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type ConsentOrRejectionReason4Choice struct {
	Cd    CounterpartyResponseStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Cd"`
	Prtry GenericIdentification30               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Prtry"`
}

type ConsentReason4 struct {
	Cd          ConsentOrRejectionReason4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Cd"`
	AddtlRsnInf Max210Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 AddtlRsnInf,omitempty"`
}

type ConsentStatus4Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 NoSpcfdRsn"`
	Rsn        []ConsentReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Rsn"`
}

// May be one of CPTR, CPCX, CPMD
type CounterpartyResponseStatusReason1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 DtTm"`
}

// May be one of VARI
type DateType3Code string

// May be one of OPEN, UKWN
type DateType4Code string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesSttlmTxCtrPtyRspn SecuritiesSettlementTransactionCounterpartyResponseV02 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 SctiesSttlmTxCtrPtyRspn"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 AmtsdVal"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 SchmeNm,omitempty"`
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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Prtry"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max210Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Adr,omitempty"`
}

// May be one of NORE
type NoReasonCode string

type NoSpecifiedReason1 struct {
	NoSpcfdRsn NoReasonCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 NoSpcfdRsn"`
}

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Tp"`
}

type PartyIdentification44Choice struct {
	AnyBIC   AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 AnyBIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Ctry"`
}

type PartyIdentification71Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 NmAndAdr"`
}

type PartyIdentification91 struct {
	Id     PartyIdentification44Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Id"`
	LEI    LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 LEI,omitempty"`
	PrcgId Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 PrcgId,omitempty"`
}

type PartyIdentification93Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 NmAndAdr"`
	Ctry     CountryCode             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Ctry"`
}

type PartyIdentification99 struct {
	Id  PartyIdentification93Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 LEI,omitempty"`
}

type PartyIdentificationAndAccount117 struct {
	Id        PartyIdentification71Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Id"`
	LEI       LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 LEI,omitempty"`
	SfkpgAcct SecuritiesAccount19         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 SfkpgAcct,omitempty"`
	PrcgId    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 PrcgId,omitempty"`
}

type PendingStatus20Choice struct {
	Fwdd        NoSpecifiedReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Fwdd"`
	UdrInvstgtn NoSpecifiedReason1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 UdrInvstgtn"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Ctry"`
}

type Quantity6Choice struct {
	Qty             FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Qty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 OrgnlAndCurFace"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type RejectionReason29 struct {
	Cd          ConsentOrRejectionReason4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Cd"`
	AddtlRsnInf Max210Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 AddtlRsnInf,omitempty"`
}

type RejectionStatus20Choice struct {
	NoSpcfdRsn NoReasonCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 NoSpcfdRsn"`
	Rsn        RejectionReason29 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Rsn"`
}

type ResponseStatus6Choice struct {
	Cnsntd ConsentStatus4Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Cnsntd"`
	Rjctd  RejectionStatus20Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Rjctd"`
	Pdg    PendingStatus20Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Pdg"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Nm,omitempty"`
}

type SecuritiesSettlementTransactionCounterpartyResponseV02 struct {
	TxId        TransactionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 TxId"`
	RspnSts     ResponseStatus6Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 RspnSts"`
	TxDtls      TransactionDetails82       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 TxDtls,omitempty"`
	SplmtryData []SupplementaryData1       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 SplmtryData,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Desc,omitempty"`
}

type SettlementDate10Choice struct {
	Dt   DateAndDateTimeChoice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Dt"`
	DtCd SettlementDateCode8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 DtCd"`
}

type SettlementDateCode8Choice struct {
	Cd    DateType4Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Prtry"`
}

type SettlementParties40 struct {
	Dpstry PartyIdentification91            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Dpstry,omitempty"`
	Pty1   PartyIdentificationAndAccount117 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Pty1,omitempty"`
	Pty2   PartyIdentificationAndAccount117 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Pty2,omitempty"`
	Pty3   PartyIdentificationAndAccount117 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Pty3,omitempty"`
	Pty4   PartyIdentificationAndAccount117 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Pty4,omitempty"`
	Pty5   PartyIdentificationAndAccount117 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Pty5,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeDate5Choice struct {
	Dt   DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Dt"`
	DtCd TradeDateCode3Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 DtCd"`
}

type TradeDateCode3Choice struct {
	Cd    DateType3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Prtry"`
}

type TransactionDetails82 struct {
	FinInstrmId     SecurityIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 FinInstrmId"`
	SctiesMvmntTp   ReceiveDelivery1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 SctiesMvmntTp"`
	Pmt             DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Pmt"`
	SttlmQty        Quantity6Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 SttlmQty"`
	SfkpgAcct       SecuritiesAccount19      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 SfkpgAcct"`
	SttlmAmt        AmountAndDirection51     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 SttlmAmt,omitempty"`
	SttlmDt         SettlementDate10Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 SttlmDt"`
	TradDt          TradeDate5Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 TradDt,omitempty"`
	DlvrgSttlmPties SettlementParties40      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 DlvrgSttlmPties,omitempty"`
	RcvgSttlmPties  SettlementParties40      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 RcvgSttlmPties,omitempty"`
	Invstr          PartyIdentification99    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Invstr,omitempty"`
}

type TransactionIdentification6 struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 AcctOwnrTxId,omitempty"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 PrcrTxId,omitempty"`
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

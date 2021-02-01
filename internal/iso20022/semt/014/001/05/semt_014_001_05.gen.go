// Code generated by main. DO NOT EDIT.

package semt_014_001_05

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AcknowledgedAcceptedStatus21Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 NoSpcfdRsn"`
	Rsn        []AcknowledgementReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Rsn"`
}

type AcknowledgementReason12Choice struct {
	Cd    AcknowledgementReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Cd"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Prtry"`
}

// May be one of ADEA, SMPG, OTHR, CDCY, CDRG, CDRE, NSTP, RQWV, LATE
type AcknowledgementReason5Code string

type AcknowledgementReason9 struct {
	Cd          AcknowledgementReason12Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Cd"`
	AddtlRsnInf Max210Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AddtlRsnInf,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CancellationReason19Choice struct {
	Cd    CancelledStatusReason13Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Prtry"`
}

type CancellationReason9 struct {
	Cd          CancellationReason19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Cd"`
	AddtlRsnInf Max210Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AddtlRsnInf,omitempty"`
}

type CancellationStatus14Choice struct {
	NoSpcfdRsn NoReasonCode          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 NoSpcfdRsn"`
	Rsn        []CancellationReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Rsn"`
}

// May be one of CANI, CANS, CSUB, CXLR, CANT, CANZ, CORP, SCEX, OTHR, CTHP
type CancelledStatusReason13Code string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 DtTm"`
}

type Document struct {
	IntraPosMvmntStsAdvc IntraPositionMovementStatusAdviceV05 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 IntraPosMvmntStsAdvc"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

// May be one of AWMO, BYIY, CLAT, ADEA, CANR, CAIS, OBJT, AWSH, PHSE, STCD, DOCY, MLAT, DOCC, BLOC, CHAS, NEWI, CLAC, MUNO, GLOB, PREA, PART, NOFX, CMON, YCOL, COLL, DEPO, FLIM, INCA, LINK, LACK, LALO, MONY, NCON, REFS, SDUT, BATC, CYCL, SBLO, CPEC, MINO, IAAD, OTHR, PHCK, BENO, BOTH, CLHT, DENO, DISA, DKNY, FROZ, LAAW, LATE, LIQU, PRCY, REGT, SETS, CERT, PRSY, INBC
type FailingReason3Code string

type FailingReason7 struct {
	Cd          FailingReason7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Cd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AddtlRsnInf,omitempty"`
}

type FailingReason7Choice struct {
	Cd    FailingReason3Code      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Prtry"`
}

type FailingStatus9Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 NoSpcfdRsn"`
	Rsn        []FailingReason7 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Rsn"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AmtsdVal"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 SchmeNm,omitempty"`
}

type GenericIdentification37 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Id"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Issr,omitempty"`
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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Prtry"`
}

type IntraPositionDetails39 struct {
	PoolId        Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 PoolId,omitempty"`
	AcctOwnr      PartyIdentification92Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AcctOwnr,omitempty"`
	SfkpgAcct     SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 SfkpgAcct"`
	FinInstrmId   SecurityIdentification19           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 FinInstrmId"`
	SttlmQty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 SttlmQty"`
	LotNb         GenericIdentification37            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 LotNb,omitempty"`
	SttlmDt       DateAndDateTimeChoice              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 SttlmDt"`
	AckdStsTmStmp ISODateTime                        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AckdStsTmStmp,omitempty"`
	BalFr         SecuritiesBalanceType7Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 BalFr,omitempty"`
	BalTo         SecuritiesBalanceType7Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 BalTo,omitempty"`
}

type IntraPositionMovementStatusAdviceV05 struct {
	TxId        TransactionIdentifications29         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 TxId"`
	PrcgSts     IntraPositionProcessingStatus5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 PrcgSts,omitempty"`
	SttlmSts    SettlementStatus16Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 SttlmSts,omitempty"`
	TxDtls      IntraPositionDetails39               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 TxDtls,omitempty"`
	SplmtryData []SupplementaryData1                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 SplmtryData,omitempty"`
}

type IntraPositionProcessingStatus5Choice struct {
	Rjctd      RejectionOrRepairStatus29Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Rjctd"`
	Rpr        RejectionOrRepairStatus29Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Rpr"`
	Canc       CancellationStatus14Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Canc"`
	AckdAccptd AcknowledgedAcceptedStatus21Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AckdAccptd"`
	Prtry      ProprietaryStatusAndReason6        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Prtry"`
}

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

// May be one of NORE
type NoReasonCode string

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Tp"`
}

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 PrtryId"`
}

// May be one of AWMO, ADEA, CAIS, REFU, AWSH, PHSE, TAMM, DOCY, DOCC, BLOC, CHAS, NEWI, CLAC, MUNO, GLOB, PREA, PART, NMAS, NOFX, CMON, YCOL, COLL, DEPO, FLIM, INCA, LINK, FUTU, LACK, LALO, MONY, NCON, REFS, SDUT, BATC, CYCL, SBLO, CPEC, MINO, IAAD, OTHR, PHCK, BENO, BOTH, CLHT, DENO, DISA, DKNY, FROZ, LAAW, LATE, LIQU, PRCY, REGT, SETS, CERT, PRSY, INBC
type PendingReason10Code string

type PendingReason14 struct {
	Cd          PendingReason26Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Cd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AddtlRsnInf,omitempty"`
}

type PendingReason26Choice struct {
	Cd    PendingReason10Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Prtry"`
}

type PendingStatus36Choice struct {
	NoSpcfdRsn NoReasonCode      `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 NoSpcfdRsn"`
	Rsn        []PendingReason14 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Rsn"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 PrtryRsn,omitempty"`
}

type RejectionAndRepairReason23Choice struct {
	Cd    RejectionReason29Code   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Prtry"`
}

type RejectionOrRepairReason23 struct {
	Cd          []RejectionAndRepairReason23Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Cd,omitempty"`
	AddtlRsnInf Max210Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairStatus29Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 NoSpcfdRsn"`
	Rsn        []RejectionOrRepairReason23 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Rsn"`
}

// May be one of SAFE, ADEA, LATE, CAEV, DDAT, REFE, OTHR, DQUA, DSEC, INVB, INVL, INVN, MINO, MUNO, VALR
type RejectionReason29Code string

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Nm,omitempty"`
}

// May be one of BLOK, AWAS, AVAI, NOMI, PLED, REGO, RSTR, OTHR, SPOS, UNRG, ISSU, QUAS, COLA
type SecuritiesBalanceType13Code string

type SecuritiesBalanceType7Choice struct {
	Cd    SecuritiesBalanceType13Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Cd"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Prtry"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Desc,omitempty"`
}

type SettlementStatus16Choice struct {
	Pdg   PendingStatus36Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Pdg"`
	Flng  FailingStatus9Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Flng"`
	Prtry ProprietaryStatusAndReason6 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TransactionIdentifications29 struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AcctOwnrTxId"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 PrcrTxId,omitempty"`
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

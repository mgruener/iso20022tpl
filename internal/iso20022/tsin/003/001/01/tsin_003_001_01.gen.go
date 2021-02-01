// Code generated by main. DO NOT EDIT.

package tsin_003_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification3Choice struct {
	IBAN      IBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 IBAN"`
	BBAN      BBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 BBAN"`
	UPIC      UPICIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 UPIC"`
	PrtryAcct SimpleIdentificationInformation2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 PrtryAcct"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern AT[0-9]{5,5}
type AustrianBankleitzahlIdentifier string

// Must match the pattern [a-zA-Z0-9]{1,30}
type BBANIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BEIIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

// Must match the pattern CP[0-9]{4,4}
type CHIPSParticipantIdentifier string

// Must match the pattern CH[0-9]{6,6}
type CHIPSUniversalIdentifier string

// Must match the pattern CA[0-9]{9,9}
type CanadianPaymentsARNIdentifier string

type CancellationRequestInformation1 struct {
	OrgnlGrpId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 OrgnlGrpId"`
	OrgnlCreDtTm  ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 OrgnlCreDtTm"`
	NbOfInvcReqs  Max15NumericText                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 NbOfInvcReqs,omitempty"`
	TtlBlkInvcAmt ActiveCurrencyAndAmount             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 TtlBlkInvcAmt,omitempty"`
	CxlRsn        Max105Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 CxlRsn"`
	FincgRqstr    PartyIdentificationAndAccount6      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 FincgRqstr,omitempty"`
	IntrmyAgt     FinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 IntrmyAgt,omitempty"`
	FrstAgt       FinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 FrstAgt,omitempty"`
}

type CashAccount7 struct {
	Id  AccountIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 Id"`
	Tp  CashAccountType2             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 Tp,omitempty"`
	Ccy CurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 Nm,omitempty"`
}

type CashAccountType2 struct {
	Cd    CashAccountType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 Prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

type ClearingSystemMemberIdentification2Choice struct {
	USCHU       CHIPSUniversalIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 USCHU"`
	NZNCC       NewZealandNCCIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 NZNCC"`
	IENSC       IrishNSCIdentifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 IENSC"`
	GBSC        UKDomesticSortCodeIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 GBSC"`
	USCH        CHIPSParticipantIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 USCH"`
	CHBC        SwissBCIdentifier                              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 CHBC"`
	USFW        FedwireRoutingNumberIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 USFW"`
	PTNCC       PortugueseNCCIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 PTNCC"`
	RUCB        RussianCentralBankIdentificationCodeIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 RUCB"`
	ITNCC       ItalianDomesticIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 ITNCC"`
	ATBLZ       AustrianBankleitzahlIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 ATBLZ"`
	CACPA       CanadianPaymentsARNIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 CACPA"`
	CHSIC       SwissSICIdentifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 CHSIC"`
	DEBLZ       GermanBankleitzahlIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 DEBLZ"`
	ESNCC       SpanishDomesticInterbankingIdentifier          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 ESNCC"`
	ZANCC       SouthAfricanNCCIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 ZANCC"`
	HKNCC       HongKongBankIdentifier                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 HKNCC"`
	AUBSBx      ExtensiveBranchNetworkIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 AUBSBx"`
	AUBSBs      SmallNetworkIdentifier                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 AUBSBs"`
	INIFSC      IndianFinancialSystemCodeIdentifier            `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 INIFSC"`
	GRHEBIC     HellenicBankIdentificationCodeIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 GRHEBIC"`
	PLKNR       PolishNationalClearingCodeIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 PLKNR"`
	OthrClrCdId Max35Text                                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 OthrClrCdId"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type Document struct {
	InvcFincgCxlReq InvoiceFinancingCancellationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 InvcFincgCxlReq"`
}

// Must match the pattern AU[0-9]{6,6}
type ExtensiveBranchNetworkIdentifier string

// Must match the pattern FW[0-9]{9,9}
type FedwireRoutingNumberIdentifier string

type FinancialInstitutionIdentification6 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 ClrSysMmbId,omitempty"`
	PrtryId     GenericIdentification4                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 PrtryId,omitempty"`
	BIC         BICIdentifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 BIC,omitempty"`
}

type GenericIdentification4 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 Id"`
	IdTp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 IdTp"`
}

// Must match the pattern BL[0-9]{8,8}
type GermanBankleitzahlIdentifier string

// Must match the pattern GR[0-9]{7,7}
type HellenicBankIdentificationCodeIdentifier string

// Must match the pattern HK[0-9]{3,3}
type HongKongBankIdentifier string

// Must match the pattern [a-zA-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBANIdentifier string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must match the pattern IN[a-zA-Z0-9]{11,11}
type IndianFinancialSystemCodeIdentifier string

type InvoiceFinancingCancellationRequestV01 struct {
	CxlReqId  MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 CxlReqId"`
	CxlReqInf CancellationRequestInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 CxlReqInf"`
}

// Must match the pattern IE[0-9]{6,6}
type IrishNSCIdentifier string

// Must match the pattern IT[0-9]{10,10}
type ItalianDomesticIdentifier string

// Must be at least 1 items long
type Max105Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 CreDtTm"`
}

// Must match the pattern NZ[0-9]{6,6}
type NewZealandNCCIdentifier string

type PartyIdentification25 struct {
	Nm      Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 Nm"`
	PrtryId GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 PrtryId,omitempty"`
	BEI     BEIIdentifier          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 BEI,omitempty"`
}

type PartyIdentificationAndAccount6 struct {
	PtyId     PartyIdentification25 `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 PtyId"`
	CdtAcct   CashAccount7          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 CdtAcct,omitempty"`
	FincgAcct CashAccount7          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 FincgAcct,omitempty"`
}

// Must match the pattern PL[0-9]{8,8}
type PolishNationalClearingCodeIdentifier string

// Must match the pattern PT[0-9]{8,8}
type PortugueseNCCIdentifier string

// Must match the pattern RU[0-9]{9,9}
type RussianCentralBankIdentificationCodeIdentifier string

type SimpleIdentificationInformation2 struct {
	Id Max34Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 Id"`
}

// Must match the pattern AU[0-9]{6,6}
type SmallNetworkIdentifier string

// Must match the pattern ZA[0-9]{6,6}
type SouthAfricanNCCIdentifier string

// Must match the pattern ES[0-9]{8,9}
type SpanishDomesticInterbankingIdentifier string

// Must match the pattern SW[0-9]{3,5}
type SwissBCIdentifier string

// Must match the pattern SW[0-9]{6,6}
type SwissSICIdentifier string

// Must match the pattern SC[0-9]{6,6}
type UKDomesticSortCodeIdentifier string

// Must match the pattern [0-9]{8,17}
type UPICIdentifier string

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
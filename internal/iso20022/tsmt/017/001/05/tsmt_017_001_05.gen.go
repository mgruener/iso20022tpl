// Code generated by main. DO NOT EDIT.

package tsmt_017_001_05

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Prtry"`
}

// May be one of SBTW, RSTW, RSBS, ARDM, ARCS, ARES, WAIT, UPDT, SBDS, ARBA, ARRO, CINR
type Action2Code string

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type Adjustment6 struct {
	Tp    AdjustmentType1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tp"`
	Drctn AdjustmentDirection1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Drctn"`
	Amt   CurrencyAndAmount        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Amt"`
}

// May be one of ADDD, SUBS
type AdjustmentDirection1Code string

type AdjustmentType1Choice struct {
	Tp             AdjustmentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tp"`
	OthrAdjstmntTp Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 OthrAdjstmntTp"`
}

// May be one of REBA, DISC, CREN, SURC
type AdjustmentType2Code string

type AirportDescription1 struct {
	Twn      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Twn"`
	AirprtNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 AirprtNm,omitempty"`
}

type AirportName1Choice struct {
	AirprtCd       Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 AirprtCd"`
	OthrAirprtDesc AirportDescription1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 OthrAirprtDesc"`
}

type AmountOrPercentage2Choice struct {
	Amt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Amt"`
	Pctg float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Pctg"`
}

type BICIdentification1 struct {
	BIC BICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 BIC"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

// May be one of PROP, CLSD, PMTC, ESTD, ACTV, COMP, AMRQ, RARQ, CLRQ, SCRQ, SERQ, DARQ
type BaselineStatus3Code string

type CashAccount24 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Id"`
	Tp  CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Nm,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Prtry"`
}

type CertificateDataSet2 struct {
	DataSetId         DocumentIdentification1         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 DataSetId"`
	CertTp            TradeCertificateType1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CertTp"`
	LineItm           []LineItemAndPOIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 LineItm,omitempty"`
	CertfdChrtcs      CertifiedCharacteristics2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CertfdChrtcs"`
	IsseDt            ISODate                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 IsseDt"`
	PlcOfIsse         PostalAddress5                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PlcOfIsse,omitempty"`
	Issr              PartyIdentification26           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Issr"`
	InspctnDt         DatePeriodDetails               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 InspctnDt,omitempty"`
	AuthrsdInspctrInd bool                            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 AuthrsdInspctrInd,omitempty"`
	CertId            Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CertId"`
	Trnsprt           SingleTransport3                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Trnsprt,omitempty"`
	GoodsDesc         Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 GoodsDesc,omitempty"`
	Consgnr           PartyIdentification26           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Consgnr,omitempty"`
	Consgn            PartyIdentification26           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Consgn,omitempty"`
	Manfctr           PartyIdentification26           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Manfctr,omitempty"`
	AddtlInf          []Max350Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 AddtlInf,omitempty"`
}

type CertifiedCharacteristics2Choice struct {
	Orgn             CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Orgn"`
	Qlty             Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Qlty"`
	Anlys            Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Anlys"`
	Wght             Quantity9   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Wght"`
	Qty              Quantity9   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Qty"`
	HlthIndctn       bool        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 HlthIndctn"`
	PhytosntryIndctn bool        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PhytosntryIndctn"`
}

type Charge25 struct {
	Tp    FreightCharges1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tp"`
	Chrgs []ChargesDetails4   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Chrgs,omitempty"`
}

// May be one of SIGN, STDE, STOR, PACK, PICK, DNGR, SECU, INSU, COLF, CHOR, CHDE, AIRF, TRPT
type ChargeType8Code string

type ChargesDetails4 struct {
	ChrgsTp ChargesType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ChrgsTp"`
	Amt     CurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Amt"`
}

type ChargesType1Choice struct {
	Tp          ChargeType8Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tp"`
	OthrChrgsTp Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 OthrChrgsTp"`
}

type CommercialDataSet5 struct {
	DataSetId    DocumentIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 DataSetId"`
	ComrclDocRef InvoiceIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ComrclDocRef"`
	Buyr         PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Buyr"`
	Sellr        PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Sellr"`
	BllTo        PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 BllTo,omitempty"`
	Goods        []LineItem15            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Goods"`
	PmtTerms     []PaymentTerms4         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PmtTerms"`
	SttlmTerms   SettlementTerms3        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 SttlmTerms"`
}

type Consignment3 struct {
	TtlQty  Quantity10 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TtlQty,omitempty"`
	TtlVol  Quantity10 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TtlVol,omitempty"`
	TtlWght Quantity10 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TtlWght,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyAndAmount struct {
	Value float64      `xml:",chardata"`
	Ccy   CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DataSetSubmissionReferences4 struct {
	TxId              Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TxId"`
	PurchsOrdrRef     DocumentIdentification7   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PurchsOrdrRef"`
	UsrTxRef          []DocumentIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 UsrTxRef,omitempty"`
	ForcdMtch         bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ForcdMtch"`
	EstblishdBaselnId DocumentIdentification3   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 EstblishdBaselnId"`
	TxSts             BaselineStatus3Code       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TxSts"`
}

type DatePeriodDetails struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ToDt"`
}

type Document struct {
	FwdDataSetSubmissnRpt ForwardDataSetSubmissionReportV05 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 FwdDataSetSubmissnRpt"`
}

type DocumentIdentification1 struct {
	Id      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Id"`
	Vrsn    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Vrsn"`
	Submitr BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Submitr"`
}

type DocumentIdentification3 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Id"`
	Vrsn float64   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Vrsn"`
}

type DocumentIdentification5 struct {
	Id     Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Id"`
	IdIssr BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 IdIssr"`
}

type DocumentIdentification7 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Id"`
	DtOfIsse ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 DtOfIsse"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{7}
type Exact7NumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalIncoterms1Code string

type FinancialInstitutionIdentification4Choice struct {
	BIC      BICIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 BIC"`
	NmAndAdr NameAndAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 NmAndAdr"`
}

type ForwardDataSetSubmissionReportV05 struct {
	RptId           MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 RptId"`
	RltdTxRefs      []DataSetSubmissionReferences4  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 RltdTxRefs"`
	CmonSubmissnRef SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CmonSubmissnRef"`
	Submitr         BICIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Submitr"`
	BuyrBk          BICIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 BuyrBk"`
	SellrBk         BICIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 SellrBk"`
	ComrclDataSet   CommercialDataSet5              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ComrclDataSet,omitempty"`
	TrnsprtDataSet  TransportDataSet5               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TrnsprtDataSet,omitempty"`
	InsrncDataSet   InsuranceDataSet1               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 InsrncDataSet,omitempty"`
	CertDataSet     []CertificateDataSet2           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CertDataSet,omitempty"`
	OthrCertDataSet []OtherCertificateDataSet2      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 OthrCertDataSet,omitempty"`
	ReqForActn      PendingActivity2                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ReqForActn,omitempty"`
}

// May be one of CLCT, PRPD
type FreightCharges1Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Issr"`
}

type GenericIdentification4 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Id"`
	IdTp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 IdTp"`
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

type Incoterms4 struct {
	IncotrmsCd Incoterms4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 IncotrmsCd"`
	Lctn       Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Lctn,omitempty"`
}

type Incoterms4Choice struct {
	Cd    ExternalIncoterms1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Prtry"`
}

// May be one of ICCA, ICCB, ICCC, ICAI, IWCC, ISCC, IREC, ICLC, ISMC, CMCC, IRCE
type InsuranceClauses1Code string

type InsuranceDataSet1 struct {
	DataSetId      DocumentIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 DataSetId"`
	Issr           PartyIdentification26       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Issr"`
	IsseDt         ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 IsseDt"`
	FctvDt         ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 FctvDt,omitempty"`
	PlcOfIsse      PostalAddress5              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PlcOfIsse,omitempty"`
	InsrncDocId    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 InsrncDocId"`
	Trnsprt        SingleTransport3            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Trnsprt,omitempty"`
	InsrdAmt       CurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 InsrdAmt"`
	InsrdGoodsDesc Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 InsrdGoodsDesc,omitempty"`
	InsrncConds    []Max350Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 InsrncConds,omitempty"`
	InsrncClauses  []InsuranceClauses1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 InsrncClauses,omitempty"`
	Assrd          PartyIdentification29Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Assrd"`
	ClmsPyblAt     PostalAddress5              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ClmsPyblAt"`
	ClmsPyblIn     CurrencyCode                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ClmsPyblIn,omitempty"`
}

type InvoiceIdentification1 struct {
	InvcNb Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 InvcNb"`
	IsseDt ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 IsseDt"`
}

type LineItem15 struct {
	PurchsOrdrRef  DocumentIdentification7   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PurchsOrdrRef"`
	FnlSubmissn    bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 FnlSubmissn"`
	ComrclLineItms []LineItemDetails14       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ComrclLineItms"`
	LineItmsTtlAmt CurrencyAndAmount         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 LineItmsTtlAmt"`
	Adjstmnt       []Adjustment6             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Adjstmnt,omitempty"`
	FrghtChrgs     Charge25                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 FrghtChrgs,omitempty"`
	Tax            []Tax22                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tax,omitempty"`
	TtlNetAmt      CurrencyAndAmount         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TtlNetAmt"`
	BuyrDfndInf    []UserDefinedInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 BuyrDfndInf,omitempty"`
	SellrDfndInf   []UserDefinedInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 SellrDfndInf,omitempty"`
	Incotrms       Incoterms4                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Incotrms,omitempty"`
}

type LineItemAndPOIdentification1 struct {
	LineItmId     []Max70Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 LineItmId"`
	PurchsOrdrRef DocumentIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PurchsOrdrRef"`
}

type LineItemDetails14 struct {
	LineItmId  Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 LineItmId"`
	Qty        Quantity9                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Qty"`
	UnitPric   UnitPrice18                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 UnitPric,omitempty"`
	PdctNm     Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PdctNm,omitempty"`
	PdctIdr    []ProductIdentifier2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PdctIdr,omitempty"`
	PdctChrtcs []ProductCharacteristics1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PdctChrtcs,omitempty"`
	PdctCtgy   []ProductCategory1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PdctCtgy,omitempty"`
	PdctOrgn   CountryCode                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PdctOrgn,omitempty"`
	Adjstmnt   []Adjustment6                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Adjstmnt,omitempty"`
	FrghtChrgs Charge25                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 FrghtChrgs,omitempty"`
	Tax        []Tax22                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tax,omitempty"`
	TtlAmt     CurrencyAndAmount               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TtlAmt"`
}

// Must be at least 1 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CreDtTm"`
}

type MultimodalTransport3 struct {
	TakngInChrg  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TakngInChrg"`
	PlcOfFnlDstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PlcOfFnlDstn"`
}

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Adr"`
}

type OtherCertificateDataSet2 struct {
	DataSetId DocumentIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 DataSetId"`
	CertId    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CertId"`
	CertTp    Exact4AlphaNumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CertTp"`
	IsseDt    ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 IsseDt"`
	Issr      PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Issr"`
	CertInf   []Max350Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CertInf,omitempty"`
}

type PartyIdentification26 struct {
	Nm      Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Nm"`
	PrtryId GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PrtryId,omitempty"`
	PstlAdr PostalAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PstlAdr"`
}

type PartyIdentification29Choice struct {
	BIC      BICIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 BIC"`
	NmAndAdr PartyIdentification26 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 NmAndAdr"`
}

type PaymentCodeOrOther1Choice struct {
	PmtCd        PaymentPeriod3 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PmtCd"`
	PmtDueDt     ISODate        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PmtDueDt"`
	OthrPmtTerms Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 OthrPmtTerms"`
}

type PaymentPeriod3 struct {
	Cd       PaymentTime3Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Cd"`
	NbOfDays float64          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 NbOfDays,omitempty"`
}

type PaymentTerms4 struct {
	PmtTerms  PaymentCodeOrOther1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PmtTerms"`
	AmtOrPctg AmountOrPercentage2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 AmtOrPctg"`
}

// May be one of EMTD, EMTR, EPBE, EPRD, PRMD, PRMR, EPIN, EPAM, EPPO, EPRR, EPSD, CASH, IREC
type PaymentTime3Code string

type PendingActivity2 struct {
	Tp   Action2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tp"`
	Desc Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Desc,omitempty"`
}

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Ctry"`
}

type PostalAddress5 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PstCdId,omitempty"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TwnNm,omitempty"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Ctry"`
}

type ProductCategory1 struct {
	Tp   ProductCategory1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tp"`
	Ctgy Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Ctgy"`
}

type ProductCategory1Choice struct {
	StrdPdctCtgy ProductCategory1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 StrdPdctCtgy"`
	OthrPdctCtgy GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 OthrPdctCtgy"`
}

// May be one of HRTR, QOTA, PRGP, LOBU, GNDR
type ProductCategory1Code string

type ProductCharacteristics1 struct {
	Tp     ProductCharacteristics1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tp"`
	Chrtcs Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Chrtcs"`
}

type ProductCharacteristics1Choice struct {
	StrdPdctChrtcs ProductCharacteristics1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 StrdPdctChrtcs"`
	OthrPdctChrtcs GenericIdentification4  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 OthrPdctChrtcs"`
}

// May be one of BISP, CHNR, CLOR, EDSP, ENNR, OPTN, ORCR, PCTV, SISP, SIZE, SZRG, SPRM, STOR, VINR
type ProductCharacteristics1Code string

type ProductIdentifier2 struct {
	Tp  ProductIdentifier2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tp"`
	Idr Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Idr"`
}

type ProductIdentifier2Choice struct {
	StrdPdctIdr ProductIdentifier2     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 StrdPdctIdr"`
	OthrPdctIdr GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 OthrPdctIdr"`
}

// May be one of BINR, COMD, EANC, HRTR, MANI, MODL, PART, QOTA, STYL, SUPI, UPCC
type ProductIdentifier2Code string

type Quantity10 struct {
	UnitOfMeasr UnitOfMeasure3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 UnitOfMeasr"`
	Val         float64              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Val"`
}

type Quantity9 struct {
	UnitOfMeasr UnitOfMeasure3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 UnitOfMeasr"`
	Val         float64              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Val"`
	Fctr        Max15NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Fctr,omitempty"`
}

type SettlementTerms3 struct {
	CdtrAgt  FinancialInstitutionIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CdtrAgt,omitempty"`
	CdtrAcct CashAccount24                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CdtrAcct"`
}

type ShipmentDate1Choice struct {
	PropsdShipmntDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PropsdShipmntDt"`
	ActlShipmntDt   ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ActlShipmntDt"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Id"`
}

type SingleTransport3 struct {
	TrnsprtByAir  TransportByAir2  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TrnsprtByAir,omitempty"`
	TrnsprtBySea  TransportBySea4  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TrnsprtBySea,omitempty"`
	TrnsprtByRoad TransportByRoad2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TrnsprtByRoad,omitempty"`
	TrnsprtByRail TransportByRail2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TrnsprtByRail,omitempty"`
}

type SingleTransport8 struct {
	TrnsprtByAir  []TransportByAir4  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TrnsprtByAir,omitempty"`
	TrnsprtBySea  []TransportBySea5  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TrnsprtBySea,omitempty"`
	TrnsprtByRoad []TransportByRoad4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TrnsprtByRoad,omitempty"`
	TrnsprtByRail []TransportByRail4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TrnsprtByRail,omitempty"`
}

type Tax22 struct {
	Tp  TaxType2Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tp"`
	Amt CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Amt"`
}

type TaxType2Choice struct {
	Tp        TaxType9Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Tp"`
	OthrTaxTp Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 OthrTaxTp"`
}

// May be one of PROV, NATI, STAT, WITH, STAM, COAX, VATA, CUST
type TaxType9Code string

// May be one of ANLY, QUAL, QUAN, WEIG, ORIG, HEAL, PHYT
type TradeCertificateType1Code string

type TransportByAir2 struct {
	DprtureAirprt AirportName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 DprtureAirprt"`
	DstnAirprt    AirportName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 DstnAirprt"`
	AirCrrierNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 AirCrrierNm,omitempty"`
}

type TransportByAir4 struct {
	DprtureAirprt AirportName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 DprtureAirprt"`
	DstnAirprt    AirportName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 DstnAirprt"`
	FlghtNb       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 FlghtNb,omitempty"`
	AirCrrierNm   Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 AirCrrierNm,omitempty"`
	AirCrrierCtry CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 AirCrrierCtry,omitempty"`
	CrrierAgtNm   Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CrrierAgtNm,omitempty"`
	CrrierAgtCtry CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CrrierAgtCtry,omitempty"`
}

type TransportByRail2 struct {
	PlcOfRct     Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PlcOfRct"`
	PlcOfDlvry   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PlcOfDlvry"`
	RailCrrierNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 RailCrrierNm,omitempty"`
}

type TransportByRail4 struct {
	PlcOfRct       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PlcOfRct"`
	PlcOfDlvry     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PlcOfDlvry"`
	RailCrrierNm   Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 RailCrrierNm,omitempty"`
	RailCrrierCtry CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 RailCrrierCtry,omitempty"`
	CrrierAgtNm    Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CrrierAgtNm,omitempty"`
	CrrierAgtCtry  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CrrierAgtCtry,omitempty"`
}

type TransportByRoad2 struct {
	PlcOfRct     Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PlcOfRct"`
	PlcOfDlvry   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PlcOfDlvry"`
	RoadCrrierNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 RoadCrrierNm,omitempty"`
}

type TransportByRoad4 struct {
	PlcOfRct       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PlcOfRct"`
	PlcOfDlvry     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PlcOfDlvry"`
	RoadCrrierNm   Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 RoadCrrierNm,omitempty"`
	RoadCrrierCtry CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 RoadCrrierCtry,omitempty"`
	CrrierAgtNm    Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CrrierAgtNm,omitempty"`
	CrrierAgtCtry  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CrrierAgtCtry,omitempty"`
}

type TransportBySea4 struct {
	PortOfLoadng  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PortOfLoadng"`
	PortOfDschrge Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PortOfDschrge"`
	VsslNm        Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 VsslNm,omitempty"`
	SeaCrrierNm   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 SeaCrrierNm,omitempty"`
}

type TransportBySea5 struct {
	PortOfLoadng  Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PortOfLoadng"`
	PortOfDschrge Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PortOfDschrge"`
	VsslNm        Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 VsslNm,omitempty"`
	SeaCrrierNm   Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 SeaCrrierNm,omitempty"`
	SeaCrrierCtry CountryCode       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 SeaCrrierCtry,omitempty"`
	CrrierAgtNm   Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CrrierAgtNm,omitempty"`
	CrrierAgtCtry CountryCode       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 CrrierAgtCtry,omitempty"`
	MstrNm        Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 MstrNm,omitempty"`
	ChrtrrNm      Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ChrtrrNm,omitempty"`
	OwnrNm        Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 OwnrNm,omitempty"`
	IMONb         Exact7NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 IMONb,omitempty"`
	VygNb         Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 VygNb,omitempty"`
}

type TransportDataSet5 struct {
	DataSetId  DocumentIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 DataSetId"`
	Buyr       PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Buyr,omitempty"`
	Sellr      PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Sellr,omitempty"`
	Consgnr    PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Consgnr"`
	Consgn     PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Consgn,omitempty"`
	ShipTo     PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ShipTo,omitempty"`
	TrnsprtInf TransportDetails4       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TrnsprtInf"`
}

type TransportDetails4 struct {
	TrnsprtDocRef []DocumentIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TrnsprtDocRef"`
	TrnsprtdGoods []TransportedGoods1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 TrnsprtdGoods"`
	Consgnmt      Consignment3              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Consgnmt,omitempty"`
	RtgSummry     TransportMeans6           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 RtgSummry"`
	ShipmntDt     ShipmentDate1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 ShipmntDt"`
	FrghtChrgs    Charge25                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 FrghtChrgs,omitempty"`
	Incotrms      Incoterms4                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Incotrms,omitempty"`
}

type TransportMeans6 struct {
	IndvTrnsprt   SingleTransport8     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 IndvTrnsprt"`
	MltmdlTrnsprt MultimodalTransport3 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 MltmdlTrnsprt,omitempty"`
}

type TransportedGoods1 struct {
	PurchsOrdrRef DocumentIdentification7   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 PurchsOrdrRef"`
	GoodsDesc     Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 GoodsDesc,omitempty"`
	BuyrDfndInf   []UserDefinedInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 BuyrDfndInf,omitempty"`
	SellrDfndInf  []UserDefinedInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 SellrDfndInf,omitempty"`
}

type UnitOfMeasure3Choice struct {
	UnitOfMeasrCd   UnitOfMeasure4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 UnitOfMeasrCd"`
	OthrUnitOfMeasr Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 OthrUnitOfMeasr"`
}

// May be one of KGM, EA, LTN, MTR, INH, LY, GLI, GRM, CMT, MTK, FOT, 1A, INK, FTK, MIK, ONZ, PTI, PT, QTI, QT, GLL, MMT, KTM, YDK, MMK, CMK, KMK, MMQ, CLT, LTR, LBR, STN, BLL, BX, BO, CT, CH, CR, INQ, MTQ, OZI, OZA, BG, BL, TNE
type UnitOfMeasure4Code string

type UnitPrice18 struct {
	UnitPric UnitOfMeasure3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 UnitPric"`
	Amt      CurrencyAndAmount    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Amt"`
	Fctr     Max15NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Fctr,omitempty"`
}

type UserDefinedInformation1 struct {
	Labl Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Labl"`
	Inf  Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.05 Inf"`
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

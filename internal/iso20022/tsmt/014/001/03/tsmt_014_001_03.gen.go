// Code generated by main. DO NOT EDIT.

package tsmt_014_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification3Choice struct {
	IBAN      IBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 IBAN"`
	BBAN      BBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 BBAN"`
	UPIC      UPICIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 UPIC"`
	PrtryAcct SimpleIdentificationInformation2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PrtryAcct"`
}

type Adjustment4 struct {
	Tp             AdjustmentType2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Tp"`
	OthrAdjstmntTp Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 OthrAdjstmntTp"`
	Drctn          AdjustmentDirection1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Drctn"`
	Amt            CurrencyAndAmount        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Amt"`
}

// May be one of ADDD, SUBS
type AdjustmentDirection1Code string

// May be one of REBA, DISC, CREN, SURC
type AdjustmentType2Code string

type AirportDescription1 struct {
	Twn      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Twn"`
	AirprtNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 AirprtNm,omitempty"`
}

type AirportName1Choice struct {
	AirprtCd       Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 AirprtCd"`
	OthrAirprtDesc AirportDescription1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 OthrAirprtDesc"`
}

// Must match the pattern [a-zA-Z0-9]{1,30}
type BBANIdentifier string

type BICIdentification1 struct {
	BIC BICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 BIC"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

type CashAccount7 struct {
	Id  AccountIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Id"`
	Tp  CashAccountType2             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Tp,omitempty"`
	Ccy CurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Nm,omitempty"`
}

type CashAccountType2 struct {
	Cd    CashAccountType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

type CertificateDataSet1 struct {
	DataSetId         DocumentIdentification1         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 DataSetId"`
	CertTp            TradeCertificateType1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CertTp"`
	LineItm           []LineItemAndPOIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 LineItm,omitempty"`
	CertfdChrtcs      CertifiedCharacteristics1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CertfdChrtcs"`
	IsseDt            ISODate                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 IsseDt"`
	PlcOfIsse         PostalAddress5                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PlcOfIsse,omitempty"`
	Issr              PartyIdentification26           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Issr"`
	InspctnDt         DatePeriodDetails               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 InspctnDt,omitempty"`
	AuthrsdInspctrInd bool                            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 AuthrsdInspctrInd,omitempty"`
	CertId            Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CertId"`
	Trnsprt           SingleTransport3                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Trnsprt,omitempty"`
	GoodsDesc         Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 GoodsDesc,omitempty"`
	Consgnr           PartyIdentification26           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Consgnr,omitempty"`
	Consgn            PartyIdentification26           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Consgn,omitempty"`
	Manfctr           PartyIdentification26           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Manfctr,omitempty"`
	AddtlInf          []Max350Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 AddtlInf,omitempty"`
}

type CertifiedCharacteristics1Choice struct {
	Orgn             CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Orgn"`
	Qlty             Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Qlty"`
	Anlys            Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Anlys"`
	Wght             Quantity4   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Wght"`
	Qty              Quantity4   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Qty"`
	HlthIndctn       bool        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 HlthIndctn"`
	PhytosntryIndctn bool        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PhytosntryIndctn"`
}

type Charge13 struct {
	Tp    FreightCharges1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Tp"`
	Chrgs []ChargesDetails2   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Chrgs,omitempty"`
}

// May be one of SIGN, STDE, STOR, PACK, PICK, DNGR, SECU, INSU, COLF, CHOR, CHDE, AIRF, TRPT
type ChargeType8Code string

type ChargesDetails2 struct {
	Tp          ChargeType8Code   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Tp"`
	OthrChrgsTp Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 OthrChrgsTp"`
	Amt         CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Amt"`
}

type CommercialDataSet3 struct {
	DataSetId    DocumentIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 DataSetId"`
	ComrclDocRef InvoiceIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 ComrclDocRef"`
	Buyr         PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Buyr"`
	Sellr        PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Sellr"`
	BllTo        PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 BllTo,omitempty"`
	Goods        []LineItem9             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Goods"`
	PmtTerms     []PaymentTerms1         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PmtTerms"`
	SttlmTerms   SettlementTerms2        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 SttlmTerms"`
}

type Consignment1 struct {
	TtlQty  Quantity3 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TtlQty,omitempty"`
	TtlVol  Quantity3 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TtlVol,omitempty"`
	TtlWght Quantity3 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TtlWght,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyAndAmount struct {
	Value float64      `xml:",chardata"`
	Ccy   CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DataSetSubmissionReferences3 struct {
	TxId          Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TxId"`
	PurchsOrdrRef DocumentIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PurchsOrdrRef"`
	SubmitrTxRef  Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 SubmitrTxRef,omitempty"`
	ForcdMtch     bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 ForcdMtch"`
}

type DataSetSubmissionV03 struct {
	SubmissnId      MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 SubmissnId"`
	RltdTxRefs      []DataSetSubmissionReferences3  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 RltdTxRefs"`
	CmonSubmissnRef SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CmonSubmissnRef"`
	Instr           InstructionType3                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Instr"`
	BuyrBk          BICIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 BuyrBk"`
	SellrBk         BICIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 SellrBk"`
	ComrclDataSet   CommercialDataSet3              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 ComrclDataSet,omitempty"`
	TrnsprtDataSet  TransportDataSet3               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TrnsprtDataSet,omitempty"`
	InsrncDataSet   InsuranceDataSet1               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 InsrncDataSet,omitempty"`
	CertDataSet     []CertificateDataSet1           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CertDataSet,omitempty"`
	OthrCertDataSet []OtherCertificateDataSet1      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 OthrCertDataSet,omitempty"`
}

type DatePeriodDetails struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 ToDt"`
}

type Document struct {
	DataSetSubmissn DataSetSubmissionV03 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 DataSetSubmissn"`
}

type DocumentIdentification1 struct {
	Id      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Id"`
	Vrsn    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Vrsn"`
	Submitr BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Submitr"`
}

type DocumentIdentification7 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Id"`
	DtOfIsse ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 DtOfIsse"`
}

type FinancialInstitutionIdentification4Choice struct {
	BIC      BICIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 BIC"`
	NmAndAdr NameAndAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 NmAndAdr"`
}

// May be one of CLCT, PRPD
type FreightCharges1Code string

type GenericIdentification4 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Id"`
	IdTp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 IdTp"`
}

// Must match the pattern [a-zA-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBANIdentifier string

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

// May be one of EXW, FCA, FAS, FOB, CFR, CIF, CPT, CIP, DAF, DES, DEQ, DDU, DDP
type Incoterms1Code string

type Incoterms2 struct {
	Cd   Incoterms1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Cd"`
	Othr Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Othr"`
	Lctn Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Lctn"`
}

type InstructionType3 struct {
	Tp InstructionType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Tp"`
}

// May be one of MTCH, PMTC
type InstructionType3Code string

// May be one of ICCA, ICCB, ICCC, ICAI, IWCC, ISCC, IREC, ICLC, ISMC, CMCC, IRCE
type InsuranceClauses1Code string

type InsuranceDataSet1 struct {
	DataSetId      DocumentIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 DataSetId"`
	Issr           PartyIdentification26       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Issr"`
	IsseDt         ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 IsseDt"`
	FctvDt         ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 FctvDt,omitempty"`
	PlcOfIsse      PostalAddress5              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PlcOfIsse,omitempty"`
	InsrncDocId    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 InsrncDocId"`
	Trnsprt        SingleTransport3            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Trnsprt,omitempty"`
	InsrdAmt       CurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 InsrdAmt"`
	InsrdGoodsDesc Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 InsrdGoodsDesc,omitempty"`
	InsrncConds    []Max350Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 InsrncConds,omitempty"`
	InsrncClauses  []InsuranceClauses1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 InsrncClauses,omitempty"`
	Assrd          PartyIdentification29Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Assrd"`
	ClmsPyblAt     PostalAddress5              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 ClmsPyblAt"`
	ClmsPyblIn     CurrencyCode                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 ClmsPyblIn,omitempty"`
}

type InvoiceIdentification1 struct {
	InvcNb Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 InvcNb"`
	IsseDt ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 IsseDt"`
}

type LineItem9 struct {
	PurchsOrdrRef  DocumentIdentification7   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PurchsOrdrRef"`
	FnlSubmissn    bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 FnlSubmissn"`
	ComrclLineItms []LineItemDetails9        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 ComrclLineItms"`
	LineItmsTtlAmt CurrencyAndAmount         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 LineItmsTtlAmt"`
	Incotrms       Incoterms2                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Incotrms,omitempty"`
	Adjstmnt       []Adjustment4             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Adjstmnt,omitempty"`
	FrghtChrgs     Charge13                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 FrghtChrgs,omitempty"`
	Tax            []Tax12                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Tax,omitempty"`
	TtlNetAmt      CurrencyAndAmount         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TtlNetAmt"`
	BuyrDfndInf    []UserDefinedInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 BuyrDfndInf,omitempty"`
	SellrDfndInf   []UserDefinedInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 SellrDfndInf,omitempty"`
}

type LineItemAndPOIdentification1 struct {
	LineItmId     []Max70Text             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 LineItmId"`
	PurchsOrdrRef DocumentIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PurchsOrdrRef"`
}

type LineItemDetails9 struct {
	LineItmId  Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 LineItmId"`
	Qty        Quantity4                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Qty"`
	UnitPric   UnitPrice9                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 UnitPric,omitempty"`
	PdctNm     Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PdctNm,omitempty"`
	PdctIdr    []ProductIdentifier2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PdctIdr,omitempty"`
	PdctChrtcs []ProductCharacteristics1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PdctChrtcs,omitempty"`
	PdctCtgy   []ProductCategory1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PdctCtgy,omitempty"`
	PdctOrgn   CountryCode                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PdctOrgn,omitempty"`
	Adjstmnt   []Adjustment4                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Adjstmnt,omitempty"`
	FrghtChrgs Charge13                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 FrghtChrgs,omitempty"`
	Tax        []Tax12                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Tax,omitempty"`
	TtlAmt     CurrencyAndAmount               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TtlAmt"`
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

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CreDtTm"`
}

type MultimodalTransport3 struct {
	TakngInChrg  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TakngInChrg"`
	PlcOfFnlDstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PlcOfFnlDstn"`
}

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Adr"`
}

type OtherCertificateDataSet1 struct {
	DataSetId DocumentIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 DataSetId"`
	CertId    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CertId"`
	CertTp    TradeCertificateType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CertTp"`
	IsseDt    ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 IsseDt"`
	Issr      PartyIdentification26     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Issr"`
	CertInf   []Max350Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CertInf,omitempty"`
}

type PartyIdentification26 struct {
	Nm      Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Nm"`
	PrtryId GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PrtryId,omitempty"`
	PstlAdr PostalAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PstlAdr"`
}

type PartyIdentification29Choice struct {
	BIC      BICIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 BIC"`
	NmAndAdr PartyIdentification26 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 NmAndAdr"`
}

type PaymentPeriod1 struct {
	Cd       PaymentTime1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Cd"`
	NbOfDays float64          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 NbOfDays,omitempty"`
}

type PaymentTerms1 struct {
	OthrPmtTerms Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 OthrPmtTerms"`
	PmtCd        PaymentPeriod1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PmtCd"`
	Pctg         float64           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Pctg"`
	Amt          CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Amt"`
}

// May be one of CASH, EMTD, EPRD, PRMD, IREC, PRMR, EPRR, EMTR
type PaymentTime1Code string

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Ctry"`
}

type PostalAddress5 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PstCdId,omitempty"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Ctry"`
}

type ProductCategory1 struct {
	Tp   ProductCategory1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Tp"`
	Ctgy Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Ctgy"`
}

type ProductCategory1Choice struct {
	StrdPdctCtgy ProductCategory1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 StrdPdctCtgy"`
	OthrPdctCtgy GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 OthrPdctCtgy"`
}

// May be one of HRTR, QOTA, PRGP, LOBU, GNDR
type ProductCategory1Code string

type ProductCharacteristics1 struct {
	Tp     ProductCharacteristics1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Tp"`
	Chrtcs Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Chrtcs"`
}

type ProductCharacteristics1Choice struct {
	StrdPdctChrtcs ProductCharacteristics1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 StrdPdctChrtcs"`
	OthrPdctChrtcs GenericIdentification4  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 OthrPdctChrtcs"`
}

// May be one of BISP, CHNR, CLOR, EDSP, ENNR, OPTN, ORCR, PCTV, SISP, SIZE, SZRG, SPRM, STOR, VINR
type ProductCharacteristics1Code string

type ProductIdentifier2 struct {
	Tp  ProductIdentifier2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Tp"`
	Idr Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Idr"`
}

type ProductIdentifier2Choice struct {
	StrdPdctIdr ProductIdentifier2     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 StrdPdctIdr"`
	OthrPdctIdr GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 OthrPdctIdr"`
}

// May be one of BINR, COMD, EANC, HRTR, MANI, MODL, PART, QOTA, STYL, SUPI, UPCC
type ProductIdentifier2Code string

type Quantity3 struct {
	UnitOfMeasrCd   UnitOfMeasure4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 UnitOfMeasrCd"`
	OthrUnitOfMeasr Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 OthrUnitOfMeasr"`
	Val             float64            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Val"`
}

type Quantity4 struct {
	UnitOfMeasrCd   UnitOfMeasure4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 UnitOfMeasrCd"`
	OthrUnitOfMeasr Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 OthrUnitOfMeasr"`
	Val             float64            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Val"`
	Fctr            Max15NumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Fctr,omitempty"`
}

type SettlementTerms2 struct {
	CdtrAgt  FinancialInstitutionIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CdtrAgt,omitempty"`
	CdtrAcct CashAccount7                              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 CdtrAcct"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Id"`
}

type SimpleIdentificationInformation2 struct {
	Id Max34Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Id"`
}

type SingleTransport3 struct {
	TrnsprtByAir  TransportByAir2  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TrnsprtByAir,omitempty"`
	TrnsprtBySea  TransportBySea4  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TrnsprtBySea,omitempty"`
	TrnsprtByRoad TransportByRoad2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TrnsprtByRoad,omitempty"`
	TrnsprtByRail TransportByRail2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TrnsprtByRail,omitempty"`
}

type SingleTransport5 struct {
	TrnsprtByAir  []TransportByAir2  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TrnsprtByAir,omitempty"`
	TrnsprtBySea  []TransportBySea4  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TrnsprtBySea,omitempty"`
	TrnsprtByRoad []TransportByRoad2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TrnsprtByRoad,omitempty"`
	TrnsprtByRail []TransportByRail2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TrnsprtByRail,omitempty"`
}

type Tax12 struct {
	Tp        TaxType9Code      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Tp"`
	OthrTaxTp Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 OthrTaxTp"`
	Amt       CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Amt"`
}

// May be one of PROV, NATI, STAT, WITH, STAM, COAX, VATA, CUST
type TaxType9Code string

// May be one of ANLY, QUAL, QUAN, WEIG, ORIG, HEAL, PHYT
type TradeCertificateType1Code string

// May be one of BENE, SHIP, UND1, UND2
type TradeCertificateType2Code string

type TransportByAir2 struct {
	DprtureAirprt AirportName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 DprtureAirprt"`
	DstnAirprt    AirportName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 DstnAirprt"`
	AirCrrierNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 AirCrrierNm,omitempty"`
}

type TransportByRail2 struct {
	PlcOfRct     Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PlcOfRct"`
	PlcOfDlvry   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PlcOfDlvry"`
	RailCrrierNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 RailCrrierNm,omitempty"`
}

type TransportByRoad2 struct {
	PlcOfRct     Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PlcOfRct"`
	PlcOfDlvry   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PlcOfDlvry"`
	RoadCrrierNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 RoadCrrierNm,omitempty"`
}

type TransportBySea4 struct {
	PortOfLoadng  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PortOfLoadng"`
	PortOfDschrge Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PortOfDschrge"`
	VsslNm        Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 VsslNm,omitempty"`
	SeaCrrierNm   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 SeaCrrierNm,omitempty"`
}

type TransportDataSet3 struct {
	DataSetId  DocumentIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 DataSetId"`
	Buyr       PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Buyr,omitempty"`
	Sellr      PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Sellr,omitempty"`
	Consgnr    PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Consgnr"`
	Consgn     PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Consgn,omitempty"`
	ShipTo     PartyIdentification26   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 ShipTo,omitempty"`
	TrnsprtInf TransportDetails2       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TrnsprtInf"`
}

type TransportDetails2 struct {
	TrnsprtDocRef   []DocumentIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TrnsprtDocRef"`
	TrnsprtdGoods   []TransportedGoods1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 TrnsprtdGoods"`
	Consgnmt        Consignment1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Consgnmt,omitempty"`
	RtgSummry       TransportMeans2           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 RtgSummry"`
	PropsdShipmntDt ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PropsdShipmntDt"`
	ActlShipmntDt   ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 ActlShipmntDt"`
	Incotrms        Incoterms2                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Incotrms,omitempty"`
	FrghtChrgs      Charge13                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 FrghtChrgs,omitempty"`
}

type TransportMeans2 struct {
	IndvTrnsprt   SingleTransport5     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 IndvTrnsprt"`
	MltmdlTrnsprt MultimodalTransport3 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 MltmdlTrnsprt,omitempty"`
}

type TransportedGoods1 struct {
	PurchsOrdrRef DocumentIdentification7   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 PurchsOrdrRef"`
	GoodsDesc     Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 GoodsDesc,omitempty"`
	BuyrDfndInf   []UserDefinedInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 BuyrDfndInf,omitempty"`
	SellrDfndInf  []UserDefinedInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 SellrDfndInf,omitempty"`
}

// Must match the pattern [0-9]{8,17}
type UPICIdentifier string

// May be one of KGM, EA, LTN, MTR, INH, LY, GLI, GRM, CMT, MTK, FOT, 1A, INK, FTK, MIK, ONZ, PTI, PT, QTI, QT, GLL, MMT, KTM, YDK, MMK, CMK, KMK, MMQ, CLT, LTR, LBR, STN, BLL, BX, BO, CT, CH, CR, INQ, MTQ, OZI, OZA, BG, BL, TNE
type UnitOfMeasure4Code string

type UnitPrice9 struct {
	UnitOfMeasrCd   UnitOfMeasure4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 UnitOfMeasrCd"`
	OthrUnitOfMeasr Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 OthrUnitOfMeasr"`
	Amt             CurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Amt"`
	Fctr            Max15NumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Fctr,omitempty"`
}

type UserDefinedInformation1 struct {
	Labl Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Labl"`
	Inf  Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.03 Inf"`
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

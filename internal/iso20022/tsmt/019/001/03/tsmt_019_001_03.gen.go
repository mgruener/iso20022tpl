// Code generated by main. DO NOT EDIT.

package tsmt_019_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification3Choice struct {
	IBAN      IBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 IBAN"`
	BBAN      BBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 BBAN"`
	UPIC      UPICIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 UPIC"`
	PrtryAcct SimpleIdentificationInformation2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PrtryAcct"`
}

type Adjustment3 struct {
	Tp             AdjustmentType2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Tp"`
	OthrAdjstmntTp Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrAdjstmntTp"`
	Amt            CurrencyAndAmount        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Amt"`
	Rate           float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Rate"`
	Drctn          AdjustmentDirection1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Drctn"`
}

// May be one of ADDD, SUBS
type AdjustmentDirection1Code string

// May be one of REBA, DISC, CREN, SURC
type AdjustmentType2Code string

type AirportDescription1 struct {
	Twn      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Twn"`
	AirprtNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 AirprtNm,omitempty"`
}

type AirportName1Choice struct {
	AirprtCd       Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 AirprtCd"`
	OthrAirprtDesc AirportDescription1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrAirprtDesc"`
}

// May be one of BUYE, SELL, BUBA, SEBA
type AssuredType1Code string

// Must match the pattern [a-zA-Z0-9]{1,30}
type BBANIdentifier string

type BICIdentification1 struct {
	BIC BICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 BIC"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

type Baseline3 struct {
	SubmitrBaselnId      DocumentIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SubmitrBaselnId"`
	SvcCd                TradeFinanceService2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SvcCd"`
	PurchsOrdrRef        DocumentIdentification7  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PurchsOrdrRef"`
	Buyr                 PartyIdentification26    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Buyr"`
	Sellr                PartyIdentification26    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Sellr"`
	BuyrBk               BICIdentification1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 BuyrBk"`
	SellrBk              BICIdentification1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SellrBk"`
	BuyrSdSubmitgBk      []BICIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 BuyrSdSubmitgBk,omitempty"`
	SellrSdSubmitgBk     []BICIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SellrSdSubmitgBk,omitempty"`
	BllTo                PartyIdentification26    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 BllTo,omitempty"`
	ShipTo               PartyIdentification26    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 ShipTo,omitempty"`
	Consgn               PartyIdentification26    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Consgn,omitempty"`
	Goods                LineItem7                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Goods"`
	PmtTerms             []PaymentTerms1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PmtTerms"`
	SttlmTerms           SettlementTerms2         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SttlmTerms,omitempty"`
	PmtOblgtn            []PaymentObligation1     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PmtOblgtn,omitempty"`
	LatstMtchDt          ISODate                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 LatstMtchDt,omitempty"`
	ComrclDataSetReqrd   RequiredSubmission2      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 ComrclDataSetReqrd"`
	TrnsprtDataSetReqrd  RequiredSubmission2      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 TrnsprtDataSetReqrd,omitempty"`
	InsrncDataSetReqrd   RequiredSubmission3      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 InsrncDataSetReqrd,omitempty"`
	CertDataSetReqrd     []RequiredSubmission4    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 CertDataSetReqrd,omitempty"`
	OthrCertDataSetReqrd []RequiredSubmission5    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrCertDataSetReqrd,omitempty"`
	InttToPayXpctd       bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 InttToPayXpctd"`
}

type CashAccount7 struct {
	Id  AccountIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Id"`
	Tp  CashAccountType2             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Tp,omitempty"`
	Ccy CurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Nm,omitempty"`
}

type CashAccountType2 struct {
	Cd    CashAccountType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

type Charge12 struct {
	Tp    FreightCharges1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Tp"`
	Chrgs []ChargesDetails1   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Chrgs,omitempty"`
}

// May be one of SIGN, STDE, STOR, PACK, PICK, DNGR, SECU, INSU, COLF, CHOR, CHDE, AIRF, TRPT
type ChargeType8Code string

type ChargesDetails1 struct {
	Tp          ChargeType8Code   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Tp"`
	OthrChrgsTp Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrChrgsTp"`
	Amt         CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Amt"`
	Rate        float64           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Rate"`
}

type ContactIdentification1 struct {
	Nm       Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Nm"`
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 GvnNm,omitempty"`
	Role     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Role,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PhneNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 EmailAdr,omitempty"`
}

type ContactIdentification3 struct {
	BIC      BICIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 BIC"`
	Nm       Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Nm"`
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 GvnNm,omitempty"`
	Role     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Role,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PhneNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 EmailAdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyAndAmount struct {
	Value float64      `xml:",chardata"`
	Ccy   CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type Document struct {
	InitlBaselnSubmissn InitialBaselineSubmissionV03 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 InitlBaselnSubmissn"`
}

type DocumentIdentification1 struct {
	Id      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Id"`
	Vrsn    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Vrsn"`
	Submitr BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Submitr"`
}

type DocumentIdentification7 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Id"`
	DtOfIsse ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 DtOfIsse"`
}

type FinancialInstitutionIdentification4Choice struct {
	BIC      BICIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 BIC"`
	NmAndAdr NameAndAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 NmAndAdr"`
}

// May be one of CLCT, PRPD
type FreightCharges1Code string

type GenericIdentification4 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Id"`
	IdTp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 IdTp"`
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

type Incoterms1 struct {
	Cd   Incoterms1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Cd"`
	Othr Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Othr"`
	Lctn Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Lctn,omitempty"`
}

// May be one of EXW, FCA, FAS, FOB, CFR, CIF, CPT, CIP, DAF, DES, DEQ, DDU, DDP
type Incoterms1Code string

type InitialBaselineSubmissionV03 struct {
	SubmissnId      MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SubmissnId"`
	SubmitrTxRef    SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SubmitrTxRef"`
	Instr           InstructionType1                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Instr"`
	Baseln          Baseline3                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Baseln"`
	BuyrCtctPrsn    []ContactIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 BuyrCtctPrsn,omitempty"`
	SellrCtctPrsn   []ContactIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SellrCtctPrsn,omitempty"`
	BuyrBkCtctPrsn  []ContactIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 BuyrBkCtctPrsn"`
	SellrBkCtctPrsn []ContactIdentification1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SellrBkCtctPrsn"`
	OthrBkCtctPrsn  []ContactIdentification3        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrBkCtctPrsn,omitempty"`
}

type InstructionType1 struct {
	Tp InstructionType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Tp"`
}

// May be one of LODG, FPTR
type InstructionType1Code string

// May be one of ICCA, ICCB, ICCC, ICAI, IWCC, ISCC, IREC, ICLC, ISMC, CMCC, IRCE
type InsuranceClauses1Code string

type LineItem7 struct {
	GoodsDesc      Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 GoodsDesc,omitempty"`
	PrtlShipmnt    bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PrtlShipmnt"`
	TrnsShipmnt    bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 TrnsShipmnt,omitempty"`
	ShipmntDtRg    ShipmentDateRange1        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 ShipmntDtRg,omitempty"`
	LineItmDtls    []LineItemDetails7        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 LineItmDtls"`
	LineItmsTtlAmt CurrencyAndAmount         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 LineItmsTtlAmt"`
	RtgSummry      TransportMeans1           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 RtgSummry,omitempty"`
	Incotrms       []Incoterms1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Incotrms,omitempty"`
	Adjstmnt       []Adjustment3             `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Adjstmnt,omitempty"`
	FrghtChrgs     Charge12                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 FrghtChrgs,omitempty"`
	Tax            []Tax13                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Tax,omitempty"`
	TtlNetAmt      CurrencyAndAmount         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 TtlNetAmt"`
	BuyrDfndInf    []UserDefinedInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 BuyrDfndInf,omitempty"`
	SellrDfndInf   []UserDefinedInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SellrDfndInf,omitempty"`
}

type LineItemDetails7 struct {
	LineItmId    Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 LineItmId"`
	Qty          Quantity4                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Qty"`
	QtyTlrnce    PercentageTolerance1            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 QtyTlrnce,omitempty"`
	UnitPric     UnitPrice9                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 UnitPric,omitempty"`
	PricTlrnce   PercentageTolerance1            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PricTlrnce,omitempty"`
	PdctNm       Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PdctNm,omitempty"`
	PdctIdr      []ProductIdentifier2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PdctIdr,omitempty"`
	PdctChrtcs   []ProductCharacteristics1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PdctChrtcs,omitempty"`
	PdctCtgy     []ProductCategory1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PdctCtgy,omitempty"`
	PdctOrgn     []CountryCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PdctOrgn,omitempty"`
	ShipmntSchdl ShipmentSchedule1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 ShipmntSchdl,omitempty"`
	RtgSummry    TransportMeans1                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 RtgSummry,omitempty"`
	Incotrms     []Incoterms1                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Incotrms,omitempty"`
	Adjstmnt     []Adjustment3                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Adjstmnt,omitempty"`
	FrghtChrgs   Charge12                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 FrghtChrgs,omitempty"`
	Tax          []Tax13                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Tax,omitempty"`
	TtlAmt       CurrencyAndAmount               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 TtlAmt"`
}

// Must be at least 1 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 CreDtTm"`
}

type MultimodalTransport3 struct {
	TakngInChrg  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 TakngInChrg"`
	PlcOfFnlDstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PlcOfFnlDstn"`
}

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Adr"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type PartyIdentification26 struct {
	Nm      Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Nm"`
	PrtryId GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PrtryId,omitempty"`
	PstlAdr PostalAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PstlAdr"`
}

type PartyIdentification27 struct {
	Nm      Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Nm"`
	PrtryId GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PrtryId,omitempty"`
	Ctry    CountryCode            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Ctry"`
}

type PaymentObligation1 struct {
	OblgrBk    BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OblgrBk"`
	RcptBk     BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 RcptBk"`
	Amt        CurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Amt"`
	Pctg       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Pctg"`
	ChrgsAmt   CurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 ChrgsAmt,omitempty"`
	ChrgsPctg  float64            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 ChrgsPctg,omitempty"`
	XpryDt     ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 XpryDt"`
	AplblLaw   CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 AplblLaw,omitempty"`
	PmtTerms   []PaymentTerms2    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PmtTerms,omitempty"`
	SttlmTerms SettlementTerms2   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SttlmTerms,omitempty"`
}

type PaymentPeriod1 struct {
	Cd       PaymentTime1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Cd"`
	NbOfDays float64          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 NbOfDays,omitempty"`
}

type PaymentPeriod2 struct {
	Cd       PaymentTime2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Cd"`
	NbOfDays float64          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 NbOfDays,omitempty"`
}

type PaymentTerms1 struct {
	OthrPmtTerms Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrPmtTerms"`
	PmtCd        PaymentPeriod1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PmtCd"`
	Pctg         float64           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Pctg"`
	Amt          CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Amt"`
}

type PaymentTerms2 struct {
	OthrPmtTerms Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrPmtTerms"`
	PmtCd        PaymentPeriod2    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PmtCd"`
	Pctg         float64           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Pctg"`
	Amt          CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Amt"`
}

// May be one of CASH, EMTD, EPRD, PRMD, IREC, PRMR, EPRR, EMTR
type PaymentTime1Code string

// May be one of CASH, EMTD, EPRD, PRMD, IREC, PRMR, EPRR, EMTR, EPAM
type PaymentTime2Code string

type PercentageTolerance1 struct {
	PlusPct float64 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PlusPct"`
	MnsPct  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 MnsPct"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Ctry"`
}

type PostalAddress5 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PstCdId,omitempty"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Ctry"`
}

type ProductCategory1 struct {
	Tp   ProductCategory1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Tp"`
	Ctgy Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Ctgy"`
}

type ProductCategory1Choice struct {
	StrdPdctCtgy ProductCategory1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 StrdPdctCtgy"`
	OthrPdctCtgy GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrPdctCtgy"`
}

// May be one of HRTR, QOTA, PRGP, LOBU, GNDR
type ProductCategory1Code string

type ProductCharacteristics1 struct {
	Tp     ProductCharacteristics1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Tp"`
	Chrtcs Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Chrtcs"`
}

type ProductCharacteristics1Choice struct {
	StrdPdctChrtcs ProductCharacteristics1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 StrdPdctChrtcs"`
	OthrPdctChrtcs GenericIdentification4  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrPdctChrtcs"`
}

// May be one of BISP, CHNR, CLOR, EDSP, ENNR, OPTN, ORCR, PCTV, SISP, SIZE, SZRG, SPRM, STOR, VINR
type ProductCharacteristics1Code string

type ProductIdentifier2 struct {
	Tp  ProductIdentifier2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Tp"`
	Idr Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Idr"`
}

type ProductIdentifier2Choice struct {
	StrdPdctIdr ProductIdentifier2     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 StrdPdctIdr"`
	OthrPdctIdr GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrPdctIdr"`
}

// May be one of BINR, COMD, EANC, HRTR, MANI, MODL, PART, QOTA, STYL, SUPI, UPCC
type ProductIdentifier2Code string

type Quantity4 struct {
	UnitOfMeasrCd   UnitOfMeasure4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 UnitOfMeasrCd"`
	OthrUnitOfMeasr Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrUnitOfMeasr"`
	Val             float64            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Val"`
	Fctr            Max15NumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Fctr,omitempty"`
}

type RequiredSubmission2 struct {
	Submitr []BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Submitr"`
}

type RequiredSubmission3 struct {
	Submitr      []BICIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Submitr"`
	MtchIssr     PartyIdentification27   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 MtchIssr,omitempty"`
	MtchIsseDt   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 MtchIsseDt"`
	MtchTrnsprt  bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 MtchTrnsprt"`
	MtchAmt      bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 MtchAmt"`
	ClausesReqrd []InsuranceClauses1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 ClausesReqrd,omitempty"`
	MtchAssrdPty AssuredType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 MtchAssrdPty,omitempty"`
}

type RequiredSubmission4 struct {
	Submitr           []BICIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Submitr"`
	CertTp            TradeCertificateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 CertTp"`
	MtchIssr          PartyIdentification27     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 MtchIssr,omitempty"`
	MtchIsseDt        bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 MtchIsseDt"`
	MtchInspctnDt     bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 MtchInspctnDt"`
	AuthrsdInspctrInd bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 AuthrsdInspctrInd"`
	MtchConsgn        bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 MtchConsgn"`
	MtchManfctr       PartyIdentification27     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 MtchManfctr,omitempty"`
	LineItmId         []Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 LineItmId,omitempty"`
}

type RequiredSubmission5 struct {
	Submitr []BICIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Submitr"`
	CertTp  TradeCertificateType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 CertTp"`
}

type SettlementTerms2 struct {
	CdtrAgt  FinancialInstitutionIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 CdtrAgt,omitempty"`
	CdtrAcct CashAccount7                              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 CdtrAcct"`
}

type ShipmentDateRange1 struct {
	EarlstShipmntDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 EarlstShipmntDt,omitempty"`
	LatstShipmntDt  ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 LatstShipmntDt,omitempty"`
}

type ShipmentDateRange2 struct {
	SubQtyVal       float64 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SubQtyVal"`
	EarlstShipmntDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 EarlstShipmntDt,omitempty"`
	LatstShipmntDt  ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 LatstShipmntDt,omitempty"`
}

type ShipmentSchedule1Choice struct {
	ShipmntDtRg     ShipmentDateRange1   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 ShipmntDtRg"`
	ShipmntSubSchdl []ShipmentDateRange2 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 ShipmntSubSchdl"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Id"`
}

type SimpleIdentificationInformation2 struct {
	Id Max34Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Id"`
}

type SingleTransport4 struct {
	TrnsprtByAir  []TransportByAir3  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 TrnsprtByAir,omitempty"`
	TrnsprtBySea  []TransportBySea3  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 TrnsprtBySea,omitempty"`
	TrnsprtByRoad []TransportByRoad3 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 TrnsprtByRoad,omitempty"`
	TrnsprtByRail []TransportByRail3 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 TrnsprtByRail,omitempty"`
}

type Tax13 struct {
	Tp        TaxType9Code      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Tp"`
	OthrTaxTp Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrTaxTp"`
	Amt       CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Amt"`
	Rate      float64           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Rate"`
}

// May be one of PROV, NATI, STAT, WITH, STAM, COAX, VATA, CUST
type TaxType9Code string

// May be one of ANLY, QUAL, QUAN, WEIG, ORIG, HEAL, PHYT
type TradeCertificateType1Code string

// May be one of BENE, SHIP, UND1, UND2
type TradeCertificateType2Code string

// May be one of LEV1, LEV2, LEV3
type TradeFinanceService2Code string

type TransportByAir3 struct {
	DprtureAirprt []AirportName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 DprtureAirprt,omitempty"`
	DstnAirprt    []AirportName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 DstnAirprt"`
	AirCrrierNm   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 AirCrrierNm,omitempty"`
}

type TransportByRail3 struct {
	PlcOfRct     []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PlcOfRct,omitempty"`
	PlcOfDlvry   []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PlcOfDlvry"`
	RailCrrierNm Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 RailCrrierNm,omitempty"`
}

type TransportByRoad3 struct {
	PlcOfRct     []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PlcOfRct,omitempty"`
	PlcOfDlvry   []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PlcOfDlvry"`
	RoadCrrierNm Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 RoadCrrierNm,omitempty"`
}

type TransportBySea3 struct {
	PortOfLoadng  []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PortOfLoadng,omitempty"`
	PortOfDschrge []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 PortOfDschrge"`
	SeaCrrierNm   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 SeaCrrierNm,omitempty"`
}

type TransportMeans1 struct {
	IndvTrnsprt   SingleTransport4     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 IndvTrnsprt"`
	MltmdlTrnsprt MultimodalTransport3 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 MltmdlTrnsprt,omitempty"`
}

// Must match the pattern [0-9]{8,17}
type UPICIdentifier string

// May be one of KGM, EA, LTN, MTR, INH, LY, GLI, GRM, CMT, MTK, FOT, 1A, INK, FTK, MIK, ONZ, PTI, PT, QTI, QT, GLL, MMT, KTM, YDK, MMK, CMK, KMK, MMQ, CLT, LTR, LBR, STN, BLL, BX, BO, CT, CH, CR, INQ, MTQ, OZI, OZA, BG, BL, TNE
type UnitOfMeasure4Code string

type UnitPrice9 struct {
	UnitOfMeasrCd   UnitOfMeasure4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 UnitOfMeasrCd"`
	OthrUnitOfMeasr Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 OthrUnitOfMeasr"`
	Amt             CurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Amt"`
	Fctr            Max15NumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Fctr,omitempty"`
}

type UserDefinedInformation1 struct {
	Labl Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Labl"`
	Inf  Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.03 Inf"`
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

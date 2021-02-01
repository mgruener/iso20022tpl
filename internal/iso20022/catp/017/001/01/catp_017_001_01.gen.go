// Code generated by main. DO NOT EDIT.

package catp_017_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// May be one of ABAL, ASTS, CFGT, CCNT, DISC, SNDM, RPTC
type ATMCommand4Code string

type ATMCommand7 struct {
	Tp        ATMCommand4Code             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Tp"`
	Urgcy     TMSContactLevel2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Urgcy"`
	DtTm      ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 DtTm,omitempty"`
	CmdId     ATMCommandIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CmdId,omitempty"`
	CmdParams ATMCommandParameters1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CmdParams,omitempty"`
}

type ATMCommandIdentification1 struct {
	Orgn Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Orgn,omitempty"`
	Ref  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Ref,omitempty"`
	Prcr Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Prcr,omitempty"`
}

type ATMCommandParameters1Choice struct {
	ATMReqrdGblSts  ATMStatus1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ATMReqrdGblSts"`
	XpctdMsgFctn    MessageFunction8Code       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 XpctdMsgFctn"`
	ReqrdCfgtnParam ATMConfigurationParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ReqrdCfgtnParam"`
}

type ATMConfigurationParameter1 struct {
	Tp   DataSetCategory7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Tp"`
	Vrsn Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Vrsn"`
}

type ATMContext19 struct {
	SsnRef Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 SsnRef,omitempty"`
	Svc    ATMService23 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Svc"`
}

type ATMCustomer5 struct {
	Prfl         ATMCustomerProfile2              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Prfl,omitempty"`
	AuthntcnRslt []TransactionVerificationResult5 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AuthntcnRslt,omitempty"`
}

type ATMCustomerProfile2 struct {
	PrflRef Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PrflRef,omitempty"`
	CstmrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CstmrId,omitempty"`
}

// May be one of CDIS, DPRN, JRNL, JPRN, RPRN, RWDR
type ATMDevice1Code string

type ATMEnvironment12 struct {
	Acqrr          Acquirer7                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Acqrr,omitempty"`
	ATMMgr         Acquirer8                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ATMMgr,omitempty"`
	HstgNtty       TerminalHosting1         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 HstgNtty,omitempty"`
	ATM            AutomatedTellerMachine2  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ATM"`
	Cstmr          ATMCustomer5             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Cstmr"`
	PrtctdCardData ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PrtctdCardData,omitempty"`
	PlainCardData  PlainCardData19          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PlainCardData,omitempty"`
}

type ATMMessageFunction2 struct {
	Fctn     MessageFunction11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Fctn"`
	ATMSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ATMSvcCd,omitempty"`
	HstSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 HstSvcCd,omitempty"`
}

type ATMService23 struct {
	SvcRef     Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 SvcRef,omitempty"`
	ATMSvcCd   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ATMSvcCd,omitempty"`
	HstSvcCd   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 HstSvcCd,omitempty"`
	SvcTp      ATMServiceType9Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 SvcTp"`
	SvcVarntId []Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 SvcVarntId,omitempty"`
}

// May be one of TRFC, TRFI, TRFP
type ATMServiceType9Code string

// May be one of INSV, OUTS
type ATMStatus1Code string

type ATMTransaction24 struct {
	TxId          TransactionIdentifier1   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 TxId"`
	RcncltnId     Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 RcncltnId,omitempty"`
	CdtrLabl      Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CdtrLabl,omitempty"`
	DbtrLabl      Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 DbtrLabl,omitempty"`
	TrfIdr        Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 TrfIdr,omitempty"`
	PmtRef        Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PmtRef,omitempty"`
	TxRspn        ResponseType7            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 TxRspn"`
	Actn          []Action7                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Actn,omitempty"`
	AcctFr        CardAccount13            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AcctFr,omitempty"`
	PrtctdAcctFr  ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PrtctdAcctFr,omitempty"`
	AcctTo        []CardAccount13          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AcctTo,omitempty"`
	PrtctdAcctTo  ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PrtctdAcctTo,omitempty"`
	TtlAuthrsdAmt AmountAndCurrency1       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 TtlAuthrsdAmt"`
	TtlReqdAmt    float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 TtlReqdAmt,omitempty"`
	DtldReqdAmt   DetailedAmount17         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 DtldReqdAmt,omitempty"`
	AddtlChrg     []DetailedAmount18       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AddtlChrg,omitempty"`
	Lmts          ATMTransactionAmounts6   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Lmts,omitempty"`
	ReqdExctnDt   ISODate                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ReqdExctnDt,omitempty"`
	PropsdExctnDt ISODate                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PropsdExctnDt,omitempty"`
	InstntTrfPrgm Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 InstntTrfPrgm,omitempty"`
	RcrngTrf      RecurringTransaction3    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 RcrngTrf,omitempty"`
	AuthstnRslt   AuthorisationResult13    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AuthstnRslt,omitempty"`
	ICCRltdData   Max10000Binary           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ICCRltdData,omitempty"`
	Cmd           []ATMCommand7            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Cmd,omitempty"`
}

type ATMTransactionAmounts6 struct {
	Ccy         ActiveCurrencyCode       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Ccy,omitempty"`
	MaxPssblAmt float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 MaxPssblAmt,omitempty"`
	MinPssblAmt float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 MinPssblAmt,omitempty"`
	AddtlAmt    []ATMTransactionAmounts7 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AddtlAmt,omitempty"`
}

type ATMTransactionAmounts7 struct {
	Tp   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Tp"`
	Amt  float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Amt"`
	Ccy  ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Ccy,omitempty"`
	Labl Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Labl,omitempty"`
}

type ATMTransferResponse1 struct {
	Envt  ATMEnvironment12 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Envt"`
	Cntxt ATMContext19     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Cntxt"`
	Tx    ATMTransaction24 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Tx"`
}

type ATMTransferResponseV01 struct {
	Hdr              Header31                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Hdr"`
	PrtctdATMTrfRspn ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PrtctdATMTrfRspn,omitempty"`
	ATMTrfRspn       ATMTransferResponse1     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ATMTrfRspn,omitempty"`
	SctyTrlr         ContentInformationType15 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 SctyTrlr,omitempty"`
}

type AccountIdentification31Choice struct {
	IBAN     IBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 IBAN"`
	BBAN     BBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 BBAN"`
	UPIC     UPICIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 UPIC"`
	DmstAcct SimpleIdentificationInformation4 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 DmstAcct"`
}

type Acquirer7 struct {
	AcqrgInstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AcqrgInstn,omitempty"`
	Brnch      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Brnch,omitempty"`
}

type Acquirer8 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Id"`
	ApplVrsn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ApplVrsn,omitempty"`
}

type Action7 struct {
	ActnTp     ActionType6Code       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ActnTp"`
	MsgToPres  ActionMessage4        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 MsgToPres,omitempty"`
	ReqToPrfrm MessageFunction11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ReqToPrfrm,omitempty"`
}

type ActionMessage4 struct {
	Frmt         OutputFormat2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Frmt,omitempty"`
	Msg          Max20000Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Msg,omitempty"`
	Ref          Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Ref,omitempty"`
	Dvc          ATMDevice1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Dvc,omitempty"`
	MsgCnttSgntr Max35Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 MsgCnttSgntr,omitempty"`
}

// May be one of DCCQ, FEES, HAMT, LAMT, BUSY, CPTR, DISP, CPNS, RQST, PINL, PINR, TRCK
type ActionType6Code string

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5
type Algorithm12Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of EA2C, E3DC, EA9C, EA5C
type Algorithm15Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Param,omitempty"`
}

type AmountAndCurrency1 struct {
	Amt float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Amt"`
	Ccy ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Ccy,omitempty"`
}

type AmountAndDirection43 struct {
	Amt CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Amt"`
	Sgn bool              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Sgn,omitempty"`
	Dt  ISODate           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Dt,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 MAC"`
}

// May be one of ICCD, AGNT, MERC, ACQR, ISSR, TRML
type AuthenticationEntity2Code string

// May be one of TOKA, BIOM, MOBL, OTHR, FPIN, NPIN, PSWD, SCRT, SCNL
type AuthenticationMethod7Code string

type AuthorisationResult13 struct {
	AuthstnNtty PartyType16Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AuthstnNtty,omitempty"`
	AuthstnRspn ResponseType7   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AuthstnRspn"`
	RspnTrac    []ResponseType8 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 RspnTrac,omitempty"`
	AuthstnCd   Min6Max8Text    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AuthstnCd,omitempty"`
	Actn        []Action7       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Actn,omitempty"`
}

type AutomatedTellerMachine2 struct {
	Id      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Id"`
	AddtlId Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AddtlId,omitempty"`
	SeqNb   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 SeqNb,omitempty"`
	BaseCcy ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 BaseCcy,omitempty"`
	Lctn    PostalAddress17    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Lctn,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{1,30}
type BBANIdentifier string

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CardAccount13 struct {
	AcctTp  CardAccountType3Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AcctTp,omitempty"`
	AcctNm  Max70Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AcctNm,omitempty"`
	Ccy     ActiveCurrencyCode            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Ccy,omitempty"`
	AcctIdr AccountIdentification31Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AcctIdr,omitempty"`
	CdtRef  Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CdtRef,omitempty"`
	Svcr    PartyIdentification72Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Svcr,omitempty"`
	BalBfr  AmountAndDirection43          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 BalBfr,omitempty"`
	BalAftr AmountAndDirection43          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 BalAftr,omitempty"`
}

// May be one of CTDP, CHCK, CRDT, CURR, CDBT, DFLT, EPRS, HEQL, ISTL, INVS, LCDT, MBNW, MNMK, MNMC, MTGL, RTRM, RVLV, SVNG, STBD, UVRL, PRPD, FLTC
type CardAccountType3Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 RltvDstngshdNm"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 EnvlpdData"`
}

type ContentInformationType15 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CnttTp"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AuthntcdData"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyAndAmount struct {
	Value float64      `xml:",chardata"`
	Ccy   CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

// May be one of ATMC, ATMP, APPR, CRAP, CPRC, OEXR, AMNT, LOCC, MNOC
type DataSetCategory7Code string

type DetailedAmount17 struct {
	AmtToTrf float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AmtToTrf"`
	Ccy      ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Ccy,omitempty"`
	Fees     []DetailedAmount18 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Fees,omitempty"`
	Dontn    []DetailedAmount18 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Dontn,omitempty"`
}

type DetailedAmount18 struct {
	Amt        float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Amt"`
	Ccy        ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Ccy,omitempty"`
	ChrgAcctTo bool               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ChrgAcctTo,omitempty"`
	Labl       Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Labl,omitempty"`
}

type Document struct {
	ATMTrfRspn ATMTransferResponseV01 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ATMTrfRspn"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 NcrptdCntt,omitempty"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, TEND
type Frequency3Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Issr,omitempty"`
}

type GenericIdentification77 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Id"`
	Tp     PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Tp"`
	Issr   PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 ShrtNm,omitempty"`
}

type GeographicCoordinates1 struct {
	Lat  Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Lat"`
	Long Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Long"`
}

type GeographicLocation1Choice struct {
	GeogcCordints GeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 GeogcCordints"`
	UTMCordints   UTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 UTMCordints"`
}

type Header31 struct {
	MsgFctn    ATMMessageFunction2 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 MsgFctn"`
	PrtcolVrsn Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PrtcolVrsn"`
	XchgId     Max3NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 XchgId"`
	CreDtTm    ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CreDtTm"`
	InitgPty   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 InitgPty"`
	RcptPty    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 RcptPty,omitempty"`
	PrcStat    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PrcStat,omitempty"`
	Tracblt    []Traceability4     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Tracblt,omitempty"`
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

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 NcrptdKey"`
}

type Max10000Binary []byte

func (t *Max10000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max104Text string

// Must be at least 1 items long
type Max10Text string

type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max20000Text string

type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max37Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

type Max5000Binary []byte

func (t *Max5000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max5000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max500Binary []byte

func (t *Max500Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max500Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max500Text string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// Must be at least 1 items long
type Max76Text string

// May be one of BALN, CMPA, CMPD, ACMD, DVCC, DIAQ, DIAP, GSTS, INQQ, INQP, KYAQ, KYAP, PINQ, PINP, RJAQ, RJAP, WITV, WITK, WITQ, WITP, INQC, H2AP, H2AQ, TMOP, CSEC, DSEC, SKSC, SSTS, DPSK, DPSV, DPSQ, DPSP, EXPK, EXPV, TRFQ, TRFP, RPTC
type MessageFunction11Code string

// May be one of BALN, GSTS, DSEC, INQC, KEYQ, SSTS
type MessageFunction8Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 6 items long
type Min6Max8Text string

// Must match the pattern [0-9]{8,28}
type Min8Max28NumericText string

// May be one of MREF, SREF, TEXT, HTML
type OutputFormat2Code string

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 BPddg,omitempty"`
}

type PartyIdentification72Choice struct {
	AnyBIC  AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AnyBIC"`
	PrtryId GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PrtryId"`
}

// May be one of ACQR, ATMG, CISP, DLIS, HSTG, ITAG, OATM
type PartyType12Code string

// May be one of ACQR, CISS, DLIS, ITAG, OTRM, BKAF, BKAT, ATMG
type PartyType16Code string

type PlainCardData19 struct {
	PAN       Min8Max28NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PAN,omitempty"`
	CardSeqNb Min2Max3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CardSeqNb,omitempty"`
	FctvDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 FctvDt,omitempty"`
	XpryDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 XpryDt,omitempty"`
	Trck1     Max76Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Trck1,omitempty"`
	Trck2     Max37Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Trck2,omitempty"`
	Trck3     Max104Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Trck3,omitempty"`
}

type PostalAddress17 struct {
	AdrLine     []Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 BldgNb,omitempty"`
	PstCd       Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PstCd,omitempty"`
	TwnNm       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 TwnNm"`
	CtrySubDvsn []Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Ctry"`
	GLctn       GeographicLocation1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 GLctn,omitempty"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 KeyIdr"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 KeyIdr"`
}

type RecurringTransaction3 struct {
	StartDt    ISODate        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 StartDt"`
	NbOfOcrncs float64        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 NbOfOcrncs,omitempty"`
	EndDt      ISODate        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 EndDt"`
	PrdUnit    Frequency3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 PrdUnit,omitempty"`
	IntrvlDay  float64        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 IntrvlDay,omitempty"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AttrVal"`
}

// May be one of APPR, DECL, PART
type Response4Code string

type ResponseType7 struct {
	Rspn         Response4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Rspn"`
	RspnRsn      ResultDetail4Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 RspnRsn,omitempty"`
	AddtlRspnInf Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AddtlRspnInf,omitempty"`
}

type ResponseType8 struct {
	RspndrId     Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 RspndrId"`
	Cdfctn       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Cdfctn,omitempty"`
	Rspn         Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Rspn"`
	RspnRsn      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 RspnRsn,omitempty"`
	AddtlRspnInf Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AddtlRspnInf,omitempty"`
}

// May be one of ACTF, ACQS, AMLV, AMTA, AUTH, BANK, CRDR, CRDF, ACTC, CTVG, DBER, FEES, TXNL, AMTD, NMBD, CRDX, FDCL, FMTR, TXNG, FNDI, ACPI, AMTI, ADDI, BRHI, CHDI, CRDI, CTFV, AMTO, PINV, TKKO, SGNI, TKID, TXNV, DATI, ISSP, ISSF, ISSO, ISST, ISSU, KEYS, LBLA, CRDL, MACR, MACK, ICCM, PINN, CRDA, LBLU, PINA, NPRA, OFFL, ONLP, NPRC, TXNM, OTHR, BALO, SEQO, PINC, PIND, PINS, PINX, PINE, QMAX, RECD, CRDT, SECV, SRVU, SFWE, SPCC, CRDS, SRCH, CNTC, FRDS, SYSP, SYSM, TRMI, ACTT, TTLV, TXNU, TXND, ORGF, UNBO, UNBP, UNBC, CMKY, CRDU, SVSU, VNDR, VNDF, AMTW, NMBW, CRDW, MEDI, SRVI
type ResultDetail4Code string

type SimpleIdentificationInformation4 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Id"`
}

// May be one of ASAP, CRIT, DTIM, ENCS
type TMSContactLevel2Code string

type TerminalHosting1 struct {
	Ctgy TransactionEnvironment3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Ctgy,omitempty"`
	Id   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Id,omitempty"`
}

type Traceability4 struct {
	RlayId      GenericIdentification77 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 RlayId"`
	SeqNb       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 SeqNb,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 TracDtTmOut"`
}

// May be one of BRCH, MERC, OTHR
type TransactionEnvironment3Code string

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 TxRef"`
}

type TransactionVerificationResult5 struct {
	Mtd         AuthenticationMethod7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Mtd"`
	VrfctnNtty  AuthenticationEntity2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 VrfctnNtty,omitempty"`
	Rslt        Verification1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Rslt,omitempty"`
	AddtlRslt   Max500Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AddtlRslt,omitempty"`
	AuthntcnTkn Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 AuthntcnTkn,omitempty"`
}

// Must match the pattern [0-9]{8,17}
type UPICIdentifier string

type UTMCoordinates1 struct {
	UTMZone    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 UTMZone"`
	UTMEstwrd  float64   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 UTMEstwrd"`
	UTMNrthwrd float64   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 UTMNrthwrd"`
}

// May be one of FAIL, MISS, NOVF, PART, SUCC, ERRR
type Verification1Code string

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
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

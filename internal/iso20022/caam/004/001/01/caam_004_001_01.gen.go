// Code generated by main. DO NOT EDIT.

package caam_004_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// ATMCommand1
//
// Maintenance command to perform on an ATM.
type ATMCommand1 struct {
	Tp        ATMCommand1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Tp"`
	Urgcy     TMSContactLevel2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Urgcy"`
	DtTm      ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DtTm,omitempty"`
	CmdId     ATMCommandIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 CmdId,omitempty"`
	CmdParams ATMCommandParameters1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 CmdParams,omitempty"`
}

// May be one of ABAL, ASTS, CFGT, CCNT, DISC, SNDM
type ATMCommand1Code string

// ATMCommandIdentification1
//
// Identification of the entity issuing the command.
type ATMCommandIdentification1 struct {
	Orgn Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Orgn,omitempty"`
	Ref  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Ref,omitempty"`
	Prcr Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Prcr,omitempty"`
}

// ATMCommandParameters1Choice
//
// Specific parameters attached to an ATM command.
type ATMCommandParameters1Choice struct {
	ATMReqrdGblSts  ATMStatus1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 ATMReqrdGblSts"`
	XpctdMsgFctn    MessageFunction8Code       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 XpctdMsgFctn"`
	ReqrdCfgtnParam ATMConfigurationParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 ReqrdCfgtnParam"`
}

// ATMConfigurationParameter1
//
// Configuration parameter version of the ATM.
type ATMConfigurationParameter1 struct {
	Tp   DataSetCategory7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Tp"`
	Vrsn Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Vrsn"`
}

// ATMEnvironment7
//
// Environment of the transaction.
type ATMEnvironment7 struct {
	Acqrr    Acquirer7               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Acqrr,omitempty"`
	ATMMgr   Acquirer8               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 ATMMgr,omitempty"`
	HstgNtty TerminalHosting1        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 HstgNtty,omitempty"`
	ATM      AutomatedTellerMachine3 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 ATM"`
}

// ATMEquipment3
//
// Hardware security module information, so called EPP for Encrypted PIN Pad.
type ATMEquipment3 struct {
	Manfctr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Manfctr,omitempty"`
	Mdl        Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Mdl,omitempty"`
	Vrsn       Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Vrsn,omitempty"`
	SrlNb      Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SrlNb,omitempty"`
	SgndSrlNb  ContentInformationType14 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SgndSrlNb,omitempty"`
	FrmwrPrvdr Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 FrmwrPrvdr,omitempty"`
	FrmwrId    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 FrmwrId,omitempty"`
	FrmwrVrsn  Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 FrmwrVrsn,omitempty"`
}

// ATMKeyDownloadResponse1
//
// Information related to the response of an ATM key download from an ATM manager.
type ATMKeyDownloadResponse1 struct {
	Envt          ATMEnvironment7     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Envt"`
	ATMSctyCntxt  ATMSecurityContext2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 ATMSctyCntxt"`
	ATMChllng     Max140Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 ATMChllng,omitempty"`
	HstSctyParams SecurityParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 HstSctyParams"`
	Cmd           []ATMCommand1       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Cmd,omitempty"`
}

// ATMKeyDownloadResponseV01
//
// The ATMKeyDownloadResponse message is sent from an acquirer to an ATM in response to an ATMKeyDownloadRequest message, to download of one or several cryptographic keys.
type ATMKeyDownloadResponseV01 struct {
	Hdr                   Header20                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Hdr"`
	PrtctdATMKeyDwnldRspn ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 PrtctdATMKeyDwnldRspn,omitempty"`
	ATMKeyDwnldRspn       ATMKeyDownloadResponse1  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 ATMKeyDwnldRspn,omitempty"`
	SctyTrlr              ContentInformationType13 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SctyTrlr,omitempty"`
}

// ATMMessageFunction1
//
// Identifies the type of process related to an ATM message.
type ATMMessageFunction1 struct {
	Fctn     MessageFunction7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Fctn"`
	ATMSvcCd Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 ATMSvcCd,omitempty"`
	HstSvcCd Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 HstSvcCd,omitempty"`
}

// ATMSecurityConfiguration1
//
// Configuration parameters in use by the security device.
type ATMSecurityConfiguration1 struct {
	Keys      ATMSecurityConfiguration2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Keys,omitempty"`
	Ncrptn    ATMSecurityConfiguration3 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Ncrptn,omitempty"`
	MACAlgo   []Algorithm12Code         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MACAlgo,omitempty"`
	DgstAlgo  []Algorithm11Code         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DgstAlgo,omitempty"`
	DgtlSgntr ATMSecurityConfiguration4 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DgtlSgntr,omitempty"`
	PIN       ATMSecurityConfiguration5 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 PIN,omitempty"`
	MsgPrtcn  []MessageProtection1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MsgPrtcn,omitempty"`
}

// ATMSecurityConfiguration2
//
// Configuration of the cryptographic keys.
type ATMSecurityConfiguration2 struct {
	MaxSmmtrcKey    float64 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MaxSmmtrcKey,omitempty"`
	MaxAsmmtrcKey   float64 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MaxAsmmtrcKey,omitempty"`
	MaxRSAKeyLngth  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MaxRSAKeyLngth,omitempty"`
	MaxRootKeyLngth float64 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MaxRootKeyLngth,omitempty"`
}

// ATMSecurityConfiguration3
//
// Configuration of the encryption or digital envelope, if the security module is able to perform encryption.
type ATMSecurityConfiguration3 struct {
	AsmmtrcNcrptn        bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 AsmmtrcNcrptn,omitempty"`
	AsmmtrcKeyStdId      bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 AsmmtrcKeyStdId,omitempty"`
	AsmmtrcNcrptnAlgo    []Algorithm7Code        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 AsmmtrcNcrptnAlgo,omitempty"`
	SmmtrcTrnsprtKey     bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SmmtrcTrnsprtKey,omitempty"`
	SmmtrcTrnsprtKeyAlgo []Algorithm13Code       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SmmtrcTrnsprtKeyAlgo,omitempty"`
	SmmtrcNcrptnAlgo     []Algorithm15Code       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SmmtrcNcrptnAlgo,omitempty"`
	NcrptnFrmt           []EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 NcrptnFrmt,omitempty"`
}

// ATMSecurityConfiguration4
//
// Configuration of the digital signatures if the security module is able to perform digital signatures with an asymmetric key.
type ATMSecurityConfiguration4 struct {
	MaxCerts      float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MaxCerts,omitempty"`
	MaxSgntrs     float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MaxSgntrs,omitempty"`
	DgtlSgntrAlgo []Algorithm14Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DgtlSgntrAlgo,omitempty"`
}

// ATMSecurityConfiguration5
//
// Configuration of the PIN online verification.
type ATMSecurityConfiguration5 struct {
	PINFrmt          []PINFormat4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 PINFrmt,omitempty"`
	PINLngthCpblties float64          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 PINLngthCpblties,omitempty"`
}

// ATMSecurityContext2
//
// Context of the ATM for the key download.
type ATMSecurityContext2 struct {
	CurSctySchme ATMSecurityScheme1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 CurSctySchme"`
	DvcPrprty    ATMEquipment3             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DvcPrprty,omitempty"`
	CurCfgtn     ATMSecurityConfiguration1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 CurCfgtn,omitempty"`
}

// May be one of APPK, CERT, FRAN, DTCH, LUXG, MANU, PKIP, SIGN, NONE
type ATMSecurityScheme1Code string

// May be one of INSV, OUTS
type ATMStatus1Code string

// Acquirer7
//
// Acquirer of the withdrawal transaction, in charge of the funds settlement with the issuer.
type Acquirer7 struct {
	AcqrgInstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 AcqrgInstn,omitempty"`
	Brnch      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Brnch,omitempty"`
}

// Acquirer8
//
// Institution in charge of managing the ATM.
type Acquirer8 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Id"`
	ApplVrsn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 ApplVrsn,omitempty"`
}

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5
type Algorithm12Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of ERS2, ERS1, RPSS
type Algorithm14Code string

// May be one of EA2C, E3DC, EA9C, EA5C
type Algorithm15Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

// AlgorithmIdentification11
//
// Cryptographic algorithms and parameters for the protection of transported keys by an asymmetric key.
type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Param,omitempty"`
}

// AlgorithmIdentification12
//
// Mask generator function cryptographic algorithm and parameters.
type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Param,omitempty"`
}

// AlgorithmIdentification13
//
// Cryptographic algorithm and parameters for the protection of the transported key.
type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Param,omitempty"`
}

// AlgorithmIdentification14
//
// Cryptographic algorithm and parameters for encryptions with a symmetric cryptographic key.
type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Param,omitempty"`
}

// AlgorithmIdentification15
//
// Identification of a cryptographic algorithm and parameters for the MAC computation.
type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Param,omitempty"`
}

// AlgorithmIdentification16
//
// Cryptographic algorithm and parameters of digests.
type AlgorithmIdentification16 struct {
	Algo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Algo"`
}

// AlgorithmIdentification17
//
// Identification of a cryptographic algorithm and parameters for digital signatures.
type AlgorithmIdentification17 struct {
	Algo  Algorithm14Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Algo"`
	Param Parameter8      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

// AuthenticatedData4
//
// Message authentication code (MAC), computed on the data to protect with an encryption key.
type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MAC"`
}

// AutomatedTellerMachine3
//
// ATM information.
type AutomatedTellerMachine3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Id"`
	AddtlId Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 AddtlId,omitempty"`
	SeqNb   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SeqNb,omitempty"`
	Lctn    PostalAddress17 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Lctn,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// CertificateIssuer1
//
// Certificate issuer name (see X.509).
type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 RltvDstngshdNm"`
}

// ContentInformationType10
//
// General cryptographic message syntax (CMS) containing encrypted data.
type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 EnvlpdData"`
}

// ContentInformationType13
//
// General cryptographic message syntax (CMS) containing data. protected by a MAC or a digital signature
type ContentInformationType13 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 CnttTp"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 AuthntcdData,omitempty"`
	SgndData     SignedData4        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SgndData,omitempty"`
}

// ContentInformationType14
//
// General cryptographic message syntax (CMS) containing data. protected by a digital signature
type ContentInformationType14 struct {
	CnttTp   ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 CnttTp"`
	SgndData SignedData4      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SgndData"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// CryptographicKey8
//
// Cryptographic key.
type CryptographicKey8 struct {
	Nm           Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Nm,omitempty"`
	Id           Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Id,omitempty"`
	SctyDomnId   Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SctyDomnId,omitempty"`
	AddtlId      Max35Binary               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 AddtlId,omitempty"`
	Vrsn         Max256Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Vrsn,omitempty"`
	SeqCntr      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SeqCntr,omitempty"`
	Tp           CryptographicKeyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Tp"`
	Fctn         []KeyUsage1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Fctn"`
	ActvtnDt     ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 ActvtnDt,omitempty"`
	DeactvtnDt   ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DeactvtnDt,omitempty"`
	KeyChckVal   Max35Binary               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 KeyChckVal,omitempty"`
	NcrptdKeyVal ContentInformationType10  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 NcrptdKeyVal,omitempty"`
	PblcKeyVal   PublicRSAKey1             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 PblcKeyVal,omitempty"`
}

// May be one of AES2, EDE3, DKP9, AES9, AES5, EDE4
type CryptographicKeyType3Code string

// May be one of ATMC, ATMP, APPR, CRAP, CPRC, OEXR, AMNT, LOCC, MNOC
type DataSetCategory7Code string

type Document struct {
	ATMKeyDwnldRspn ATMKeyDownloadResponseV01 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 ATMKeyDwnldRspn"`
}

// EncapsulatedContent3
//
// Data to authenticate.
type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Cntt,omitempty"`
}

// EncryptedContent3
//
// Encrypted data with an encryption key.
type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

// EnvelopedData4
//
// Encrypted data with encryption key.
type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 NcrptdCntt,omitempty"`
}

// GenericIdentification77
//
// Identification of an entity.
type GenericIdentification77 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Id"`
	Tp     PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Tp"`
	Issr   PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 ShrtNm,omitempty"`
}

// GeographicCoordinates1
//
// Location on the Earth specified by two numbers representing vertical and horizontal position.
type GeographicCoordinates1 struct {
	Lat  Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Lat"`
	Long Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Long"`
}

// GeographicLocation1Choice
//
// Geographic location of the ATM specified by geographic coordinates or UTM coordinates.
type GeographicLocation1Choice struct {
	GeogcCordints GeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 GeogcCordints"`
	UTMCordints   UTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 UTMCordints"`
}

// Header20
//
// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
type Header20 struct {
	MsgFctn    ATMMessageFunction1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MsgFctn"`
	PrtcolVrsn Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 PrtcolVrsn"`
	XchgId     Max3NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 XchgId"`
	CreDtTm    ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 CreDtTm"`
	InitgPty   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 InitgPty"`
	RcptPty    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 RcptPty,omitempty"`
	PrcStat    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 PrcStat,omitempty"`
	Tracblt    []Traceability4     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Tracblt,omitempty"`
}

// ISODateTime
//
// A particular point in the progression of time defined by a mandatory date and a mandatory time component, expressed in either UTC time format (YYYY-MM-DDThh:mm:ss.sssZ), local time with UTC offset format (YYYY-MM-DDThh:mm:ss.sss+/-hh:mm), or local time format (YYYY-MM-DDThh:mm:ss.sss). These representations are defined in "XML Schema Part 2: Datatypes Second Edition - W3C Recommendation 28 October 2004" which is aligned with ISO 8601.
// Note on the time format:
// 1) beginning / end of calendar day
// 00:00:00 = the beginning of a calendar day
// 24:00:00 = the end of a calendar day
// 2) fractions of second in time format
// Decimal fractions of seconds may be included. In this case, the involved parties shall agree on the maximum number of digits that are allowed.
type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// IssuerAndSerialNumber1
//
// Certificate issuer name and serial number  (see X.509).
type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SrlNb"`
}

// KEK4
//
// Key encryption key (KEK), using previously distributed symmetric key.
type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 NcrptdKey"`
}

// KEKIdentifier2
//
// Identification of a key encryption key (KEK), using previously distributed symmetric key.
type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DerivtnId,omitempty"`
}

// KeyTransport4
//
// Key encryption key (KEK), encrypted with a previously distributed asymmetric public key.
type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 NcrptdKey"`
}

// May be one of ENCR, DCPT, DENC, DDEC, TRNI, TRNX, MACG, MACV, SIGG, SUGV, PINE, PIND, PINV, KEYG, KEYI, KEYX, KEYD
type KeyUsage1Code string

// Max100KBinary
//
// Binary data of 100K maximum.
type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Max140Binary
//
// Specifies a binary string with a maximum length of 140 binary bytes.
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
type Max256Text string

// Max3000Binary
//
// Specifies a binary string with a maximum length of 3000 binary bytes.
type Max3000Binary []byte

func (t *Max3000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max3000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Max35Binary
//
// Specifies a binary string with a maximum length of 35 binary bytes.
type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Max5000Binary
//
// Specifies a binary string with a maximum length of 5000 binary bytes.
type Max5000Binary []byte

func (t *Max5000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max5000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Max500Binary
//
// Specifies a binary string with a maximum length of 500 binary bytes.
type Max500Binary []byte

func (t *Max500Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max500Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// May be one of BALN, CMPA, CMPD, ACMD, DVCC, DIAQ, DIAP, GSTS, INQQ, INQP, KYAQ, KYAP, PINQ, PINP, RJAQ, RJAP, WITV, WITK, WITQ, WITP, INQC, H2AP, H2AQ, TMOP, CSEC, DSEC, SKSC, SSTS
type MessageFunction7Code string

// May be one of BALN, GSTS, DSEC, INQC, KEYQ, SSTS
type MessageFunction8Code string

// May be one of EVLP, MACB, MACM, UNPR
type MessageProtection1Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// Min5Max16Binary
//
// Specifies a binary string with a minimum length of 5 bytes, and a maximum length of 16 bytes.
type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be one of ANSI, BNCM, BKSY, DBLD, DBLC, ECI2, ECI3, EMVS, IBM3, ISO0, ISO1, ISO2, ISO3, ISO4, ISO5, VIS2, VIS3
type PINFormat4Code string

// Parameter4
//
// Parameters of the asymmetric encryption algorithm.
type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MskGnrtrAlgo,omitempty"`
}

// Parameter5
//
// Parameters associated to a mask generator cryptographic function.
type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DgstAlgo,omitempty"`
}

// Parameter6
//
// Parameters associated to a cryptographic encryption algorithm.
type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 BPddg,omitempty"`
}

// Parameter7
//
// Parameters associated to the MAC algorithm.
type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 BPddg,omitempty"`
}

// Parameter8
//
// Parameters of the RSASSA-PSS digital signature algorithm (RSA signature algorithm with appendix: Probabilistic Signature Scheme).
type Parameter8 struct {
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DgstAlgo"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 MskGnrtrAlgo"`
	SaltLngth    float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SaltLngth"`
	TrlrFld      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 TrlrFld,omitempty"`
}

// May be one of ACQR, ATMG, CISP, DLIS, HSTG, ITAG, OATM
type PartyType12Code string

// PostalAddress17
//
// Information that locates and identifies a specific address, as defined by postal services or in free format text.
type PostalAddress17 struct {
	AdrLine     []Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 BldgNb,omitempty"`
	PstCd       Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 PstCd,omitempty"`
	TwnNm       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 TwnNm"`
	CtrySubDvsn []Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Ctry"`
	GLctn       GeographicLocation1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 GLctn,omitempty"`
}

// PublicRSAKey1
//
// Value of the public component of a RSA key.
type PublicRSAKey1 struct {
	Mdlus Max5000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Mdlus"`
	Expnt Max5000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Expnt"`
}

// Recipient4Choice
//
// Transport key or key encryption key (KEK) for the recipient.
type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 KeyIdr"`
}

// Recipient5Choice
//
// Identification of a cryptographic asymmetric key.
type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 KeyIdr"`
}

// RelativeDistinguishedName1
//
// Relative distinguished name defined by X.500 and X.509.
type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 AttrVal"`
}

// SecurityParameters5
//
// Security parameters of the host downloading the key.
type SecurityParameters5 struct {
	HstChllng Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 HstChllng,omitempty"`
	Key       []CryptographicKey8      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Key,omitempty"`
	DgtlSgntr ContentInformationType14 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DgtlSgntr,omitempty"`
}

// SignedData4
//
// Digital signatures of data from one or several signers.
type SignedData4 struct {
	Vrsn        float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 NcpsltdCntt"`
	Cert        []Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Cert,omitempty"`
	Sgnr        []Signer3                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Sgnr"`
}

// Signer3
//
// Entity who has signed the data and its digital signature.
type Signer3 struct {
	Vrsn      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Vrsn,omitempty"`
	SgnrId    Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SgnrId,omitempty"`
	DgstAlgo  AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 DgstAlgo"`
	SgntrAlgo AlgorithmIdentification17 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SgntrAlgo"`
	Sgntr     Max3000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Sgntr"`
}

// May be one of ASAP, CRIT, DTIM, ENCS
type TMSContactLevel2Code string

// TerminalHosting1
//
// Entity hosting the ATM terminal.
type TerminalHosting1 struct {
	Ctgy TransactionEnvironment3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Ctgy,omitempty"`
	Id   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Id,omitempty"`
}

// Traceability4
//
// Identification of partners involved in exchange from the ATM to the issuer, with the relative timestamp of their exchanges.
type Traceability4 struct {
	RlayId      GenericIdentification77 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 RlayId"`
	SeqNb       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 SeqNb,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 TracDtTmOut"`
}

// May be one of BRCH, MERC, OTHR
type TransactionEnvironment3Code string

// UTMCoordinates1
//
// Location on the Earth specified by the Universal Transverse Mercator coordinate system, using the WGS84 geodesic system.
type UTMCoordinates1 struct {
	UTMZone    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 UTMZone"`
	UTMEstwrd  float64   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 UTMEstwrd"`
	UTMNrthwrd float64   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 UTMNrthwrd"`
}

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
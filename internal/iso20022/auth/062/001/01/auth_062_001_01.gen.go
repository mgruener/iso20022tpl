// Code generated by main. DO NOT EDIT.

package auth_062_001_01

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type CCPLiquidityStressTestingDefinitionReportV01 struct {
	LqdtyStrssScnroDef []LiquidityStressScenarioDefinition1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 LqdtyStrssScnroDef"`
	SplmtryData        []SupplementaryData1                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 SplmtryData,omitempty"`
}

type Document struct {
	CCPLqdtyStrssTstgDefRpt CCPLiquidityStressTestingDefinitionReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 CCPLqdtyStrssTstgDefRpt"`
}

type GenericIdentification168 struct {
	Id      Max256Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 Id"`
	Desc    Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 Desc,omitempty"`
	Issr    Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 Issr,omitempty"`
	SchmeNm Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 SchmeNm,omitempty"`
}

type LiquidityStressScenarioDefinition1 struct {
	Id       GenericIdentification168 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 Id"`
	Desc     Max2000Text              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 Desc"`
	Tp       Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 Tp,omitempty"`
	StrssCcy ActiveCurrencyCode       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 StrssCcy"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max2000Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.062.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
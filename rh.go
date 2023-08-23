package molit

// RH : 연립다세대 정보
type RH struct {
	// 공통
	RegionalCode        string `xml:"지역코드" json:"지역코드"`
	ApartmentName       string `xml:"연립다세대" json:"연립다세대"`
	Dong                string `xml:"법정동" json:"법정동"`
	BuildYear           string `xml:"건축년도" json:"건축년도"`
	Jibun               string `xml:"지번" json:"지번"`
	AreaForExclusiveUse string `xml:"전용면적" json:"전용면적"`
	LandArea            string `xml:"대지권면적" json:"대지권면적"`
	Floor               string `xml:"층" json:"층"`

	// 매매
	DealAmount     string `xml:"거래금액,omitempty" json:"거래금액,omitempty"`
	DealYear       string `xml:"년,omitempty" json:"년,omitempty"`
	DealMon        string `xml:"월,omitempty" json:"월,omitempty"`
	DealDay        string `xml:"일,omitempty" json:"일,omitempty"`
	CancelDealType string `xml:"해제여부,omitempty" json:"해제여부,omitempty"`
	CancelDealDay  string `xml:"해제사유발생일,omitempty" json:"해제사유발생일,omitempty"`
	REQ_GBN        string `xml:"거래유형,omitempty" json:"거래유형,omitempty"`
	Rdealer_Lawdnm string `xml:"중개사소재지,omitempty" json:"중개사소재지,omitempty"`

	// 전월세
	ContractType                   string `xml:"계약구분,omitempty" json:"계약구분,omitempty"`
	TermOfContract                 string `xml:"계약기간,omitempty" json:"계약기간,omitempty"`
	Deposit                        string `xml:"보증금액,omitempty" json:"보증금액,omitempty"`
	MonthlyRent                    string `xml:"월세금액,omitempty" json:"월세금액,omitempty"`
	PreviousDeposit                string `xml:"종전계약보증금,omitempty" json:"종전계약보증금,omitempty"`
	PreviousMonthlyRent            string `xml:"종전계약월세,omitempty" json:"종전계약월세,omitempty"`
	UseRequestRenewalContractRight string `xml:"갱신요구권사용,omitempty" json:"갱신요구권사용,omitempty"`
}

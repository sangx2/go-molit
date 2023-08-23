package molit

// SH : 단독다가구 정보
type SH struct {
	// 공통
	RegionalCode   string `xml:"지역코드" json:"지역코드"`
	Dong           string `xml:"법정동" json:"법정동"`
	Jibun          string `xml:"지번" json:"지번"`
	HouseType      string `xml:"주택유형" json:"주택유형"`
	BuildYear      string `xml:"건축년도" json:"건축년도"`
	TotalFloorArea string `xml:"연면적" json:"연면적"`
	Plottage       string `xml:"대지면적" json:"대지면적"`
	DealYear       string `xml:"년" json:"년"`
	DealMon        string `xml:"월" json:"월"`
	DealDay        string `xml:"일" json:"일"`

	// 매매
	DealAmount     string `xml:"거래금액,omitempty" json:"거래금액,omitempty"`
	CancelDealType string `xml:"해제여부,omitempty" json:"해제여부,omitempty"`
	CancelDealDay  string `xml:"해제사유발생일,omitempty" json:"해제사유발생일,omitempty"`
	REQ_GBN        string `xml:"거래유형,omitempty" json:"거래유형,omitempty"`
	Rdealer_Lawdnm string `xml:"중개사소재지,omitempty" json:"중개사소재지,omitempty"`

	// 전월세
	ContractType                   string `xml:"계약구분,omitempty" json:"계약구분,omitempty"`
	TermOfContract                 string `xml:"계약기간,omitempty" json:"계약기간,omitempty"`
	DealBuildingArea               string `xml:"계약면적,omitempty" json:"계약면적,omitempty"`
	Deposit                        string `xml:"보증금액,omitempty" json:"보증금액,omitempty"`
	MonthlyRent                    string `xml:"월세금액,omitempty" json:"월세금액,omitempty"`
	PreviousDeposit                string `xml:"종전계약보증금,omitempty" json:"종전계약보증금,omitempty"`
	PreviousMonthlyRent            string `xml:"종전계약월세,omitempty" json:"종전계약월세,omitempty"`
	UseRequestRenewalContractRight string `xml:"갱신요구권사용,omitempty" json:"갱신요구권사용,omitempty"`
}

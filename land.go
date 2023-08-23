package molit

// Land : 땅 정보
type Land struct {
	//매물정보
	Sigungu      string `xml:"시군구" json:"시군구"`
	Dong         string `xml:"법정동" json:"법정동"`
	Zoning       string `xml:"용도지역" json:"용도지역"`
	LandUse      string `xml:"지목" json:"지목"`
	Jibun        string `xml:"지번" json:"지번"`
	RegionalCode string `xml:"지역코드" json:"지역코드"`

	// 거래정보
	DealAmount                      string `xml:"거래금액" json:"거래금액"`
	DealArea                        string `xml:"거래면적" json:"거래면적"`
	REQ_GBN                         string `xml:"거래유형" json:"거래유형"`
	DealYear                        string `xml:"년" json:"년"`
	DealMonth                       string `xml:"월" json:"월"`
	DealDay                         string `xml:"일" json:"일"`
	CancelDealType                  string `xml:"해제여부" json:"해제여부"`
	CancelDealDay                   string `xml:"해제사유발생일" json:"해제사유발생일"`
	ClassificationsOfPartialDealing string `xml:"지분거래구분" json:"지분거래구분"`
	RdealerLawdnm                   string `xml:"중개사소재지" json:"중개사소재지"`
}

package btwholesale

type SuggestedMessage struct {
	M1  string
	M2  string
	M3  string
	M4  string
	M5  string
	M6  string
	M7  string
	M8  string
	M9  string
	M10 string
	M11 string
	M12 string
	M13 string
	M14 string
	M15 string
	M16 string
	M17 string
	M18 string
	M19 string
	M20 string
	M21 string
	M22 string
	M23 string
}

type SupplementaryMessage struct {
	M1         string
	M2         string
	M3         string
	M4         string
	M5         string
	M6         string
	M7         string
	M8         string
	M9         string
	M10        string
	Bet_BR_MSG string
	Text1      string
}

type PremiseEnvTable struct {
	PremiseEnvTableFlag  string
	BridgeTapStatus      string
	VrlStatus            string
	NteFacePlateStatus   string
	LastRoutineDate      string
	RequestJsonDto       string
	FttpPriorityExchange string
	WlrWithdrawal        string
	SotapRestriction     string
}

type OtherOfferingTable struct {
	OtherOfferingTableFlag bool
	VdslAvailDate          string
	AdslAvailDate          string
	GfastAvailDate         string
}

type ObservedSpeed struct {
	GFast              bool
	Adsl               bool
	Vdsl               bool
	MaxObsDnSpeedGfast string
	MaxObsDnSpeedVdsl  string
	MaxObsDnSpeedadsl  string
	MaxObsUpSpeedGfast string
	MaxObsUpSpeedVdsl  string
	MaxObsUpSpeedadsl  string
	ObsDateGfast       string
	ObsDateVdsl        string
	ObsDateadsl        string
}

type FeaturedProductTable struct {
	FeaturedProductTableFlag     bool
	VdslRangeAHighDnStreamSpeed  string
	VdslRangeALowDnStreamSpeed   string
	VdslRangeAHighUpStreamSpeed  string
	VdslRangeALowUpStreamSpeed   string
	VdslRangeADnstreamThreshold  string
	VdslRangeALeftJumper         string
	VdslRangeBHighDnStreamSpeed  string
	VdslRangeBLowDnStreamSpeed   string
	VdslRangeBHighUpStreamSpeed  string
	VdslRangeBLowUpStreamSpeed   string
	VdslRangeBDnstreamThreshold  string
	VdslRangeBLeftJumper         string
	GfastRangeAHighDnStreamSpeed string
	GfastRangeALowDnStreamSpeed  string
	GfastRangeAHighUpStreamSpeed string
	GfastRangeALowUpStreamSpeed  string
	GfastRangeADnstreamThreshold string
	GfastRangeALeftJumper        string
	GfastRangeBHighDnStreamSpeed string
	GfastRangeBLowDnStreamSpeed  string
	GfastRangeBHighUpStreamSpeed string
	GfastRangeBLowUpStreamSpeed  string
	GfastRangeBDnstreamThreshold string
	GfastRangeBLeftJumper        string
	WbcFttpDnStreamSpeed         string
	WbcFttpUpStreamSpeed         string
	WbcFttpDnStreamRange         string
	WbcFttpAvailDate             string
	WbcFttpInstallProcess        string
	FttpOnDemandDnStreamSpeed    string
	FttpOnDemandUpStreamSpeed    string
	FttpOnDemandDnStreamRange    string
	FttpOnDemandAvailDate        string
	FttpOnDemandInstallProcess   string
	Fttp                         bool
	Fod                          bool
	VdslRangeAWBCFttcAvailDate   string
	VdslRangeAWBCSogeaAvailDate  string
	VdslRangeBWBCSogeaAvailDate  string
	VdslRangeBWBCFttcAvailDate   string
	GfastRangeAWBCFttcAvailDate  string
	GfastRangeAWBCSogeaAvailDate string
	GfastRangeBWBCSogeaAvailDate string
	GfastRangeBWBCFttcAvailDate  string
	WbcFttpAvail                 string
	FttpOnDemandAvail            string
}

type AdslTableFlag struct {
	AdslTableFlag               bool
	Adsl2DnStreamSpeed          string
	Adsl2UpStreamSpeed          string
	Adsl2DnStreamRange          string
	Adsl2UpStreamRange          string
	Adsl2AvailDate              string
	Soadsl2AvailDate            string
	Adsl2LeftJumper             string
	Adsl2AnnexMDnStreamSpeed    string
	Adsl2AnnexMUpStreamSpeed    string
	Adsl2AnnexMDnStreamRange    string
	Adsl2AnnexMUpStreamRange    string
	Adsl2AnnexMAvailDate        string
	Soadsl2AnnexMAvailDate      string
	Adsl2AnnexMLeftJumper       string
	AdslMaxDnStreamSpeed        string
	AdslMaxUpStreamSpeed        string
	AdslMaxDnStreamRange        string
	AdslMaxUpStreamRange        string
	AdslMaxAvailDate            string
	SoadslMaxAvailDate          string
	AdslMaxLeftJumper           string
	WbcFixedRateDnStreamSpeed   string
	WbcFixedRateUpStreamSpeed   string
	WbcFixedRateDnStreamRange   string
	WbcFixedRateUpStreamRange   string
	WbcFixedRateAvailDate       string
	SoadslWbcFixedRateAvailDate string
	WbcFixedRateLeftJumper      string
	FixedRateDnStreamSpeed      string
	FixedRateUpStreamSpeed      string
	FixedRateDnStreamRange      string
	FixedRateUpStreamRange      string
	FixedRateAvailDate          string
	SoadslFixedRateAvailDate    string
	FixedRateLeftJumper         string
	BetDnStreamSpeed            string
	BetUpStreamSpeed            string
	BetDnStreamRange            string
	BetUpStreamRange            string
	BetAvailDate                string
	SoadslBetAvailDate          string
	BetLeftJumper               string
	AdslTableLeftJumper         string
}

type BTWholeBoi struct {
	ErrorId                  int
	TelephoneNumber          string
	AccessLineId             string
	Uprn                     string
	BuildingNumber           string
	BuildingName             string
	Street                   string
	Town                     string
	PostCode                 string
	AdslTableFlag            AdslTableFlag
	FeaturedProductTable     FeaturedProductTable
	ObservedSpeed            ObservedSpeed
	OtherOfferingTable       OtherOfferingTable
	PremiseEnvTable          PremiseEnvTable
	SupplementaryMessage     SupplementaryMessage
	SuggestedMessage         SuggestedMessage
	ApigwTrackingHeader      string
	ResultStatus             string
	ResultText               string
	Sll                      string
	MarketSplit              string
	NetworkType              string
	ExchangeDistrictID       string
	ExchangeCode             string
	ExchangeName             string
	ExchangeBMF              string
	BridgeTapDetected        string
	VriDetected              string
	NtefaceplateDetected     string
	NteTypeRecorded          string
	NteFaceplateTypeRecorded string
	NteLocationRecorded      string
	DateNteDetailsRecorded   string
	CabinetNumber            string
	ListOfOpenOprder         string
	OrderTargetDate          string
	OrderType                string
	AddressMismatch          string
	DisplayHostAddress       string
	Location                 string
	Lijstatus                string
}

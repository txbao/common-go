package holidays

const (
	NewYearDay = iota
	SpringFestivalDay
	TombSweepingDay
	LaborDay
	DragonBoatFestivalDay
	NationalDay
	MidAutumnFestivalDay
)

var ChHolidays = [...]string{
	"元旦",
	"春节",
	"清明节",
	"劳动节",
	"端午节",
	"中秋节",
	"国庆节",
}
var EnHolidays = [...]string{
	"New Year\\'s Day",
	"Spring Festival",
	"Tomb-sweeping Day",
	"Labour Day",
	"Dragon Boat Festival",
	"Mid-autumn Festival",
	"National Day",
}

type CollectionYearHistory struct {
	Data [][]OneCollection `json:"data"`
}

func (h *CollectionYearHistory) Add(one []OneCollection) {
	h.Data = append(h.Data, one)
}

type OneCollection struct {
	Start  string `json:"start"`
	End    string `json:"end"`
	ChName string `json:"ch_name"`
	EnName string `json:"en_name"`
}
type YearCollection struct {
	Data []OneCollection `json:"data"`
}

func (y *YearCollection) Add(one OneCollection) {
	y.Data = append(y.Data, one)
}

func FetchCollectionYearHistory() CollectionYearHistory {
	return CollectionYearHistory{
		Data: [][]OneCollection{
			holiday2021,
		},
	}
}

var holiday2021 = []OneCollection{
	{
		Start:  "2022/01/01",
		End:    "2022/01/03",
		ChName: ChHolidays[NewYearDay], //元旦
		EnName: EnHolidays[NewYearDay],
	},
	{
		Start:  "2022/01/31",
		End:    "2022/02/06",
		ChName: ChHolidays[SpringFestivalDay], //春节
		EnName: EnHolidays[SpringFestivalDay],
	},
	{
		Start:  "2022/04/03",
		End:    "2022/04/05",
		ChName: ChHolidays[TombSweepingDay], //清明节
		EnName: EnHolidays[TombSweepingDay],
	},
	{
		Start:  "2022/05/01",
		End:    "2022/05/04",
		ChName: ChHolidays[LaborDay], //劳动节
		EnName: EnHolidays[LaborDay],
	},
	{
		Start:  "2022/06/03",
		End:    "2022/06/05",
		ChName: ChHolidays[DragonBoatFestivalDay], //端午节
		EnName: EnHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2022/09/10",
		End:    "2022/09/12",
		ChName: ChHolidays[MidAutumnFestivalDay], //中秋节
		EnName: EnHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2022/10/01",
		End:    "2022/10/07",
		ChName: ChHolidays[NationalDay], //国庆节
		EnName: EnHolidays[NationalDay],
	},
}

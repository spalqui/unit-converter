package pages

type Converter struct {
	Title    string
	Units    []string
	UnitTo   string
	UnitFrom string
	Value    float64
	Result   float64
	Error    string
}

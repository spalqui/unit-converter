package pages

type Converter struct {
	Type     string
	URLPath  string
	Units    []string
	UnitTo   string
	UnitFrom string
	Value    float64
	Result   float64
	Error    string
}

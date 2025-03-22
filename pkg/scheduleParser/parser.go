package scheduleParser

type Parser struct {
	IFefuParser
}

func NewParser() *Parser {
	return &Parser{
		IFefuParser: NewFefuParser(),
	}
}

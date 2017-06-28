package sofort

type Errors struct {
	XMLName              xml.Name `xml:"errors"`
	Field                []string `xml:"error>field"`
	CodeErrorErrorsSu    []string `xml:"su>errors>error>code"`
	MessageErrorErrorsSu []string `xml:"su>errors>error>message"`
	FieldErrorErrorsSu   []string `xml:"su>errors>error>field"`
	Code                 []string `xml:"error>code"`
	Message              []string `xml:"error>message"`
}

package sofort

type Warnings struct {
	XMLName xml.Name `xml:"warnings"`
	Code    []string `xml:"warning>code"`
	Message []string `xml:"warning>message"`
	Field   []string `xml:"warning>field"`
}

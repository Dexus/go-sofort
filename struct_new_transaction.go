package sofort

type New_transaction struct {
	XMLName     xml.Name `xml:"new_transaction"`
	Payment_url string   `xml:"payment_url"`
	Code        []string `xml:"warnings>warning>code"`
	Message     []string `xml:"warnings>warning>message"`
	Field       []string `xml:"warnings>warning>field"`
	Transaction string   `xml:"transaction"`
}

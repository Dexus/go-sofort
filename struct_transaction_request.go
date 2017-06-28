package sofort

type Transaction_request struct {
	XMLName                   xml.Name `xml:"transaction_request"`
	Version                   int      `xml:"version,attr"`
	Status                    string   `xml:"status"`
	Product                   string   `xml:"product"`
	From_time                 string   `xml:"from_time"`
	To_time                   string   `xml:"to_time"`
	Number                    string   `xml:"number"`
	From_status_modified_time string   `xml:"from_status_modified_time"`
	Transaction               []string `xml:"transaction"`
	To_status_modified_time   string   `xml:"to_status_modified_time"`
	Status_reason             string   `xml:"status_reason"`
	Page                      string   `xml:"page"`
}

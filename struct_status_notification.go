package sofort

type Status_notification struct {
	XMLName     xml.Name `xml:"status_notification"`
	Time        string   `xml:"time"`
	Transaction string   `xml:"transaction"`
}

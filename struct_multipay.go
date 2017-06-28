package sofort

type Multipay struct {
	XMLName                 xml.Name             `xml:"multipay"`
	Notification_url        []Notification_url   `xml:"notification_urls>notification_url"`
	Notification_email      []Notification_email `xml:"notification_emails>notification_email"`
	User_variable           []string             `xml:"user_variables>user_variable"`
	Success_url             string               `xml:"success_url"`
	Timeout_url             string               `xml:"timeout_url"`
	Bic                     string               `xml:"sender>bic"`
	Reason                  []string             `xml:"reasons>reason"`
	Abort_url               string               `xml:"abort_url"`
	Phone_customer          string               `xml:"phone_customer"`
	Project_id              string               `xml:"project_id"`
	Amount                  string               `xml:"amount"`
	Currency_code           string               `xml:"currency_code"`
	Holder                  string               `xml:"sender>holder"`
	Account_number          string               `xml:"sender>account_number"`
	Bank_code               string               `xml:"sender>bank_code"`
	Timeout                 string               `xml:"timeout"`
	Email_customer          string               `xml:"email_customer"`
	Country_code            string               `xml:"sender>country_code"`
	Country_codeBeneficiary string               `xml:"beneficiary>country_code"`
	Customer_protection     string               `xml:"su>customer_protection"`
	Language_code           string               `xml:"language_code"`
	Iban                    string               `xml:"sender>iban"`
	Success_link_redirect   string               `xml:"success_link_redirect"`
	Identifier              string               `xml:"beneficiary>identifier"`
	Interface_version       string               `xml:"interface_version"`
}

type Notification_url struct {
	Notify_on string `xml:"notify_on,attr"`
	Text      string `xml:",chardata"`
}
type Notification_email struct {
	Notify_on string `xml:"notify_on,attr"`
	Text      string `xml:",chardata"`
}

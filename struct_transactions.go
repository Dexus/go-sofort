package sofort

type Transactions struct {
	XMLName                                                                 xml.Name   `xml:"transactions"`
	Account_number                                                          []string   `xml:"transaction_details>sender>account_number"`
	Country_codeRecipientTransaction_details                                []string   `xml:"transaction_details>recipient>country_code"`
	BicRecipientTransaction_details                                         []string   `xml:"transaction_details>recipient>bic"`
	Currency_code                                                           []string   `xml:"transaction_details>costs>currency_code"`
	Phone_customer                                                          []string   `xml:"transaction_details>phone_customer"`
	Reason                                                                  [][]string `xml:"transaction_details>reasons>reason"`
	Project_id                                                              []string   `xml:"transaction_details>project_id"`
	User_variable                                                           [][]string `xml:"transaction_details>user_variables>user_variable"`
	Bank_codeRecipientTransaction_details                                   []string   `xml:"transaction_details>recipient>bank_code"`
	HolderRecipientTransaction_details                                      []string   `xml:"transaction_details>recipient>holder"`
	Status                                                                  []string   `xml:"transaction_details>status_history_items>status_history_item>status"`
	Status_reasonStatus_history_itemStatus_history_itemsTransaction_details []string   `xml:"transaction_details>status_history_items>status_history_item>status_reason"`
	Status_modified                                                         []string   `xml:"transaction_details>status_modified"`
	Bank_code                                                               []string   `xml:"transaction_details>sender>bank_code"`
	Iban                                                                    []string   `xml:"transaction_details>sender>iban"`
	Exchange_rate                                                           []string   `xml:"transaction_details>costs>exchange_rate"`
	Currency_codeTransaction_details                                        []string   `xml:"transaction_details>currency_code"`
	Language_code                                                           []string   `xml:"transaction_details>language_code"`
	Exchange_rateTransaction_details                                        []string   `xml:"transaction_details>exchange_rate"`
	Email_customer                                                          []string   `xml:"transaction_details>email_customer"`
	Amount                                                                  []string   `xml:"transaction_details>amount"`
	Bank_nameRecipientTransaction_details                                   []string   `xml:"transaction_details>recipient>bank_name"`
	Status_reason                                                           []string   `xml:"transaction_details>status_reason"`
	Time                                                                    []string   `xml:"transaction_details>status_history_items>status_history_item>time"`
	Holder                                                                  []string   `xml:"transaction_details>sender>holder"`
	Payment_method                                                          []string   `xml:"transaction_details>payment_method"`
	Amount_refunded                                                         []string   `xml:"transaction_details>amount_refunded"`
	TimeTransaction_details                                                 []string   `xml:"transaction_details>time"`
	StatusTransaction_details                                               []string   `xml:"transaction_details>status"`
	Test                                                                    []string   `xml:"transaction_details>test"`
	Bic                                                                     []string   `xml:"transaction_details>sender>bic"`
	Bank_name                                                               []string   `xml:"transaction_details>sender>bank_name"`
	Country_code                                                            []string   `xml:"transaction_details>sender>country_code"`
	Account_numberRecipientTransaction_details                              []string   `xml:"transaction_details>recipient>account_number"`
	Customer_protection                                                     []string   `xml:"transaction_details>su>customer_protection"`
	Transaction                                                             []string   `xml:"transaction_details>transaction"`
	IbanRecipientTransaction_details                                        []string   `xml:"transaction_details>recipient>iban"`
	Fees                                                                    []string   `xml:"transaction_details>costs>fees"`
}

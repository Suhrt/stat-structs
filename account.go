package structs

type Account struct {
	AccountNumber      string `json:"account_number" bson:"account_number"`
	Verified           bool   `json:"verified" bson:"verified"`
	VerificationStatus string `json:"verification_status" bson:"verification_status"`
	AccountName        string `json:"account_name" bson:"account_name"`
	ISFC               string `json:"isfc" bson:"isfc"`
	FundAccount        string `json:"fund_account" bson:"fund_account"`
	RazorPayContactId  string `json:"razorpay_contact_id" bson:"razorpay_contact_id"`
	Vpa                string `json:"vpa" bson:"vpa"`
}

package ubill

/*
TransactionInfo - DescribeOrderDetailInfo
*/
type TransactionInfo struct {
	TradeNo         string
	TransactionNo   string
	TransactionType string
	State           string
	PayPlatform     string
	//IncomeAmount    string
	//ExpenseAmount   string
	CreateTime      int
	BalanceSurplus  string
}


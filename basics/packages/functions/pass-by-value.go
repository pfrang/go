package main

func main() {
	costPerSend := 10
	numLastMonth := 100
	numThisMonth := 100

	monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth)
}

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	lastMonthBill := getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill := getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}

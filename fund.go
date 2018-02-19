package main

type Fund struct {
	balance int
}

func NewFund(initBal int) Fund {
	return Fund{
		balance: initBal,
	}
}

func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}

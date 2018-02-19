package main

import "fmt"

type FundServer struct {
	commands chan interface{}
	fund     Fund
}

type WithdrawCommand struct {
	Amount int
}

type BalanceCommand struct {
	Response chan int
}

type Transactor func(fund *Fund)

type TransactionCommand struct {
	Transactor Transactor
	Done       chan bool
}

func NewFundServer(initBal int) *FundServer {
	server := &FundServer{
		commands: make(chan interface{}),
		fund:     NewFund(initBal),
	}

	go server.loop()
	return server
}

func (s *FundServer) Balance() int {
	balanceResponseChan := make(chan int)
	s.commands <- BalanceCommand{Response: balanceResponseChan}
	return <-balanceResponseChan
}

func (s *FundServer) Withdraw(Amount int) {
	s.commands <- WithdrawCommand{Amount: 1}
}

func (s *FundServer) Transact(transactor Transactor) bool {
	done := make(chan bool)
	transCommand := TransactionCommand{
		Transactor: transactor,
		Done:       done,
	}
	s.commands <- transCommand
	return <-done
}

func (s *FundServer) loop() {
	for command := range s.commands {
		switch command.(type) {
		case WithdrawCommand:
			withdrawl := command.(WithdrawCommand)
			s.fund.Withdraw(withdrawl.Amount)
		case BalanceCommand:
			getBalance := command.(BalanceCommand)
			bl := s.fund.Balance()
			getBalance.Response <- bl
		case TransactionCommand:
			transCommand := command.(TransactionCommand)
			transCommand.Transactor(&s.fund)
			transCommand.Done <- true
		default:
			panic(fmt.Sprintf("Unrecongnized command : %v ", command))
		}

	}
}

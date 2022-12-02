package blockchain

import "strconv"

type Transaction struct {
	Id          int
	Sender      string //dni
	Receiver    string //dni
	Quantity    int
	SenderAcc   string
	ReceiverAcc string
}

func (t *Transaction) ToString() string {
	return "Transaction: " + strconv.Itoa(t.Id) + t.Sender + t.Receiver +
		strconv.Itoa(t.Quantity) + t.SenderAcc + t.ReceiverAcc
}

func CreateTransaction(_id int, _sender string, recv string, q int, sacc string, racc string) *Transaction {
	return &Transaction{_id, _sender, recv, q, sacc, racc}
}

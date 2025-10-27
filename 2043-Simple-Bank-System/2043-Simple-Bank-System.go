package leetcode

type Bank struct {
	accounts []int64
}

func Constructor(balance []int64) Bank {
	accounts := make([]int64, len(balance))
	copy(accounts, balance)

	return Bank{
		accounts,
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	account1--
	account2--

	if (account1 < 0 || account1 >= len(this.accounts)) ||
		(account2 < 0 || account2 >= len(this.accounts)) ||
		(this.accounts[account1] < money) {
		return false
	}

	this.accounts[account1] -= money
	this.accounts[account2] += money

	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	account--

	if account < 0 || account >= len(this.accounts) {
		return false
	}

	this.accounts[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	account--

	if account < 0 || account >= len(this.accounts) ||
		this.accounts[account] < money {
		return false
	}

	this.accounts[account] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */

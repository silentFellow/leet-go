package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		balance []int64
	}
	tests := []struct {
		name string
		args args
		want Bank
	}{
		{
			name: "example",
			args: args{balance: []int64{10, 100, 20, 50, 30}},
			want: Bank{accounts: []int64{10, 100, 20, 50, 30}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.balance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBank_Transfer(t *testing.T) {
	type fields struct {
		accounts []int64
	}
	type args struct {
		account1 int
		account2 int
		money    int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "valid transfer",
			fields: fields{accounts: []int64{10, 100, 20, 50, 30}},
			args:   args{account1: 5, account2: 1, money: 20},
			want:   true,
		},
		{
			name:   "invalid transfer (insufficient funds)",
			fields: fields{accounts: []int64{10, 100, 10, 50, 10}},
			args:   args{account1: 3, account2: 4, money: 15},
			want:   false,
		},
		{
			name:   "invalid transfer (account out of range)",
			fields: fields{accounts: []int64{10, 100, 20, 50, 30}},
			args:   args{account1: 10, account2: 1, money: 50},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Bank{
				accounts: tt.fields.accounts,
			}
			if got := this.Transfer(tt.args.account1, tt.args.account2, tt.args.money); got != tt.want {
				t.Errorf("Bank.Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBank_Deposit(t *testing.T) {
	type fields struct {
		accounts []int64
	}
	type args struct {
		account int
		money   int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "valid deposit",
			fields: fields{accounts: []int64{10, 100, 20, 50, 10}},
			args:   args{account: 5, money: 20},
			want:   true,
		},
		{
			name:   "invalid deposit (account out of range)",
			fields: fields{accounts: []int64{10, 100, 20, 50, 30}},
			args:   args{account: 10, money: 50},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Bank{
				accounts: tt.fields.accounts,
			}
			if got := this.Deposit(tt.args.account, tt.args.money); got != tt.want {
				t.Errorf("Bank.Deposit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBank_Withdraw(t *testing.T) {
	type fields struct {
		accounts []int64
	}
	type args struct {
		account int
		money   int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "valid withdraw",
			fields: fields{accounts: []int64{10, 100, 20, 50, 30}},
			args:   args{account: 3, money: 10},
			want:   true,
		},
		{
			name:   "invalid withdraw (insufficient funds)",
			fields: fields{accounts: []int64{10, 100, 10, 50, 30}},
			args:   args{account: 3, money: 15},
			want:   false,
		},
		{
			name:   "invalid withdraw (account out of range)",
			fields: fields{accounts: []int64{10, 100, 20, 50, 30}},
			args:   args{account: 10, money: 50},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Bank{
				accounts: tt.fields.accounts,
			}
			if got := this.Withdraw(tt.args.account, tt.args.money); got != tt.want {
				t.Errorf("Bank.Withdraw() = %v, want %v", got, tt.want)
			}
		})
	}
}

package pointers

import (
	"testing"
)

func TestWallet_Balance(t *testing.T) {
	tests := []struct {
		name string
		w    *Wallet
		want float64
	}{
		{
			name: "should start with balance set to 0",
			w:    &Wallet{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Balance(); got != tt.want {
				t.Errorf("Wallet.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWallet_Deposit(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		w       *Wallet
		args    args
		wantErr bool
	}{
		{
			name: "should accept positive amount for deposit",
			w:    &Wallet{},
			args: args{
				amount: 10,
			},
			wantErr: false,
		},
		{
			name: "should NOT accept negative amount for deposit",
			w:    &Wallet{},
			args: args{
				amount: -10,
			},
			wantErr: true,
		},
		{
			name: "should NOT accept zero amount for deposit",
			w:    &Wallet{},
			args: args{
				amount: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Deposit(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Wallet.Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWallet_Withdraw(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		w       *Wallet
		args    args
		wantErr bool
	}{
		{
			name: "should accept the withdraw if it's smaller then the balance",
			w:    &Wallet{50},
			args: args{
				amount: 10,
			},
			wantErr: false,
		},
		{
			name: "should reject the withdraw if it's lager then the balance",
			w:    &Wallet{50},
			args: args{
				amount: 100,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Withdraw(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Wallet.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBalanceWithWithdraw(t *testing.T) {
	w := &Wallet{}

	w.Deposit(100)
	w.Deposit(50)
	w.Withdraw(10)
	w.Withdraw(20)

	got := w.Balance()
	want := 120.0

	if got != want {
		t.Errorf("Failed the deposit and withdraw, got = %v, want = %v", got, want)
	}
}

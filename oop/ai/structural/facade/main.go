package main

type WalletFacade struct {
	account *Account
	wallet *Wallet
	notif *Norification
}

func (w *WalletFacade) DepositMoney(amount int) {
	w.account.Check()
	w.wallet.Credit(amount)
	w.notif.SencSuccess()
}
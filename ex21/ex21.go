package main

import "fmt"

/*Создаем константные переменные курса биткойна к валюте*/
const (
	bitcoinToDollarExchangeRate float64 = 23557.80
	bitcoinToRubleExchangeRate          = 1392297.44
)

/*Интерфейс, который конвертирует баланс в валюте в доллар*/
type CurrencyConvert interface {
	GetConvertBalance() float64
}

type Wallet struct{}

/*Метод структуры wallet*/
func (w *Wallet) showBalance(convert CurrencyConvert) {
	balance := convert.GetConvertBalance()
	fmt.Printf("On your account ₿: %.2f\n", balance)
}

type DollarWallet struct {
	balance float64
}

/*Имплементируем метод интерфейса структурой DollarWallet*/
func (d *DollarWallet) GetConvertBalance() float64 {
	fmt.Printf("On your account $: %.2f\n", d.balance)
	return d.balance / bitcoinToDollarExchangeRate
}

type RubleWallet struct {
	balance float64
}

/*Метод конвертации рублей в биткойн*/
func (r *RubleWallet) ConvertRubleToBitcoin() float64 {
	fmt.Printf("On your account ₽: %.2f\n", r.balance)
	return r.balance / bitcoinToRubleExchangeRate
}

/*Адаптер, который содержит указатель на кошел в рублях*/
type RubleWalletAdapter struct {
	rubleWallet *RubleWallet
}

/*Адаптер имплементирует метод интерфейса, тем самым мы можем в главном методе Wallet вызывать наш адаптер и конвертировать рубли в биткойн*/
func (adapter *RubleWalletAdapter) GetConvertBalance() float64 {
	return adapter.rubleWallet.ConvertRubleToBitcoin()
}

func main() {
	walet := Wallet{}
	dollarWallet := &DollarWallet{
		balance: 11111.11,
	}
	walet.showBalance(dollarWallet)
	adapter := &RubleWalletAdapter{
		&RubleWallet{
			balance: 1234523.23,
		},
	}
	walet.showBalance(adapter)
}

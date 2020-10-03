package wallet

import (
	"github.com/hirs-500/wallet/pkg/types"
	"github.com/google/uuid"
	"errors"
)

// ErrPhoneRegistered этот номер уже заригестрирован.
var ErrPhoneRegistered = errors.New("phone already registered")
//ErrAmountMustBePositive баланс дольжен бить выше 0.
var ErrAmountMustBePositive = errors.New("amount must be greater zero")
//ErrAccountNotFond аккаунт не найден 
var ErrAccountNotFond = errors.New("account not fond")
//ErrNotEgnoughBalance not egnough balance
var ErrNotEgnoughBalance = errors.New( "not egnough balance")
// Service информация 
type Service struct {
nextAccountID int64
accounts   []*types.Account
payments   []*types.Payment
}
//RegisterAccount регистрация новых аккунтов.
func (s *Service)RegisterAccount(phone types.Phone)(*types.Account, error)  {
for _, account := range s.accounts {
	if account.Phone == phone{
		return nil, ErrPhoneRegistered

	}	
}
s.nextAccountID++
account := &types.Account{
	ID:   s.nextAccountID,
	Phone:  phone,
	Balance:  0,
}

s.accounts= append(s.accounts,account )
return account,  nil

}
//Deposit информация 
func (s *Service) Deposit(accountID int64, amount types.Money) error {
if amount <=0 {
	return ErrAmountMustBePositive

}
var account *types.Account
for _, acc := range s.accounts {
	if acc.ID == accountID{
		account =acc
		break
	}
	if account ==nil {
		return ErrAccountNotFond
	}
	
}
account.Balance+=amount 
return nil 
}

//Pay информация 
func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory)(*types.Payment, error)  {
if amount <=0 {
	return nil, ErrAmountMustBePositive
}

var account *types.Account
for _, acc := range s.accounts {
if acc.ID == accountID{
	account = acc
	break
}

}
if account == nil{
return nil, ErrAccountNotFond
}
if account.Balance < amount{return nil, ErrNotEgnoughBalance
}
account.Balance -=amount
paymentID := uuid.New().String()
payment := &types.Payment{
ID:        paymentID,
AccountID: accountID,
Amount:    amount,
Category: category,
Status: types.PaymentStatusInProgress,
}
s.payments = append(s.payments, payment)
return payment, nil

}
//FindAccountByID функция поиск аккаунтов
func (s *Service)FindAccountByID(accountID int64,)(*types.Account, error)  {
var account *types.Account
for _, acc := range s.accounts {
	
	if acc.ID == accountID{
		account=acc
		break
	}
	if account ==nil {
		return nil, ErrAccountNotFond 
	}
}

return account, nil 
}

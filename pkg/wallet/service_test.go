package wallet

import ("testing")

func TestService_FindAccountByID_succes_user(t *testing.T) {
	var svc Service 
	svc.RegisterAccount("+992918630008")
   account, err := svc.FindAccountByID(1)
if err !=nil {
	t.Errorf("method returned not nil error, account => %v ", account)

}
}

func TestService_FindAccountByID_notFond_user(t *testing.T) {
	var svc Service 
	svc.RegisterAccount("+992918630008")
   account, err := svc.FindAccountByID(2)
if err ==nil {
	t.Errorf("method returned not nil error, account => %v ", account)

}
}
func TestService_Reject_success_user(t *testing.T){
	var svc Service
	svc.RegisterAccount("+99291800008")
	account, err := svc.FindAccountByID(1)

	if err != nil{
		t.Errorf("method RegisterAccount returned not nil error, error => %v", err)
	}


	err = svc.Deposit(account.ID, 100_00)
	if err != nil{
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}



	payment, err := svc.Pay(account.ID, 10_00,"Cafe")

	if err != nil{
		t.Errorf("method Pay returned not nil error, error => %v", err)
	}

	pay, err := svc.FindPaymentByID(payment.ID)

	if err != nil{
		t.Errorf("method FindPaymentByID returned not nil error, error => %v", err)
	}

	err = svc.Reject(pay.ID)

	if err != nil{
		t.Errorf("method Reject returned not nil error, error => %v", err)
	}



}

func TestService_Reject_fail_user(t *testing.T){
	var svc Service
	svc.RegisterAccount("+99918630008")
	account, err := svc.FindAccountByID(1)

	if err != nil{
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 999_99)
	if err != nil{
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}


	payment, err := svc.Pay(account.ID, 20_00,"Cafe")

	if err != nil{
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}

	pay, err := svc.FindPaymentByID(payment.ID)

	if err != nil{
		t.Errorf("method FindPaymentByID returned not nil error, payment => %v", payment)
	}

	err = svc.Reject(pay.ID+"uu")

	if err == nil{
		t.Errorf("method Reject returned not nil error, pay => %v", pay)
	}



}

func TestService_Repeat_success_user(t *testing.T){
	var svc Service
	
	account, err := svc.RegisterAccount("+992918630008")

	if err != nil{
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 88_888_88)
	if err != nil{
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}


	payment, err := svc.Pay(account.ID, 11_111,"fun")

	if err != nil{
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}

	newPayment, err := svc.FindPaymentByID(payment.ID)

	if err != nil{
		t.Errorf("method FindPaymentByID returned not nil error, payment => %v", payment)
	}

	paymentRepeat, err := svc.Repeat(newPayment.ID)

	if err != nil{
		t.Errorf("method Repeat returned not nil error, paymentRepeat => %v", paymentRepeat)
	}

}
func TestService_Favorite_success_user(t *testing.T){
	var svc Service
	
	account, err := svc.RegisterAccount("+992918630008")

	if err != nil{
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 99_999_99)
	if err != nil{
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}


	payment, err := svc.Pay(account.ID, 11_111_111,"club")

	if err != nil{
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}



	favorite, err := svc.FavoritePayment(payment.ID, "Hot Sex Tajik girl")

	if err != nil{
		t.Errorf("method FavoritePayment returned not nil error, favorite => %v", favorite)
	}

	favoritePay, err := svc.PayFromFavorite(favorite.ID)
	if err != nil{
		t.Errorf("method PayFromFavorite returned not nil error, paymentFavorite => %v", favoritePay)
	}



}
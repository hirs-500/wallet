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

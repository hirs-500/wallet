package main

import (
	"github.com/hirs-500/wallet/pkg/wallet"
)




func main ()  {
svc := &wallet.Service{}
svc.RegisterAccount("+22665225")
svc.RegisterAccount("+12665225")
svc.RegisterAccount("+02665225")
svc.ExportToFile("data/export")
}
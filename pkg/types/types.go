package types
// Money представляет собой денежную сумму в минимальных единицах (дирамы, рубль, центы, и т.д ).
type Money int64

// PaymentCategory представляет собой котегорию, который быль совершён платёж (авто, аптека, ресторан, т.д.).
type PaymentCategory string

// PaymentStatus представлет собой статус платжа.
type PaymentStatus string
// Предопределённый статус платежей.
const(
	PaymentStatusOk PaymentStatus = "OK"
	PaymentStatusFail PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

// Payment представляет информацию о платеже.
type Payment struct {
	ID     string
	AccountID int64
	Amount  Money//использовали Money 
	Category  PaymentCategory
	Status    PaymentStatus
}
// Phone номер 
 type Phone string 

//Account отстаток баланса
 type Account struct{
	 ID int64
	 Phone Phone
	 Balance Money 
 }


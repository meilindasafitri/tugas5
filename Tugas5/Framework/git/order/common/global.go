package common

//Struct API
// Order struct (Model) ...

type Message struct {
	Code    int     `json:"code"`
	Remark  string  `json:"remark"`
	OrderID string  `json:"orderID"`
	Orders  *Orders `json:"orders,omitempty"`
	Result  *Result `json:"result,omitempty"`
}

type Orders struct {
	OrderID    string         `json:"orderID"`
	CustomerID string         `json:"customerID"`
	EmployeeID string         `json:"employeeID"`
	OrderDate  string         `json:"orderDate"`
	OrdersDet  []OrdersDetail `json:"ordersDetail"`
}

type OrdersDetail struct {
	OrderID     string  `json:"orderID"`
	ProductID   string  `json:"ProductID"`
	ProductName string  `json:"ProductName"`
	UnitPrice   float64 `json:"UnitPrice"`
	Quantity    int     `json:"Quantity"`
}

type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}
type Customers struct {
	CustomerID   string `json:"CustomerID"`
	CompanyName  string `json:"CompanyName"`
	ContactName  string `json:"ContactName"`
	ContactTitle string `json:"ContactTitle"`
	Address      string `json:"Address"`
	City         string `json:"City"`
	Country      string `json:"Country"`
	Phone        string `json:"Phone"`
	PostalCode   string `json:"PostalCode"`
}
type Products struct {
	ProductID       string `json:"ProductID"`
	ProductName     string `json:"ProductName"`
	SupplierID      string `json:"SupplierID"`
	CategoryID      string `json:"CategoryID"`
	QuantityPerUnit string `json:"QuantityPerUnit"`
	UnitPrice       string `json:"UnitPrice"`
	UnitsInStock    string `json:"UnitsInStock"`
}

type FastPayRequest struct {
	Request    string `json:"request"`
	Merchant_ID string `json:"merchant_id"`
	Merchant   string `json:"merchant"`
	Signature  string `json:"signature"`
}

type FastPayResponse struct {
	Response     	string `json:"response"`
	Merchant_ID   string `json:"merchant_id"`
	Merchant      string `json:"merchant"`
	PaymentChan []PaymentChannel   `json:"payment_channel"`	
	ResponseCode string `json:"response_code"`
	ResponseDesc string `json:"response_desc"`
}

type PaymentChannel struct {
	PgCode    string `json:"pg_code"`
	PgName   	string `json:"pg_name"`
}
//End Struct API

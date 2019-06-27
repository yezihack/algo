package data

type Ticket struct {
	Name         string
	Denomination int //元
	Cost         int //多少个币
	Number       int
}

//初使优惠券列表
func InitTicket() (result []Ticket) {
	result = append(result, Ticket{
		Name:         "优惠券1",
		Denomination: 1,
		Cost:         10,
		Number:       1,
	})
	result = append(result, Ticket{
		Name:         "优惠券2",
		Denomination: 5,
		Cost:         50,
		Number:       0,
	})
	result = append(result, Ticket{
		Name:         "优惠券3",
		Denomination: 10,
		Cost:         100,
		Number:       0,
	})
	result = append(result, Ticket{
		Name:         "优惠券4",
		Denomination: 15,
		Cost:         150,
		Number:       0,
	})
	result = append(result, Ticket{
		Name:         "优惠券5",
		Denomination: 20,
		Cost:         200,
		Number:       0,
	})
	result = append(result, Ticket{
		Name:         "优惠券6",
		Denomination: 25,
		Cost:         250,
		Number:       0,
	})
	return
}

package main

import (
	"fmt"
	"sort"
)

var CoinNumber = 1077

func main() {

	//PackBest(-1)
	//PackBest(0)
	//PackBest(1)
	//PackBest(120)
	PackBest(300, 820)
	//PackBest(200)
	//PackBest(310)
	//PackBest(610)
	//PackBest(5000)
	//PackBest(10000)
	//PackBest(80000)
	//PackBest(80230)
}

func PackBest(coinReq, fee int) {
	CoinNumber = coinReq
	fmt.Println("fee", fee)
	tickets := InitTicket()
	best := make([]Ticket, 0)
	// 对tickets排序，大的在前
	sort.Slice(tickets, func(i, j int) bool { return tickets[i].Cost > tickets[j].Cost })
	lastTicket := tickets[len(tickets)-1]
	for {
		if fee <= 0 {
			break
		}
		storeNumber := 0 //每次循环计算一下库存数
		for k, ticket := range tickets {
			storeNumber += ticket.Number
			if ticket.Number <= 0 {
				continue
			}
			c := CoinNumber / ticket.Cost
			c1 := fee / (ticket.Denomination * 100)
			if c > c1 {
				c = c1
			}
			//如果倍数大于库存数量则取库存
			if c > ticket.Number {
				c = ticket.Number
			}
			fmt.Println("倍数: ", c, "number", ticket.Number, k, "last", len(tickets)-1)
			if c == ticket.Number && k == len(tickets)-1 {
				fmt.Println("------", k, ticket, "c", c)
				for iii := k - 1; iii >= 0; iii-- {
					if tickets[iii].Number > 0 {
						if fee > c*ticket.Denomination*100 {
							// fmt.Println("**************************")
							best = append(best, Ticket{
								Name:         tickets[iii].Name,
								Denomination: tickets[iii].Denomination,
								Cost:         tickets[iii].Cost,
								Number:       1,
							})
							fee -= tickets[iii].Denomination * 100
							CoinNumber -= tickets[iii].Cost
							tickets[iii].Number -= 1
							c = 0
						}
						break
					}
				}
				if fee <= 0 {
					break
				}
			}
			// 个数为0跳过
			if c == 0 {
				continue
			}
			//库存需要减去当前求的倍数
			tickets[k].Number -= c
			//如何倍数大于当前券的库存,重新对券组赋值
			if c >= ticket.Number {
				newTickets := make([]Ticket, 0)
				for _, vvv := range tickets {
					if vvv.Denomination != ticket.Denomination {
						newTickets = append(newTickets, vvv)
					}
				}
				tickets = newTickets
			}
			//获取已经得到的总值,总价值, 总币数
			total, coin := getTotal(best)
			//预算下一次的情况
			total += ticket.Denomination * c * 100
			coin += ticket.Cost * c

			// 证明币可能够用
			if coin <= CoinNumber && ticket.Denomination*c*100 <= fee {
				best = append(best, Ticket{
					Name:         ticket.Name,
					Denomination: ticket.Denomination,
					Cost:         ticket.Cost,
					Number:       c,
				})
				//减去已选中的价值和数量
				fee -= ticket.Denomination * c * 100
				CoinNumber -= ticket.Number * c
				break
			}
		}
		fmt.Println("CoinNumber", CoinNumber, "fee", fee, "last", lastTicket)
		//跳出条件
		if CoinNumber <= lastTicket.Cost || fee < lastTicket.Denomination*100 || len(tickets) == 0 || storeNumber <= 0 {
			break
		}
	}
	// 如果剩余的抵扣券还有余量则扣最小的
	if len(tickets) > 0 && fee > 0 {
		// 正序排序
		sort.Slice(tickets, func(i, j int) bool { return tickets[i].Cost < tickets[j].Cost })
		for _, v := range tickets {
			fmt.Println(v.Denomination, v.Number)
			if v.Number > 0 {
				//当前的币还允许购买一个的情况下,则购买
				if CoinNumber >= v.Cost {
					best = append(best, Ticket{
						Name:         v.Name,
						Denomination: v.Denomination,
						Cost:         v.Cost,
						Number:       1,
					})
					break
				}
			}
		}
	}
	// 合并相同卷
	newBest := make([]Ticket, 0)
	for _, v := range best {
		isExist := false
		for k1, v1 := range newBest {
			if v.Denomination == v1.Denomination {
				newBest[k1].Number++
				isExist = true
			}
		}
		if isExist == false {
			newBest = append(newBest, v)
		}
	}

	fmt.Println("最终统计卷")
	for _, item := range newBest {
		fmt.Printf("name:%s, Denomination:%d, cost:%d, number:%d\n", item.Name, item.Denomination, item.Cost, item.Number)
	}
	total, coin := getTotal(newBest)
	fmt.Printf("总共抵扣 %d 分，消耗 %d 个币\n", total, coin)
	fmt.Println("-----------------")
}

//返回:价值:分, 币个数
func getTotal(best []Ticket) (int, int) {
	total := 0
	coin := 0
	for _, item := range best {
		total += item.Number * item.Denomination * 100
		coin += item.Number * item.Cost
	}
	return total, coin
}

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
		Number:       1,
	})
	result = append(result, Ticket{
		Name:         "优惠券5",
		Denomination: 20,
		Cost:         200,
		Number:       1,
	})

	return
}

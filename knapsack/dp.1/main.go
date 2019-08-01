package main

import (
	"fmt"
	"sort"
)

//考虑库存数量,库存不足,不能选则最佳优惠券组合中,必须另选一个
func main() {
	BestTicket(300, 820)
	BestTicket(1000, 3600)
	BestTicket(310, 820)
	BestTicket(310, 2)
	BestTicket(310, 120)
}

//
func BestTicket(coinTotal, feeTotal int) {
	fmt.Println("------------------------------------------")
	//保留原始数据
	oriCoin, oriFee := coinTotal, feeTotal
	tickets := InitTicket()
	newTicket := make([]Ticket, 0)
	//倒序
	sort.Slice(tickets, func(i, j int) bool {
		return tickets[i].Denomination > tickets[j].Denomination
	})
	for i := 0; i < len(tickets); i++ {
		//只有当库存为0时,将库存小于0的移除掉
		tickets = resetTicket(tickets)
		// fmt.Println("key", k)
		for key, ticket := range tickets {
			// fmt.Printf("kkkkkkkkkkkkkk:%+v\n", key)
			//当库存等于0时,则重置切片
			if ticket.Number <= 0 {
				continue
			}
			//当币小于库存币则退出,或抵扣费用小于券费用
			cc := feeTotal / (ticket.Denomination * 100)
			// mc := coinTotal % ticket.Cost
			//当求出的倍数大于库存,则使用当前物品的库存
			flag := false //当倍数大于库存则表未,这个券不够用.必须要找下一个券
			if cc > ticket.Number {
				cc = ticket.Number
				flag = true
			}
			if flag == true {
				//当倍数大于当前数量时,也就说明当前券不够使用,则需要去寻找上一个券(因为切片是倒序)
				for j := key - 1; j >= 0; j-- {
					if tickets[j].Number > 0 {
						if feeTotal > cc*ticket.Denomination*100 {
							// fmt.Println("**************************")
							newTicket = append(newTicket, Ticket{
								Name:         tickets[j].Name,
								Denomination: tickets[j].Denomination,
								Cost:         tickets[j].Cost,
								Number:       1,
							})
							feeTotal -= tickets[j].Denomination * 100
							coinTotal -= tickets[j].Cost
							tickets[j].Number -= 1
							cc = 0
						}
						break
					}
				}
			}
			// tt := coinTotal / ticket.Cost
			//  fmt.Printf("倍数:%d, ticket:%v\n", cc, ticket)
			//当券数量充足,币也充足,抵扣金额也大于或等于当前券价值则进行抵扣
			if cc > 0 && ticket.Number > 0 &&
				ticket.Cost*cc <= coinTotal &&
				ticket.Denomination*cc*100 <= feeTotal {
				//fmt.Println("ssss", ticket)
				newTicket = append(newTicket, Ticket{
					Name:         ticket.Name,
					Denomination: ticket.Denomination,
					Cost:         ticket.Cost,
					Number:       cc,
				})
				//减去已选中的价值和数量
				feeTotal -= ticket.Denomination * cc * 100
				coinTotal -= ticket.Cost * cc
				//减库存
				tickets[key].Number -= cc
			}
		}
	}
	if len(tickets) > 0 {
		lastTicket := tickets[len(tickets)-1]
		//若未抵扣完全,则再补一个券(条件必须币够用)
		if feeTotal > 0 && coinTotal >= lastTicket.Cost {
			newTicket = append(newTicket, Ticket{
				Name:         lastTicket.Name,
				Denomination: lastTicket.Denomination,
				Cost:         lastTicket.Cost,
				Number:       1,
			})
			feeTotal -= lastTicket.Denomination * 1 * 100
			coinTotal -= lastTicket.Cost * 1
		}
	}
	//寻找最优方案
	//对数组倒序
	sort.Slice(newTicket, func(i, j int) bool {
		return newTicket[i].Denomination > newTicket[j].Denomination
	})
	fmt.Println("次优", newTicket)
	bestValue := 0
	for k, v := range newTicket {
		if bestValue > oriFee {
			fmt.Println("多余的数据", v, "value", bestValue, k)
			newTicket = append(newTicket[:k], newTicket[k+1:]...)
		}
		bestValue += v.Denomination * 100
	}
	//合并相同券的数量,时间复杂度:O(n^2)
	bestTicket := make([]Ticket, 0)
	for _, item := range newTicket {
		flag := false
		for k, b := range bestTicket {
			if b.Name == item.Name {
				bestTicket[k].Number += item.Number
				flag = true
			}
		}
		if flag == false {
			bestTicket = append(bestTicket, item)
		}
	}
	PrintTicket(bestTicket, oriCoin, oriFee, coinTotal, feeTotal)
}

func PrintTicket(list []Ticket, coin, fee, YuCoin, YuFee int) {
	fmt.Printf("预抵扣费用:%d(分), 兑换前币数量:%d(个), 余费用:%d(分), 余币数:%d(个)\n", fee, coin, YuFee, YuCoin)
	coinTotal, feeTotal := 0, 0
	for _, item := range list {
		coinTotal += item.Cost * item.Number
		feeTotal += item.Denomination * item.Number * 100
		fmt.Printf("券名称:%s, 券价值:%d 分, 券售价: %d 币, 数量:%d\n", item.Name, item.Denomination*100, item.Cost, item.Number)
	}
	fmt.Printf("实际抵扣费用:%d(分), 实际消耗币数量:%d(个)\n", feeTotal, coinTotal)
}
func ShowTicket(list []Ticket) {
	for _, item := range list {
		fmt.Printf("券名称:%s, 券价值:%d 分, 券售价: %d 币, 数量:%d\n", item.Name, item.Denomination*100, item.Cost, item.Number)
	}
}

//移除库存小于0的
func resetTicket(t []Ticket) (result []Ticket) {
	result = make([]Ticket, 0)
	for _, item := range t {
		if item.Number > 0 {
			result = append(result, item)
		}
	}
	//	ShowTicket(result)
	return
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
		Number:       2,
	})
	result = append(result, Ticket{
		Name:         "优惠券2",
		Denomination: 5,
		Cost:         50,
		Number:       1,
	})
	result = append(result, Ticket{
		Name:         "优惠券3",
		Denomination: 10,
		Cost:         100,
		Number:       1,
	})
	result = append(result, Ticket{
		Name:         "优惠券4",
		Denomination: 15,
		Cost:         150,
		Number:       4,
	})
	result = append(result, Ticket{
		Name:         "优惠券5",
		Denomination: 20,
		Cost:         200,
		Number:       2,
	})
	return
}

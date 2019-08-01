package main

import (
	"fmt"
	"sort"

	"git.sgfoot.com/studyALG/demo1/dp/data"
)

//不考虑库存数量,选则最优优惠券组合
//时间复杂度: O(n^2)
func main() {
	BestTicket(100, 820)
	BestTicket(10, 820)
	BestTicket(310, 820)
	BestTicket(310, 2)
	BestTicket(310, 120)
}

//
func BestTicket(coinTotal, feeTotal int) {
	c, f := coinTotal, feeTotal
	tickets := data.InitTicket()
	newTicket := make([]data.Ticket, 0)
	//倒序
	sort.Slice(tickets, func(i, j int) bool {
		return tickets[i].Denomination > tickets[j].Denomination
	})
	lastTicket := tickets[len(tickets)-1]
	for i := 0; i < len(tickets); i++ {
		for _, ticket := range tickets {
			//当币小于库存币则退出,或抵扣费用小于券费用
			cc := feeTotal / (ticket.Denomination * 100)
			if cc > 0 && ticket.Cost*cc <= coinTotal && ticket.Denomination*cc*100 <= feeTotal {
				newTicket = append(newTicket, data.Ticket{
					Name:         ticket.Name,
					Denomination: ticket.Denomination,
					Cost:         ticket.Cost,
					Number:       cc,
				})
				//减去已选中的价值和数量
				feeTotal -= ticket.Denomination * cc * 100
				coinTotal -= ticket.Cost * cc
			}
		}
	}
	if feeTotal > 0 && coinTotal >= lastTicket.Cost {
		newTicket = append(newTicket, data.Ticket{
			Name:         lastTicket.Name,
			Denomination: lastTicket.Denomination,
			Cost:         lastTicket.Cost,
			Number:       1,
		})
		feeTotal -= lastTicket.Denomination * 1 * 100
		coinTotal -= lastTicket.Cost * 1
	}
	bestTicket := make([]data.Ticket, 0)
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
	PrintTicket(bestTicket, c, f, coinTotal, feeTotal)
}

func PrintTicket(list []data.Ticket, coin, fee, YuCoin, YuFee int) {
	fmt.Println("------------------------------------------")
	fmt.Printf("预抵扣费用:%d, 兑换前币数量:%d, 余费用:%d, 余币数:%d\n", fee, coin, YuFee, YuCoin)
	coinTotal, feeTotal := 0, 0
	for _, item := range list {
		coinTotal += item.Cost * item.Number
		feeTotal += item.Denomination * item.Number * 100
		fmt.Printf("券名称:%s, 券价值:%d 分, 券售价: %d 币, 数量:%d\n", item.Name, item.Denomination*100, item.Cost, item.Number)
	}
	fmt.Printf("实际抵扣费用:%d, 实际消耗币数量:%d\n", feeTotal, coinTotal)
}
func ShowTicket(list []data.Ticket) {
	for _, item := range list {
		fmt.Printf("券名称:%s, 券价值:%d 分, 券售价: %d 币, 数量:%d\n", item.Name, item.Denomination*100, item.Cost, item.Number)
	}
}

//返回:价值:分, 币个数
func getTotal(best []data.Ticket) (int, int) {
	total := 0
	coin := 0
	for _, item := range best {
		total += item.Number * item.Denomination * 100
		coin += item.Number * item.Cost
	}
	return total, coin
}

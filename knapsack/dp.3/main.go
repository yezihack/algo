package main

/******************
* 功能:实现最佳优惠券组合,考虑库存
* 版本: v0.3
* 时间:2018/12/31 9:27
********************/
import (
	"fmt"
	"sort"

	. "github.com/yezihack/studyALG/demo1/dp.3/data"
)

//考虑库存数量,库存不足,不能选则最佳优惠券组合中,必须另选一个
func main() {
	BestTicket(800, 132)
	BestTicket(132, 800)
	BestTicket(100, 10)
	BestTicket(10, 100)
	BestTicket(200, 12000)
	BestTicket(2000, 12000)
	BestTicket(800, 1200)
	BestTicket(0, 1200)
	BestTicket(10, 1)
	BestTicket(0, 0)
	BestTicket(5000, 4990)
}

//计算最优组合
//coinTotal 多少个币;feeTotal 需要抵扣多少钱(分)
func BestTicket(coinTotal, feeTotal int) {
	fmt.Println("-------------------开始选择最优组合-----------------------")
	//保留原始数据
	oriCoin, oriFee := coinTotal, feeTotal
	//加载已有的优惠券
	tickets := InitTicket()
	//声明一个组合切片
	newTicket := make([]Ticket, 0)
	//倒序
	sort.Slice(tickets, func(i, j int) bool {
		return tickets[i].Denomination > tickets[j].Denomination
	})
	//检查库存,踢出库存小于0的元素
	tickets = resetTicket(tickets)
	//ShowTicket(tickets)
	//设置最外层循环,时间复杂度:O(n^2),当flag=true,O(n^3)
	for i := 0; i < len(tickets); i++ {
		//只有当库存为0时,将库存小于0的移除掉
		if tickets[i].Number <= 0 {
			tickets = resetTicket(tickets)
		}
		if feeTotal <= 0 {
			break
		}
		//设置内层循环
		for key, ticket := range tickets {
			if feeTotal <= 0 {
				break
			}
			//当库存等于0时,则跳过
			if ticket.Number <= 0 {
				continue
			}
			//当币小于库存币则退出,或抵扣费用小于券费用
			cc := feeTotal / (ticket.Denomination * 100)
			cc2 := coinTotal / ticket.Cost
			//当cc或cc2其中有一个为0时则选择最大值;如何cc,cc2都大于0时,选择最小值
			if cc > 0 && cc2 > 0 {
				if cc > cc2 {
					cc = cc2
				}
			} else {
				if cc < cc2 {
					cc = cc2
				}
			}
			//当求出的倍数大于库存,则使用当前物品的库存
			flag := false //当倍数大于库存则表未,这个券不够用.必须要找下一个券
			if cc > ticket.Number {
				cc = ticket.Number
				flag = true
			}
			if flag == true {
				//当倍数大于当前数量时,也就说明当前券不够使用,则需要去寻找上一个券(因为切片是倒序)
				for j := key - 1; j >= 0; j-- {
					//寻找上一个元素时,要求:数量>0, feeTotal大于价值
					if tickets[j].Number > 0 && coinTotal > tickets[j].Cost && feeTotal > cc*ticket.Denomination*100 {
						newTicket = append(newTicket, Ticket{
							Name:         tickets[j].Name,
							Denomination: tickets[j].Denomination,
							Cost:         tickets[j].Cost,
							Number:       1,
						})
						feeTotal -= tickets[j].Denomination * 100
						coinTotal -= tickets[j].Cost
						tickets[j].Number -= 1
						//将倍数设置为0
						cc = 0
						break
					}
				}
			}
			//当券数量充足,币也充足,抵扣金额也大于或等于当前券价值则进行抵扣,费用需要一个2的系数,因为当币充足,优先大额抵扣.
			if cc > 0 && ticket.Number > 0 &&
				ticket.Cost*cc <= coinTotal &&
				ticket.Denomination*cc*100 <= feeTotal*2 {
				newTicket = append(newTicket, Ticket{
					Name:         ticket.Name,
					Denomination: ticket.Denomination,
					Cost:         ticket.Cost,
					Number:       cc,
				})
				//减去已选中的价值和数量,减库存
				feeTotal -= ticket.Denomination * cc * 100
				coinTotal -= ticket.Cost * cc
				tickets[key].Number -= cc
			}
		}
	}
	//当双循环结束后,再判断是否完全抵扣完,前提:必须tickets还存在数据
	if len(tickets) > 0 {
		//获取最后一个元素
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
			tickets[len(tickets)-1].Number -= 1
		}
	}
	//进一步优化最优方案
	//对数组倒序
	sort.Slice(newTicket, func(i, j int) bool {
		return newTicket[i].Denomination > newTicket[j].Denomination
	})
	bestValue := 0 //设置一个最佳总值
	goodTicket := make([]Ticket, 0)
	for _, item := range newTicket {
		if bestValue > oriFee { //累加结果大于需要抵扣的则跳出
			//fmt.Println("多余的数据", item, "value", bestValue)
			//退回币和抵扣
			coinTotal += item.Cost * item.Number
			feeTotal += item.Denomination * 100 * item.Number
		} else {
			goodTicket = append(goodTicket, item)
		}
		bestValue += item.Denomination * 100
	}
	//合并相同券的数量,时间复杂度:O(n)
	bestTicket := make([]Ticket, 0)
	type tmpS struct {
		number int
		index  int
	}
	ticketMap := make(map[int]tmpS)
	for _, item := range goodTicket {
		if m, ok := ticketMap[item.Denomination]; ok {
			bestTicket[m.index].Number = m.number + item.Number
		} else {
			bestTicket = append(bestTicket, item)
			ticketMap[item.Denomination] = tmpS{
				number: item.Number,
				index:  len(bestTicket) - 1,
			}
		}
	}
	//打印结果
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
	fmt.Printf("实际抵扣费用:%d(分), 实际消耗币数量:%d(个)\n\n", feeTotal, coinTotal)
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

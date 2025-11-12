package main

import (
	"fmt"
	"time"
)

type canlook interface { //方法通知
	sendAlert(people string) error
}
type EmailNotifier struct{}

func (e EmailNotifier) sendAlert(people string) error {
	fmt.Printf("'%s'250,wow,哈哈哈\n", people)
	return nil
}

type SMSNotify struct{}

func (s SMSNotify) sendAlert(people string) error {
	fmt.Printf("好难呀，这怎么写的")
	return nil
}

type total struct {
	people   string
	sex      string
	age      int
	IQ       int
	normaliq int
}
type peoplemonitor struct {
	look []canlook //便于添加通知器
	to   []total   //更改数据
}

func (p *peoplemonitor) AddNotifier(n canlook) {
	p.look = append(p.look, n) //添加通知
}
func (p *peoplemonitor) AddpeopleMonitor(n total) {
	p.to = append(p.to, n) //添加人物
}
func (p *peoplemonitor) CheckIQ(people string) {
	for _, People := range p.to {
		if People.IQ < People.normaliq {
			fmt.Printf("'%v'智商太低\n", People.people)
		}
		if People.IQ > People.normaliq {
			fmt.Printf("'%s'是我男神\n", People.people)
		}
		if People.IQ == People.normaliq {
			fmt.Printf("'%s'是我女神\n", People.people)
		}
	}
}
func (p *peoplemonitor) sendAlerts(people string) {
	for _, Can := range p.look {
		err := Can.sendAlert(people)
		if err != nil {
			fmt.Printf("发送警报错误:%v\n", err)
		}
	}
}

func (p *peoplemonitor) Updatetotal(people string, newIQ int) {
	for i := range p.to {
		if p.to[i].people == people {
			oldiq := p.to[i].IQ
			p.to[i].IQ = newIQ
			fmt.Printf("'%s'智商已从'%d'是提升到'%d'\n", people, oldiq, newIQ)
			if newIQ < p.to[i].normaliq {
				p.sendAlerts(people)
			}
			if newIQ > p.to[i].normaliq {
				fmt.Printf("'%s'是我男神\n", people)
			}
			if newIQ == p.to[i].normaliq {
				fmt.Printf("'%s'是我女神\n", people)
			}
			break
		}
	}
}
func main() {
	fmt.Println("==智商测试启动==")
	emailNotifier := EmailNotifier{}
	smsNotifier := SMSNotify{}
	monitor := &peoplemonitor{}
	monitor.AddNotifier(emailNotifier)
	monitor.AddNotifier(smsNotifier)
	monitor.AddpeopleMonitor(total{
		people:   "唐晓丽",
		sex:      "small girl",
		age:      18,
		IQ:       249,
		normaliq: 250,
	})
	monitor.AddpeopleMonitor(total{
		people:   "朱一铭",
		sex:      "big boy",
		age:      18,
		IQ:       249,
		normaliq: 250,
	})
	fmt.Println("正在进行智商检查")
	monitor.CheckIQ("唐晓丽")
	fmt.Println()
	fmt.Println("智商重测中")
	monitor.Updatetotal("唐晓丽", 250)
	monitor.CheckIQ("朱一铭")
	monitor.Updatetotal("朱一铭", 128)
	time.Sleep(2 * time.Second)
	monitor.CheckIQ("朱一铭")
	monitor.Updatetotal("朱一铭", 251)
	time.Sleep(2 * time.Second)
	fmt.Println()
	monitor.CheckIQ("朱一铭")
	monitor.CheckIQ("唐晓丽")
	fmt.Println("智商测试完成")
	fmt.Println()

}

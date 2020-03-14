package main

import (
	"fmt"
)

//参考文章：https://mp.weixin.qq.com/s/MqQ6b-Z_wvYe9YpNI5LDeA

/*
开闭原则定义:
一个软件实体如类、模块和函数应该对扩展开放，对修改关闭。
简单的说就是在修改需求的时候，应该尽量通过扩展来实现变化，而不是通过修改已有代码来实现变化。
下面的例子： 因为银行业务员可能会新加很多个新的功能，新加的功能不能影响到已有的

另外一个：也可以基于抽象层进行业务封装-针对interface接口进行封装
func BankerBusiness(banker AbstractBanker) {
    //通过接口来向下调用，(多态现象)
    banker.DoBusi()
}

func main() {
    //进行存款
    BankerBusiness(&SaveBanker{})
    //进行存款
    BankerBusiness(&TransferBanker{})
    //进行存款
    BankerBusiness(&PayBanker{})
*/

//抽象的银行业务员
type AbstractBanker interface {
	DoBusi() //抽象的处理业务接口
}

//存款的业务员
type SaveBanker struct {
	//AbstractBanker
}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("进行了存款")
}

//转账的业务员
type TransferBanker struct {
	//AbstractBanker
}

func (tb *TransferBanker) DoBusi() {
	fmt.Println("进行了转账")
}

//支付的业务员
type PayBanker struct {
	//AbstractBanker
}

func (pb *PayBanker) DoBusi() {
	fmt.Println("进行了支付")
}

func main() {
	//进行存款
	sb := &SaveBanker{}
	sb.DoBusi()

	//进行转账
	tb := &TransferBanker{}
	tb.DoBusi()

	//进行支付
	pb := &PayBanker{}
	pb.DoBusi()

}

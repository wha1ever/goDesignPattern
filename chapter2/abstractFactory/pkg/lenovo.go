//++++++++++++++++++++++++++++++++++++++++
// 《Go语言设计模式》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 作者知乎：https://www.zhihu.com/people/shirdonl
// 仓库地址：https://gitee.com/shirdonl/goDesignPattern
// 仓库地址：https://github.com/shirdonl/goDesignPattern
// 交流咨询，请关注公众号"源码大数据"
//++++++++++++++++++++++++++++++++++++++++

package pkg

//联想品牌工厂
type Lenovo struct {
}

//生产手机
func (n *Lenovo) MakePhone() InterfacePhone {
	return &LenovoPhone{
		Phone: Phone{
			color: "Black",
			size:  5,
		},
	}
}

//生产电脑
func (n *Lenovo) MakeComputer() InterfaceComputer {
	return &LenovoComputer{
		Computer: Computer{
			color: "White",
			size:  14,
		},
	}
}

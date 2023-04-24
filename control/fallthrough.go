package main

import "fmt"

func main() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

// 每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。（ Go 语言使用快速的查找算法来测试 switch 条件与 case 分支的匹配情况，直到算法匹配到某个 case 或者进入 default 条件为止。）
//
// 一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块，也就是说您不需要特别使用 break 语句来表示结束。
//
// 因此，程序也不会自动地去执行下一个分支的代码。如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字来达到目的。
//
// ————————————————
// 原文作者：Go &#25216;&#26415;&#35770;&#22363;&#25991;&#26723;&#65306;&#12298;Go &#20837;&#38376;&#25351;&#21335;&#65288;&#65289;&#12299;
// 转自链接：https://learnku.com/docs/the-way-to-go/switch-structure/3594
// &#29256;&#26435;&#22768;&#26126;&#65306;&#32763;&#35793;&#25991;&#26723;&#33879;&#20316;&#26435;&#24402;&#35793;&#32773;&#21644; LearnKu &#31038;&#21306;&#25152;&#26377;&#12290;&#36716;&#36733;&#35831;&#20445;&#30041;&#21407;&#25991;&#38142;&#25509;

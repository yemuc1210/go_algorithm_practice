package main
import (
    "fmt"
    "bufio"
    "os"
//     "strings"
)
/*
坚持刷题可以获得牛币和实物勋章等多种奖励，点击此处立即进入活动页面

HJ2 计算某字符出现次数
较难  通过率：28.59%  时间限制：1秒  空间限制：32M
知识点	字符串	哈希
warning 校招时部分企业笔试将禁止编程题跳出页面，为提前适应，练习时请使用在线自测，而非本地IDE。
描述
写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字符，然后输出输入字符串中该字符的出现次数。（不区分大小写字母）

数据范围：  ，输入的数据有可能包含大小写字母、数字和空格
输入描述：
第一行输入一个由字母和数字以及空格组成的字符串，第二行输入一个字符。

输出描述：
输出输入字符串中含有该字符的个数。（不区分大小写字母）
*/

func read()(line string, ch byte){
    reader := bufio.NewReader(os.Stdin)
    //读取第一行
    var err error
    line,err = reader.ReadString('\n')
    if err != nil {
        line = ""
    }
    //如果是字母，转小写
    lineBytes := []byte(line)
    for i:=0;i<len(lineBytes);i++ {
        if lineBytes[i]>='A' && lineBytes[i]<='Z'{
            lineBytes[i] += 32
        }//数字和小写字母不变
    }
    line = string(lineBytes)
    // fmt.Println(line)
    err = nil
    ch,err = reader.ReadByte()
    if err != nil {
        ch = ' '
    }
    if ch>='A'&&ch<='Z' {
        ch +=32
    }
    return 
}

func main(){
    line,ch := read()
	// line :="655b203khDc79w76239SC60Yx5O634244O8oNKw2g5c122dt9M17197FQ9fR0fC8619q00cC156NJQ958ZAH8qe76G024Xvy74ETiO10WgQmOuNx104B9L8Euod682x6hVX941DUalUwb20nyD6IX6IM57Jvypu2mtZ8fpk9Z72AU9g0zob32W7J427VT50P7x40d6QF1067yYMioy6E037e2fTDTse6Z983873849KSB3t93t1wR14717iDi9ybw76h9410q345Z4882zWY3v7C0jD33NlLiF53Hp69067Dzvu1s4Er0G24MhcAHEq151Y339jyfbP6w9j1dX72DEv7L02764f2K7XVf7tWNYGqU5iNlS5RpV619vg7zR96D8jMsO5xN223W97CMuHHWBX806AD732q42h1bwpSb923V3Q73D3g3O0100tk948693d67nxrd7LB6He5O9qG2i9xJ0U1x9054y20039V9uD61HF9eP825JJ79M200X3zH1pW99QeG40kR4275p5P96S3J6mj60z278C1K685As6CX35rcY70qI5810i3rt5WB1C6H74118433Z8wsfE9s2M82q7TN7Q6Zz3047W99g56kLK58NqPfSuhSfEs0W4u1D720A3jqpG7aD0Be3N18395zS037T0Rg47K3O90H5w5660bNW58O8TK3l38xyN7R36kH4l21VYDx3waZK8kf5O3ySWJQdN7eVnf778GIa70Fan34v0owxkg4N5keNbaxh044l67Vimm110i7N8614Xj4667rQd6R67aHlm5z6N51jtbAm22305l9p83Uu6l8pBOaxQL56c5M76SUUKv9538FVD28WT1046sPQ1eHf5e65tk3Q6Lub314FBNv2C4Q12e6oPcX95jC06g8oFj7q7V20ciEO73m00U8v8342H1w86yS298r094x2y8O24e8p6j685D9a539930DQ6s6813uF8D34i6E1x5j2v9"
	// var ch byte
	// ch = '1'
	fmt.Println(ch)
    res := 0
    for i:=0;i<len(line);i++{
        if line[i] == ch {
            res ++
        }
    }
    fmt.Println(res)
}
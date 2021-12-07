package lt541

func reverseStr(s string, k int) string {
    //每隔k个翻转k个
    bt_arr := []byte(s)  //转字符数组
    for i:=0;i<len(bt_arr); i+= 2*k  {
        // sub := bt_arr[i: i+k]
        sub := bt_arr[i: min(i+k, len(bt_arr))]   //最后可能特殊
        //翻转sub
        n := len(sub)
        for j:=0;j<n/2;j++{
            sub[j],sub[n-1-j] = sub[n-1-j],sub[j]
        }
    }
    return string(bt_arr)

}

func min(a,b int)int{
    if a<b {
        return a
    }
    return b
}
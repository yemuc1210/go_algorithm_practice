package lt599

func findRestaurant(list1 []string, list2 []string) []string {
    m1 := make(map[string]int)
    m2 := make(map[string]int)
    n1,n2 := len(list1),len(list2)
    for k,v := range list1 {
        m1[v] = k+1
    }
    for k,v := range list2 {
        m2[v] = k+1
    }
    minIndexSum := n1+n2
    res := []string{}
    for _,v := range list1 {
        if m2[v] != 0 && m1[v] !=0{
            if m2[v]+m1[v] < minIndexSum {
                minIndexSum = m1[v] + m2[v]
            }
        }
    }
    for k,v := range m1 {
        if m2[k] != 0 {
            if m2[k]+v == minIndexSum {
                res = append(res, k)
            }
        }
    }
    return res
}
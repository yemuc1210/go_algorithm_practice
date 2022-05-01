package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func is(vv bool) bool {
    if vv == true{
        return false
    }
    if vv == false{
        return true
    }
    return false
}

func main(){
    s := bufio.NewScanner(os.Stdin)

    for s.Scan(){
        if s.Text() == ""{
            break
        }
        var enter bool
        lens := len(strings.Split(s.Text(), " "))
        
        for _,v := range strings.Split(s.Text(), " "){
            if strings.Contains(v, "\""){
                if strings.Count(v, "\"") == 1{
                    enter = is(enter)
                }
            }
            if !enter{
            }else{
                lens --
            }
        }
        fmt.Println(lens)
        
        for _,v := range strings.Split(s.Text(), " "){
            var value string
            value = v
            if strings.Contains(v, "\""){
                value = strings.Replace(v, "\"","",-1)
                
                 if strings.Count(v, "\"") == 1{
                    enter = is(enter)
                }
            }
            if !enter{
                fmt.Println(value)
            }else{
                fmt.Printf("%s ", value)
            }
            
        }
    }
    
}


import java.util.ArrayList;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String s = in.nextLine();
        // 判断是否可以变成模糊回文
        // 全部变成小写
        char[] bs= s.toLowerCase().toCharArray();
        // 非字母变成*
        for(int i=0;i<bs.length;i++) {
            if(bs[i]<'a' || bs[i]>'z'){
                bs[i]='*';                
            }
        }
        int[] count = new int[127];
        // 奇数位即可
        for(char c:bs){
            if(c=='*') {
                count[26]++;
            }else{
                count[c-'a']++;
            }
        }
        int even = 0;
        for(int i=0;i<27;i++){
            if(count[i]%2!=0){
                even++;
            }
        }
        System.out.println(even==0 || even==1);
        in.close();
    }
}
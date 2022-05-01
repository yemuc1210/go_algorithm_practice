import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

class Solution {
    public List<Integer> getRow(int rowIndex) {
        if(rowIndex==0) {
            List<Integer> a = Arrays.asList(1);
            return a;
        }
        // 计算只需要知晓前一行即可，但问题是每行长度不同
        // 频繁申请空间？
        // 每一行得长度与索引之间得关系： 1-1 2-2 i-i
        int[] a=new int[1],b=new int[2];
        for(int i=1;i<=rowIndex;i++){
            b = new int[i+1];
            // 计算新的一行
            // 首位是1
            b[0]=1;
            b[i]=1;
            for(int j=1;j<i;j++){
                b[j] = a[j-1]+a[j];
            }
            // b-a
            a = new int[i+1];
            for(int j=0;j<=i;j++){
                a[j] = b[j];
            }
        }
        return Arrays.stream(a).boxed().collect(Collectors.toList());
    }
}



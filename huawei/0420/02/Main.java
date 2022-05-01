import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        int[] root = {1,1,2,0,0,4,5};
        String path = "/1/1";
        int[] sub = {5,3,0};
        int[] res = solve(root, path, sub);
        in.close();
    }

    public static int[] solve(int[] root, String path, int[] sub ){

        // 1. 处理path
        String[] paths = path.split("/");
        // 转int
        int[] nums = new int[paths.length-1];
        for(int i=1;i<paths.length;i++){
            nums[i-1] =Integer.parseInt(paths[i]);
            System.out.println(nums[i-1]);
        }
        //2 遍历
        int i=0,j=0 ;
        int n=root.length;
        int m = sub.length;
        while(i<n && j < m) {
            if (root[i] == nums[j] ){
                if ((j+1) < m) {
                    j++;
                }else{
                    break;
                }
                if( root[2*i+1] == nums[j] ){
                    i = 2*i+1;
                }else if( root[2*i+2] == nums[j] ){
                    i = 2*i+2;
                }
            }
        }
        for(int k=0;k<n;k++){
            System.out.println(root[i]);
        }
        // 3. 清除 root 以i为根的部分
        clear(root, i, n);
        for(int k=0;k<n;k++){
            System.out.println(root[i]);
        }
        // 4. 对root扩容，因为sub.length > root[i].lengt 
        int[] tmp = new int[n];
        for(int k=0;k<n;k++){tmp[k] = root[k];}
        // 重新申请扩建
        root = new int[n+m];
        // copy回去
        for(int k=0;k<n;k++){
            root[i] = tmp[i];
        }
        return sub;
    }
    public static void clear(int[] root, int i, int n) {
        if(i<n) {
            root[i] = 0;
            clear(root, i*2+1, n);
            clear(root, i*2+2, n);
        }
    }
}
import java.util.Scanner;

public class Main{
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        // int n = in.nextInt();
        // // id 数组
        // int[] id = new int[n];
        // for(int i=0;i<n;i++){
        //     id[i] = in.nextInt();
        //     // System.out.println(id[i]);
        // }
        // // pid
        // int[] pid = new int[n];
        // for(int i=0;i<n;i++){
        //     pid[i] = in.nextInt();
        //     // System.out.println(pid[i]);
        // }
        // // task
        // int task = in.nextInt();
        // // System.out.println(task);

        int n=8;
        int[] id={16,12,3,10,5,2,4,1};
        int[] pid={0,16,16,12,3,3,2,2};
        int task=7;

        // 求解
        // 1 构造树
        // 假设输入比较规范 pid[0]=0
        int[] tree = new int[2*n];
        tree[0] = id[0];
  
        // 16 12 3 10 5 2 4 1   id   k
        // 0 16 16 12 3 3 2 2   pid  j 
        // 构建树
        int j = 1;
        int i=1;
        while(i<pid.length && j<pid.length){
            if(pid[j] == pid[i+1]) {
                tree[i] = id[j];
                tree[i+1] = id[j+1];
                j+=2;
            }else{
                tree[i] = id[j];
                tree[i+1] = 0;
                j++;
            }            
            i+=2;
        }
            

        for(int kk=0;kk<tree.length;kk++){
            System.out.print(tree[kk]+" ");
        }
        in.close();
    }
}
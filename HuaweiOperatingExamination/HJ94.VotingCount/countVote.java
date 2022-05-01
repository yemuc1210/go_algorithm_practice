import java.util.*;

public class countVote {
    public static void main(String[] args) {
        int n,m;   //n-候选人   m-投票人个数
        Scanner in = new Scanner(System.in);

        while(in.hasNextInt()) {
            n = in.nextInt();
            // System.out.println(n);
            in.nextLine();  //指针移动下一行开头
            // n个候选的名字
            String[] names = in.nextLine().split(" ");
            // System.out.println(names);
            // 输入投票人个数
            m = in.nextInt();
            // System.out.println(m);
            in.nextLine();
            //输入投票
            String[] votes = in.nextLine().split(" ");
            // System.out.println(votes);
            // 下面计算
            Map<String, Integer> dict = new HashMap<>();
            int inValid = 0;
            for(int i=0;i<n;i ++) {
                dict.put(names[i], 0);
            }
            // 计数
            for(int i=0;i<m;i++) {
                if(!dict.containsKey(votes[i])) {
                    // 非法投票
                    inValid ++;
                }else {
                    Integer curVote = dict.get(votes[i]);
                    dict.put(votes[i], curVote+1);
                }
            }
            // out
            StringBuilder res = new StringBuilder();
            for(String s: names) {
                res.append(s);
                res.append(" : ");
                res.append(dict.get(s));
                res.append("\n");
            }
            res.append("Invalid : ");
            res.append(inValid);
            System.out.println(res);
            
        }
        in.close();
    }
}
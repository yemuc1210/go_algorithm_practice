import java.util.*;
public class Solution {
    /**
     * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
     *
     * 返回最小的花费代价使得这n户人家连接起来
     * @param n int n户人家的村庄
     * @param m int m条路
     * @param cost int二维数组 一维3个参数，表示连接1个村庄到另外1个村庄的花费的代价
     * @return int
     */
    public int miniSpanningTree (int n, int m, int[][] cost) {
        // write code here
        HashSet<Integer> getpoint = new HashSet<>();
        //存储点
        ArrayList<int[]> list = new ArrayList<>();
        //存储边
        Arrays.sort(cost, (l1,l2)->{
            //排序
            return l1[2] - l2[2];
        });
        for (int[] data : cost) {//将数据添加到list中
            list.add(data);
        }
        int res = 0;
        res += cost[0][2];
        getpoint.add(cost[0][0]);
        getpoint.add(cost[0][1]);
        while (true) {
            for (int i = 1; i < list.size(); i++) {//下一条为权值最小的可连接的边
                if ((getpoint.contains(list.get(i)[0]) && !getpoint.contains(list.get(i)[1])) || (!getpoint.contains(list.get(i)[0]) && getpoint.contains(list.get(i)[1]))) {
                    res += list.get(i)[2];
                    getpoint.add(list.get(i)[0]);
                    getpoint.add(list.get(i)[1]);
                    list.remove(list.get(i));
                    //删除已经遍历过的这条边
                    break;
                }
            }
            if (getpoint.size() == n) {
                //所有边都遍历完成，结束
                break;
            }
        }
        return res;
    }
}
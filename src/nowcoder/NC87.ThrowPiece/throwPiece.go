package NC87


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 返回最差情况下扔棋子的最小次数
 * 一座大楼有 n+1 层，地面算作第0层，最高的一层为第 n 层。
 已知棋子从第0层掉落肯定不会摔碎，从第 i 层掉落可能会摔碎，也可能不会摔碎。
给定整数 n 作为楼层数，再给定整数 k 作为棋子数，
返回如果想找到棋子不会摔碎的最高层数，
即使在最差的情况下扔的最小次数。一次只能扔一个棋子。
 * @param n int整型 楼层数
 * @param k int整型 棋子数
 * @return int整型
*/

/*
copy

*/
func solve( n int ,  k int ) int {
    // write code here
	t := 1;
	for (calF(k, t) < n + 1) {
		t++
	}// 凑够n层的时候输出尝试的次数
	return t;
}

func calF(k,t int) int {
	if(t == 1 || k == 1) {
		return t + 1;
	} // 如果只有一次尝试机会或者棋子只有一个，则只能确定尝试机会+1数量的楼层哪一层棋子会碎
	return calF(k - 1, t - 1) + calF(k, t - 1);   // 非上述情况则递归（棋子-1，则机会-1）和（棋子没碎，机会-1）
}
// class Solution {
// 	public:
// 		/**
// 		 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
// 		 * 返回最差情况下扔棋子的最小次数
// 		 * @param n int整型 楼层数
// 		 * @param k int整型 棋子数
// 		 * @return int整型
// 		 */
// 		int calF(int k, int t) {
// 			if(t == 1 || k == 1) return t + 1;            // 如果只有一次尝试机会或者棋子只有一个，则只能确定尝试机会+1数量的楼层哪一层棋子会碎
// 			return calF(k - 1, t - 1) + calF(k, t - 1);   // 非上述情况则递归（棋子-1，则机会-1）和（棋子没碎，机会-1）
// 		}
	
// 		int solve(int n, int k) {
// 			// write code here
// 			int t = 1;
// 			while(calF(k, t) < n + 1) t++;                // 凑够n层的时候输出尝试的次数
// 			return t;
// 		}
// 	};

class Solution {
    int[][] dirs = { { -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 } };

    public boolean exist(char[][] board, String word) {
        // 搜索
        boolean res = false;
        int m = board.length, n = board[0].length;
        boolean[][] vis = new boolean[m][n];

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (dfs(i, j, 0, word, board, vis)) {
                    return true;
                }
            }
        }
        return res;
    }

    boolean dfs(int i, int j, int k, String target, char[][] board, boolean[][] visited) {
        if (board[i][j] != target.charAt(k)) {
            return false;
        }
        if (k == target.length() - 1) {
            return true;
        }
        visited[i][j] = true;

        for (int[] dir : dirs) {
            int x = i + dir[0], y = j + dir[1];
            if (x >= 0 && x < board.length && 0 <= y && y < board[0].length && !visited[x][y]) {
                // 递归
                if (dfs(x, y, k + 1, target, board, visited)) {
                    return true;
                }
            }
        }

        visited[i][j] = false;
        return false;
    }
}
class Solution {
    // 433. 最小基因变化 Java
    public int minMutation(String start, String end, String[] bank) {
        // 根据start end筛选出可选的基因序列
        Set<String> set = new HashSet<String>();
        set.add(start);
        for(String str:bank){
            set.add(str);
        }
        int cnt=0;
        Queue<String> queue = new LinkedList<>();
        queue.offer(start);
        set.remove(start);
        while(!queue.isEmpty()){
            int curWidth = queue.size();
            for(int i=0;i<curWidth;i++){
                String top = queue.poll();
                if(top.equals(end)) {
                    return cnt;
                }
                for(String str : bank) {
                    if(isMutation(str, top) && set.contains(str)) {
                        queue.offer(str);
                        set.remove(str);
                    }
                }
            }
            cnt++;
        }
        return -1;
    }
    private boolean isMutation(String mutation, String old) {
        boolean f = false;
        for(int i = 0; i < old.length(); i++) {
            if(mutation.charAt(i) != old.charAt(i)) {
                if(f) return false;
                f = true;
            }
        }
        return f;
    }
}
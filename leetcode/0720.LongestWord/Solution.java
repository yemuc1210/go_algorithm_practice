import java.util.HashMap;

class Solution {
    public String longestWord(String[] words) {
        HashMap<String,Boolean> map = new HashMap<>();
        String res = "";
        for(String word:words){
            map.put(word, true);
        }
        outer:
        for(String word:words) {
            for(int i=1;i<word.length();i++){
                if(!map.getOrDefault(word.substring(0, i), false)){
                    continue outer;
                }
            }
            // 那么，word是一个候选
            int len1=word.length(),len2=res.length();
            if( len1>len2 || (len1==len2 && word.compareTo(res)<0) ){
                res = word;
            }
        }
        return res;
    }
}

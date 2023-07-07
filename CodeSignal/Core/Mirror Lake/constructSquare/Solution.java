public class Solution {

  int[] pattern(String s,int min){
    int[] ans=new int[26];
    for(int i=0;i<s.length();i++) ans[s.charAt(i)-min]++;
    Arrays.sort(ans);
    return ans;
  }

  int solution(String s) {
    int[] pattern=pattern(s,'a');
    int ans=-1;

    for(int i=3;true;i++){
      int x=i*i;
      String ss=""+x;
      if(ss.length()>s.length()) return ans;
      if(ss.length()==s.length()){
        int[] p=pattern(ss,'0');
        boolean good=true;
        for(int j=0;j<26;j++) if(p[j]!=pattern[j]) good=false;
        if(good) ans=x;
      }
    }
  }

}

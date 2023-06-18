public class Solution {

  int solution(int a0) {
    HashSet<Integer> seen=new HashSet<Integer>();
    int it=0;
    int curr=a0;
    while (!seen.contains(curr)) {
      System.out.println(curr);
      it++;
      seen.add(curr);
      int next=0;
      while (curr>0) {
        int digit=curr%10;
        curr/=10;
        next+=digit*digit;
      }
      curr=next;
    }
    return it+1;
  }

}

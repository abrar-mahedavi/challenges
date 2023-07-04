public class Solution {

  int solution(String a, String b) {
    int[] arr1 = new int[26];
    int[] arr2 = new int[26];
    for(int i = 0; i< a.length();i++) {
      arr1[a.charAt(i)-97]++;
    }
    for(int i = 0; i< b.length();i++) {
      arr2[b.charAt(i)-97]++;
    }
    int min = b.length();
    for(int i=0; i<26;i++) {
      if(arr1[i] !=0) {
        arr1[i] = arr2[i]/arr1[i];
        if(arr1[i]<min) min = arr1[i];
      }
    }
    return min;
  }


}

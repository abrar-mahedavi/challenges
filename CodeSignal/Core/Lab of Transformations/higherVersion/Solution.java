public class Solution {

  boolean solution(String ver1, String ver2) {
    int[] v1 = Arrays.stream(ver1.split("\\.")).mapToInt(i->Integer.parseInt(i)).toArray();
    int[] v2 = Arrays.stream(ver2.split("\\.")).mapToInt(i->Integer.parseInt(i)).toArray();
    int i = 0;
    while (i<v1.length-1&&v1[i]==v2[i]) {i++;}
    return v1[i]>v2[i];
  }
}

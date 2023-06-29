public class Solution {

  String solution(String address) {
    return address.substring(address.lastIndexOf("@") + 1);
  }
}

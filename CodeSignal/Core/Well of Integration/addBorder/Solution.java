public class Solution {

  String[] solution(String[] picture) {
    String[] borderArray = new String[picture.length + 2];
    borderArray[0] = "*";
    for(int i = 0; i < picture[0].length() + 1; i ++)
    {
      borderArray[0] += "*";
    }
    for(int j = 0; j < picture.length; j ++)
    {
      borderArray[j + 1] = "*" + picture[j] + "*";
    }
    borderArray[picture.length + 1] = "*";
    for(int x = 0; x < picture[0].length() + 1; x ++)
    {
      borderArray[picture.length + 1] += "*";
    }
    return borderArray;
  }


}

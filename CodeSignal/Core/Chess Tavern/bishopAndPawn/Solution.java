public class Solution {

  boolean solution(String bishop, String pawn) {
    Set<String> squares = new HashSet<>();
    char startRow = bishop.charAt(1);
    char startCol = bishop.charAt(0);
    for (int i = 0; i < 8; i++) {
      char rowUp = (char) (startRow + i);
      char rowDown = (char) (startRow - i);
      char colLeft = (char) (startCol - i);
      char colRight = (char) (startCol + i);
      squares.add("" + colLeft + rowUp);
      squares.add("" + colLeft + rowDown);
      squares.add("" + colRight + rowUp);
      squares.add("" + colRight + rowDown);
    }

    return squares.contains(pawn);
  }

}

public class Solution {

  int[] solution(String king, String amazon) {
    char[] k = king.toCharArray();
    int k1 = (int) (k[1] - '0') - 1;
    int k2 = (int) (k[0] - 'a');

    char[] a = amazon.toCharArray();
    int a1 = (int) (a[1] - '0') - 1;
    int a2 = (int) (a[0] - 'a');

    ArrayList<TwoIntTuple> check = new ArrayList<>();
    ArrayList<TwoIntTuple> checkmate = new ArrayList<>();
    ArrayList<TwoIntTuple> stalemate = new ArrayList<>();
    ArrayList<TwoIntTuple> safe = new ArrayList<>();

    for (int row = 0; row < 8; row++) {
      for (int column = 0; column < 8; column++) {
        if (InCheck(row, column, k1, k2, a1, a2)) {
          check.add(CreateArray(row, column));
        } else {
          if (Math.abs(row - k1) > 1 || Math.abs(column - k2) > 1) {
            safe.add(CreateArray(row, column));
          }
        }
      }
    }
    // check is checkmate
    for (TwoIntTuple check1 : check) {
      int row = check1.Integer1();
      int column = check1.Integer2();
      boolean isCheckMate = true;
      if ((safe.contains(CreateArray(row - 1, column))) ||
        (safe.contains(CreateArray(row + 1, column))) ||
        (safe.contains(CreateArray(row, column - 1))) ||
        (safe.contains(CreateArray(row, column + 1))) ||
        (safe.contains(CreateArray(row - 1, column - 1))) ||
        (safe.contains(CreateArray(row - 1, column + 1))) ||
        (safe.contains(CreateArray(row + 1, column - 1))) ||
        (safe.contains(CreateArray(row + 1, column + 1)))) {
        isCheckMate = false;
      }
      if (isCheckMate) {
        checkmate.add(check1);
      }
    }
    check.removeAll(checkmate);
    safe.remove(CreateArray(a1, a2));

    // safe has a safe neighbor
    for (TwoIntTuple safe1 : safe) {
      int row = safe1.Integer1();
      int column = safe1.Integer2();
      boolean hasASafeNeighbor = false;
      if ((safe.contains(CreateArray(row - 1, column))) ||
        (safe.contains(CreateArray(row + 1, column))) ||
        (safe.contains(CreateArray(row, column - 1))) ||
        (safe.contains(CreateArray(row, column + 1))) ||
        (safe.contains(CreateArray(row - 1, column - 1))) ||
        (safe.contains(CreateArray(row - 1, column + 1))) ||
        (safe.contains(CreateArray(row + 1, column - 1))) ||
        (safe.contains(CreateArray(row + 1, column + 1)))) {
        hasASafeNeighbor = true;
      }
      if (!hasASafeNeighbor) {
        stalemate.add(safe1);
      }
    }
    safe.removeAll(stalemate);
    int[] answer = new int[4];
    answer[0] = checkmate.size();
    answer[1] = check.size();
    answer[2] = stalemate.size();
    answer[3] = safe.size();
    return answer;
  }

  boolean InCheck(int row, int column, int k1, int k2, int a1, int a2) {
    // Too close to the king
    if (Math.abs(row - k1) < 2 && Math.abs(column - k2) < 2) {
      return false;
    }
    // On the Amazon
    if (row == a1 && column == a2) {
      return false;
    }
    // Too close to the Amazon
    if (Math.abs(row - a1) < 3 && Math.abs(column - a2) < 3) {
      return true;
    }
    // On same row
    if (row == a1) {
      // King in between on same row
      if (row == k1) {
        if ((column < k2 && k2 < a2) ||
          (column > k2 && k2 > a2)) {
          return false;
        }
      }
      return true;
    }
    // On same column
    if (column == a2) {
      // King in between on same column
      if (column == k2) {
        if ((row < k1 && k1 < a1) ||
          (row > k1 && k1 > a1)) {
          return false;
        }
      }
      return true;
    }
    // On same SW-NE diagonal
    if (row - a1 == column - a2) {
      // King in between on same diagonal
      if (row - k1 == column - k2) {
        if ((column < k2 && k2 < a2) ||
          (column > k2 && k2 > a2)) {
          return false;
        }
      }
      return true;
    }
    // On same NW-SE diagonal
    if (row - a1 == a2 - column) {
      // King in between on same diagonal
      if (row - k1 == k2 - column) {
        if ((column < k2 && k2 < a2) ||
          (column > k2 && k2 > a2)) {
          return false;
        }
      }
      return true;
    }
    return false;
  }

  TwoIntTuple CreateArray(int row, int column) {
    //   return new Integer[]{row, column};
    return new TwoIntTuple(row, column);
  }

  class TwoIntTuple {

    private final int int1;
    private final int int2;

    public TwoIntTuple(int int1, int int2) {
      this.int1 = int1;
      this.int2 = int2;
    }

    @Override
    public int hashCode() {
      int hash = 5;
      hash = 57 * hash + this.int1;
      hash = 57 * hash + this.int2;
      return hash;
    }

    @Override
    public boolean equals(Object object) {
      if (object != null && object instanceof TwoIntTuple) {
        TwoIntTuple tuple = (TwoIntTuple) object;
        return (int1 == tuple.int1) && (int2 == tuple.int2);
      }
      return false;
    }

    public int Integer1() {
      return int1;
    }

    public int Integer2() {
      return int2;
    }
  }


}

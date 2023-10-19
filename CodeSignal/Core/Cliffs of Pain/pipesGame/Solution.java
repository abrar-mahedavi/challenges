public class Solution {

  int solution(String[] state) {
    ArrayList<ArrayList<String>> st = new ArrayList<>();
    boolean ch = true;
    int min = state.length * state[0].length();
    for (int y = 0; y < state.length; y++) {
      for (int x = 0; x < state[y].length(); x++) {
        if (state[y].charAt(x) >= 'a' && state[y].charAt(x) <= 'z') {
          if (x < state[y].length() - 1) {
            st.add(0, pipes(state, y, x, 0, 1));
            if (st.get(0).get(0).equals("false")) {
              ch = false;
              min = st.get(0).size() < min && st.get(0).size() != 1 ?
                st.get(0).size() : min;
            }
          }
          if (y < state.length - 1) {
            st.add(0, pipes(state, y, x, 1, 0));
            if (st.get(0).get(0).equals("false")) {
              ch = false;
              min = st.get(0).size() < min && st.get(0).size() != 1 ?
                st.get(0).size() : min;
            }
          }
          if (x != 0) {
            st.add(0, pipes(state, y, x, 0, -1));
            if (st.get(0).get(0).equals("false")) {
              ch = false;
              min = st.get(0).size() < min && st.get(0).size() != 1 ?
                st.get(0).size() : min;
            }
          }
          if (y != 0) {
            st.add(0, pipes(state, y, x, -1, 0));
            if (st.get(0).get(0).equals("false")) {
              ch = false;
              min = st.get(0).size() < min && st.get(0).size() != 1 ?
                st.get(0).size() : min;
            }
          }
        }
      }

    }

    return check(st, min, ch);
  }

  int check(ArrayList<ArrayList<String>> st, int min, boolean ch) {
    ArrayList full = new ArrayList();
    for (int i = 0; i < st.size(); i++) {
      if (st.get(i).size() > 1) {
        int tmp = st.get(i).size() < min ? st.get(i).size() : min;
        full.addAll(st.get(i).subList(1, ch == true ? st.get(i).size() : tmp));
      }
    }
    return repeat(full, ch, min);
  }

  int repeat(ArrayList full, boolean ch, int min) {
    String[] arr = new String[full.size()];

    int count = 1;
    for (int i = 0; i < arr.length; i++) {
      arr[i] = full.get(i).toString();
    }
    Arrays.sort(arr);
    for (int i = 1; i < arr.length; i++) {
      if (arr[i].equals(arr[i - 1]) || arr[i].equals("")) {
        continue;
      }
      count++;
    }
    count = arr.length == 0 ? 0 : count;
    return ch == true ? count : -count;
  }

  ArrayList pipes(String[] state, int y, int x, int indY, int indX) {
    ArrayList pipes = new ArrayList<>();
    pipes.add(0, "true");
    char s = (char) (state[y].charAt(x) - 32);
    while (true) {
      if (x + indX >= state[y].length() || x + indX < 0 ||
        y + indY >= state.length || y + indY < 0) {
        if (pipes.size() != 1) {
          pipes.set(0, "false");
        }

        return pipes;
      }
      x += indX;
      y += indY;
      char check = state[y].charAt(x);
      switch (check) {
        case '0':
          if (pipes.size() != 1) {
            pipes.set(0, "false");
          }
          return pipes;
        case '1':
          if (indX != 0) {
            if (pipes.size() != 1) {
              pipes.set(0, "false");
            }
            return pipes;
          }
          break;
        case '2':
          if (indY != 0) {
            if (pipes.size() != 1) {
              pipes.set(0, "false");
            }
            return pipes;
          }
          break;
        case '3':
          if (indY > 0 || indX > 0) {
            if (pipes.size() != 1) {
              pipes.set(0, "false");
            }
            return pipes;
          }
          if (indY < 0) {
            indY = 0;
            indX = 1;
          } else {
            indY = 1;
            indX = 0;
          }
          break;
        case '4':
          if (indY > 0 || indX < 0) {
            if (pipes.size() != 1) {
              pipes.set(0, "false");
            }
            return pipes;
          }
          if (indY < 0) {
            indY = 0;
            indX = -1;
          } else {
            indY = 1;
            indX = 0;
          }
          break;
        case '5':
          if (indY < 0 || indX < 0) {
            if (pipes.size() != 1) {
              pipes.set(0, "false");
            }
            return pipes;
          }
          if (indY > 0) {
            indY = 0;
            indX = -1;
          } else {
            indY = -1;
            indX = 0;
          }
          break;
        case '6':
          if (indY < 0 || indX > 0) {
            if (pipes.size() != 1) {
              pipes.set(0, "false");
            }
            return pipes;
          }
          if (indY > 0) {
            indY = 0;
            indX = 1;
          } else {
            indY = -1;
            indX = 0;
          }
          break;
        case '7':
          break;
        default:
          if (state[y].charAt(x) == s) {
            return pipes;
          }
          if (state[y].charAt(x) >= 'a' && state[y].charAt(x) <= 'z') {
            if (pipes.size() != 1) {
              pipes.set(0, "false");
            }
          }
          return pipes;
      }
      pipes.add(y + "-" + x);
    }
  }

}

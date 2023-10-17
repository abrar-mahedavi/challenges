public class Solution {

  String solution(String[] program) {
    char[][] instructions = new char[program.length][];
    for (int i = 0; i < program.length; i++) {
      instructions[i] = program[i].toCharArray();
    }
    int x = 0;
    int y = 0;
    int xInc = 0;
    int yInc = 1;
    StringBuilder result = new StringBuilder();
    Stack<Integer> st = new Stack<Integer>();
    int count = 0;
    boolean stringMode = false;
    while ((stringMode || instructions[x][y] != '@') && count < 10000 && result.length() < 100) {
      if (stringMode) {
        if (instructions[x][y] == '"') {
          stringMode = false;
        } else {
          st.push((int) instructions[x][y]);
        }
      } else {
        switch (instructions[x][y]) {
          case '>':
            xInc = 0;
            yInc = 1;
            break;
          case '<':
            xInc = 0;
            yInc = -1;
            break;
          case 'v':
            xInc = 1;
            yInc = 0;
            break;
          case '^':
            xInc = -1;
            yInc = 0;
            break;
          case '#':
            x += xInc;
            y += yInc;
            if (x >= instructions.length) {
              x = 0;
            } else if (x < 0) {
              x = instructions.length - 1;
            }
            if (y >= instructions[0].length) {
              y = 0;
            } else if (y < 0) {
              y = instructions[0].length - 1;
            }
            break;
          case '_':
            xInc = 0;
            if (pop(st) == 0) {
              yInc = 1;
            } else {
              yInc = -1;
            }
            break;
          case '|':
            yInc = 0;
            if (pop(st) == 0) {
              xInc = 1;
            } else {
              xInc = -1;
            }
            break;
          case '+':
            st.push(pop(st) + pop(st));
            break;
          case '-':
            st.push(-pop(st) + pop(st));
            break;
          case '*':
            int a = pop(st);
            int b = pop(st);
            st.push(a * b);
            break;
          case '/':
            a = pop(st);
            b = pop(st);
            st.push(b / a);
            break;
          case '%':
            a = pop(st);
            b = pop(st);
            st.push(b % a);
            break;
          case '!':
            a = pop(st);
            st.push(a == 0 ? 1 : 0);
            break;
          case '`':
            a = pop(st);
            b = pop(st);
            st.push(b > a ? 1 : 0);
            break;
          case ':':
            if (!st.isEmpty()) {
              st.push(st.peek());
            }
            break;
          case '\\':
            a = pop(st);
            b = pop(st);
            st.push(a);
            st.push(b);
            break;
          case '$':
            pop(st);
            break;
          case '.':
            a = pop(st);
            result.append(a + " ");
            break;
          case ',':
            result.append((char) pop(st));
            break;
          case '0':
          case '1':
          case '2':
          case '3':
          case '4':
          case '5':
          case '6':
          case '7':
          case '8':
          case '9':
            st.push(instructions[x][y] - '0');
            break;
          case '"':
            stringMode = true;
            break;
        }
      }
      x += xInc;
      y += yInc;
      if (x >= instructions.length) {
        x = 0;
      } else if (x < 0) {
        x = instructions.length - 1;
      }
      if (y >= instructions[0].length) {
        y = 0;
      } else if (y < 0) {
        y = instructions[0].length - 1;
      }
      count++;
    }
    return result.toString();
  }

  int pop(Stack<Integer> st) {
    if (st.isEmpty()) {
      return 0;
    } else {
      return st.pop();
    }
  }

}

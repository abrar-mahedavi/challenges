public class Solution {

  int solution(String[] words) {
    int t = 0;
    for (int i = 0; i < words.length; i++)
      for (int j = 0; j < words.length; j++)
        for (int k = 0; k < words.length; k++)
          for (int l = 0; l < words.length; l++)
            if (i != j && i != k && i != l &&
              j != k && j != l && k != l)
              t+=check(words[i],words[j],words[k],words[l]);
    return t;
  }

  int check (String a, String b, String c, String d) {
    int t = 0;
    for (int a1 = 0; a1 < a.length(); a1++)
      for (int a2 = a1+2; a2 < a.length(); a2++)
        for (int b1 = 0; b1 < b.length(); b1++)
          for (int b2 = b1+2; b2 < b.length(); b2++)
            for (int c1 = 0; c1 < c.length(); c1++)
              for (int d1 = 0; d1 < d.length(); d1++) {
                int c2 = c1 + (a2-a1);
                int d2 = d1 + (b2-b1);
                if (c2 < c.length() && d2 < d.length()) {
                  if (a.charAt(a1)==b.charAt(b1)
                    && a.charAt(a2)==d.charAt(d1)
                    && c.charAt(c1)==b.charAt(b2)
                    && c.charAt(c2)==d.charAt(d2))
                    t++;
                }
              }
    return t;
  }

}

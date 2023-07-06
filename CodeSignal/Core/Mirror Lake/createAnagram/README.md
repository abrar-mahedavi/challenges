`Medium`	`Codewriting` 	`300`

## Description

---

You are given two strings <code>s</code> and <code>t</code> of the same length, consisting of uppercase English letters. Your task is to find the minimum number of "_replacement operations_" needed to get some anagram (a string <code>x</code> is an **anagram** of a string <code>y</code> if one can get <code>y</code> by rearranging the letters of <code>x</code>. For example, the strings <code>"MITE"</code> and <code>"TIME"</code> are anagrams, so are <code>"BABA"</code> and <code>"AABB"</code>, but <code>"ABBAC"</code> and <code>"CAABA"</code> are not) of the string <code>t</code> from the string <code>s</code>. A replacement operation is performed by picking exactly one character from the string <code>s</code> and replacing it by some other character.

**Example**

- For <code>s = "AABAA"</code> and <code>t = "BBAAA"</code>, the output should be
  <code>createAnagram(s, t) = 1</code>;

- For <code>s = "OVGHK"</code> and <code>t = "RPGUC"</code>, the output should be
  <code>createAnagram(s, t) = 4</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] string s**

  _Guaranteed constraints:_<br>
  <code>5 ≤ s.length ≤ 35</code>.

- **[input] string t**

  _Guaranteed constraints:_<br>
  <code>t.length = s.length</code>.

- **[output] integer**
  - The minimum number of replacement operations needed to get an anagram of the string <code>t</code> from the string <code>s</code>.

## [Java] Syntax Tips

``` java
// Prints help message to the console
// Returns a string
// 
// Globals declared here will cause a compilation error,
// declare variables inside the function instead!
String helloWorld(String name) {
    System.out.println("This prints to the console when you Run Tests");
    return "Hello, " + name;
}
```

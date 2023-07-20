`Medium`	`Codewriting` 	`300`

## Description

---

Timed Reading is an educational tool used in many schools to improve and advance reading skills. A young elementary student has just finished his very first timed reading exercise. Unfortunately he's not a very good reader yet, so whenever he encountered a word longer than <code>maxLength</code>, he simply skipped it and read on.

Help the teacher figure out how many words the boy has read by calculating the number of words in the <code>text</code> he has read, no longer than <code>maxLength</code>.
Formally, a word is a substring consisting of English letters, such that characters to the left of the leftmost letter and to the right of the rightmost letter are not letters.

**Example**

For <code>maxLength = 4</code> and
<code>text = "The Fox asked the stork, 'How is the soup?'"</code>,<br>the output should be
<code>timedReading(maxLength, text) = 7</code>.

The boy has read the following words: <code>"The", "Fox", "the", "How", "is", "the", "soup"</code>.

</code>

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] integer maxLength**

  A positive integer, the maximum length of the word the boy can read.<br>

  _Guaranteed constraints:_<br>
  <code>1 ≤ maxLength ≤ 10</code>.

- **[input] string text**

  A non-empty string of English letters and punctuation marks.<br>

  _Guaranteed constraints:_<br>
  <code>3 ≤ text.length ≤ 110</code>.

- **[output] integer**
  - The number of words the boy has read.

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

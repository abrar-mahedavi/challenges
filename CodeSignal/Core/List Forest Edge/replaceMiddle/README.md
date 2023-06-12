`Medium`	`Codewriting` 	`300`

## Description

---

We define the _middle_ of the array <code>arr</code> as follows:

- if <code>arr</code> contains an odd number of elements, its _middle_ is the element whose index number is the same when counting from the beginning of the array and from its end;
- if <code>arr</code> contains an even number of elements, its _middle_ is the sum of the two elements whose index numbers when counting from the beginning and from the end of the array differ by one.

Given array <code>arr</code>, your task is to find its _middle_, and, if it consists of two elements, replace those elements with the value of _middle_. Return the resulting array as the answer.

**Example**

- For <code>arr = [7, 2, 2, 5, 10, 7]</code>, the output should be
  <code>replaceMiddle(arr) = [7, 2, 7, 10, 7].</code>

  The _middle_ consists of two elements, <code>2</code> and 5. These two elements should be replaced with their sum, i.e. <code>7</code>.

- For <code>arr = [-5, -5, 10]</code>, the output should be
  <code>replaceMiddle(arr) = [-5, -5, 10]</code>.

  The middle is defined as a single element <code>-5</code>, so the initial array with no changes should be returned.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.integer arr**

  The given array.<br>

  _Guaranteed constraints:_<br>
  <code>2 ≤ arr.length ≤ 10<sup>4</sup></code>,<br>
  <code>-10<sup>3</sup> ≤ arr[i] ≤ 10<sup>3</sup></code>.

- **[output] array.integer**

  - <code>arr</code> with its _middle_ replaced by a single element, or the initial array if the _middle_ consisted of a single element to begin with.

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

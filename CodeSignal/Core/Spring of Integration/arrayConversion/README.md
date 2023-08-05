`Medium`	`Codewriting` 	`300`

## Description

---

Given an array of <code>2<sup>k</sup></code> integers (for some integer <code>k</code>), perform the following operations until the array contains only one element:

- On the <code>1<sup>st</sup></code>, <code>3<sup>rd</sup></code>, <code>5<sup>th</sup></code>, etc. iterations (1-based) replace each pair of consecutive elements with their sum;
- On the <code>2<sup>nd</sup></code>, <code>4<sup>th</sup></code>, 6th, etc. iterations replace each pair of consecutive elements with their product.

After the algorithm has finished, there will be a single element left in the array. Return that element.

**Example**

For <code>inputArray = [1, 2, 3, 4, 5, 6, 7, 8]</code>, the output should be
<code>arrayConversion(inputArray) = 186</code>.

We have <code>[1, 2, 3, 4, 5, 6, 7, 8] -> [3, 7, 11, 15] -> [21, 165] -> [186]</code>, so the answer is <code>186</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.integer inputArray**

  _Guaranteed constraints:_<br>
  <code>1 ≤ inputArray.length ≤ 2<sup>7</sup></code>,<br> <code>-100 ≤ inputArray[i] ≤ 100</code>.

- **[output] integer**

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

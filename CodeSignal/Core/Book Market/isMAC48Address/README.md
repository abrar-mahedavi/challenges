`Medium`	`Codewriting` 	`300`

## Description

---

A media access control address (MAC address) is a unique identifier assigned to network interfaces for communications on the physical network segment.

The standard (IEEE 802) format for printing MAC-48 addresses in human-friendly form is six groups of two hexadecimal digits (<code>0</code> to <code>9</code> or <code>A</code> to <code>F</code>), separated by hyphens (e.g. <code>01-23-45-67-89-AB</code>).

Your task is to check by given string <code>inputString</code> whether it corresponds to MAC-48 address or not.

**Example**

- For <code>inputString = "00-1B-63-84-45-E6"</code>, the output should be
  <code>isMAC48Address(inputString) = true</code>;
- For <code>inputString = "Z1-1B-63-84-45-E6"</code>, the output should be
  <code>isMAC48Address(inputString) = false</code>;
- For <code>inputString = "not a MAC-48 address"</code>, the output should be
  <code>isMAC48Address(inputString) = false</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] string inputString**

  _Guaranteed constraints:_<br>
  <code>15 ≤ inputString.length ≤ 20</code>.

- **[output] boolean**
  - <code>true</code> if <code>inputString</code> corresponds to MAC-48 address naming rules, <code>false</code> otherwise.

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

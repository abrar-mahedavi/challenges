`Medium`	`Codewriting` 	`300`

## Description

---

Given two version strings composed of several non-negative decimal fields separated by periods (<code>"."</code>), both strings contain equal number of numeric fields. Return <code>true</code> if the first version is higher than the second version and <code>false</code> otherwise.

The syntax follows the regular semver ordering rules:

<code>
1.0.5 < 1.1.0 < 1.1.5 < 1.1.10 < 1.2.0 < 1.2.2
< 1.2.10 < 1.10.2 < 2.0.0 < 10.0.0
</code>

There are no leading zeros in any of the numeric fields, i.e. you do not have to handle inputs like <code>100.020.003</code> (it would instead be given as <code>100.20.3</code>).

**Example**

- For <code>ver1 = "1.2.2"</code> and <code>ver2 = "1.2.0"</code>, the output should be
  <code>higherVersion(ver1, ver2) = true</code>;
- For <code>ver1 = "1.0.5"</code> and <code>ver2 = "1.1.0"</code>, the output should be
  <code>higherVersion(ver1, ver2) = false</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] string ver1**

  _Guaranteed constraints:_<br>
  <code>1 ≤ ver1.length ≤ 15</code>.

- **[input] string ver2**

  _Guaranteed constraints:_<br>
  <code>1 ≤ ver2.length ≤ 15</code>.

- **[output] boolean**

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

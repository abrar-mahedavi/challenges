`Medium`	`Codewriting` 	`300`

## Description

---

An email address such as <code>"John.Smith@example.com"</code> is made up of a local part (<code>"John.Smith"</code>), an <code>"@"</code> symbol, then a domain part (<code>"example.com"</code>).

The domain name part of an email address may only consist of letters, digits, hyphens and dots. The local part, however, also allows a lot of different special characters. [Here](https://en.wikipedia.org/wiki/Email_address#Examples) you can look at several examples of correct and incorrect email addresses.

Given a valid email address, find its domain part.

**Example**

- For <code>address = "prettyandsimple@example.com"</code>, the output should be
  <code>findEmailDomain(address) = "example.com"</code>;
- For <code>address = "fully-qualified-domain@codesignal.com"</code>, the output should be
  <code>findEmailDomain(address) = "codesignal.com"</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] string address**

  _Guaranteed constraints:_<br>
  <code>10 ≤ address.length ≤ 50</code>.

- **[output] string**
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

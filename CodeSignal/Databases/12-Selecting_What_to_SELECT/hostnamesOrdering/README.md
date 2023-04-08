`Medium`	`Codewriting` 	`300`

Implement the missing code, denoted by ellipses. You may not modify the pre-existing code.
Yesterday you wrote down the links for some resources that you are going to use as references in your academic paper, and now you want to sort them. You didn't bother to write down the entire link for any of them, so all you have is a bunch of unique hostnames.

You stored the information about these hostnames in the table **hostnames**, which has the structure:

- **id**: the unique hostname id;
- **hostname**: the unique hostname.

Each hostname contain several domains and can be written in the format <code>hostname = domain1.domain2. ...</code>. To sort the hostnames, you've decided that hostname <code>A</code> should go before hostname <code>B</code> in the sorted list if its reversed list of domains is [lexicographically smaller](keyword://lexicographical-order-for-arrays) than the reversed list of <code>B</code>'s domains.

Given the table **hostnames**, build the resulting table as follows: The table should have columns <code>id</code> and <code>hostname</code>, and its values should be ordered as described above. It is guaranteed that all the hostnames are different and that there are no more than <code>3</code> domains in each hostname.

## Example

- For the following tables **hostnames**

| id  | hostname             |
| --- | -------------------- |
| 1   | workbench.mysql.com  |
| 2   | codesignal.slack.com |
| 3   | codesignal.com       |
| 4   | snarknews.info       |
| 5   | sololearn.com        |
| 6   | dev.mysql.com        |

the output should be

| id  | hostname             |
| --- | -------------------- |
| 3   | codesignal.com       |
| 6   | dev.mysql.com        |
| 1   | workbench.mysql.com  |
| 2   | codesignal.slack.com |
| 5   | sololearn.com        |
| 4   | snarknews.info       |

# Output
- [execution time limit] 10 seconds (mysql)


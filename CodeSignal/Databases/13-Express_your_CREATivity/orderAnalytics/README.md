`Medium`	`Codewriting` 	`300`

You're working in a big delivery company that has to handle a lot of orders.

The information about all orders is stored in the table **orders**. Here is its structure:

- **id**: unique order id;
- **order_date**: order date (having DATE type) in the format YYYY-MM-DD;
- **type**: the type of the product in the order;
- **quantity**: the quantity of the product in the order;
- **price**: the price of one item of the product.

In order to get a better view of your data you decided to create a new table called **orders_analytics**. It should contain the same orders as the **orders** and consist of the following columns:

- **id**: order id;
- **year**: the year of the order;
- **quarter**: the quarter of the year in which the order was taken;
- **type**: the type of the product in the order;
- **total_price**: the total price of the order (<code>total_price = price \* quantity</code>).

The _quarter_ of the year is defined as follows:

- 01 January till 31 March refers to the <code>1<sup>st</sup></code> quarter;
- 01 April till 30 June refers to the <code>2<sup>nd</sup></code> quarter;
- 01 July till 30 September refers to the <code>3<sup>rd</sup></code> quarter;
- 01 October till 31 December refers to the <code>4<sup>th</sup></code> quarter.

Given the table **orders**, build the new table **order_analytics** and print all its rows ordered by order <code>id</code>s.

**Example**

For the following tables **orders**

| id  | order_date | type  | quantity | price |
| --- | ---------- | ----- | -------- | ----- |
| 1   | 2015-08-15 | Pizza | 2        | 25    |
| 2   | 2016-05-11 | Sushi | 1        | 37    |
| 3   | 2015-01-01 | Wok   | 5        | 8     |
| 4   | 2016-12-31 | Cake  | 1        | 49    |

the output should be

| id  | year | quarter | type  | total_price |
| --- | ---- | ------- | ----- | ----------- |
| 1   | 2015 | 3       | Pizza | 50          |
| 2   | 2016 | 2       | Sushi | 37          |
| 3   | 2015 | 1       | Wok   | 40          |
| 4   | 2016 | 4       | Cake  | 49          |

# Output
- [execution time limit] 10 seconds (mysql)


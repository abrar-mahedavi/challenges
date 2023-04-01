CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT name as names
    FROM (
        (SELECT *, 1 AS dep FROM pr_department ORDER BY date_joined DESC, name LIMIT 5)
        UNION
        (SELECT *, 2 AS dep FROM it_department ORDER BY date_joined DESC, name LIMIT 5)
        UNION
        (SELECT *, 3 AS dep FROM sales_department ORDER BY date_joined DESC, name LIMIT 5)
        ) u
    ORDER BY dep, name;
END

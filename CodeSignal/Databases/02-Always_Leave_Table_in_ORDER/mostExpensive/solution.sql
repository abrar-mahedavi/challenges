CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT name FROM Products ORDER BY price*quantity DESC, name ASC LIMIT 1;
END

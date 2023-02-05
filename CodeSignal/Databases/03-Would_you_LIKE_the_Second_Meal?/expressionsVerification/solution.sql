CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT *
    FROM expressions
    WHERE (operation='+' AND c=a+b)
        OR(operation='-' AND c=a-b)
        OR(operation='/' AND c=a/b)
        OR(operation='*' AND c=a*b);
END

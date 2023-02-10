CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT continent,sum(users) as users from countries  GROUP BY continent ORDER by users DESC ;
END

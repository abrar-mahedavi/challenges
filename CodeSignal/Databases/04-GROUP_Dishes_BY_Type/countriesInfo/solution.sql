CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT COUNT(*) as number, AVG(population) as average, SUM(population) as total FROM countries;
END

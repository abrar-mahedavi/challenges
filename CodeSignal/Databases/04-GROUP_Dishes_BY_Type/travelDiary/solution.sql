CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT GROUP_CONCAT(t.country SEPARATOR ';') AS countries
  FROM (SELECT DISTINCT country FROM diary ORDER BY country) t;
END

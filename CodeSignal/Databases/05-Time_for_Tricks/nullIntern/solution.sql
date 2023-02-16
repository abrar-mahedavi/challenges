CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT count(*) AS number_of_nulls
  FROM departments
  WHERE description REGEXP '^[[:space:]]*(NULL|nil|-)[[:space:]]*$' OR description IS NULL;
END

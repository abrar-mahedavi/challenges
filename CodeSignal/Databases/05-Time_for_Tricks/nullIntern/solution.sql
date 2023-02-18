CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT COUNT(id) AS number_of_nulls
    FROM departments
    WHERE description IS NULL
    OR UPPER(TRIM(description))='NULL'
    OR UPPER(TRIM(description))='NIL'
    OR UPPER(TRIM(description))='-';
END

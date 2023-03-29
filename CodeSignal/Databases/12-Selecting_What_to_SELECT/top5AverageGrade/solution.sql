CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT ROUND(AVG(grade),2) AS average_grade
    FROM (
        SELECT grade
        FROM students
        ORDER BY grade DESC
        LIMIT 5) t;
END

CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT dep_name
    FROM departments
    WHERE NOT EXISTS (
    SELECT * FROM employees
    WHERE employees.department = departments.id)
    ORDER BY id;
END

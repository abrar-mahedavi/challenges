CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT dep_name,  emp_name
  FROM departments, employees
  ORDER BY dep_name, emp_name;
END

CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT IFNULL(maxd - mind,0) AS difference FROM
(SELECT mins*COUNT(*) AS mind FROM employees e
    JOIN (SELECT MIN(Salary) AS mins FROM employees LIMIT 1) tmin
    ON e.salary = mins) t1,

(SELECT maxs*COUNT(*) AS maxd FROM employees e
    JOIN (SELECT MAX(Salary) AS maxs FROM employees LIMIT 1) tmAx
    ON e.salary = maxs) t2;

END

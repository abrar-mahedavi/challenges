CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT IF(gender='F',CONCAT("Queen ",name),CONCAT("King ",name)) AS name FROM Successors ORDER BY birthday;
END

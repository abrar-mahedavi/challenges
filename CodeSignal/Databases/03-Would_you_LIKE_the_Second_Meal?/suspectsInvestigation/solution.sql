CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT id,name,surname FROM Suspect WHERE upper(name) LIKE 'b%' AND height<=170 AND upper(surname) LIKE 'gre_n';
END

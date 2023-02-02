CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT id,name,surname FROM Suspect WHERE height <= 170 OR upper(surname) NOT LIKE 'gre_n' OR upper(name) NOT LIKE 'B%' ORDER BY id;
END

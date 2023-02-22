CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT name
    FROM people_hobbies
    WHERE hobbies & 3  =3
    ORDER BY name;
END

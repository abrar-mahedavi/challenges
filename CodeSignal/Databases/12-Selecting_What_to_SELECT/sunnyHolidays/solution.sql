CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT holiday_date as ski_date
    FROM holidays INNER JOIN weather ON holiday_date = sunny_date;
END

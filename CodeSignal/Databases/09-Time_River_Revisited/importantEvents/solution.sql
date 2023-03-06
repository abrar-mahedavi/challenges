CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT *
    FROM events
    ORDER BY (DAYOFWEEK(event_date) + 5) % 7, participants DESC;
END

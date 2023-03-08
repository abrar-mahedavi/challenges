CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT name, event_date
    FROM Events,
        (
            SELECT MAX(event_date) AS maxdate
            FROM Events
        ) AS recent_date
    WHERE DATEDIFF(recent_date.maxdate,event_date) <= 7 AND recent_date.maxdate != event_date
    ORDER BY event_date DESC;
END

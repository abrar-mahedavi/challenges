CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT event_id, DATE_FORMAT(ADDDATE(date, INTERVAL timeshift MINUTE), CONCAT('%Y-%m-%d ',IF(hours=12,'%h:%i %p','%H:%i'))) AS formatted_date
    FROM events JOIN settings USING(user_id)
    ORDER BY event_id;
END

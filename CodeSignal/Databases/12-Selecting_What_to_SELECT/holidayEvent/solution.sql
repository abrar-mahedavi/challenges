CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SET @purchase = 0;
	SELECT DISTINCT buyer_name AS winners
    FROM(
        SELECT * FROM purchases
        WHERE (@purchase:=@purchase + 1)%4 = 0
        ORDER BY timestamp) t
    ORDER BY buyer_name;
END

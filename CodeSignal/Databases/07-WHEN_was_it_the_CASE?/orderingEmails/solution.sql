CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT id, email_title,
    CASE
        WHEN size >= POW(2,20)
        THEN CONCAT(ROUND(FLOOR(size/POW(2,20)))," Mb")
        ELSE CONCAT(ROUND(FLOOR(size/POW(2,10)))," Kb")
    END AS short_size
    FROM emails
    ORDER BY size DESC;
END

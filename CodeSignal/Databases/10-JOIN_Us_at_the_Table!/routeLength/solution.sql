CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT ROUND(SUM(ST_Distance(Point(c.x, c.y), Point(d.x, d.y))), 9) AS total
    FROM cities AS c LEFT JOIN cities AS d ON (c.id + 1) = d.id ;
END

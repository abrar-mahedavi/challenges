CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT IFNULL(country, 'Total:') as country, COUNT(competitor) as competitors
  FROM foreignCompetitors
  GROUP BY country WITH ROLLUP;
END

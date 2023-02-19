CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SET @c = 1;
  SELECT @c := @c * LENGTH(characters) AS combinations
  FROM discs
  ORDER BY combinations DESC
  LIMIT 1;
END

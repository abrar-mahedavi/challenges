CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT GROUP_CONCAT(CONCAT_WS(' ', first_name, surname, CONCAT('#', player_number)) ORDER BY player_number SEPARATOR '; ') AS players
  FROM soccer_team
  ORDER BY player_number;
END

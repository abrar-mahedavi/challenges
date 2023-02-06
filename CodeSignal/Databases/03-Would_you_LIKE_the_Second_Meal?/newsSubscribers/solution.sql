CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	 SELECT subscriber FROM full_year WHERE instr(newspaper,'Daily')
   UNION
   SELECT subscriber FROM half_year WHERE instr(newspaper,'Daily') ORDER BY subscriber;
END

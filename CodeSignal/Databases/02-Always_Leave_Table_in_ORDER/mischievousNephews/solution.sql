CREATE PROCEDURE solution()
BEGIN
	  /* Write your SQL here. Terminate each statement with a semicolon. */
		SELECT weekday(mischief_date) AS weekday, mischief_date, author,title
    FROM mischief
    ORDER BY weekday,field(author,"Huey","Dewey","Louie"),mischief_date,title;
END

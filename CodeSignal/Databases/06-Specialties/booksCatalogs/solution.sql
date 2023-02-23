CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT ExtractValue(xml_doc, '/descendant-or-self::author[1]') AS author
    FROM catalogs
    GROUP BY author
    ORDER BY author;
END

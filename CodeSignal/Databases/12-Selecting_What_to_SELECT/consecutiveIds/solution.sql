CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SET @id = 0;
	SELECT  id AS oldId, @id:=@id + 1 AS newId
    FROM itemIds
    ORDER BY id;
END

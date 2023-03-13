CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT candidate_id AS student_id
    FROM candidates LEFT JOIN detentions ON candidate_id = student_id
    WHERE student_id is NULL;
END

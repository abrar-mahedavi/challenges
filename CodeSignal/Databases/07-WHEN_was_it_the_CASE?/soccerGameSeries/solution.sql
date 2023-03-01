CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT CASE
        WHEN SUM(IF(fts > sts,1,0)) > SUM(IF(fts < sts,1,0)) THEN 1
        WHEN SUM(IF(fts > sts,1,0)) < SUM(IF(fts < sts,1,0)) THEN 2
        WHEN SUM(fts) > SUM(sts) THEN 1
        WHEN SUM(fts) < SUM(sts) THEN 2
        WHEN SUM(IF(mh = 2,fts,0)) > SUM(IF(mh = 1,sts,0)) THEN 1
        WHEN SUM(IF(mh = 2,fts,0)) < SUM(IF(mh = 1,sts,0)) THEN 2
        ELSE 0
        END as winner
    FROM
    (
        SELECT first_team_score as fts, second_team_score as sts, match_host as mh
        FROM scores
    ) x;
END

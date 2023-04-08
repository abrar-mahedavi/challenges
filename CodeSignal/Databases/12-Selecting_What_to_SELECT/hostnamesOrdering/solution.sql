CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SELECT id, hostname
FROM
    (SELECT
        id,
        hostname,
        domain3,
        CASE
            WHEN LENGTH(hostname) = LENGTH(domain3) THEN NULL
            ELSE SUBSTRING_INDEX(domain2domain3, '.', 1)
        END
            AS domain2,
        CASE
            WHEN LENGTH(hostname) = LENGTH(domain2domain3) THEN NULL
            ELSE SUBSTRING_INDEX(hostname, '.', 1)
        END
            AS domain1
    FROM
        (SELECT
            id,
            hostname,
            SUBSTRING_INDEX(hostname, '.', -1) AS domain3,
            SUBSTRING_INDEX(hostname, '.', -2) AS domain2domain3
        FROM hostnames) AS t1) AS t2
ORDER BY domain3, domain2, domain1;
END

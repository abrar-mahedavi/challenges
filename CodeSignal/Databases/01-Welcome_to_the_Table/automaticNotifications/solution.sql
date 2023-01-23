CREATE PROCEDURE solution()
    SELECT email
    FROM users
    WHERE role NOT IN ("admin", "premium")

    ORDER BY email;

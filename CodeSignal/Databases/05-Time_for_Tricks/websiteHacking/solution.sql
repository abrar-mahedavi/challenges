CREATE PROCEDURE solution()
    SELECT id,login,name
    FROM users
    WHERE type='user'
    UNION (SELECT id, login,name FROM users WHERE type<>'user')
    ORDER BY id

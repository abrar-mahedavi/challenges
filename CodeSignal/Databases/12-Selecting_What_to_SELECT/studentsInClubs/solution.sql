CREATE PROCEDURE solution()
    SELECT * FROM students
    WHERE EXISTS (
        SELECT * FROM clubs WHERE clubs.id = students.club_id
    )
    ORDER BY students.id;

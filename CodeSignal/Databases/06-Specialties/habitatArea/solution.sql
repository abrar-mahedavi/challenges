CREATE PROCEDURE solution()
BEGIN
	/* Write your SQL here. Terminate each statement with a semicolon. */
	SET @g = (SELECT CONCAT('MULTIPOINT(', GROUP_CONCAT( CONCAT(x, ' ', y) SEPARATOR ','),')')      FROM places);
  SELECT ST_Area(ST_ConvexHull(ST_GeomFromText(@g))) AS area;
END

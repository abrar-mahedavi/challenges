CREATE PROCEDURE solution()
BEGIN
	SELECT
		ae.id,
		IF(sum(ep.expenditure_sum)-sum(distinct ae.value)<=0
			,0
			,sum(ep.expenditure_sum)-sum(distinct ae.value)) loss
	FROM
	expenditure_plan ep
	INNER JOIN allowable_expenditure ae ON 1=1
		AND WEEK(ep.monday_date, 2) >= ae.left_bound
		AND WEEK(ep.monday_date, 2) <= ae.right_bound
	GROUP BY
		ae.id
	ORDER BY
		ae.id;
END

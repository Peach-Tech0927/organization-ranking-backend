SELECT
     o.id AS organization_id,
     o.name AS organization_name,
     COALESCE(SUM(u.contributions),0) AS total_score
 FROM 
     organizations o
 LEFT JOIN 
     user_organizations uo ON o.id = uo.organization_id
 LEFT JOIN 
     users u ON u.id = uo.user_id
 GROUP BY 
     o.id, o.name
 ORDER BY 
     total_score DESC;
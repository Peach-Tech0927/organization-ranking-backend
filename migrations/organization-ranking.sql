SELECT
     o.id AS organization_id,
     o.name AS organization_name,
     COALESCE(SUM(u.contributions),0) AS total_contributions
 FROM 
     organizations o
-- 全団体を出すために左結合にする
 LEFT JOIN 
     user_organizations uo ON o.id = uo.organization_id
 LEFT JOIN 
     users u ON u.id = uo.user_id
 GROUP BY 
     o.id, o.name
 ORDER BY 
     total_contributions DESC;
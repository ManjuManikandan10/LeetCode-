-- Write your PostgreSQL query statement below
SELECT 
    user_id
    , reaction as dominant_reaction
    , round(n / n_total, 2) as reaction_ratio
FROM (
    SELECT 
        user_id
        , reaction
        , count(user_id) as n
        , sum(count(user_id)) OVER (PARTITION BY user_id) as n_total
    FROM reactions
    GROUP BY user_id, reaction
)
WHERE n / n_total >= 0.6 AND n_total >= 5
ORDER BY n / n_total DESC, user_id ASC
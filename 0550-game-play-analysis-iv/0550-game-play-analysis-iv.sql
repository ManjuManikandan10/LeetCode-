SELECT ROUND(
    COUNT(*) / (SELECT COUNT(DISTINCT player_id) FROM Activity),
    2
) AS fraction
FROM (
    SELECT
        player_id,
        event_date,
        LAG(event_date) OVER (
            PARTITION BY player_id
            ORDER BY event_date
        ) AS prev,
        ROW_NUMBER() OVER (
            PARTITION BY player_id
            ORDER BY event_date
        ) AS rnk
    FROM Activity
) t
WHERE rnk = 2
  AND prev = event_date - INTERVAL 1 DAY;
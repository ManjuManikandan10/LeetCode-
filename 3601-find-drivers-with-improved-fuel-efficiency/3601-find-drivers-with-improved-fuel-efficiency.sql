-- Write your PostgreSQL query statement below
WITH half_avgs AS(
    SELECT
        driver_id,
        AVG(distance_km / fuel_consumed) 
            FILTER(WHERE trip_date BETWEEN '2023-01-01' AND '2023-06-30')
            AS first_half_avg,
        AVG(distance_km / fuel_consumed) 
            FILTER(WHERE trip_date BETWEEN '2023-07-01' AND '2023-12-31') 
            AS second_half_avg
    FROM trips
    GROUP BY driver_id
)
SELECT 
    driver_id,
    driver_name,
    ROUND(first_half_avg, 2) AS first_half_avg,
    ROUND(second_half_avg, 2) AS second_half_avg,
    ROUND(second_half_avg - first_half_avg, 2) AS efficiency_improvement
FROM half_avgs
JOIN drivers
    USING(driver_id)
WHERE
    first_half_avg IS NOT NULL
    AND second_half_avg IS NOT NULL
    AND second_half_avg > first_half_avg
ORDER BY
    efficiency_improvement DESC,
    driver_name
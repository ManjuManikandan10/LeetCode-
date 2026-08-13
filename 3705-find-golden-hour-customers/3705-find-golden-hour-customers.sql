-- Step 1: Aggregate customer order stats
WITH customer_stats AS (
    SELECT 
        customer_id,
        COUNT(*) AS total_orders,  -- total number of orders
        COUNT(*) FILTER (
            WHERE (EXTRACT(HOUR FROM order_timestamp) BETWEEN 11 AND 13)
               OR (EXTRACT(HOUR FROM order_timestamp) BETWEEN 18 AND 20)
        ) AS peak_orders,  -- number of orders during peak hours
        COUNT(order_rating) AS rated_orders,  -- number of rated orders
        ROUND(AVG(order_rating), 2) AS average_rating  -- average rating of rated orders
    FROM restaurant_orders
    GROUP BY customer_id
),

-- Step 2: Filter customers who meet all golden hour criteria
golden_hour_customers AS (
    SELECT 
        customer_id,
        total_orders,
        ROUND(peak_orders::NUMERIC * 100 / total_orders, 0) AS peak_hour_percentage,  -- % of orders during peak hours
        average_rating
    FROM customer_stats
    WHERE 
        total_orders >= 3  -- at least 3 orders
        AND peak_orders::NUMERIC / total_orders >= 0.6  -- at least 60% during peak hours
        AND rated_orders::NUMERIC / total_orders >= 0.5  -- at least 50% rated
        AND average_rating >= 4.0  -- average rating at least 4.0
)

-- Step 3: Return sorted result
SELECT *
FROM golden_hour_customers
ORDER BY average_rating DESC, customer_id DESC;
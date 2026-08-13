WITH ranked_events AS (
    SELECT
        user_id,
        plan_name,
        event_type,
        event_date,
        monthly_amount,

        ROW_NUMBER() OVER (
            PARTITION BY user_id
            ORDER BY event_date DESC
        ) AS rn,

        MAX(monthly_amount) OVER (
            PARTITION BY user_id
        ) AS max_historical_amount,

        MIN(event_date) OVER (
            PARTITION BY user_id
        ) AS first_event_date

    FROM subscription_events
)

SELECT
    user_id,
    plan_name AS current_plan,
    monthly_amount AS current_monthly_amount,
    max_historical_amount,
    event_date - first_event_date AS days_as_subscriber
FROM ranked_events
WHERE rn = 1
  AND event_type = 'downgrade'
  AND (monthly_amount * 100.0 / max_historical_amount) < 50
  AND (event_date - first_event_date) >= 60
ORDER BY days_as_subscriber DESC, user_id;
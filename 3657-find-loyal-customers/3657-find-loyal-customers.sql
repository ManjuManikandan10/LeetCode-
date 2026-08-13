WITH final AS(SELECT customer_id,
    COUNT(transaction_type) FILTER(WHERE transaction_type = 'purchase') AS purch_count,
    COUNT(transaction_type) FILTER(WHERE transaction_type = 'refund') AS ref_count,
    ROUND(COUNT(transaction_type) FILTER(WHERE transaction_type = 'refund')::decimal 
        / COUNT(transaction_type), 4) AS refund_rate,
    MAX(transaction_date) - MIN(transaction_date) AS active
FROM customer_transactions
GROUP BY customer_id)

SELECT customer_id
FROM final
WHERE purch_count > 2 AND active > 29 AND refund_rate < 0.2
ORDER BY customer_id
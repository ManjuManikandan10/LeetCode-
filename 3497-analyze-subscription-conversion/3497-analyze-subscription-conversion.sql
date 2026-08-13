# Write your MySQL query statement below
SELECT u1.user_id, ROUND(AVG(u1.activity_duration), 2) AS 'trial_avg_duration', ROUND(AVG(u2.activity_duration), 2) AS 'paid_avg_duration' FROM UserActivity u1 JOIN UserActivity u2 ON u1.activity_type = 'free_trial' AND u2.activity_type = 'paid' AND u1.user_id = u2.user_id GROUP BY u1.user_id ORDER BY user_id; 


# More numbers (for 1 it's multiplied by 3 (*3) but similarly by calculating average it'll be /3 as well so they'll calcel out as it'll be like 3/3 vs 9/9)
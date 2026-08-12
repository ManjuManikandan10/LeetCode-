# Write your MySQL query statement below
with t1 as (select *, sum(weight) over (order by turn) as running_weight from Queue)
select person_name from t1 where running_weight<=1000
order by running_weight desc
limit 1;
# Write your MySQL query statement below
with cte as (
    select u.name, ifnull(sum(r.distance),0) as travelled_distance  from users as u
    left join rides as r
    on u.id=r.user_id
    group by u.id
)select name, travelled_distance from cte
order by travelled_distance desc , name asc;
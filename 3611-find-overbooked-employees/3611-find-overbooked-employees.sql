-- Write your PostgreSQL query statement below
with tot_dur as(
    select e.employee_id , e.employee_name ,e.department ,
    date_trunc('week',m.meeting_date)as week_start,
    sum(duration_hours)as tot_hour
    from employees  e join meetings m on e.employee_id =m.employee_id 
    group by e.employee_id, e.employee_name, e.department, week_start
)
select employee_id , employee_name ,department,
     COUNT(employee_id) AS meeting_heavy_weeks
     from tot_dur
     where tot_hour>20
     group by employee_id , employee_name ,department
     having COUNT(employee_id)>=2
order by meeting_heavy_weeks desc , employee_name asc


select a.store_id,
store_name ,
location,
ii.product_name as most_exp_product,
i.product_name as cheapest_product,
round(i.quantity*1.0/ii.quantity,2) as imbalance_ratio   
from 
(select store_id,
min(price) as min_value,
max(price) as max_value,
count(*) as countt
from inventory 
group by store_id ) a
join stores s on
a.store_id = s.store_id 
join inventory i on 
a.store_id = i.store_id and 
a.min_value = i.price
join inventory ii on 
a.store_id = ii.store_id and 
a.max_value = ii.price
where ii.quantity  < i.quantity and 
countt >= 3
order by round(i.quantity*1.0/ii.quantity,2) desc, store_name
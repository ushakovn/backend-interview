-- 1. Список уникальных клиентов и купленных ими товаров в феврале

select
    distinct
    user.id,
    user.firstname,
    user.lastname,
    purchase.sku
from user
    inner join purchase on user.id = purchase.user_id
    left join ban_list on user.id = ban_list.user_id
where
    date_part('month', timestamp purchase.date) = 1
and
    user.id != ban_list.user_id

-- 2. Список уникальных клиентов,
-- совершивших покупку товаров на сумму больше 5000 и не имеющих бан

select
    distinct
    user.id,
    user.firstname,
    user.lastname,
    sum(purchase.sku) as total_sum
from user as u
    inner join purchase as p on u.id = p.user_id
    left join ban_list as b on u.id = b.user_id
where u.id != b.user_id
    group by u.id
        having(total_sum) > 5000

# 写一条 SQL 查询语句获取每位玩家 第一次登陆平台的日期
# group by+agg聚合函数
select 
    player_id,
    min(event_date) first_login
from activity
group by player_id

-- 请编写一个 SQL 查询，描述每一个玩家首次登陆的设备名称

select 
player_id, device_id
from activity
where (player_id, event_date) in 
(
select player_id, min(event_date)
from activity
group by player_id
)

-- 编写一个 SQL 查询，同时报告每组玩家和日期，以及玩家到目前为止玩了多少游戏。也就是说，在此日期之前玩家所玩的游戏总数。详细情况请查看示例。
-- 534. 游戏玩法分析 III
-- 
select 
    a1.player_id,
    a1.event_date,
    sum((a1.event_date>=a2.event_date)*a2.games_played) as games_played_so_far
from Activity a1,Activity a2
# 自连接 全连接
where a1.player_id=a2.player_id
group by a1.player_id,a1.event_date
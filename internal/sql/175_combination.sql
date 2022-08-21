
-- 编写一个SQL查询来报告 Person 表中每个人的姓、名、城市和州。如果 personId 的地址不在 Address 表中，则报告为空  null 。
select FirstName, LastName, City, State
from Person, left join Address
on Person.PersonId = Address.PersonId
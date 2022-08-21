# Write your MySQL query statement below
-- IFNULL函数用于判断第一个表达式是否为NULL，否则取第二个参数值
SELECT
    IFNULL(
      (SELECT DISTINCT Salary
       FROM Employee
       ORDER BY Salary DESC
        LIMIT 1 OFFSET 1),
    NULL) AS SecondHighestSalary


-- 编写一个SQL查询来报告 Employee 表中第 n 高的工资。如果没有第 n 个最高工资，查询应该报告为 null 
-- 1
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    SET N := N-1;
  RETURN (
      # Write your MySQL query statement below.
      SELECT 
            salary
      FROM 
            employee
      GROUP BY 
            salary
      ORDER BY 
            salary DESC
      LIMIT N, 1 #跳过前N个数,取一个第N个数
  );
END

-- 2
SELECT 
            DISTINCT salary
        FROM 
            (SELECT 
                salary, dense_rank() over(ORDER BY salary DESC) AS rnk
             FROM 
                employee) tmp
        WHERE rnk = N
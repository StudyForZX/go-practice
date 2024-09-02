-- table
-- | user_id | activity_date |
-- |  user1  |   2023-03-01  |
-- |  user2  |   2023-03-02  |
-- |  user3  |   2023-03-03  |
-- |  user1  |   2023-03-08  |
-- |  user1  |   2023-03-09  |
-- |  user1  |   2023-03-10  |

-- 方法一: 使用row_number
SELECT
  user_id,
  min(activity_date) as min_date,
  max(activity_date) as max_date,
  count(1) as login_times
FROM
(
  SELECT
    user_id,
    activity_date,
    DATE_SUB(activity_date, INTERVAL rn DAY) as sub_date
  FROM
  (
    SELECT
      user_id,
      activity_date,
      ROW_NUMBER() OVER (partition by user_id order by activity_date) as rn
    FROM
      table
  )t1
)t2
GROUP BY
  user_id,
  sub_date
having login_times >= 3

-- 方法二: 使用lag和lead函数
SELECT
  user_id,
  min(current_login_date) as min_date,
  max(current_login_date) as max_date,
  count(1) as login_times
  (
  SELECT
    user_id,
    current_login_date,
    lag_login_date,
    lead_login_date
  FROM
    (
    SELECT
      user_id,
      activity_date AS current_login_date,
      LAG(activity_date, 1, activity_date) OVER (PARTITION BY user_id ORDER BY activity_date) AS lag_login_date,
      LEAD(activity_date，1， activity_date) OVER (PARTITION BY user_id ORDER BY activity_date) AS lead_login_date
    FROM
      TABLE
    )t1
  WHERE
    DATEDIFF(current_login_date, lag_login_date) = 1
  AND
    DATEDIFF(lead_login_date, current_login_date) = 1
  )t2
GROUP BY
  user_id,
  current_login_date

-- 方法三：使用lead/lead函数
-- 去重表
WITH distinct_table(user_id, activity_date)
AS
(
  SELECT
    user_id,
    DISTINCT() DATE(activity_date) AS activity_date
  FROM
    TABLE
)
-- lead 向下两天的表
WITH lead_table(user_id, activity_date, next_2_days)
AS
(
  SELECT
    user_id,
    activity_date,
    LEAD(activity_date, 2, activity_date) OVER (PARTITION BY user_id ORDER BY activity_date) AS next_2_days
  FROM
    distinct_table
)
-- 求出diff
WITH diff_days_table(user_id, activity_date, diff)
AS
(
  SELECT
    user_id,
    activity_date,
    DATEDIFF(next_2_days) AS diff
  FROM
    lead_table
)
-- 最终结果
SELECT
  user_id,
  min(activity_date) AS min_date,
  max(activity_date) AS max_date
FROM
  diff_days_table
WHERE
  diff = 2

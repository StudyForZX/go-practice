-- 第一步 去重
WITH distinct_role_login(role_id, login_date)
AS
(
  SELECT
    DISTINCT STR_TO_DATE(part_date, "%Y-%m-%d") login_date,
    role_id
  FROM
    role_login
)
-- 第二步 为了计算某条登录日志是该用户创建账号后的第几天登录
-- 我们可以用用户登录日志和账号创建日志进行inner join
-- 用登陆日期字段和创建账户的字段进行差值比较获取第几天登陆

WITH diff_day_table(role_id, login_date, create_date, diff)
AS
(
  SELECT
    drl.role_id,
    drl.login_date，
    rc.create_date,
    DATEDIFF(drl.login_date, rc.create_date) AS diff
  FROM
    distinct_role_login drl
  INNER JOIN
    role_create as rc ON drl.role_id = rc.role_id
)
-- 第三步 计算总数
WITH count_table(create_date, `新增用户数`, `次日留存`, `3日留存`, `7日留存`
) AS
(
  SELECT
    create_date,
    count((CASE WHEN (diff = 0) THEN role_id END )) AS `新增用户数`,
    count((CASE WHEN (diff = 1) THEN role_id END )) AS `次日留存`,
    count((CASE WHEN (diff = 3) THEN role_id END )) AS `3日留存`,
    count((CASE WHEN (diff = 7) THEN role_id END )) AS `7日留存`
  FROM
    diff_day_table
  GROUP BY
    create_date
)
-- 第四步 计算结果
SELECT
  create_date,
  `新增用户数`,
  concat(CAST(ROUND((100 * 次日留存) / `新增用户数`), 2) AS varchar, '%') AS `次日留存率`
  concat(CAST(ROUND((100 * 3日留存) / 新增用户数, 2) AS varchar), '%') AS `3日留存率`
  concat(CAST(ROUND((100 * 7日留存) / 新增用户数, 2) AS varchar), '%') AS `7日留存率`
FROM
  count_table
ORDER BY
  create_date ASC

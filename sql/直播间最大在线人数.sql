-- -- 举例如下：
-- -- 输入
-- uid  op_time  op_type  oid
-- 1001  11:11      1      A
-- 1002  11:12      1      A
-- 1002  11:13      0      A
-- 1003  11:14      1      A
-- 1005  11:19      1      A
-- 1006  11:23      1      A
-- 1001  11:30      0      A
-- -- 输出
-- oid  max_cnt  dura_ts
-- A      3        7 

WITH t1 AS (
    SELECT
        uid,
        op_time,
        oid,
        sum(case when op_type = 1 THEN 1 ELSE -1) OVER (PARTITION BY oid ORDER BY op_time) AS cnt
    FROM
        table
),
t2 AS (
    SELECT
        uid,
        op_time,
        oid,
        cnt,
        max(cnt) over (PARTITION BY oid) as max_cnt,
        lead(op_time, 1) over (PARTITION BY oid ORDER BY op_time) as end_time
    FROM
        t1
),
SELECT
    oid,
    max_cnt,
    end_time - op_time as dura_ts
FROM
    t2
WHERE
    cnt = max_cnt

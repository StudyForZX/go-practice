-- Table
-- +-----------+------------+--------+
-- | player_id | match_day  | result |
-- +-----------+------------+--------+
-- | 1         | 2022-01-17 | Win    |
-- | 1         | 2022-01-18 | Win    |
-- | 1         | 2022-01-25 | Win    |
-- | 1         | 2022-01-31 | Draw   |
-- | 1         | 2022-02-08 | Win    |
-- | 2         | 2022-02-06 | Lose   |
-- | 2         | 2022-02-08 | Lose   |
-- | 3         | 2022-03-30 | Win    |
-- +-----------+------------+--------+

WITH RankedResults AS (
    SELECT
        player_id,
        match_day,
        result,
        CASE
            WHEN result = 'Win' THEN 1
            ELSE 0
        END AS win_flag,
        ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY match_day) AS rn,
        SUM(CASE WHEN result = 'Win' THEN 1 ELSE 0 END)
            OVER (PARTITION BY player_id ORDER BY match_day) AS cum_win,
        ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY match_day)
            - SUM(CASE WHEN result = 'Win' THEN 1 ELSE 0 END)
            OVER (PARTITION BY player_id ORDER BY match_day) AS win_streak_group
    FROM matches
)
SELECT * FROM RankedResults;

-- +-----------+------------+--------+----------+----+----------+------------------+
-- | player_id | match_day  | result | win_flag | rn | cum_win  | win_streak_group |
-- +-----------+------------+--------+----------+----+----------+------------------+
-- | 1         | 2022-01-17 | Win    | 1        |  1 | 1        | 0                |
-- | 1         | 2022-01-18 | Win    | 1        |  2 | 2        | 0                |
-- | 1         | 2022-01-25 | Win    | 1        |  3 | 3        | 0                |
-- | 1         | 2022-01-31 | Draw   | 0        |  4 | 3        | 1                |
-- | 1         | 2022-02-08 | Win    | 1        |  5 | 4        | 1                |
-- | 2         | 2022-02-06 | Lose   | 0        |  1 | 0        | 1                |
-- | 2         | 2022-02-08 | Lose   | 0        |  2 | 0        | 2                |
-- | 3         | 2022-03-30 | Win    | 1        |  1 | 1        | 0                |
-- +-----------+------------+--------+----------+----+----------+------------------+

WinStreaks AS (
    SELECT
        player_id,
        COUNT(*) AS streak_length
    FROM RankedResults
    WHERE win_flag = 1
    GROUP BY player_id, win_streak_group
)

-- +-----------+---------------+
-- | player_id | streak_length  |
-- +-----------+---------------+
-- | 1         | 3             |
-- | 1         | 1             |
-- | 3         | 1             |
-- +-----------+---------------+


SELECT
    player_id,
    MAX(streak_length) AS max_win_streak
FROM WinStreaks
GROUP BY player_id;

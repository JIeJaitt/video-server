# video_server
流媒体点播服务

# 数据库建表如下
```sql
CREATE TABLE users
(
    id         INT(10) UNSIGNED PRIMARY KEY,
    login_name VARCHAR(64) NULL UNIQUE,
    pwd        TEXT        NOT NULL
);
CREATE TABLE video_info
(
    id            VARCHAR(64) NOT NULL PRIMARY KEY,
    author_id     INT(10)     NULL,
    name          TEXT        NULL,
    display_ctime TEXT        NULL,
    create_time   DATETIME    NULL
);
CREATE TABLE comments
(
    id        VARCHAR(64) NOT NULL PRIMARY KEY,
    video_id  VARCHAR(64) NULL,
    author_id INT(10)     NULL,
    content   TEXT        NULL,
    time      DATETIME    NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE sessions
(
    session_id tinytext NOT NULL,
    TTL        tinytext,
    login_name varchar(64)
);
ALTER TABLE sessions
    ADD PRIMARY KEY (session_id(63));

DESCRIBE users;
DESCRIBE video_info;
DESCRIBE comments;
DESCRIBE sessions;
```

UNSIGNED INT： 无符号整数，意味该字段都是非负整数，这样可以扩大正数范围。比如TINYINT类型表示： -128～128， 无符号的TINYINT则表示0~256

TABLE: users
  - id UNSIGNED INT, PRIMARY KEY, AUTO_INCREMENT
  - login_name VARCHAR(64), UNIQUE KEY
  - pwd TEXT

TABLE: video_info
  - id VARCHAR(64), PRIMARY KEY, NOT NULL
  - author_id UNSIGNED INT
  - name TEXT
  - display_ctime TEXT
  -create_time DATETIME

TABLE: comments
  - id VARCHAR(64), PRIMARY KEY, NOT NULL
  - video_id VARCHAR(64)
  - author_id UNSIGNED INT
  - content TEXT
  - time DATETIME

TABLE: sessions
  - session_id TINYTEXT, PRIMARY KEY, NOT NULL
  - TTL TINYTEXT
  - login_name VARCHAR(64)


* 视频资源量肯定远大于用户数量


Question：
  - 什么是外键，什么是外键约束
  - 数据库设计第三范式？
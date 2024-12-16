


```sql
create table  if not exists short.short_url
(
    id int auto_increment
    primary key,
    short_url varchar(256) default '' not null,
    origin_url VARCHAR(2048) default '' not null,
    url_type int  default 1 not null,
    description varchar(256) default '' not null ,
    expire_at timestamp default null,
    created_at  timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null,
    deleted_at timestamp default null
    );

CREATE INDEX `idx_short_url_origin_url` ON
    `short_url` (`short_url`, `origin_url`(255));
CREATE INDEX `idx_short_url_deleted_at` ON
    `short_url` (`short_url`, `deleted_at`);

create table if not exists short.url_access_log
(
    id int auto_increment
    primary key,
    short_url varchar(256) default '' not null,
    user_agent VARCHAR(2048) default '',
    referrer VARCHAR(2048) default '',
    created_at  timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null,
    deleted_at timestamp default null
    );

CREATE INDEX `idx_url_access_log_short_url` ON
    `url_access_log` (`short_url`);
CREATE INDEX `idx_url_access_log_url_deleted_at` ON
    `url_access_log` (`short_url`, `deleted_at`);

```
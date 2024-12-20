# GoShortURL

GoShortURL is a high-performance short link service system developed based on the Go-Zero microservice framework.

## Main features

- Based on the Go-Zero framework, providing RESTful API services
- Using GORM as the ORM framework
- Multi-level cache design
- Local cache (Ristretto)
- Redis cache
- Efficient short link generation algorithm
- MurmurHash + Base62 encoding
- Conflict handling mechanism
- Support short link expiration time setting
- Support short link access statistics

## System architecture

- API layer: Go-Zero REST API
- Business layer: Manager + Logic
- Data layer: Repository Pattern
- Cache layer:
- Local cache (Ristretto)
- Redis cache
- Storage layer: MySQL

## Technology stack

- [Go-Zero](https://github.com/zeromicro/go-zero): Microservice framework
- [GORM](https://gorm.io): ORM framework
- [Ristretto](https://github.com/dgraph-io/ristretto): High-performance memory cache
- [Redis](https://redis.io): Distributed cache
- MySQL: Data storage

## SQL
```sql
create table  if not exists short.short_url
(
    id          int auto_increment
    primary key,
    short_url   varchar(256) default ''                not null,
    origin_url  text                                   not null,
    url_type    int          default 1                 not null,
    description varchar(256) default ''                not null,
    expire_at   timestamp                              null,
    created_at  timestamp    default CURRENT_TIMESTAMP not null,
    updated_at  timestamp    default CURRENT_TIMESTAMP not null,
    deleted_at  timestamp                              null
    );

create index idx_short_url_deleted_at
    on short_url (short_url, deleted_at);

create index idx_short_url_origin_url
    on short_url (short_url, origin_url(255));



create table if not exists short.url_access_log
(
    id         int auto_increment
    primary key,
    short_url  varchar(256)  default ''                not null,
    user_agent varchar(2048) default ''                null,
    referrer   varchar(2048) default ''                null,
    created_at timestamp     default CURRENT_TIMESTAMP not null,
    updated_at timestamp     default CURRENT_TIMESTAMP not null,
    deleted_at timestamp                               null,
    ip         varchar(255)  default ''                not null
    );

create index idx_url_access_log_short_url
    on url_access_log (short_url);

create index idx_url_access_log_url_deleted_at
    on url_access_log (short_url, deleted_at);
```
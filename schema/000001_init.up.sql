CREATE TABLE users 
(
    id      serial      NOT NULL UNIQUE,
    user_id int         UNIQUE
);

CREATE TABLE segments 
(
    id      serial      NOT NULL UNIQUE,
    segment varchar(30) UNIQUE
);

CREATE TABLE comparison 
(
    id          serial                                              NOT NULL UNIQUE,
    user_id     int REFERENCES users (id)       ON DELETE CASCADE   NOT NULL,
    segment_id  int REFERENCES segments (id)    ON DELETE CASCADE   NOT NULL
);
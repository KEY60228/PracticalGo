DROP TABLE IF EXISTS users;
CREATE TABLE users (
    user_id varchar(32) NOT NULL,
    user_name varchar(100) NOT NULL,
    created_at timestamp with time zone,
    CONSTRAINT pk_users PRIMARY KEY (user_id)
);

INSERT INTO users VALUES ('0001', 'Gopher', now()), ('0002', 'Ferris', now());

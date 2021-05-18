-- +goose Up
-- +goose StatementBegin
create table if not exists banner
(
    id          serial PRIMARY KEY,
    description text NOT NULL
);

create table if not exists slot
(
    id          serial PRIMARY KEY,
    description text NOT NULL
);

create table if not exists user_group
(
    id          serial PRIMARY KEY,
    description text NOT NULL
);

CREATE TABLE IF NOT EXISTS banner_slot
(
    banner_id serial references banner (id),
    slot_id   serial references slot (id),
    PRIMARY KEY (banner_id, slot_id)
);

CREATE TABLE IF NOT EXISTS banner_showing
(
    banner_id     serial references banner (id),
    slot_id       serial references slot (id),
    user_group_id serial references user_group (id),
    date          timestamptz NOT NULL,
    PRIMARY KEY (banner_id, slot_id, user_group_id, date)
);

CREATE TABLE IF NOT EXISTS banner_click
(
    banner_id     serial references banner (id),
    slot_id       serial references slot (id),
    user_group_id serial references user_group (id),
    date          timestamptz NOT NULL,
    PRIMARY KEY (banner_id, slot_id, user_group_id, date)
);

INSERT INTO banner (id, description) VALUES (1, 'banner 1'), (2, 'banner 2'), (3, 'banner 3');
INSERT INTO slot (id, description) VALUES (1, 'slot 1'), (2, 'slot 2'), (3, 'slot 3');
INSERT INTO user_group (id, description) VALUES (1, 'Дети'), (2,'Старики');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table banner_showing;
drop table banner_click;
drop table banner_slot;
drop table banner;
drop table slot;
drop table user_group;
-- +goose StatementEnd
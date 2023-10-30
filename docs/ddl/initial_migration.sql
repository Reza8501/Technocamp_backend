insert into roles (`name`) values ('admin');
insert into roles (`name`) values ('client');

insert into users (id, full_name, username, email, `password`, `role`, created_at, updated_at) values ('af380699-9714-46d3-9429-7b88eca313d9', 'admin', 'admin', 'admin@gmail.com', '$2a$12$1yywWzt1MfghWc5v/yUTkOPKaTyGMNcup8smCqSR24oBQDKXpVZYK', 'admin', now(), now());
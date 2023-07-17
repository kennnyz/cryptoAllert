CREATE TABLE IF NOT EXISTS users (
                                     id int primary key,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



create table if not exists transfers (
                                         id int primary key,
                                         type varchar(10) not null,
                                         amount decimal not null,
                                         price decimal not null,
                                         user_id int references users(id),
                                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

create table if not exists user_coins (
                                          name varchar(10),
                                          user_id int references users(id),
                                          amount decimal not null default 0,
                                          usd_amount decimal not null default 0
);

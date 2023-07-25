CREATE TABLE IF NOT EXISTS users (
                                     id int primary key,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



create table if not exists user_coins (
                                          name varchar(10),
                                          user_id int references users(id),
	                                      frequency int,
);

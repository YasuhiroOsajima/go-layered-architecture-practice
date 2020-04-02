CREATE TABLE user(id TEXT PRIMARY KEY, name TEXT NOT NULL, mailaddress TEXT NOT NULL, usertype TEXT NOT NULL);
CREATE TABLE circle(id TEXT PRIMARY KEY, name TEXT NOT NULL, owner TEXT NOT NULL);
CREATE TABLE user_circle(user_id TEXT NOT NULL, circle_id TEXT NOT NULL, PRIMARY KEY(user_id, circle_id));

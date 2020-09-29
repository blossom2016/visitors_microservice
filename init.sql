create table if not exists visitors(
id serial constraint visitors_pk primary key,
user_agent varchar(250),
datetime timestamp
);
Create INDEX IF NOT exists visitors_user_agent_index on visitors(user_agent);

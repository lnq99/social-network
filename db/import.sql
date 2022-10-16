truncate Photo cascade;
truncate Album cascade;
truncate Notification cascade;
truncate Reaction cascade;
truncate Relationship cascade;
truncate Comment cascade;
truncate Post cascade;
truncate Profile cascade;

copy Profile from '/tmp/csv/profile.csv' delimiter ',' csv header;
copy Post from '/tmp/csv/post.csv' delimiter ',' csv header;
copy Comment from '/tmp/csv/comment.csv' delimiter ',' csv header;
copy Reaction from '/tmp/csv/reaction.csv' delimiter ',' csv header;
copy Notification from '/tmp/csv/notification.csv' delimiter ',' csv header;
copy Album from '/tmp/csv/album.csv' delimiter ',' csv header;
copy Photo from '/tmp/csv/photo.csv' delimiter ',' csv header;

-- copy Relationship from '/tmp/DB/SocialNetwork/db/csv/relationship.csv' delimiter ',' csv header;

begin;
create temp table ttmp on commit drop
as select * from Relationship with no data;

copy ttmp from '/tmp/csv/relationship.csv' delimiter ',' csv header;

insert into Relationship
select distinct on (user1, user2) * from ttmp;
commit;


begin;
create temp table ttmp
as select * from Profile with no data;

copy ttmp from '/tmp/DB/SocialNetwork/db/csv/profile.csv' delimiter ',' csv header;

select email
from ttmp
group by email
having count(email) > 1;
commit;


select setval('photo_id_seq', (select max(id) from Photo));
select setval('album_id_seq', (select max(id) from Album));
select setval('notification_id_seq', (select max(id) from Notification));
select setval('comment_id_seq', (select max(id) from Comment));
select setval('post_id_seq', (select max(id) from Post));
select setval('profile_id_seq', (select max(id) from Profile));


update Profile
set avatars = 'https://eu.ui-avatars.com/api/?name=SM&background=random&size=32'
where id = 1;
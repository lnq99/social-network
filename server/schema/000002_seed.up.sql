insert into Profile(name, gender, birthdate, email, salt, hash, postcount)
values
    ('User1', 'M', '1999-1-1', 'user1@gmail.com', 'SXmZdHRT', '23h3nlI-gObbXQvg2DFDHClP4YA=', 2),
    ('User2', 'F', '2000-2-1', 'user2@gmail.com', 'Qqwoyajo', 'TtokyP6wRYoF6T09LpBHDfs1YIQ=', 1),
    ('User3', 'M', '2001-3-1', 'user3@gmail.com', 'peSzrrPI', 'MCo_TBd49VG3cVAFqJH0isHhHvE=', 1)
on conflict
do nothing;


insert into Post(userId, tags, content)
values
    (1, 'food', 'food content'),
    (1, '', 'no tag content'),
    (2, 'dog', 'dog content'),
    (3, 'cat', 'cat content')
on conflict
do nothing;

-- insert into Comment(userId, postId, parentId, content, created)
-- values (1, 8, 0, 'ok', now());

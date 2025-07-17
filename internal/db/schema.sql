CREATE TABLE IF NOT EXISTS public.admins (
    id serial4 NOT NULL,
    login varchar(32) NOT NULL,
    password varchar(32) NOT NULL,
    name varchar(32) NOT NULL,
    surname varchar(32) NOT NULL,
    CONSTRAINT admins_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.users (
    id serial4 NOT NULL,
    name varchar(32),
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.messages (
    id serial4 NOT NULL,
    user_id serial4 REFERENCES users(id),
    admin_id serial4 REFERENCES admins(id),
    time TIMESTAMP NOT NULL,
    text text NOT NULL,
    CONSTRAINT messages_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.questions_answers (
    id serial4 NOT NULL,
    question text NOT NULL,
    answer text NOT NULL,
    CONSTRAINT questions_answers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.template_answers (
    id serial4 NOT NULL,
    user_id serial4 REFERENCES users(id),
    question_id serial4 REFERENCES questions_answers(id),
    time TIMESTAMP NOT NULL,
    CONSTRAINT template_answers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.jwt_tokens (
    id serial4 NOT NULL,
    token text NOT NULL,
    admin_id serial4 REFERENCES admins(id),
    time TIME NOT NULL,
    is_expired BOOL DEFAULT FALSE,
    CONSTRAINT jwt_tokens_pkey PRIMARY KEY (id)
);
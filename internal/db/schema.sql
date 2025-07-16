CREATE TABLE IF NOT EXISTS public.admins (
    id serial NOT NULL,
    login varchar(32) NOT NULL,
    password varchar(32) NOT NULL,
    name varchar(32) NOT NULL,
    surname varchar(32) NOT NULL,
    CONSTRAINT admins_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.users (
    id serial NOT NULL,
    name varchar(32),
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.messages (
    id serial NOT NULL,
    user_id int REFERENCES users(id),
    admin_id int REFERENCES admins(id),
    time TIMESTAMP NOT NULL,
    text text NOT NULL,
    CONSTRAINT messages_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.questions_answers (
    question_id int NOT NULL,
    question text NOT NULL,
    answer text NOT NULL,
    CONSTRAINT questions_answers_pkey PRIMARY KEY (question_id)
);

CREATE TABLE IF NOT EXISTS public.template_answers (
    id serial NOT NULL,
    user_id int REFERENCES users(id),
    question_id int REFERENCES questions_answers(question_id),
    time TIMESTAMP NOT NULL,
    CONSTRAINT template_answers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.GWT_tokens (
    id serial NOT NULL,
    token text NOT NULL,
    user_id int REFERENCES users(id),
    time TIME NOT NULL,
    is_expired BOOL DEFAULT FALSE,
    CONSTRAINT GWT_tokens_pkey PRIMARY KEY (id)
);
create type gender as ENUM('M','F');
CREATE TABLE post
(
  id        bigserial     NOT NULL,
  grade     integer       NOT NULL,
  subject   varchar[]     NOT NULL,
  address   varchar       NOT NULL,
  extra     varchar,
  post_time      timestamp     NOT NULL DEFAULT NOW(),
  available boolean       NOT NULL DEFAULT True,
  owner_id  bigint     NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE post_register
(
  approved   boolean   NOT NULL DEFAULT False,
  teacher_id bigint NOT NULL,
  post_id    bigint NOT NULL,
  PRIMARY KEY (teacher_id,post_id)
);

CREATE TABLE teacher
(
  id        bigserial     NOT NULL,
  email     varchar       NOT NULL,
  password  varchar       NOT NULL,
  full_name varchar       NOT NULL,
  tel       varchar       NOT NULL,
  address  varchar      ,
  grade     integer[]     NOT NULL,
  subject   varchar[]     NOT NULL,
  gender    gender NOT NULL,
  confirmed bool          NOT NULL DEFAULT False,
  PRIMARY KEY (id)
);

CREATE TABLE customer
(
  id        bigserial NOT NULL,
  email     varchar   NOT NULL,
  password  varchar   NOT NULL,
  full_name varchar   NOT NULL,
  tel       varchar   NOT NULL,
  address   varchar  ,
  confirmed bool      NOT NULL DEFAULT False,
  PRIMARY KEY (id)
);

ALTER TABLE post
  ADD CONSTRAINT FK_customer_TO_post
    FOREIGN KEY (owner_id)
    REFERENCES customer (id) on delete cascade;

ALTER TABLE post_register
  ADD CONSTRAINT FK_teacher_TO_post_register
    FOREIGN KEY (teacher_id)
    REFERENCES teacher (id) on delete cascade;

ALTER TABLE post_register
  ADD CONSTRAINT FK_post_TO_post_register
    FOREIGN KEY (post_id)
    REFERENCES post (id) on delete cascade;

        
      
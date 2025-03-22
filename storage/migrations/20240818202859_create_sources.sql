-- +goose Up
-- +goose StatementBegin
CREATE TABLE homework (
    id SERIAL NOT NULL UNIQUE ,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(1024),
    images TEXT[],
    create_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deadline TIMESTAMP NOT NULL DEFAULT NOW(),
    update_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE tags (
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE homeworks_tags (
    id SERIAL PRIMARY KEY UNIQUE,
    homework_id INT REFERENCES homework (id) ON DELETE CASCADE NOT NULL,
    tag_id INT REFERENCES tags (id) ON DELETE CASCADE NOT NULL
);


-- +goose Up
-- +goose StatementBegin
CREATE TABLE tg_users (
                          id uuid PRIMARY KEY,
                          username VARCHAR(255) NOT NULL UNIQUE,
                          user_id serial NOT NULL,
                          associates VARCHAR(255),
                          amount VARCHAR(255) ,
                          tron_amount VARCHAR(255),
                          tron_address VARCHAR(50),
                          eth_address VARCHAR(50),
                          eth_amount VARCHAR(255),
                          create_at TIMESTAMP NOT NULL DEFAULT NOW(),
                          deadline TIMESTAMP NOT NULL DEFAULT NOW(),
                          update_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE homeworks_tags;

DROP TABLE homework;

DROP TABLE tags;
-- +goose StatementEnd

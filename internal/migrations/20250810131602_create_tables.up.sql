-- migrations/001_create_tables.up.sql
CREATE TABLE IF NOT EXISTS jobs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    apply_url TEXT,
    last_date DATE,
    date_posted TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS results (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    result_link TEXT,
    exam_date DATE,
    date_posted TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS admit_cards (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    download_link TEXT,
    exam_date DATE,
    date_posted TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS vimps_links (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    link TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS imp_links (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    link TEXT NOT NULL
);

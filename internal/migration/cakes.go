package migration

var CakeSchema = `
CREATE TABLE IF NOT EXISTS cakes (
	id BIGINT UNSIGNED auto_increment NOT NULL,
	title varchar(100) NOT NULL,
	description text NULL,
	rating FLOAT NULL,
	image varchar(100) NULL,
	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
	CONSTRAINT cakes_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;`

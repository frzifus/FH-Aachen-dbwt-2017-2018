DROP TABLE IF EXISTS products_ingredients;
DROP TABLE IF EXISTS ingredients;
DROP TABLE IF EXISTS prices;
DROP TABLE IF EXISTS products_orders;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categorys;
DROP TABLE IF EXISTS images;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS employees;
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS guests;
DROP TABLE IF EXISTS members;
DROP TABLE IF EXISTS users;


CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED AUTO_INCREMENT,
    active TINYINT(1) NOT NULL DEFAULT TRUE,
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    mail VARCHAR(100) NOT NULL UNIQUE,
    loginname VARCHAR(100) NOT NULL,
    created_at DATE DEFAULT CURRENT_DATE,
    last_login TIMESTAMP,
    stretch INT UNSIGNED NOT NULL,
    algo VARCHAR(6) CHECK(algo = 'sha256'),
    salt VARCHAR(32) NOT NULL,
    `hash` VARCHAR(64) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS members (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS guests (
    reason VARCHAR(200) NOT NULL,
    expiry_date DATE DEFAULT CURRENT_DATE,
    id INT UNSIGNED NOT NULL,
    FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS students (
    id INT UNSIGNED NOT NULL UNIQUE CHECK( id > 10000
        AND id < 9999999),
    course VARCHAR(100) NOT NULL,
    member_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (id) REFERENCES members(id) ON DELETE CASCADE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS employees (
    phone_number INT UNSIGNED,
    office VARCHAR(4),
    id INT UNSIGNED NOT NULL,
    FOREIGN KEY (id) REFERENCES members(id) ON DELETE CASCADE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS orders (
    id INT UNSIGNED AUTO_INCREMENT,
    `time` TIMESTAMP NOT NULL,
    user_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS images (
    id INT UNSIGNED AUTO_INCREMENT,
    blob_data BLOB NOT NULL,
    alttext VARCHAR(60),
    title VARCHAR(60),
    caption VARCHAR(80),
    PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS categorys (
    id INT UNSIGNED AUTO_INCREMENT,
    designation VARCHAR(100),
    upper_category_id INT UNSIGNED DEFAULT NULL,
    image_id INT UNSIGNED,
    CONSTRAINT upper_category FOREIGN KEY (upper_category_id)
      REFERENCES categorys(id),
    FOREIGN KEY (image_id) REFERENCES images(id),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS products (
    id INT UNSIGNED AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    vegetarian TINYINT(1),
    vegan TINYINT(1),
    image_id INT UNSIGNED NOT NULL,
    category_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (image_id) REFERENCES images(id),
    FOREIGN KEY (category_id) REFERENCES categorys(id),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS products_orders (
    product_id INT UNSIGNED NOT NULL,
    order_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

CREATE TABLE IF NOT EXISTS prices (
    guest INT UNSIGNED NOT NULL,
    student INT UNSIGNED NOT NULL,
    employee INT UNSIGNED NOT NULL,
    id INT UNSIGNED NOT NULL,
    FOREIGN KEY (id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS ingredients (
    id INT UNSIGNED AUTO_INCREMENT,
    gluten_free TINYINT(1),
    bio TINYINT(1),
    vegetarian TINYINT(1),
    vegan TINYINT(1),
    description TEXT,
    `name` VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS products_ingredients (
    product_id INT UNSIGNED NOT NULL,
    ingredient_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id)
);

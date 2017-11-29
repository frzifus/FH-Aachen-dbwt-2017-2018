DROP TABLE IF EXISTS trans_produkt_zutat;
DROP TABLE IF EXISTS zutat;
DROP TABLE IF EXISTS preis;
DROP TABLE IF EXISTS trans_produkt_bestellung;
DROP TABLE IF EXISTS produkt;
DROP TABLE IF EXISTS kategorie;
DROP TABLE IF EXISTS bild;
DROP TABLE IF EXISTS bestellung;
DROP TABLE IF EXISTS mitarbeiter;
DROP TABLE IF EXISTS student;
DROP TABLE IF EXISTS gast;
DROP TABLE IF EXISTS angehoerige;
DROP TABLE IF EXISTS nutzer;


CREATE TABLE IF NOT EXISTS users (
    nr INT AUTO_INCREMENT,
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
    PRIMARY KEY(Nr)
);

CREATE TABLE IF NOT EXISTS members (
    nr INT NOT NULL AUTO_INCREMENT,
    FOREIGN KEY (nr) REFERENCES users(nr) ON DELETE CASCADE,
    PRIMARY KEY (nr)
);

CREATE TABLE IF NOT EXISTS guests (
    reason VARCHAR(200) NOT NULL,
    expiry_date DATE DEFAULT CURRENT_DATE,
    nr INT NOT NULL,
    FOREIGN KEY (nr) REFERENCES users(nr) ON DELETE CASCADE,
    PRIMARY KEY (nr)
);

CREATE TABLE IF NOT EXISTS students (
    student_id INT UNSIGNED NOT NULL UNIQUE CHECK( matrikelnummer > 10000
        AND matrikelnummer < 9999999),
    course VARCHAR(100) NOT NULL,
    nr INT NOT NULL,
    FOREIGN KEY (nr) REFERENCES angehoerige(nr) ON DELETE CASCADE,
    PRIMARY KEY (nr)
);

-- -------------------------------------------------------------

CREATE TABLE IF NOT EXISTS mitarbeiter (
    telefonnummer INT,
    buero VARCHAR(4),
    nr INT NOT NULL,
    FOREIGN KEY (nr) REFERENCES angehoerige(nr) ON DELETE CASCADE,
    PRIMARY KEY (nr)
);

CREATE TABLE IF NOT EXISTS bestellung (
    id INT AUTO_INCREMENT,
    zeitpunkt TIMESTAMP NOT NULL,
    nutzernr INT NOT NULL,
    FOREIGN KEY (nutzernr) REFERENCES nutzer(Nr),
    PRIMARY KEY (Id)
);

CREATE TABLE IF NOT EXISTS bild (
    id INT AUTO_INCREMENT,
    binaerdaten BLOB NOT NULL,
    alttext VARCHAR(60),
    titel VARCHAR(60),
    bildunterschrift VARCHAR(80),
    PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS kategorie (
    id INT AUTO_INCREMENT,
    bezeichnung VARCHAR(100),
    oberkategorie INT DEFAULT NULL,
    kategoriebild INT,
    CONSTRAINT `oberkat` FOREIGN KEY (oberkategorie) REFERENCES kategorie(id),
    FOREIGN KEY (kategoriebild) REFERENCES bild(id),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS produkt (
    id INT AUTO_INCREMENT,
    beschreibung TEXT NOT NULL,
    vegetarisch TINYINT(1),
    vegan TINYINT(1),
    produktbildId INT NOT NULL,
    kategorieId INT NOT NULL,
    FOREIGN KEY (produktbildId) REFERENCES bild(id),
    FOREIGN KEY (kategorieId) REFERENCES kategorie(id),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS trans_produkt_bestellung (
    produktid INT NOT NULL,
    bestellid INT NOT NULL,
    FOREIGN KEY (produktid) REFERENCES produkt(id),
    FOREIGN KEY (bestellid) REFERENCES bestellung(id)
);

CREATE TABLE IF NOT EXISTS preis (
    gastbetrag INT NOT NULL,
    studentenbetrag INT NOT NULL,
    mitarbeiterbetrag INT NOT NULL,
    produkt INT NOT NULL,
    FOREIGN KEY (produkt) REFERENCES produkt(id)
);

CREATE TABLE IF NOT EXISTS zutat (
    id INT AUTO_INCREMENT,
    glutenfrei TINYINT(1),
    bio TINYINT(1),
    vegetarisch TINYINT(1),
    vegan TINYINT(1),
    beschreibung TEXT,
    Name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS trans_produkt_zutat (
    produktId INT NOT NULL,
    zutatId INT NOT NULL,
    FOREIGN KEY (produktid) REFERENCES produkt(id),
    FOREIGN KEY (zutatid) REFERENCES zutat(id)
);

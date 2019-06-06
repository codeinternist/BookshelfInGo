START TRANSACTION;

CREATE TABLE IF NOT EXISTS library
(
    id              INT         NOT NULL AUTO_INCREMENT,
    title           VARCHAR(255),
    author          VARCHAR(255),
    publisher       VARCHAR(255),
    publish_date    TIMESTAMP,
    rating          TINYINT,
    status          TINYINT,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

COMMIT;
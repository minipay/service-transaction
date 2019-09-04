CREATE TABLE transactions (
	id INT(11) NOT NULL AUTO_INCREMENT,
	id_user bigint(20) NOT NULL,
    unique_id_transaction varchar(200) NOT NULL,
    type enum('in','out') NOT NULL,
    amount bigint(20) NOT NULL,
    description text NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(unique_id_transaction),
    PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;
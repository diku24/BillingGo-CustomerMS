use billingapplication;
DROP TABLE IF EXISTS customer;
CREATE TABLE customer (
    customer_id INT AUTO_INCREMENT NOT NULL,
    customer_name VARCHAR(128) NOT NULL,
    contact_number VARCHAR(255) NOT NULL,
    address VARCHAR(500) NOT NULL,
    prority VARCHAR(266) NOT NULL,
    primary key (`customer_id`)
);

select * from customer;

INSERT INTO `billingapplication`.`customer` (`customer_id`, `customer_name`, `contact_number`, `address`, `prority`) VALUES ('1', 'diku', '7020228551', 'Ahmednagar', '1');
CREATE DATABASE business;
USE business;
CREATE TABLE orders (id bigint(8) NOT NULL AUTO_INCREMENT PRIMARY KEY, storeId int(3), storeOrderId int(7), productId int(5), productName varchar(200), eachPrice float, quantity int, totalPrice float);
ALTER TABLE orders AUTO_INCREMENT=10000001;

#INSERT INTO orders (storeId, productId, productName, eachPrice, quantity, totalPrice) VALUES (101, 1001, "Pulsar side mirror", 105.34, 2, 210.68);
#SELECT id, storeId, storeOrderId, productId, productName, eachPrice, quantity, totalPrice FROM orders WHERE storeId=101;
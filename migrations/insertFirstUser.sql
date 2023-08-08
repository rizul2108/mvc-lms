-- +migrate Up

INSERT INTO users (username,full_name,hash,type) VALUES ("admin","iamadmin","$2a$10$dTT9W.1YqBhy760t54QUV.3ueDACZG/LbELi2PKaXE3PE3/83EYwW","admin");

-- +migrate Down

DELETE FROM users WHERE username="admin";
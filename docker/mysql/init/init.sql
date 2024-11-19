CREATE DATABASE IF NOT EXISTS normal_db;
CREATE DATABASE IF NOT EXISTS view_db;

CREATE USER 'app_user'@'%' IDENTIFIED BY 'dev';

GRANT ALL PRIVILEGES ON normal_db.* TO 'app_user'@'%';
GRANT ALL PRIVILEGES ON view_db.* TO 'app_user'@'%';

FLUSH PRIVILEGES;

USE normal_db;

-- tenants テーブルの作成
CREATE TABLE tenants
(
    id      INT AUTO_INCREMENT PRIMARY KEY,
    name    VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    tel     VARCHAR(20)
);

-- employees テーブルの作成
CREATE TABLE employees
(
    id        INT AUTO_INCREMENT PRIMARY KEY,
    name      VARCHAR(255) NOT NULL,
    email     VARCHAR(255) NOT NULL, -- emailをカラム名に設定
    tenant_id INT,
    FOREIGN KEY (tenant_id) REFERENCES tenants (id)
);

-- tenants テーブルにダミーデータを挿入
INSERT INTO tenants (name, address, tel)
VALUES ('Tenant A', '123 Main St, City A', '012-345-6789'),
       ('Tenant B', '456 Side St, City B', '098-765-4321');

-- employees テーブルにダミーデータを挿入
INSERT INTO employees (name, email, tenant_id)
VALUES ('Alice Smith', 'alice.smith@example.com', 1),
       ('Bob Jones', 'bob.jones@example.com', 1),
       ('Carol White', 'carol.white@example.com', 2),
       ('David Black', 'david.black@example.com', 2);

USE view_db;

CREATE FUNCTION get_target_tenant_id()
    RETURNS INT
    DETERMINISTIC
    RETURN @target_tenant_id;

CREATE VIEW employee_view AS
SELECT id, name, email, tenant_id
FROM normal_db.employees
WHERE tenant_id = get_target_tenant_id();


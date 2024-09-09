CREATE DATABASE project_asset;

CREATE TABLE employee (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100),
    gender VARCHAR(1),
    address VARCHAR(100),
    phone_number VARCHAR(15) UNIQUE  
);

CREATE TABLE asset_categories (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE asset_location (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE vendors (
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(100) NOT NULL,
    phone VARCHAR(100) NOT NULL,
);

CREATE TABLE asset (
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    category_id VARCHAR(100) NOT NULL,
    transaction_detail_id VARCHAR(100),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    image_url VARCHAR(100) NULL,
    qty INT NOT NULL,
    created_at DATE NOT NULL,
    CONSTRAINT fk_category_id FOREIGN KEY(category_id) REFERENCES asset_categories(id),
    CONSTRAINT fk_transaction_detail_id FOREIGN KEY(transaction_detail_id) REFERENCES transaction_detail(id)
);

CREATE TABLE asset_details (
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    asset_id VARCHAR(100) NOT NULL,
    location_id VARCHAR(100) NOT NULL,
    status int, 
    updated_at DATE NULL,
    removed_at DATE null,
    CONSTRAINT fk_asset_id FOREIGN KEY(asset_id) REFERENCES asset(id),
    CONSTRAINT fk_asset_loc_id FOREIGN KEY(location_id) REFERENCES asset_location(id)
);
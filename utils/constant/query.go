package constant

const (
	EMPLOYEE_INSERT = "INSERT INTO employee(id,name,gender,address,phone_number)VALUES($1, $2, $3, $4, $5)"
	EMPLOYEE_LIST   = "SELECT * FROM employee"
	EMPLOYEE_GET    = "SELECT * FROM employee where id=$1"
	EMPLOYEE_UPDATE = "UPDATE employee SET name=$1,gender=$2,phone_number=$3, address=$4 WHERE id=$5"
	EMPLOYEE_DELETE = "DELETE FROM employee WHERE id=$1"

	ASSET_CATEGORIES_INSERT = "INSERT INTO asset_categories(id,name)VALUES($1, $2)"
	ASSET_CATEGORIES_LIST   = "SELECT * FROM asset_categories"
	ASSET_CATEGORIES_GET    = "SELECT * FROM asset_categories where id=$1"
	ASSET_CATEGORIES_UPDATE = "UPDATE asset_categories SET name=$1 WHERE id=$2"
	ASSET_CATEGORIES_DELETE = "DELETE FROM asset_categories WHERE id=$1"

	ASSET_LOCATION_INSERT = "INSERT INTO asset_location(id, name) VALUES ($1, $2);"
	ASSET_LOCATION_LIST   = "SELECT id, name FROM asset_location;"
	ASSET_LOCATION_SEARCH = "SELECT id, name FROM asset_location WHERE id=$1;"
	ASSET_LOCATION_UPDATE = "UPDATE asset_location SET name=$2 WHERE id=$1;"
	ASSET_LOCATION_DELETE = "DELETE FROM asset_location WHERE id=$1;"
)

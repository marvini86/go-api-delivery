package customer

import (
	"log"

	"github.com/marvini86/go-api-delivery/internal/commom/entity"
	"github.com/marvini86/go-api-delivery/pkg/database"
)

func GetAll() (customers []entity.Customer, err error) {
	conn, err := database.OpenConnection()

	if err != nil {
		log.Fatalf("Error to open connection: %v", err)
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query(SQL_LIST_ALL)

	if err != nil {
		log.Fatalf("Error to fetch all data: %v", err)
		return
	}

	for rows.Next() {
		var customer entity.Customer
		_ = rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.Address.Id,
			&customer.Address.Address,
			&customer.Address.District)

		customers = append(customers, customer)
	}

	return
}

func Get(id int64) (customer entity.Customer, err error) {
	conn, err := database.OpenConnection()

	if err != nil {
		log.Fatalf("Error to open connection: %v", err)
		return
	}

	defer conn.Close()

	err = conn.QueryRow(SQL_BY_ID, id).Scan(
		&customer.Id,
		&customer.Name,
		&customer.Address.Id,
		&customer.Address.Address,
		&customer.Address.District)

	if err != nil {
		log.Fatalf("Error to fetch single data: %v", err)
		return
	}

	return
}

func Save(customer entity.Customer) (saved entity.Customer, err error) {
	conn, err := database.OpenConnection()

	if err != nil {
		log.Fatalf("Error to open connection: %v", err)
		return
	}

	defer conn.Close()

	err = conn.QueryRow(SQL_SAVE_CUSTOMER, customer.Name).Scan(&saved.Id)

	if err != nil {
		log.Fatalf("Error to save customer found: %v", err)
		return
	}

	_, err = conn.Exec(SQL_SAVE_ADDRESS, customer.Address.Address, customer.Address.District, saved.Id)

	if err != nil {
		log.Fatalf("Error to save address: %v", err)
		return
	}

	err = conn.QueryRow(SQL_BY_ID, saved.Id).Scan(&saved.Id, &saved.Name, &saved.Address.Id, &saved.Address.Address, &saved.Address.District)

	if err != nil {
		log.Fatalf("Error to retrieve data: %v", err)
		return
	}

	return
}

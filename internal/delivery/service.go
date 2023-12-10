package delivery

import (
	"log"

	"github.com/marvini86/go-api-delivery/internal/commom/entity"
	"github.com/marvini86/go-api-delivery/pkg/database"
)

func GetAll() (deliveries []entity.Delivery, err error) {
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
		var delivery entity.Delivery
		_ = rows.Scan(
			&delivery.Id, &delivery.CreatedDate,
			&delivery.Customer.Id,
			&delivery.Customer.Name,
			&delivery.Customer.Address.Id,
			&delivery.Customer.Address.Address,
			&delivery.Customer.Address.District)

		items, _ := GetItems(delivery.Id)
		delivery.Items = items

		deliveries = append(deliveries, delivery)
	}

	return
}

func Get(id int64) (delivery entity.Delivery, err error) {
	conn, err := database.OpenConnection()

	if err != nil {
		log.Fatalf("Error to open connection: %v", err)
		return
	}

	defer conn.Close()

	err = conn.QueryRow(SQL_BY_ID, id).Scan(
		&delivery.Id, &delivery.CreatedDate,
		&delivery.Customer.Id,
		&delivery.Customer.Name,
		&delivery.Customer.Address.Id,
		&delivery.Customer.Address.Address,
		&delivery.Customer.Address.District)

	items, _ := GetItems(delivery.Id)
	delivery.Items = items

	if err != nil {
		log.Fatalf("Error to fetch single data: %v", err)
		return
	}

	return
}

func Save(delivery entity.Delivery) (saved entity.Delivery, err error) {
	conn, err := database.OpenConnection()

	if err != nil {
		log.Fatalf("Error to open connection: %v", err)
		return
	}

	defer conn.Close()

	var receiver entity.Customer

	err = conn.QueryRow(SQL_RECEIVER, delivery.Customer.Id).Scan(&receiver.Id, &receiver.Name)

	if err != nil {
		log.Fatalf("No customer found: %v", err)
		return
	}

	var id = 0

	err = conn.QueryRow(SQL_SAVE_DELIVERY, receiver.Id).Scan(&id)

	if err != nil {
		log.Fatalf("Error to save data: %v", err)
		return
	}

	for _, item := range delivery.Items {
		_, _ = conn.Exec(SQL_SAVE_ITEM, id, item.Description, item.Quantity)
	}

	err = conn.QueryRow(SQL_BY_ID, id).Scan(
		&saved.Id, &delivery.CreatedDate,
		&saved.Customer.Id,
		&saved.Customer.Name,
		&saved.Customer.Address.Id,
		&saved.Customer.Address.Address,
		&saved.Customer.Address.District)

	items, _ := GetItems(saved.Id)
	saved.Items = items

	if err != nil {
		log.Fatalf("Error to fetch single data: %v", err)
		return
	}

	return
}

func GetItems(deliveryId int64) (items []entity.Item, err error) {
	conn, err := database.OpenConnection()

	if err != nil {
		log.Fatalf("Error to open connection: %v", err)
		return
	}

	defer conn.Close()

	rows, err := conn.Query(SQL_ITEMS, deliveryId)

	if err != nil {
		log.Fatalf("Error to retieve data: %v", err)
		return
	}

	for rows.Next() {
		var item entity.Item
		_ = rows.Scan(&item.Id, &item.Description, &item.Quantity)

		items = append(items, item)
	}

	return
}

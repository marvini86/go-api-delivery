package customer

const SQL_LIST_ALL = `select c.id, c.name, a.id, 
							a.address, a.district  
						from customer c
						join address a on c.id = a.customer_id `

const SQL_BY_ID = `select c.id, c.name, a.id, 
						a.address, a.district  
						from customer c
						join address a on c.id = a.customer_id
						where c.id = $1`

const SQL_SAVE_CUSTOMER = `INSERT INTO customer(name) VALUES($1) RETURNING id`

const SQL_SAVE_ADDRESS = `INSERT INTO address(address, district, customer_id) VALUES($1, $2, $3)`

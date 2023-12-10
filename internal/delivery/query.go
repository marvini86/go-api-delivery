package delivery

const SQL_LIST_ALL = `select d.id, d.created_date, 
							r.id, r.name, a.id, 
							a.address, a.district  
						from delivery d
						join receiver r on r.id = d.receiver_id 
						join address a on r.id = a.receiver_id `

const SQL_BY_ID = `select d.id, d.created_date, 
							r.id, r.name, a.id, 
							a.address, a.district  
						from delivery d
						join receiver r on r.id = d.receiver_id 
						join address a on r.id = a.receiver_id
						where d.id = $1`

const SQL_SAVE_DELIVERY = `INSERT INTO delivery(receiver_id, created_date) VALUES($1, now()) RETURNING id`

const SQL_SAVE_ITEM = `INSERT INTO item(delivery_id, description, quantity) VALUES($1, $2, $3)`

const SQL_ITEMS = `SELECT id, description, quantity FROM item where delivery_id = $1`

const SQL_RECEIVER = `SELECT id, name FROM receiver where id = $1`

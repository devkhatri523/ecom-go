package query

const ORDER_QUERY = "select a.Id,a.name as product_name,a.description as product_description ,a.available_quantity as available_quantity," +
	"a.price as price ,b.Id as category_id ,b.name as category_name,b.description as category_description from product a join category b on  a.category_id = b.id "

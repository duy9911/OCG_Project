SELECT products.id, products.name, products.description, product_variances.color, product_variances.size, images.image
FROM products
INNER JOIN product_variances 
ON products.id=product_variances.product_id 
INNER JOIN images ON products.id=images.product_id
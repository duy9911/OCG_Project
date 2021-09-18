SELECT pages.id, pages.name, collections.name
FROM pages INNER JOIN collections 
ON pages.id=collections.page_id  WHERE pages.id=1
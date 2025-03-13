SELECT 
    u.id AS user_id,
    c.id AS company_id,
    u.nama,
    u.email,
    u.telp,
    c.company_code,
    c.company_name
FROM
    users u
LEFT JOIN
    company c
ON
    u.id = c.user_id;

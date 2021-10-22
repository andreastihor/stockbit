SELECT Parent.id    AS "ID",
       Parent.NAME  AS "UserName",
       Child."name" AS "ParentUserName"
FROM   USER AS Parent
       LEFT JOIN USER AS Child
              ON Parent.parent_id = Child.id
ORDER  BY "ID" 
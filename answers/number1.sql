SELECT Parent.id    AS "ID",
       Parent.NAME  AS "UserName",
       Child."name" AS "ParentUserName"
FROM   USER AS Child
       RIGHT JOIN USER AS Parent
              ON Parent.parent_id = Child.id
ORDER  BY "ID" ;
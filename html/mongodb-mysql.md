查询： 
MySQL: 
SELECT * FROM user 
Mongo: 
db.user.find() 

MySQL: 
SELECT * FROM user WHERE name = ’starlee’ 
Mongo: 
db.user.find({‘name’ : ’starlee’}) 

插入： 
MySQL: 
INSERT INOT user (`name`, `age`) values (’starlee’,25) 
Mongo: 
db.user.insert({‘name’ : ’starlee’, ‘age’ : 25}) 

如果你想在MySQL里添加一个字段，你必须： 
ALTER TABLE user…. 
但在MongoDB里你只需要： 
db.user.insert({‘name’ : ’starlee’, ‘age’ : 25, ‘email’ : ’starlee@starlee.com’}) 

删除： 
MySQL: 
DELETE * FROM user 
Mongo: 
db.user.remove({}) 

MySQL: 
DELETE FROM user WHERE age < 30 
Mongo: 
db.user.remove({‘age’ : {$lt : 30}}) 

$gt : > ; $gte : >= ; $lt : < ; $lte : <= ; $ne : != 

更新: 

MySQL: 
UPDATE user SET `age` = 36 WHERE `name` = ’starlee’ 
Mongo: 
db.user.update({‘name’ : ’starlee’}, {$set : {‘age’ : 36}}) 

MySQL: 
UPDATE user SET `age` = `age` + 3 WHERE `name` = ’starlee’ 
Mongo: 
db.user.update({‘name’ : ’starlee’}, {$inc : {‘age’ : 3}}) 

MySQL: 
SELECT COUNT(*) FROM user WHERE `name` = ’starlee’ 
Mongo: 
db.user.find({‘name’ : ’starlee’}).count() 

MySQL: 
SELECT * FROM user limit 10,20 
Mongo: 
db.user.find().skip(10).limit(20) 

MySQL: 
SELECT * FROM user WHERE `age` IN (25, 35,45) 
Mongo: 
db.user.find({‘age’ : {$in : [25, 35, 45]}}) 

MySQL: 
SELECT * FROM user ORDER BY age DESC 
Mongo: 
db.user.find().sort({‘age’ : -1}) 

MySQL: 
SELECT DISTINCT(name) FROM user WHERE age > 20 
Mongo: 
db.user.distinct(‘name’, {‘age’: {$lt : 20}}) 

MySQL: 
SELECT name, sum(marks) FROM user GROUP BY name 
Mongo: 
db.user.group({ 
key : {‘name’ : true}, 
cond: {‘name’ : ‘foo’}, 
reduce: function(obj,prev) { prev.msum += obj.marks; }, 
initial: {msum : 0} 
}); 

MySQL: 
SELECT name FROM user WHERE age < 20 
Mongo: 
db.user.find(‘this.age < 20′, {name : 1}) 

发现很多人在搜MongoDB循环插入数据，下面把MongoDB循环插入数据的方法添加在下面： 

for(var i=0;i<100;i++)db.test.insert({uid:i,uname:’nosqlfan’+i}); 

上面一次性插入一百条数据，大概结构如下： 
{ “_id” : ObjectId(“4c876e519e86023a30dde6b8″), “uid” : 55, “uname” : “nosqlfan55″ } 
{ “_id” : ObjectId(“4c876e519e86023a30dde6b9″), “uid” : 56, “uname” : “nosqlfan56″ } 
{ “_id” : ObjectId(“4c876e519e86023a30dde6ba”), “uid” : 57, “uname” : “nosqlfan57″ } 
{ “_id” : ObjectId(“4c876e519e86023a30dde6bb”), “uid” : 58, “uname” : “nosqlfan58″ } 
{ “_id” : ObjectId(“4c876e519e86023a30dde6bc”), “uid” : 59, “uname” : “nosqlfan59″ } 
{ “_id” : ObjectId(“4c876e519e86023a30dde6bd”), “uid” : 60, “uname” : “nosqlfan60″ } 


简易对照表 
SQL Statement                                                  Mongo Query Language Statement 
CREATE TABLE USERS (a Number, b Number)         implicit; can be done explicitly 

INSERT INTO USERS VALUES(1,1)                             db.users.insert({a:1,b:1}) 
SELECT a,b FROM users                                           db.users.find({}, {a:1,b:1}) 
SELECT * FROM users                                              db.users.find() 
SELECT * FROM users WHERE age=33                      db.users.find({age:33}) 
SELECT a,b FROM users WHERE age=33                   db.users.find({age:33}, {a:1,b:1}) 
SELECT * FROM users WHERE age=33 ORDER BY name                db.users.find({age:33}).sort({name:1}) 
SELECT * FROM users WHERE age>33                     db.users.find({'age':{$gt:33}})}) 
SELECT * FROM users WHERE age<33                     db.users.find({'age':{$lt:33}})}) 
SELECT * FROM users WHERE name LIKE "%Joe%"                                   db.users.find({name:/Joe/}) 
SELECT * FROM users WHERE name LIKE "Joe%"                               db.users.find({name:/^Joe/}) 
SELECT * FROM users WHERE age>33 AND age<=40                                   db.users.find({'age':{$gt:33,$lte:40}})}) 
SELECT * FROM users ORDER BY name DESC                                   db.users.find().sort({name:-1}) 
CREATE INDEX myindexname ON users(name)                                   db.users.ensureIndex({name:1}) 
CREATE INDEX myindexname ON users(name,ts DESC)                                   db.users.ensureIndex({name:1,ts:-1}) 
SELECT * FROM users WHERE a=1 and b='q'                                   db.users.find({a:1,b:'q'}) 
SELECT * FROM users LIMIT 10 SKIP 20                                   db.users.find().limit(10).skip(20) 
SELECT * FROM users WHERE a=1 or b=2                          db.users.find( { $or : [ { a : 1 } , { b : 2 } ] } ) 
SELECT * FROM users LIMIT 1                                          db.users.findOne() 
EXPLAIN SELECT * FROM users WHERE z=3                                   db.users.find({z:3}).explain() 
SELECT DISTINCT last_name FROM users                                   db.users.distinct('last_name') 
SELECT COUNT(*y) FROM users                                            db.users.count() 
SELECT COUNT(*y) FROM users where AGE > 30                             db.users.find({age: {'$gt': 30}}).count() 
SELECT COUNT(AGE) from users                                       db.users.find({age: {'$exists': true}}).count() 
UPDATE users SET a=1 WHERE b='q'                                   db.users.update({b:'q'}, {$set:{a:1}}, false, true) 
UPDATE users SET a=a+2 WHERE b='q'                                   db.users.update({b:'q'}, {$inc:{a:2}}, false, true) 
DELETE FROM users WHERE z="abc"                                    db.users.remove({z:'abc'}); 

### 活动创建 
商户（admin）发布秒杀活动，根据时间，admin_id和商品类型生成随机活动id，并返回给前端
### 创建订单
用户进行抢购，利用channel对查询和创建的操作加锁，获取到活动id，在redis中执行事务，查询&更新redis中的商品库存，成功生成20s(这里是便于测试)支付的订单，并返回订单号，释放锁
ps：此处首次获取库存数有误
### 订单
查询订单是否在规定时间支付，未支付将商品加入库存
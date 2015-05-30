经纬度为uint64

Http APIs   
1. 地理位置报告接口，返回所有玩家信息和豆子信息
URL: /player/report
Params: id=123&longitude=117&latitude=34

Return: 
{
    "errno": 0,
     "errmsg": "",
     "players": [
     {
         "id": 123,
         "longitude": 117.01,
         "latitude": 23.1
     },
     {
         "id": 456,
         "longitude": 117.01,
         "latitude": 23.1
     }
     ],
     "beans": [
     {
         "id": 1,
         "state": 1,
         "longitude": 117.01,
         "latitude": 23.1
     },
     {
         "id": 2,
         "state": 1,
         "longitude": 117.01,
         "latitude": 23.1
     }
     ]
}

2. create/update 豆子接口
URL: /bean/manipulate
Params: id=123&state=0&longitude=117&latitude=34

Return: 
{
    "errno": 0,
     "errmsg": "",
     "players": [
     {
         "id": 123,
         "longitude": 117.01,
         "latitude": 23.1
     },
     {
         "id": 456,
         "longitude": 117.01,
         "latitude": 23.1
     }
     ],
     "beans": [
     {
         "id": 1,
         "state": 1,
         "longitude": 117.01,
         "latitude": 23.1
     },
     {
         "id": 2,
         "state": 1,
         "longitude": 117.01,
         "latitude": 23.1
     }
     ]
}

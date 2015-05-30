struct Player{
    1: required i64 id;
    2: required string name;
    3: optional Point location;
}

exception BeanException{
    1: i32 errNo;
    2: string errMsg;
}

struct Game {
    1: required i64 id;            
    2: required string name;            
    3: required i64 hostPlayerId;
    4: required Rectangle rect;
    5: required i32 players;
    6: required i32 maxPlayers;
}

struct Point {
    1: required i64 longitude;
    2: required i64 latitude;
}

struct Rectangle {
    1: required Point pointMin;
    2: required Point pointMax;
}

service GameManagerService {
    //CreateGame returns game id
    i64 CreateGame(1:Player player, 2:i32 maxPlayers, 3:i32 cityId, 4:Rectangle rect) throws (1: BeanException excep),

    //List games which belong to this cityId
    list<Game> ListGames(1:i32 cityId) throws (1: BeanException excep),

    //Usage:
    //1. Join the game.
    //2. If already joined the game, you could query players status.
    //3. Throws the exception if the game quit(errno = 1).
    //4. Throws the exception if the game starts(errno = 2).
    list<Player> JoinGame(1:i64 gameId) throws (1: BeanException excep)
}

service PlayerService {
    //Report my location and returns other players' location
    list<Player> Report(1: i64 game, 2: Player me) throws(1: BeanException excep)
    list<Point> ListBeans(1: i64 game) throws(1: BeanException excep)
} 

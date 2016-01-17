var io = require('socket.io')();
var players = new Array();
io.on('connection', function(socket) {
    socket.emit('getId', socket.conn.id);
    console.log(socket.conn.id);
    socket.on('newPlayer', function(data) {
        data.id = socket.conn.id;
        socket.broadcast.emit('newPlayer', data);
        socket.emit('dump', players);
        players.push(data);
        
        
    });
    socket.on('update', function(data) {
        data.id = socket.conn.id;
        console.log(data)
        socket.broadcast.emit('update', data);
        
    });
    socket.on('disconnect', function() {
        for (var i = players.length - 1; i >= 0; i--) {
        	if(players[i].id == socket.id){
        		scoket.broadcast.emit('death', socket.id)
        		players.splice(i);
        	}
        };
    });
});


io.listen(3000);

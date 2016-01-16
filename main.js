var io = require('socket.io')();
var players;
io.on('connection', function(socket) {
            socket.emit('getId', socket.conn.id);
            console.log(socket.conn.id);
            socket.on('newPlayer', function(data) {
            	data.id = socket.conn.id;
            	console.log(data);
                });
            });

        io.listen(3000);

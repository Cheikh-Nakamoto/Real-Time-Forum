let socket;

const connectWebSocket = () => {
    const token = localStorage.getItem('token');
    if (!token) {
        alert('You must be logged in to chat.');
        return;
    }

    socket = new WebSocket(`ws://localhost:${location.port}/ws?token=${token}`);

    socket.onmessage = function(event) {
        const msg = JSON.parse(event.data);
        const li = document.createElement('li');
        li.textContent = `From ${msg.sender_id}: ${msg.content}`;
        document.getElementById('messages').appendChild(li);
    };

    socket.onclose = function(event) {
        alert('WebSocket connection closed.');
    };
}
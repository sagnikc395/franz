import socket

def send_message(topic, message, host='localhost', port=5000):
    with socket.socket() as s:
        s.connect((host, port))
        s.sendall(f"PRODUCE {topic} {message}".encode())

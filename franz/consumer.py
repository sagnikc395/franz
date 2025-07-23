import socket

def consume_message(topic, offset=0, host='localhost', port=5000):
    with socket.socket() as s:
        s.connect((host, port))
        s.sendall(f"CONSUME {topic} {offset}".encode())
        data = s.recv(4096).decode()
        return data

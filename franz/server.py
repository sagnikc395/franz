import socket
from threading import Thread
from broker import Broker

broker = Broker()
broker.create_topic('demo')

def handle_client(conn):
    with conn:
        request = conn.recv(1024).decode()
        parts = request.split()
        if parts[0] == 'PRODUCE':
            _, topic, msg = parts[0], parts[1], ' '.join(parts[2:])
            broker.produce(topic, msg)
            conn.send(b'OK')
        elif parts[0] == 'CONSUME':
            _, topic, offset = parts
            msgs, new_offset = broker.consume(topic, int(offset))
            response = ''.join(msgs)
            conn.send(response.encode())

def run_server(host='localhost', port=5000):
    with socket.socket() as s:
        s.bind((host, port))
        s.listen()
        print("Broker running...")
        while True:
            conn, _ = s.accept()
            Thread(target=handle_client, args=(conn,)).start()



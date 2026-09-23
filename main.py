import socket

target = input("Enter IP address: ")

# TODO - Implement port scanning logic (manually supply the range of ports to scan or use a list of common ports)
for port in range(1, 1025):
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.settimeout(0.2)

    result = sock.connect_ex((target, port))

    if result == 0:
        print(f"Port {port} is OPEN")

    sock.close()
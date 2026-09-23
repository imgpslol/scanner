## How to Run

### 1. Navigate to the Project

Open **WSL** or **PowerShell** and navigate to the directory containing `main.py`.

**Example:**

    cd path/to/project

### 2. Run the Program

    python3 ./main.py

> On some Windows installations, you may need to use `python` instead of `python3`.

### 3. Enter an IP Address

When prompted, enter the IP address you want to scan.

**Example:**

    Enter IP address: 127.0.0.1

The scanner will then check the first **1024 TCP ports** and report any ports that are open.

> **Important:** Only scan systems you own or have permission to test.

---

## Current Features

- [x] Scan TCP ports 1–1024
- [x] Detect open ports
- [x] Accept an IP address from the user
- [x] Basic connection timeout

---

## TODO

### Basic Improvements

- [ ] Allow the user to choose the port range
- [ ] Display a message when no open ports are found
- [ ] Add better error handling for invalid IP addresses
- [ ] Add a progress indicator
- [ ] Display how long the scan took

### Networking

- [ ] Scan a hostname instead of only an IP address
- [ ] Identify common services running on open ports
- [ ] Add support for scanning multiple IP addresses
- [ ] Add configurable connection timeouts

### Performance

- [ ] Use threads to scan multiple ports simultaneously
- [ ] Compare threaded vs. non-threaded scan times

### Output

- [ ] Add a `--help` option
- [ ] Add command-line arguments
- [ ] Save scan results to a file
- [ ] Add a cleaner terminal output format

### Later / Advanced

- [ ] Detect the service/version running on common ports
- [ ] Add UDP scanning
- [ ] Add a simple configuration file
- [ ] Add automated tests
- [ ] Package the scanner as a proper command-line application

---


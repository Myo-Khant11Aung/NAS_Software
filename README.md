
# Personal NAS — Go Systems Engineering

A self-hosted **Network Attached Storage (NAS)** and personal cloud storage system built using **Go and Linux**.

The goal is to provide secure, reliable access to files stored on a central Linux server from Linux, macOS, and Windows devices, both over local networks and remotely through the internet.

This project is developed from scratch as a hands-on systems software engineering learning experience, focusing on networking, concurrency, file storage, security, reliability, and software architecture.

**Project Status:** Planning & Foundations — Active Development

---

## Project Objectives

- Build a functional, reliable, self-hosted NAS system.
- Access files securely from anywhere with an internet connection.
- Support Linux, macOS, and Windows clients.
- Implement file uploads, downloads, and directory management.
- Handle multiple client connections concurrently.
- Ensure file integrity and prevent data corruption.
- Implement authentication and authorization.
- Handle interrupted connections and unexpected failures.
- Learn Linux systems programming and network communication.
- Develop professional software testing, debugging, and documentation practices.
- Create a portfolio project demonstrating practical systems software engineering skills.

---

## Technology Stack

| Component | Technology |
|-----------|------------|
| Core Server | Go |
| Server Operating System | Linux |
| Client Platforms | Linux, macOS, Windows |
| Networking | TCP/IP, HTTP where appropriate |
| Remote Access | Tailscale (planned) |
| Unit Testing | Go `testing` |
| Integration Testing | Python + pytest |
| End-to-End Testing | Python + pytest |
| Version Control | Git & GitHub |
| Continuous Integration | GitHub Actions |
| Deployment | Linux systemd |

Additional technologies will be introduced only when justified by project requirements.

---

## System Architecture

The system follows a **client-server architecture**.

A Linux computer acts as the central storage server. Linux, macOS, and Windows clients connect to it through a local network or a secure remote connection.

### Planned Architecture

```text
                Linux / macOS / Windows Clients
                             |
                   +---------+---------+
                   |                   |
              Local Network       Remote Internet
                   |                   |
                   |             Encrypted VPN
                   |              (Tailscale)
                   |                   |
                   +---------+---------+
                             |
                        Go NAS Server
                             |
              +--------------+--------------+
              |              |              |
           Network       Authentication    Storage
           Handler       & Authorization   Manager
              |              |              |
              +--------------+--------------+
                             |
                      Linux Filesystem
                             |
                          HDD / SSD
```

The Go server will be responsible for:

- Network communication
- File operations
- Authentication and authorization
- Concurrent request handling
- Data integrity verification
- Error handling and recovery

Linux will provide the underlying operating system and filesystem capabilities.

The project will not attempt to recreate an operating system, filesystem, or SMB implementation from scratch.

---

## Planned Features

### Core File Management

- [ ] Upload files to the NAS
- [ ] Download files from the NAS
- [ ] List files and directories
- [ ] Create and manage directories
- [ ] Delete files and directories
- [ ] Support large file transfers through streaming

### Networking and Concurrency

- [ ] Implement network communication between clients and server
- [ ] Define a reliable client-server communication protocol
- [ ] Handle multiple simultaneous client connections
- [ ] Implement connection timeouts and resource limits
- [ ] Support graceful server shutdown

### Security

- [ ] Implement user authentication
- [ ] Enforce file access permissions
- [ ] Validate requests and file paths
- [ ] Prevent unauthorized filesystem access
- [ ] Protect network communication
- [ ] Support secure remote access through a VPN

### Reliability and Data Integrity

- [ ] Verify transferred file integrity using checksums
- [ ] Handle interrupted uploads and downloads
- [ ] Prevent incomplete files from replacing valid data
- [ ] Handle disk and network errors safely
- [ ] Implement logging and error reporting
- [ ] Test recovery behavior after failures

### Future Enhancements

- [ ] Native network-drive mounting experience on Linux, macOS, and Windows
- [ ] File synchronization between devices
- [ ] Web-based file management interface
- [ ] Storage usage monitoring
- [ ] Backup and recovery management
- [ ] Additional remote-access options

---

## Development Roadmap

The project follows an incremental development process organized into six phases.

### Phase 1 — Foundations

**Topics:** Linux fundamentals, Go programming, filesystems, networking concepts, Git, and testing.

**Deliverables:**

- [ ] Configure development environments
- [ ] Initialize the GitHub repository
- [ ] Review Linux processes, files, and permissions
- [ ] Practice Go file I/O and error handling
- [ ] Understand TCP/IP and client-server architecture

### Phase 2 — Network Server

**Topics:** TCP sockets, client-server communication, network protocols, and file streaming.

**Deliverables:**

- [ ] Implement a basic TCP server and client
- [ ] Establish network connections
- [ ] Design an application communication protocol
- [ ] Transfer data between computers
- [ ] Write initial Python integration tests

### Phase 3 — File Storage

**Topics:** Filesystem APIs, directory management, streaming, data integrity, and safe file operations.

**Deliverables:**

- [ ] Implement file uploads and downloads
- [ ] Add directory listing and management
- [ ] Verify file integrity
- [ ] Prevent path traversal vulnerabilities
- [ ] Handle interrupted file transfers safely

### Phase 4 — Concurrency and Reliability

**Topics:** Goroutines, synchronization, resource management, fault tolerance, and performance.

**Deliverables:**

- [ ] Support multiple concurrent clients
- [ ] Prevent race conditions
- [ ] Implement connection and resource limits
- [ ] Add graceful shutdown
- [ ] Test failure and recovery scenarios

### Phase 5 — Security and Remote Access

**Topics:** Authentication, authorization, secure communication, VPN networking, and threat modeling.

**Deliverables:**

- [ ] Implement authentication and access controls
- [ ] Validate network requests and file paths
- [ ] Secure communication between clients and server
- [ ] Configure private remote access through Tailscale
- [ ] Test access from outside the local network

### Phase 6 — Integration and Deployment

**Topics:** Cross-platform compatibility, automated testing, Linux deployment, benchmarking, and documentation.

**Deliverables:**

- [ ] Test Linux, macOS, and Windows clients
- [ ] Deploy the NAS as a Linux systemd service
- [ ] Configure automated testing with GitHub Actions
- [ ] Measure transfer speeds and resource usage
- [ ] Complete technical documentation
- [ ] Demonstrate local and secure remote file access

---

## Testing Strategy

Testing will be a core part of the development process.

### Go Unit Tests

Go's built-in `testing` package will verify individual functions and internal components.

Tests will cover:

- Normal behavior
- Edge cases
- Error handling
- Relevant concurrency scenarios

### Python Integration Tests

Python and pytest will simulate real clients communicating with the NAS server through its public network interfaces.

Integration tests will verify:

- Client connections
- File uploads and downloads
- File integrity
- Invalid requests
- Authentication and permissions
- Concurrent connections
- Network interruptions and timeouts

### End-to-End Tests

End-to-end tests will validate complete file management workflows from a client's perspective.

### Reliability and Performance Testing

Testing will include:

- Go race detection
- Concurrent request simulation
- Transfer speed measurements
- CPU and memory usage
- Server stability under load
- Failure recovery behavior

All destructive tests will operate within isolated test directories to protect personal data.

---

## Development Environment

Development will take place across multiple computers.

| Device | Responsibility |
|--------|----------------|
| MacBook | Go development, unit testing, client testing |
| ThinkPad | Go development, Linux and Windows client compatibility testing |
| Linux Desktop | NAS server hosting and deployment |
| GitHub | Source control and continuous integration |

Go's cross-platform capabilities allow client development and testing on Linux, macOS, and Windows while targeting Linux for server deployment.

Python integration and end-to-end tests will be maintained separately.

---

## Engineering Principles

This project follows professional software engineering practices.

- **Correctness:** Prioritize predictable and correct behavior.
- **Data Integrity:** Never sacrifice stored data for performance or convenience.
- **Reliability:** Handle failures without unnecessary data loss.
- **Security:** Apply authentication, authorization, and secure communication.
- **Maintainability:** Keep the architecture understandable and modular.
- **Testability:** Design components with automated testing in mind.
- **Documentation:** Record important design choices and limitations.
- **Performance:** Measure and improve performance based on evidence.

Established security libraries and protocols will be used instead of implementing custom cryptographic algorithms.

Experimental services will not be directly exposed to the public internet.

The NAS must not be used as the only copy of important personal files.

---

## Learning Objectives

Beyond building a functional NAS, this project aims to develop practical understanding of:

- Linux operating systems and system interfaces
- TCP/IP networking
- Client-server architecture
- Filesystem operations
- Concurrent programming in Go
- Network security
- Data integrity and failure recovery
- Automated testing
- Debugging and performance profiling
- Software architecture
- Deployment and system administration

The student is responsible for designing and implementing the software independently, using documentation, experimentation, and code review to guide development.

---

## Project Status

**Current Stage:** Planning & Foundations

The architecture, implementation, and feature set will evolve throughout development.

Milestone progress, testing results, and technical decisions will be documented as the project advances.

---

## License

To be determined.

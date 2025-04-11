# GoLang Date-Time Web App with Docker & Kubernetes

A simple web application written in Go that displays the current date and time. This project demonstrates containerization using Docker, and deployment on Kubernetes with 2 replicas, exposed to the internet via a `NodePort` service.


##  Stack

- Go (Golang)
- Docker
- Kubernetes (Minikube)
- YAML (K8s declarative configs)

##  Features

- Displays current date and time
- Dockerized for easy deployment
- Kubernetes deployment with 2 replicas
- Exposed to the internet via `NodePort`

---

##  Project Structure

```
.
├── Dockerfile
├── main.go
├── go.mod
├── deployment.yaml
└── service.yaml
```

---

## Setup Instructions

### 1. Clone the Repo

```bash
git clone https://github.com/diamachi/datetime-app.git
cd datetime-app
```

### 2. Run Locally (without Docker)

```bash
go run main.go
```

Access at: `http://localhost:8080`

---

### 3. Build Docker Image

```bash
docker build -t diamachi/datetime-app:latest .
```

### 4. Push to DockerHub

```bash
docker login
docker push diamachi/datetime-app:latest
```

---

## ☸️ Deploy on Kubernetes

### 5. Start Minikube

```bash
minikube start
```

### 6. Apply Kubernetes Resources

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

This deploys 2 replicas of the app and exposes it on a NodePort (e.g., `30080`).

---

### 7. Access the App

```bash
minikube service datetime-service --url
```

Visit the returned URL in your browser — should look like:

```
http://192.168.49.2:30080
```

---

## 🧹 Cleanup

To stop and clean everything:

```bash
kubectl delete deployment datetime-app
kubectl delete service datetime-service
minikube stop
```

---

## ✨ Output Example

```
Current Date and Time: 2025-04-11 17:09:30
```


## 📦 DockerHub

➡️ [DockerHub: diamachi/datetime-app](https://hub.docker.com/repository/docker/diamachi/datetime-app/general)



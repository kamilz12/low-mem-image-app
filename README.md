
# Achieving the Goal of Image Size Minimization

The main objective of this project was to create the smallest possible Docker image for a server application written in Go. To achieve this, several optimization techniques were applied:

- **Static linking** of the Go binary during compilation,  
- Use of a **multi-stage build** process,  
- Deployment of a minimal base image such as `scratch` or `alpine`,  
- **Removal of unnecessary dependencies and auxiliary files**.

As a result of these efforts, the final Docker image size was reduced to just **2.53 MB**:

![image](https://github.com/user-attachments/assets/8d85b8e5-52b2-43ff-85e9-2dbcb131e8b7)

This is significantly smaller than typical container images and demonstrates the effectiveness of the optimization strategy.

Such a minimal image size brings several practical advantages:

- **Faster application downloads and deployments**,  
- **Lower disk and network resource usage**,  
- **Improved security** by reducing the potential attack surface.

The image was built using the command:

```bash
docker buildx build --platform linux/amd64 --no-cache -t kziolkowski/weather-go:latest --load .
```


In conclusion, the final result fully meets the original goal – delivering a **lightweight, portable, and production-ready containerized application** with an impressively small footprint.

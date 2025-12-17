FROM python:latest

WORKDIR /app
COPY . .

# BAD: No requirements pinning
RUN pip install flask

# BAD: No USER nonroot

CMD ["python", "app.py"]

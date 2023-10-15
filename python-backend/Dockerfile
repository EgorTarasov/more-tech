FROM python:3.11.5-bullseye

# Select working directory
WORKDIR /code

# Copy requirements.txt to working directory
COPY requirements.txt requirements.txt


# Install dependencies
RUN pip install -r requirements.txt

# Copy source code to working directory
COPY . /code

# Create data directory
RUN mkdir -p /data/logs

# Run the application
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8001"]
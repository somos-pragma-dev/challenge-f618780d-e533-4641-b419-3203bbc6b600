version: '3.8'

services:
  auth-service:
    build:
      context:./auth-service
    ports:
      - '50051:50051'
  transaction-service:
    build:
      context:./transaction-service
    ports:
      - '50052:50052'
  notification-service:
    build:
      context:./notification-service
    ports:
      - '50053:50053'

#1.
FROM alpine:latest

#2. Run command mkdir on the new small docker image
RUN mkdir /app

#3. copy
COPY brokerApp /app

#4. Run the command
CMD [ "/app/brokerApp" ]


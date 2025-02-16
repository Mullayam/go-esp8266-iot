# ReSmArt: Recycle. Smartify. Automate.

## ESP8266 WebSocket IoT Controller with Go, MongoDB, and Google Nest Mini
This project can Turn your old phone into a smart camera and seamlessly automate  wiht real-time IoT automation using ESP8266, Google Nest Mini, and WiFi Smart Devices (like Smitch Smart Lights and motors). It integrates a Go server with MongoDB logging and a WebSocket-based control system for seamless smart home management.

ReSMart breathes new life into outdated tech, transforming it into a real-time IoT powerhouse. Recycle, Smartify, Automate.
### Features
- ✅ Turn your old phone into a smart camera 
- ✅ Real-time WebSocket Communication between ESP8266 and Go server
- ✅ MongoDB Logging for tracking all commands, sensor data, and events
- ✅ Relay & Motor Control via WebSocket commands
- ✅ WiFi Smart Light (Smitch) Control using Go integration
- ✅ Google Nest Mini Voice Commands to control devices (SDK)
- ✅ Sensor Monitoring (motion detection, light levels using photo diodes)

### How It Works
- ESP8266 runs a WebSocket server, listens for commands, and sends sensor data.
- Go server acts as a WebSocket client, sending control commands and receiving sensor data.
- MongoDB logs all activity, including incoming/outgoing messages and device states.
- Google Nest Mini is integrated for voice-based control of relays, lights, and motors.
- Smitch WiFi Smart Light & Motor are controlled via HTTP/WebSocket commands from the Go server.

###  Smart Device Integration
- ESP8266: Controls relays, motors, and reads sensor data.
- Google Nest Mini: Listens to voice commands and triggers API requests.
- Smitch WiFi Smart Light: Controlled via API or Go server integration.
- Photo Diode Sensors: Detects motion and ambient light levels for automation.
- This project provides a real-time, fully logged, and voice-controlled IoT automation system without relying on paid external services like IFTTT.

### Requirements                      
ESP8266 Module
Arduino IDE and `ESP8266WebServer` and `WebSocketsServer` library  (arduinoWebSockets)
Flash ESP8266 with WebSocket server firmware
Connect to WiFi and establish WebSocket connection to Go server
Send sensor data and control relay/motor

### Example Use Cases
- 1️⃣ Turn ON/OFF Smart Light using ESP8266
- 2️⃣ Control Relay Motor via WebSocket
- 3️⃣ Detect Motion and Adjust Smart Light Brightness
- 4️⃣ Log All Events to MongoDB
- 5️⃣ Use Google Nest Mini for Voice Commands


# This  Project is still Under Development Phase
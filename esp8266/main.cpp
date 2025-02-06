#include <ESP8266WiFi.h>
#include <ESP8266WebServer.h>

// Replace with your Wi-Fi credentials
const char* ssid = "your_wifi_ssid";
const char* password = "your_wifi_password";

// Relay Pin connected to the smart plug
const int relayPin = D1; // Pin where relay is going to be connected
const int photodiodePin = A0; // Pin where photodiode  is going to be connected
const int motorPin = A1; // Pin where photodiode  is going to be connected
int photodiodeValue = 0;
ESP8266WebServer server(80); // Create a web server on port 80

void setup() {
  // Start the Serial Monitor for debugging
  Serial.begin(115200);
  
  // Set relay pin as output
  pinMode(relayPin, OUTPUT);
  digitalWrite(relayPin, LOW); // Initially turn off the relay

  // Connect to Wi-Fi
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.println("Connecting to WiFi...");
  }
  Serial.println("Connected to WiFi");

  // Print ESP8266 IP address
  Serial.print("ESP IP Address: ");
  Serial.println(WiFi.localIP());

  // Define the routes
  server.on("/on", HTTP_GET, turnOn);
  server.on("/off", HTTP_GET, turnOff);

  // Start the web server
  server.begin();
  Serial.println("Server started");
}

void loop() {
  // Read the analog value from the photodiode
  photodiodeValue = analogRead(photodiodePin);

  // Print the photodiode value to Serial Monitor
  Serial.println(photodiodeValue);

  // If the value is above a threshold (indicating infrared light detected)
  if (photodiodeValue > 500) {
    Serial.println("IR light detected!");
    // Perform your task here (e.g., control smart plug, turn on a light, etc.)
  }

  delay(500); // Small delay for stability
  server.handleClient(); // Handle incoming client requests
}

void turnOn() {
  digitalWrite(relayPin, HIGH); // Turn on the relay
  server.send(200, "text/plain", "Smart Plug is ON");
}

void turnOff() {
  digitalWrite(relayPin, LOW); // Turn off the relay
  server.send(200, "text/plain", "Smart Plug is OFF");
}

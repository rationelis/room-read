#include <DHT.h>
#include <WiFi.h>
#include <PubSubClient.h>

const char *SSID = nullptr;
const char *PWD = nullptr;

const char *SERVER = "192.168.178.214";
const int PORT = 1883;

#define DHT_SENSOR_PIN 4
#define DHT_SENSOR_TYPE DHT22

DHT dht_sensor(DHT_SENSOR_PIN, DHT_SENSOR_TYPE);

WiFiClient wifiClient;
PubSubClient mqttClient(wifiClient);

void connectToWiFi() {
  Serial.print("Connecting to ");

  WiFi.begin(SSID, PWD);
  Serial.print(SSID);

  while (WiFi.status() != WL_CONNECTED) {
    Serial.print(".");
    delay(500);
  }

  Serial.println("Connected.");
}

void setupMQTT() {
  mqttClient.setServer(SERVER, PORT);
}

void reconnect() {
  Serial.println("Connecting to MQTT Broker");
  while (!mqttClient.connected()) {
    Serial.println("Reconnecting to MQTT Broker..");
    String clientId = "ESP32Client-";
    clientId += String(random(0xffff), HEX);

    if (mqttClient.connect(clientId.c_str())) {
      Serial.println("Connected.");
    }

    delay(500);
  }
}

void setup() {
  Serial.begin(115200);

  SSID = getenv("WIFI_SSID");
  PWD = getenv("WIFI_PASSWORD");

  if (SSID == nullptr || PWD == nullptr) {
    Serial.println("Error: SSID and password not set in environment variables.");
    while (1) {
      delay(1000);
    }
  }

  connectToWiFi();

  setupMQTT();

  dht_sensor.begin();
}

void loop() {
  if (!mqttClient.connected())
    reconnect();
  mqttClient.loop();

  float humi = dht_sensor.readHumidity();
  float tempC = dht_sensor.readTemperature();
  float abs = dht_sensor.computeHeatIndex(tempC, humi, false);

  if (isnan(abs)) {
    Serial.println("Failed to read from DHT sensor!");
  } else {
    mqttClient.publish("home/temperature", String(abs).c_str());
  }

  delay(5000);
}

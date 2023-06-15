package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"os"
	"os/signal"
	"server/config"
	mqtt2 "server/pkg/mqtt"
	"server/pkg/scylla"
	"strconv"
	"strings"
	"time"
)

func main() {
	c := config.NewConfig()

	mqttClient, err := mqtt2.NewMqttClient(mqtt2.Config{
		Host:           c.MqttBrokerConfig.Host,
		Port:           c.MqttBrokerConfig.Port,
		ClientId:       c.MqttBrokerConfig.ClientId,
		Username:       c.MqttBrokerConfig.Username,
		Password:       c.MqttBrokerConfig.Password,
		KeepAliveSec:   c.MqttBrokerConfig.KeepAliveSec,
		PingTimeoutSec: c.MqttBrokerConfig.PingTimeoutSec,
	})
	if err != nil {
		panic(err)
	}

	session, err := scylla.NewScyllaSession(scylla.Config{
		Host:            c.ScyllaConfig.Host,
		Port:            c.ScyllaConfig.Port,
		Username:        c.ScyllaConfig.Username,
		Password:        c.ScyllaConfig.Password,
		DefaultKeyspace: c.ScyllaConfig.DefaultKeyspace,
	})
	if err != nil {
		panic(err)
	}
	var tn = c.TopicNames
	mqttClient.Subscribe(fmt.Sprintf("%s/+", tn.Request), mqtt2.ExactlyOnce, func(client mqtt.Client, message mqtt.Message) {
		topic := message.Topic()
		topicParts := strings.Split(topic, "/")
		deviceId := topicParts[1]

		payload := string(message.Payload())
		payloadParts := strings.Split(payload, ",")
		secondAfterStart, err := strconv.Atoi(payloadParts[0])
		if err != nil {
			// log the error using zap
			return
		}
		cardUid := payloadParts[1]

		fmt.Println("request device_id:", deviceId, " secondAfterStart:", secondAfterStart, " cardUid:", cardUid)

		// check authorization and then send response
		//lockPermitted: LOCK_OPEN_PERMITED, LOCK_OPEN_NOT_PERMITED
		var isPermitted = true
		lockPermitted := "LOCK_OPEN_NOT_PERMITTED"
		if isPermitted {
			lockPermitted = "LOCK_OPEN_PERMITTED"
		}
		lcdMessage := "your welcome mahmoud ;)"

		messagePayload := fmt.Sprintf("%d,%s,%s", time.Now().Unix(), lockPermitted, lcdMessage)
		mqttClient.Publish(fmt.Sprintf("response/%s/%s", deviceId, cardUid), mqtt2.ExactlyOnce, false, messagePayload)
	})
	mqttClient.Subscribe(fmt.Sprintf("%s/+", tn.LockOpened), mqtt2.ExactlyOnce, func(client mqtt.Client, message mqtt.Message) {
		topic := message.Topic()
		topicParts := strings.Split(topic, "/")
		deviceId := topicParts[1]

		payload := string(message.Payload())
		payloadParts := strings.Split(payload, ",")
		secondAfterStart, err := strconv.Atoi(payloadParts[0])
		if err != nil {
			// log the error using zap
			return
		}

		fmt.Println("lock-opened deviceId:", deviceId, " secondAfterStart:", secondAfterStart)

	})
	mqttClient.Subscribe(fmt.Sprintf("%s/+", tn.DeviceState), mqtt2.ExactlyOnce, func(client mqtt.Client, message mqtt.Message) {
		topic := message.Topic()
		topicParts := strings.Split(topic, "/")
		deviceId := topicParts[1]

		payload := string(message.Payload())
		payloadParts := strings.Split(payload, ",")
		secondAfterStart, err := strconv.Atoi(payloadParts[0])
		if err != nil {
			// log the error using zap
			return
		}
		rgbString := strings.TrimPrefix(payloadParts[1], "LED_COLOR_")
		red, err := strconv.Atoi(fmt.Sprintf("%c", rgbString[0]))
		if err != nil {
			// log the error using zap
			return
		}
		green, err := strconv.Atoi(fmt.Sprintf("%c", rgbString[1]))
		if err != nil {
			// log the error using zap
			return
		}
		blue, err := strconv.Atoi(fmt.Sprintf("%c", rgbString[2]))
		if err != nil {
			// log the error using zap
			return
		}

		fmt.Println("device-state deviceId:", deviceId, " secondsAfterStart:", secondAfterStart, " red:", red, " green:", green, " blue:", blue)

	})

	//mqttClient.Publish(topicNames.AdminCommand, mqtt2.ExactlyOnce, false, "")
	//mqttClient.Publish(fmt.Sprintf("%s-%s", topicNames.Response, "USER_ID"), mqtt2.ExactlyOnce, false, "")

	fmt.Println(mqttClient)
	fmt.Println(session)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("gracefully shutdown ...")
	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//defer cancel()
	mqttClient.Disconnect(250)
	session.Close()
	time.Sleep(3 * time.Second)
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal(err)
	//}

}

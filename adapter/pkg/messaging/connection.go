/*
 *  Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 *  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

// Package messaging holds the implementation for event listeners functions
package messaging

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
	logger "github.com/wso2/product-microgateway/adapter/pkg/loggers"
)

var (
	// NotificationChannel stores the Events for notifications
	NotificationChannel chan amqp.Delivery
	// KeyManagerChannel stores the key manager eventsv
	KeyManagerChannel chan amqp.Delivery
	// RevokedTokenChannel stores the revoked token events
	RevokedTokenChannel chan amqp.Delivery
	// ThrottleDataChannel stores the throttling related events
	ThrottleDataChannel chan amqp.Delivery
)

func init() {
	NotificationChannel = make(chan amqp.Delivery)
	KeyManagerChannel = make(chan amqp.Delivery)
	RevokedTokenChannel = make(chan amqp.Delivery)
	ThrottleDataChannel = make(chan amqp.Delivery)
}

// EventListeningEndpoints represents the list of endpoints
var EventListeningEndpoints []string

// ConnectToRabbitMQ function tries to connect to the RabbitMQ server as long as it takes to establish a connection
func connectToRabbitMQ() (*amqp.Connection, error) {
	var err error = nil
	var conn *amqp.Connection
	amqpURIArray = retrieveAMQPURLList()
	for index, url := range amqpURIArray {
		logger.LoggerMsg.Infof("Dialing URI [%d] %q", index, maskURL(url.URL)+"/")
		conn, err = amqp.Dial(url.URL)
		if err == nil {
			logger.LoggerMsg.Infof("Successfully established the AMQP connection on URI [%d], %q", index, maskURL(url.URL)+"/")
			return conn, nil
		}
	}
	_, conn, err = connectionRetry("")
	return conn, err
}

// reconnect reconnects to server if the connection or a channel
// is closed unexpectedly.
func (c *Consumer) reconnect(key string) {
	var err error
	shouldReconnect := false
	connClose := <-c.Conn.NotifyClose(make(chan *amqp.Error))
	connBlocked := c.Conn.NotifyBlocked(make(chan amqp.Blocking))
	chClose := c.Channel.NotifyClose(make(chan *amqp.Error))

	if connClose != nil {
		shouldReconnect = true
		logger.LoggerMsg.Errorf("CRITICAL: Connection dropped for %s, reconnecting...", key)
	}

	if connBlocked != nil {
		shouldReconnect = true
		logger.LoggerMsg.Errorf("CRITICAL: Connection blocked for %s, reconnecting...", key)
	}

	if chClose != nil {
		shouldReconnect = true
		logger.LoggerMsg.Errorf("CRITICAL: Channel closed for %s, reconnecting...", key)
	}

	if shouldReconnect {
		c.Conn.Close()
		c, RabbitConn, err = connectionRetry(key)
		if err != nil {
			logger.LoggerMsg.Errorf("Cannot establish connection for topic %s", key)
		}
	} else {
		logger.LoggerMsg.Infof("NotifyClose from the connection and channel are %v and %v respectively, NotifyBlocked from the connection is %v",
			connClose, chClose, connBlocked)
	}
}

// connectionRetry
func connectionRetry(key string) (*Consumer, *amqp.Connection, error) {
	var err error = nil
	var i int

	for j := 0; j < len(amqpURIArray); j++ {
		var maxAttempt int = amqpURIArray[j].retryCount
		var retryInterval time.Duration = time.Duration(amqpURIArray[j].connectionDelay) * time.Second

		if maxAttempt == 0 {
			maxAttempt = 5
		}

		if retryInterval == 0 {
			retryInterval = 10 * time.Second
		}
		logger.LoggerMsg.Infof("Retrying to connect with %s (URI %d) in every %d seconds until exceed %d attempts",
			maskURL(amqpURIArray[j].URL), j, amqpURIArray[j].connectionDelay, maxAttempt)

		for i = 1; i <= maxAttempt; i++ {

			RabbitConn, err = amqp.Dial(amqpURIArray[j].URL + "/")
			if err == nil {
				logger.LoggerMsg.Infof("Successfully connected to %s (URI %d) after %d attempts", maskURL(amqpURIArray[j].URL), j, i)
				if key != "" && len(key) > 0 {
					logger.LoggerMsg.Infof("Reconnected to topic %s", key)
					// startup pull
					c := startConsumer(key)
					return c, RabbitConn, nil
				}
				return nil, RabbitConn, nil
			}

			if key != "" && len(key) > 0 {
				logger.LoggerMsg.Infof("Retry attempt %d for the %s (URI %d) to connect with topic %s has failed. Retrying after %d seconds", i,
					maskURL(amqpURIArray[j].URL), j, key, amqpURIArray[j].connectionDelay)
			} else {
				logger.LoggerMsg.Infof("Retry attempt %d for the %s (URI %d) has failed. Retrying after %d seconds", i, maskURL(amqpURIArray[j].URL), j, amqpURIArray[j].connectionDelay)
			}
			time.Sleep(retryInterval)
		}
		// Retried all the uris in the amqp URI array. Start from begining.
		if j == len(amqpURIArray)-1 {
			j = -1
		}
	}
	return nil, RabbitConn, err
}

// retrieveAMQPURLList function extract AMQPURLList from EventListening connection url
func retrieveAMQPURLList() []amqpFailoverURL {
	var connectionURLList []string
	connectionURLList = EventListeningEndpoints

	amqlURLList := []amqpFailoverURL{}

	for _, conURL := range connectionURLList {
		var delay int = 0
		var retries int = 0
		amqpConnectionURL := strings.Split(conURL, "?")[0]
		u, err := url.Parse(conURL)
		if err != nil {
			logger.LoggerMsg.Errorf("Error occured %s", maskURL(err.Error()))
		} else {
			m, _ := url.ParseQuery(u.RawQuery)
			if m["connectdelay"] != nil {
				connectdelay := m["connectdelay"][0]
				delay, _ = strconv.Atoi(connectdelay[1 : len(connectdelay)-1])
			}

			if m["retries"] != nil {
				retrycount := m["retries"][0]
				retries, _ = strconv.Atoi(retrycount[1 : len(retrycount)-1])
			}

			failoverurlObj := amqpFailoverURL{URL: amqpConnectionURL, retryCount: retries,
				connectionDelay: delay}
			amqlURLList = append(amqlURLList, failoverurlObj)
		}
	}
	return amqlURLList
}

// maskURL function mask the incoming url
func maskURL(url string) string {
	pattern := regexp.MustCompile(`\/\/([a-zA-Z].*)@\b`)
	matches := pattern.FindStringSubmatch(url)
	if len(matches) > 1 {
		return strings.ReplaceAll(url, matches[1], "******")
	}
	return url
}

// amqpFailoverURL defines the structure of an amqp failover url
type amqpFailoverURL struct {
	URL             string
	retryCount      int
	connectionDelay int
}

func handleEvent(c *Consumer, key string) error {
	var err error

	logger.LoggerMsg.Debugf("Getting Channel for %s events.", key)
	c.Channel, err = c.Conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	logger.LoggerMsg.Debugf("got Channel, declaring Exchange (%q)", exchange)
	if err = c.Channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	logger.LoggerMsg.Infof("declared Exchange, declaring Queue %q", key+"queue")
	queue, err := c.Channel.QueueDeclare(
		"",    // name of the queue
		false, // durable
		true,  // delete when usused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("Error while declaring queue: %s", err)
	}

	logger.LoggerMsg.Debugf("Binding to Exchange (key %q) after declaring the Queue (%q %d messages, %d consumers)",
		key, queue.Name, queue.Messages, queue.Consumers)

	if err = c.Channel.QueueBind(
		queue.Name, // name of the queue
		key,        // bindingKey
		exchange,   // sourceExchange
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return fmt.Errorf("Queue Bind: %s", err)
	}
	logger.LoggerMsg.Infof("Queue bound to Exchange, starting Consume (consumer tag %q) events", c.Tag)
	deliveries, err := c.Channel.Consume(
		queue.Name, // name
		c.Tag,      // consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	if strings.EqualFold(key, Notification) {
		for event := range deliveries {
			NotificationChannel <- event
		}
	} else if strings.EqualFold(key, keymanager) {
		for event := range deliveries {
			KeyManagerChannel <- event
		}
	} else if strings.EqualFold(key, TokenRevocation) {
		for event := range deliveries {
			RevokedTokenChannel <- event
		}
	} else if strings.EqualFold(key, throttleData) {
		for event := range deliveries {
			ThrottleDataChannel <- event
		}
	}
	return nil
}

// InitiateJMSConnection to pass event consumption
func InitiateJMSConnection(eventListeningEndpoints []string) error {
	var err error
	EventListeningEndpoints = eventListeningEndpoints
	bindingKeys := []string{Notification, keymanager, TokenRevocation, throttleData}
	RabbitConn, err = connectToRabbitMQ()

	if err == nil {
		for i, key := range bindingKeys {
			logger.LoggerMsg.Infof("Establishing consumer index %v for key %s ", i, key)
			go func(key string) {
				startConsumer(key)
				select {}
			}(key)
		}
	}
	return err
}

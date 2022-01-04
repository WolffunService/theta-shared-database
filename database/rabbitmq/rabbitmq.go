package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func InitRabbitMq(config RabbitMQConfig) (*amqp.Connection, *amqp.Channel) {
	host := config.Host
	port := config.Port
	username := config.UserName
	password := config.Password
	vHost := config.Vhost
	rabbitmqConnString := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", username, password, host, port, vHost)
	conn, ch := connect(rabbitmqConnString, config.Exchange)
	return conn, ch
}

func connect(url string, exchange string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "[ERROR] FAILED TO CONNECT TO RABBITMQ")

	ch, err := conn.Channel()
	failOnError(err, "[ERROR] FAILED TO OPEN A CHANNEL RABBITMQ")

	_ = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
	return conn, ch
}

func DeclareQueue(ch *amqp.Channel, name string, exchange string, durable bool) amqp.Queue {
	q, err := ch.QueueDeclare(
		name, // name
		durable,                         // durable
		false,                         // delete when unused
		false,                         // exclusive
		false,                         // no-wait
		nil,                           // arguments
	)
	failOnError(err, "[ERROR] FAILED TO DECLARE A QUEUE")
	err = ch.QueueBind(
		q.Name,                                 // queue name
		fmt.Sprintf("%s.%s", exchange, q.Name), // routing key
		exchange,                               // exchange
		false,
		nil,
	)
	failOnError(err, "[ERROR] FAILED TO BINDING A QUEUE")
	log.Printf("[INFO] DECLARED QUEUE " + name)
	return q
}

func Publish(ch *amqp.Channel, q amqp.Queue, data []byte) error  {
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", q.Name)
	return err
}

func Consume(ch *amqp.Channel, q amqp.Queue, opt Option) <-chan amqp.Delivery {
	log.Printf("[INFO] CONSUMING " + q.Name)
	messages, err := ch.Consume(
		q.Name, // queue
		q.Name, // consumer
		opt.AutoACK,   // auto-ack
		opt.Exclusive,  // exclusive
		opt.NoLocal,  // no-local
		opt.NoWait,  // no-wait
		opt.Args,    // args
	)
	failOnError(err, "[ERROR] FAILED TO REGISTER A CONSUMER")
	return messages
}

func HandleMessages(qName string,isProduction bool, messages <-chan amqp.Delivery, f func(d []byte) error) {
	forever := make(chan bool)

	go func() {
		for d := range messages {
			fmt.Printf("[INFO] RECEIVED A MESSAGE OF %s: %s", qName, d.Body)
			err := f(d.Body)

			if err != nil {
				fmt.Printf("[ERR] ERROR WHEN CONSUME A MESSAGE OF %s: %s", qName, d.Body)
				errReject := d.Reject(isProduction) //không phải production thì không cần requeue

				if errReject != nil {
					fmt.Printf("[ERR] ERROR REQUEUE MESSAGE OF %s: %s", qName, d.Body)
				}
			} else {
				//Không phải production thì sẽ autoAck
				if isProduction {
					errAck := d.Ack(false)

					if errAck != nil {
						fmt.Printf("[ERR] ERROR SENT POSITIVE ACK MESSAGE OF %s: %s", qName, d.Body)
					}
				}

			}
		}
	}()

	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Option struct{
	AutoACK bool
	Exclusive bool
	NoLocal bool
	NoWait bool
	Args amqp.Table
}

var DefaultOption = Option{
	AutoACK:   false,
	Exclusive: false,
	NoLocal:   false,
	NoWait:    false,
	Args:      nil,
}
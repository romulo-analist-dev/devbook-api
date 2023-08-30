package main

import (
	"api/src/config"
	controller "api/src/controllers"
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	emitter "github.com/emitter-io/go"
	"github.com/rs/cors"
)

const channelKey = "alguma-coisa"

func main() {
	config.Carregar()
	r := router.Gerar()

	opts := emitter.NewClientOptions()
	opts.AddBroker("tcp://broker.hivemq.com:1883")
	opts.SetOnMessageHandler(onMessage)

	client := emitter.NewClient(opts)
	wait(client.Connect())
	wait(client.Subscribe(channelKey, "demo/"))
	//wait(client.Publish(channelKey, "demo/", "hello, carai!"))

	fmt.Printf("Escutando na porta %d", config.Porta)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Permitir todas as origens
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := corsHandler.Handler(r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), handler))

	for {
	}
}

func onMessage(client emitter.Emitter, msg emitter.Message) {
	fmt.Printf("topic: %v, message: %s\n", msg.Topic(), msg.Payload())
	controller.CreateRecord(string(msg.Payload()))
}

func wait(t emitter.Token) {
	t.Wait()
	if t.Error() != nil {
		panic(t.Error())
	}
}

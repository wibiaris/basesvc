package mokabox

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/iwanjunaid/basesvc/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	mConfig "github.com/iwanjunaid/mokabox/config"
	events "github.com/iwanjunaid/mokabox/event"
	event "github.com/iwanjunaid/mokabox/internal/interfaces/event"
	"github.com/iwanjunaid/mokabox/manager"

	"go.mongodb.org/mongo-driver/mongo"
)

type MokaBoxImpl struct {
	kc *kafka.Consumer
	mc *mongo.Client
}

type Event interface {
	fmt.Stringer
	GetPickerGroupID() string
	GetTimestamp() time.Time
}

func NewMokaBox(kc *kafka.Consumer, mdb *mongo.Client) *MokaBoxImpl {
	//registry := registry.NewRegistry(db, registry.NewMongoConn(
	//	mdb.Collection(config.GetString("database.mongo.collection"))))
	//appController := registry.NewAppController()

	return &MokaBoxImpl{
		kc: kc,
		mc: mdb,
	}
}

func (c *MokaBoxImpl) Listen(topic []string) {
	eventHandler := func(e event.Event) {
		switch event := e.(type) {
		case events.PickerStarted:
			fmt.Printf("%v\n", event)
		case events.Picked:
			fmt.Printf("%v\n", event)
		case events.Sent:
			fmt.Printf("%v\n", event)
		case events.StatusChanged:
			fmt.Printf("%v\n", event)
		case events.PickerPaused:
			fmt.Printf("%v\n", event)
		case events.ZombiePickerStarted:
			fmt.Printf("%v\n", event)
		case events.ZombiePicked:
			fmt.Printf("%v\n", event)
		case events.ZombieAcquired:
			fmt.Printf("%v\n", event)
		case events.ZombiePickerPaused:
			fmt.Printf("%v\n", event)
		case events.RemoverStarted:
			fmt.Printf("%v\n", event)
		case events.Removed:
			fmt.Printf("%v\n", event)
		case events.RemoverPaused:
			fmt.Printf("%v\n", event)
		case events.ErrorOccured:
			fmt.Printf("%v\n", event)
		}
	}

	outboxConfig := mConfig.NewDefaultCommonOutboxConfig(config.GetString("kafka.group_id"), config.GetString("mongo.db"))

	kafkaConfig := mConfig.NewCommonKafkaConfig(&kafka.ConfigMap{
		"bootstrap.servers": config.GetString("kafka.host"),
		"acks":              "all",
	})

	manager, err := manager.New(outboxConfig, kafkaConfig, c.mc)

	if err != nil {
		log.Fatal(err)
	}

	manager.SetEventHandler(eventHandler)
	manager.Start()
	manager.Await()
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{})
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 page not found"))
}
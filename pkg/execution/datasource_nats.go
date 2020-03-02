package execution

import (
	"encoding/json"
	log "github.com/jensneuse/abstractlogger"
	"github.com/jensneuse/graphql-go-tools/pkg/lexer/literal"
	"github.com/nats-io/nats.go"
	"io"
	"sync"
	"time"
)

type NatsDataSourceConfig struct {
	Addr  string
	Topic string
}

type NatsDataSourcePlannerFactoryFactory struct {
}

func (n NatsDataSourcePlannerFactoryFactory) Initialize(base BaseDataSourcePlanner, configReader io.Reader) (DataSourcePlannerFactory, error) {
	factory := &NatsDataSourcePlannerFactory{
		base: base,
	}
	return factory, json.NewDecoder(configReader).Decode(&factory.config)
}

type NatsDataSourcePlannerFactory struct {
	base   BaseDataSourcePlanner
	config NatsDataSourceConfig
}

func (n NatsDataSourcePlannerFactory) DataSourcePlanner() DataSourcePlanner {
	return SimpleDataSourcePlanner(&NatsDataSourcePlanner{
		BaseDataSourcePlanner: n.base,
		dataSourceConfig:      n.config,
	})
}

type NatsDataSourcePlanner struct {
	BaseDataSourcePlanner
	dataSourceConfig NatsDataSourceConfig
}

func (n *NatsDataSourcePlanner) Plan([]Argument) (DataSource, []Argument) {
	n.args = append(n.args, &StaticVariableArgument{
		Name:  literal.ADDR,
		Value: []byte(n.dataSourceConfig.Addr),
	})
	n.args = append(n.args, &StaticVariableArgument{
		Name:  literal.TOPIC,
		Value: []byte(n.dataSourceConfig.Topic),
	})
	return &NatsDataSource{
		log: n.log,
	}, n.args
}

type NatsDataSource struct {
	log  log.Logger
	conn *nats.Conn
	sub  *nats.Subscription
	once sync.Once
}

func (d *NatsDataSource) Resolve(ctx Context, args ResolvedArgs, out io.Writer) (n int, err error) {
	d.once.Do(func() {

		addrArg := args.ByKey(literal.ADDR)
		topicArg := args.ByKey(literal.TOPIC)

		addr := nats.DefaultURL
		topic := string(topicArg)

		if len(addrArg) != 0 {
			addr = string(addrArg)
		}

		go func() {
			<-ctx.Done()
			if d.sub != nil {
				d.log.Debug("NatsDataSource.unsubscribing",
					log.String("addr", addr),
					log.String("topic", topic),
				)
				err := d.sub.Unsubscribe()
				if err != nil {
					d.log.Error("Unsubscribe", log.Error(err))
				}
			}
			if d.conn != nil {
				d.log.Debug("NatsDataSource.closing",
					log.String("addr", addr),
					log.String("topic", topic),
				)
				d.conn.Close()
			}
		}()

		d.log.Debug("NatsDataSource.connecting",
			log.String("addr", addr),
			log.String("topic", topic),
		)

		d.conn, err = nats.Connect(addr)
		if err != nil {
			panic(err)
		}

		d.log.Debug("NatsDataSource.subscribing",
			log.String("addr", addr),
			log.String("topic", topic),
		)

		d.sub, err = d.conn.SubscribeSync(topic)
		if err != nil {
			panic(err)
		}
	})

	if err != nil {
		return n, err
	}

	message, err := d.sub.NextMsg(time.Minute)
	if err != nil {
		return n, err
	}

	return out.Write(message.Data)
}

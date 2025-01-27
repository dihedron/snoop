package record

/*
import (
	"context"
	"errors"

	"github.com/bancaditalia/brokerd/dataflow"
	"github.com/bancaditalia/brokerd/message"
	"github.com/bancaditalia/brokerd/model"
	"github.com/dgraph-io/ristretto"
	"go.uber.org/zap"
)

type Correlator struct {
	cache *ristretto.Cache
}

func NewCorrelator() (*Correlator, error) {
	var err error
	result := &Correlator{}
	result.cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e4,     // numbser of keys to track frequency of (10k).
		MaxCost:     1 << 25, // maximum cost of cache (32MB).
		BufferItems: 64,      // number of keys per Get buffer.
		OnEvict: func(item *ristretto.Item) {
			zap.S().Debugf("evicting %s: %+v", item.Key, item.Value)
		},
	})
	if err != nil {
		zap.S().With(zap.Error(err)).Error("error instantiating cache")
		return nil, err
	}
	zap.S().Debug("cache ready")
	return result, nil
}

func (c *Correlator) Close() {
	if c.cache != nil {
		zap.S().Debug("closing cache")
		defer c.cache.Close()
	}
}

func (c *Correlator) Name() string {
	return "github.com/bancaditalia/brokerd/command/record/Correlator"
}

var ErrInvalidMessageType = errors.New("invalid input message type")

// Process stores the result of the previous request id into the Context,
// for further processing by following filters.
func (c *Correlator) Process(ctx context.Context, quit chan<- bool, msg dataflow.Message) (dataflow.Message, context.Context, error) {
	zap.S().Debugf("processing message of type %T", msg)
	select {
	case <-ctx.Done():
		zap.S().Debug("context cancelled")
		return msg, ctx, dataflow.ErrInterrupted
	default:
		if msg, ok := msg.(*message.OpenStackMessage); !ok {
			zap.S().With(zap.Error(ErrInvalidMessageType)).Errorf("input message type is %T", msg)
			return msg, ctx, ErrInvalidMessageType
		} else {
			zap.S().Debugf("OpenStack notification of type %T with event type %q", msg.Message, msg.Message.Info().EventType)
			switch object := msg.Message.(type) {
			case *message.ComputeInstanceNotification:
				// try to retrieve the item from the cache
				var (
					cached interface{}
					ok     bool
				)
				if cached, ok = c.cache.Get(object.ContextRequestID); ok {
					zap.S().Debugf("item for request ID %s found in cache (type %T)", object.ContextRequestID, cached)
				} else {
					zap.S().Debugf("item for request ID %s not found in cache", object.ContextRequestID)
				}
				switch object.EventType {
				case
					"scheduler.select_destinations.start",
					"scheduler.select_destinations.end":
					if not, ok := msg.Message.(*message.ComputeTaskNotification); ok {
						var vm *model.VirtualMachine
						if cached != nil {
							if vm, ok = cached.(*model.VirtualMachine); !ok {
								vm = &model.VirtualMachine{}
							}
						}
						vm.InstanceID = model.String(not.Payload.InstanceID)

						vm.ImageID = model.String(not.Payload.RequestSpec.Image.ID)
						vm.ImageName = model.String(not.Payload.RequestSpec.Image.Name)
						vm.ImageSize = model.Int64(not.Payload.RequestSpec.Image.Size)
						vm.ImageChecksum = model.String(not.Payload.RequestSpec.Image.Checksum)
						vm.InstanceID = model.String(not.Payload.InstanceProperties.UUID) // TODO: check!
						// TODO: go on...

						// payload.image.id (UUID), e.g. "55dec1d5-1e6b-4d97-818b-7f0c6df1ff7d".
						// 5. payload.image.image_name (string), e.g. "rhel-8.3".
						// 6. payload.image.image_checksum (shasum), e.g."dd554c059e0910379fff88f677f4a4b3".
						// 7. payload.image.image_size (int) e.g. 1316683776.
						id := not.Payload.RequestSpec.Image.ID
						name := not.Payload.RequestSpec.Image.Name
						checksum := not.Payload.RequestSpec.Image.Checksum
						size := not.Payload.RequestSpec.Image.Size
						zap.S().Debugf("image: ID: %v, name: %v, checksum: %v, size: %d", id, name, checksum, size)
						// TODO: go on trating message
						// store into cache, acknowledge message
					}
				default:
					// in all other compute messages, we should already have
					// an entry in the cache; if we don't, we must query OpenStack
					// to populate it first.
				}
			}
		}
		return msg, ctx, nil
	}
}
*/

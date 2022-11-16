package events

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/openshift/assisted-service/internal/common"
	eventsapi "github.com/openshift/assisted-service/internal/events/api"
	"github.com/openshift/assisted-service/models"
	logutil "github.com/openshift/assisted-service/pkg/log"
	"github.com/openshift/assisted-service/restapi"
	"github.com/openshift/assisted-service/restapi/operations/events"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var _ restapi.EventsAPI = &Api{}

type Api struct {
	handler eventsapi.Handler
	log     logrus.FieldLogger
}

func NewApi(handler eventsapi.Handler, log logrus.FieldLogger) *Api {
	return &Api{
		handler: handler,
		log:     log,
	}
}

func (a *Api) V2ListEvents(ctx context.Context, params events.V2ListEventsParams) middleware.Responder {
	log := logutil.FromContext(ctx, a.log)

	evs, err := a.handler.V2GetEvents(ctx, params.ClusterID, params.HostID, params.InfraEnvID, params.Categories...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.NewApiError(http.StatusNotFound, err)
		}
		log.WithError(err).Errorf("failed to get events")
		return common.NewApiError(http.StatusInternalServerError, err)
	}
	ret := make(models.EventList, len(evs))
	for i, ev := range evs {
		ret[i] = &models.Event{
			Name:       ev.Name,
			ClusterID:  ev.ClusterID,
			HostID:     ev.HostID,
			InfraEnvID: ev.InfraEnvID,
			Severity:   ev.Severity,
			EventTime:  ev.EventTime,
			Message:    ev.Message,
			Props:      ev.Props,
		}
	}
	return events.NewV2ListEventsOK().WithPayload(ret)
}

// V2EventsSubscribe will register a call back url on a given even
// This url will be called in case the event occur
func (a *Api) V2EventsSubscribe(ctx context.Context, params events.V2EventsSubscribeParams) middleware.Responder {
	log := logutil.FromContext(ctx, a.log)
	log.Infof("EventsSubscribe: %+v", params)

	eventSubscription, err := a.handler.CreateEventsSubscribe(ctx,
		*params.NewEventSubscriptionParams.ClusterID,
		*params.NewEventSubscriptionParams.EventName,
		*params.NewEventSubscriptionParams.URL,
	)
	if err != nil {

	}
	return events.NewV2EventsSubscribeCreated().WithPayload(&eventSubscription)
}

func (a *Api) V2EventsSubscriptionGet(ctx context.Context, params events.V2EventsSubscriptionGetParams) middleware.Responder {
	log := logutil.FromContext(ctx, a.log)
	log.Infof("EventsSubscribe: %+v", params)

	eventSubscription, err := a.handler.GetEventsSubscription(ctx,
		params.SubscriptionID,
	)
	if err != nil {

	}
	return events.NewV2EventsSubscribeCreated().WithPayload(&eventSubscription)
}

func (a *Api) V2EventsSubscriptionDelete(ctx context.Context, params events.V2EventsSubscriptionDeleteParams) middleware.Responder {
	log := logutil.FromContext(ctx, a.log)
	log.Debugf("V2EventsSubscriptionDelete doing noting for no", params)
	return events.NewV2EventsSubscriptionDeleteInternalServerError()
}

func (a *Api) V2EventsSubscriptionList(ctx context.Context, params events.V2EventsSubscriptionListParams) middleware.Responder {
	log := logutil.FromContext(ctx, a.log)
	log.Debugf("V2EventsSubscriptionGet doing noting for no", params)
	return events.NewV2EventsSubscriptionListInternalServerError()
}

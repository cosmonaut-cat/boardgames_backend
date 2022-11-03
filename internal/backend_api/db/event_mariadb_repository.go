package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	eventTypes "github.com/cosmonaut-cat/boardgames_backend/internal/backend_api/domain/event"
	"github.com/cosmonaut-cat/boardgames_backend/pkg/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type mariadbEvent struct {
	ID          string `db:"event_id"`
	Type        string `db:"event_type"`
	Entity      string `db:"event_entity"`
	Version     int64  `db:"event_version"`
	Timestamp   string `db:"event_timestamp"`
	PayloadType string `db:"event_payload_type"`
	Payload     []byte `db:"event_payload"`
}

type MariaDBEventRepository struct {
	db *sqlx.DB
}

func NewMariaDBEventRepository(db *sqlx.DB) *MariaDBEventRepository {
	if db == nil {
		log.Fatalf("Missing database in event repository\n")
	}

	return &MariaDBEventRepository{db: db}
}

func (m MariaDBEventRepository) AppendEvents(ctx context.Context, eventId string, events []*api.Event) error {
	if len(events) == 0 {
		return errors.New(fmt.Sprintf("without items to append"))
	}

	eventCurrentVer := events[0].EventVersion
	storedEventLatestVer, err := m.getEventLatestVersion(ctx, eventId)

	if err != nil {
		return err
	}

	if storedEventLatestVer == 0 {
		lastEvent := events[len(events)-1]
		events = append([]*api.Event{
			{
				EventId:        eventId,
				EventType:      string(eventTypes.LatestVersion),
				EventEntity:    lastEvent.EventEntity,
				EventVersion:   lastEvent.EventVersion,
				EventTimestamp: lastEvent.EventTimestamp,
				EventPayload:   lastEvent.EventPayload,
			},
		}, events...)
	}

	if eventCurrentVer != storedEventLatestVer+1 {
		return errors.New(fmt.Sprintf("version expected %d received %d\n", storedEventLatestVer+1, eventCurrentVer))
	}

	tx, err := m.db.Beginx()

	if err != nil {
		return err
	}

	defer func() {
		err = m.finishTransaction(err, tx)
	}()

	err = m.appendEvents(tx, eventId, storedEventLatestVer, events)

	if err != nil {
		return err
	}

	return nil
}

func (m MariaDBEventRepository) appendEvents(tx *sqlx.Tx, eventId string, eventLatestVersion int64, events []*api.Event) error {
	newEvents := []mariadbEvent{}

	for _, event := range events {
		newEvents = append(newEvents, mariadbEvent{
			ID:          eventId,
			Type:        event.EventType,
			Entity:      event.EventEntity,
			Version:     event.EventVersion,
			Timestamp:   event.EventTimestamp.AsTime().String(),
			PayloadType: event.EventPayload.GetTypeUrl(),
			Payload:     event.EventPayload.GetValue(),
		})
	}

	_, err := tx.NamedExec(`INSERT INTO events (event_id, event_type, event_entity, event_version, event_timestamp, event_payload_type, event_payload) VALUES (:event_id, :event_type, :event_entity, :event_version, :event_timestamp, :event_payload_type, :event_payload)`, newEvents)

	if eventLatestVersion > 0 {
		_, err = tx.NamedExec(`UPDATE events SET event_version=:version, event_timestamp=:timestamp, event_payload=:payload WHERE event_id=:id AND event_type=:type`, map[string]interface{}{
			"id":        eventId,
			"type":      string(eventTypes.LatestVersion),
			"version":   newEvents[len(newEvents)-1].Version,
			"timestamp": newEvents[len(newEvents)-1].Timestamp,
			"payload":   newEvents[len(newEvents)-1].Payload,
		})

		if err != nil {
			return errors.New(fmt.Sprintf("Failed to update latest version because %s\n", err))
		}
	}

	if err != nil {
		return errors.New(fmt.Sprintf("Failed to insert event because %s\n", err))
	}

	return nil
}

func (m MariaDBEventRepository) Latest(ctx context.Context, eventId string) (*api.Event, error) {
	event := &mariadbEvent{}

	err := m.db.Get(event, "SELECT * FROM events WHERE event_id=? AND event_version=?", eventId, eventTypes.LatestVersion)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get event because %s\n", err))
	}

	eventTimestamp, err := time.Parse(time.RFC3339, event.Timestamp)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to parse event timestamp because %s\n", err))
	}

	return &api.Event{
		EventId:        event.ID,
		EventType:      event.Type,
		EventEntity:    event.Type,
		EventVersion:   event.Version,
		EventTimestamp: timestamppb.New(eventTimestamp),
		EventPayload: &anypb.Any{
			TypeUrl: event.Type,
			Value:   event.Payload,
		},
	}, nil
}

func (m MariaDBEventRepository) getEventLatestVersion(ctx context.Context, eventId string) (int64, error) {
	var eventLatestVersion int64 = 0
	err := m.db.Get(&eventLatestVersion, "SELECT event_version FROM events WHERE event_id=? AND event_type=?", eventId, eventTypes.LatestVersion)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return eventLatestVersion, nil
		default:
			return eventLatestVersion, errors.New(fmt.Sprintf("Failed to get event because %s\n", err))

		}
	}
	return eventLatestVersion, nil
}

func (m MariaDBEventRepository) finishTransaction(err error, tx *sqlx.Tx) error {
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return errors.New(fmt.Sprintf("Error: %s\nRollbackError:%s\n", err, rollbackErr))
		}

		return err
	}
	commitErr := tx.Commit()
	if commitErr != nil {
		return errors.New(fmt.Sprintf("Failed to commit tx because %s\n", commitErr))
	}

	return nil
}

func NewMariaDBConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:123root@tcp(db)/event_store")

	if err != nil {
		return nil, err
	}

	return db, nil
}

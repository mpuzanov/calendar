package grpcserver

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//EventProtoMsgToEvent переконвертируем событие protoMsg в структуру golang
func EventProtoMsgToEvent(in *Event) (*model.Event, error) {

	id, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, err
	}
	startTime, err := ptypes.Timestamp(in.GetStartTime())
	if err != nil {
		return nil, err
	}
	endTime, err := ptypes.Timestamp(in.GetEndTime())
	if err != nil {
		return nil, err
	}
	reminderBefore, err := ptypes.Duration(in.GetReminderBefore())
	if err != nil {
		return nil, err
	}

	event := model.Event{
		ID:             id,
		Header:         in.GetHeader(),
		Text:           in.GetText(),
		StartTime:      startTime,
		EndTime:        endTime,
		UserName:       in.GetUser(),
		ReminderBefore: reminderBefore,
	}
	return &event, nil
}

//EventToEventProtoMsg переконвертируем событие в структуре golang в protoMsg
func EventToEventProtoMsg(in *model.Event) (*Event, error) {

	startTime, err := ptypes.TimestampProto(in.StartTime)
	if err != nil {
		return nil, err
	}
	endTime, err := ptypes.TimestampProto(in.EndTime)
	if err != nil {
		return nil, err
	}

	eventProto := Event{
		Id:             in.ID.String(),
		Header:         in.Header,
		Text:           in.Text,
		StartTime:      startTime,
		EndTime:        endTime,
		User:           in.UserName,
		ReminderBefore: ptypes.DurationProto(in.ReminderBefore),
	}
	return &eventProto, nil
}

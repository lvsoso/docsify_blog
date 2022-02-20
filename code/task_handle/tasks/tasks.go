package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// A list of queues.
const (
	QueueCommand = "command"
	QueueBundle  = "bundle"
	QueueMessage = "message"
	QueueLow     = "low"
)

// A list of task types.
const (
	TypeEmailDelivery = "email:deliver"
	TypeImageResize   = "image:resize"
	TypeBundleSplit   = "bundle:split"
	TypeBundlePack    = "bundle:pack"
)

type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
}

type ImageResizePayload struct {
	SourceURL string
}

type WorkloadPayload struct {
	Count int
}

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//----------------------------------------------

func NewEmailDeliveryTask(userID int, tmplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{UserID: userID, TemplateID: tmplID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func NewImageResizeTask(src string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageResizePayload{SourceURL: src})
	if err != nil {
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.NewTask(TypeImageResize, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}

func NewTestWorkloadTask(count int) (*asynq.Task, error) {
	payload, err := json.Marshal(WorkloadPayload{Count: count})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeBundlePack, payload, asynq.MaxRetry(-1), asynq.Queue(QueueBundle), asynq.Timeout(3*time.Minute)), nil
}

//---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.
// Note that it satisfies the asynq.HandlerFunc interface.
//
// Handler doesn't need to be a function. You can define a type
// that satisfies asynq.Handler interface. See examples below.
//---------------------------------------------------------------

func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Sending Email to User: user_id=%d, template_id=%s", p.UserID, p.TemplateID)
	// Email delivery code ...
	return nil
}

// ImageProcessor implements asynq.Handler interface.
type ImageProcessor struct {
	// ... fields for struct
}

func (processor *ImageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p ImageResizePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Resizing image: src=%s", p.SourceURL)
	// Image resizing code ...
	return nil
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}

func HandleTestWorkloadTask(ctx context.Context, t *asynq.Task) error {
	c := make(chan error, 1)
	isCanceled := make(chan int64, 1)
	go func() {
		c <- handleTestWorkloadTask(isCanceled, t)
	}()
	select {
	case <-ctx.Done():
		// cancelation signal received, abandon this work.
		isCanceled <- 1
		log.Println("task canceled")
		return ctx.Err()
	case res := <-c:
		return res
	}
}

func handleTestWorkloadTask(isCanceled <-chan int64, t *asynq.Task) error {
	var p WorkloadPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("TestWorkload: count=%d", p.Count)

	for i := 0; i < p.Count; i++ {
		if len(isCanceled) >0  {
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Printf("count %d \n", i)
	}
	fmt.Println("task end")
	return nil
}

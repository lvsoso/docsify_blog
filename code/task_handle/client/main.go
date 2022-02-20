package main

import (
	"flag"
	"log"
	"time"

	"task_handle/taskmanager"
	"task_handle/tasks"

	"github.com/davecgh/go-spew/spew"
	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

var taskType int64
var count int
var queue string
var inspect bool
var taskId string
var op string

func init() {
	flag.Int64Var(&taskType, "taskType", 0, "0 : test workload, 1 : email, 2 : future task , 3 : handle image")
	flag.IntVar(&count, "count", 1, "count, test wrok load count")
	flag.BoolVar(&inspect, "inspect", false, "inspect task")
	flag.StringVar(&queue, "queue", tasks.QueueBundle, "task queue name")
	flag.StringVar(&taskId, "taskId", "", "task id")
	flag.StringVar(&op, "op", "", "task op, q: query info , c: cancle task, d: delete task")
}

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	tm := taskmanager.NewTaskManager(redisAddr)

	flag.Parse()
	if inspect {
		if len(taskId) == 0 {
			log.Println("taskId empty")
		}

		switch op {
		case "q":
			log.Println("query task info")
			taskInfo, err := tm.GetTaskInfo(queue, taskId)
			if err != nil {
				log.Println(err.Error())
				return
			}
			spew.Printf("%+v\n", taskInfo)
		case "c":
			log.Println("cancel task ")
			err := tm.CancelProcessing(taskId)
			if err != nil {
				log.Println(err.Error())
				return
			}
			log.Println("cancel task signal had send")
		case "d":
			log.Println("delete task ")
			err := tm.DeleteTask(queue, taskId)
			if err != nil {
				log.Println(err.Error())
				return
			}
			log.Println("delete task succed")
		default:
			log.Println("unknown task op")
		}
	} else {
		switch taskType {
		case 0:
			task, err := tasks.NewTestWorkloadTask(count)
			if err != nil {
				log.Fatalf("could not create task: %v", err)
			}
			info, err := client.Enqueue(task)
			if err != nil {
				log.Fatalf("could not enqueue task: %v", err)
			}
			log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
		case 1:
			// ------------------------------------------------------
			// Example 1: Enqueue task to be processed immediately.
			//            Use (*Client).Enqueue method.
			// ------------------------------------------------------

			task, err := tasks.NewEmailDeliveryTask(42, "some:template:id")
			if err != nil {
				log.Fatalf("could not create task: %v", err)
			}
			info, err := client.Enqueue(task)
			if err != nil {
				log.Fatalf("could not enqueue task: %v", err)
			}
			log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
		case 2:
			// ------------------------------------------------------------
			// Example 2: Schedule task to be processed in the future.
			//            Use ProcessIn or ProcessAt option.
			// ------------------------------------------------------------

			task, err := tasks.NewEmailDeliveryTask(42, "some:template:id")
			if err != nil {
				log.Fatalf("could not create task: %v", err)
			}

			info, err := client.Enqueue(task, asynq.ProcessIn(24*time.Hour))
			if err != nil {
				log.Fatalf("could not schedule task: %v", err)
			}
			log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
		case 3:
			// ----------------------------------------------------------------------------
			// Example 3: Set other options to tune task processing behavior.
			//            Options include MaxRetry, Queue, Timeout, Deadline, Unique etc.
			// ----------------------------------------------------------------------------

			task, err := tasks.NewImageResizeTask("https://example.com/myassets/image.jpg")
			if err != nil {
				log.Fatalf("could not create task: %v", err)
			}
			info, err := client.Enqueue(task, asynq.MaxRetry(10), asynq.Timeout(3*time.Minute))
			if err != nil {
				log.Fatalf("could not enqueue task: %v", err)
			}
			log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
		default:
			log.Fatalf("unknown task type %d", taskType)
		}
	}
}

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"

	"github.com/gin-gonic/gin"
)

type DownloadController struct {
	Controller

	jobUrl     string
	queuePath  string
	taskClient *cloudtasks.Client
}

type DownloadControllerOptions struct {
	ProjectID  string
	LocationID string
	QueueID    string
}

func NewDownloadController(
	dOpt DownloadControllerOptions, opt ControllerOptions,
) (*DownloadController, error) {

	ctrl := Controller{
		opt.ErrorReportCallback,
	}

	jobUrl := fmt.Sprintf("%s/download/url", opt.Host)

	ctx := context.Background()
	taskClient, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("cloudtasks.NewClient: %v", err)
	}

	queuePath := fmt.Sprintf(
		"projects/%s/locations/%s/queues/%s",
		dOpt.ProjectID, dOpt.LocationID, dOpt.QueueID,
	)

	return &DownloadController{
		ctrl,

		jobUrl,
		queuePath,
		taskClient,
	}, nil
}

// @Description Param to send to the download job
type downloadJobParam struct {
	Url  string `json:"url"` // the youtube url to use for download
	Meta struct {
		Artist string `json:"artist"` // the artist
		Album  string `json:"album"`  // the album
		Track  string `json:"track"`  // the music track
	} `json:"meta"` // metadata infos to format the file
}

// @Summary     Download and save a new music file
// @Description Execute the Youtube-DL job using Cloud Task
// @Accept      json
// @Produces    json
// @Param       q   body     true "Main user query"
// @Success     200 {object} pb.Results
// @Router      /music-researcher/search [get]
func (c *DownloadController) DownloadJob(g *gin.Context) {
	var dParam downloadJobParam
	if err := g.ShouldBind(&dParam); err != nil {
		c.badRequest(fmt.Errorf("gin.Context.ShouldBind: %v", err), g)
		return
	}

	body, err := json.Marshal(&dParam)
	if err != nil {
		c.internalError(fmt.Errorf("json.Marshal: %v", err), g)
		return
	}

	req := &taskspb.CreateTaskRequest{
		Parent: c.queuePath,
		Task: &taskspb.Task{
			MessageType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_POST,
					Url:        c.jobUrl,
					Body:       body,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
				},
			},
		},
	}

	ctx := context.Background()
	createdTask, err := c.taskClient.CreateTask(ctx, req)
	if err != nil {
		c.internalError(fmt.Errorf("cloudtasks.CreateTask: %v", err), g)
		return
	}

	g.JSON(http.StatusOK, &createdTask)
}
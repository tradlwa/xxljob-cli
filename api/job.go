package api

import (
	"github.com/tradlwa/xj/api/urlcodec"
	"strings"
)

type JobResponse struct {
	RecordsFiltered int `json:"recordsFiltered"`
	Data            []struct {
		ID                     int    `json:"id"`
		JobGroup               int    `json:"jobGroup"`
		JobCron                string `json:"jobCron"`
		JobDesc                string `json:"jobDesc"`
		AddTime                string `json:"addTime"`
		UpdateTime             string `json:"updateTime"`
		Author                 string `json:"author"`
		AlarmEmail             string `json:"alarmEmail"`
		ExecutorRouteStrategy  string `json:"executorRouteStrategy"`
		ExecutorHandler        string `json:"executorHandler"`
		ExecutorParam          string `json:"executorParam"`
		ExecutorBlockStrategy  string `json:"executorBlockStrategy"`
		ExecutorTimeout        int    `json:"executorTimeout"`
		ExecutorFailRetryCount int    `json:"executorFailRetryCount"`
		GlueType               string `json:"glueType"`
		GlueSource             string `json:"glueSource"`
		GlueRemark             string `json:"glueRemark"`
		GlueUpdatetime         string `json:"glueUpdatetime"`
		ChildJobID             string `json:"childJobId"`
		TriggerStatus          int    `json:"triggerStatus"`
		TriggerLastTime        int    `json:"triggerLastTime"`
		TriggerNextTime        int    `json:"triggerNextTime"`
	} `json:"data"`
	RecordsTotal int `json:"recordsTotal"`
}

type JobOptions struct {
	Group   int    `url:"jobGroup"`
	Status  int    `url:"triggerStatus"`
	Desc    string `url:"jobDesc"`
	Handler string `url:"executorHandler"`
	Start   int    `url:"start"`
	Length  int    `url:"length"`
}

func NewJobOptions() *JobOptions {
	return &JobOptions{
		Status: -1,
		Start:  0,
		Length: 10,
	}
}

func JobPage(client *Client, opts *JobOptions) (*JobResponse, error) {
	values := urlcodec.StructToValues(opts)
	var response JobResponse
	if err := client.Post("jobinfo/pageList", strings.NewReader(values.Encode()), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

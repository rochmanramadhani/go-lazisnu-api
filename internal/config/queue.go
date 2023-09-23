package config

type Queue struct {
	QueueImageMaxLine int `json:"queue_image_max_line"`
	QueueFileMaxLine  int `json:"queue_file_max_line"`
}

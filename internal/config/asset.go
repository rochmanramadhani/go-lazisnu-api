package config

type Asset struct {
	ImageTempDir   string `json:"image_temp_dir"`
	ImageExtension string `json:"image_extension"`
	ImageQuality   string `json:"image_quality"`
	FileTempDir    string `json:"file_temp_dir"`
	FileExtension  string `json:"file_extension"`
}

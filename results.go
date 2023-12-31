package sabnzbd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type BytesFromGB int

func (b *BytesFromGB) UnmarshalJSON(data []byte) error {
	var gb float64
	var gbStr string
	err := json.Unmarshal(data, &gb)
	if err != nil {
		err = json.Unmarshal(data, &gbStr)
		if err != nil {
			return err
		}
		gb, err = strconv.ParseFloat(gbStr, 32)
		if err != nil {
			return err
		}
	}
	*b = BytesFromGB(gb * float64(GByte))
	return nil
}

func (b BytesFromGB) String() string {
	return bytesToHumanReadable(int64(b))
}

type BytesFromMB int

func (b *BytesFromMB) UnmarshalJSON(data []byte) error {
	var mb float64
	var mbStr string
	err := json.Unmarshal(data, &mb)
	if err != nil {
		err = json.Unmarshal(data, &mbStr)
		if err != nil {
			return err
		}
		mb, err = strconv.ParseFloat(mbStr, 32)
		if err != nil {
			return err
		}
	}
	*b = BytesFromMB(mb * float64(MByte))
	return nil
}

func (b BytesFromMB) String() string {
	return bytesToHumanReadable(int64(b))
}

type BytesFromKB int

func (b *BytesFromKB) UnmarshalJSON(data []byte) error {
	var kb float64
	var kbStr string
	err := json.Unmarshal(data, &kb)
	if err != nil {
		err = json.Unmarshal(data, &kbStr)
		if err != nil {
			return err
		}
		kb, err = strconv.ParseFloat(kbStr, 32)
		if err != nil {
			return err
		}
	}
	*b = BytesFromKB(kb * float64(KByte))
	return nil
}

func (b BytesFromKB) String() string {
	return bytesToHumanReadable(int64(b))
}

type BytesFromB int

func (b *BytesFromB) UnmarshalJSON(data []byte) error {
	var bytes float64
	var bytesStr string
	err := json.Unmarshal(data, &bytes)
	if err != nil {
		err = json.Unmarshal(data, &bytesStr)
		if err != nil {
			return err
		}
		bytes, err = strconv.ParseFloat(bytesStr, 32)
		if err != nil {
			return err
		}
	}
	*b = BytesFromB(bytes)
	return nil
}

func (b BytesFromB) String() string {
	return bytesToHumanReadable(int64(b))
}

type SabDuration time.Duration

func (d SabDuration) String() string {
	return time.Duration(d).String()
}

func (d *SabDuration) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	var h, m, s time.Duration
	if _, err = fmt.Sscanf(str, `%d:%2d:%2d`, &h, &m, &s); err != nil {
		return err
	}
	*d = SabDuration(h*time.Hour + m*time.Minute + s*time.Second)
	return nil
}

type apiError struct {
	ErrorMsg string `json:"error,omitempty"`
}

func (e apiError) Error() string {
	return e.ErrorMsg
}

type versionResponse struct {
	Version string `json:"version"`
}

type authResponse struct {
	Auth string `json:"auth"`
}

type QueueResponse struct {
	Version                string      `json:"version"`
	Paused                 bool        `json:"paused"`
	PauseInt               string      `json:"pause_int"`
	PausedAll              bool        `json:"paused_all"`
	DownloadDiskFreeSpace  BytesFromGB `json:"diskspace1"`
	CompleteDiskFreeSpace  BytesFromGB `json:"diskspace2"`
	Diskspace1Norm         string      `json:"diskspace1_norm"`
	Diskspace2Norm         string      `json:"diskspace2_norm"`
	DownloadDiskTotalSpace BytesFromGB `json:"diskspacetotal1"`
	CompleteDiskTotalSpace BytesFromGB `json:"diskspacetotal2"`
	SpeedLimitPercentage   int         `json:"speedlimit,string"`
	SpeedLimit             BytesFromB  `json:"speedlimit_abs"`
	HaveWarnings           string      `json:"have_warnings"`
	Finishaction           *string     `json:"finishaction"`
	Quota                  string      `json:"quota"`
	HaveQuota              bool        `json:"have_quota"`
	LeftQuota              string      `json:"left_quota"`
	CacheArt               string      `json:"cache_art"`
	CacheSize              string      `json:"cache_size"`
	BytesPerSecond         BytesFromKB `json:"kbpersec"`
	Speed                  string      `json:"speed"`
	BytesLeft              BytesFromMB `json:"mbleft"`
	Bytes                  BytesFromMB `json:"mb"`
	BytesMissing           BytesFromMB
	SizeLeft               string      `json:"sizeleft"`
	Size                   string      `json:"size"`
	NoOfSlotsTotal         int         `json:"noofslots_total"`
	NoOfSlots              int         `json:"noofslots"`
	Start                  int         `json:"start"`
	Limit                  int         `json:"limit"`
	Finish                 int         `json:"finish"`
	Status                 string      `json:"status"`
	TimeLeft               SabDuration `json:"timeleft"`
	Slots                  []QueueSlot `json:"slots"`
	apiError
}

type queueResponse *QueueResponse

func (r *QueueResponse) UnmarshalJSON(data []byte) error {
	var queue struct {
		Queue json.RawMessage `json:"queue"`
	}
	if err := json.Unmarshal(data, &queue); err != nil {
		return err
	}
	err := json.Unmarshal(queue.Queue, queueResponse(r))
	r.BytesMissing = r.Bytes - r.BytesLeft
	return err
}

type QueueSlot struct {
	Index        int         `json:"index"`
	NzoID        string      `json:"nzo_id"`
	UnpackOpts   string      `json:"unpackopts"`
	Priority     string      `json:"priority"`
	Script       string      `json:"script"`
	Filename     string      `json:"filename"`
	Labels       []string    `json:"labels"`
	Password     string      `json:"password"`
	Category     string      `json:"cat"`
	BytesLeft    BytesFromMB `json:"mbleft"`
	Bytes        BytesFromMB `json:"mb"`
	Size         string      `json:"size"`
	SizeLeft     string      `json:"sizeleft"`
	Percentage   string      `json:"percentage"`
	BytesMissing BytesFromMB `json:"mbmissing"`
	DirectUnpack string      `json:"direct_unpack"`
	Status       string      `json:"status"`
	TimeLeft     SabDuration `json:"timeleft"`
	AverageAge   string      `json:"avg_age"`
}

type HistoryResponse struct {
	TotalSize         string        `json:"total_size"`
	MonthSize         string        `json:"month_size"`
	WeekSize          string        `json:"week_size"`
	DaySize           string        `json:"day_size"`
	Slots             []HistorySlot `json:"slots"`
	NoOfSlots         int           `json:"noofslots"`
	Version           string        `json:"version"`
	LastHistoryUpdate int           `json:"last_history_update"`
	apiError
}

type historyResponse *HistoryResponse

func (r *HistoryResponse) UnmarshalJSON(data []byte) error {
	var history struct {
		History json.RawMessage `json:"history"`
	}
	if err := json.Unmarshal(data, &history); err != nil {
		return err
	}
	err := json.Unmarshal(history.History, historyResponse(r))

	// Parsing time/duration fields
	if r.NoOfSlots > 0 {
		for i, slot := range r.Slots {
			slot.Completed = time.Unix(slot.CompletedUnix, 0)
			slot.DownloadDuration = time.Duration(slot.DownloadTime * int64(time.Second))
			slot.PostProcessingDuration = time.Duration(slot.PostProcessingTime * int64(time.Second))
			r.Slots[i] = slot
		}
	}
	return err
}

type HistorySlot struct {
	ID                     int   `json:"id"`
	CompletedUnix          int64 `json:"completed"`
	Completed              time.Time
	Name                   string `json:"name"`
	NZBName                string `json:"nzb_name"`
	Category               string `json:"category"`
	PP                     string `json:"pp"`
	Script                 string `json:"script"`
	Report                 string `json:"report"`
	URL                    string `json:"url"`
	Status                 string `json:"status"`
	NzoID                  string `json:"nzo_id"`
	Storage                string `json:"storage"`
	Path                   string `json:"path"`
	ScriptLog              string `json:"script_log"`
	ScriptLine             string `json:"script_line"`
	DownloadTime           int64  `json:"download_time"`
	DownloadDuration       time.Duration
	PostProcessingTime     int64 `json:"postproc_time"`
	PostProcessingDuration time.Duration
	StageLogs              []HistoryStageLog `json:"stage_log"`
	Downloaded             int64             `json:"downloaded"` // represents downloaded bytes, not time
	Completeness           int               `json:"completeness"`
	FailMessage            string            `json:"fail_message"`
	URLInfo                string            `json:"url_info"`
	Bytes                  int               `json:"bytes"`
	Meta                   string            `json:"meta"`
	Series                 string            `json:"series"`
	MD5Sum                 string            `json:"md5sum"`
	Password               string            `json:"password"`
	ActionLine             string            `json:"action_line"`
	Size                   string            `json:"size"`
	Loaded                 bool              `json:"loaded"`
	Retry                  int
}

type HistoryStageLog struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}

type ServerStatsResponse struct {
	Total   BytesFromB                   `json:"total"`
	Month   BytesFromB                   `json:"month"`
	Week    BytesFromB                   `json:"week"`
	Day     BytesFromB                   `json:"day"`
	Servers map[string]ServerStatsServer `json:"servers"`
}

type ServerStatsServer struct {
	Total           BytesFromB                     `json:"total"`
	Month           BytesFromB                     `json:"month"`
	Week            BytesFromB                     `json:"week"`
	Day             BytesFromB                     `json:"day"`
	Daily           map[ServerStatsDate]BytesFromB `json:"daily"`
	ArticlesTried   map[ServerStatsDate]int        `json:"articles_tried"`
	ArticlesSuccess map[ServerStatsDate]int        `json:"articles_success"`
}

type ServerStatsDate string

type warningsResponse struct {
	Warnings []string `json:"warnings"`
	apiError
}

type categoriesResponse struct {
	Categories []string `json:"categories"`
	apiError
}

type scriptsResponse struct {
	Scripts []string `json:"scripts"`
	apiError
}

type addFileResponse struct {
	NzoIDs []string `json:"nzo_ids"`
	apiError
}

type ItemFilesResponse struct {
	Files []ItemFile `json:"files"`
	apiError
}

type ItemFile struct {
	ID        string      `json:"id"`
	NzfID     string      `json:"nzf_id"`
	Status    string      `json:"status"`
	Filename  string      `json:"filename"`
	Age       string      `json:"age"`
	Bytes     BytesFromB  `json:"bytes"`
	BytesLeft BytesFromMB `json:"mbleft"`
}

type itemFile *ItemFile

func (f *ItemFile) UnmarshalJSON(data []byte) (err error) {
	err = json.Unmarshal(data, itemFile(f))
	if err != nil {
		return err
	}
	if int(f.BytesLeft) > int(f.Bytes) {
		f.BytesLeft = BytesFromMB(f.Bytes)
	}
	return nil
}

func bytesToHumanReadable(i int64) string {
	if i < 1000 {
		return fmt.Sprintf("%d B", i)
	} else if i < 1000*1000 {
		return fmt.Sprintf("%.2f KB", float64(i)/1000)
	} else if i < 1000*1000*1000 {
		return fmt.Sprintf("%.2f MB", float64(i)/(1000*1000))
	} else if i < 1000*1000*1000*1000 {
		return fmt.Sprintf("%.2f GB", float64(i)/(1000*1000*1000))
	} else if i < 1000*1000*1000*1000*1000 {
		return fmt.Sprintf("%.2f TB", float64(i)/(1000*1000*1000*1000))
	} else {
		return fmt.Sprintf("%.2f PB", float64(i)/(1000*1000*1000*1000*1000))
	}
}

package models

type (
	User struct {
		UserId int `json:"user_id"`
	}

	Segment struct {
		Segment string `json:"segment"`
	}

	UserSegment struct {
		UserId   int    `json:"user_id"`
		Segments string `json:"segments"`
	}

	UserSetSegment struct {
		SegmentsSet    []string `json:"segments_set"`
		SegmentsDelete []string `json:"segments_delete"`
		UserId         int      `json:"user_id"`
	}
)

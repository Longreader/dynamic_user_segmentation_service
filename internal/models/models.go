package models

type (
	User struct {
		UserID int `json:"user_id"`
	}

	Segment struct {
		Segment string `json:"segment"`
		Percent int    `json:"percent"`
	}

	UserSegments struct {
		UserID   int      `json:"user_id"`
		Segments []string `json:"segments"`
	}

	UserSetSegment struct {
		SegmentsSet    []string `json:"segments_set"`
		SegmentsDelete []string `json:"segments_delete"`
		UserID         int      `json:"user_id"`
	}
)

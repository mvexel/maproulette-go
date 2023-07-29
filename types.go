package maproulette

import "net/http"

// MapRoulette represents a client for the MapRoulette API.
// To create a new client, use the NewMapRouletteClient function.
type MapRoulette struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

// MapRouletteClientOptions represents options for a MapRoulette client.
type MapRouletteClientOptions struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

// GeoJSON represents a GeoJSON object in MapRoulette.
// TODO: we need a real spatial type here
type GeoJSON struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"` // []float64 for Point, [][]float64 for Polygon
}

// Challenge represents a challenge in MapRoulette.
type Challenge struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Created              string `json:"created"`
	Modified             string `json:"modified"`
	Description          string `json:"description"`
	Deleted              bool   `json:"deleted"`
	InfoLink             string `json:"infoLink"`
	General              General
	Creation             Creation
	Priority             Priority
	Extra                Extra
	Status               int     `json:"status"`
	StatusMessage        string  `json:"statusMessage"`
	LastTaskRefresh      string  `json:"lastTaskRefresh"`
	DataOriginDate       string  `json:"dataOriginDate"`
	Location             GeoJSON `json:"location"`
	Bounding             GeoJSON `json:"bounding"`
	CompletionPercentage int     `json:"completionPercentage"`
	TasksRemaining       int     `json:"tasksRemaining"`
}

// General represents the general section of a challenge with basic information
// about its owner, parent project, instructions, etc.
type General struct {
	Owner           int     `json:"owner"`
	Parent          Project `json:"parent"`
	Instruction     string  `json:"instruction"`
	Difficulty      int     `json:"difficulty"`
	Blurb           string  `json:"blurb"`
	Enabled         bool    `json:"enabled"`
	Featured        bool    `json:"featured"`
	CooperativeType int     `json:"cooperativeType"`
	Popularity      int     `json:"popularity"`
	CheckinComment  string  `json:"checkinComment"`
	CheckinSource   string  `json:"checkinSource"`
	VirtualParents  []int   `json:"virtualParents"`
	RequiresLocal   bool    `json:"requiresLocal"`
}

// Creation represents the creation section of a challenge with information
// about the overpass query or remote GeoJSON link.
type Creation struct {
	OverpassQL         string `json:"overpassQL"`
	RemoteGeoJson      string `json:"remoteGeoJson"`
	OverpassTargetType string `json:"overpassTargetType"`
}

// Priority represents the priority section of a challenge with information
// about the default priority and priority rules.
type Priority struct {
	DefaultPriority    int    `json:"defaultPriority"`
	HighPriorityRule   string `json:"highPriorityRule"`
	MediumPriorityRule string `json:"mediumPriorityRule"`
	LowPriorityRule    string `json:"lowPriorityRule"`
}

// Extra represents the extra section of a challenge with information
// which are optional to set when creating a challenge.
type Extra struct {
	DefaultZoom          int      `json:"defaultZoom"`
	MinZoom              int      `json:"minZoom"`
	MaxZoom              int      `json:"maxZoom"`
	DefaultBasemap       int      `json:"defaultBasemap"`
	DefaultBasemapId     string   `json:"defaultBasemapId"`
	CustomBasemap        string   `json:"customBasemap"`
	UpdateTasks          bool     `json:"updateTasks"`
	ExportableProperties string   `json:"exportableProperties"`
	OsmIdProperty        string   `json:"osmIdProperty"`
	PreferredTags        string   `json:"preferredTags"`
	PreferredReviewTags  string   `json:"preferredReviewTags"`
	LimitTags            bool     `json:"limitTags"`
	LimitReviewTags      bool     `json:"limitReviewTags"`
	TaskStyles           string   `json:"taskStyles"`
	TaskBundleIdProperty string   `json:"taskBundleIdProperty"`
	IsArchived           bool     `json:"isArchived"`
	ReviewSetting        int      `json:"reviewSetting"`
	SystemArchivedAt     int      `json:"systemArchivedAt"`
	Presets              []string `json:"presets"`
}

// Project represents a project in MapRoulette.
//
// This struct is used by various methods in the Client to receive and send
// project data.
type Project struct {
	ID          int     `json:"id"`
	Owner       int     `json:"owner"`
	Name        string  `json:"name"`
	Created     string  `json:"created"`
	Modified    string  `json:"modified"`
	Description string  `json:"description"`
	Grants      []Grant `json:"grants"`
	Enabled     bool    `json:"enabled"`
	DisplayName string  `json:"displayName"`
	Deleted     bool    `json:"deleted"`
	Featured    bool    `json:"featured"`
}

// Grant represents a grant in MapRoulette.
//
// This struct is used by various methods in the Client to receive and send
// grant data.
type Grant struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Grantee Grantee `json:"grantee"`
	Role    int     `json:"role"`
	Target  Target  `json:"target"`
}

// Grantee represents a grantee in MapRoulette.
type Grantee struct {
	GranteeType GranteeType `json:"granteeType"`
	GranteeID   int         `json:"granteeId"`
}

// GranteeType represents a grantee type in MapRoulette.
type GranteeType struct {
	ID int `json:"id"`
}

// Target represents a target in MapRoulette.
type Target struct {
	ObjectType ObjectType `json:"objectType"`
	ObjectID   int        `json:"objectId"`
}

// ObjectType represents an object type in MapRoulette.
type ObjectType struct {
	ID int `json:"id"`
}

// Task represents a task in MapRoulette.
type Task struct {
	ID                  int64   `json:"id"`
	Name                string  `json:"name"`
	Created             string  `json:"created"`
	Modified            string  `json:"modified"`
	Parent              int64   `json:"parent"`
	Instruction         string  `json:"instruction"`
	Location            GeoJSON `json:"location"`
	Geometries          GeoJSON `json:"geometries"`
	CooperativeWork     string  `json:"cooperativeWork"`
	Status              int     `json:"status"`
	MappedOn            string  `json:"mappedOn"`
	CompletedTimeSpent  int64   `json:"completedTimeSpent"`
	CompletedBy         int64   `json:"completedBy"`
	Review              Review  `json:"review"`
	Priority            int     `json:"priority"`
	ChangesetId         int64   `json:"changesetId"`
	CompletionResponses string  `json:"completionResponses"`
	BundleId            int64   `json:"bundleId"`
	IsBundlePrimary     bool    `json:"isBundlePrimary"`
	MapillaryImages     []Image `json:"mapillaryImages"`
	ErrorTags           string  `json:"errorTags"`
}

// Review represents review data for a Task in MapRoulette.
type Review struct {
	ReviewStatus        int    `json:"reviewStatus"`
	ReviewRequestedBy   int    `json:"reviewRequestedBy"`
	ReviewedBy          int    `json:"reviewedBy"`
	ReviewedAt          string `json:"reviewedAt"`
	MetaReviewedBy      int    `json:"metaReviewedBy"`
	MetaReviewStatus    int    `json:"metaReviewStatus"`
	MetaReviewedAt      string `json:"metaReviewedAt"`
	ReviewStartedAt     string `json:"reviewStartedAt"`
	ReviewClaimedBy     int    `json:"reviewClaimedBy"`
	ReviewClaimedAt     string `json:"reviewClaimedAt"`
	AdditionalReviewers []int  `json:"additionalReviewers"`
}

// Image represents a set of URLs pointing to Mapillary images for a Task.
type Image struct {
	Key     string  `json:"key"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	URL320  string  `json:"url_320"`
	URL640  string  `json:"url_640"`
	URL1024 string  `json:"url_1024"`
	URL2048 string  `json:"url_2048"`
}

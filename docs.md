# maproulette
--
    import "github.com/mvexel/maproulette-go"

Package maproulette provides a Go client for the MapRoulette API.

The starting point for most applications will be the NewClient function, which
creates a new API client instance. The client provides methods to interact with
MapRoulette API resources, such as challenges.

## Usage

```go
var Version string
```

#### type Challenge

```go
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
```

Challenge represents a challenge in MapRoulette.

This struct is used by various methods in the Client to receive and send
challenge data.

#### type Creation

```go
type Creation struct {
	OverpassQL         string `json:"overpassQL"`
	RemoteGeoJson      string `json:"remoteGeoJson"`
	OverpassTargetType string `json:"overpassTargetType"`
}
```

Creation represents the creation section of a challenge with information about
the overpass query or remote GeoJSON link.

#### type Extra

```go
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
```

Extra represents the extra section of a challenge with information which are
optional to set when creating a challenge.

#### type General

```go
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
```

General represents the general section of a challenge with basic information
about its owner, parent project, instructions, etc.

#### type GeoJSON

```go
type GeoJSON struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"` // []float64 for Point, [][]float64 for Polygon
}
```

TODO - use real GeoJSON types from twpayne/go-geojson?

#### type Grant

```go
type Grant struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Grantee Grantee `json:"grantee"`
	Role    int     `json:"role"`
	Target  Target  `json:"target"`
}
```

Grant represents a grant in MapRoulette.

This struct is used by various methods in the Client to receive and send grant
data.

#### type Grantee

```go
type Grantee struct {
	GranteeType GranteeType `json:"granteeType"`
	GranteeID   int         `json:"granteeId"`
}
```

Grantee represents a grantee in MapRoulette.

This struct is used by various methods in the Client to receive and send grantee
data.

#### type GranteeType

```go
type GranteeType struct {
	ID int `json:"id"`
}
```

GranteeType represents a grantee type in MapRoulette.

This struct is used by various methods in the Client to receive and send grantee
type data.

#### type MapRoulette

```go
type MapRoulette struct {
	APIKey string
}
```

MapRoulette represents a client for the MapRoulette API.

To create a new client, use the NewMapRouletteClient function.

#### func  NewMapRouletteClient

```go
func NewMapRouletteClient(apiKey string) *MapRoulette
```
NewMapRouletteClient creates a new MapRoulette API client.

The client communicates with the MapRoulette API at the specified baseURL, and
uses the specified API key for authentication.

#### func (*MapRoulette) GetChallenge

```go
func (mr *MapRoulette) GetChallenge(id int) (Challenge, error)
```
GetChallenge returns a challenge from the MapRoulette API.

The id parameter specifies the ID of the challenge to return.

#### func (*MapRoulette) GetChallenges

```go
func (mr *MapRoulette) GetChallenges(limit int) ([]Challenge, error)
```

#### type ObjectType

```go
type ObjectType struct {
	ID int `json:"id"`
}
```

ObjectType represents an object type in MapRoulette.

This struct is used by various methods in the Client to receive and send object
type data.

#### type Priority

```go
type Priority struct {
	DefaultPriority    int    `json:"defaultPriority"`
	HighPriorityRule   string `json:"highPriorityRule"`
	MediumPriorityRule string `json:"mediumPriorityRule"`
	LowPriorityRule    string `json:"lowPriorityRule"`
}
```

Priority represents the priority section of a challenge with information about
the default priority and priority rules.

#### type Project

```go
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
```

Project represents a project in MapRoulette.

This struct is used by various methods in the Client to receive and send project
data.

#### type Target

```go
type Target struct {
	ObjectType ObjectType `json:"objectType"`
	ObjectID   int        `json:"objectId"`
}
```

Target represents a target in MapRoulette.

This struct is used by various methods in the Client to receive and send target
data.

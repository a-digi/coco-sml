package performance

import "time"

// PerformanceTrack hält Informationen über den Start- und Endzeitpunkt sowie eine ID für einen Prozess.
type PerformanceTrack struct {
	StartedAt   time.Time `json:"started_at"`
	EndedAt     time.Time `json:"ended_at"`
	Id          string    `json:"id"`
	DurationNS  float64   `json:"duration_ns"` // Dauer in Nanosekunden als float64
}

// CreatePerformanceTrack erstellt einen neuen PerformanceTrack mit Startzeitpunkt und ID.
func CreatePerformanceTrack(id string) *PerformanceTrack {
	return &PerformanceTrack{
		StartedAt: time.Now(),
		Id:        id,
	}
}

// End markiert das Ende des Prozesses.
func (pt *PerformanceTrack) End() {
	pt.EndedAt = time.Now()
	pt.DurationNS = float64(pt.EndedAt.Sub(pt.StartedAt).Nanoseconds())
}

// Milliseconds gibt die Dauer in Millisekunden zurück.
func (pt *PerformanceTrack) Milliseconds() float64 {
	return float64(pt.DurationNS) / 1e6
}

// PerformanceTracker manages a collection of PerformanceTrack and allows lookup by ID.
type PerformanceTracker struct {
	trackPoints map[string]*PerformanceTrack `json:"trackPoints"`
}

// NewPerformanceTracker creates a new PerformanceTracker.
func NewPerformanceTracker() *PerformanceTracker {
	return &PerformanceTracker{
		trackPoints: make(map[string]*PerformanceTrack),
	}
}

// AddTrackPoint creates a new PerformanceTrack with the given ID and adds it to the tracker. Returns the created PerformanceTrack.
func (pt *PerformanceTracker) AddTrackPoint(id string) *PerformanceTrack {
	track := &PerformanceTrack{
		StartedAt: time.Now(),
		Id:        id,
	}
	pt.trackPoints[id] = track
	return track
}

// GetTrackPointByID returns the PerformanceTrack with the given ID, or nil if not found.
func (pt *PerformanceTracker) GetTrackPointByID(id string) *PerformanceTrack {
	track, ok := pt.trackPoints[id]
	if !ok {
		return nil
	}
	return track
}

// EndTrackPoint ends the tracking for the given ID by calling End() on the corresponding PerformanceTrack.
// Returns true if the track point was found and ended, false otherwise.
func (pt *PerformanceTracker) EndTrackPoint(id string) bool {
	track, ok := pt.trackPoints[id]
	if !ok {
		return false
	}

	track.End()
	return true
}

package Sources

type Source interface {
	GetVideos() (string, error)
}
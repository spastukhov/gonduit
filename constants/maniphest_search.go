package constants

// ManiphestSearchOrder is the order in which search results should be ordered.
type ManiphestSearchOrder string

const (
	ManiphestSearchOrderPriority ManiphestSearchOrder = "priority"
	ManiphestSearchOrderUpdated  ManiphestSearchOrder = "updated"
	ManiphestSearchOrderOutdated ManiphestSearchOrder = "outdated"

	ManiphestSearchOrderNewest ManiphestSearchOrder = "newest"
	ManiphestSearchOrderOldest ManiphestSearchOrder = "oldest"
	ManiphestSearchOrderClosed ManiphestSearchOrder = "closed"

	ManiphestSearchOrderTitle     ManiphestSearchOrder = "title"
	ManiphestSearchOrderRelevance ManiphestSearchOrder = "relevance"
)
